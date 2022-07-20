package main

import (
	"github.com/nakji-network/connector"
	"github.com/nakji-network/connectors/deribit"
	"github.com/nakji-network/connectors/deribit/market"

	"github.com/rs/zerolog/log"
)

func main() {

	conn, err := connector.NewConnector()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to instantiate connector")
	}

	// Register topic and protobuf type mappings
	conn.RegisterProtos(
		&market.OpenInterest{},
	)

	c := deribit.DeribitConnector{
		Connector: conn,
	}

	c.Start()
}
