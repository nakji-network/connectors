package main

import (
	"github.com/nakji-network/woofi-connectors/woofi"
	"github.com/nakji-network/woofi-connectors/woofi/WooCrossChainRouterV1"
	"github.com/nakji-network/woofi-connectors/woofi/WooPPV3"
	"github.com/nakji-network/woofi-connectors/woofi/WooPPV4"
	"github.com/nakji-network/woofi-connectors/woofi/WooRouterV2"
	"github.com/nakji-network/woofi-connectors/woofi/WooRouterV3"

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
	c.AddContract(WooPPV3.NewContract("0x9503E7517D3C5bc4f9E4A1c6AE4f8B33AC2546f2"))               // WooPPV1_1
	c.AddContract(WooPPV4.NewContract("0x286ab107c5E9083dBed35A2B5fb0242538F4f9bf"))               // WooPPV2_1
	c.AddContract(WooRouterV2.NewContract("0x37B5a5A730dAD670874f26Cc5507bb1b9705e447"))           // WooRouterV1_1
	c.AddContract(WooRouterV3.NewContract("0x382A9b0bC5D29e96c3a0b81cE9c64d6C8F150Efb"))           // WooRouterV2_1
	c.AddContract(WooCrossChainRouterV1.NewContract("0xcF6Ce5Fd6bf28bB1AeAc88A55251f6c840059De5")) // WooCrossChainRouterV1_1
	c.AddContract(WooCrossChainRouterV1.NewContract("0x28D2B949024FE50627f1EbC5f0Ca3Ca721148E40")) // WooCrossChainRouterV1_2
	c.Start()
}
