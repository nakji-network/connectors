package evm

import (
	"context"
	"math/big"
	"strings"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/chain/ethereum"
	"github.com/nakji-network/connector/common"
	"github.com/nakji-network/connector/kafkautils"
	"github.com/nakji-network/connectors/evm/chain"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

type Config struct {
	ConnectorName string
	KafkaURL      string
	NetworkName   string
	FromBlock     uint64
	NumBlocks     uint64
}

type Connector struct {
	*Config
	*connector.Connector
	client *ethclient.Client
	sub    ethereum.ISubscription
}

var TopicTypes = []proto.Message{
	&chain.Block{},
	&chain.Transaction{},
}

func New(c *connector.Connector, config *Config) *Connector {
	return &Connector{
		Connector: c,
		Config:    config,
	}
}

func (c *Connector) Start() {
	ctx, cancel := context.WithCancel(context.Background())

	c.setup(ctx)
	go c.backfill(ctx, cancel)

	//	Only subscribe to the blockchain events when it is not a backfill job
	if c.FromBlock == 0 && c.NumBlocks == 0 {
		go c.listenBlocks(ctx, cancel)
	}

	<-ctx.Done()
	c.sub.Close()
}

func (c *Connector) setup(ctx context.Context) {
	// Read config from config yaml under `rpcs.[chain].full`
	rpcs := c.RPCMap[c.NetworkName].Full

	// go-ethereum client only supports 1 rpc connection currently, so we do this hack
	var RPCURL string
	for _, u := range rpcs {
		if strings.HasPrefix(u, "ws") {
			RPCURL = u
			break
		}
	}
	log.Info().Str("chain", c.NetworkName).Str("url", RPCURL).Msg("connecting to RPC")

	client, err := ethclient.DialContext(ctx, RPCURL)
	if err != nil {
		log.Fatal().Err(err).Msg("RPC connection error")
	}
	c.client = client

	sub, err := ethereum.NewSubscription(client, c.NetworkName, nil)
	if err != nil {
		log.Fatal().Err(err).Str("chain", c.NetworkName).Msg("subscription failed")
	}

	c.sub = sub

	// Register topic and protobuf type mappings
	c.RegisterProtos(kafkautils.MsgTypeFct, TopicTypes...)
}

func (c *Connector) backfill(ctx context.Context, cancel context.CancelFunc) {

	if c.FromBlock == 0 && c.NumBlocks == 0 {
		return
	}

	c.RegisterProtos(kafkautils.MsgTypeBf, TopicTypes...)

	fromBlock := c.FromBlock
	toBlock, err := c.client.BlockNumber(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get current block number")
	}

	if c.FromBlock > 0 && c.NumBlocks > 0 {
		lastBlock := c.FromBlock + c.NumBlocks

		if lastBlock < toBlock {
			toBlock = lastBlock
		}

	} else if c.NumBlocks > 0 {
		fromBlock = toBlock - c.NumBlocks
	}

	for fromBlock < toBlock {
		block, err := c.client.BlockByNumber(ctx, big.NewInt(int64(toBlock)))
		if err != nil {
			log.Error().Err(err).Msg("failed to retrieve block")
		}

		err = c.process(block)
		if err != nil {
			log.Error().Err(err).Uint64("block", block.Number().Uint64()).Msg("failed to process block")
		}
		toBlock--
	}

	if c.FromBlock != 0 || c.NumBlocks != 0 {
		log.Info().Msg("backfill completed. shutting down connector.")
		cancel()
	}
}

func (c *Connector) listenBlocks(ctx context.Context, cancel context.CancelFunc) {
	c.sub.Subscribe(ctx)

	go c.listenCloseSignal(cancel)

	for h := range c.sub.Headers() {
		block, err := c.client.BlockByNumber(ctx, h.Number)
		if err != nil {
			log.Error().Err(err).Msg("failed to retrieve block")
		}

		err = c.process(block)
		if err != nil {
			log.Error().Err(err).Uint64("block", block.Number().Uint64()).Msg("failed to process block")
		}
	}
}

func (c *Connector) listenCloseSignal(cancel context.CancelFunc) {
	select {
	//	Listen to error channel
	case err := <-c.sub.Err():
		log.Error().Err(err).Str("network", c.NetworkName).Msg("subscription failed")
		cancel()

	case <-c.sub.Done():
		cancel()
	}
}

func (c *Connector) process(block *types.Block) error {

	header := block.Header()
	messages := make([]*kafkautils.Message, len(block.Transactions())+1)

	for i, t := range block.Transactions() {
		ts := common.UnixToTimestampPb(int64(header.Time))
		messages[i] = &kafkautils.Message{
			MsgType:  kafkautils.MsgTypeFct,
			ProtoMsg: chain.ParseTransaction(t, ts),
		}
	}

	messages[len(messages)-1] = &kafkautils.Message{
		MsgType:  kafkautils.MsgTypeFct,
		ProtoMsg: chain.ParseHeader(header),
	}

	return c.ProduceWithTransaction(messages)
}
