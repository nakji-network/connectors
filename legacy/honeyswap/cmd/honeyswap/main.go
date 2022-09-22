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
	conf.SetDefault("honeyswap.kafka.txID", "honeyswap-source")
	conf.SetDefault("honeyswap.factoryAddress", "0xA818b4F111Ccac7AA31D0BCc0806d64F2E0737D7")
	conf.SetDefault("xdai.rpc", []string{
		"https://rpc.xdaichain.com/",
		"wss://rpc.xdaichain.com/wss",
		"wss://xdai.poanetwork.dev/wss",
		//"ws://xdai.poanetwork.dev:8546",
	})
	conf.SetDefault("honeyswap.grpc.port", 50051)
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

	kp, err := kafkautils.NewProducer(conf.GetString("kafka.url"), conf.GetString("honeyswap.kafka.txID"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create new kafka producer")
	}

	connector := uniswapv2.New(uniswapv2.Config{
		Namespace:      "honeyswap",
		TokenNamespace: "xdai",
		KP:             kp,
		RPCURLs:        conf.GetStringSlice("xdai.rpc"),
		LPTopic:        kafkautils.MustParseTopic(conf.GetString("uniswapv2.kafka.topic.liquiditypool_reserve"), conf.GetString("kafka.env")),
		LPEventTopic:   kafkautils.MustParseTopic(conf.GetString("uniswapv2.kafka.topic.liquiditypool_change"), conf.GetString("kafka.env")),
		TradeTopic:     kafkautils.MustParseTopic(conf.GetString("uniswapv2.kafka.topic.market_trade"), conf.GetString("kafka.env")),
		Db:             db,
		FactoryAddress: ethcommon.HexToAddress(conf.GetString("honeyswap.factoryAddress")),
	})

	monitor.StartMonitor("honeyswap")
	go uniswapv2.GrpcServer(conf.GetString("honeyswap.grpc.port"), connector)

	connectorLoaded = true
	connector.Start()
}
