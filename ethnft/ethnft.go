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
	contracts  []*Contract
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

	ec := ethereum.NewConnector(context.Background(), nil, conf.BlockchainName, opts...)

	return &Connector{
		Config:     conf,
		Connector:  ec,
		blockRetry: blocksToRetry{blocks: make([]uint64, 0)},
		contracts:  GetContracts(),
	}
}

func (c *Connector) Start() {

	ctx, cancel := context.WithCancel(context.Background())

	go c.setup(ctx, cancel)

	<-ctx.Done()

	c.Sub.Close()

	log.Info().Msg("shutting down connector.")
}

func (c *Connector) setup(ctx context.Context, cancel context.CancelFunc) {
	var wg sync.WaitGroup

	backfillSignalERC721 := make(chan struct{})
	backfillSignalERC1155 := make(chan struct{})
	go c.retry(ctx, &wg, backfillSignalERC721, c.contracts[0])
	go c.retry(ctx, &wg, backfillSignalERC1155, c.contracts[1])

	go c.listenCloseSignal(cancel)

	c.Connector.RegisterProtos(kafkautils.MsgTypeBf, protos...)

	if c.FromBlock == 0 && c.NumBlocks == 0 {
		// LIVE DATA

		// Backfill last few blocks at every start
		const defaultBackfill = 10
		go c.backfill(ctx, &wg, backfillSignalERC721, c.contracts[0], 0, defaultBackfill)
		go c.backfill(ctx, &wg, backfillSignalERC1155, c.contracts[1], 0, defaultBackfill)

		// Register topic and protobuf type mappings
		c.RegisterProtos(kafkautils.MsgTypeFct, protos...)

		c.Sub.Subscribe(ctx)

		// Listen live data
		go c.listenLogs(ctx, &wg, c.contracts[0])
		go c.listenLogs(ctx, &wg, c.contracts[1])

	} else {
		//	HISTORICAL DATA

		go c.backfill(ctx, &wg, backfillSignalERC721, c.contracts[0], c.FromBlock, c.NumBlocks)
		go c.backfill(ctx, &wg, backfillSignalERC1155, c.contracts[1], c.FromBlock, c.NumBlocks)
	}

	wg.Wait()
	cancel()
}

// backfill queries for historical data and pushes them to Kafka.
func (c *Connector) backfill(ctx context.Context, wg *sync.WaitGroup, sig chan struct{}, contract *Contract, fromBlock, numBlocks uint64) {
	wg.Add(1)
	defer wg.Done()

	var blockNumber uint64

	messages := make([]*kafkautils.Message, 0)

	if logs, err := ethereum.BackfillEventsWithQueryParams(ctx, c.Client, nil, fromBlock, numBlocks); err == nil {
		for bfLog := range logs {
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

		//	Flush out last messages
		c.Connector.ProduceWithTransaction(messages)
	}

	if sig != nil {
		close(sig)
	}

	log.Info().Uint64("from", fromBlock).Uint64("num blocks", numBlocks).Msg("backfill completed")
}

// listenLogs subscribes to live data and pushes incoming logs to Kafka.
func (c *Connector) listenLogs(ctx context.Context, wg *sync.WaitGroup, contract *Contract) {
	wg.Add(1)
	defer wg.Done()

	abiEvents := make([]common.Hash, len(contract.ABI.Events))
	i := 0
	for _, v := range contract.ABI.Events {
		abiEvents[i] = v.ID
		i++
	}

	q := geth.FilterQuery{
		Topics: [][]common.Hash{abiEvents},
	}

	eventLogs := make(chan types.Log)
	sub, err := c.Client.SubscribeFilterLogs(ctx, q, eventLogs)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to subscribe to event logs")
	}

	defer sub.Unsubscribe()

	var blockNumber uint64

	messages := make([]*kafkautils.Message, 0)
	for vLog := range eventLogs {

		if msg := c.parse(kafkautils.MsgTypeFct, vLog, contract); msg != nil {
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

// listenCloseSignal signals the program to terminate.
func (c *Connector) listenCloseSignal(cancel context.CancelFunc) {
	select {
	//	Listen to error channel
	case err := <-c.Sub.Err():
		log.Error().Err(err).Str("blockchain", c.BlockchainName).Msg("subscription failed")
		cancel()

	case <-c.Sub.Done():
		cancel()
	}
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

func (c *Connector) retry(ctx context.Context, wg *sync.WaitGroup, backfillSignal chan struct{}, contract *Contract) {
	wg.Add(1)
	defer wg.Done()

	const (
		initialBackoff = time.Second
		maxBackoff     = 24 * time.Hour //	1 day in minutes
	)

	backoff := initialBackoff

	for {
		select {
		case <-ctx.Done():
			return
		case <-backfillSignal:
			if len(c.blockRetry.blocks) == 0 {
				return
			}
		default:
			if len(c.blockRetry.blocks) > 0 {
				blockNo := c.blockRetry.pop()
				logs, err := ethereum.HistoricalEventsWithQueryParams(ctx, c.Client, nil, blockNo, 1)
				if err != nil {
					select {
					case <-ctx.Done():
						return
					default:
						time.Sleep(backoff)
						log.Warn().Err(err).Uint64("block", blockNo).Uint("backoff seconds", uint(backoff.Seconds())).Msg("failed to retrieve block, skipping...")
						c.blockRetry.push(blockNo)
						backoff *= 2
					}
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
