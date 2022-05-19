package main

import (
	"fmt"
	"net/http"

	"blep.ai/data/monitor"

	"blep.ai/data/common"
	"blep.ai/data/config"
	"blep.ai/data/connectors/source/uniswapv2"
	"blep.ai/data/database"
	"github.com/nakji-network/connector/kafkautils"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/heptiolabs/healthcheck"
	"github.com/rs/zerolog/log"
)

var conf = config.GetConfig()

func init() {
	conf.SetDefault("uniswapv2.kafka.topic.liquiditypool_reserve", ".fct.nakji.common.0_0_0.liquiditypool_reserve")
	conf.SetDefault("uniswapv2.kafka.topic.liquiditypool_change", ".fct.nakji.common.0_0_0.liquiditypool_change")
	conf.SetDefault("uniswapv2.kafka.topic.market_trade", ".fct.nakji.common.0_0_0.market_trade")
	conf.SetDefault("mdex.kafka.txID", "mdex-source")
	conf.SetDefault("mdex.factoryAddress", "0xb0b670fc1f7724119963018db0bfa86adb22d941")
	conf.SetDefault("heco.rpc", []string{
		"https://http-mainnet-node.huobichain.com",
		"wss://ws-mainnet-node.huobichain.com",
	})
	conf.SetDefault("mdex.grpc.port", 50051)
}

func main() {
	// HEALTHCHECK
	connectorLoaded := false
	health := healthcheck.NewHandler()
	//health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))
	health.AddReadinessCheck("connector-loaded", func() error {
		if !connectorLoaded {
			return fmt.Errorf("connector not loaded")
		}
		return nil
	})
	go http.ListenAndServe("0.0.0.0:8080", health)
	log.Info().Str("addr", "0.0.0.0:8080").Msg("healthcheck listening on /live and /ready")

	// Connect Database
	db, err := database.New(conf.GetString("timescaledb.connection"))
	if err != nil {
		log.Fatal().Err(err).Str("dsn", conf.GetString("timescaledb.connection")).Msg("Timescaledb connection failed")
	}
	defer db.Close()

	// Load Topic registry
	kafkautils.TopicTypeRegistry.Load(uniswapv2.TopicTypes)
	kafkautils.TopicTypeRegistry.Load(common.TopicTypes)

	kp, err := kafkautils.NewProducer(conf.GetString("kafka.url"), conf.GetString("mdex.kafka.txID"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create new kafka producer")
	}

	connector := uniswapv2.New(uniswapv2.Config{
		KP:             kp,
		Db:             db,
		LPTopic:        kafkautils.MustParseTopic(conf.GetString("uniswapv2.kafka.topic.liquiditypool_reserve"), conf.GetString("kafka.env")),
		LPEventTopic:   kafkautils.MustParseTopic(conf.GetString("uniswapv2.kafka.topic.liquiditypool_change"), conf.GetString("kafka.env")),
		TradeTopic:     kafkautils.MustParseTopic(conf.GetString("uniswapv2.kafka.topic.market_trade"), conf.GetString("kafka.env")),
		Namespace:      "mdex",
		TokenNamespace: "heco",
		RPCURLs:        conf.GetStringSlice("heco.rpc"),
		FactoryAddress: ethcommon.HexToAddress(conf.GetString("mdex.factoryAddress")),
	})

	monitor.StartMonitor("mdex")
	go uniswapv2.GrpcServer(conf.GetString("mdex.grpc.port"), connector)

	connectorLoaded = true
	connector.Start()
}
