package woofi_polygon

import (
	"context"
	"fmt"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/common"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Config struct {
	NetworkName string
	FromBlock   uint64
	NumBlocks   uint64
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
	var (
		events    []proto.Message
		addresses []ethcommon.Address
	)

	for _, contract := range c.contracts {
		events = append(events, contract.Events()...)
		addresses = append(addresses, ethcommon.HexToAddress(contract.Address()))
	}

	c.RegisterProtos(events...)

	if sub, err := connector.NewSubscription(context.Background(), c.Connector, c.NetworkName, addresses, c.FromBlock, c.NumBlocks); err == nil {
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

func (c *Connector) parse(vLog types.Log) protoreflect.ProtoMessage {
	contract := c.GetContract(vLog.Address.String())
	if contract == nil {
		log.Info().Str("address", vLog.Address.String()).Msg("event from unsupported address")
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	t, err := c.sub.GetBlockTime(ctx, vLog)
	if err != nil {
		log.Error().Err(err).Msg("failed to retrieve timestamp")
	}
	ts := common.UnixToTimestampPb(int64(t * 1000))

	return contract.Message(vLog, ts)
}
