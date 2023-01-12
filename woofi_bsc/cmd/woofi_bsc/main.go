package main

import (
	"github.com/nakji-network/woofi-connectors/woofi"
	"github.com/nakji-network/woofi-connectors/woofi/WooCrossChainRouterV1"
	"github.com/nakji-network/woofi-connectors/woofi/WooPPV1"
	"github.com/nakji-network/woofi-connectors/woofi/WooPPV2"
	"github.com/nakji-network/woofi-connectors/woofi/WooPPV3"
	"github.com/nakji-network/woofi-connectors/woofi/WooPPV4"
	"github.com/nakji-network/woofi-connectors/woofi/WooRouterV1"
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
		NetworkName: "bsc",
		FromBlock:   *fromBlock,
		NumBlocks:   *numBlocks,
	}

	c := woofi.New(conf)
	c.AddContract(WooPPV1.NewContract("0x10C24658815585851a8888f059Cb4338790023F1"))               // WooPPV1_0
	c.AddContract(WooPPV2.NewContract("0x8489d142Da126F4Ea01750e80ccAa12FD1642988"))               // WooPPV1_1
	c.AddContract(WooPPV3.NewContract("0xbf365Ce9cFcb2d5855521985E351bA3bcf77FD3F"))               // WooPPV1_2
	c.AddContract(WooPPV4.NewContract("0xEc054126922a9a1918435c9072c32f1B60cB2B90"))               // WooPPV2_1
	c.AddContract(WooRouterV1.NewContract("0x114f84658c99aa6EA62e3160a87A16dEaF7EFe83"))           // WooRouterV1_1
	c.AddContract(WooRouterV2.NewContract("0xcEf5BE73ae943B77f9Bc08859367D923C030a269"))           // WooRouterV1_2
	c.AddContract(WooRouterV3.NewContract("0xC90bFE9951a4Efbf20aCa5ECd9966b2bF8A01294"))           // WooRouterV2_1
	c.AddContract(WooCrossChainRouterV1.NewContract("0x53E255e8Bbf4EDF16797f9885291B3Ca0C70B59f")) // WooCrossChainRouterV1_1
	c.AddContract(WooCrossChainRouterV1.NewContract("0xd12D239b781e34E0AAa106159940803A07E31a67")) // WooCrossChainRouterV1_2
	c.Start()
}
