package deribit

//wss://www.deribit.com/ws/api/v2                                                                                                                                                                                                                                             ✭ ✱ ◼
//> {"method":"public/subscribe","params":{"channels":["ticker.ETH-PERPETUAL.raw"]},"jsonrpc":"2.0","id":0}

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connectors/deribit/market"

	"github.com/golang/protobuf/ptypes"
	"github.com/rs/zerolog/log"
	gctconfig "github.com/thrasher-corp/gocryptotrader/config"
	"github.com/thrasher-corp/gocryptotrader/currency"
	"github.com/thrasher-corp/gocryptotrader/exchanges/deribit"
	"github.com/thrasher-corp/gocryptotrader/exchanges/ticker"
	_ "go.uber.org/automaxprocs"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type DeribitConnector struct {
	*connector.Connector
	sink chan protoreflect.ProtoMessage
}

func (c *DeribitConnector) Start() {
	c.sink = make(chan protoreflect.ProtoMessage)
	go c.InitProduceChannel(c.sink)

	// gocryptotrader setup
	exch := new(deribit.Deribit)
	exch.SetDefaults()

	cfg := &gctconfig.Cfg
	err := cfg.LoadConfig("", false)
	if err != nil {
		log.Error().Err(err).Msg("failed to load config")
	}
	exchConfig, err := cfg.GetExchangeConfig("Deribit")
	if err != nil {
		log.Error().Err(err).Msg("get config")
	}

	err = exch.Setup(exchConfig)
	if err != nil {
		log.Error().Err(err).Msg("setup")
	}

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
		//exchName := exch.Websocket.GetName()
		switch data.(type) {
		case *ticker.Ticker:
			inst := data.(*ticker.Ticker)
			if inst.DerivStatus.OpenInterest != cache[inst.Pair].OpenInterest {
				cache[inst.Pair] = UpdateData{inst.LastUpdated, inst.DerivStatus.OpenInterest}

				// write to Kafka
				ts, err := ptypes.TimestampProto(inst.LastUpdated)
				if err != nil {
					log.Panic().Err(err).Msg("")
				}

				// write to Kafka
				c.sink <- &market.OpenInterest{
					Ts:           ts,
					OpenInterest: inst.DerivStatus.OpenInterest,
					Asset:        inst.Pair.String(),
				}

				//fmt.Println(cache.String())
				log.Info().
					Float64("oi", inst.DerivStatus.OpenInterest).
					Str("sym", inst.Pair.String()).
					Msg("Update")
			}
		default:
			log.Info().Str("type", reflect.TypeOf(data).String()).Msg("unhandled type")
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
