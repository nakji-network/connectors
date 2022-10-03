package woofi

import (
	"context"
	"fmt"
	"sync"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/chain"
	"github.com/nakji-network/connector/common"
	"github.com/nakji-network/connector/kafkautils"
	"github.com/nakji-network/connectors/woofi/WOOPP"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
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
	addresses []ethcommon.Address
	contracts map[string]*Contract
	bfLogs    chan types.Log
}

func New(c *connector.Connector, config *Config) *Connector {
	return &Connector{
		Connector: c,
		Config:    config,
		bfLogs:    make(chan types.Log),
	}
}

func (c *Connector) Start() {
	ctx, cancel := context.WithCancel(context.Background())

	c.setup(ctx)

	go c.backfill(ctx, cancel)

	go func() {
		<-c.sub.Done()
		cancel()
	}()

	if c.FromBlock == 0 && c.NumBlocks == 0 {
		//	Only subscribe to the blockchain events when it is not a backfill job
		c.sub.Subscribe(ctx)
		go c.listen(ctx)
	}

	<-ctx.Done()
	c.sub.Close()
}

func (c *Connector) setup(ctx context.Context) {
	c.addresses = GetAddresses(ContractAddresses)
	c.contracts = GetContracts(ContractAddresses)

	if sub, err := connector.NewSubscription(ctx, c.Connector, c.NetworkName, c.addresses); err == nil {
		c.sub = sub
	} else {
		log.Fatal().Err(err).Msg(fmt.Sprintf("%s connection error", c.NetworkName))
	}
}

func (c *Connector) backfill(ctx context.Context, cancel context.CancelFunc) {
	// TODO - Get block times from DB

	var toBlock uint64
	var lastBlock uint64

	if c.NumBlocks != 0 {
		if c.FromBlock != 0 {
			toBlock = c.FromBlock + c.NumBlocks
		}

		currentBlockNumber, err := c.sub.Client().BlockNumber(ctx)
		if err != nil {
			log.Error().Err(err).Msg("failed to get current block number")
		} else {
			lastBlock = currentBlockNumber - c.NumBlocks
		}
	}

	if c.FromBlock != 0 {
		lastBlock = c.FromBlock
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for bfLog := range c.bfLogs {
			if msg := c.parse(ctx, kafkautils.MsgTypeBf, bfLog); msg != nil {
				c.EventSink <- msg
			}
		}
		wg.Done()
	}()

	chain.Backfill(ctx, c.sub.Client(), c.addresses, c.bfLogs, lastBlock, toBlock)
	wg.Wait()

	if c.FromBlock != 0 || c.NumBlocks != 0 {
		log.Info().Msg("backfill completed. shutting down connector.")
		cancel()
	}
}

func (c *Connector) listen(ctx context.Context) {

	for {
		select {

		//	Listen to error channel
		case err := <-c.sub.Err():
			log.Error().Err(err).Str("network", c.NetworkName).Msg("subscription failed")
			return

		//	Listen to event logs
		case vLog := <-c.sub.Logs():
			if msg := c.parse(ctx, kafkautils.MsgTypeFct, vLog); msg != nil {
				c.EventSink <- msg
			}
		}
	}
}

func (c *Connector) parse(ctx context.Context, msgType kafkautils.MsgType, vLog types.Log) *kafkautils.Message {
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

	if smartContract := getContract(contractType); smartContract != nil {
		return &kafkautils.Message{
			MsgType:  msgType,
			ProtoMsg: smartContract.Message(abiEvent.Name, &contractAbi, vLog, timestamp),
		}
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
