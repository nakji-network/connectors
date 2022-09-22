package main

import (
	"blep.ai/data/config"
	"blep.ai/data/connectors/source"
	"blep.ai/data/monitor"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/golang/protobuf/ptypes"
	"github.com/rs/zerolog/log"
	gctconfig "github.com/thrasher-corp/gocryptotrader/config"
	"github.com/thrasher-corp/gocryptotrader/exchanges/asset"
	"github.com/thrasher-corp/gocryptotrader/exchanges/bitmex"
	"github.com/thrasher-corp/gocryptotrader/exchanges/ticker"
)

var (
	transactionalID = "bitmex-source"
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

	// Init exchange
	exch := new(bitmex.Bitmex)
	exch.SetDefaults()

	cfg := &gctconfig.Cfg
	err = cfg.LoadConfig("", false)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	exchConfig, err := cfg.GetExchangeConfig("Bitmex")
	if err != nil {
		log.Error().Err(err).Msg("get config")
	}

	exch.Setup(exchConfig)
	if err != nil {
		log.Error().Err(err).Msg("setup")
	}

	log.Info().Interface("futures", exchConfig.CurrencyPairs.Pairs[asset.Futures]).Msg("")
	log.Info().Interface("perps", exchConfig.CurrencyPairs.Pairs[asset.PerpetualContract]).Msg("")

	ws, err := exch.GetWebsocket()
	if err != nil {
		log.Error().Err(err).Msg("get websocket")
	}
	err = ws.Connect()
	if err != nil {
		log.Error().Err(err).Msg("connect websocket")
	}

	kp.EnableTransactions()
	monitor.StartMonitor("bitmex")

	for data := range exch.Websocket.ToRoutine {
		//exchName := exch.Websocket.GetName()
		switch data.(type) {
		case *ticker.Price:
			inst := data.(*ticker.Price)
			if inst.OpenInterest != 0 {
				log.Info().
					Time("ts", inst.LastUpdated).
					Float64("oi", inst.OpenInterest).
					Str("sym", inst.Pair.String()).
					Msg("Update")

				ts, err := ptypes.TimestampProto(inst.LastUpdated)
				if err != nil {
					log.Panic().Err(err).Msg("")
				}

				key := kafkautils.NewKey(exch.Name, inst.Pair.String())
				err = kp.WriteAndCommit(oiTopic, key.Bytes(), &source.OpenInterest{
					Ts:           ts,
					OpenInterest: inst.OpenInterest,
				})
				if err != nil {
					log.Fatal().Err(err)
				}
			}
		default:
		}
	}
}

//{
//"name": "Bitmex",
//"enabled": true,
//"verbose": false,
//"httpTimeout": 15000000000,
//"websocketResponseCheckTimeout": 30000000,
//"websocketResponseMaxLimit": 7000000000,
//"websocketTrafficTimeout": 30000000000,
//"websocketOrderbookBufferLimit": 5,
//"baseCurrencies": "USD",
//"currencyPairs": {
//"pairs": {
//"futures": {
//"assetEnabled": false,
//"enabled": "",
//"available": "",
//"requestFormat": {
//"uppercase": true
//},
//"configFormat": {
//"uppercase": true
//}
//},
//"perpetualcontract": {
//"assetEnabled": true,
//"enabled": "XBTUSD,ETHUSD",
//"available": "XBTUSD,ETHUSD",
//"requestFormat": {
//"uppercase": true
//},
//"configFormat": {
//"uppercase": true
//}
//}
//}
//},
//"api": {
//"authenticatedSupport": false,
//"authenticatedWebsocketApiSupport": false,
//"endpoints": {
//"url": "https://www.bitmex.com/api/v1",
//"urlSecondary": "NON_DEFAULT_HTTP_LINK_TO_EXCHANGE_API",
//"websocketURL": "wss://www.bitmex.com/realtime"
//},
//"credentials": {
//"key": "",
//"secret": ""
//},
//"credentialsValidator": {
//"requiresKey": true,
//"requiresSecret": true
//}
//},
//"features": {
//"supports": {
//"restAPI": true,
//"restCapabilities": {
//"tickerBatching": true,
//"autoPairUpdates": true
//},
//"websocketAPI": true,
//"websocketCapabilities": {}
//},
//"enabled": {
//"autoPairUpdates": true,
//"websocketAPI": true
//}
//},
//"bankAccounts": [
//{
//"enabled": false,
//"bankName": "",
//"bankAddress": "",
//"bankPostalCode": "",
//"bankPostalCity": "",
//"bankCountry": "",
//"accountName": "",
//"accountNumber": "",
//"swiftCode": "",
//"iban": "",
//"supportedCurrencies": ""
//}
//]
//},
