package main

import (
	"fmt"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/config"
	"github.com/nakji-network/connectors/woofi_fantom"
	"github.com/nakji-network/connectors/woofi_fantom/WooCrossChainRouterV1"
	"github.com/nakji-network/connectors/woofi_fantom/WooPPV3"
	"github.com/nakji-network/connectors/woofi_fantom/WooRouterV2"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	_ "go.uber.org/automaxprocs"
)

func main() {
	c, err := connector.NewConnector()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to instantiate connector")
	}

	// TODO: remove backfill from main binary
	if err := processBackfillFlags(c.Config); err != nil {
		log.Fatal().Err(err).Msg("input is not correct")
	}

	conf := &woofi_fantom.Config{
		NetworkName: "fantom",
		FromBlock:   c.Config.GetUint64("from-block"),
		NumBlocks:   c.Config.GetUint64("num-blocks"),
	}

	m := woofi_fantom.New(c, conf)
	m.AddContract(WooPPV3.NewContract("0x9503E7517D3C5bc4f9E4A1c6AE4f8B33AC2546f2"))               // WooPPV3
	m.AddContract(WooRouterV2.NewContract("0x37B5a5A730dAD670874f26Cc5507bb1b9705e447"))           // WooRouterV2
	m.AddContract(WooCrossChainRouterV1.NewContract("0xcF6Ce5Fd6bf28bB1AeAc88A55251f6c840059De5")) // WooCrossChainRouterV1
	m.Start()
}

// Deprecated: backfill should be moved outside the main process
func processBackfillFlags(conf *viper.Viper) error {
	conf.SetDefault("blockTime", 15*time.Second)
	conf.SetDefault("waitBlocks", 4)

	//	Woo initial block is #18675185
	pflag.Int64P("from-block", "f", 0, "block number to start backfill from (optional")
	pflag.Int64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")

	pflag.Parse()
	conf.BindPFlags(pflag.CommandLine)

	return validateFlags(conf)
}

// Deprecated: backfill should be moved outside the main process
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
