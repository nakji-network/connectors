package main

import (
	"context"

	"github.com/nakji-network/connector/kafkautils"
	"github.com/nakji-network/connectors/bitcoin"

	_ "go.uber.org/automaxprocs"
)

func main() {
	connector := bitcoin.NewConnector(func() {})

	// Register topic and protobuf type mappings
	connector.RegisterProtos(kafkautils.MsgTypeFct,
		&bitcoin.Block{},
		&bitcoin.Transaction{},
	)

	connector.Start(context.Background())
}
