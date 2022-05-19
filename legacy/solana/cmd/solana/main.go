package main

import (
	"context"

	"blep.ai/data/config"
	"blep.ai/data/connectors/source/solana"
	"blep.ai/data/connectors/source/solana/chain"
	"blep.ai/data/database"
	"blep.ai/data/monitor"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/rs/zerolog/log"
)

var conf = config.GetConfig()

func init() {
	kafkautils.TopicTypeRegistry.Load(chain.TopicTypes)
	conf.SetDefault("solana.rpc.url", "https://api.devnet.solana.com")
	conf.SetDefault("solana.ws.url", "wss://api.devnet.solana.com")
	conf.SetDefault("solana.kafka.topic.solanaaccount", ".fct.nakji.solana.0_0_0.solana_account")
	conf.SetDefault("solana.kafka.topic.solanblock", ".fct.nakji.solana.0_0_0.solana_block")
	conf.SetDefault("solana.kafka.topic.solanamessage", ".fct.nakji.solana.0_0_0.solana_message")
	conf.SetDefault("solana.kafka.topic.solanareward", ".fct.nakji.solana.0_0_0.solana_reward")
	conf.SetDefault("solana.kafka.topic.solanatransaction", ".fct.nakji.solana.0_0_0.solana_transaction")
	conf.SetDefault("solana.kafka.txID", "solana")
	conf.SetDefault("solana.rpcErrorsTable", `"nakji.solana.0_0_0.rpc_errors"`)
	conf.SetDefault("solana.maxRetries", 10)
	conf.SetDefault("solana.retryDelay", 500)
}

func main() {
	wsClient, err := ws.Connect(context.Background(), conf.GetString("solana.ws.url"))
	if err != nil {
		log.Fatal().Err(err).Msg("Solana: WS Client connection error")
	}

	client := rpc.New(conf.GetString("solana.rpc.url"))

	kp, err := kafkautils.NewProducer(conf.GetString("kafka.url"), conf.GetString("solana.kafka.txID"))
	if err != nil {
		log.Fatal().Err(err).Msg("BTC New Producer failing")
	}

	db, err := database.New(conf.GetString("timescaledb.connection"))
	if err != nil {
		log.Fatal().Err(err).Str("dsn", conf.GetString("timescaledb.connection")).Msg("Timescaledb connection failed")
	}
	defer db.Close()

	topics := map[string]kafkautils.Topic{
		"solanaaccount":     kafkautils.MustParseTopic(conf.GetString("solana.kafka.topic.solanaaccount"), conf.GetString("kafka.env")),
		"solanablock":       kafkautils.MustParseTopic(conf.GetString("solana.kafka.topic.solanablock"), conf.GetString("kafka.env")),
		"solanamessage":     kafkautils.MustParseTopic(conf.GetString("solana.kafka.topic.solanamessage"), conf.GetString("kafka.env")),
		"solanareward":      kafkautils.MustParseTopic(conf.GetString("solana.kafka.topic.solanareward"), conf.GetString("kafka.env")),
		"solanatransaction": kafkautils.MustParseTopic(conf.GetString("solana.kafka.topic.solanatransaction"), conf.GetString("kafka.env")),
	}

	monitor.StartMonitor("solana")
	newDb := &database.Db{DbInterface: db}
	connector := solana.NewConnector(client, wsClient, topics, kp, conf, newDb)
	connector.Start(context.Background())
}
