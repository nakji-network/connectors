// This connector ingests real time data from any EVM compatible chain.
package main

import (
	"github.com/nakji-network/connector"
	"github.com/nakji-network/connectors/ethereum"
	"github.com/nakji-network/connectors/ethereum/evm"
)

func main() {
	c := connector.NewConnector()

	// Register topic and protobuf type mappings
	c.RegisterProtos(
		&evm.Block{},
		&evm.Transaction{},
	)

	evmConnector := ethereum.EthereumConnector{
		Connector: c,

		// Any additional command line arguments, such as chain selection override. Set up via https://pkg.go.dev/github.com/spf13/viper#readme-working-with-flags
		Chain: "avalanche",

		GetBlockByNumber: true,
	}

	evmConnector.Start()
}
