package opensea

import (
	"context"
	"strings"
	"testing"

	ethmocks "blep.ai/data/chain/ethereum/ethclient/mocks"
	"blep.ai/data/connectors/source/opensea/wyvernexchangev1"
	"blep.ai/data/connectors/source/opensea/wyvernexchangev2"
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
		"wyvernexchangev1_orderapprovedparttwo": kafkautils.MustParseTopic(".fct.nakji.opensea.0_0_0.wyvernexchangev1_orderapprovedparttwo", "test"),
		"wyvernexchangev1_orderapprovedpartone": kafkautils.MustParseTopic(".fct.nakji.opensea.0_0_0.wyvernexchangev1_orderapprovedpartone", "test"),
		"wyvernexchangev1_ordercancelled":       kafkautils.MustParseTopic(".fct.nakji.opensea.0_0_0.wyvernexchangev1_ordercancelled", "test"),
		"wyvernexchangev1_ordersmatched":        kafkautils.MustParseTopic(".fct.nakji.opensea.0_0_0.wyvernexchangev1_ordersmatched", "test"),
		"wyvernexchangev1_ownershiprenounced":   kafkautils.MustParseTopic(".fct.nakji.opensea.0_0_0.wyvernexchangev1_ownershiprenounced", "test"),
		"wyvernexchangev1_ownershiptransferred": kafkautils.MustParseTopic(".fct.nakji.opensea.0_0_0.wyvernexchangev1_ownershiptransferred", "test"),
		"wyvernexchangev2_nonceincremented":     kafkautils.MustParseTopic(".fct.nakji.opensea.0_0_0.wyvernexchangev2_nonceincremented", "test"),
		"wyvernexchangev2_orderapprovedpartone": kafkautils.MustParseTopic(".fct.nakji.opensea.0_0_0.wyvernexchangev2_orderapprovedpartone", "test"),
		"wyvernexchangev2_orderapprovedparttwo": kafkautils.MustParseTopic(".fct.nakji.opensea.0_0_0.wyvernexchangev2_orderapprovedparttwo", "test"),
		"wyvernexchangev2_ordercancelled":       kafkautils.MustParseTopic(".fct.nakji.opensea.0_0_0.wyvernexchangev2_ordercancelled", "test"),
		"wyvernexchangev2_ordersmatched":        kafkautils.MustParseTopic(".fct.nakji.opensea.0_0_0.wyvernexchangev2_ordersmatched", "test"),
		"wyvernexchangev2_ownershiprenounced":   kafkautils.MustParseTopic(".fct.nakji.opensea.0_0_0.wyvernexchangev2_ownershiprenounced", "test"),
		"wyvernexchangev2_ownershiptransferred": kafkautils.MustParseTopic(".fct.nakji.opensea.0_0_0.wyvernexchangev2_ownershiptransferred", "test"),
	}

	wyvernexchangev1Address := ethcommon.HexToAddress("0x7Be8076f4EA4A4AD08075C2508e481d6C946D12b")
	wyvernexchangev2Address := ethcommon.HexToAddress("0x7f268357A8c2552623316e2562D90e642bB538E5")

	addresses := map[string][]ethcommon.Address{
		"wyvernexchangev1": {wyvernexchangev1Address},
		"wyvernexchangev2": {wyvernexchangev2Address},
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

func (c *OpenseaConnector) setupMocks(cancel context.CancelFunc) {
	mockETHClientPool := c.ClientPool.(*ethmocks.MockETHClientPool)
	mockProducer := c.KP.(*kafkamocks.MockProducerInterface)

	mockETHClientPool.EXPECT().ConsumeHeaders(gomock.Any())

	mockSubscription := MockSubscription{
		errorChan: make(chan error),
	}

	mockETHClientPool.EXPECT().ChunkedFilterLogs(context.Background(), gomock.Any(), int64(100), int64(1), gomock.Any(), gomock.Any()).AnyTimes()
	mockETHClientPool.EXPECT().GetLogTimestamp(gomock.Any(), c.blockCache).AnyTimes().Return(uint64(1000), nil)

	wyvernexchangev1Query := geth.FilterQuery{
		Addresses: c.addresses["wyvernexchangev1"],
	}
	mockETHClientPool.EXPECT().SubscribeFilterLogs(gomock.Any(), wyvernexchangev1Query, gomock.Any()).DoAndReturn(
		func(ctx context.Context, query geth.FilterQuery, eventLogs chan<- ethtypes.Log) (event.Subscription, error) {
			wyvernexchangev1Abi, err := abi.JSON(strings.NewReader(wyvernexchangev1.Wyvernexchangev1ABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Wyvernexchangev1 ABI")
			}

			go tests.WriteEventLogs(eventLogs, &wyvernexchangev1Abi)
			return mockSubscription, nil
		},
	)

	wyvernexchangev2Query := geth.FilterQuery{
		Addresses: c.addresses["wyvernexchangev2"],
	}
	mockETHClientPool.EXPECT().SubscribeFilterLogs(gomock.Any(), wyvernexchangev2Query, gomock.Any()).DoAndReturn(
		func(ctx context.Context, query geth.FilterQuery, eventLogs chan<- ethtypes.Log) (event.Subscription, error) {
			wyvernexchangev2Abi, err := abi.JSON(strings.NewReader(wyvernexchangev2.Wyvernexchangev2ABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Wyvernexchangev2 ABI")
			}

			go tests.WriteEventLogs(eventLogs, &wyvernexchangev2Abi)
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

					if counter >= 13 {
						cancel()
					}
				}
			}()
		},
	)
}
