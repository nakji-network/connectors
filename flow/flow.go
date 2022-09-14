package flow

import (
	"context"

	"github.com/nakji-network/connector"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

type Config struct {
	Host      string
	FromBlock uint64
	NumBlocks uint64
}

type Connector struct {
	*connector.Connector
	*Config
	contracts map[string]ISmartContract
}

func New(c *connector.Connector, config *Config) *Connector {
	return &Connector{
		Connector: c,
		Config:    config,
		contracts: make(map[string]ISmartContract),
	}
}

func (c *Connector) AddContract(sc ISmartContract) {
	c.contracts[sc.Address()] = sc
}

func (c *Connector) GetContract(addr string) ISmartContract {
	return c.contracts[addr]
}

func (c *Connector) Start() {
	var events []string
	for _, contract := range c.contracts {
		events = append(events, contract.Events()...)
	}

	ctx := context.Background()

	sub, err := NewSubscription(ctx, c.Host, events, c.FromBlock, c.NumBlocks)
	if err != nil {
		log.Fatal().Err(err).Str("host", c.Host).Msg("connection error")
	}

	for {
		select {
		case <-sub.Done():
			log.Info().Msg("connector shutdown")
			return

		//	Listen to error channel
		case err := <-sub.Err():
			log.Error().Err(err).Str("host", c.Host).Msg("subscription failed")

		//	Listen to event logs
		case vLog := <-sub.Logs():
			if msg := c.parse(vLog); msg != nil {
				c.EventSink <- msg
			}
		}
	}
}

func (c *Connector) parse(vLog Log) proto.Message {
	contract := c.GetContract(vLog.Type.ContractAddr)
	if contract == nil {
		log.Info().Str("address", vLog.Type.ContractAddr).Msg("event from unsupported address")
		return nil
	}
	return contract.Message(vLog)
}
