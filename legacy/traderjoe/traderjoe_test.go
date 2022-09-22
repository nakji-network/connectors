package traderjoe

import (
	"context"
	"math/big"
	"strings"
	"testing"

	"blep.ai/data/chain/avalanche/avaxclient"
	"blep.ai/data/connectors/source/evm"
	"blep.ai/data/connectors/source/traderjoe/joefactory"
	"blep.ai/data/connectors/source/traderjoe/joepair"
	uniswapv2mocks "blep.ai/data/connectors/source/uniswapv2/mocks"
	"github.com/ava-labs/coreth/accounts/abi"
	"github.com/ava-labs/coreth/interfaces"

	"blep.ai/data/connectors/source/traderjoe/joebar"
	"blep.ai/data/connectors/source/traderjoe/joehattoken"
	"blep.ai/data/connectors/source/traderjoe/joetoken"
	"blep.ai/data/connectors/source/traderjoe/joetroller"
	"blep.ai/data/connectors/source/traderjoe/masterchefjoev2"
	"blep.ai/data/connectors/source/traderjoe/masterchefjoev3"
	"blep.ai/data/database"
	"github.com/nakji-network/connector/kafkautils"
	kafkamocks "github.com/nakji-network/connector/kafkautils/mocks"

	"github.com/ava-labs/coreth/core/types"
	"github.com/ethereum/go-ethereum/common"
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
	mockETHClientPool := avaxclient.NewMockETHClientPool(mockCtrl)
	mockTokenCache := uniswapv2mocks.NewMockTokenCacheInterface(mockCtrl)

	db, err := database.New("host=localhost port=5432 user=postgres password=password database=blep_test")
	if err != nil {
		log.Fatal().Err(err).Msg("Timescaledb connection failed")
	}
	defer db.Close()

	kafkautils.TopicTypeRegistry.Load(TopicTypes)

	topics := map[string]kafkautils.Topic{
		"joebar_approval":                      kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joebar.approval.0_0_0", "test"),
		"joebar_transfer":                      kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joebar.transfer.0_0_0", "test"),
		"joefactory_paircreated":               kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joefactory.paircreated.0_0_0", "test"),
		"joehattoken_approval":                 kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joehattoken.approval.0_0_0", "test"),
		"joehattoken_transfer":                 kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joehattoken.transfer.0_0_0", "test"),
		"joepair_approval":                     kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joepair.approval.0_0_0", "test"),
		"joepair_burn":                         kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joepair.burn.0_0_0", "test"),
		"joepair_mint":                         kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joepair.mint.0_0_0", "test"),
		"joepair_swap":                         kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joepair.swap.0_0_0", "test"),
		"joepair_sync":                         kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joepair.sync.0_0_0", "test"),
		"joepair_transfer":                     kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joepair.transfer.0_0_0", "test"),
		"joetoken_approval":                    kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joetoken.approval.0_0_0", "test"),
		"joetoken_delegatechanged":             kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joetoken.delegatechanged.0_0_0", "test"),
		"joetoken_delegatevoteschanged":        kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joetoken.delegatevoteschanged.0_0_0", "test"),
		"joetoken_ownershiptransferred":        kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joetoken.ownershiptransferred.0_0_0", "test"),
		"joetoken_transfer":                    kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joetoken.transfer.0_0_0", "test"),
		"joetroller_failure":                   kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joetroller.failure.0_0_0", "test"),
		"joetroller_newadmin":                  kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joetroller.newadmin.0_0_0", "test"),
		"joetroller_newimplementation":         kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joetroller.newimplementation.0_0_0", "test"),
		"joetroller_newpendingadmin":           kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joetroller.newpendingadmin.0_0_0", "test"),
		"joetroller_newpendingimplementation":  kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.joetroller.newpendingimplementation.0_0_0", "test"),
		"masterchefjoev2_add":                  kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.masterchefjoev2.add.0_0_0", "test"),
		"masterchefjoev2_deposit":              kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.masterchefjoev2.deposit.0_0_0", "test"),
		"masterchefjoev2_emergencywithdraw":    kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.masterchefjoev2.emergencywithdraw.0_0_0", "test"),
		"masterchefjoev2_harvest":              kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.masterchefjoev2.harvest.0_0_0", "test"),
		"masterchefjoev2_ownershiptransferred": kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.masterchefjoev2.ownershiptransferred.0_0_0", "test"),
		"masterchefjoev2_set":                  kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.masterchefjoev2.set.0_0_0", "test"),
		"masterchefjoev2_setdevaddress":        kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.masterchefjoev2.setdevaddress.0_0_0", "test"),
		"masterchefjoev2_updateemissionrate":   kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.masterchefjoev2.updateemissionrate.0_0_0", "test"),
		"masterchefjoev2_updatepool":           kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.masterchefjoev2.updatepool.0_0_0", "test"),
		"masterchefjoev2_withdraw":             kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.masterchefjoev2.withdraw.0_0_0", "test"),
		"masterchefjoev3_add":                  kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.masterchefjoev3.add.0_0_0", "test"),
		"masterchefjoev3_deposit":              kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.masterchefjoev3.deposit.0_0_0", "test"),
		"masterchefjoev3_emergencywithdraw":    kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.masterchefjoev3.emergencywithdraw.0_0_0", "test"),
		"masterchefjoev3_harvest":              kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.masterchefjoev3.harvest.0_0_0", "test"),
		"masterchefjoev3_init":                 kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.masterchefjoev3.init.0_0_0", "test"),
		"masterchefjoev3_ownershiptransferred": kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.masterchefjoev3.ownershiptransferred.0_0_0", "test"),
		"masterchefjoev3_set":                  kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.masterchefjoev3.set.0_0_0", "test"),
		"masterchefjoev3_updatepool":           kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.masterchefjoev3.updatepool.0_0_0", "test"),
		"masterchefjoev3_withdraw":             kafkautils.MustParseTopic(".fct.nakji.traderjoe.0_0_0.masterchefjoev3.withdraw.0_0_0", "test"),
	}

	joebarAddress := common.HexToAddress("0x57319d41F71E81F3c65F2a47CA4e001EbAFd4F33")
	joefactoryAddress := common.HexToAddress("0x9Ad6C38BE94206cA50bb0d90783181662f0Cfa10")
	joehattokenAddress := common.HexToAddress("0x82FE038Ea4b50f9C957da326C412ebd73462077C")

	joetokenAddress := common.HexToAddress("0x6e84a6216eA6dACC71eE8E6b0a5B7322EEbC0fDd")
	joetrollerAddress := common.HexToAddress("0xdc13687554205E5b89Ac783db14bb5bba4A1eDaC")
	masterchefjoev2Address := common.HexToAddress("0xd6a4F121CA35509aF06A0Be99093d08462f53052")
	masterchefjoev3Address := common.HexToAddress("0x188bED1968b795d5c9022F6a0bb5931Ac4c18F00")

	addresses := map[string][]common.Address{
		"joebar":      {joebarAddress},
		"joefactory":  {joefactoryAddress},
		"joehattoken": {joehattokenAddress},
		"joepair":     {common.Address{}},

		"joetoken":        {joetokenAddress},
		"joetroller":      {joetrollerAddress},
		"masterchefjoev2": {masterchefjoev2Address},
		"masterchefjoev3": {masterchefjoev3Address},
	}

	connector := NewConnector(
		mockProducer,
		addresses,
		topics,
		mockETHClientPool,
		db,
		mockTokenCache,
	)

	ctx, cancel := context.WithCancel(context.Background())
	connector.setupMocks(cancel)
	connector.Start(ctx, 10)
}

func (c *Connector) setupMocks(cancel context.CancelFunc) {
	mockETHClientPool := c.AvaxClientPool.(*avaxclient.MockETHClientPool)
	mockProducer := c.KP.(*kafkamocks.MockProducerInterface)
	mockTokenCache := c.tokenCache.(*uniswapv2mocks.MockTokenCacheInterface)

	mockETHClientPool.EXPECT().ConsumeHeaders(gomock.Any())
	mockTokenCache.EXPECT().Get(gomock.Any()).Times(2).Return(&evm.ERC20{}, true)

	mockSubscription := MockSubscription{
		errorChan: make(chan error),
	}

	mockETHClientPool.EXPECT().BlockNumber(context.Background()).AnyTimes()
	mockETHClientPool.EXPECT().ChunkedFilterLogs(context.Background(), gomock.Any(), int64(100), int64(1), gomock.Any(), gomock.Any()).AnyTimes()
	mockETHClientPool.EXPECT().GetLogTimestamp(gomock.Any(), c.blockCache).AnyTimes().Return(uint64(1000), nil)

	joepairQuery := interfaces.FilterQuery{
		Addresses: c.addresses["joepair"],
	}
	mockETHClientPool.EXPECT().SubscribeFilterLogs(gomock.Any(), joepairQuery, gomock.Any()).DoAndReturn(
		func(ctx context.Context, query interfaces.FilterQuery, eventLogs chan<- types.Log) (event.Subscription, error) {
			joepairAbi, err := abi.JSON(strings.NewReader(joepair.JoePairMetaData.ABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Joepair ABI")
			}

			go writeEventLogs(eventLogs, &joepairAbi)
			return mockSubscription, nil
		},
	)

	mockETHClientPool.EXPECT().ChunkedSubscribeFilterLogs(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, query interfaces.FilterQuery, eventLogs chan<- types.Log, chunksize int) (<-chan error, error) {
			joebarAbi, err := abi.JSON(strings.NewReader(joebar.JoebarABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Joebar ABI")
			}

			go writeEventLogs(eventLogs, &joebarAbi)
			errchan := make(chan error, 1)
			return errchan, nil
		},
	)

	mockETHClientPool.EXPECT().ChunkedSubscribeFilterLogs(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, query interfaces.FilterQuery, eventLogs chan<- types.Log, chunksize int) (<-chan error, error) {
			joefactoryAbi, err := abi.JSON(strings.NewReader(joefactory.JoefactoryABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Joefactory ABI")
			}

			go writeEventLogs(eventLogs, &joefactoryAbi)
			errchan := make(chan error, 1)
			return errchan, nil
		},
	)

	mockETHClientPool.EXPECT().ChunkedSubscribeFilterLogs(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, query interfaces.FilterQuery, eventLogs chan<- types.Log, chunksize int) (<-chan error, error) {
			joepairAbi, err := abi.JSON(strings.NewReader(joepair.JoePairMetaData.ABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Joepair ABI")
			}

			go writeEventLogs(eventLogs, &joepairAbi)
			errchan := make(chan error, 1)
			return errchan, nil
		},
	)

	mockETHClientPool.EXPECT().ChunkedSubscribeFilterLogs(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, query interfaces.FilterQuery, eventLogs chan<- types.Log, chunksize int) (<-chan error, error) {
			joehattokenAbi, err := abi.JSON(strings.NewReader(joehattoken.JoehattokenABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Joehattoken ABI")
			}

			go writeEventLogs(eventLogs, &joehattokenAbi)
			errchan := make(chan error, 1)
			return errchan, nil
		},
	)

	mockETHClientPool.EXPECT().ChunkedSubscribeFilterLogs(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, query interfaces.FilterQuery, eventLogs chan<- types.Log, chunksize int) (<-chan error, error) {
			joetokenAbi, err := abi.JSON(strings.NewReader(joetoken.JoetokenABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Joetoken ABI")
			}

			go writeEventLogs(eventLogs, &joetokenAbi)
			errchan := make(chan error, 1)
			return errchan, nil
		},
	)

	mockETHClientPool.EXPECT().ChunkedSubscribeFilterLogs(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, query interfaces.FilterQuery, eventLogs chan<- types.Log, chunksize int) (<-chan error, error) {
			joetrollerAbi, err := abi.JSON(strings.NewReader(joetroller.JoetrollerABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Joetroller ABI")
			}

			go writeEventLogs(eventLogs, &joetrollerAbi)
			errchan := make(chan error, 1)
			return errchan, nil
		},
	)

	mockETHClientPool.EXPECT().ChunkedSubscribeFilterLogs(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, query interfaces.FilterQuery, eventLogs chan<- types.Log, chunksize int) (<-chan error, error) {
			masterchefjoev2Abi, err := abi.JSON(strings.NewReader(masterchefjoev2.Masterchefjoev2ABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Masterchefjoev2 ABI")
			}

			go writeEventLogs(eventLogs, &masterchefjoev2Abi)
			errchan := make(chan error, 1)
			return errchan, nil
		},
	)

	mockETHClientPool.EXPECT().ChunkedSubscribeFilterLogs(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, query interfaces.FilterQuery, eventLogs chan<- types.Log, chunksize int) (<-chan error, error) {
			masterchefjoev3Abi, err := abi.JSON(strings.NewReader(masterchefjoev3.Masterchefjoev3ABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Masterchefjoev3 ABI")
			}

			go writeEventLogs(eventLogs, &masterchefjoev3Abi)
			errchan := make(chan error, 1)
			return errchan, nil
		},
	)

	mockETHClientPool.EXPECT().FilterLogs(context.Background(), gomock.Any()).Return([]types.Log{}, nil)
	mockProducer.EXPECT().EnableTransactions()
	counter := 0
	mockProducer.EXPECT().WriteAndCommitSink(gomock.Any()).DoAndReturn(
		func(sink <-chan *kafkautils.Message) {
			go func() {
				for range sink {
					counter++

					if counter >= 40 {
						cancel()
					}
				}
			}()
		},
	)
}

// Helper for getting all events from a contract ABI and writing all those
// events with fake data into a given event logs channel.
// This is the main function we use for emulating subscribing to logs
// of an EVM blockchain.
func writeEventLogs(eventLogs chan<- types.Log, contractAbi *abi.ABI) {
	for _, event := range contractAbi.Events {
		topics := []common.Hash{event.ID}
		data := []byte{}

		for _, arg := range event.Inputs {
			if arg.Indexed {
				hash := common.BigToHash(big.NewInt(0))
				topics = append(topics, hash)
			} else {
				newData := big.NewInt(0).Bytes()
				for len(newData) < 32 {
					newData = append(newData, 0)
				}
				data = append(data, newData...)
			}
		}

		eventLogs <- types.Log{
			Topics: topics,
			Data:   data,
		}
	}
}
