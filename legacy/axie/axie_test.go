package axie

import (
	"context"
	"strings"
	"testing"

	ethmocks "blep.ai/data/chain/ethereum/ethclient/mocks"
	"blep.ai/data/connectors/source/axie/axienft"
	"blep.ai/data/connectors/source/axie/axs"
	"blep.ai/data/connectors/source/axie/bridge"
	"blep.ai/data/connectors/source/axie/slp"
	"blep.ai/data/connectors/tests"

	"github.com/nakji-network/connector/kafkautils"
	kafkamocks "github.com/nakji-network/connector/kafkautils/mocks"

	geth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/golang/mock/gomock"
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
		"axienft_approval":       kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.axienft_approval", "test"),
		"axienft_approvalforall": kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.axienft_approvalforall", "test"),
		"axienft_axieevolved":    kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.axienft_axieevolved", "test"),
		"axienft_axierebirthed":  kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.axienft_axierebirthed", "test"),
		"axienft_axieretired":    kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.axienft_axieretired", "test"),
		"axienft_axiespawned":    kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.axienft_axiespawned", "test"),
		"axienft_transfer":       kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.axienft_transfer", "test"),
		"axs_approval":           kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.axs_approval", "test"),
		"axs_transfer":           kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.axs_transfer", "test"),
		"bridge_adminchanged":    kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.bridge_adminchanged", "test"),
		"bridge_adminremoved":    kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.bridge_adminremoved", "test"),
		"bridge_paused":          kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.bridge_paused", "test"),
		"bridge_proxyupdated":    kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.bridge_proxyupdated", "test"),
		"bridge_tokendeposited":  kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.bridge_tokendeposited", "test"),
		"bridge_tokenwithdrew":   kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.bridge_tokenwithdrew", "test"),
		"bridge_unpaused":        kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.bridge_unpaused", "test"),
		"slp_adminchanged":       kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.slp_adminchanged", "test"),
		"slp_adminremoved":       kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.slp_adminremoved", "test"),
		"slp_approval":           kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.slp_approval", "test"),
		"slp_minteradded":        kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.slp_minteradded", "test"),
		"slp_minterremoved":      kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.slp_minterremoved", "test"),
		"slp_transfer":           kafkautils.MustParseTopic(".fct.nakji.axie.0_0_0.slp_transfer", "test"),
	}

	axienftAddress := ethcommon.HexToAddress("0xF5b0A3eFB8e8E4c201e2A935F110eAaF3FFEcb8d")
	axsAddress := ethcommon.HexToAddress("0xbb0e17ef65f82ab018d8edd776e8dd940327b28b")
	bridgeAddress := ethcommon.HexToAddress("0x1A2a1c938CE3eC39b6D47113c7955bAa9DD454F2")
	slpAddress := ethcommon.HexToAddress("0xCC8Fa225D80b9c7D42F96e9570156c65D6cAAa25")

	addresses := map[string][]ethcommon.Address{
		"axienft": {axienftAddress},
		"axs":     {axsAddress},
		"bridge":  {bridgeAddress},
		"slp":     {slpAddress},
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

func (c *AxieConnector) setupMocks(cancel context.CancelFunc) {
	mockETHClientPool := c.ClientPool.(*ethmocks.MockETHClientPool)
	mockProducer := c.KP.(*kafkamocks.MockProducerInterface)

	mockETHClientPool.EXPECT().ConsumeHeaders(gomock.Any())

	mockSubscription := MockSubscription{
		errorChan: make(chan error),
	}

	mockETHClientPool.EXPECT().ChunkedFilterLogs(context.Background(), gomock.Any(), int64(100), int64(1), gomock.Any(), gomock.Any()).AnyTimes()
	mockETHClientPool.EXPECT().GetLogTimestamp(gomock.Any(), c.blockCache).AnyTimes().Return(uint64(1000), nil)

	axienftQuery := geth.FilterQuery{
		Addresses: c.addresses["axienft"],
	}
	mockETHClientPool.EXPECT().SubscribeFilterLogs(gomock.Any(), axienftQuery, gomock.Any()).DoAndReturn(
		func(ctx context.Context, query geth.FilterQuery, eventLogs chan<- ethtypes.Log) (event.Subscription, error) {
			axienftAbi, err := abi.JSON(strings.NewReader(axienft.AxienftABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Axienft ABI")
			}

			go tests.WriteEventLogs(eventLogs, &axienftAbi)
			return mockSubscription, nil
		},
	)

	axsQuery := geth.FilterQuery{
		Addresses: c.addresses["axs"],
	}
	mockETHClientPool.EXPECT().SubscribeFilterLogs(gomock.Any(), axsQuery, gomock.Any()).DoAndReturn(
		func(ctx context.Context, query geth.FilterQuery, eventLogs chan<- ethtypes.Log) (event.Subscription, error) {
			axsAbi, err := abi.JSON(strings.NewReader(axs.AxsABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Axs ABI")
			}

			go tests.WriteEventLogs(eventLogs, &axsAbi)
			return mockSubscription, nil
		},
	)

	bridgeQuery := geth.FilterQuery{
		Addresses: c.addresses["bridge"],
	}
	mockETHClientPool.EXPECT().SubscribeFilterLogs(gomock.Any(), bridgeQuery, gomock.Any()).DoAndReturn(
		func(ctx context.Context, query geth.FilterQuery, eventLogs chan<- ethtypes.Log) (event.Subscription, error) {
			bridgeAbi, err := abi.JSON(strings.NewReader(bridge.BridgeABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Bridge ABI")
			}

			go tests.WriteEventLogs(eventLogs, &bridgeAbi)
			return mockSubscription, nil
		},
	)

	slpQuery := geth.FilterQuery{
		Addresses: c.addresses["slp"],
	}
	mockETHClientPool.EXPECT().SubscribeFilterLogs(gomock.Any(), slpQuery, gomock.Any()).DoAndReturn(
		func(ctx context.Context, query geth.FilterQuery, eventLogs chan<- ethtypes.Log) (event.Subscription, error) {
			slpAbi, err := abi.JSON(strings.NewReader(slp.SlpABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Slp ABI")
			}

			go tests.WriteEventLogs(eventLogs, &slpAbi)
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

					if counter >= 22 {
						cancel()
					}
				}
			}()
		},
	)
}
