package woofi

import (
	"context"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/common"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Config struct {
	ConnectorName string
	NetworkName   string
	FromBlock     uint64
	NumBlocks     uint64
}

type ContractKey struct {
	Network string
	Address string
}

type Connector struct {
	*connector.Connector
	*Config
	sub       *SubscriptionGroup
	contracts map[ContractKey]ISmartContract
}

func New(c *connector.Connector, config *Config) *Connector {
	return &Connector{
		Connector: c,
		Config:    config,
		sub:       NewSubscriptionGroup(),
		contracts: map[ContractKey]ISmartContract{},
	}
}

func (c *Connector) AddContract(sc ISmartContract) {
	c.contracts[ContractKey{sc.Network(), sc.Address()}] = sc
}

func (c *Connector) GetContract(network string, address string) ISmartContract {
	return c.contracts[ContractKey{network, address}]
}

func (c *Connector) Start() {
	networks := make(map[string][]ethcommon.Address)
	for _, contract := range c.contracts {
		address := ethcommon.HexToAddress(contract.Address())
		networks[contract.Network()] = append(networks[contract.Network()], address)
	}

	ctx := context.Background()

	for network, addresses := range networks {
		if err := c.sub.AddSubscription(ctx, c.Connector, network, addresses, c.FromBlock, c.NumBlocks); err != nil {
			log.Fatal().Err(err).Str("network", network).Msg("connection error")
		}
	}

	c.sub.Init()

	for {
		select {
		case <-c.sub.Done():
			log.Info().Msg("connector shutdown")
			return

		//	Listen to error channel
		case err := <-c.sub.Err():
			log.Error().Err(err).Str("network", c.NetworkName).Msg("subscription failed")
			return

		//	Listen to event logs
		case vLog := <-c.sub.Logs():
			if msg := c.parse(vLog); msg != nil {
				c.EventSink <- msg
			}
		}
	}
}

func (c *Connector) parse(vLog Log) protoreflect.ProtoMessage {
	contract := c.GetContract(vLog.Network, vLog.Address.String())
	if contract == nil {
		log.Error().Str("network", vLog.Network).Str("address", vLog.Address.String()).Msg("unknown event")
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	t, err := c.sub.GetBlockTime(ctx, vLog)
	if err != nil {
		log.Error().Err(err).Msg("failed to retrieve timestamp")
	}
	ts := common.UnixToTimestampPb(int64(t * 1000))

	return contract.Message(vLog.Log, ts)
}
