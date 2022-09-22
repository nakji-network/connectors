package eth2

import (
	"context"
	"strings"
	"testing"

	ethmocks "blep.ai/data/chain/ethereum/ethclient/mocks"
	"blep.ai/data/connectors/source/eth2/eth2"
	"blep.ai/data/connectors/tests"
	"github.com/nakji-network/connector/kafkautils"
	kafkamocks "github.com/nakji-network/connector/kafkautils/mocks"

	geth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog/log"
)

type MockSubscription struct {
	errorChan chan error
}

func (m MockSubscription) Err() <-chan error {
	return m.errorChan
}

func (MockSubscription) Unsubscribe() {
	// noop
}

func TestStartConnector(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProducer := kafkamocks.NewMockProducerInterface(mockCtrl)
	mockETHClientPool := ethmocks.NewMockETHClientPool(mockCtrl)

	kafkautils.TopicTypeRegistry.Load(TopicTypes)

	topics := map[string]kafkautils.Topic{
		"eth2_depositevent": kafkautils.MustParseTopic(".fct.nakji.eth2.0_0_0.eth2_depositevent", "test"),
	}

	eth2Address := ethcommon.HexToAddress("0x00000000219ab540356cBB839Cbe05303d7705Fa")

	addresses := map[string][]ethcommon.Address{
		"eth2": {eth2Address},
	}

	connector := NewConnector(
		mockProducer,
		addresses,
		topics,
		mockETHClientPool,
	)

	ctx, cancel := context.WithCancel(context.Background())
	connector.setupMocks(cancel)
	connector.Start(ctx, 10)
}

func (c *Eth2Connector) setupMocks(cancel context.CancelFunc) {
	mockETHClientPool := c.ClientPool.(*ethmocks.MockETHClientPool)
	mockProducer := c.KP.(*kafkamocks.MockProducerInterface)

	mockETHClientPool.EXPECT().ConsumeHeaders(gomock.Any())

	mockSubscription := MockSubscription{
		errorChan: make(chan error),
	}

	mockETHClientPool.EXPECT().ChunkedFilterLogs(context.Background(), gomock.Any(), int64(100), int64(1), gomock.Any(), gomock.Any()).AnyTimes()
	mockETHClientPool.EXPECT().GetLogTimestamp(gomock.Any(), c.blockCache).AnyTimes().Return(uint64(1000), nil)

	eth2Query := geth.FilterQuery{
		Addresses: c.addresses["eth2"],
	}
	mockETHClientPool.EXPECT().SubscribeFilterLogs(gomock.Any(), eth2Query, gomock.Any()).DoAndReturn(
		func(ctx context.Context, query geth.FilterQuery, eventLogs chan<- ethtypes.Log) (event.Subscription, error) {
			eth2Abi, err := abi.JSON(strings.NewReader(eth2.Eth2ABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Eth2 ABI")
			}

			go tests.WriteEventLogs(eventLogs, &eth2Abi)
			return mockSubscription, nil
		},
	)

	mockProducer.EXPECT().EnableTransactions()
	counter := 0
	mockProducer.EXPECT().WriteAndCommitSink(gomock.Any()).DoAndReturn(
		func(sink <-chan *kafkautils.Message) {
			go func() {
				for range sink {
					counter++

					if counter >= 1 {
						cancel()
					}
				}
			}()
		},
	)
}
