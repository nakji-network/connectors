package main

import (
	"blep.ai/data/common"
	"blep.ai/data/config"
	"blep.ai/data/connectors/source/uniswapv2"
	"blep.ai/data/database"
	"blep.ai/data/monitor"
	"github.com/nakji-network/connector/kafkautils"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
)

var conf = config.GetConfig()

func init() {
	conf.SetDefault("uniswapv2.kafka.topic.liquiditypool_reserve", ".fct.nakji.common.0_0_0.liquiditypool_reserve")
	conf.SetDefault("uniswapv2.kafka.topic.liquiditypool_change", ".fct.nakji.common.0_0_0.liquiditypool_change")
	conf.SetDefault("uniswapv2.kafka.topic.market_trade", ".fct.nakji.common.0_0_0.market_trade")
	conf.SetDefault("bscswap.kafka.txID", "bscswap-source")
	conf.SetDefault("bscswap.factoryAddress", "0xCe8fd65646F2a2a897755A1188C04aCe94D2B8D0")
	conf.SetDefault("bscswap.grpc.port", 50051)
}

func main() {
	db, err := database.New(conf.GetString("timescaledb.connection"))
	if err != nil {
		log.Fatal().Err(err).Str("dsn", conf.GetString("timescaledb.connection")).Msg("Timescaledb connection failed")
	}
	defer db.Close()

	// Load Topic registry
	kafkautils.TopicTypeRegistry.Load(uniswapv2.TopicTypes)
	kafkautils.TopicTypeRegistry.Load(common.TopicTypes)

	kp, err := kafkautils.NewProducer(conf.GetString("kafka.url"), conf.GetString("bscswap.kafka.txID"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create new kafka producer")
	}

	monitor.StartMonitor("bscswap")

	connector := uniswapv2.New(uniswapv2.Config{
		Namespace:      "bscswap",
		TokenNamespace: "bsc",
		KP:             kp,
		RPCURLs:        conf.GetStringSlice("bsc.rpc"),
		LPTopic:        kafkautils.MustParseTopic(conf.GetString("uniswapv2.kafka.topic.liquiditypool_reserve"), conf.GetString("kafka.env")),
		LPEventTopic:   kafkautils.MustParseTopic(conf.GetString("uniswapv2.kafka.topic.liquiditypool_change"), conf.GetString("kafka.env")),
		TradeTopic:     kafkautils.MustParseTopic(conf.GetString("uniswapv2.kafka.topic.market_trade"), conf.GetString("kafka.env")),
		Db:             db,
		FactoryAddress: ethcommon.HexToAddress(conf.GetString("bscswap.factoryAddress")),
	})

	monitor.StartMonitor("bscswap")
	go uniswapv2.GrpcServer(conf.GetString("bscswap.grpc.port"), connector)

	connector.Start()
}
