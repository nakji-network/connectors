package main

import (
	"blep.ai/data/config"
	"blep.ai/data/connectors/source/candymachine"
	"blep.ai/data/protoregistry"
	"github.com/nakji-network/connector/kafkautils"
)

var conf = config.GetConfig()

func init() {
	kafkautils.TopicTypeRegistry.Load(candymachine.TopicTypes)
	conf.SetDefault("solana.rpc.url", "https://api.mainnet-beta.solana.com")
	conf.SetDefault("solana.ws.url", "wss://api.mainnet-beta.solana.com")
	conf.SetDefault("solana.maxRetries", 10)
	conf.SetDefault("solana.retryDelay", 500)
	conf.SetDefault("solana.lowPrioRetryDelay", 3)
	conf.SetDefault("solana.lowPrioMaxRetries", 100)
	conf.SetDefault("candymachine.kafka.txID", "candymachine")
	conf.SetDefault("candymachine.id", "candymachine")
	conf.SetDefault("candymachine.programid", "cndy3Z4yapfJBmL3ShUp5exZKqR3z33thTzeNMm2gRZ")
	conf.SetDefault("candymachine.kafka.topic.mintnft", ".fct.nakji.candymachine.0_0_0.nft_mintnft")
	conf.SetDefault("candymachine.kafka.topic.updatecandymachine", ".fct.nakji.candymachine.0_0_0.nft_updatecandymachine")
	conf.SetDefault("candymachine.kafka.topic.addconfiglines", ".fct.nakji.candymachine.0_0_0.nft_addconfiglines")
	conf.SetDefault("candymachine.kafka.topic.initializecandymachine", ".fct.nakji.candymachine.0_0_0.nft_initializecandymachine")
	conf.SetDefault("candymachine.kafka.topic.updateauthority", ".fct.nakji.candymachine.0_0_0.nft_updateauthority")
	conf.SetDefault("candymachine.kafka.topic.withdrawfunds", ".fct.nakji.candymachine.0_0_0.nft_withdrawfunds")
	conf.SetDefault("candymachine.backfillLimit", 250) // Maximum value is 1000
	conf.SetDefault("protoregistry.host", "localhost:9191")
}

func main() {
	topics := map[string]kafkautils.Topic{
		"mintnft":                kafkautils.MustParseTopic(conf.GetString("candymachine.kafka.topic.mintnft"), conf.GetString("kafka.env")),
		"updatecandymachine":     kafkautils.MustParseTopic(conf.GetString("candymachine.kafka.topic.updatecandymachine"), conf.GetString("kafka.env")),
		"addconfiglines":         kafkautils.MustParseTopic(conf.GetString("candymachine.kafka.topic.addconfiglines"), conf.GetString("kafka.env")),
		"initializecandymachine": kafkautils.MustParseTopic(conf.GetString("candymachine.kafka.topic.initializecandymachine"), conf.GetString("kafka.env")),
		"updateauthority":        kafkautils.MustParseTopic(conf.GetString("candymachine.kafka.topic.updateauthority"), conf.GetString("kafka.env")),
		"withdrawfunds":          kafkautils.MustParseTopic(conf.GetString("candymachine.kafka.topic.withdrawfunds"), conf.GetString("kafka.env")),
	}
	connector := candymachine.NewConnector(conf, topics)

	prc := protoregistry.NewClient(conf.GetString("protoregistry.host"))
	connector.Prc = prc

	connector.Start()
}
