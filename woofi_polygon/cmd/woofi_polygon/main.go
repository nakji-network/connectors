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
		NetworkName: "polygon",
		FromBlock:   *fromBlock,
		NumBlocks:   *numBlocks,
	}

	c := woofi.New(conf)
	c.AddContract(WooPPV3.NewContract("0x7400B665C8f4f3a951a99f1ee9872efb8778723d"))               // WooPPV1_1
	c.AddContract(WooPPV4.NewContract("0x7081A38158BD050Ae4a86e38E0225Bc281887d7E"))               // WooPPV2_1
	c.AddContract(WooRouterV2.NewContract("0x9D1A92e601db0901e69bd810029F2C14bCCA3128"))           // WooRouterV1_1
	c.AddContract(WooRouterV3.NewContract("0x817Eb46D60762442Da3D931Ff51a30334CA39B74"))           // WooRouterV2_1
	c.AddContract(WooCrossChainRouterV1.NewContract("0x376d567C5794cfc64C74852A9DB2105E0b5B482C")) // WooCrossChainRouterV1_1
	c.AddContract(WooCrossChainRouterV1.NewContract("0x574b9cec19553435B360803D8B4De2a5b2C008Fd")) // WooCrossChainRouterV1_2
	c.Start()
}
