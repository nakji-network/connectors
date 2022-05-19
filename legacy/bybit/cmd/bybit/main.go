package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"blep.ai/data/monitor"

	"blep.ai/data/config"
	"blep.ai/data/connectors/source"
	"blep.ai/data/connectors/source/bybit"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/golang/protobuf/ptypes"
	"github.com/rs/zerolog/log"
)

var (
	transactionalID = "bybit-source"
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
	monitor.StartMonitor("bybit")

	// gocryptotrader setup
	exch := new(bybit.Bybit)
	err = exch.SetDefaults()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to set defaults")
	}
	// call https://api.bybit.com/v2/public/tickers to get all tickers
	// ws.send('{"op": "subscribe", "args": ["instrument_info.100ms.BTCUSD"]}')

	cache := make(Cache)

	// Since open_interest updates only possible 1/min, read every few seconds? or use sync
	for range time.NewTicker(100 * time.Millisecond).C {
		ts, instruments, err := exch.GetInstruments(&bybit.GenericRequestParams{})
		if err != nil {
			log.Error().Err(err).Msg("GetInstruments")
		}

		for _, d := range instruments {
			// update cache if it's new
			if d.OpenInterest != cache[d.Symbol].OpenInterest {
				cache[d.Symbol] = UpdateData{ts, d.OpenInterest}

				// write to Kafka
				ts, err := ptypes.TimestampProto(ts)
				if err != nil {
					log.Panic().Err(err).Msg("")
				}

				ov, err := strconv.ParseFloat(d.OpenValue, 64)
				if err != nil {
					log.Panic().Err(err).Msg("")
				}

				key := kafkautils.NewKey(exch.Name, d.Symbol)
				err = kp.WriteAndCommit(oiTopic, key.Bytes(), &source.OpenInterest{
					Ts:                ts,
					OpenInterest:      d.OpenInterest,
					OpenInterestValue: ov,
				})
				if err != nil {
					log.Error().Err(err)
				}

				log.Info().
					Float64("oi", d.OpenInterest).
					Float64("oiv", ov).
					Str("sym", d.Symbol).
					Msg("Update")

				//fmt.Printf("%s\t%s\t%0f  \t%f\n", ts, d.Symbol, d.OpenInterest, ov)
			}
		}
	}
}

type Cache map[string]UpdateData
type UpdateData struct {
	LastUpdated  time.Time
	OpenInterest float64
}

func (c *Cache) String() string {
	var res strings.Builder
	for k, v := range *c {
		res.WriteString(fmt.Sprintf("%15s %s\t%f\n", k, v.LastUpdated, v.OpenInterest))
	}
	return res.String()
}
