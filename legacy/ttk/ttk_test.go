package ttk

import (
	"context"
	"strings"
	"testing"

	ethmocks "blep.ai/data/chain/ethereum/ethclient/mocks"
	"blep.ai/data/connectors/source/ttk/ttk"
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
		"ttk_approval": kafkautils.MustParseTopic(".fct.nakji.ttk.0_0_0.ttk_approval", "test"),
		"ttk_transfer": kafkautils.MustParseTopic(".fct.nakji.ttk.0_0_0.ttk_transfer", "test"),
	}

	ttkAddress := ethcommon.HexToAddress("0x39703A67bAC0E39f9244d97f4c842D15Fbad9C1f")

	addresses := map[string][]ethcommon.Address{
		"ttk": {ttkAddress},
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

func (c *TtkConnector) setupMocks(cancel context.CancelFunc) {
	mockETHClientPool := c.ClientPool.(*ethmocks.MockETHClientPool)
	mockProducer := c.KP.(*kafkamocks.MockProducerInterface)

	mockETHClientPool.EXPECT().ConsumeHeaders(gomock.Any())

	mockSubscription := MockSubscription{
		errorChan: make(chan error),
	}

	mockETHClientPool.EXPECT().ChunkedFilterLogs(context.Background(), gomock.Any(), int64(100), int64(1), gomock.Any(), gomock.Any()).AnyTimes()
	mockETHClientPool.EXPECT().GetLogTimestamp(gomock.Any(), c.blockCache).AnyTimes().Return(uint64(1000), nil)

	ttkQuery := geth.FilterQuery{
		Addresses: c.addresses["ttk"],
	}
	mockETHClientPool.EXPECT().SubscribeFilterLogs(gomock.Any(), ttkQuery, gomock.Any()).DoAndReturn(
		func(ctx context.Context, query geth.FilterQuery, eventLogs chan<- ethtypes.Log) (event.Subscription, error) {
			ttkAbi, err := abi.JSON(strings.NewReader(ttk.TtkABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Ttk ABI")
			}

			go tests.WriteEventLogs(eventLogs, &ttkAbi)
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

					if counter >= 2 {
						cancel()
					}
				}
			}()
		},
	)
}
