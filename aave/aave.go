package aave

import (
	"context"
	"fmt"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/common"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

type Config struct {
	ConnectorName string
	NetworkName   string
	FromBlock     uint64
	NumBlocks     uint64
	Namespace     string
}

type Connector struct {
	*connector.Connector
	*Config
	sub       connector.ISubscription
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
	addresses := make([]ethcommon.Address, 0, len(c.contracts))
	for _, v := range c.contracts {
		addresses = append(addresses, ethcommon.HexToAddress(v.Address()))
	}

	ctx := context.Background()

	if sub, err := connector.NewSubscription(ctx, c.Connector, c.NetworkName, addresses, c.FromBlock, c.NumBlocks); err == nil {
		c.sub = sub
	} else {
		log.Fatal().Err(err).Msg(fmt.Sprintf("%s connection error", c.NetworkName))
	}

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

func (c *Connector) parse(vLog types.Log) proto.Message {
	contract := c.GetContract(vLog.Address.String())
	if contract == nil {
		log.Info().Str("address", vLog.Address.String()).Msg("Event from unsupported address")
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	t, err := c.sub.GetBlockTime(ctx, vLog)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve timestamp")
	}
	ts := common.UnixToTimestampPb(int64(t * 1000))

	return contract.Message(vLog, ts)
}
