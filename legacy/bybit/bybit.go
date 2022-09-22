package bybit

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/thrasher-corp/gocryptotrader/common"
	"github.com/thrasher-corp/gocryptotrader/currency"
	exchange "github.com/thrasher-corp/gocryptotrader/exchanges"
	"github.com/thrasher-corp/gocryptotrader/exchanges/asset"
	"github.com/thrasher-corp/gocryptotrader/exchanges/bitmex"
	"github.com/thrasher-corp/gocryptotrader/exchanges/protocol"
	"github.com/thrasher-corp/gocryptotrader/exchanges/request"
	"github.com/thrasher-corp/gocryptotrader/exchanges/stream"
	"github.com/thrasher-corp/gocryptotrader/log"
)

type Bybit struct {
	exchange.Base
	endpoints *exchange.Endpoints
}

const (
	bybitAPIURL = "https://api.bybit.com"
	bybitWSURL  = ""

	bybitEndpointInstruments = "/v2/public/tickers"
)

func (b *Bybit) GetInstruments(params *GenericRequestParams) (time.Time, []Instrument, error) {
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
func (b *Bybit) SetDefaults() error {
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
		log.Errorln(log.ExchangeSys, err)
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
func (b *Bybit) SendHTTPRequest(path string, params bitmex.Parameter, result interface{}) error {
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
