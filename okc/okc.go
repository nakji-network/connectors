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

	ctx := context.TODO()

	c.Connector.RegisterProtos(kafkautils.MsgTypeBf, protos...)

	backfillSignal := make(chan struct{})

	if c.FromBlock == 0 && c.NumBlocks == 0 {
		// LIVE DATA

		// Backfill last few blocks at every start
		const defaultBackfill = 100
		go c.backfill(ctx, nil, 0, defaultBackfill)

		// Listen live data
		go c.listenLogs(ctx)

	} else {
		//	HISTORICAL DATA
		go c.backfill(ctx, backfillSignal, c.FromBlock, c.NumBlocks)

	}

	// Retry failed blocks
	c.retry(ctx, backfillSignal)

	log.Info().Msg("shutting down connector...")
	time.Sleep(5 * time.Second)

	c.Sub.Close()
}

// backfill queries for historical data and pushes them to Kafka.
func (c *Connector) backfill(ctx context.Context, sig chan struct{}, fromBlock, numBlocks uint64) {
	if sig != nil {
		defer close(sig)
	}

	var blockNumber uint64

	messages := make([]*kafkautils.Message, 0)

	if logs, err := ethereum.BackfillEventsWithQueryParams(ctx, c.Client, c.addresses, fromBlock, numBlocks); err == nil {
		for bfLog := range logs {
			select {
			case <-c.Sub.Done():
				return
			default:

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

		//	Flush out last messages
		c.Connector.ProduceWithTransaction(messages)
	}

	log.Info().Uint64("from", fromBlock).Uint64("num blocks", numBlocks).Msg("backfill completed")
}

// listenLogs subscribes to live data and pushes incoming logs to Kafka.
func (c *Connector) listenLogs(ctx context.Context) {

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

	//	Flush out last messages
	c.Connector.ProduceWithTransaction(messages)
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

func (c *Connector) retry(ctx context.Context, backfillSignal chan struct{}) {
	const (
		initialBackoff = time.Second
		maxBackoff     = 24 * time.Hour
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

						if msg := c.parse(kafkautils.MsgTypeBf, ethereum.Log{Log: bfLog}); msg != nil {
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
