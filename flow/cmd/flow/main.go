package main

import (
	"fmt"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/config"
	"github.com/nakji-network/connectors/flow"
	"github.com/nakji-network/connectors/flow/flowtoken"
	"google.golang.org/protobuf/proto"

	"github.com/onflow/flow-go-sdk/access/grpc"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	_ "go.uber.org/automaxprocs"
)

func main() {
	c, err := connector.NewConnector()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to instantiate connector")
	}

	c.Config.SetDefault("flow.author", "nakji")
	c.Config.SetDefault("flow.version", "0_0_0")
	c.Config.SetDefault("blockTime", 15*time.Second)
	c.Config.SetDefault("waitBlocks", 4)

	pflag.Int64P("from-block", "f", 0, "block number to start backfill from (optional")
	pflag.Int64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")

	pflag.Parse()
	c.Config.BindPFlags(pflag.CommandLine)

	if err := validateFlags(c.Config); err != nil {
		log.Fatal().Err(err).Msg("input is not correct")
	}

	var topicTypes []proto.Message
	// topicTypes = append(topicTypes, flowcontractaudits.Types...)
	// topicTypes = append(topicTypes, flowfees.Types...)
	// topicTypes = append(topicTypes, flowidtablestaking.Types...)
	// topicTypes = append(topicTypes, flowserviceaccount.Types...)
	// topicTypes = append(topicTypes, flowstakingcollection.Types...)
	// topicTypes = append(topicTypes, flowstoragefees.Types...)
	topicTypes = append(topicTypes, flowtoken.Types...)
	// topicTypes = append(topicTypes, lockedtokens.Types...)

	// Register topic and protobuf type mappings
	c.RegisterProtos(topicTypes...)

	conf := &flow.Config{
		Host:      grpc.MainnetHost,
		FromBlock: c.Config.GetUint64("from-block"),
		NumBlocks: c.Config.GetUint64("num-blocks"),
	}

	m := flow.New(c, conf)
	// m.AddContract(flowcontractaudits.SmartContract{})
	// m.AddContract(flowidtablestaking.SmartContract{})
	// m.AddContract(flowidtablestaking.SmartContract{})
	// m.AddContract(flowserviceaccount.SmartContract{})
	// m.AddContract(flowstakingcollection.SmartContract{})
	// m.AddContract(flowstoragefees.SmartContract{})
	m.AddContract(flowtoken.SmartContract{})
	// m.AddContract(lockedtokens.SmartContract{})
	m.Start()
}

func validateFlags(conf config.IConfig) error {
	fromBlock := conf.GetInt64("from-block")
	numBlocks := conf.GetInt64("num-blocks")

	if fromBlock < 0 {
		return fmt.Errorf("backfill input value cannot be negative. from-block: %d", fromBlock)
	}

	if numBlocks < 0 {
		return fmt.Errorf("backfill input value cannot be negative. num-blocks: %d", numBlocks)
	}
	return nil
}
