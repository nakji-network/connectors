package main

import (
	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/kafkautils"
	"github.com/nakji-network/connectors/bybit"
	"github.com/nakji-network/connectors/bybit/market"

	"github.com/rs/zerolog/log"
)

func main() {

	conn, err := connector.NewConnector()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to instantiate connector")
	}

	// Register topic and protobuf type mappings
	conn.RegisterProtos(kafkautils.MsgTypeFct, &market.OpenInterest{})

	c := bybit.BybitConnector{
		Connector: conn,
	}

	c.Start()
}
