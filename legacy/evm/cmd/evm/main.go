// This connector ingests real time data from any EVM compatible chain. After
// compiling this program, you can specify which chain to connect to as follows:
//
//     bin/ethereum -n avalanche
package main

import (
	"fmt"

	"blep.ai/data/config"
	"blep.ai/data/connectors/source/evm"
	"blep.ai/data/monitor"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
)

var conf = config.GetConfig()

func init() {
	conf.SetDefault("evm.kafka.txSuffix", "evm-source")
	conf.SetDefault("evm.kafka.topics.block", ".fct.nakji.evm.0_0_0.chain_block")
	conf.SetDefault("evm.kafka.topics.tx", ".fct.nakji.evm.0_0_0.chain_tx")

	pflag.StringP("network", "n", "ethereum", "network to connect to e.g. ethereum")

	pflag.Parse()
	conf.BindPFlags(pflag.CommandLine)
}

func main() {
	network := conf.GetString("network")
	kafkaTransactionID := fmt.Sprintf("%s-%s", network, conf.GetString("evm.kafka.txsuffix"))
	RPCURLs := conf.GetStringSlice(fmt.Sprintf("rpcs.%s.full", network))
	getBlockByNumber := conf.GetBool(fmt.Sprintf("rpcs.%s.getBlockByNumber", network))

	log.Info().
		Str("network", network).
		Strs("RPCs", RPCURLs).
		Str("kafkaTxId", kafkaTransactionID).
		Bool("getBlockByNumber", getBlockByNumber).
		Msg("Starting EVM connector")

	// Load Topic registry
	kafkautils.TopicTypeRegistry.Load(evm.TopicTypes)

	kp, err := kafkautils.NewProducer(conf.GetString("kafka.url"), kafkaTransactionID)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create new kafka producer")
	}

	connector := evm.Connector{
		KP:               kp,
		RPCURLs:          RPCURLs,
		BlockTopic:       kafkautils.MustParseTopic(conf.GetString("evm.kafka.topics.block"), conf.GetString("kafka.env")),
		TXTopic:          kafkautils.MustParseTopic(conf.GetString("evm.kafka.topics.tx"), conf.GetString("kafka.env")),
		GetBlockByNumber: getBlockByNumber,
		Network:          network,
	}

	monitor.StartMonitor(fmt.Sprintf("%s-%s", network, "evm-source"))

	connector.Start()
}
