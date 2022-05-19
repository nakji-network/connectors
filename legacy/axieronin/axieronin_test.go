package axieronin

import (
	"context"
	"strings"
	"testing"

	ethmocks "blep.ai/data/chain/ethereum/ethclient/mocks"
	"blep.ai/data/connectors/source/axie/axienft"
	"blep.ai/data/connectors/source/axie/axs"
	"blep.ai/data/connectors/source/axie/slp"
	"blep.ai/data/connectors/source/axieronin/clockauction"
	"blep.ai/data/connectors/source/axieronin/land"
	"blep.ai/data/connectors/source/axieronin/landitem"
	"blep.ai/data/connectors/source/axieronin/roninweth"
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
		"axienft_approval":                   kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.axienft_approval", "test"),
		"axienft_approvalforall":             kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.axienft_approvalforall", "test"),
		"axienft_axieevolved":                kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.axienft_axieevolved", "test"),
		"axienft_axierebirthed":              kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.axienft_axierebirthed", "test"),
		"axienft_axieretired":                kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.axienft_axieretired", "test"),
		"axienft_axiespawned":                kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.axienft_axiespawned", "test"),
		"axienft_transfer":                   kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.axienft_transfer", "test"),
		"axs_approval":                       kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.axs_approval", "test"),
		"axs_transfer":                       kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.axs_transfer", "test"),
		"clockauction_adminchanged":          kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.clockauction_adminchanged", "test"),
		"clockauction_adminremoved":          kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.clockauction_adminremoved", "test"),
		"clockauction_auctioncancelled":      kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.clockauction_auctioncancelled", "test"),
		"clockauction_auctioncreated":        kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.clockauction_auctioncreated", "test"),
		"clockauction_auctionsuccessful":     kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.clockauction_auctionsuccessful", "test"),
		"clockauction_paused":                kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.clockauction_paused", "test"),
		"clockauction_tokenauctioncancelled": kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.clockauction_tokenauctioncancelled", "test"),
		"clockauction_unpaused":              kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.clockauction_unpaused", "test"),
		"land_adminchanged":                  kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.land_adminchanged", "test"),
		"land_adminremoved":                  kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.land_adminremoved", "test"),
		"land_approval":                      kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.land_approval", "test"),
		"land_approvalforall":                kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.land_approvalforall", "test"),
		"land_minteradded":                   kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.land_minteradded", "test"),
		"land_minterremoved":                 kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.land_minterremoved", "test"),
		"land_nonceupdated":                  kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.land_nonceupdated", "test"),
		"land_paused":                        kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.land_paused", "test"),
		"land_spenderunwhitelisted":          kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.land_spenderunwhitelisted", "test"),
		"land_spenderwhitelisted":            kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.land_spenderwhitelisted", "test"),
		"land_transfer":                      kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.land_transfer", "test"),
		"land_unpaused":                      kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.land_unpaused", "test"),
		"landitem_adminchanged":              kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.landitem_adminchanged", "test"),
		"landitem_adminremoved":              kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.landitem_adminremoved", "test"),
		"landitem_approval":                  kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.landitem_approval", "test"),
		"landitem_approvalforall":            kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.landitem_approvalforall", "test"),
		"landitem_minteradded":               kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.landitem_minteradded", "test"),
		"landitem_minterremoved":             kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.landitem_minterremoved", "test"),
		"landitem_nonceupdated":              kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.landitem_nonceupdated", "test"),
		"landitem_paused":                    kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.landitem_paused", "test"),
		"landitem_spenderunwhitelisted":      kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.landitem_spenderunwhitelisted", "test"),
		"landitem_spenderwhitelisted":        kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.landitem_spenderwhitelisted", "test"),
		"landitem_transfer":                  kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.landitem_transfer", "test"),
		"landitem_unpaused":                  kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.landitem_unpaused", "test"),
		"roninweth_approval":                 kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.roninweth_approval", "test"),
		"roninweth_minteradded":              kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.roninweth_minteradded", "test"),
		"roninweth_minterremoved":            kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.roninweth_minterremoved", "test"),
		"roninweth_transfer":                 kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.roninweth_transfer", "test"),
		"slp_adminchanged":                   kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.slp_adminchanged", "test"),
		"slp_adminremoved":                   kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.slp_adminremoved", "test"),
		"slp_approval":                       kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.slp_approval", "test"),
		"slp_minteradded":                    kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.slp_minteradded", "test"),
		"slp_minterremoved":                  kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.slp_minterremoved", "test"),
		"slp_transfer":                       kafkautils.MustParseTopic(".fct.nakji.axieronin.0_0_0.slp_transfer", "test"),
	}

	axienftAddress := ethcommon.HexToAddress("0x32950db2a7164ae833121501c797d79e7b79d74c")
	axsAddress := ethcommon.HexToAddress("0x97a9107c1793bc407d6f527b77e7fff4d812bece")
	clockauctionAddress := ethcommon.HexToAddress("0x213073989821f738a7ba3520c3d31a1f9ad31bbd")
	landAddress := ethcommon.HexToAddress("0x8c811e3c958e190f5ec15fb376533a3398620500")
	landitemAddress := ethcommon.HexToAddress("0xa96660f0e4a3e9bc7388925d245a6d4d79e21259")
	roninwethAddress := ethcommon.HexToAddress("0xc99a6a985ed2cac1ef41640596c5a5f9f4e19ef5")
	slpAddress := ethcommon.HexToAddress("0xa8754b9fa15fc18bb59458815510e40a12cd2014")

	addresses := map[string][]ethcommon.Address{
		"axienft":      {axienftAddress},
		"axs":          {axsAddress},
		"clockauction": {clockauctionAddress},
		"land":         {landAddress},
		"landitem":     {landitemAddress},
		"roninweth":    {roninwethAddress},
		"slp":          {slpAddress},
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

func (c *AxieroninConnector) setupMocks(cancel context.CancelFunc) {
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

	clockauctionQuery := geth.FilterQuery{
		Addresses: c.addresses["clockauction"],
	}
	mockETHClientPool.EXPECT().SubscribeFilterLogs(gomock.Any(), clockauctionQuery, gomock.Any()).DoAndReturn(
		func(ctx context.Context, query geth.FilterQuery, eventLogs chan<- ethtypes.Log) (event.Subscription, error) {
			clockauctionAbi, err := abi.JSON(strings.NewReader(clockauction.ClockauctionABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Clockauction ABI")
			}

			go tests.WriteEventLogs(eventLogs, &clockauctionAbi)
			return mockSubscription, nil
		},
	)

	landQuery := geth.FilterQuery{
		Addresses: c.addresses["land"],
	}
	mockETHClientPool.EXPECT().SubscribeFilterLogs(gomock.Any(), landQuery, gomock.Any()).DoAndReturn(
		func(ctx context.Context, query geth.FilterQuery, eventLogs chan<- ethtypes.Log) (event.Subscription, error) {
			landAbi, err := abi.JSON(strings.NewReader(land.LandABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Land ABI")
			}

			go tests.WriteEventLogs(eventLogs, &landAbi)
			return mockSubscription, nil
		},
	)

	landitemQuery := geth.FilterQuery{
		Addresses: c.addresses["landitem"],
	}
	mockETHClientPool.EXPECT().SubscribeFilterLogs(gomock.Any(), landitemQuery, gomock.Any()).DoAndReturn(
		func(ctx context.Context, query geth.FilterQuery, eventLogs chan<- ethtypes.Log) (event.Subscription, error) {
			landitemAbi, err := abi.JSON(strings.NewReader(landitem.LanditemABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Landitem ABI")
			}

			go tests.WriteEventLogs(eventLogs, &landitemAbi)
			return mockSubscription, nil
		},
	)

	roninwethQuery := geth.FilterQuery{
		Addresses: c.addresses["roninweth"],
	}
	mockETHClientPool.EXPECT().SubscribeFilterLogs(gomock.Any(), roninwethQuery, gomock.Any()).DoAndReturn(
		func(ctx context.Context, query geth.FilterQuery, eventLogs chan<- ethtypes.Log) (event.Subscription, error) {
			roninwethAbi, err := abi.JSON(strings.NewReader(roninweth.RoninwethABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Roninweth ABI")
			}

			go tests.WriteEventLogs(eventLogs, &roninwethAbi)
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

					if counter >= 51 {
						cancel()
					}
				}
			}()
		},
	)
}
