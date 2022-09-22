package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"blep.ai/data/monitor"

	"blep.ai/data/config"
	"blep.ai/data/connectors/source"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/golang/protobuf/ptypes"
	"github.com/rs/zerolog/log"
	gctconfig "github.com/thrasher-corp/gocryptotrader/config"
	"github.com/thrasher-corp/gocryptotrader/exchanges/ftx"
)

var (
	transactionalID   = "ftx-source"
	OITOPIC           = ".fct.nakji.common.0_0_0.market_openinterest"
	subscribedFutures = []string{"ADA", "ALGO", "ALT", "ATOM", "BAL", "BCH", "BTC", "BNB", "BRZ", "BSV"}
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
	monitor.StartMonitor("ftx")

	// gocryptotrader setup
	exch := new(ftx.FTX)
	exch.SetDefaults()

	cfg := &gctconfig.Cfg
	err = cfg.LoadConfig("", false)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	exchConfig, err := cfg.GetExchangeConfig("FTX")
	if err != nil {
		log.Error().Err(err).Msg("get config")
	}

	err = exch.Setup(exchConfig)
	if err != nil {
		log.Error().Err(err).Msg("setup")
	}

	//ws, err := exch.GetWebsocket()
	//if err != nil {
	//log.WithField("err", err).Error("get websocket")
	//}
	//err = ws.Connect()
	//if err != nil {
	//log.WithField("err", err).Error("connect websocket")
	//}

	cache := make(Cache)

	//for data := range exch.Websocket.ToRoutine {
	////exchName := exch.Websocket.GetName()
	//switch data.(type) {
	//case *ticker.Ticker:
	//inst := data.(*ticker.Ticker)
	//if inst.DerivStatus.OpenInterest != cache[inst.Pair].OpenInterest {
	//cache[inst.Pair] = UpdateData{inst.LastUpdated, inst.DerivStatus.OpenInterest}
	//fmt.Println(cache.String())
	//}
	//default:
	//log.Info(reflect.TypeOf(data))
	//}
	//}

	minInterval, subbedFutures := UpdateFutures(exch)

	// Add 50% additional since OI doesn't actually change that fast
	for range time.NewTicker(time.Duration(1.5*minInterval) * time.Millisecond).C {
		for _, f := range subbedFutures {
			futureStatsData, err := exch.GetFutureStats(f)
			if err != nil {
				log.Info().Str("Pair", f).Err(err).Msg("GetFutureStats")
				log.Info().Msg("Attempting to updated futures")
				minInterval, subbedFutures = UpdateFutures(exch)
			}
			if futureStatsData.OpenInterest != cache[f].OpenInterest && futureStatsData.OpenInterest != 0 {
				ts := time.Now()
				cache[f] = UpdateData{LastUpdated: ts, OpenInterest: futureStatsData.OpenInterest}

				// write to Kafka
				tsp, err := ptypes.TimestampProto(ts)
				if err != nil {
					log.Panic().Err(err).Msg("")
				}

				key := kafkautils.NewKey(exch.Name, f)
				err = kp.WriteAndCommit(oiTopic, key.Bytes(), &source.OpenInterest{
					Ts:           tsp,
					OpenInterest: futureStatsData.OpenInterest,
				})
				if err != nil {
					log.Error().Err(err)
				}

				//fmt.Println(cache.OrderedString())
				log.Info().
					Time("ts", ts).
					Float64("oi", futureStatsData.OpenInterest). // in usd
					Str("sym", f).
					Msg("Update")
			}
		}
	}
}

func UpdateFutures(exch *ftx.FTX) (float64, []string) {
	updatedFutures := []string{}
	futures, err := exch.GetFutures()
	if err != nil {
		log.Error().Err(err).Msg("GetFutures")
		return 0, nil
	}
	for _, f := range futures {
		for _, sf := range subscribedFutures {
			if strings.HasPrefix(f.Name, sf) {
				updatedFutures = append(updatedFutures, f.Name)
				break
			}
		}
	}

	// ftx max 30 req/second. We limit to 100 req/second.
	minInterval := 1000 / 30 * float64(len(updatedFutures))
	if minInterval < 100 {
		minInterval = 100
	}
	log.Info().Strs("Subbed Futures", updatedFutures).Msg("")
	return minInterval, updatedFutures
}

type Cache map[string]UpdateData
type UpdateData struct {
	LastUpdated  time.Time
	OpenInterest float64
}

type ByString []string

func (c Cache) OrderedString() string {
	var res strings.Builder
	keys := make([]string, 0)
	for k := range c {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		res.WriteString(fmt.Sprintf("%15s %s\t%20f\n", k, c[k].LastUpdated, c[k].OpenInterest))
	}
	return res.String()
}

func (c Cache) String() string {
	var res strings.Builder
	for k, v := range c {
		res.WriteString(fmt.Sprintf("%15s %s\t%f\n", k, v.LastUpdated, v.OpenInterest))
	}
	return res.String()
}
