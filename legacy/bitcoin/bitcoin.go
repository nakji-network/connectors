package bitcoin

import (
	"context"
	"fmt"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/kafkautils"
	"github.com/rs/zerolog/log"
)

type BitcoinConnector struct {
	*connector.Connector // embed Nakji connector.Connector into your custom connector to get access to all its methods
	client               BTCClient
	blockTopic           kafkautils.Topic
	txTopic              kafkautils.Topic
	callback             func()
}

//go:generate mockgen -destination=mocks/mock_client.go -package=mocks . BTCClient
type BTCClient interface {
	GetBlock(*chainhash.Hash) (*wire.MsgBlock, error)
	GetBlockVerbose(*chainhash.Hash) (*btcjson.GetBlockVerboseResult, error)
	GetRawTransactionVerbose(*chainhash.Hash) (*btcjson.TxRawResult, error)
	Shutdown()
	WaitForShutdown()
}

const Namespace = "bitcoin"

func NewConnector(client BTCClient, blockTopic, txTopic kafkautils.Topic, callback func()) *BitcoinConnector {
	c := connector.NewConnector()
	return &BitcoinConnector{
		Connector:  c,
		client:     client,
		blockTopic: blockTopic,
		txTopic:    txTopic,
		callback:   callback,
	}
}

func (c BitcoinConnector) Start(ctx context.Context, hashChan <-chan chainhash.Hash) {
	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("BTC: transactions worker cancelled and shutting down")
			c.client.Shutdown()
			return

		case hash := <-hashChan:
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
			c.ProduceAndCommitMessage(Namespace, verboseBlock.Hash, &blockData)

			block := btcutil.NewBlock(wireBlock)
			for _, tx := range block.Transactions() {
				rawTx, err := c.client.GetRawTransactionVerbose(tx.Hash())
				if err != nil {
					log.Error().Err(err).Msg("BTC: Failed to get RawTxVerbose")
				}

				var txData Transaction
				txData.UnmarshalBTCTransaction(rawTx)
				c.ProduceAndCommitMessage(Namespace, txData.Txid, &txData)
			}
			c.callback()
		}
	}
}
