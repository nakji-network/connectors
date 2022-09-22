// TODO: handle the fact that binance websocket connections are limited at 24 hours.

package binance

//import (
//"context"
//"encoding/json"
//"fmt"
//"log"
//"os"
//"strings"
//"time"

//"net/url"
//"os/signal"

//"github.com/gorilla/websocket"
//kafka "github.com/segmentio/kafka-go"
//"github.com/tidwall/gjson"
//"google.golang.org/protobuf/proto"

//"blep.ai/data/common"
//"blep.ai/data/currency"
//"github.com/nakji-network/connector/kafkautils"
//"blep.ai/data/stream"
//)

//func convertStreamName(s string) []byte {
//a := strings.FieldsFunc(s, func(r rune) bool {
//return r == '@' || r == '_'
//})
//if len(a) == 3 {
//return stream.NewStreamName("binance", a[0], a[1], a[2]).Bytes()
//}
//return stream.NewStreamName("binance", a[0], a[1]).Bytes()
//}

//func Init(kconn *kafkautils.Conn) {
//interrupt := make(chan os.Signal, 1)
//signal.Notify(interrupt, os.Interrupt)

//w := kconn.Writer()

//defer w.Close()

//// connect to websockets
//streams := []string{
//"mBTCUSDT@kline_1m", // mark price
//"iBTCUSDT@kline_1m", // index price
////"btcusdt@trade",
////"ethusdt@trade",
//}
//u := url.URL{Scheme: "wss", Host: "fstream.binance.com", Path: "/stream", RawQuery: "streams=" + strings.Join(streams, "/")}
//log.Printf("connecting to %s", u.String())

//c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
//if err != nil {
//log.Fatal("dial:", err)
//}
//defer c.Close()

//done := make(chan struct{})

//go func() {
//defer close(done)
//for {
//var res MultiStreamResponse

//err := c.ReadJSON(&res)
//if err != nil {
//log.Println("read:", err)
//return
//}
//log.Printf("recv: %s", res)

//// Unmarshall json to protobuf
//t0 := time.Now()
//var data proto.Message
//eventType := gjson.GetBytes(res.Data, "e").Str
//switch eventType {
//case "trade":
//data = UnmarshalJsonTrade(res.Data)
//case "kline":
//data = UnmarshalJsonKline(res.Data)
//default:
//fmt.Printf("unhandled EventType: %s\n", eventType)
//}

//if data != nil {
//pfData, err := proto.Marshal(data)
//if err != nil {
//fmt.Println(err)
//return
//}

//// write to kafka
//msg := kafka.Message{
//Key:   convertStreamName(res.Stream),
//Value: pfData,
//}

//err = w.WriteMessages(context.Background(), msg)
//t1 := time.Now()
//fmt.Printf("Kafka write (%v)\n", t1.Sub(t0))

//if err != nil {
//log.Fatalln(err)
//}
//}
//}
//}()

//for {
//select {
//case <-done:
//return
//case <-interrupt:
//log.Println("interrupt")

//// Cleanly close the connection by sending a close message and then
//// waiting (with timeout) for the server to close the connection.
//err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
//if err != nil {
//log.Println("write close:", err)
//return
//}
//return
//}
//}
//}

//func UnmarshalJsonKline(raw json.RawMessage) *common.Kline {
//[>
//type KlineStream struct {
//EventType string `json:"e"`
//EventTime int64  `json:"E"`
//Symbol    string `json:"s"`
//Kline     struct {
//StartTime                int64   `json:"t"`
//CloseTime                int64   `json:"T"`
//Symbol                   string  `json:"s"`
//Interval                 string  `json:"i"`
//FirstTradeID             int64   `json:"f"`
//LastTradeID              int64   `json:"L"`
//OpenPrice                float64 `json:"o,string"`
//ClosePrice               float64 `json:"c,string"`
//HighPrice                float64 `json:"h,string"`
//LowPrice                 float64 `json:"l,string"`
//Volume                   float64 `json:"v,string"`
//NumberOfTrades           int64   `json:"n"`
//KlineClosed              bool    `json:"x"`
//Quote                    float64 `json:"q,string"`
//TakerBuyBaseAssetVolume  float64 `json:"V,string"`
//TakerBuyQuoteAssetVolume float64 `json:"Q,string"`
//} `json:"k"`
//*/

//results := gjson.GetMany(
//string(raw),
//"E",   // EventTime
//"s",   // Symbol
//"k.t", // StartTime
//"k.T", // CloseTime
//"k.i", // Interval
//"k.o", // OpenPrice
//"k.c",
//"k.h", // HighPrice
//"k.l",
//"k.v",
//)

////currencyPair := currency.NewPairFromString(results[1].Str)

//return &common.Kline{
//Ts:  common.UnixToTimestampPb(results[3].Int()),
////Pair:       &currencyPair,
////AssetType:  "asdf", //asset.Spot,
////Exchange:   "binance",
//Namespace:  "binance",
////StartTime:  common.UnixToTimestampPb(results[2].Int()),
////CloseTime:  common.UnixToTimestampPb(results[3].Int()),
////Interval:   results[4].Str,
//Open:  results[5].Float(),
//Close: results[6].Float(),
//High:  results[7].Float(),
//Low:   results[8].Float(),
//Volume:     results[9].Float(),
//}
//}

//func UnmarshalJsonTrade(raw json.RawMessage) *common.Trade {
//// Reference: https://github.com/thrasher-corp/gocryptotrader/blob/2c7e531c5c9848e2051afc36f87034631d8fafb7/exchanges/binance/binance_websocket.go
//[>
//EventType      string `json:"e"`
//EventTime      int64  `json:"E"`
//Symbol         string `json:"s"`
//TradeID        int64  `json:"t"`
//Price          string `json:"p"`
//Quantity       string `json:"q"`
//BuyerOrderID   int64  `json:"b"`
//SellerOrderID  int64  `json:"a"`
//TimeStamp      int64  `json:"T"`
//Maker          bool   `json:"m"`
//BestMatchPrice bool   `json:"M"`
//*/
//results := gjson.GetMany(
//string(raw),
//"T",
//"s",
//"p",
//"q",
//)

//currencyPair := currency.NewPairFromString(results[1].Str)

//return &common.Trade{
//Ts:    common.UnixToTimestampPb(results[0].Int()),
//CurrencyPair: &currencyPair,
//AssetType:    "asdf",
//Exchange:     "binance",
//Price:        results[2].Float(),
//Amount:       results[3].Float(),
//}
//}
