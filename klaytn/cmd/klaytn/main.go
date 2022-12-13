// This connector ingests real time data from any EVM compatible chain.
package main

import "github.com/nakji-network/connectors/klaytn"

func main() {
	ec := klaytn.NewConnector("klaytn")
	ec.Start()
}
