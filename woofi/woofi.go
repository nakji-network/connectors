package woofi

import (
	"context"
	"fmt"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/common"
	"github.com/nakji-network/connectors/woofi/WooCrossChainRouterV1"
	"github.com/nakji-network/connectors/woofi/WooPP"
	"github.com/nakji-network/connectors/woofi/WooPPV1"
	"github.com/nakji-network/connectors/woofi/WooPPV2"
	"github.com/nakji-network/connectors/woofi/WooRouterV1"
	"github.com/nakji-network/connectors/woofi/WooRouterV2"
	"github.com/nakji-network/connectors/woofi/WooRouterV3"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var contracts = map[string]ISmartContract{
	"WooPP":                 &WooPP.SmartContract{},
	"WooPPV1":               &WooPPV1.SmartContract{},
	"WooPPV2":               &WooPPV2.SmartContract{},
	"WooRouterV1":           &WooRouterV1.SmartContract{},
	"WooRouterV2":           &WooRouterV2.SmartContract{},
	"WooRouterV3":           &WooRouterV3.SmartContract{},
	"WooCrossChainRouterV1": &WooCrossChainRouterV1.SmartContract{},
}

type Config struct {
	NetworkName       string
	ContractAddresses map[string][]string
	FromBlock         uint64
	NumBlocks         uint64
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
	c.RegisterProtos(
		&WooPP.FeeManagerUpdated{},
		&WooPP.OwnershipTransferPrepared{},
		&WooPP.OwnershipTransferred{},
		&WooPP.ParametersUpdated{},
		&WooPP.Paused{},
		&WooPP.RewardManagerUpdated{},
		&WooPP.StrategistUpdated{},
		&WooPP.Unpaused{},
		&WooPP.Withdraw{},
		&WooPP.WooGuardianUpdated{},
		&WooPP.WooracleUpdated{},
		&WooPP.WooSwap{},

		&WooPPV1.OwnershipTransferPrepared{},
		&WooPPV1.OwnershipTransferred{},
		&WooPPV1.ParametersUpdated{},
		&WooPPV1.Paused{},
		&WooPPV1.RewardManagerUpdated{},
		&WooPPV1.StrategistUpdated{},
		&WooPPV1.Unpaused{},
		&WooPPV1.Withdraw{},
		&WooPPV1.WooGuardianUpdated{},
		&WooPPV1.WooracleUpdated{},
		&WooPPV1.WooSwap{},

		&WooPPV2.OwnershipTransferPrepared{},
		&WooPPV2.OwnershipTransferred{},
		&WooPPV2.ParametersUpdated{},
		&WooPPV2.Paused{},
		&WooPPV2.RewardManagerUpdated{},
		&WooPPV2.StrategistUpdated{},
		&WooPPV2.Unpaused{},
		&WooPPV2.Withdraw{},
		&WooPPV2.WooGuardianUpdated{},
		&WooPPV2.WooracleUpdated{},
		&WooPPV2.WooSwap{},

		&WooRouterV1.OwnershipTransferred{},
		&WooRouterV1.WooPoolChanged{},
		&WooRouterV1.WooRouterSwap{},

		&WooRouterV2.OwnershipTransferred{},
		&WooRouterV2.WooPoolChanged{},
		&WooRouterV2.WooRouterSwap{},

		&WooRouterV3.OwnershipTransferred{},
		&WooRouterV3.WooPoolChanged{},
		&WooRouterV3.WooRouterSwap{},

		&WooCrossChainRouterV1.WooCrossSwapOnSrcChain{},
		&WooCrossChainRouterV1.WooCrossSwapOnDstChain{},
		&WooCrossChainRouterV1.OwnershipTransferred{},
	)

	addresses := GetAddresses(c.ContractAddresses)
	c.contracts = GetContracts(c.ContractAddresses)

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

	time, err := c.sub.GetBlockTime(context.Background(), vLog)
	if err != nil {
		log.Error().Str("contract name", contractName).Err(err).Msg("Failed to retrieve timestamp")
	}
	timestamp := common.UnixToTimestampPb(int64(time * 1000))

	if smartContract, ok := contracts[contractType]; ok {
		return smartContract.Message(abiEvent.Name, &contractAbi, vLog, timestamp)
	}
	return nil
}
