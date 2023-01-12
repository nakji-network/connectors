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
		NetworkName: "avalanche",
		FromBlock:   *fromBlock,
		NumBlocks:   *numBlocks,
	}

	c := woofi.New(conf)
	c.AddContract(WooPPV3.NewContract("0xF8cE0D043891b62c55380fB1EFBfB4F186153D96"))               // WooPPV1_1
	c.AddContract(WooPPV3.NewContract("0x1df3009c57a8B143c6246149F00B090Bce3b8f88"))               // WooPPV1_2
	c.AddContract(WooPPV4.NewContract("0x3b3E4b4741e91aF52d0e9ad8660573E951c88524"))               // WooPPV2_1
	c.AddContract(WooRouterV2.NewContract("0x160020B09DeD3d862f7f851B5c50632BcF2062FF"))           // WooRouterV1_1
	c.AddContract(WooRouterV2.NewContract("0x3e0Da0A9e4139B32b37710784b8DCa643c152001"))           // WooRouterV1_2
	c.AddContract(WooRouterV2.NewContract("0x5AA6a4E96A9129562e2fc06660D07FEdDAAf7854"))           // WooRouterV1_3
	c.AddContract(WooRouterV3.NewContract("0xC22FBb3133dF781E6C25ea6acebe2D2Bb8CeA2f9"))           // WooRouterV2_1
	c.AddContract(WooCrossChainRouterV1.NewContract("0xdF37F7A85D4563f39A78494568824b4dF8669B7a")) // WooCrossChainRouterV1_1
	c.AddContract(WooCrossChainRouterV1.NewContract("0xd12D239b781e34E0AAa106159940803A07E31a67")) // WooCrossChainRouterV1_2
	c.Start()
}
