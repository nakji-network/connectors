package evm

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"blep.ai/data/chain/ethereum/ethclient"
	"github.com/nakji-network/connector/kafkautils"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

type Connector struct {
	KP               *kafkautils.Producer
	RPCURLs          []string
	BlockTopic       kafkautils.Topic
	TXTopic          kafkautils.Topic
	GetBlockByNumber bool // some evm compatible chains use different block hashes than geth's built-in hash
	Network          string
}

func (c Connector) Start() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	unsubscribe := make(chan interface{})
	// connect to websockets
	log.Info().Str("network", c.Network).Strs("url", c.RPCURLs).Msg("connecting to RPC")
	ethClientPool, err := ethclient.DialPoolContext(context.Background(), c.RPCURLs)
	if err != nil {
		log.Fatal().Err(err).Str("network", c.Network).Msg("RPC connection error")
	}

	// kafkaKey examples: eth.eth, arbitrum.arbitrum
	kafkaKey := []byte(fmt.Sprintf("%s.%s", c.Network, c.Network))

	headers := ethClientPool.ConsumeHeaders(unsubscribe)
	var blocks <-chan *ethtypes.Block

	if c.GetBlockByNumber {
		blocks = ethClientPool.ConsumeBlocksByNumber(headers)
	} else {
		blocks = ethClientPool.ConsumeBlocks(headers)
	}

	defer ethClientPool.Close()

	c.KP.EnableTransactions()

	go func() {
		for block := range blocks {

			t0 := time.Now()

			// EthBlock -> Block -> Protobuf
			var blockData Block
			blockData.UnmarshalEthBlock(block)

			err = c.KP.WriteKafkaMessages(c.BlockTopic, kafkaKey, &blockData)
			if err != nil {
				log.Error().Err(err).Str("topic", c.BlockTopic.String()).Str("key", string(kafkaKey)).Msg("Kafka write block proto")
			}

			// EthTransaction -> Transaction -> Protobuf
			var wg sync.WaitGroup
			for _, tx := range block.Transactions() {
				wg.Add(1)
				go func(tx *ethtypes.Transaction, wg *sync.WaitGroup) {
					defer wg.Done()
					txData := Transaction{}
					err := txData.UnmarshalEthTransaction(tx)
					if err != nil {
						// Don't write invalid transactions to kafka
						return
					}

					txData.Ts = blockData.Ts // Timestamp isn't in the raw transaction from geth

					err = c.KP.WriteKafkaMessages(c.TXTopic, kafkaKey, &txData)
					if err != nil {
						log.Error().Err(err).Str("topic", c.TXTopic.String()).Str("key", string(kafkaKey)).Msg("Kafka write transaction proto")
					}
				}(tx, &wg)
			}
			wg.Wait()

			// Commit Kafka Transaction
			t1 := time.Now()
			err = c.KP.CommitTransaction(nil)
			if err != nil {
				log.Error().Err(err).Msg("Processor: Failed to commit transaction")

				err = c.KP.AbortTransaction(nil)
				if err != nil {
					log.Fatal().Err(err).Msg("")
				}
			}
			// Start a new transaction
			err = c.KP.BeginTransaction()
			if err != nil {
				log.Fatal().Err(err).Msg("")
			}

			log.Info().
				TimeDiff("latency", t1, t0).
				Msg("kafka write")
		}
	}()

	for {
		select {
		case <-interrupt:
			log.Debug().Msg("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			ethClientPool.Close()
			c.KP.Close()
			return
		}
	}
}

func PrintBlock(block *ethtypes.Block) {
	fmt.Printf("hash: %s\n", block.Hash().Hex())       // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
	fmt.Printf("num: %v\n", block.Number().Uint64())   // 3477413
	fmt.Printf("time: %v\n", block.Time())             // 1529525947
	fmt.Printf("nonce: %v\n", block.Nonce())           // 130524141876765836
	fmt.Printf("#tx: %v\n", len(block.Transactions())) // 7
	fmt.Printf("gaslim: %v\n", block.GasLimit())       // 1529525947
	fmt.Printf("gasuse: %v\n", block.GasUsed())        // 1529525947
	fmt.Printf("diff: %v\n", block.Difficulty())       // 1529525947
}
