package main

import (
	"blep.ai/data/config"
	"blep.ai/data/connectors/source/solanatoken"
	"github.com/nakji-network/connector/kafkautils"
)

var conf = config.GetConfig()

func init() {
	kafkautils.TopicTypeRegistry.Load(solanatoken.TopicTypes)
	conf.SetDefault("solana.rpc.url", "https://api.mainnet-beta.solana.com")
	conf.SetDefault("solana.ws.url", "wss://api.mainnet-beta.solana.com")
	conf.SetDefault("solana.maxRetries", 10)
	conf.SetDefault("solana.retryDelay", 500)
	conf.SetDefault("solana.lowPrioRetryDelay", 100)
	conf.SetDefault("solana.lowPrioMaxRetries", 2)
	conf.SetDefault("solanatoken.kafka.txID", "solanatoken")
	conf.SetDefault("solanatoken.id", "solanatoken")
	conf.SetDefault("solanatoken.tokenAddress", "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	conf.SetDefault("solanatoken.kafka.topic.tokentransfer", ".fct.nakji.solanatoken.0_0_0.token_transfer")
	conf.SetDefault("solanatoken.kafka.topic.tokeninitializemint", ".fct.nakji.solanatoken.0_0_0.token_initializemint")
	conf.SetDefault("solanatoken.kafka.topic.tokeninitializeaccount", ".fct.nakji.solanatoken.0_0_0.token_initializeaccount")
	conf.SetDefault("solanatoken.kafka.topic.tokenmintto", ".fct.nakji.solanatoken.0_0_0.token_mintto")
	conf.SetDefault("solanatoken.kafka.topic.tokensetauthority", ".fct.nakji.solanatoken.0_0_0.token_setauthority")
	conf.SetDefault("solanatoken.kafka.topic.tokencloseaccount", ".fct.nakji.solanatoken.0_0_0.token_closeaccount")
	conf.SetDefault("solanatoken.kafka.topic.tokenburn", ".fct.nakji.solanatoken.0_0_0.token_burn")
	conf.SetDefault("solanatoken.kafka.topic.tokenapprove", ".fct.nakji.solanatoken.0_0_0.token_approve")
	conf.SetDefault("solanatoken.kafka.topic.tokenrevoke", ".fct.nakji.solanatoken.0_0_0.token_revoke")
	conf.SetDefault("solanatoken.kafka.topic.tokenfreezeaccount", ".fct.nakji.solanatoken.0_0_0.token_freezeaccount")
	conf.SetDefault("solanatoken.kafka.topic.tokeninitializemultisig", ".fct.nakji.solanatoken.0_0_0.token_initializemultisig")
	conf.SetDefault("solanatoken.kafka.topic.tokensyncnative", ".fct.nakji.solanatoken.0_0_0.token_syncnative")
	conf.SetDefault("solanatoken.kafka.topic.tokenthawaccount", ".fct.nakji.solanatoken.0_0_0.token_thawaccount")
	conf.SetDefault("solanatoken.kafka.topic.tokentransferchecked", ".fct.nakji.solanatoken.0_0_0.token_transferchecked")
	conf.SetDefault("solanatoken.kafka.topic.tokenapprovechecked", ".fct.nakji.solanatoken.0_0_0.token_approvechecked")
	conf.SetDefault("solanatoken.kafka.topic.tokenminttochecked", ".fct.nakji.solanatoken.0_0_0.token_minttochecked")
	conf.SetDefault("solanatoken.kafka.topic.tokenburnchecked", ".fct.nakji.solanatoken.0_0_0.token_burnchecked")
	conf.SetDefault("solanatoken.kafka.topic.tokeninitializeaccount2", ".fct.nakji.solanatoken.0_0_0.token_initializeaccount2")
	conf.SetDefault("solanatoken.kafka.topic.tokeninitializeaccount3", ".fct.nakji.solanatoken.0_0_0.token_initializeaccount3")
	conf.SetDefault("solanatoken.kafka.topic.tokeninitializemint2", ".fct.nakji.solanatoken.0_0_0.token_initializemint2")
	conf.SetDefault("solanatoken.kafka.topic.tokeninitializemultisig2", ".fct.nakji.solanatoken.0_0_0.token_initializemultisig2")
	conf.SetDefault("solanatoken.kafka.topic.tokennft", ".fct.nakji.solanatoken.0_0_0.token_nft")
	conf.SetDefault("solanatoken.kafka.topic.tokentrade", ".fct.nakji.solanatoken.0_0_0.token_trade")
	conf.SetDefault("solanatoken.backfillLimit", 250) // Maximum value is 1000
}

func main() {
	topics := map[string]kafkautils.Topic{
		"tokentrasfer":             kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokentransfer"), conf.GetString("kafka.env")),
		"tokeninitializemint":      kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokeninitializemint"), conf.GetString("kafka.env")),
		"tokeninitializeaccount":   kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokeninitializeaccount"), conf.GetString("kafka.env")),
		"tokenmintto":              kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokenmintto"), conf.GetString("kafka.env")),
		"tokensetauthority":        kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokensetauthority"), conf.GetString("kafka.env")),
		"tokencloseaccount":        kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokencloseaccount"), conf.GetString("kafka.env")),
		"tokenburn":                kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokenburn"), conf.GetString("kafka.env")),
		"tokenapprove":             kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokenapprove"), conf.GetString("kafka.env")),
		"tokenrevoke":              kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokenrevoke"), conf.GetString("kafka.env")),
		"tokenfreezeaccount":       kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokenfreezeaccount"), conf.GetString("kafka.env")),
		"tokeninitializemultisig":  kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokeninitializemultisig"), conf.GetString("kafka.env")),
		"tokensyncnative":          kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokensyncnative"), conf.GetString("kafka.env")),
		"tokenthawaccount":         kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokenthawaccount"), conf.GetString("kafka.env")),
		"tokentransferchecked":     kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokentransferchecked"), conf.GetString("kafka.env")),
		"tokenapprovechecked":      kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokenapprovechecked"), conf.GetString("kafka.env")),
		"tokenminttochecked":       kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokenminttochecked"), conf.GetString("kafka.env")),
		"tokenburnchecked":         kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokenburnchecked"), conf.GetString("kafka.env")),
		"tokeninitializeaccount2":  kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokeninitializeaccount2"), conf.GetString("kafka.env")),
		"tokeninitializeaccount3":  kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokeninitializeaccount3"), conf.GetString("kafka.env")),
		"tokeninitializemint2":     kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokeninitializemint2"), conf.GetString("kafka.env")),
		"tokeninitializemultisig2": kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokeninitializemultisig2"), conf.GetString("kafka.env")),
		"tokennft":                 kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokennft"), conf.GetString("kafka.env")),
		"tokentrade":               kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokentrade"), conf.GetString("kafka.env")),
		// "tokennftbackfill":         kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokennftbackfill"), conf.GetString("kafka.env")),
		// "tokentradebackfill":       kafkautils.MustParseTopic(conf.GetString("solanatoken.kafka.topic.tokentradebackfill"), conf.GetString("kafka.env")),
	}

	connector := solanatoken.NewConnector(conf, topics)
	connector.Start()
}
