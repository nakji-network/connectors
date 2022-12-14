// subscribe to new Blocks and Transactions and writes the results to Nakji.

// It also works for other evm-compatible chains, as long as they use the ethclient.Client.
package klaytn

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/klaytn/klaytn/blockchain/types"
	klayclient "github.com/klaytn/klaytn/client"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

type Connector struct {
	*connector.Connector // embed Nakji connector.Connector into your custom connector to get access to all its methods

	// Any additional custom connections not supported natively by Nakji, replace it as you see fit.
	// eg: client: DogecoinClient(context.Background()),
	client *klayclient.Client

	// Any additional config vars from the config yaml, using functions from Viper (https://pkg.go.dev/github.com/spf13/viper#readme-getting-values-from-viper)
	// This is namespaced via connector id (author-name-version)
	// CustomOption: c.Config.GetString("custom_option"),
}

func NewConnector(chain string) *Connector {
	c, err := connector.NewConnector()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to instantiate nakji connector")
	}

	// Read config from config yaml under `rpcs.[chain].full`
	rpcs := c.RPCMap[chain].Full

	// go-klaytn client only supports 1 rpc connection currently, so we do this hack
	var RPCURL string
	for _, u := range rpcs {
		if strings.HasPrefix(u, "ws") {
			RPCURL = u
			break
		}
	}
	log.Info().
		Str("chain", chain).
		Str("url", RPCURL).
		Msg("connecting to RPC")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := klayclient.DialContext(ctx, RPCURL)
	if err != nil {
		log.Fatal().Err(err).Msg("RPC connection error")
	}

	return &Connector{
		c,
		client,
	}
}

func (c *Connector) Start() {
	// Register topic and protobuf type mappings
	c.RegisterProtos(kafkautils.MsgTypeFct, protos...)

	// Listen for interrupt in order to cleanly close connections later
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Subscribe to headers
	headers := make(chan *types.Header, 100)
	sub, err := c.client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal().Err(err)
	}

	// Main loop to process errors and headers
	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatal().Err(err)
			case header := <-headers:
				// Header doesn't contain full block information, so get block
				block, err := c.client.BlockByHash(context.Background(), header.Hash())
				if err != nil {
					log.Fatal().Err(err).Msg("BlockByHash error: " + header.Hash().Hex() + " " + header.Number.String())
				}

				LogBlock(block)

				// KlaytnBlock -> Block -> Protobuf -> kafka
				var blockData Block
				blockData.UnmarshalKlaytnBlock(block)

				c.EventSink <- buildKafkaMsg(kafkautils.MsgTypeFct, &blockData)

				// KlaytnTransaction -> Transaction -> Protobuf -> Kafka
				for _, tx := range block.Transactions() {

					txData := Transaction{}
					txData.UnmarshalKlaytnTransaction(tx)
					txData.Ts = blockData.Ts // Timestamp isn't in the raw transaction from geth

					c.EventSink <- buildKafkaMsg(kafkautils.MsgTypeFct, &txData)
				}
			}
		}
	}()

	for {
		select {
		case <-interrupt:
			log.Info().Msg("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			c.client.Close()
			c.Close()
			return
		}
	}
}

func buildKafkaMsg(msgType kafkautils.MsgType, msg proto.Message) *kafkautils.Message {
	return &kafkautils.Message{
		MsgType:  msgType,
		ProtoMsg: msg,
	}
}

func LogBlock(block *types.Block) {
	log.Debug().
		Str("hash", block.Hash().Hex()).              // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
		Uint64("num", block.Number().Uint64()).       // 3477413
		Uint64("time", block.Time().Uint64()).        // 1529525947
		Int("#tx", len(block.Transactions())).        // 7
		Uint64("gau used", block.GasUsed()).          // 1529525947
		Interface("block score", block.BlockScore()). // 1529525947
		Msg("new block")
}
