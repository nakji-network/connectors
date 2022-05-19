package main

import (
	"blep.ai/data/config"
	"blep.ai/data/connectors/source/briq"
	"blep.ai/data/database"
	"blep.ai/data/monitor"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/rs/zerolog/log"
)

var conf = config.GetConfig()

// All events:
//TransferSingle
//Mutate
//ConvertToFT
//ConvertToNFT
//Transfer
//URI

// Smart contracts addresses: https://github.com/briqNFT/briq-builder/blob/main/src/Contracts.ts
// Make sure to remove the third '0' in the addresses

func init() {
	kafkautils.TopicTypeRegistry.Load(briq.TopicTypes)
	// TODO: change it to mainnet
	conf.SetDefault("starknet.network", "alpha-goerli")
	conf.SetDefault("briq.kafka.txID", "briq")
	conf.SetDefault("briq.id", "briq")
	conf.SetDefault("briq.addresses", []string{"0x1317354276941f7f799574c73fd8fe53fa3f251084b4c04d88cf601b6bd915e", "0x266b1276d23ffb53d99da3f01be7e29fa024dd33cd7f7b1eb7a46c67891c9d0", "0x578fd3377d865b7798140731b53258b1270ac19c37a3916645f80e0e4c8ad69"})
	conf.SetDefault("briq.kafka.topic.transfersingle", ".fct.nakji.briq.0_0_0.storage_transfersingle")
	conf.SetDefault("briq.kafka.topic.mutate", ".fct.nakji.briq.0_0_0.storage_mutate")
	conf.SetDefault("briq.kafka.topic.converttoft", ".fct.nakji.briq.0_0_0.storage_converttoft")
	conf.SetDefault("briq.kafka.topic.converttonft", ".fct.nakji.briq.0_0_0.storage_converttonft")
	conf.SetDefault("briq.kafka.topic.transfer", ".fct.nakji.briq.0_0_0.storage_transfer")
	conf.SetDefault("briq.kafka.topic.uri", ".fct.nakji.briq.0_0_0.storage_uri")
}

func main() {
	topics := map[string]kafkautils.Topic{
		"transfersingle": kafkautils.MustParseTopic(conf.GetString("briq.kafka.topic.transfersingle"), conf.GetString("kafka.env")),
		"mutate":         kafkautils.MustParseTopic(conf.GetString("briq.kafka.topic.mutate"), conf.GetString("kafka.env")),
		"converttoft":    kafkautils.MustParseTopic(conf.GetString("briq.kafka.topic.converttoft"), conf.GetString("kafka.env")),
		"converttonft":   kafkautils.MustParseTopic(conf.GetString("briq.kafka.topic.converttonft"), conf.GetString("kafka.env")),
		"transfer":       kafkautils.MustParseTopic(conf.GetString("briq.kafka.topic.transfer"), conf.GetString("kafka.env")),
		"uri":            kafkautils.MustParseTopic(conf.GetString("briq.kafka.topic.uri"), conf.GetString("kafka.env")),
	}

	kp, err := kafkautils.NewProducer(conf.GetString("kafka.url"), conf.GetString("briq.kafka.txID"))
	if err != nil {
		log.Fatal().Err(err).Msg("briq: Failed to create new kafka producer")
	}

	// Connect to Database
	db, err := database.New(conf.GetString("timescaledb.connection"))
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot connect to TimescaleDB")
	}

	network := conf.GetString("starknet.network")

	// Sink that publishes transactional messages to Kafka
	qs := kp.MakeQueueTransactionSink()
	newDb := &database.Db{DbInterface: db}

	connector := briq.NewConnector(topics, kp, qs, newDb, network)

	addresses := conf.GetStringSlice("briq.addresses")
	addrMap := make(map[string]bool)
	for _, addr := range addresses {
		addrMap[addr] = true
	}
	connector.Addresses = addrMap
	abi := briq.ParseABIPayload()
	connector.BuildSelectorName(abi)
	monitor.StartMonitor("briq")
	connector.Start()
}
