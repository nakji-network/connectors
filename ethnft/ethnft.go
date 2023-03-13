package ethnft

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/chain/ethereum"
	nakjicommon "github.com/nakji-network/connector/common"
	"github.com/nakji-network/connector/kafkautils"

	geth "github.com/ethereum/go-ethereum"
	"github.com/rs/zerolog/log"
)

type Config struct {
	BlockchainName string
	FromBlock      uint64
	NumBlocks      uint64
}

type Connector struct {
	*ethereum.Connector
	*Config
	blockRetry blocksToRetry
	listeners  []*listener
}

type listener struct {
	contract *Contract
	query    geth.FilterQuery
}

type blocksToRetry struct {
	blocks []uint64
	mu     sync.Mutex
}

func New(conf *Config) *Connector {
	opts := make([]connector.Option, 0)
	if conf.FromBlock > 0 || conf.NumBlocks > 0 {
		opts = append(opts, connector.BackfillOption())
	}

	contracts := GetContracts()

	listeners := make([]*listener, len(contracts))

	topics := make([][]common.Hash, len(contracts))
	for i, contract := range contracts {
		abiEvents := make([]common.Hash, len(contract.ABI.Events))
		j := 0
		for _, v := range contract.ABI.Events {
			abiEvents[j] = v.ID
			j++
		}

		topics[i] = abiEvents

		listeners[i] = &listener{
			contract: contract,
			query: geth.FilterQuery{
				Topics: [][]common.Hash{abiEvents},
			},
		}
	}

	q := geth.FilterQuery{
		Topics: topics,
	}

	ec := ethereum.NewConnectorWithQuery(context.Background(), q, conf.BlockchainName, opts...)

	return &Connector{
		Config:     conf,
		Connector:  ec,
		blockRetry: blocksToRetry{blocks: make([]uint64, 0)},
		listeners:  listeners,
	}
}

func (c *Connector) Start() {

	ctx := context.TODO()

	// Register topic and protobuf type mappings
	c.Connector.RegisterProtos(kafkautils.MsgTypeBf, protos...)
	c.Connector.RegisterProtos(kafkautils.MsgTypeFct, protos...)

	var wg sync.WaitGroup
	for _, lst := range c.listeners {
		wg.Add(1)
		go func(listener *listener) {
			c.launch(ctx, listener)
			wg.Done()
		}(lst)
	}

	wg.Wait()

	log.Info().Msg("shutting down connector...")
	time.Sleep(5 * time.Second)

	c.Sub.Close()
}

func (c *Connector) launch(ctx context.Context, listener *listener) {

	backfillSignal := make(chan struct{})

	if c.FromBlock == 0 && c.NumBlocks == 0 {
		// LIVE DATA

		// Backfill last few blocks at every start
		const defaultBackfill = 10
		go c.backfill(ctx, nil, listener.contract, 0, defaultBackfill)

		// Listen live data
		go c.listenLogs(ctx, listener)

	} else {
		//	HISTORICAL DATA
		go c.backfill(ctx, backfillSignal, listener.contract, c.FromBlock, c.NumBlocks)

	}

	// Retry failed blocks
	c.retry(ctx, backfillSignal, listener.contract)
}

// backfill queries for historical data and pushes them to Kafka.
func (c *Connector) backfill(ctx context.Context, sig chan struct{}, contract *Contract, fromBlock, numBlocks uint64) {
	if sig != nil {
		defer close(sig)
	}

	var blockNumber uint64

	messages := make([]*kafkautils.Message, 0)

	if logs, err := ethereum.BackfillEventsWithQueryParams(ctx, c.Client, nil, fromBlock, numBlocks); err == nil {
		for bfLog := range logs {
			select {
			case <-c.Sub.Done():
				return
			default:

				if msg := c.parse(kafkautils.MsgTypeBf, bfLog, contract); msg != nil {
					messages = append(messages, msg)

					// Commit messages at every new block
					if blockNumber != bfLog.BlockNumber {
						c.Connector.ProduceWithTransaction(messages)
						blockNumber = bfLog.BlockNumber
						messages = make([]*kafkautils.Message, 0)
					}
				}
			}
		}

		//	Flush out last messages
		c.Connector.ProduceWithTransaction(messages)
	}

	log.Info().Uint64("from", fromBlock).Uint64("num blocks", numBlocks).Msg("backfill completed")
}

// listenLogs subscribes to live data and pushes incoming logs to Kafka.
func (c *Connector) listenLogs(ctx context.Context, listener *listener) {
	eventLogs := make(chan types.Log)
	sub, err := c.Client.SubscribeFilterLogs(ctx, listener.query, eventLogs)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to subscribe to event logs")
	}

	defer sub.Unsubscribe()

	var blockNumber uint64

	messages := make([]*kafkautils.Message, 0)
	for vLog := range eventLogs {

		if msg := c.parse(kafkautils.MsgTypeFct, vLog, listener.contract); msg != nil {
			messages = append(messages, msg)

			// Commit messages at every new block
			if blockNumber != vLog.BlockNumber {
				c.Connector.ProduceWithTransaction(messages)
				blockNumber = vLog.BlockNumber
				messages = make([]*kafkautils.Message, 0)
			}
		}
	}

	//	Flush out last messages
	c.Connector.ProduceWithTransaction(messages)
}

// parse extracts data from incoming event log and converts into a Kafka message.
func (c *Connector) parse(msgType kafkautils.MsgType, vLog types.Log, contract *Contract) *kafkautils.Message {
	contractAbi := contract.ABI
	contractName := contract.Name
	eventParser := contract.MessageParser

	abiEvent, err := contractAbi.EventByID(vLog.Topics[0])
	if err != nil {
		// log.Warn().Str("contract name", contractName).Err(err).Msg("failed to get event from ABI")
		return nil
	}

	//	ERC20 event
	if len(vLog.Topics) == 3 && (abiEvent.Name == EventApproval || abiEvent.Name == EventTransfer) {
		return nil
	}

	bt, err := c.Sub.GetBlockTime(context.Background(), vLog)
	if err != nil {
		log.Error().Str("contract name", contractName).Err(err).Msg("failed to retrieve timestamp")
	}
	timestamp := nakjicommon.UnixToTimestampPb(int64(bt * 1000))

	msg := eventParser.Message(abiEvent.Name, contractAbi, vLog, timestamp)
	if msg == nil {
		log.Warn().Str("event", abiEvent.Name).Msg("event is not defined")
		return nil
	}

	return &kafkautils.Message{
		MsgType:  msgType,
		ProtoMsg: msg,
	}
}

func (c *Connector) retry(ctx context.Context, backfillSignal chan struct{}, contract *Contract) {
	const (
		initialBackoff = time.Second
		maxBackoff     = 24 * time.Hour //	1 day in minutes
	)

	backoff := initialBackoff

	for {
		select {
		case <-c.Sub.Done():
			return
		case <-backfillSignal:
			// Wait until all blocks in the queue are retried
			if len(c.blockRetry.blocks) == 0 {
				return
			}
		default:
			if len(c.blockRetry.blocks) > 0 {
				blockNo := c.blockRetry.pop()
				logs, err := ethereum.HistoricalEventsWithQueryParams(ctx, c.Client, nil, blockNo, 1)
				if err != nil {
					time.Sleep(backoff)
					log.Warn().Err(err).Uint64("block", blockNo).Uint("backoff seconds", uint(backoff.Seconds())).Msg("failed to retrieve block, skipping...")
					c.blockRetry.push(blockNo)
					backoff *= 2

				} else {

					messages := make([]*kafkautils.Message, 0)
					for bfLog := range logs {

						if msg := c.parse(kafkautils.MsgTypeBf, bfLog, contract); msg != nil {
							messages = append(messages, msg)
						}
						c.Connector.ProduceWithTransaction(messages)
					}

					backoff = initialBackoff
				}
			} else {
				time.Sleep(backoff)
			}
		}

		if backoff > maxBackoff {
			for i := 0; i < len(c.blockRetry.blocks); i++ {
				blockNo := c.blockRetry.pop()
				log.Error().Uint64("block", blockNo).Msg("failed to retrieve block permanently")
			}
			return
		}
	}
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
