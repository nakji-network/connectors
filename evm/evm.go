package evm

import (
	"context"
	"math/big"
	"strings"
	"sync"
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
	client     *ethclient.Client
	sub        ethereum.ISubscription
	blockRetry blocksToRetry
}

type blocksToRetry struct {
	blocks []uint64
	mu     sync.Mutex
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

	backfillSignal := make(chan struct{})
	retrySignal := c.retry(ctx, backfillSignal)

	if c.FromBlock == 0 && c.NumBlocks == 0 {
		// LIVE DATA

		// Backfill last 100 blocks at every start
		go c.backfill(ctx, nil, 0, 100, retrySignal, make(chan struct{}))

		// Listen live data
		go c.listenBlocks(ctx, cancel)
	} else {
		//	HISTORICAL DATA

		go c.backfill(ctx, cancel, c.FromBlock, c.NumBlocks, retrySignal, backfillSignal)
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

	c.blockRetry = blocksToRetry{blocks: make([]uint64, 0)}
}

// backfill queries for historical data and pushes them to Kafka.
func (c *Connector) backfill(ctx context.Context, cancel context.CancelFunc, fromBlock, numBlocks uint64, retrySignal, backfillSignal chan struct{}) {
	if fromBlock == 0 && numBlocks == 0 {
		return
	}

	// Calculate block interval for historical data
	startingBlock := fromBlock
	toBlock, err := c.client.BlockNumber(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get latest block number")
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
			log.Warn().Err(err).Uint64("block", toBlock).Msg("failed to retrieve block, skipping...")
			c.blockRetry.push(toBlock)
		} else {
			err = c.process(block, kafkautils.MsgTypeBf)
			if err != nil {
				log.Warn().Err(err).Uint64("block", block.Number().Uint64()).Msg("failed to process block")
				c.blockRetry.push(toBlock)
			}
		}
		toBlock--
	}

	if backfillSignal != nil {
		close(backfillSignal)
	}

	<-retrySignal
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
				log.Warn().Err(err).Msg("failed to retrieve block")
				c.blockRetry.push(h.Number.Uint64())
				continue
			}
		}

		err = c.process(block, kafkautils.MsgTypeFct)
		if err != nil {
			log.Warn().Err(err).Uint64("block", block.Number().Uint64()).Msg("failed to process block")
			c.blockRetry.push(h.Number.Uint64())
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

func (c *Connector) retry(ctx context.Context, backfillSignal chan struct{}) chan struct{} {
	var backoff uint = 1
	const backoffLimit = 1440 //	1 day in minutes
	ch := make(chan struct{})

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			case <-backfillSignal:
				if len(c.blockRetry.blocks) == 0 {
					close(ch)
					return
				}
			default:
				if len(c.blockRetry.blocks) > 0 {
					blockNo := c.blockRetry.pop()
					block, err := c.client.BlockByNumber(ctx, big.NewInt(int64(blockNo)))
					if err != nil {
						select {
						case <-ctx.Done():
							close(ch)
							return
						case <-time.After(time.Duration(backoff) * time.Minute):
							log.Warn().Err(err).Uint64("block", blockNo).Uint("backoff minutes", backoff).Msg("failed to retrieve block, skipping...")
							c.blockRetry.push(blockNo)
							backoff *= 2
						}
					} else {
						err = c.process(block, kafkautils.MsgTypeBf)
						if err != nil {
							log.Warn().Err(err).Uint64("block", block.Number().Uint64()).Msg("failed to process block")
							c.blockRetry.push(block.NumberU64())
						}

						backoff = 1
					}
				} else {
					time.Sleep(time.Duration(backoff) * time.Minute)
				}
			}

			if backoff > backoffLimit {
				for i := 0; i < len(c.blockRetry.blocks); i++ {
					blockNo := c.blockRetry.pop()
					log.Error().Uint64("block", blockNo).Msg("failed to retrieve block permanently")
				}
				close(ch)
				return
			}
		}
	}()
	return ch
}

func (b *blocksToRetry) push(blockNo uint64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.blocks = append(b.blocks, blockNo)
}

func (b *blocksToRetry) pop() uint64 {
	b.mu.Lock()
	defer b.mu.Unlock()

	blockNo := b.blocks[0]
	b.blocks = b.blocks[1:]
	return blockNo
}
