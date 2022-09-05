package main

import (
	"fmt"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/config"
	"github.com/nakji-network/connectors/aave"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	_ "go.uber.org/automaxprocs"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func main() {
	c, err := connector.NewConnector()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to instantiate connector")
	}

	c.Config.SetDefault("aave.author", "nakji")
	c.Config.SetDefault("aave.version", "0_0_0")
	c.Config.SetDefault("blockTime", 15*time.Second)
	c.Config.SetDefault("waitBlocks", 4)

	pflag.Int64P("from-block", "f", 0, "block number to start backfill from (optional")
	pflag.Int64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")

	pflag.Parse()
	c.Config.BindPFlags(pflag.CommandLine)

	if err := validateFlags(c.Config); err != nil {
		log.Fatal().Err(err).Msg("input is not correct")
	}

	// Register topic and protobuf type mappings
	protos := make([]protoreflect.ProtoMessage, 0, len(aave.TopicTypes))
	for _, topicProto := range aave.TopicTypes {
		protos = append(protos, topicProto)
	}

	c.RegisterProtos(protos...)

	conf := &aave.Config{
		ConnectorName: "aave",
		NetworkName:   "ethereum",
		FromBlock:     c.Config.GetUint64("from-block"),
		NumBlocks:     c.Config.GetUint64("num-blocks"),
		Namespace:     "aave",
	}

	m := aave.New(c, conf)
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
