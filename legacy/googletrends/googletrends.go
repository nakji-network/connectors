// Poll Google Trends for data.
// References:
// - https://github.com/groovili/gogtrends/tree/694001cbd88873cee94b9cad76d135860790155e
// - https://towardsdatascience.com/reconstruct-google-trends-daily-data-for-extended-period-75b6ca1d3420
// - https://github.com/qztseng/google-trends-daily/blob/master/gtrend.py
package googletrends

import (
	"context"
	reflect "reflect"
	"strconv"
	"time"

	"blep.ai/data/database"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/golang/protobuf/ptypes"
	"github.com/groovili/gogtrends"
	"github.com/rs/zerolog/log"
)

func New(conf Config) GoogleTrends {
	return GoogleTrends{
		Producer:   conf.KP,
		BaseURL:    conf.BaseURL,
		Keywords:   conf.Keywords,
		NS:         conf.NS,
		DB:         conf.DB,
		Table:      conf.Table,
		TrendTopic: conf.TrendTopic,
		Sleep:      conf.Sleep,
		Interval:   conf.Interval,
		Backoff:    conf.Backoff,
	}
}

const (
	//Increases size of db query, but removes bug
	//where old query of 10 highest values in records had no overlap
	//also allows us to handle the sporadic 0 value returned by API

	//historyLimit = 10
	historyLimit = 240
)

// Start polling Google Trends.
// How it works:
// - Get last 4 hours from API.
// - Find records with overlapping timestamps from DB.
// - Scaling factor = old value / new value
func (t *GoogleTrends) Start() {
	// Timed fetch (current refresh period is 1 hour)
	for ; true; <-time.Tick(t.Interval) {
		// Fetch by keyword
		for _, keyword := range t.Keywords {

			//If reconstructing historical data outside of records due to
			//adding new keyword or server downtime we can use the following query:
			//results := fetchGoogleInterest(context.Background(), keyword, "2021-06-05T01 2021-06-11T05")
			//Change 2021-06-05T01 2021-06-05T05 to approriate 4 hour interval (will still provide per minute data)

			//API results is not guaranteed to return a value every minute (i.e. "glassnode")
			//On the google trends site visualiation, timestamps without value, have default value 0
			//This doesn't appear to be an API error as subsequent calls to the same interval doesn't return
			//values for the timestamps that are 0 on the visualization site....
			results := fetchGoogleInterest(context.Background(), keyword, "now 4-H", t.Backoff)

			if len(results) < 2 {
				continue
			}

			start := results[0].Ts
			end := results[len(results)-1].Ts

			// Get last batch that was fetched for normalization
			// SelectLastFetched, returns sorted list of records in descending order
			// by normalized value (i.e. max normalized value timestamp is first)
			lastFetched, err := t.DB.SelectLastFetched(database.QueryOptions{
				Table:     t.Table,
				Namespace: t.NS,
				Subject:   keyword,
				ToType:    reflect.TypeOf(TrendScore{}),
			}, start.AsTime(), end.AsTime(), historyLimit)
			if err != nil {
				log.Error().Err(err).Msg("db.SelectLast errorstring")
				return
			}

			// Normalize new result values based on largest stored value
			// If not found or no existing records, do not normalize
			multiplier := 1.0
			lastRecords := lastFetched.([]TrendScore)

			multiplier = findHighestMatch(multiplier, lastRecords, results)

			// Write normalized struct to kafka.
			// All are written, including duplicate Timestamps so normalization can be modified later.
			for _, r := range results {
				t.normalizeValues(r, multiplier, keyword, lastRecords)
			}

			// Commit Kafka Transaction
			// Requires timescaledb to be running to consume transactions
			// and throw into the sql database (make start-timescaledb)
			err = t.CommitTransaction(nil)
			if err != nil {
				log.Error().Err(err).Msg("Processor: Failed to commit transaction")

				err = t.AbortTransaction(nil)
				if err != nil {
					log.Fatal().Err(err).Msg("")
				}
			}

			// Start a new transaction
			err = t.BeginTransaction()
			if err != nil {
				log.Fatal().Err(err).Msg("")
			}

			log.Info().
				Str("keyword", keyword).
				Msg("kafka write + commit")

			// Prevent hitting google trends api rate limiter (undocumented, do not know rate limit)
			time.Sleep(t.Sleep)
		}
	}
}

func findHighestMatch(multiplier float64, lastRecords []TrendScore, results []*TrendScore) float64 {
	for i := range lastRecords {
		for _, r := range results {
			if r.Ts.AsTime() == lastRecords[i].Ts.AsTime() {

				// Handle API returning 0 for actually none 0 value and potentially causing divide by 0 error
				// Future timestamps in results would be larger, so match next record
				if r.Value == 0 {
					break
				}

				return lastRecords[i].ValueNormalized / r.Value
			}
		}
	}
	return multiplier
}

func (t *GoogleTrends) normalizeValues(r *TrendScore, multiplier float64, keyword string, lastRecords []TrendScore) {

	//Multiplier can be > 1 so normalized value can be >100
	r.ValueNormalized = r.Value * multiplier

	// Handle API returning 0 when it should return a number
	if r.Value == 0 {
		// Hacky method where we just use previously stored normalized value
		// Won't actually be completely normalized, but should be more normalized then 0
		// If previously stored normalized value was 0 then it must be because it was originally 0...
		for i := range lastRecords {
			if r.Ts.AsTime() == lastRecords[i].Ts.AsTime() {
				r.ValueNormalized = lastRecords[i].ValueNormalized
				break
			}
		}
	}

	err := t.WriteKafkaMessages(
		t.TrendTopic,
		kafkautils.NewKey(t.NS, keyword).Bytes(),
		r)

	if err != nil {
		log.Error().Err(err).Msg("write kafka error from proto marshal")
	}
}

// Convert row timestamp value from pgtype.Timestamptz to integer
func rowTimestampAsUnix(r *gogtrends.Timeline) int64 {
	u, err := strconv.ParseInt(r.Time, 0, 64)
	if err != nil {
		panic(err)
	}
	return u
}

// https://github.com/groovili/gogtrends/blob/694001cbd88873cee94b9cad76d135860790155e/vars.go//L195
func fetchGoogleInterest(ctx context.Context, keyword, timeframe string, backoff time.Duration) []*TrendScore {
	maxBackOff := 10 * time.Minute
	tsFetched := ptypes.TimestampNow()

	// Create widget to search
	explore, err := gogtrends.Explore(ctx,
		&gogtrends.ExploreRequest{
			ComparisonItems: []*gogtrends.ComparisonItem{
				{
					Keyword: keyword,
					Geo:     "", // "" is worldwide
					Time:    timeframe,
				},
			},
			Category: 0,
			Property: "",
		}, "EN")

	if err != nil {
		log.Error().Err(err).Msg("fetch explore error")
		if backoff > maxBackOff {
			log.Info().Str("Keyword", keyword).Msg("Hit maximum backoff, skipping")
			return []*TrendScore{}
		} else {
			log.Info().Str("Keyword", keyword).Msg("Retrying with increased backoff")
			return fetchGoogleInterest(ctx, keyword, timeframe, backoff*2)
		}
	}

	// Fetch data
	data, err := gogtrends.InterestOverTime(ctx, explore[0], "EN")
	if err != nil {
		log.Error().Err(err).Msg("fetch google interest error")
		if backoff > maxBackOff {
			log.Info().Str("Keyword", keyword).Msg("Hit maximum backoff, skipping")
			return []*TrendScore{}
		} else {
			log.Info().Str("Keyword", keyword).Msg("Retrying with increased backoff")
			return fetchGoogleInterest(ctx, keyword, timeframe, backoff*2)
		}
	}

	// Unmarshal gogtrends.Timeline to TrendScore
	return UnmarshalGogtrendsTimeline(data, tsFetched)
}
