package huobi

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connectors/huobi/market"

	"github.com/rs/zerolog/log"
	gctconfig "github.com/thrasher-corp/gocryptotrader/config"
	"github.com/thrasher-corp/gocryptotrader/currency"
	"github.com/thrasher-corp/gocryptotrader/exchanges/huobi"
	_ "go.uber.org/automaxprocs"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type HuobiConnector struct {
	*connector.Connector
}

func (c *HuobiConnector) Start() {
	// gocryptotrader setup
	exch := new(huobi.HUOBI)
	exch.SetDefaults()

	cfg := &gctconfig.Cfg
	err := cfg.LoadConfig("", false)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	exchConfig, err := cfg.GetExchangeConfig("Huobi")
	if err != nil {
		log.Error().Err(err).Msg("get config")
	}

	err = exch.Setup(exchConfig)
	if err != nil {
		log.Error().Err(err).Msg("setup")
	}

	cache := make(Cache)

	for range time.NewTicker(100 * time.Millisecond).C {
		ts, oi, err := exch.GetSwapOpenInterest(huobi.OpenInterestRequestParams{})
		if err != nil {
			log.Error().Err(err).Msg("GetSwapOpenInterest")
		}

		//fmt.Printf("%s\n", ts)

		for _, d := range oi {
			p, err := currency.NewPairFromString(d.ContractCode)
			if err != nil {
				log.Fatal().Err(err).Msg("")
			}
			if d.Amount != cache[p].OpenInterest && d.Amount != 0 {
				cache[p] = UpdateData{LastUpdated: ts, OpenInterest: d.Amount}

				c.EventSink <- &market.OpenInterest{
					Ts:                timestamppb.New(ts),
					OpenInterest:      d.Volume,
					OpenInterestValue: d.Amount,
					Asset:             p.String(),
				}

				log.Info().
					Float64("oi", d.Volume). // in usd
					Float64("oiv", d.Amount).
					Str("sym", p.String()).
					Msg("Update")
			}
		}
	}
}

type Cache map[currency.Pair]UpdateData
type UpdateData struct {
	LastUpdated  time.Time
	OpenInterest float64
}

type ByPairString []currency.Pair

func (a ByPairString) Len() int           { return len(a) }
func (a ByPairString) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPairString) Less(i, j int) bool { return a[i].String() < a[j].String() }

func (c Cache) OrderedString() string {
	var res strings.Builder
	keys := make([]currency.Pair, 0)
	for k := range c {
		keys = append(keys, k)
	}
	sort.Sort(ByPairString(keys))
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
