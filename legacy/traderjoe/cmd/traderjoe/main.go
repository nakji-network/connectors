// Code generated by connectorgen - Edit as necessary.

package main

import (
	"context"
	"net/http"

	"blep.ai/data/chain/avalanche/avaxclient"
	"blep.ai/data/chain/ethereum/ethclient"
	"blep.ai/data/config"
	"blep.ai/data/connectors/source/traderjoe"
	"blep.ai/data/database"
	"blep.ai/data/monitor"
	"blep.ai/data/tokencache"
	"github.com/nakji-network/connector/kafkautils"

	"cloud.google.com/go/profiler"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/heptiolabs/healthcheck"
	"github.com/rs/zerolog/log"
	_ "go.uber.org/automaxprocs"
)

var conf = config.InitConfig()

func init() {
	kafkautils.TopicTypeRegistry.Load(traderjoe.TopicTypes)
	conf.SetDefault("traderjoe.kafka.topic.joebar_approval", ".fct.nakji.traderjoe.0_0_0.joebar_approval")
	conf.SetDefault("traderjoe.kafka.topic.joebar_transfer", ".fct.nakji.traderjoe.0_0_0.joebar_transfer")
	conf.SetDefault("traderjoe.kafka.topic.joefactory_paircreated", ".fct.nakji.traderjoe.0_0_0.joefactory_paircreated")
	conf.SetDefault("traderjoe.kafka.topic.joehattoken_approval", ".fct.nakji.traderjoe.0_0_0.joehattoken_approval")
	conf.SetDefault("traderjoe.kafka.topic.joehattoken_transfer", ".fct.nakji.traderjoe.0_0_0.joehattoken_transfer")
	conf.SetDefault("traderjoe.kafka.topic.joepair_approval", ".fct.nakji.traderjoe.0_0_0.joepair_approval")
	conf.SetDefault("traderjoe.kafka.topic.joepair_burn", ".fct.nakji.traderjoe.0_0_0.joepair_burn")
	conf.SetDefault("traderjoe.kafka.topic.joepair_mint", ".fct.nakji.traderjoe.0_0_0.joepair_mint")
	conf.SetDefault("traderjoe.kafka.topic.joepair_swap", ".fct.nakji.traderjoe.0_0_0.joepair_swap")
	conf.SetDefault("traderjoe.kafka.topic.joepair_sync", ".fct.nakji.traderjoe.0_0_0.joepair_sync")
	conf.SetDefault("traderjoe.kafka.topic.joepair_transfer", ".fct.nakji.traderjoe.0_0_0.joepair_transfer")
	conf.SetDefault("traderjoe.kafka.topic.joetoken_approval", ".fct.nakji.traderjoe.0_0_0.joetoken_approval")
	conf.SetDefault("traderjoe.kafka.topic.joetoken_delegatechanged", ".fct.nakji.traderjoe.0_0_0.joetoken_delegatechanged")
	conf.SetDefault("traderjoe.kafka.topic.joetoken_delegatevoteschanged", ".fct.nakji.traderjoe.0_0_0.joetoken_delegatevoteschanged")
	conf.SetDefault("traderjoe.kafka.topic.joetoken_ownershiptransferred", ".fct.nakji.traderjoe.0_0_0.joetoken_ownershiptransferred")
	conf.SetDefault("traderjoe.kafka.topic.joetoken_transfer", ".fct.nakji.traderjoe.0_0_0.joetoken_transfer")
	conf.SetDefault("traderjoe.kafka.topic.joetroller_failure", ".fct.nakji.traderjoe.0_0_0.joetroller_failure")
	conf.SetDefault("traderjoe.kafka.topic.joetroller_newadmin", ".fct.nakji.traderjoe.0_0_0.joetroller_newadmin")
	conf.SetDefault("traderjoe.kafka.topic.joetroller_newimplementation", ".fct.nakji.traderjoe.0_0_0.joetroller_newimplementation")
	conf.SetDefault("traderjoe.kafka.topic.joetroller_newpendingadmin", ".fct.nakji.traderjoe.0_0_0.joetroller_newpendingadmin")
	conf.SetDefault("traderjoe.kafka.topic.joetroller_newpendingimplementation", ".fct.nakji.traderjoe.0_0_0.joetroller_newpendingimplementation")
	conf.SetDefault("traderjoe.kafka.topic.masterchefjoev2_add", ".fct.nakji.traderjoe.0_0_0.masterchefjoev2_add")
	conf.SetDefault("traderjoe.kafka.topic.masterchefjoev2_deposit", ".fct.nakji.traderjoe.0_0_0.masterchefjoev2_deposit")
	conf.SetDefault("traderjoe.kafka.topic.masterchefjoev2_emergencywithdraw", ".fct.nakji.traderjoe.0_0_0.masterchefjoev2_emergencywithdraw")
	conf.SetDefault("traderjoe.kafka.topic.masterchefjoev2_harvest", ".fct.nakji.traderjoe.0_0_0.masterchefjoev2_harvest")
	conf.SetDefault("traderjoe.kafka.topic.masterchefjoev2_ownershiptransferred", ".fct.nakji.traderjoe.0_0_0.masterchefjoev2_ownershiptransferred")
	conf.SetDefault("traderjoe.kafka.topic.masterchefjoev2_set", ".fct.nakji.traderjoe.0_0_0.masterchefjoev2_set")
	conf.SetDefault("traderjoe.kafka.topic.masterchefjoev2_setdevaddress", ".fct.nakji.traderjoe.0_0_0.masterchefjoev2_setdevaddress")
	conf.SetDefault("traderjoe.kafka.topic.masterchefjoev2_updateemissionrate", ".fct.nakji.traderjoe.0_0_0.masterchefjoev2_updateemissionrate")
	conf.SetDefault("traderjoe.kafka.topic.masterchefjoev2_updatepool", ".fct.nakji.traderjoe.0_0_0.masterchefjoev2_updatepool")
	conf.SetDefault("traderjoe.kafka.topic.masterchefjoev2_withdraw", ".fct.nakji.traderjoe.0_0_0.masterchefjoev2_withdraw")
	conf.SetDefault("traderjoe.kafka.topic.masterchefjoev3_add", ".fct.nakji.traderjoe.0_0_0.masterchefjoev3_add")
	conf.SetDefault("traderjoe.kafka.topic.masterchefjoev3_deposit", ".fct.nakji.traderjoe.0_0_0.masterchefjoev3_deposit")
	conf.SetDefault("traderjoe.kafka.topic.masterchefjoev3_emergencywithdraw", ".fct.nakji.traderjoe.0_0_0.masterchefjoev3_emergencywithdraw")
	conf.SetDefault("traderjoe.kafka.topic.masterchefjoev3_harvest", ".fct.nakji.traderjoe.0_0_0.masterchefjoev3_harvest")
	conf.SetDefault("traderjoe.kafka.topic.masterchefjoev3_init", ".fct.nakji.traderjoe.0_0_0.masterchefjoev3_init")
	conf.SetDefault("traderjoe.kafka.topic.masterchefjoev3_ownershiptransferred", ".fct.nakji.traderjoe.0_0_0.masterchefjoev3_ownershiptransferred")
	conf.SetDefault("traderjoe.kafka.topic.masterchefjoev3_set", ".fct.nakji.traderjoe.0_0_0.masterchefjoev3_set")
	conf.SetDefault("traderjoe.kafka.topic.masterchefjoev3_updatepool", ".fct.nakji.traderjoe.0_0_0.masterchefjoev3_updatepool")
	conf.SetDefault("traderjoe.kafka.topic.masterchefjoev3_withdraw", ".fct.nakji.traderjoe.0_0_0.masterchefjoev3_withdraw")

	conf.SetDefault("traderjoe.kafka.txID", "traderjoe")

	conf.SetDefault("traderjoe.joebarAddress", "0x57319d41F71E81F3c65F2a47CA4e001EbAFd4F33")
	conf.SetDefault("traderjoe.joefactoryAddress", "0x9Ad6C38BE94206cA50bb0d90783181662f0Cfa10")
	conf.SetDefault("traderjoe.joehattokenAddress", "0x82FE038Ea4b50f9C957da326C412ebd73462077C")

	conf.SetDefault("traderjoe.joetokenAddress", "0x6e84a6216eA6dACC71eE8E6b0a5B7322EEbC0fDd")
	conf.SetDefault("traderjoe.joetrollerAddress", "0xdc13687554205E5b89Ac783db14bb5bba4A1eDaC")
	conf.SetDefault("traderjoe.masterchefjoev2Address", "0xd6a4F121CA35509aF06A0Be99093d08462f53052")
	conf.SetDefault("traderjoe.masterchefjoev3Address", "0x188bED1968b795d5c9022F6a0bb5931Ac4c18F00")
}

func main() {
	cfg := profiler.Config{
		Service:        "traderjoe",
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

	kp, err := kafkautils.NewProducer(conf.GetString("kafka.url"), conf.GetString("traderjoe.kafka.txID"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create new kafka producer")
	}

	urls := conf.GetStringSlice("rpcs.avalanche.full")

	log.Info().Strs("url", urls).Msg("connecting to Avalanche RPC")

	ethClientPool, err := ethclient.DialPoolContext(context.Background(), urls)
	if err != nil {
		log.Fatal().Err(err).Msg("Ethereum RPC connection error")
	}

	avaxClientPool, err := avaxclient.DialPoolContext(context.Background(), urls)
	if err != nil {
		log.Fatal().Err(err).Msg("Ethereum RPC connection error")
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
	joebarAddress := ethcommon.HexToAddress(conf.GetString("traderjoe.joebarAddress"))
	joefactoryAddress := ethcommon.HexToAddress(conf.GetString("traderjoe.joefactoryAddress"))
	joehattokenAddress := ethcommon.HexToAddress(conf.GetString("traderjoe.joehattokenAddress"))
	joetokenAddress := ethcommon.HexToAddress(conf.GetString("traderjoe.joetokenAddress"))
	joetrollerAddress := ethcommon.HexToAddress(conf.GetString("traderjoe.joetrollerAddress"))
	masterchefjoev2Address := ethcommon.HexToAddress(conf.GetString("traderjoe.masterchefjoev2Address"))
	masterchefjoev3Address := ethcommon.HexToAddress(conf.GetString("traderjoe.masterchefjoev3Address"))

	addresses := map[string][]ethcommon.Address{
		"joebar":      {joebarAddress},
		"joefactory":  {joefactoryAddress},
		"joehattoken": {joehattokenAddress},
		"joepair":     traderjoe.GetJoePairAddresses(db),

		"joetoken":        {joetokenAddress},
		"joetroller":      {joetrollerAddress},
		"masterchefjoev2": {masterchefjoev2Address},
		"masterchefjoev3": {masterchefjoev3Address},
	}

	topics := map[string]kafkautils.Topic{
		"joebar_approval":                      kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joebar_approval"), conf.GetString("kafka.env")),
		"joebar_transfer":                      kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joebar_transfer"), conf.GetString("kafka.env")),
		"joefactory_paircreated":               kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joefactory_paircreated"), conf.GetString("kafka.env")),
		"joehattoken_approval":                 kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joehattoken_approval"), conf.GetString("kafka.env")),
		"joehattoken_transfer":                 kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joehattoken_transfer"), conf.GetString("kafka.env")),
		"joepair_approval":                     kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joepair_approval"), conf.GetString("kafka.env")),
		"joepair_burn":                         kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joepair_burn"), conf.GetString("kafka.env")),
		"joepair_mint":                         kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joepair_mint"), conf.GetString("kafka.env")),
		"joepair_swap":                         kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joepair_swap"), conf.GetString("kafka.env")),
		"joepair_sync":                         kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joepair_sync"), conf.GetString("kafka.env")),
		"joepair_transfer":                     kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joepair_transfer"), conf.GetString("kafka.env")),
		"joetoken_approval":                    kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joetoken_approval"), conf.GetString("kafka.env")),
		"joetoken_delegatechanged":             kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joetoken_delegatechanged"), conf.GetString("kafka.env")),
		"joetoken_delegatevoteschanged":        kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joetoken_delegatevoteschanged"), conf.GetString("kafka.env")),
		"joetoken_ownershiptransferred":        kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joetoken_ownershiptransferred"), conf.GetString("kafka.env")),
		"joetoken_transfer":                    kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joetoken_transfer"), conf.GetString("kafka.env")),
		"joetroller_failure":                   kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joetroller_failure"), conf.GetString("kafka.env")),
		"joetroller_newadmin":                  kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joetroller_newadmin"), conf.GetString("kafka.env")),
		"joetroller_newimplementation":         kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joetroller_newimplementation"), conf.GetString("kafka.env")),
		"joetroller_newpendingadmin":           kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joetroller_newpendingadmin"), conf.GetString("kafka.env")),
		"joetroller_newpendingimplementation":  kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.joetroller_newpendingimplementation"), conf.GetString("kafka.env")),
		"masterchefjoev2_add":                  kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.masterchefjoev2_add"), conf.GetString("kafka.env")),
		"masterchefjoev2_deposit":              kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.masterchefjoev2_deposit"), conf.GetString("kafka.env")),
		"masterchefjoev2_emergencywithdraw":    kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.masterchefjoev2_emergencywithdraw"), conf.GetString("kafka.env")),
		"masterchefjoev2_harvest":              kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.masterchefjoev2_harvest"), conf.GetString("kafka.env")),
		"masterchefjoev2_ownershiptransferred": kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.masterchefjoev2_ownershiptransferred"), conf.GetString("kafka.env")),
		"masterchefjoev2_set":                  kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.masterchefjoev2_set"), conf.GetString("kafka.env")),
		"masterchefjoev2_setdevaddress":        kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.masterchefjoev2_setdevaddress"), conf.GetString("kafka.env")),
		"masterchefjoev2_updateemissionrate":   kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.masterchefjoev2_updateemissionrate"), conf.GetString("kafka.env")),
		"masterchefjoev2_updatepool":           kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.masterchefjoev2_updatepool"), conf.GetString("kafka.env")),
		"masterchefjoev2_withdraw":             kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.masterchefjoev2_withdraw"), conf.GetString("kafka.env")),
		"masterchefjoev3_add":                  kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.masterchefjoev3_add"), conf.GetString("kafka.env")),
		"masterchefjoev3_deposit":              kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.masterchefjoev3_deposit"), conf.GetString("kafka.env")),
		"masterchefjoev3_emergencywithdraw":    kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.masterchefjoev3_emergencywithdraw"), conf.GetString("kafka.env")),
		"masterchefjoev3_harvest":              kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.masterchefjoev3_harvest"), conf.GetString("kafka.env")),
		"masterchefjoev3_init":                 kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.masterchefjoev3_init"), conf.GetString("kafka.env")),
		"masterchefjoev3_ownershiptransferred": kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.masterchefjoev3_ownershiptransferred"), conf.GetString("kafka.env")),
		"masterchefjoev3_set":                  kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.masterchefjoev3_set"), conf.GetString("kafka.env")),
		"masterchefjoev3_updatepool":           kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.masterchefjoev3_updatepool"), conf.GetString("kafka.env")),
		"masterchefjoev3_withdraw":             kafkautils.MustParseTopic(conf.GetString("traderjoe.kafka.topic.masterchefjoev3_withdraw"), conf.GetString("kafka.env")),
	}

	tokenCache, err := tokencache.NewTokenCache(traderjoe.TokenNamespace, db, ethClientPool, 18000)
	if err != nil {
		log.Fatal().Err(err).Msg("token cache failed to create")
	}

	connector := traderjoe.NewConnector(
		kp,
		addresses,
		topics,
		avaxClientPool,
		db,
		tokenCache,
	)

	monitor.StartMonitor("traderjoe")
	connector.Start(context.Background(), conf.GetUint64("traderjoe.backfillNumBlocks"))
}
