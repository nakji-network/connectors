package main

import (
	"context"

	"blep.ai/data/chain/ethereum/ethclient"
	"blep.ai/data/config"
	"blep.ai/data/connectors/source/compound"
	"blep.ai/data/monitor"
	"github.com/nakji-network/connector/kafkautils"

	"cloud.google.com/go/profiler"
	"github.com/rs/zerolog/log"
	_ "go.uber.org/automaxprocs"
)

func main() {
	cfg := profiler.Config{
		Service:        "compound",
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

	conf := config.GetConfig()
	conf.SetDefault("compound.kafka.topic.mint", ".fct.nakji.compound.0_0_0.compound_mint")
	conf.SetDefault("compound.kafka.topic.redeem", ".fct.nakji.compound.0_0_0.compound_redeem")
	conf.SetDefault("compound.kafka.topic.borrow", ".fct.nakji.compound.0_0_0.compound_borrow")
	conf.SetDefault("compound.kafka.topic.repayborrow", ".fct.nakji.compound.0_0_0.compound_repayborrow")
	conf.SetDefault("compound.kafka.topic.liquidateborrow", ".fct.nakji.compound.0_0_0.compound_liquidateborrow")
	conf.SetDefault("compound.kafka.txID", "compound-source")
	conf.SetDefault("compound.contracts.address", []string{
		"0x4Ddc2D193948926D02f9B1fE9e1daa0718270ED5",
		"0x6C8c6b02E7b2BE14d4fA6022Dfd6d75921D90E4E",
		"0x70e36f6BF80a52b3B46b3aF8e106CC0ed743E8e4",
		"0x5d3a536E4D6DbD6114cc1Ead35777bAB948E3643",
		"0x4Ddc2D193948926D02f9B1fE9e1daa0718270ED5",
		"0xFAce851a4921ce59e912d19329929CE6da6EB0c7",
		"0x158079Ee67Fce2f58472A96584A73C7Ab9AC95c1",
		"0xF5DCe57282A584D2746FaF1593d3121Fcac444dC",
		"0x12392F67bdf24faE0AF363c24aC620a2f67DAd86",
		"0x35A18000230DA775CAc24873d00Ff85BccdeD550",
		"0x39AA39c021dfbaE8faC545936693aC917d5E7563",
		"0xf650C3d88D12dB855b8bf7D11Be6C55A4e07dCC9",
		"0xC11b1268C1A384e55C48c2391d8d480264A3A7F4",
		"0xccF4429DB6322D5C611ee964527D42E5d685DD6a",
		"0xB3319f5D18Bc0D84dD1b4825Dcde5d5f7266d407",
	})

	kafkautils.TopicTypeRegistry.Load(compound.TopicTypes)

	kp, err := kafkautils.NewProducer(conf.GetString("kafka.url"), conf.GetString("compound.kafka.txID"))
	if err != nil {
		log.Fatal().Err(err).Msg("UniswapV3: Failed to create new kafka producer")
	}

	urls := conf.GetStringSlice("ethereum.rpc")
	log.Info().Strs("url", urls).Msg("connecting to Ethereum RPC")
	ethClientPool, err := ethclient.DialPoolContext(context.Background(), urls)
	if err != nil {
		log.Fatal().Err(err).Msg("Ethereum RPC connection error")
	}

	topics := map[string]kafkautils.Topic{
		"mint":            kafkautils.MustParseTopic(conf.GetString("compound.kafka.topic.mint"), conf.GetString("kafka.env")),
		"redeem":          kafkautils.MustParseTopic(conf.GetString("compound.kafka.topic.redeem"), conf.GetString("kafka.env")),
		"borrow":          kafkautils.MustParseTopic(conf.GetString("compound.kafka.topic.borrow"), conf.GetString("kafka.env")),
		"repayborrow":     kafkautils.MustParseTopic(conf.GetString("compound.kafka.topic.repayborrow"), conf.GetString("kafka.env")),
		"liquidateborrow": kafkautils.MustParseTopic(conf.GetString("compound.kafka.topic.liquidateborrow"), conf.GetString("kafka.env")),
	}

	rawAddrs := conf.GetStringSlice("compound.contracts.address")
	addresses := compound.ConvertRawAddress(rawAddrs...)

	connector := compound.NewConnector(
		kp,
		topics,
		ethClientPool,
		addresses,
	)

	monitor.StartMonitor("compound")
	connector.Start()
}
