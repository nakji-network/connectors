// Code generated by connectorgen - Edit as necessary.

package main

import (
	"context"
	"net/http"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/heptiolabs/healthcheck"
	"github.com/rs/zerolog/log"

	"blep.ai/data/chain/ethereum/ethclient"
	"blep.ai/data/config"
	"blep.ai/data/connectors/source/eth2"
	"blep.ai/data/monitor"
	"github.com/nakji-network/connector/kafkautils"
)

var conf = config.InitConfig()

func init() {
	kafkautils.TopicTypeRegistry.Load(eth2.TopicTypes)
	conf.SetDefault("eth2.kafka.topic.eth2_depositevent", ".fct.nakji.eth2.0_0_0.eth2_depositevent")
	conf.SetDefault("eth2.kafka.topic.eth2_depositcount", ".fct.nakji.eth2.0_0_0.eth2_depositcount")

	conf.SetDefault("eth2.kafka.txID", "eth2")

	conf.SetDefault("eth2.eth2Address", "0x00000000219ab540356cBB839Cbe05303d7705Fa")
}

func main() {
	kp, err := kafkautils.NewProducer(conf.GetString("kafka.url"), conf.GetString("eth2.kafka.txID"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create new kafka producer")
	}

	urls := conf.GetStringSlice("ethereum.rpc")

	log.Info().Strs("url", urls).Msg("connecting to Ethereum RPC")
	ethClientPool, err := ethclient.DialPoolContext(context.Background(), urls)
	if err != nil {
		log.Fatal().Err(err).Msg("Ethereum RPC connection error")
	}

	// For Liveness and Readiness Probe checks
	health := healthcheck.NewHandler()
	go http.ListenAndServe("0.0.0.0:8080", health)
	log.Info().Str("addr", "0.0.0.0:8080").Msg("health check listening on /live and /ready")
	eth2Address := ethcommon.HexToAddress(conf.GetString("eth2.eth2Address"))

	addresses := map[string][]ethcommon.Address{
		"eth2": {eth2Address},
	}

	topics := map[string]kafkautils.Topic{
		"eth2_depositevent": kafkautils.MustParseTopic(conf.GetString("eth2.kafka.topic.eth2_depositevent"), conf.GetString("kafka.env")),
		"eth2_depositcount": kafkautils.MustParseTopic(conf.GetString("eth2.kafka.topic.eth2_depositcount"), conf.GetString("kafka.env")),
	}

	connector := eth2.NewConnector(
		kp,
		addresses,
		topics,
		ethClientPool,
	)

	monitor.StartMonitor("eth2")
	connector.Start(context.Background(), conf.GetUint64("eth2.backfillNumBlocks"))
}
