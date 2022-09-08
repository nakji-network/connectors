package bitcoin

import (
	"context"
	"fmt"

	"github.com/nakji-network/connector"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/rs/zerolog/log"
)

type BitcoinConnector struct {
	*connector.Connector // embed Nakji connector.Connector into your custom connector to get access to all its methods
	client               BTCClient
	hashChan             chan chainhash.Hash
	callback             func()
}

//go:generate mockgen -destination=mocks/mock_client.go -package=mocks . BTCClient
type BTCClient interface {
	GetBlock(*chainhash.Hash) (*wire.MsgBlock, error)
	GetBlockVerbose(*chainhash.Hash) (*btcjson.GetBlockVerboseResult, error)
	GetRawTransactionVerbose(*chainhash.Hash) (*btcjson.TxRawResult, error)
	Shutdown()
	WaitForShutdown()
	NotifyBlocks() error
	GetBlockCount() (int64, error)
}

// NewConnector creates new BitcoinConnector and connects to the bitcoin RPC
func NewConnector(callback func()) *BitcoinConnector {
	c, err := connector.NewConnector()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to instantiate connector")
	}

	hashChan := make(chan chainhash.Hash, 100)

	ntfnHandlers := rpcclient.NotificationHandlers{
		OnFilteredBlockConnected: func(height int32, header *wire.BlockHeader, txns []*btcutil.Tx) {
			hashChan <- header.BlockHash()
		},
		OnFilteredBlockDisconnected: func(height int32, header *wire.BlockHeader) {
			log.Info().Interface("block hash", header.BlockHash()).Int32("height", height).
				Time("timestamp", header.Timestamp).Msg("block disconnected")
		},
	}

	return &BitcoinConnector{
		Connector: c,
		hashChan:  hashChan,
		client:    c.ChainClients.Bitcoin(&ntfnHandlers),
		callback:  callback,
	}
}

func (c *BitcoinConnector) Start(ctx context.Context) {
	// Get the current block count.
	blockCount, err := c.client.GetBlockCount()
	if err != nil {
		log.Fatal().Err(err).Msg("BTC: Get Block Count error")
	}
	log.Info().Int64("count", blockCount).Msg("block count")

	if err := c.client.NotifyBlocks(); err != nil {
		log.Fatal().Err(err).Msg("Block Notifications Register failed")
	}
	log.Info().Msg("BTC: NotifyBlocks: Registration Complete")

	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("BTC: transactions worker cancelled and shutting down")
			c.client.Shutdown()
			return

		case hash := <-c.hashChan:
			log.Info().Str("blockHash", hash.String()).Msg("BTC: writing block to Kafka")
			wireBlock, err := c.client.GetBlock(&hash)
			if err != nil {
				msg := fmt.Sprintf("BTC: Failed to get WireBlock %v", hash)
				log.Error().Err(err).Msg(msg)
			}

			verboseBlock, err := c.client.GetBlockVerbose(&hash)
			if err != nil {
				msg := fmt.Sprintf("BTC: Failed to get VerboseBlock %v", hash)
				log.Error().Err(err).Msg(msg)
			}

			var blockData Block
			blockData.UnmarshalBTCBlock(verboseBlock)
			c.EventSink <- &blockData

			block := btcutil.NewBlock(wireBlock)
			for _, tx := range block.Transactions() {
				rawTx, err := c.client.GetRawTransactionVerbose(tx.Hash())
				if err != nil {
					log.Error().Err(err).Msg("BTC: Failed to get RawTxVerbose")
				}

				var txData Transaction
				txData.UnmarshalBTCTransaction(rawTx)
				c.EventSink <- &txData
			}
			c.callback()
		}
	}
}
