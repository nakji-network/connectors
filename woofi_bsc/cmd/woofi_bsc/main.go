package main

import (
	"fmt"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/config"
	"github.com/nakji-network/connectors/woofi_bsc"
	"github.com/nakji-network/connectors/woofi_bsc/WooCrossChainRouterV1"
	"github.com/nakji-network/connectors/woofi_bsc/WooPPV2"
	"github.com/nakji-network/connectors/woofi_bsc/WooPPV3"
	"github.com/nakji-network/connectors/woofi_bsc/WooRouterV1"
	"github.com/nakji-network/connectors/woofi_bsc/WooRouterV2"

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

	conf := &woofi_bsc.Config{
		NetworkName: "bsc",
		FromBlock:   c.Config.GetUint64("from-block"),
		NumBlocks:   c.Config.GetUint64("num-blocks"),
	}

	m := woofi_bsc.New(c, conf)
	m.AddContract(WooPPV2.NewContract("0x10C24658815585851a8888f059Cb4338790023F1"))               // WooPPV1
	m.AddContract(WooPPV2.NewContract("0x8489d142Da126F4Ea01750e80ccAa12FD1642988"))               // WooPPV2
	m.AddContract(WooPPV3.NewContract("0xbf365Ce9cFcb2d5855521985E351bA3bcf77FD3F"))               // WooPPV3
	m.AddContract(WooRouterV1.NewContract("0x114f84658c99aa6EA62e3160a87A16dEaF7EFe83"))           // WooRouterV1
	m.AddContract(WooRouterV2.NewContract("0xcEf5BE73ae943B77f9Bc08859367D923C030a269"))           // WooRouterV2
	m.AddContract(WooCrossChainRouterV1.NewContract("0x53E255e8Bbf4EDF16797f9885291B3Ca0C70B59f")) // WooCrossChainRouterV1
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
