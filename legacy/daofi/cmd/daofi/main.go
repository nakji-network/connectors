package main

import (
	"fmt"
	"net/http"

	"blep.ai/data/monitor"

	"blep.ai/data/common"
	"blep.ai/data/config"
	"blep.ai/data/connectors/source/daofi"
	"blep.ai/data/connectors/source/uniswapv2"
	"blep.ai/data/database"
	"github.com/nakji-network/connector/kafkautils"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/heptiolabs/healthcheck"
	"github.com/rs/zerolog/log"
)

var conf = config.GetConfig()

func init() {
	conf.SetDefault("daofi.kafka.topic.liquiditypool_reserve", ".fct.nakji.common.0_0_0.liquiditypool_reserve")
	conf.SetDefault("daofi.kafka.topic.liquiditypool_change", ".fct.nakji.common.0_0_0.liquiditypool_change")
	conf.SetDefault("daofi.kafka.topic.market_trade", ".fct.nakji.common.0_0_0.market_trade")
	conf.SetDefault("daofi.kafka.txID", "daofi-source")
	conf.SetDefault("daofi.factoryAddress", "0xEaC9260C59693f180936779451B996b303a0A488")
	conf.SetDefault("ethereum.rpc", []string{
		//"https://sparkling-wild-water.quiknode.pro/886d6e59b11a916a403259ad83de327d15ad0c33/",
		//"wss://sparkling-wild-water.quiknode.pro/886d6e59b11a916a403259ad83de327d15ad0c33/",
		"https://cocky-hugle:cosmic-snack-lion-front-sneer-come@nd-653-072-430.p2pify.com",
		"wss://cocky-hugle:cosmic-snack-lion-front-sneer-come@ws-nd-653-072-430.p2pify.com",
		"https://hopeful-jackson:pledge-tissue-rocky-hurray-canal-mutual@nd-557-839-247.p2pify.com",
		"wss://hopeful-jackson:pledge-tissue-rocky-hurray-canal-mutual@ws-nd-557-839-247.p2pify.com",
		"https://eth-mainnet.alchemyapi.io/v2/nB37zoZNMRPGRVgDJ3s0UBTTViSxUhk5",
		"wss://eth-mainnet.ws.alchemyapi.io/v2/nB37zoZNMRPGRVgDJ3s0UBTTViSxUhk5",
	})
	conf.SetDefault("daofi.grpc.port", 50051)
}

func main() {
	// Connect Database
	db, err := database.New(conf.GetString("timescaledb.connection"))
	if err != nil {
		log.Fatal().Err(err).Str("dsn", conf.GetString("timescaledb.connection")).Msg("Timescaledb connection failed")
	}
	defer db.Close()

	// Load Topic registry
	kafkautils.TopicTypeRegistry.Load(uniswapv2.TopicTypes)
	kafkautils.TopicTypeRegistry.Load(common.TopicTypes)

	// Init Producer
	kp, err := kafkautils.NewProducer(conf.GetString("kafka.url"), conf.GetString("daofi.kafka.txID"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create new kafka producer")
	}

	connector := daofi.New(daofi.Config{
		Namespace:      "daofi",
		TokenNamespace: "ethereum",
		KP:             kp,
		RPCURLs:        conf.GetStringSlice("ethereum.rpc"),
		LPTopic:        kafkautils.MustParseTopic(conf.GetString("daofi.kafka.topic.liquiditypool_reserve"), conf.GetString("kafka.env")),
		LPEventTopic:   kafkautils.MustParseTopic(conf.GetString("daofi.kafka.topic.liquiditypool_change"), conf.GetString("kafka.env")),
		TradeTopic:     kafkautils.MustParseTopic(conf.GetString("daofi.kafka.topic.market_trade"), conf.GetString("kafka.env")),
		Db:             db,
		FactoryAddress: ethcommon.HexToAddress(conf.GetString("daofi.factoryAddress")),
	})

	monitor.StartMonitor("daofi")
	go daofi.GrpcServer(conf.GetString("daofi.grpc.port"), connector)

	// HEALTHCHECK
	health := healthcheck.NewHandler()
	// this uses a lot more than 100 goroutines
	//health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))
	health.AddReadinessCheck("connector-loaded", func() error {
		if !connector.Status.Ready {
			return fmt.Errorf("connector not ready")
		}
		return nil
	})
	go http.ListenAndServe("0.0.0.0:8080", health)
	log.Info().Str("addr", "0.0.0.0:8080").Msg("healthcheck listening on /live and /ready")

	connector.Start()
}
