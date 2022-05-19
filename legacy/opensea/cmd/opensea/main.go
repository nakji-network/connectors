// Code generated by connectorgen - Edit as necessary.

package main

import (
	"context"
	"net/http"

	"blep.ai/data/chain/ethereum/ethclient"
	"blep.ai/data/config"
	"blep.ai/data/connectors/source/opensea"
	"blep.ai/data/monitor"
	"github.com/nakji-network/connector/kafkautils"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/heptiolabs/healthcheck"
	"github.com/rs/zerolog/log"
)

var conf = config.InitConfig()

func init() {
	kafkautils.TopicTypeRegistry.Load(opensea.TopicTypes)
	conf.SetDefault("opensea.kafka.topic.wyvernexchangev1_orderapprovedpartone", ".fct.nakji.opensea.0_0_0.wyvernexchangev1_orderapprovedpartone")
	conf.SetDefault("opensea.kafka.topic.wyvernexchangev1_orderapprovedparttwo", ".fct.nakji.opensea.0_0_0.wyvernexchangev1_orderapprovedparttwo")
	conf.SetDefault("opensea.kafka.topic.wyvernexchangev1_ordercancelled", ".fct.nakji.opensea.0_0_0.wyvernexchangev1_ordercancelled")
	conf.SetDefault("opensea.kafka.topic.wyvernexchangev1_ordersmatched", ".fct.nakji.opensea.0_0_0.wyvernexchangev1_ordersmatched")
	conf.SetDefault("opensea.kafka.topic.wyvernexchangev1_ownershiprenounced", ".fct.nakji.opensea.0_0_0.wyvernexchangev1_ownershiprenounced")
	conf.SetDefault("opensea.kafka.topic.wyvernexchangev1_ownershiptransferred", ".fct.nakji.opensea.0_0_0.wyvernexchangev1_ownershiptransferred")
	conf.SetDefault("opensea.kafka.topic.wyvernexchangev2_nonceincremented", ".fct.nakji.opensea.0_0_0.wyvernexchangev2_nonceincremented")
	conf.SetDefault("opensea.kafka.topic.wyvernexchangev2_orderapprovedpartone", ".fct.nakji.opensea.0_0_0.wyvernexchangev2_orderapprovedpartone")
	conf.SetDefault("opensea.kafka.topic.wyvernexchangev2_orderapprovedparttwo", ".fct.nakji.opensea.0_0_0.wyvernexchangev2_orderapprovedparttwo")
	conf.SetDefault("opensea.kafka.topic.wyvernexchangev2_ordercancelled", ".fct.nakji.opensea.0_0_0.wyvernexchangev2_ordercancelled")
	conf.SetDefault("opensea.kafka.topic.wyvernexchangev2_ordersmatched", ".fct.nakji.opensea.0_0_0.wyvernexchangev2_ordersmatched")
	conf.SetDefault("opensea.kafka.topic.wyvernexchangev2_ownershiprenounced", ".fct.nakji.opensea.0_0_0.wyvernexchangev2_ownershiprenounced")
	conf.SetDefault("opensea.kafka.topic.wyvernexchangev2_ownershiptransferred", ".fct.nakji.opensea.0_0_0.wyvernexchangev2_ownershiptransferred")

	conf.SetDefault("opensea.kafka.txID", "opensea")

	conf.SetDefault("opensea.wyvernexchangev1Address", "0x7Be8076f4EA4A4AD08075C2508e481d6C946D12b")
	conf.SetDefault("opensea.wyvernexchangev2Address", "0x7f268357A8c2552623316e2562D90e642bB538E5")
}

func main() {
	kp, err := kafkautils.NewProducer(conf.GetString("kafka.url"), conf.GetString("opensea.kafka.txID"))
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
	log.Info().Str("addr", "0.0.0.0:8080").Msg("healthcheck listening on /live and /ready")
	wyvernexchangev1Address := ethcommon.HexToAddress(conf.GetString("opensea.wyvernexchangev1Address"))
	wyvernexchangev2Address := ethcommon.HexToAddress(conf.GetString("opensea.wyvernexchangev2Address"))

	addresses := map[string][]ethcommon.Address{
		"wyvernexchangev1": {wyvernexchangev1Address},
		"wyvernexchangev2": {wyvernexchangev2Address},
	}

	topics := map[string]kafkautils.Topic{
		"wyvernexchangev1_orderapprovedpartone": kafkautils.MustParseTopic(conf.GetString("opensea.kafka.topic.wyvernexchangev1_orderapprovedpartone"), conf.GetString("kafka.env")),
		"wyvernexchangev1_orderapprovedparttwo": kafkautils.MustParseTopic(conf.GetString("opensea.kafka.topic.wyvernexchangev1_orderapprovedparttwo"), conf.GetString("kafka.env")),
		"wyvernexchangev1_ordercancelled":       kafkautils.MustParseTopic(conf.GetString("opensea.kafka.topic.wyvernexchangev1_ordercancelled"), conf.GetString("kafka.env")),
		"wyvernexchangev1_ordersmatched":        kafkautils.MustParseTopic(conf.GetString("opensea.kafka.topic.wyvernexchangev1_ordersmatched"), conf.GetString("kafka.env")),
		"wyvernexchangev1_ownershiprenounced":   kafkautils.MustParseTopic(conf.GetString("opensea.kafka.topic.wyvernexchangev1_ownershiprenounced"), conf.GetString("kafka.env")),
		"wyvernexchangev1_ownershiptransferred": kafkautils.MustParseTopic(conf.GetString("opensea.kafka.topic.wyvernexchangev1_ownershiptransferred"), conf.GetString("kafka.env")),
		"wyvernexchangev2_nonceincremented":     kafkautils.MustParseTopic(conf.GetString("opensea.kafka.topic.wyvernexchangev2_nonceincremented"), conf.GetString("kafka.env")),
		"wyvernexchangev2_orderapprovedpartone": kafkautils.MustParseTopic(conf.GetString("opensea.kafka.topic.wyvernexchangev2_orderapprovedpartone"), conf.GetString("kafka.env")),
		"wyvernexchangev2_orderapprovedparttwo": kafkautils.MustParseTopic(conf.GetString("opensea.kafka.topic.wyvernexchangev2_orderapprovedparttwo"), conf.GetString("kafka.env")),
		"wyvernexchangev2_ordercancelled":       kafkautils.MustParseTopic(conf.GetString("opensea.kafka.topic.wyvernexchangev2_ordercancelled"), conf.GetString("kafka.env")),
		"wyvernexchangev2_ordersmatched":        kafkautils.MustParseTopic(conf.GetString("opensea.kafka.topic.wyvernexchangev2_ordersmatched"), conf.GetString("kafka.env")),
		"wyvernexchangev2_ownershiprenounced":   kafkautils.MustParseTopic(conf.GetString("opensea.kafka.topic.wyvernexchangev2_ownershiprenounced"), conf.GetString("kafka.env")),
		"wyvernexchangev2_ownershiptransferred": kafkautils.MustParseTopic(conf.GetString("opensea.kafka.topic.wyvernexchangev2_ownershiptransferred"), conf.GetString("kafka.env")),
	}

	connector := opensea.NewConnector(
		kp,
		addresses,
		topics,
		ethClientPool,
	)

	monitor.StartMonitor("opensea")
	connector.Start(context.Background(), conf.GetUint64("opensea.backfillNumBlocks"))
}
