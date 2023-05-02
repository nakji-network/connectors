package bybit

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/kafkautils"
	"github.com/nakji-network/connectors/bybit/market"

	"github.com/rs/zerolog/log"
	"github.com/thrasher-corp/gocryptotrader/common"
	"github.com/thrasher-corp/gocryptotrader/currency"
	exchange "github.com/thrasher-corp/gocryptotrader/exchanges"
	"github.com/thrasher-corp/gocryptotrader/exchanges/asset"
	"github.com/thrasher-corp/gocryptotrader/exchanges/bitmex"
	"github.com/thrasher-corp/gocryptotrader/exchanges/protocol"
	"github.com/thrasher-corp/gocryptotrader/exchanges/request"
	"github.com/thrasher-corp/gocryptotrader/exchanges/stream"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BybitConnector struct {
	*connector.Connector
	exchange.Base
	endpoints *exchange.Endpoints
}

const (
	bybitAPIURL = "https://api.bybit.com"
	bybitWSURL  = ""

	bybitEndpointInstruments = "/v2/public/tickers"
)

func (c *BybitConnector) Start() {
	// gocryptotrader setup
	exch := new(BybitConnector)
	err := exch.SetDefaults()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to set defaults")
	}
	// call https://api.bybit.com/v2/public/tickers to get all tickers
	// ws.send('{"op": "subscribe", "args": ["instrument_info.100ms.BTCUSD"]}')

	cache := make(Cache)

	// Since open_interest updates only possible 1/min, read every few seconds? or use sync
	for range time.NewTicker(100 * time.Millisecond).C {
		ts, instruments, err := exch.GetInstruments(&GenericRequestParams{})
		if err != nil {
			log.Error().Err(err).Msg("GetInstruments")
		}

		for _, d := range instruments {
			// update cache if it's new
			if d.OpenInterest != cache[d.Symbol].OpenInterest {
				cache[d.Symbol] = UpdateData{ts, d.OpenInterest}

				var ov float64
				if d.OpenValue == "" {
					ov = 0
				} else {
					ov, err = strconv.ParseFloat(d.OpenValue, 64)
					if err != nil {
						log.Error().Err(err).Msg("failed to convert to float")
						continue
					}
				}

				msg := &market.OpenInterest{
					Ts:                timestamppb.New(ts),
					OpenInterest:      d.OpenInterest,
					OpenInterestValue: ov,
					Asset:             d.Symbol,
				}

				// write to Kafka
				c.EventSink <- &kafkautils.Message{
					MsgType:  kafkautils.MsgTypeFct,
					ProtoMsg: msg,
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

func (b *BybitConnector) GetInstruments(params *GenericRequestParams) (time.Time, []Instrument, error) {
	var result = struct {
		ResultWrapper
		Result []Instrument `json:"result"`
	}{}

	err := b.SendHTTPRequest(bybitEndpointInstruments,
		params,
		&result)

	if err != nil {
		return time.Time{}, nil, err
	}

	// Parse the integer and decimal portions of the Unix timestamp into Time.time object
	timeNowIntStr := strings.Split(result.TimeNow, ".")[0]
	timeNowDecimalStr := strings.Split(result.TimeNow, ".")[1]

	timeNowInt, err := strconv.ParseInt(timeNowIntStr, 10, 64)

	if err != nil {
		return time.Time{}, nil, err
	}
	timeNowDecimal, err := strconv.ParseInt(timeNowDecimalStr, 10, 64)

	if err != nil {
		return time.Time{}, nil, err
	}

	timeNow := time.Unix(timeNowInt, timeNowDecimal)

	return timeNow, result.Result, nil
}

// SetDefaults sets the basic defaults for Bitmex
func (b *BybitConnector) SetDefaults() error {
	b.Name = "Bybit"
	b.Enabled = true
	b.Verbose = true
	b.API.CredentialsValidator.RequiresKey = true
	b.API.CredentialsValidator.RequiresSecret = true

	requestFmt := &currency.PairFormat{Uppercase: true}
	configFmt := &currency.PairFormat{Uppercase: true}
	err := b.SetGlobalPairsManager(requestFmt,
		configFmt,
		asset.PerpetualContract,
		asset.Futures,
		asset.Index)
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}

	b.Features = exchange.Features{
		Supports: exchange.FeaturesSupported{
			REST:      true,
			Websocket: true,
			RESTCapabilities: protocol.Features{
				TickerBatching:      true,
				TickerFetching:      true,
				TradeFetching:       true,
				OrderbookFetching:   true,
				AutoPairUpdates:     true,
				AccountInfo:         true,
				GetOrder:            true,
				GetOrders:           true,
				CancelOrders:        true,
				CancelOrder:         true,
				SubmitOrder:         true,
				SubmitOrders:        true,
				ModifyOrder:         true,
				DepositHistory:      true,
				WithdrawalHistory:   true,
				UserTradeHistory:    true,
				CryptoDeposit:       true,
				CryptoWithdrawal:    true,
				TradeFee:            true,
				CryptoWithdrawalFee: true,
			},
			WebsocketCapabilities: protocol.Features{
				TradeFetching:          true,
				OrderbookFetching:      true,
				Subscribe:              true,
				Unsubscribe:            true,
				AuthenticatedEndpoints: true,
				AccountInfo:            true,
				DeadMansSwitch:         true,
				GetOrders:              true,
				GetOrder:               true,
			},
			WithdrawPermissions: exchange.AutoWithdrawCryptoWithAPIPermission |
				exchange.WithdrawCryptoWithEmail |
				exchange.WithdrawCryptoWith2FA |
				exchange.NoFiatWithdrawals,
		},
		Enabled: exchange.FeaturesEnabled{
			AutoPairUpdates: true,
		},
	}

	b.Requester = request.New(b.Name,
		common.NewHTTPClientWithTimeout(exchange.DefaultHTTPTimeout),
		//request.WithLimiter(SetRateLimit()))
	)
	b.endpoints = b.NewEndpoints()
	b.endpoints.SetDefaultEndpoints(map[exchange.URL]string{
		exchange.RestSpot:      bybitAPIURL,
		exchange.WebsocketSpot: bybitWSURL,
	})

	b.Websocket = stream.New()
	b.WebsocketResponseMaxLimit = exchange.DefaultWebsocketResponseMaxLimit
	b.WebsocketResponseCheckTimeout = exchange.DefaultWebsocketResponseCheckTimeout
	b.WebsocketOrderbookBufferLimit = exchange.DefaultWebsocketOrderbookBufferLimit
	return nil
}

// SendHTTPRequest sends an unauthenticated HTTP request
func (b *BybitConnector) SendHTTPRequest(path string, params bitmex.Parameter, result interface{}) error {
	var respCheck interface{}
	endpoint, err := b.endpoints.GetURL(exchange.RestSpot)
	if err != nil {
		return err
	}
	path = endpoint + path
	if params != nil {
		if !params.IsNil() {
			encodedPath, err := params.ToURLVals(path)
			if err != nil {
				return err
			}
			err = b.SendPayload(context.Background(), &request.Item{
				Method:        http.MethodGet,
				Path:          encodedPath,
				Result:        &respCheck,
				Verbose:       b.Verbose,
				HTTPDebugging: b.HTTPDebugging,
				HTTPRecording: b.HTTPRecording,
			})
			if err != nil {
				return err
			}
			//return b.CaptureError(respCheck, result)
			marshalled, err := json.Marshal(respCheck)
			if err != nil {
				return err
			}

			return json.Unmarshal(marshalled, result)
		}
	}

	err = b.SendPayload(context.Background(), &request.Item{
		Method:        http.MethodGet,
		Path:          path,
		Result:        &respCheck,
		Verbose:       b.Verbose,
		HTTPDebugging: b.HTTPDebugging,
		HTTPRecording: b.HTTPRecording,
	})
	if err != nil {
		return err
	}
	//return b.CaptureError(respCheck, result)
	marshalled, err := json.Marshal(respCheck)
	if err != nil {
		return err
	}

	return json.Unmarshal(marshalled, result)
}
