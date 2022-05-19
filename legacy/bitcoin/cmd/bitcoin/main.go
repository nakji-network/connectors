package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"blep.ai/data/config"
	"blep.ai/data/connectors/source/bitcoin"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var (
	BLOCKTOPIC = ".fct.nakji.bitcoin.0_0_0.chain_block"
	TXTOPIC    = ".fct.nakji.bitcoin.0_0_0.chain_tx"
	conf       *viper.Viper
)

func init() {
	conf = config.GetConfig()
	conf.SetDefault("bitcoin.rpc.url", "localhost:8334")
	conf.SetDefault("bitcoin.rpc.user", "blep")
	conf.SetDefault("bitcoin.rpc.password", "password")
}

func main() {
	// Load Topic registry
	kafkautils.TopicTypeRegistry.Load(bitcoin.TopicTypes)
	blockTopic := kafkautils.MustParseTopic(BLOCKTOPIC, conf.GetString("kafka.env"))
	txTopic := kafkautils.MustParseTopic(TXTOPIC, conf.GetString("kafka.env"))

	hashChan := make(chan chainhash.Hash, 100)
	ntfnHandlers := rpcclient.NotificationHandlers{
		OnFilteredBlockConnected: func(height int32, header *wire.BlockHeader, txns []*btcutil.Tx) {
			hashChan <- header.BlockHash()
		},
	}

	// Connect to local btcd RPC server using websockets.
	btcdHomeDir := btcutil.AppDataDir("btcd", false)
	certs, err := ioutil.ReadFile(filepath.Join(btcdHomeDir, "rpc.cert"))
	if err != nil {
		log.Info().Msg("BTC: RPC Certificate not found")
		certs = []byte(conf.GetString("bitcoin.rpc.cert"))
	}
	connCfg := &rpcclient.ConnConfig{
		Host:         conf.GetString("bitcoin.rpc.url"),
		Endpoint:     "ws",
		User:         conf.GetString("bitcoin.rpc.user"),
		Pass:         conf.GetString("bitcoin.rpc.password"),
		Certificates: certs,
	}
	client, err := rpcclient.New(connCfg, &ntfnHandlers)
	if err != nil {
		log.Fatal().Err(err).Msg("BTC: RPC Client connection error")
	}

	if err := client.NotifyBlocks(); err != nil {
		log.Fatal().Err(err).Msg("Block Notifications Register failed")
	}
	log.Info().Msg("BTC: NotifyBlocks: Registration Complete")

	// Get the current block count.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal().Err(err).Msg("BTC: Get Block Count error")
	}
	msg := fmt.Sprintf("Block count: %d", blockCount)
	log.Info().Msg(msg)

	connector := bitcoin.NewConnector(client, blockTopic, txTopic, func() {})
	connector.Start(context.Background(), hashChan)
}
