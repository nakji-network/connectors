package main

import (
	"fmt"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/config"
	"github.com/nakji-network/connectors/near"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
)

func main() {
	c, err := connector.NewConnector()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to instantiate connector")
	}

	c.Config.SetDefault("near.author", "nakji")
	c.Config.SetDefault("near.version", "0_0_0")
	c.Config.SetDefault("blockTime", 1*time.Second)
	c.Config.SetDefault("waitBlocks", 0)

	// Genesis height is 9820210. Setting from-block to 0 will default to latest block.
	pflag.Int64P("from-block", "f", 0, "block number to start backfill from (optional")
	pflag.Int64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")

	pflag.Parse()
	c.Config.BindPFlags(pflag.CommandLine)

	if err := validateFlags(c.Config); err != nil {
		log.Fatal().Err(err).Msg("input is not correct")
	}

	config := &near.Config{
		ConnectorName: "near",
		NetworkName:   "near",
		FromBlock:     c.Config.GetUint64("from-block"),
		NumBlocks:     c.Config.GetUint64("num-blocks"),
		Namespace:     "near",
		WsPort:        "9944",
	}

	connector := near.NewConnector(c, config)

	// Register topic and protobuf type mapping
	connector.RegisterProtos(
		&near.Block{},
		&near.Transaction{},
	)

	connector.Start()
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
