package evm

import (
	"context"
	"math/big"
	"strings"
	"time"

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

var ProtobufDefinitions = []proto.Message{
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

	go c.listenCloseSignal(cancel)

	c.RegisterProtos(kafkautils.MsgTypeBf, ProtobufDefinitions...)

	go c.backfill(ctx, cancel, c.FromBlock, c.NumBlocks, 1)

	//	Only subscribe to the blockchain events when it is not a backfill job
	if c.FromBlock == 0 && c.NumBlocks == 0 {

		// Backfill last 100 blocks at every start
		go c.backfill(ctx, nil, 0, 100, 1)

		// Listen live data
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
}

// backfill queries for historical data and pushes them to Kafka.
func (c *Connector) backfill(ctx context.Context, cancel context.CancelFunc, fromBlock, numBlocks, backoff uint64) {
	if fromBlock == 0 && numBlocks == 0 {
		return
	}

	if backoff == 0 {
		backoff = 1
	}

	// Calculate block interval for historical data
	startingBlock := fromBlock
	toBlock, err := c.client.BlockNumber(ctx)
	if err != nil {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Duration(backoff) * time.Minute):
			log.Error().Err(err).Uint64("block", toBlock).Uint64("backoff minutes", backoff).Msg("failed to get current block number, retrying...")
			c.backfill(ctx, cancel, c.FromBlock, c.NumBlocks, backoff<<1)
			return
		}
	}

	if fromBlock > 0 && numBlocks > 0 {
		lastBlock := fromBlock + numBlocks

		if lastBlock < toBlock {
			toBlock = lastBlock
		}

	} else if numBlocks > 0 && numBlocks < toBlock {
		startingBlock = toBlock - numBlocks
	}

	for startingBlock < toBlock {
		block, err := c.client.BlockByNumber(ctx, big.NewInt(int64(toBlock)))
		if err != nil {
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Duration(backoff) * time.Minute):
				log.Error().Err(err).Uint64("block", toBlock).Uint64("backoff minutes", backoff).Msg("failed to retrieve block, retrying...")
				c.backfill(ctx, cancel, startingBlock, toBlock-startingBlock, backoff<<1)
				return
			}
		} else {
			err = c.process(block, kafkautils.MsgTypeBf)
			if err != nil {
				log.Error().Err(err).Uint64("block", block.Number().Uint64()).Msg("failed to process block")
			}
			toBlock--
			if backoff > 1 {
				backoff = 1
			}
		}
	}

	if cancel != nil {
		log.Info().Msg("backfill completed. shutting down connector.")
		cancel()
	}
}

// listenBlocks subscribes to live data and pushes incoming logs to Kafka.
func (c *Connector) listenBlocks(ctx context.Context, cancel context.CancelFunc) {
	// Register topic and protobuf type mappings
	c.RegisterProtos(kafkautils.MsgTypeFct, ProtobufDefinitions...)

	c.sub.Subscribe(ctx)

	for h := range c.sub.Headers() {
		block, err := c.client.BlockByNumber(ctx, h.Number)
		if err != nil {
			block, err = c.client.BlockByHash(ctx, h.Hash())
			if err != nil {
				log.Error().Err(err).Msg("failed to retrieve block")
				continue
			}
		}

		err = c.process(block, kafkautils.MsgTypeFct)
		if err != nil {
			log.Error().Err(err).Uint64("block", block.Number().Uint64()).Msg("failed to process block")
		}
	}
}

// listenCloseSignal signals the program to terminate.
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

func (c *Connector) process(block *types.Block, msgType kafkautils.MsgType) error {

	header := block.Header()
	messages := make([]*kafkautils.Message, len(block.Transactions())+1)
	ts := common.UnixToTimestampPb(int64(header.Time * 1000))

	for i, t := range block.Transactions() {
		messages[i] = &kafkautils.Message{
			MsgType:  msgType,
			ProtoMsg: chain.ParseTransaction(t, ts),
		}
	}

	messages[len(messages)-1] = &kafkautils.Message{
		MsgType:  msgType,
		ProtoMsg: chain.ParseHeader(header),
	}

	return c.ProduceWithTransaction(messages)
}
