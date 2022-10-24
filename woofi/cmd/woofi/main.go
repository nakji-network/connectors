package main

import (
	"fmt"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/config"
	"github.com/nakji-network/connectors/woofi"
	"github.com/nakji-network/connectors/woofi/bscWooPP"
	"github.com/nakji-network/connectors/woofi/polygonWooPP"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	_ "go.uber.org/automaxprocs"
)

func main() {
	c, err := connector.NewConnector()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to instantiate connector")
	}

	c.Config.SetDefault("woofi.author", "nakji")
	c.Config.SetDefault("woofi.version", "0_0_0")
	c.Config.SetDefault("blockTime", 15*time.Second)
	c.Config.SetDefault("waitBlocks", 4)

	//	Woo initial block is #18675185
	pflag.Int64P("from-block", "f", 0, "block number to start backfill from (optional")
	pflag.Int64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")

	pflag.Parse()
	c.Config.BindPFlags(pflag.CommandLine)

	if err := validateFlags(c.Config); err != nil {
		log.Fatal().Err(err).Msg("input is not correct")
	}

	// Register topic and protobuf type mappings
	c.RegisterProtos(woofi.TopicTypes...)

	conf := &woofi.Config{
		ConnectorName: "woofi",
		NetworkName:   "bsc",
		FromBlock:     c.Config.GetUint64("from-block"),
		NumBlocks:     c.Config.GetUint64("num-blocks"),
	}

	bscWooPPContract, err := bscWooPP.NewContract(woofi.BscNetwork, woofi.BscWooPPContractAddr)
	if err != nil {
		log.Err(err).Msg("cannot create bsc WooPP contract")
	}

	polygonWooPPContract, err := polygonWooPP.NewContract(woofi.PolygonNetwork, woofi.PolygonWooPPContractAddr)
	if err != nil {
		log.Err(err).Msg("cannot create polygon WooPP contract")
	}

	m := woofi.New(c, conf)
	m.AddContract(bscWooPPContract)
	m.AddContract(polygonWooPPContract)
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
