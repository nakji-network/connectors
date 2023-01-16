package main

import (
	"github.com/nakji-network/woofi-connectors/woofi"
	"github.com/nakji-network/woofi-connectors/woofi/WooCrossChainRouterV1"
	"github.com/nakji-network/woofi-connectors/woofi/WooPPV4"
	"github.com/nakji-network/woofi-connectors/woofi/WooRouterV3"
	"github.com/nakji-network/woofi-connectors/woofi/WooracleV2"

	"github.com/spf13/pflag"
	_ "go.uber.org/automaxprocs"
)

func main() {
	fromBlock := pflag.Uint64P("from-block", "f", 0, "block number to start backfill from (optional)")
	numBlocks := pflag.Uint64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")

	pflag.Parse()

	conf := &woofi.Config{
		NetworkName: "arbitrum",
		FromBlock:   *fromBlock,
		NumBlocks:   *numBlocks,
	}

	c := woofi.New(conf)
	c.AddContract(WooPPV4.NewContract("0xeFF23B4bE1091b53205E35f3AfCD9C7182bf3062"))               // WooPPV2_2
	c.AddContract(WooracleV2.NewContract("0x37a9dE70b6734dFCA54395D8061d9411D9910739"))            // WooracleV2
	c.AddContract(WooRouterV3.NewContract("0x9aEd3A8896A85FE9a8CAc52C9B402D092B629a30"))           // WooRouterV2_1
	c.AddContract(WooCrossChainRouterV1.NewContract("0x0972A0fa37984E7FF2aEFA53A0Bb10dCE535aa73")) // WooCrossChainRouterV1_1
	c.Start()
}
