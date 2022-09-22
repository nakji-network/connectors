package main

import (
	"context"
	"net/http"

	"blep.ai/data/chain/ethereum/ethclient"
	"blep.ai/data/config"
	"blep.ai/data/connectors/source/starkex"
	"blep.ai/data/database"
	"blep.ai/data/monitor"
	"github.com/nakji-network/connector/kafkautils"

	"cloud.google.com/go/profiler"
	"github.com/heptiolabs/healthcheck"
	"github.com/rs/zerolog/log"
	_ "go.uber.org/automaxprocs"
)

var conf = config.GetConfig()

func init() {
	conf.SetDefault("starkex.kafka.topic.positions", ".fct.nakji.starkex.0_0_0.starkex_position")
	conf.SetDefault("starkex.web3HttpProvider", "https://mainnet.infura.io/v3/7eb8b83da0054b7c8f2694eee7426cf0")
	conf.SetDefault("starkex.interval", 28800) // 8-hour interval default for pulling Perpetual on-chain data
	conf.SetDefault("starkex.kafka.txID", "starkex-source")
	conf.SetDefault("starkex.initialBlock", 13496455)
	conf.SetDefault("starkex.pythonCmd", []string{
		"pipenv",
		"run",
		"python",
		"perpetual_onchain_data.py",
	})
	conf.SetDefault("starkex.outputPath", "output.json")
}

func main() {
	cfg := profiler.Config{
		Service:        "starkex",
		ServiceVersion: "0.0.0",
		// ProjectID must be set if not running on GCP.
		//ProjectID: "my-project",

		// For OpenCensus users:
		// To see Profiler agent spans in APM backend,
		// set EnableOCTelemetry to true
		// EnableOCTelemetry: true,
	}

	// Profiler initialization, best done as early as possible.
	if err := profiler.Start(cfg); err != nil {
		log.Fatal().Err(err).Msg("failed to start GCP profiler")
	}

	// Load Topic registry
	kafkautils.TopicTypeRegistry.Load(starkex.TopicTypes)

	urls := conf.GetStringSlice("ethereum.rpc")
	ethClientPool, err := ethclient.DialPoolContext(context.Background(), urls)
	if err != nil {
		log.Fatal().Err(err).Msg("Ethereum RPC connection error")
	}

	kp, err := kafkautils.NewProducer(conf.GetString("kafka.url"), conf.GetString("starkex.kafka.txID"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create new kafka producer")
	}

	db, err := database.New(conf.GetString("timescaledb.connection"))
	if err != nil {
		log.Fatal().Err(err).Str("dsn", conf.GetString("timescaledb.connection")).Msg("Timescaledb connection failed")
	}
	defer db.Close()

	// For Liveness and Readiness Probe checks
	health := healthcheck.NewHandler()
	go http.ListenAndServe("0.0.0.0:8080", health)
	log.Info().Str("addr", "0.0.0.0:8080").Msg("healthcheck listening on /live and /ready")

	client, err := ethClientPool.RandClient(false)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get random client")
	}

	connector := starkex.NewConnector(
		kp,
		db,
		client,
		kafkautils.MustParseTopic(conf.GetString("starkex.kafka.topic.positions"), conf.GetString("kafka.env")),
		conf.GetString("starkex.web3HttpProvider"),
		conf.GetInt("starkex.interval"),
		conf.GetUint64("starkex.initialBlock"),
	)

	monitor.StartMonitor("starkex")

	ctx := context.Background()
	connector.Start(ctx, conf)
}
