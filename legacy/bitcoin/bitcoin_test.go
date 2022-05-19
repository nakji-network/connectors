package bitcoin

// import (
// 	"context"
// 	"crypto/rand"
// 	"testing"

// 	"github.com/btcsuite/btcd/btcjson"
// 	"github.com/btcsuite/btcd/chaincfg/chainhash"
// 	"github.com/btcsuite/btcd/wire"
// 	"github.com/golang/mock/gomock"

// 	"blep.ai/data/connectors/source/bitcoin/mocks"
// 	"github.com/nakji-network/connector/kafkautils"
// )

// const (
// 	kafkaUrl   = "localhost:9092"
// 	BLOCKTOPIC = ".fct.nakji.bitcoin.0_0_0.chain_block"
// 	TXTOPIC    = ".fct.nakji.bitcoin.0_0_0.chain_tx"
// )

// func TestStartConnector(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)

// 	kafkautils.TopicTypeRegistry.Load(TopicTypes)
// 	blockTopic := kafkautils.MustParseTopic(BLOCKTOPIC, "test")
// 	txTopic := kafkautils.MustParseTopic(TXTOPIC, "test")

// 	randBytes := make([]byte, 32)
// 	rand.Read(randBytes)
// 	hash, err := chainhash.NewHash(randBytes)
// 	if err != nil {
// 		t.Fatalf("New Hash failed")
// 	}

// 	mockClient := mocks.NewMockBTCClient(mockCtrl)
// 	ctx, cancel := context.WithCancel(context.Background())
// 	connector := NewConnector(mockClient, blockTopic, txTopic, func() {
// 		cancel()
// 	})

// 	connector.Connector.Producer =

// 	setupMocks(connector, hash)

// 	hashChan := make(chan chainhash.Hash, 5)
// 	hashChan <- *hash

// 	connector.Start(ctx, hashChan)
// }

// func setupMocks(c *BitcoinConnector, hash *chainhash.Hash) {
// 	mockClient := c.client.(*mocks.MockBTCClient)
// 	mockClient.EXPECT().Shutdown().AnyTimes()

// 	blockHeader := wire.NewBlockHeader(1, hash, hash, 32, 32)
// 	wireBlock := wire.NewMsgBlock(blockHeader)
// 	wireBlock.Transactions = []*wire.MsgTx{
// 		&wire.MsgTx{
// 			Version:  1,
// 			TxIn:     []*wire.TxIn{},
// 			TxOut:    []*wire.TxOut{},
// 			LockTime: 1,
// 		},
// 	}

// 	verboseBlock := &btcjson.GetBlockVerboseResult{
// 		Hash:          hash.String(),
// 		Confirmations: 1,
// 		StrippedSize:  1,
// 		Size:          1,
// 	}

// 	verboseTx := &btcjson.TxRawResult{
// 		Hex:  "test",
// 		Txid: "testtxid",
// 	}

// 	mockClient.EXPECT().GetBlock(gomock.Any()).Times(1).Return(wireBlock, nil)
// 	mockClient.EXPECT().GetBlockVerbose(gomock.Any()).AnyTimes().Return(verboseBlock, nil)
// 	mockClient.EXPECT().GetRawTransactionVerbose(gomock.Any()).AnyTimes().Return(verboseTx, nil)
// }
