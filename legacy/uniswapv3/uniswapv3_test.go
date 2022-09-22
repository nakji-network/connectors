package uniswapv3

import (
	"context"
	"strings"
	"testing"

	ethmocks "blep.ai/data/chain/ethereum/ethclient/mocks"
	"blep.ai/data/connectors/source/evm"
	uniswapv2mocks "blep.ai/data/connectors/source/uniswapv2/mocks"
	"blep.ai/data/connectors/source/uniswapv3/factory"
	"blep.ai/data/connectors/source/uniswapv3/pool"
	"blep.ai/data/connectors/tests"
	"blep.ai/data/database"
	"github.com/nakji-network/connector/kafkautils"
	kafkamocks "github.com/nakji-network/connector/kafkautils/mocks"

	sq "github.com/Masterminds/squirrel"
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
	mockTokenCache := uniswapv2mocks.NewMockTokenCacheInterface(mockCtrl)

	db, err := database.New("host=localhost port=5432 user=postgres password=password database=blep_test")
	if err != nil {
		t.Fatal("Timescaledb connection failed")
	}

	// Cleanup pair table in the DB
	var count int
	sql, args, err := sq.Delete("").From("pair").ToSql()
	if err != nil {
		t.Fatal("Failed to delete pairs")
	}
	db.QueryRow(context.Background(), sql, args...).Scan(&count)
	log.Info().Int("count", count).Msg("pairs deleted")

	defer db.Close()

	kafkautils.TopicTypeRegistry.Load(TopicTypes)

	topics := map[string]kafkautils.Topic{
		"factoryfeeamountenabled":                kafkautils.MustParseTopic(".fct.nakji.uniswapv3.0_0_0.factory_feeamountenabled", "test"),
		"factoryownerchanged":                    kafkautils.MustParseTopic(".fct.nakji.uniswapv3.0_0_0.factory_ownerchanged", "test"),
		"factorypoolcreated":                     kafkautils.MustParseTopic(".fct.nakji.uniswapv3.0_0_0.factory_poolcreated", "test"),
		"poolburn":                               kafkautils.MustParseTopic(".fct.nakji.uniswapv3.0_0_0.pool_burn", "test"),
		"poolcollect":                            kafkautils.MustParseTopic(".fct.nakji.uniswapv3.0_0_0.pool_collect", "test"),
		"poolcollectprotocol":                    kafkautils.MustParseTopic(".fct.nakji.uniswapv3.0_0_0.pool_collectprotocol", "test"),
		"poolflash":                              kafkautils.MustParseTopic(".fct.nakji.uniswapv3.0_0_0.pool_flash", "test"),
		"poolincreaseobservationcardinalitynext": kafkautils.MustParseTopic(".fct.nakji.uniswapv3.0_0_0.pool_increaseobservationcardinalitynext", "test"),
		"poolinitialize":                         kafkautils.MustParseTopic(".fct.nakji.uniswapv3.0_0_0.pool_initialize", "test"),
		"poolmint":                               kafkautils.MustParseTopic(".fct.nakji.uniswapv3.0_0_0.pool_mint", "test"),
		"poolsetfeeprotocol":                     kafkautils.MustParseTopic(".fct.nakji.uniswapv3.0_0_0.pool_setfeeprotocol", "test"),
		"poolswap":                               kafkautils.MustParseTopic(".fct.nakji.uniswapv3.0_0_0.pool_swap", "test"),
	}

	factoryAddress := ethcommon.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")

	addresses := map[string][]ethcommon.Address{
		"factory": {factoryAddress},
		"pool":    {ethcommon.Address{}},
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

func (c *Uniswapv3Connector) setupMocks(cancel context.CancelFunc) {
	mockETHClientPool := c.ClientPool.(*ethmocks.MockETHClientPool)
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

	poolQuery := geth.FilterQuery{
		Addresses: c.addresses["pool"],
	}
	mockETHClientPool.EXPECT().SubscribeFilterLogs(gomock.Any(), poolQuery, gomock.Any()).DoAndReturn(
		func(ctx context.Context, query geth.FilterQuery, eventLogs chan<- ethtypes.Log) (event.Subscription, error) {
			poolAbi, err := abi.JSON(strings.NewReader(pool.PoolABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Pool ABI")
			}

			go tests.WriteEventLogs(eventLogs, &poolAbi)
			return mockSubscription, nil
		},
	)

	mockETHClientPool.EXPECT().ChunkedSubscribeFilterLogs(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, query geth.FilterQuery, eventLogs chan<- ethtypes.Log, chunksize int) (<-chan error, error) {
			factoryAbi, err := abi.JSON(strings.NewReader(factory.FactoryABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Factory ABI")
			}

			go tests.WriteEventLogs(eventLogs, &factoryAbi)

			errchan := make(chan error, 1)
			return errchan, nil
		},
	)

	mockETHClientPool.EXPECT().ChunkedSubscribeFilterLogs(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, query geth.FilterQuery, eventLogs chan<- ethtypes.Log, chunksize int) (<-chan error, error) {
			poolAbi, err := abi.JSON(strings.NewReader(pool.PoolABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Pool ABI")
			}

			go tests.WriteEventLogs(eventLogs, &poolAbi)

			errchan := make(chan error, 1)
			return errchan, nil
		},
	)

	mockETHClientPool.EXPECT().FilterLogs(context.Background(), gomock.Any()).AnyTimes().Return([]ethtypes.Log{}, nil)
	mockProducer.EXPECT().EnableTransactions()
	counter := 0
	mockProducer.EXPECT().WriteAndCommitSink(gomock.Any()).DoAndReturn(
		func(sink <-chan *kafkautils.Message) {
			go func() {
				for range sink {
					counter++

					if counter >= 12 {
						cancel()
					}
				}
			}()
		},
	)
}
