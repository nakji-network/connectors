// Poll Google Trends for data.
// References:
// - https://github.com/groovili/gogtrends/tree/694001cbd88873cee94b9cad76d135860790155e
// - https://towardsdatascience.com/reconstruct-google-trends-daily-data-for-extended-period-75b6ca1d3420
// - https://github.com/qztseng/google-trends-daily/blob/master/gtrend.py
package main

import (
	"blep.ai/data/config"
	"blep.ai/data/connectors/source/googletrends"
	"blep.ai/data/database"
	"blep.ai/data/monitor"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/rs/zerolog/log"
)

var conf = config.GetConfig()

func init() {
	conf.SetDefault("googletrends.baseURL", "https://trends.google.com/trends/api/")
	conf.SetDefault("googletrends.keywords", []string{
		"btc",
		"bitcoin",
		"eth",
		"ethereum",
		"blockchain",
		"defi",
		"crypto",
		"cryptocurrency",
		"uniswap",
		"coinbase",
		"binance",
		"kraken",
		"okex",
		"huobi",
		"ftx",
		"tradingview",
		"chainlink",
		"xrp",
		"dextools",
		"glassnode",
		"skew",
		"yfi",
		"yearn",
		"coinmarketcap",
		"coingecko",
		"synthetix",
		"aave",
		//"非小号",
		"비트 코인",
		"이더리움",
		"코인원",
		//"업비트", // not enough
		"区块链",
		"比特币",
		"以太坊",
		"数字货币",
		"DCEP",
		"digital currency",
		"imtoken",
		"metamask",
		"bitcoin mining",
		//"chainnews",
		//"http://www.marsfinance.net/ 火星财经"

	})
	conf.SetDefault("googletrends.transactionalID", "googletrends-source")
	conf.SetDefault("googletrends.trendTopic", ".fct.nakji.google.0_0_0.trend_score")
	conf.SetDefault("googletrends.namespace", "google")
	conf.SetDefault("googletrends.table", "nakji.google.0_0_0.trend_score")
	conf.SetDefault("googletrends.sleep", "1m30s")
	conf.SetDefault("googletrends.interval", "60m")
	conf.SetDefault("googletrends.Backoff", "1m")
}

func main() {
	kafkautils.TopicTypeRegistry.Load(googletrends.TopicTypes)

	// Connect to Kafka Producer
	kp, err := kafkautils.NewProducer(conf.GetString("kafka.url"), conf.GetString("googletrends.transactionalID"))
	if err != nil {
		log.Fatal().Err(err).Msg("Kafka producer failed init")
	}
	kp.EnableTransactions()

	monitor.StartMonitor("googletrends")

	// Connect to database
	db, err := database.New(conf.GetString("timescaledb.connection"))
	if err != nil {
		log.Fatal().Err(err).Msg("Database failed to connect")
	}
	defer db.Close()

	connector := googletrends.New(googletrends.Config{
		BaseURL:    conf.GetString("googletrends.baseURL"),
		Keywords:   conf.GetStringSlice("googletrends.keywords"),
		NS:         conf.GetString("googletrends.namespace"),
		TrendTopic: kafkautils.MustParseTopic(conf.GetString("googletrends.trendTopic"), conf.GetString("kafka.env")),
		KP:         kp,
		DB:         db,
		Table:      conf.GetString("googletrends.table"),
		Sleep:      conf.GetDuration("googletrends.sleep"),
		Interval:   conf.GetDuration("googletrends.interval"),
		Backoff:    conf.GetDuration("googletrends.Backoff"),
	})

	connector.Start()

}
