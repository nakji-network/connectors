package main

import (
	"github.com/nakji-network/woofi-connectors/woofi"
	"github.com/nakji-network/woofi-connectors/woofi/WooCrossChainRouterV1"
	"github.com/nakji-network/woofi-connectors/woofi/WooPPV3"
	"github.com/nakji-network/woofi-connectors/woofi/WooRouterV2"

	"github.com/spf13/pflag"
	_ "go.uber.org/automaxprocs"
)

func main() {
	fromBlock := pflag.Uint64P("from-block", "f", 0, "block number to start backfill from (optional)")
	numBlocks := pflag.Uint64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")

	pflag.Parse()

	conf := &woofi.Config{
		NetworkName: "fantom",
		FromBlock:   *fromBlock,
		NumBlocks:   *numBlocks,
	}

	c := woofi.New(conf)
	c.AddContract(WooPPV3.NewContract("0x9503E7517D3C5bc4f9E4A1c6AE4f8B33AC2546f2"))               // WooPPV3
	c.AddContract(WooRouterV2.NewContract("0x37B5a5A730dAD670874f26Cc5507bb1b9705e447"))           // WooRouterV2
	c.AddContract(WooCrossChainRouterV1.NewContract("0xcF6Ce5Fd6bf28bB1AeAc88A55251f6c840059De5")) // WooCrossChainRouterV1
	c.Start()
}
