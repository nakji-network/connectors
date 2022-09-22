package main

import (
	"fmt"
	"net/http"

	"blep.ai/data/common"
	"blep.ai/data/config"
	"blep.ai/data/connectors/source/aave"
	"blep.ai/data/database"
	"blep.ai/data/monitor"
	"github.com/nakji-network/connector/kafkautils"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/heptiolabs/healthcheck"
	"github.com/rs/zerolog/log"
)

var conf = config.GetConfig()

func init() {
	conf.SetDefault("aave.kafka.txID", "aave-source")
	conf.SetDefault("aave.factoryAddress", "0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9")
	//conf.SetDefault("aave.kafka.topic.lendingpool", ".fct.lendingpool.0")
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
	conf.SetDefault("aave.grpc.port", 50051)
}

func main() {
	// HEALTHCHECK
	var status *aave.Status
	health := healthcheck.NewHandler()
	// this uses a lot more than 100 goroutines
	//health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))
	health.AddReadinessCheck("connector-loaded", func() error {
		if status == nil || !status.Ready {
			return fmt.Errorf("connector not ready")
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
	kafkautils.TopicTypeRegistry.Load(aave.TopicTypes)
	kafkautils.TopicTypeRegistry.Load(common.TopicTypes)

	// Init Producer
	kp, err := kafkautils.NewProducer(conf.GetString("kafka.url"), conf.GetString("aave.kafka.txID"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create new kafka producer")
	}
	topics := kafkautils.MustParseTopicsMap(aave.TopicMap, conf.GetString("kafka.env"))

	connector := aave.New(aave.Config{
		Namespace:      "aave",
		KP:             kp,
		RPCURLs:        conf.GetStringSlice("ethereum.rpc"),
		Db:             db,
		Topics:         topics,
		FactoryAddress: ethcommon.HexToAddress(conf.GetString("aave.factoryAddress")),
	})

	status = connector.Status

	// TODO Optional: Support AAVE custom status checks (e.g. status == false for anything)
	monitor.StartMonitor("aave")
	connector.Start()
}
