package aave

import (
	"context"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/chain"
	"github.com/nakji-network/connector/common"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
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
	addresses []ethcommon.Address
	contracts map[string]ISmartContract
	bfLogs    chan types.Log
}

func New(c *connector.Connector, config *Config) *Connector {
	return &Connector{
		Connector: c,
		Config:    config,
		contracts: make(map[string]ISmartContract),
		bfLogs:    make(chan types.Log),
	}
}

func (c *Connector) AddContract(sc ISmartContract) {
	c.contracts[sc.Address()] = sc
}

func (c *Connector) GetContract(addr string) ISmartContract {
	return c.contracts[addr]
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
	addresses := make([]ethcommon.Address, 0, len(c.contracts))
	for _, v := range c.contracts {
		addresses = append(addresses, ethcommon.HexToAddress(v.Address()))
	}
	c.addresses = addresses

	if sub, err := connector.NewSubscription(ctx, c.Connector, c.NetworkName, addresses); err == nil {
		c.sub = sub
	} else {
		log.Fatal().Err(err).Str("network", c.NetworkName).Msg("connection error")
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

	go func() {
		for bfLog := range c.bfLogs {
			if msg := c.parse(ctx, kafkautils.MsgTypeBf, bfLog); msg != nil {
				c.EventSink <- msg
			}
		}
	}()

	chain.Backfill(ctx, c.sub.Client(), c.addresses, c.bfLogs, lastBlock, toBlock)

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
	contract := c.GetContract(vLog.Address.String())
	if contract == nil {
		log.Info().Str("address", vLog.Address.String()).Msg("Event from unsupported address")
		return nil
	}

	t, err := c.sub.GetBlockTime(ctx, vLog)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve timestamp")
	}
	ts := common.UnixToTimestampPb(int64(t * 1000))

	return &kafkautils.Message{
		MsgType:  msgType,
		ProtoMsg: contract.Message(vLog, ts),
	}
}
