package near

import (
	"context"

	"github.com/nakji-network/connector"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

type Config struct {
	FromBlock uint64
	NumBlocks uint64
	Host      string
	MsgTypes  []proto.Message
}

type Connector struct {
	*connector.Connector
	*Config
}

func New(c *connector.Connector, config *Config) *Connector {
	return &Connector{
		Connector: c,
		Config:    config,
	}
}

func (c *Connector) Start() {
	var events []string
	ctx := context.Background()

	sub, err := NewSubscription(ctx, c.Host, c.MsgTypes, events, c.FromBlock, c.NumBlocks)
	if err != nil {
		log.Fatal().Err(err).Msg("connection error")
	}

	for {
		select {
		case <-sub.Done():
			log.Info().Msg("connector shutdown")
			return

		// Listen to error channel
		case err := <-sub.Err():
			log.Error().Err(err).Msg("subscription failed")

		// Listen for new blocks
		case block := <-sub.Blocks():
			c.EventSink <- block

		// Listen for new transactions
		case tx := <-sub.Transactions():
			c.EventSink <- tx

		// Listen for new transactions
		case receipt := <-sub.Receipts():
			c.EventSink <- receipt

		// Listen for new transactions
		case outcome := <-sub.ExecutionOutcomes():
			c.EventSink <- outcome
		}
	}
}
