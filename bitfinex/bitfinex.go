package bitfinex

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connectors/bitfinex/market"

	"github.com/rs/zerolog/log"
	gctconfig "github.com/thrasher-corp/gocryptotrader/config"
	"github.com/thrasher-corp/gocryptotrader/currency"
	"github.com/thrasher-corp/gocryptotrader/exchanges/bitfinex"
	"github.com/thrasher-corp/gocryptotrader/exchanges/ticker"
	_ "go.uber.org/automaxprocs"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BitfinexConnector struct {
	*connector.Connector
}

func (c *BitfinexConnector) Start() {
	// gocryptotrader setup
	exch := new(bitfinex.Bitfinex)
	exch.SetDefaults()

	cfg := &gctconfig.Cfg
	err := cfg.LoadConfig("", false)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	exchConfig, err := cfg.GetExchangeConfig("Bitfinex")
	if err != nil {
		log.Error().Err(err).Msg("get config")
	}

	exch.Setup(exchConfig)
	if err != nil {
		log.Error().Err(err).Msg("setup")
	}

	//log.Infof("%+v\n", exchConfig.CurrencyPairs.Pairs[asset.Futures])
	//log.Infof("%+v\n", exchConfig.CurrencyPairs.Pairs[asset.PerpetualContract])

	ws, err := exch.GetWebsocket()
	if err != nil {
		log.Error().Err(err).Msg("get websocket")
	}
	err = ws.Connect()
	if err != nil {
		log.Error().Err(err).Msg("connect websocket")
	}

	cache := make(Cache)

	for data := range exch.Websocket.ToRoutine {
		switch data.(type) {
		case *ticker.Ticker:
			inst := data.(*ticker.Ticker)
			if inst.DerivStatus.OpenInterest != cache[inst.Pair].OpenInterest && inst.DerivStatus.OpenInterest != 0 {
				cache[inst.Pair] = UpdateData{inst.LastUpdated, inst.DerivStatus.OpenInterest}

				// write to Kafka
				c.EventSink <- &market.OpenInterest{
					Ts:           timestamppb.New(inst.LastUpdated),
					OpenInterest: inst.DerivStatus.OpenInterest,
					Asset:        inst.Pair.String(),
				}

				log.Info().
					Time("ts", inst.LastUpdated).
					Float64("oi", inst.DerivStatus.OpenInterest).
					Str("sym", inst.Pair.String()).
					Msg("Update")
			}
		default:
			log.Info().Str("type", reflect.TypeOf(data).String()).Msg("unhandled ws response")
		}
	}
}

type Cache map[currency.Pair]UpdateData
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
