package binance

import (
	"fmt"
	"net/url"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/common"
	"github.com/nakji-network/connectors/binance/market"

	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
	_ "go.uber.org/automaxprocs"
)

type BinanceConnector struct {
	*connector.Connector
}

var (
	baseURL  = "https://fapi.binance.com/futures/data/openInterestHist?"
	exchange = "binance"
	symbols  = []string{"BTCUSDT", "ETHUSDT"}
	period   = "5m"
)

func (c *BinanceConnector) Start() {
	// Timed fetch
	for _, symbol := range symbols {
		params := url.Values{}
		params.Add("limit", "5")
		params.Add("symbol", symbol)
		params.Add("period", period)

		tfOptions := TimedFetchOptions{
			Interval: 5 * 60 * 1000, // 5 minutes
			Margin:   550,
			Delay:    4 * 60 * 1000, // 4 minutes
			Decay:    6,
		}

		TimedFetch(baseURL, params, tfOptions, func(val *gjson.Result, err error) {
			if err != nil {
				log.Fatal().Err(err).Msg("")
			}

			msInt := val.Get("timestamp").Int()
			ts := common.UnixToTimestampPb(msInt)

			// write to Kafka
			c.EventSink <- &market.OpenInterest{
				Ts:                ts,
				OpenInterest:      val.Get("sumOpenInterest").Float(),
				OpenInterestValue: val.Get("sumOpenInterestValue").Float(),
				Asset:             symbol,
			}
		})
	}

	select {}
}

// fetchAll automatically crawls entire dataset from `from` to `to`, making
// multiple requests as necessary. This has several variables that are specific
// to Binance; hence refactoring may not be trivial.
func fetchAll(symbol, period string, from, to time.Time, callback func(*market.OpenInterest)) error {
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
		resJSON, err := GetJSON(baseURL, params)
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
				Row := &market.OpenInterest{
					Ts:                common.UnixToTimestampPb(msInt),
					OpenInterest:      val.Get("sumOpenInterest").Float(),
					OpenInterestValue: val.Get("sumOpenInterestValue").Float(),
				}

				callback(Row)
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
