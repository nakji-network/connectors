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
		NetworkName: "avalanche",
		FromBlock:   *fromBlock,
		NumBlocks:   *numBlocks,
	}

	c := woofi.New(conf)
	c.AddContract(WooPPV3.NewContract("0xF8cE0D043891b62c55380fB1EFBfB4F186153D96"))               // WooPPV2
	c.AddContract(WooPPV3.NewContract("0x1df3009c57a8B143c6246149F00B090Bce3b8f88"))               // WooPPV3
	c.AddContract(WooRouterV2.NewContract("0x160020B09DeD3d862f7f851B5c50632BcF2062FF"))           // WooRouterV1
	c.AddContract(WooRouterV2.NewContract("0x3e0Da0A9e4139B32b37710784b8DCa643c152001"))           // WooRouterV2
	c.AddContract(WooRouterV2.NewContract("0x5AA6a4E96A9129562e2fc06660D07FEdDAAf7854"))           // WooRouterV3
	c.AddContract(WooCrossChainRouterV1.NewContract("0xdF37F7A85D4563f39A78494568824b4dF8669B7a")) // WooCrossChainRouterV1
	c.Start()
}
