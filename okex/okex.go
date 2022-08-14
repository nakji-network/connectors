package okex

import (
	"fmt"
	"strings"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connectors/okex/market"

	"github.com/rs/zerolog/log"
	gctconfig "github.com/thrasher-corp/gocryptotrader/config"
	"github.com/thrasher-corp/gocryptotrader/currency"
	"github.com/thrasher-corp/gocryptotrader/exchanges/okex"
	"github.com/thrasher-corp/gocryptotrader/exchanges/ticker"
	_ "go.uber.org/automaxprocs"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OkexConnector struct {
	*connector.Connector
}

func (c *OkexConnector) Start() {
	// gocryptotrader setup
	exch := new(okex.OKEX)
	exch.SetDefaults()

	cfg := &gctconfig.Cfg
	err := cfg.LoadConfig("", false)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	exchConfig, err := cfg.GetExchangeConfig("OKEX")
	if err != nil {
		log.Error().Err(err).Msg("get config")
	}

	err = exch.Setup(exchConfig)
	if err != nil {
		log.Error().Err(err).Msg("setup")
	}

	//log.Infof("%+v\n", exchConfig.CurrencyPairs.Pairs[asset.Futures])
	//log.Infof("%+v\n", exchConfig.CurrencyPairs.Pairs[asset.PerpetualSwap])

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
		case *ticker.Price:
			inst := data.(*ticker.Price)
			if cache[inst.Pair].OpenInterest != inst.OpenInterest {
				cache[inst.Pair] = UpdateData{inst.LastUpdated, inst.OpenInterest}

				// write to Kafka
				log.Info().Str("exch.Name", exch.Name).Str("pair", inst.Pair.String())

				c.EventSink <- &market.OpenInterest{
					Ts:           timestamppb.New(inst.LastUpdated),
					OpenInterest: inst.OpenInterest,
					Asset:        inst.Pair.String(),
				}

				log.Info().Msgf("%s: %+v\tOI: %f", inst.LastUpdated, inst.Pair, inst.OpenInterest)
			}
		default:
			//log.Info(reflect.TypeOf(data))
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

//{
//"name": "OKEX",
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
//"assetEnabled": null,
//"enabled": "BTC-USD_200731",
//"available": "BTC-USD_200731,XRP-USD_200703,XRP-USD_200710,XRP-USD_200925,XRP-USD_201225,BTC-USD_200703,BTC-USD_200710,BTC-USD_200925,BTC-USD_201225,BTC-USDT_200703,BTC-USDT_200710,BTC-USDT_200925,BTC-USDT_201225,LTC-USD_200703,LTC-USD_200710,LTC-USD_200925,LTC-USD_201225,LTC-USDT_200703,LTC-USDT_200710,LTC-USDT_200925,LTC-USDT_201225,ETH-USD_200703,ETH-USD_200710,ETH-USD_200925,ETH-USD_201225,ETH-USDT_200703,ETH-USDT_200710,ETH-USDT_200925,ETH-USDT_201225,TRX-USDT_200703,TRX-USDT_200710,TRX-USDT_200925,TRX-USDT_201225,ETC-USD_200703,ETC-USD_200710,ETC-USD_200925,ETC-USD_201225,BCH-USD_200703,BCH-USD_200710,BCH-USD_200925,BCH-USD_201225,BCH-USDT_200703,BCH-USDT_200710,BCH-USDT_200925,BCH-USDT_201225,BSV-USD_200703,BSV-USD_200710,BSV-USD_200925,BSV-USD_201225,BSV-USDT_200703,BSV-USDT_200710,BSV-USDT_200925,BSV-USDT_201225,EOS-USDT_200703,EOS-USDT_200710,EOS-USDT_200925,EOS-USDT_201225,XRP-USDT_200703,XRP-USDT_200710,XRP-USDT_200925,XRP-USDT_201225,ETC-USDT_200703,ETC-USDT_200710,ETC-USDT_200925,ETC-USDT_201225,EOS-USD_200703,EOS-USD_200710,EOS-USD_200925,EOS-USD_201225,TRX-USD_200703,TRX-USD_200710,TRX-USD_200925,TRX-USD_201225",
//"requestFormat": {
//"uppercase": true,
//"delimiter": "-"
//},
//"configFormat": {
//"uppercase": true,
//"delimiter": "_"
//}
//},
//"index": {
//"assetEnabled": null,
//"enabled": "BTC-USD",
//"available": "XRP-USD,XRP-USD,XRP-USD,BTC-USD,BTC-USD,BTC-USD,BTC-USDT,BTC-USDT,BTC-USDT,LTC-USD,LTC-USD,LTC-USD,LTC-USDT,LTC-USDT,LTC-USDT,ETH-USD,ETH-USD,ETH-USD,ETH-USDT,ETH-USDT,ETH-USDT,TRX-USDT,TRX-USDT,TRX-USDT,ETC-USD,ETC-USD,ETC-USD,BCH-USD,BCH-USD,BCH-USD,BCH-USDT,BCH-USDT,BCH-USDT,BSV-USD,BSV-USD,BSV-USD,BSV-USDT,BSV-USDT,BSV-USDT,EOS-USDT,EOS-USDT,EOS-USDT,XRP-USDT,XRP-USDT,XRP-USDT,ETC-USDT,ETC-USDT,ETC-USDT,EOS-USD,EOS-USD,EOS-USD,TRX-USD,TRX-USD,TRX-USD",
//"requestFormat": {
//"uppercase": true,
//"delimiter": "-"
//},
//"configFormat": {
//"uppercase": true
//}
//},
//"perpetualswap": {
//"assetEnabled": true,
//"enabled": "ETH-USD_SWAP",
//"available": "NEO-USDT_SWAP,LINK-USDT_SWAP,DASH-USDT_SWAP,ADA-USDT_SWAP,ZEC-USDT_SWAP,XTZ-USDT_SWAP,ONT-USDT_SWAP,ATOM-USDT_SWAP,QTUM-USDT_SWAP,XLM-USDT_SWAP,XMR-USDT_SWAP,IOTA-USDT_SWAP,ALGO-USDT_SWAP,IOST-USDT_SWAP,THETA-USDT_SWAP,KNC-USDT_SWAP,NEO-USD_SWAP,LINK-USD_SWAP,DASH-USD_SWAP,ADA-USD_SWAP,ZEC-USD_SWAP,XTZ-USD_SWAP,ONT-USD_SWAP,ATOM-USD_SWAP,QTUM-USD_SWAP,XLM-USD_SWAP,XMR-USD_SWAP,IOTA-USD_SWAP,ALGO-USD_SWAP,IOST-USD_SWAP,THETA-USD_SWAP,KNC-USD_SWAP,BTC-USDT_SWAP,LTC-USDT_SWAP,ETH-USDT_SWAP,TRX-USDT_SWAP,BCH-USDT_SWAP,BSV-USDT_SWAP,EOS-USDT_SWAP,XRP-USDT_SWAP,ETC-USDT_SWAP,BTC-USD_SWAP,LTC-USD_SWAP,ETH-USD_SWAP,TRX-USD_SWAP,BCH-USD_SWAP,BSV-USD_SWAP,EOS-USD_SWAP,XRP-USD_SWAP,ETC-USD_SWAP",
//"requestFormat": {
//"uppercase": true,
//"delimiter": "-"
//},
//"configFormat": {
//"uppercase": true,
//"delimiter": "_"
//}
//},
//"spot": {
//"assetEnabled": null,
//"enabled": "EOS-USDT",
//"available": "XPO-USDT,SPND-BTC,ROAD-USDK,RVN-USDT,HDAO-USDK,BAT-USDT,OXT-USDT,OXT-BTC,BCH-BTC,BSV-BTC,USDC-BTC,DASH-BTC,ADA-BTC,AE-BTC,ALGO-BTC,APM-BTC,ARDR-BTC,ATOM-BTC,BAT-BTC,BHP-BTC,BTT-BTC,COMP-BTC,CRO-BTC,CTC-BTC,CTXC-BTC,CVT-BTC,DCR-BTC,DNA-BTC,EGT-BTC,GUSD-BTC,HBAR-BTC,HYC-BTC,IQ-BTC,KAN-BTC,LBA-BTC,LEO-BTC,LET-BTC,LSK-BTC,ORS-BTC,PAX-BTC,PMA-BTC,RVN-BTC,SC-BTC,TMTG-BTC,TUSD-BTC,VITE-BTC,VSYS-BTC,WAVES-BTC,WXT-BTC,XTZ-BTC,YOU-BTC,ZIL-BTC,XRP-BTC,ELF-BTC,LRC-BTC,MCO-BTC,NULS-BTC,BCX-BTC,CMT-BTC,ITC-BTC,SBTC-BTC,ZEC-BTC,NEO-BTC,GAS-BTC,HC-BTC,QTUM-BTC,IOTA-BTC,EOS-BTC,SNT-BTC,OMG-BTC,LTC-BTC,ETH-BTC,ETC-BTC,BCD-BTC,BTG-BTC,ACT-BTC,PAY-BTC,BTM-BTC,GNT-BTC,LINK-BTC,WTC-BTC,ZRX-BTC,BNT-BTC,CVC-BTC,MANA-BTC,KNC-BTC,GNX-BTC,ICX-BTC,XEM-BTC,ARK-BTC,YOYO-BTC,FUN-BTC,TRX-BTC,DGB-BTC,SWFTC-BTC,XMR-BTC,XLM-BTC,KCASH-BTC,MDT-BTC,NAS-BTC,AAC-BTC,VIB-BTC,QUN-BTC,INT-BTC,IOST-BTC,MOF-BTC,TCT-BTC,THETA-BTC,PST-BTC,SNC-BTC,MKR-BTC,TRUE-BTC,SOC-BTC,ZEN-BTC,NANO-BTC,GTO-BTC,CHAT-BTC,MITH-BTC,ABT-BTC,TRIO-BTC,ONT-BTC,OKB-BTC,ADA-ETH,AE-ETH,ALGO-ETH,ATOM-ETH,BTT-ETH,CTXC-ETH,EGT-ETH,HYC-ETH,KAN-ETH,LEO-ETH,ORS-ETH,SC-ETH,WAVES-ETH,YOU-ETH,ZIL-ETH,ELF-ETH,LTC-ETH,CMT-ETH,LRC-ETH,MCO-ETH,NULS-ETH,STORJ-ETH,BTM-ETH,EOS-ETH,OMG-ETH,DASH-ETH,XRP-ETH,ZEC-ETH,NEO-ETH,GAS-ETH,HC-ETH,QTUM-ETH,IOTA-ETH,ETC-ETH,LINK-ETH,WTC-ETH,ZRX-ETH,CVC-ETH,MANA-ETH,GNX-ETH,XEM-ETH,TRX-ETH,SWFTC-ETH,XMR-ETH,XLM-ETH,KCASH-ETH,MDT-ETH,NAS-ETH,AAC-ETH,FAIR-ETH,TOPC-ETH,INT-ETH,IOST-ETH,MOF-ETH,MKR-ETH,TRUE-ETH,ZEN-ETH,NANO-ETH,GTO-ETH,MITH-ETH,ABT-ETH,TRIO-ETH,ONT-ETH,OKB-ETH,BTC-USDK,LTC-USDK,ETH-USDK,OKB-USDK,ETC-USDK,BCH-USDK,EOS-USDK,XRP-USDK,TRX-USDK,BSV-USDK,USDT-USDK,ALGO-USDK,CRO-USDK,DEP-USDK,DOGE-USDK,EC-USDK,EM-USDK,FSN-USDK,FTM-USDK,HBAR-USDK,LAMB-USDK,LEO-USDK,NDN-USDK,ORBS-USDK,PLG-USDK,PMA-USDK,VSYS-USDK,WGRT-USDK,WXT-USDK,BCH-USDT,BSV-USDT,USDC-USDT,ADA-USDT,AE-USDT,ALGO-USDT,ALV-USDT,APM-USDT,ATOM-USDT,BHP-USDT,BLOC-USDT,BTT-USDT,COMP-USDT,CRO-USDT,CTC-USDT,CTXC-USDT,CVT-USDT,DAI-USDT,DCR-USDT,DEP-USDT,DNA-USDT,DOGE-USDT,EC-USDT,EGT-USDT,EM-USDT,ETM-USDT,FSN-USDT,FTM-USDT,GUSD-USDT,HBAR-USDT,HDAO-USDT,HYC-USDT,IQ-USDT,KAN-USDT,LAMB-USDT,LBA-USDT,LEO-USDT,LET-USDT,LSK-USDT,NDN-USDT,ORBS-USDT,ORS-USDT,PAX-USDT,PLG-USDT,ROAD-USDT,SC-USDT,TMTG-USDT,TUSD-USDT,VSYS-USDT,WAVES-USDT,WGRT-USDT,WXT-USDT,XTZ-USDT,YOU-USDT,ZIL-USDT,TRX-OKB,EGT-OKB,SC-OKB,WXT-OKB,BTC-DAI,ETH-DAI,ELF-USDT,DASH-USDT,BTG-USDT,LRC-USDT,MCO-USDT,NULS-USDT,DASH-OKB,XRP-USDT,ZEC-USDT,NEO-USDT,GAS-USDT,HC-USDT,QTUM-USDT,IOTA-USDT,BTC-USDT,BCD-USDT,XUC-USDT,CMT-USDT,ITC-USDT,ETH-USDT,LTC-USDT,ETC-USDT,EOS-USDT,OMG-USDT,ACT-USDT,BTM-USDT,GNT-USDT,PAY-USDT,STORJ-USDT,SNT-USDT,LINK-USDT,WTC-USDT,ZRX-USDT,BNT-USDT,CVC-USDT,MANA-USDT,KNC-USDT,ICX-USDT,XEM-USDT,ARK-USDT,YOYO-USDT,AST-USDT,TRX-USDT,MDA-USDT,DGB-USDT,PPT-USDT,SWFTC-USDT,XMR-USDT,XLM-USDT,KCASH-USDT,MDT-USDT,NAS-USDT,RNT-USDT,AAC-USDT,FAIR-USDT,UBTC-USDT,VIB-USDT,UTK-USDT,TOPC-USDT,QUN-USDT,INT-USDT,IOST-USDT,YEE-USDT,MOF-USDT,TCT-USDT,THETA-USDT,PST-USDT,MKR-USDT,TRUE-USDT,SOC-USDT,ZEN-USDT,ZIP-USDT,NANO-USDT,GTO-USDT,CHAT-USDT,BEC-USDT,MITH-USDT,ABT-USDT,TRIO-USDT,ONT-USDT,OKB-USDT,NEO-OKB,LTC-OKB,ETC-OKB,XRP-OKB,ZEC-OKB,IOTA-OKB,EOS-OKB,BTC-USDC,LTC-USDC,ETH-USDC,OKB-USDC,ETC-USDC,BCH-USDC,EOS-USDC,XRP-USDC,TRX-USDC,BSV-USDC",
//"requestFormat": {
//"uppercase": true,
//"delimiter": "-"
//},
//"configFormat": {
//"uppercase": true,
//"delimiter": "-"
//}
//}
//}
//},
//"api": {
//"authenticatedSupport": false,
//"authenticatedWebsocketApiSupport": false,
//"endpoints": {
//"url": "NON_DEFAULT_HTTP_LINK_TO_EXCHANGE_API",
//"urlSecondary": "NON_DEFAULT_HTTP_LINK_TO_EXCHANGE_API",
//"websocketURL": "wss://real.okex.com:8443/ws/v3"
//},
//"credentials": {
//"key": "Key",
//"secret": "Secret"
//},
//"credentialsValidator": {
//"requiresKey": true,
//"requiresSecret": true,
//"requiresClientID": true
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
