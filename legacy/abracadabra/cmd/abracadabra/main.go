package main

import (
	"blep.ai/data/connectors/source/abracadabra"

	"github.com/nakji-network/connector/kafkautils"
)

func init() {
	kafkautils.TopicTypeRegistry.Load(abracadabra.TopicTypes)
}

func main() {
	c := abracadabra.NewConnector()

	c.Config.SetDefault("abracadabra.kafka.txID", "abracadabra")
	c.Config.SetDefault("abracadabra.backfillNumBlocks", 100)
	c.Config.SetDefault("abracadabra.kafka.topic.magiccrvtransfer", ".fct.nakji.abracadabra.0_0_0.magiccrv_transfer")
	c.Config.SetDefault("abracadabra.kafka.topic.magiccrvapproval", ".fct.nakji.abracadabra.0_0_0.magiccrv_approval")

	env := c.Config.GetString("kafka.env")
	c.Topics = map[string]kafkautils.Topic{
		"magiccrvapproval": kafkautils.MustParseTopic(c.Config.GetString("abracadabra.kafka.topic.magiccrvapproval"), env),
		"magiccrvtransfer": kafkautils.MustParseTopic(c.Config.GetString("abracadabra.kafka.topic.magiccrvtransfer"), env),
	}

	c.Start()
}
