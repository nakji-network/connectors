// Binance open interest is not served over websockets, so it must short poll.
// This is a good template for all future connectors that need to implement
// short polling.
package main

import (
	"fmt"
	"net/url"
	"sync"
	"time"

	"blep.ai/data/common"
	"blep.ai/data/config"
	"blep.ai/data/connectors/source"
	"blep.ai/data/database"
	"blep.ai/data/request"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
)

var (
	baseURL  = "https://fapi.binance.com/futures/data/openInterestHist?"
	exchange = "binance"
	symbols  = []string{"BTCUSDT", "ETHUSDT"}
	period   = "5m"
	table    = "openinterest-1m0s"

	transactionalID = "binance-openinterest-source"
	OITOPIC         = ".fct.nakji.common.0_0_0.market_openinterest"
)

func main() {
	conf := config.GetConfig()

	// Load Topic registry
	kafkautils.TopicTypeRegistry.Load(source.TopicTypes)
	oiTopic := kafkautils.MustParseTopic(OITOPIC, conf.GetString("kafka.env"))
	kp, err := kafkautils.NewProducer(conf.GetString("kafka.url"), transactionalID)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	kp.EnableTransactions()

	// Connect to database
	db, err := database.New(conf.GetString("timescaledb.connection"))
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	defer db.Close()

	// Backfill all symbols
	var wg sync.WaitGroup
	for _, symbol := range symbols {
		go backfill(db, symbol, period, &wg)
	}
	wg.Wait()

	// Timed fetch
	for _, s := range symbols {
		go func(symbol string) {
			params := url.Values{}
			params.Add("limit", "5")
			params.Add("symbol", symbol)
			params.Add("period", period)

			tfOptions := request.TimedFetchOptions{
				Interval: 5 * 60 * 1000, // 5 minutes
				Margin:   550,
				Delay:    4 * 60 * 1000, // 4 minutes
				Decay:    6,
			}

			request.TimedFetch(baseURL, params, tfOptions, func(val *gjson.Result, err error) {
				if err != nil {
					log.Fatal().Err(err).Msg("")
				}

				msInt := val.Get("timestamp").Int()
				ts := common.UnixToTimestampPb(msInt)

				key := kafkautils.NewKey(exchange, symbol)
				err = kp.WriteAndCommit(oiTopic, key.Bytes(), &source.OpenInterest{
					Ts:                ts,
					OpenInterest:      val.Get("sumOpenInterest").Float(),
					OpenInterestValue: val.Get("sumOpenInterestValue").Float(),
				})
				if err != nil {
					log.Error().Err(err)
				}
			})
		}(s)
	}

	select {}
}

// fetchAll automatically crawls entire dataset from `from` to `to`, making
// multiple requests as necessary. This has several variables that are specific
// to Binance; hence refactoring may not be trivial.
func fetchAll(symbol, period string, from, to time.Time, callback func(*source.OpenInterest)) error {
	fromms := from.Unix() * 1000
	toms := to.Unix() * 1000

	for {
		// Create request parameters
		params := url.Values{}
		params.Add("limit", "500")
		params.Add("symbol", symbol)
		params.Add("period", period)
		params.Add("endTime", fmt.Sprintf("%v", toms))

		// Fetch open interest data from exchange
		resJSON, err := request.GetJSON(baseURL, params)
		if err != nil {
			return err
		}
		res := resJSON.Array()

		if len(res) == 0 {
			break
		}

		// Insert each result into database
		// TODO: ignore duplicate inserts
		for _, val := range res {
			msInt := val.Get("timestamp").Int()

			if msInt >= fromms {
				Row := source.OpenInterest{
					Ts:                common.UnixToTimestampPb(msInt),
					OpenInterest:      val.Get("sumOpenInterest").Float(),
					OpenInterestValue: val.Get("sumOpenInterestValue").Float(),
				}
				callback(&Row)
			}
		}

		startms := res[0].Get("timestamp").Int()
		if startms <= fromms {
			break
		}

		toms = startms
	}
	return nil
}

// Backfill data between now and last timestamp found in db
func backfill(db *database.Database, symbol, period string, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	startTime, err := db.SelectLastTimestamp(database.QueryOptions{
		Table:     table,
		Namespace: exchange,
		Subject:   symbol})
	if err == nil {
		// Increment startTime by 5 min for new data
		startTime = startTime.Add(5 * time.Minute)
	} else {
		// Table has no rows: fetch data from the last month
		// Note: Binance only serves 1 month of data.
		startTime = time.Now().Add(-24 * 32 * time.Hour)
	}

	if startTime.Before(time.Now()) {
		err = fetchAll(symbol, period, startTime, time.Now(), func(Row *source.OpenInterest) {
			data := common.OHLC{
				Ts:    Row.Ts,
				Open:  Row.OpenInterest,
				High:  Row.OpenInterest,
				Low:   Row.OpenInterest,
				Close: Row.OpenInterest,
				//ValueOpen:  Row.OpenInterestValue,
				//ValueHigh:  Row.OpenInterestValue,
				//ValueLow:   Row.OpenInterestValue,
				//ValueClose: Row.OpenInterestValue,
			}
			db.InsertProto2(table, &data, kafkautils.NewKey(exchange, symbol).Tuple()...)
		})
		if err != nil {
			log.Fatal().Err(err).
				Str("symbol", symbol).
				Str("period", period).
				Time("startTime", startTime).
				Msg("fetchAll error")
		}
	}
}
