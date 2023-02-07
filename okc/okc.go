package okc

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/chain/ethereum"
	nakjicommon "github.com/nakji-network/connector/common"
	"github.com/nakji-network/connector/kafkautils"

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
	addresses  []common.Address
	blockRetry blocksToRetry
	contracts  map[string]*Contract
}

type blocksToRetry struct {
	blocks []uint64
	mu     sync.Mutex
}

func New(conf *Config) *Connector {
	addresses := GetAddresses(ContractAddresses)
	contracts := GetContracts(ContractAddresses)

	opts := make([]connector.Option, 0)
	if conf.FromBlock > 0 || conf.NumBlocks > 0 {
		opts = append(opts, connector.BackfillOption())
	}

	ec := ethereum.NewConnector(context.Background(), addresses, conf.BlockchainName, opts...)

	return &Connector{
		Config:     conf,
		Connector:  ec,
		addresses:  addresses,
		blockRetry: blocksToRetry{blocks: make([]uint64, 0)},
		contracts:  contracts,
	}
}

func (c *Connector) Start() {
	ctx, cancel := context.WithCancel(context.Background())

	go c.listenCloseSignal(cancel)

	c.Connector.RegisterProtos(kafkautils.MsgTypeBf, protos...)

	backfillSignal := make(chan struct{})
	retrySignal := c.retry(ctx, backfillSignal)

	if c.FromBlock == 0 && c.NumBlocks == 0 {
		// LIVE DATA

		// Backfill last 100 blocks at every start
		go c.backfill(ctx, nil, 0, 100, retrySignal, make(chan struct{}))

		// Listen live data
		go c.listenLogs(ctx, cancel)
	} else {
		//	HISTORICAL DATA

		go c.backfill(ctx, cancel, c.FromBlock, c.NumBlocks, retrySignal, backfillSignal)
	}

	<-ctx.Done()
	c.Sub.Close()
}

// backfill queries for historical data and pushes them to Kafka.
func (c *Connector) backfill(ctx context.Context, cancel context.CancelFunc, fromBlock, numBlocks uint64, retrySignal, backfillSignal chan struct{}) {
	var blockNumber uint64

	messages := make([]*kafkautils.Message, 0)

	if logs, err := ethereum.BackfillEventsWithQueryParams(ctx, c.Client, c.addresses, fromBlock, numBlocks); err == nil {
		for bfLog := range logs {
			if msg := c.parse(kafkautils.MsgTypeBf, ethereum.Log{Log: bfLog}); msg != nil {
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

	if backfillSignal != nil {
		close(backfillSignal)
	}

	<-retrySignal
	if cancel != nil {
		log.Info().Msg("backfill completed. shutting down connector.")
		cancel()
	}
}

// listenLogs subscribes to live data and pushes incoming logs to Kafka.
func (c *Connector) listenLogs(ctx context.Context, cancel context.CancelFunc) {

	// Register topic and protobuf type mappings
	c.RegisterProtos(kafkautils.MsgTypeFct, protos...)

	c.Sub.Subscribe(ctx)

	var blockNumber uint64

	messages := make([]*kafkautils.Message, 0)
	for vLog := range c.Sub.Logs() {

		if msg := c.parse(kafkautils.MsgTypeFct, vLog); msg != nil {
			messages = append(messages, msg)

			// Commit messages at every new block
			if blockNumber != vLog.BlockNumber {
				c.Connector.ProduceWithTransaction(messages)
				blockNumber = vLog.BlockNumber
				messages = make([]*kafkautils.Message, 0)
			}
		}
	}
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
func (c *Connector) parse(msgType kafkautils.MsgType, vLog ethereum.Log) *kafkautils.Message {
	address := vLog.Address.String()
	if c.contracts[address] == nil {
		log.Info().Str("address", address).Msg("Event from unsupported address")
		return nil
	}

	contract := c.contracts[address]

	contractAbi := contract.ABI
	contractName := contract.Name
	eventParser := contract.MessageParser

	abiEvent, err := contractAbi.EventByID(vLog.Topics[0])
	if err != nil {
		log.Warn().Str("contract name", contractName).Err(err).Msg("Failed to get event from ABI")
		return nil
	}

	bt, err := c.Sub.GetBlockTime(context.Background(), vLog.Log)
	if err != nil {
		log.Error().Str("contract name", contractName).Err(err).Msg("Failed to retrieve timestamp")
	}
	timestamp := nakjicommon.UnixToTimestampPb(int64(bt * 1000))

	msg := eventParser.Message(abiEvent.Name, contractAbi, vLog.Log, timestamp)
	if msg == nil {
		log.Warn().Str("event", abiEvent.Name).Msg("event is not defined")
		return nil
	}

	return &kafkautils.Message{
		MsgType:  msgType,
		ProtoMsg: msg,
	}
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
					logs, err := ethereum.HistoricalEventsWithQueryParams(ctx, c.Client, c.addresses, blockNo, 1)
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

						messages := make([]*kafkautils.Message, 0)
						for bfLog := range logs {

							if msg := c.parse(kafkautils.MsgTypeBf, ethereum.Log{Log: bfLog}); msg != nil {
								messages = append(messages, msg)
							}
							c.Connector.ProduceWithTransaction(messages)
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
