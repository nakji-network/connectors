package woofi

import (
	"context"
	"fmt"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/common"
	"github.com/nakji-network/connectors/woofi/WOOPP"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Config struct {
	ConnectorName string
	NetworkName   string
	FromBlock     uint64
	NumBlocks     uint64
}

type Connector struct {
	*connector.Connector
	*Config
	sub       connector.ISubscription
	contracts map[string]*Contract
}

func New(c *connector.Connector, config *Config) *Connector {
	return &Connector{
		Connector: c,
		Config:    config,
	}
}

func (c *Connector) Start() {
	addresses := GetAddresses(ContractAddresses)
	c.contracts = GetContracts(ContractAddresses)

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

func (c *Connector) parse(vLog types.Log) protoreflect.ProtoMessage {
	address := vLog.Address.String()
	if c.contracts[address] == nil {
		log.Info().Str("address", address).Msg("Event from unsupported address")
		return nil
	}
	contractAbi := *c.contracts[address].ABI
	contractName := c.contracts[address].Name
	contractType := c.contracts[address].Type

	abiEvent, err := contractAbi.EventByID(vLog.Topics[0])
	if err != nil {
		log.Warn().Str("contract name", contractName).Err(err).Msg("Failed to get event from ABI")
		return nil
	}

	time, err := c.sub.GetBlockTime(vLog)
	if err != nil {
		log.Error().Str("contract name", contractName).Err(err).Msg("Failed to retrieve timestamp")
	}
	timestamp := common.UnixToTimestampPb(int64(time * 1000))

	if smartContract := getContract(contractType); smartContract != nil {
		return smartContract.Message(abiEvent.Name, &contractAbi, vLog, timestamp)
	}
	return nil
}

func getContract(contractType string) ISmartContract {
	switch contractType {
	case "WOOPP":
		return &WOOPP.SmartContract{}
	}

	return nil
}
