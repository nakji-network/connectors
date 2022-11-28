package main

import (
	"github.com/nakji-network/connectors/woofi_polygon"
	"github.com/nakji-network/connectors/woofi_polygon/WooCrossChainRouterV1"
	"github.com/nakji-network/connectors/woofi_polygon/WooPPV3"
	"github.com/nakji-network/connectors/woofi_polygon/WooRouterV2"

	"github.com/spf13/pflag"
	_ "go.uber.org/automaxprocs"
)

func main() {
	fromBlock := pflag.Uint64P("from-block", "f", 0, "block number to start backfill from (optional)")
	numBlocks := pflag.Uint64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")

	pflag.Parse()

	conf := &woofi_polygon.Config{
		NetworkName: "polygon",
		FromBlock:   *fromBlock,
		NumBlocks:   *numBlocks,
	}

	c := woofi_polygon.New(conf)
	c.AddContract(WooPPV3.NewContract("0x7400B665C8f4f3a951a99f1ee9872efb8778723d"))               // WooPPV3
	c.AddContract(WooRouterV2.NewContract("0x9D1A92e601db0901e69bd810029F2C14bCCA3128"))           // WooRouterV2
	c.AddContract(WooCrossChainRouterV1.NewContract("0x376d567C5794cfc64C74852A9DB2105E0b5B482C")) // WooCrossChainRouterV1
	c.Start()
}
