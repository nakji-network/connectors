// TODO: upon start, backfill x amount
package daofi

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"blep.ai/data/chain/ethereum/ethclient"
	"blep.ai/data/common"
	"blep.ai/data/connectors/source/daofi/factory"
	. "blep.ai/data/connectors/source/daofi/pair"
	"blep.ai/data/connectors/source/uniswapv2"
	"blep.ai/data/database"
	"blep.ai/data/tokencache"
	"github.com/nakji-network/connector/kafkautils"

	geth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Config struct {
	KP               *kafkautils.Producer
	RPCURLs          []string
	LPTopic          kafkautils.Topic
	LPEventTopic     kafkautils.Topic
	TradeTopic       kafkautils.Topic
	Db               *database.Database
	FactoryAddress   ethcommon.Address
	Namespace        string //uniswapv2, honeyswap, etc
	TokenNamespace   string //ethereum, xdai, etc
	EnableBlockCache bool   // load all historical blocks into memory. useful for backfill only I think
}

type UniswapV2 struct {
	Config
	*ethclient.ClientPool
	*Status

	pairCache  *PairCache
	tokenCache *tokencache.TokenCache
	blockCache map[uint64]uint64

	backfillNumBlocks uint64
}

// For kubernetes healthcheck
type Status struct {
	Live  bool
	Ready bool
	Data  map[string]interface{}
}

// Connect to eth rpc and initialize pair+token cache
func New(config Config) *UniswapV2 {
	if config.Namespace == "" || config.TokenNamespace == "" {
		log.Fatal().Msg("config.Namespace and TokenNamespace must be set.")
	}

	log.Info().Strs("url", config.RPCURLs).Msg("connecting to Ethereum RPC")
	ethClientPool, err := ethclient.DialPoolContext(context.Background(), config.RPCURLs)
	if err != nil {
		log.Fatal().Err(err).Msg("Ethereum RPC connection error")
	}

	client, err := ethClientPool.RandClient(false)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to find available client")
	}

	ethconn := client.Client

	tokenCache, err := tokencache.NewTokenCache(config.TokenNamespace, config.Db, ethClientPool, 18000)
	if err != nil {
		log.Fatal().Err(err).Msg("token cache failed to create")
	}

	pairCache, err := NewPairCache(config.TokenNamespace, config.Namespace, config.Db, ethconn, 1000, tokenCache)
	if err != nil {
		log.Fatal().Err(err).Msg("Pairs failed to initialize")
	}

	// Hacky way to load all blocks into memory to reduce db hits
	var blockCache map[uint64]uint64
	if config.EnableBlockCache {
		blockCache, err = config.Db.GetAllBlockTimes()
		if err != nil {
			log.Fatal().Err(err).Msg("block cache failed to initialize")
		}
	}

	return &UniswapV2{
		Config:            config,
		pairCache:         pairCache,
		tokenCache:        tokenCache,
		ClientPool:        ethClientPool,
		blockCache:        blockCache,
		Status:            new(Status),
		backfillNumBlocks: 100,
	}
}

// Listeners

// Start all listeners
func (c *UniswapV2) Start() {
	c.Status.Live = true
	c.InsertPairAddresses()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	defer c.ClientPool.Close()

	// Keepalive by subscribing to new heads
	headers := make(chan *ethtypes.Header)
	defer close(headers)

	client, err := c.ClientPool.RandClient(true)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to find available client")
	}

	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal().Err(err).Msg("SubscribeNewHead error")
	}

	// All Listeners
	logc1, errc1 := c.FactoryListener()
	logc2, errc2, err := c.PairListener()
	if err != nil {
		log.Fatal().Err(err).
			//Interface("query", query).
			Msg("Subscribe pairs failed to subscribe")
	}

	out := common.MergeErrChans(errc1, errc2, sub.Err())

	c.KP.EnableTransactions()
	// We want to set this to Ready at least after Websocket listeners have been setup so
	// that new rollouts will only terminate the old pod once the new pod has started listening
	// for events. This way, we minimize missing any events during deployments.
	c.Status.Ready = true

	// Initialize uniswap pair ABI for reading logs
	contractAbi, err := abi.JSON(strings.NewReader(IDAOfiV1PairABI))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read uniswapv2 pair abi")
	}

	sink := make(chan *kafkautils.Message) // 0 length for real time streams

	// out chan written to kafka
	go c.KP.WriteAndCommitSink(sink)

	var once sync.Once

	for {
		select {
		case err = <-out:
			c.Status.Live = false
			c.Status.Ready = false
			log.Error().Err(err).Msg("sub error")
			return
		case <-interrupt:
			log.Debug().Msg("interrupt")
			return
		case header := <-headers:
			log.Debug().
				Str("block", header.Number.String()).
				Uint64("ts", header.Time).
				Msg("header received")
		case vLog := <-logc1: // All newly added pairs
			if vLog.Removed {
				continue
			}

			client, err := c.ClientPool.RandClient(true)
			if err != nil {
				log.Warn().Err(err).Msg("failed to find available client")
			} else {
				c.processPairLogTs(client, contractAbi, vLog, sink)
			}
		case vLog := <-logc2: // All initial pairs
			if vLog.Removed {
				continue
			}

			// start backfill from the first log received. Logs come in later than headers if the query is large
			go once.Do(func() {
				c.backfill(vLog.BlockNumber, c.backfillNumBlocks)
			})

			client, err := c.ClientPool.RandClient(true)
			if err != nil {
				log.Warn().Err(err).Msg("failed to find available client")
			} else {
				c.processPairLogTs(client, contractAbi, vLog, sink)
			}
		}
	}
}

// Get timestamp of the log's block and then process log for output.
// Used only in Start().Prevent clogging of vLogs by using concurrency.
func (c *UniswapV2) processPairLogTs(client *ethclient.Client, contractAbi abi.ABI, vLog ethtypes.Log, sink chan *kafkautils.Message) {
	// Add timestamp from block since logs don't include timestamp
	ts, err := ethclient.GetLogTimestamp(client.Client, vLog, nil)
	if err != nil {
		log.Error().Err(err).
			Interface("blockNumber", vLog.BlockNumber).
			Msg("GetLogTimetsamp error")
	}

	// Writes to out chan
	err = c.ProcessPairLog(contractAbi, vLog, ts, sink, false)
	if err != nil {
		log.Error().Err(err).
			Interface("vLog", vLog).
			Msg("ProcessPairLog error")
	}

	// Even if there are issues, it should not clog anything
}

// Backfill last 100 blocks
func (c *UniswapV2) backfill(latestBlockNumber, backfillNumBlocks uint64) {
	messageCh := c.PairHistoryPipeline(geth.FilterQuery{
		FromBlock: big.NewInt(int64(latestBlockNumber - backfillNumBlocks)),
		ToBlock:   big.NewInt(int64(latestBlockNumber)),
	}, 100, 1)
	c.Db.InsertBatchKafkaMessagesBuffered(messageCh, 10000)
}

// Subscribe to UniswapV2Factory contract to listen to PairCreated events.
// Subscribes to logs for all new pairs created.
func (u *UniswapV2) FactoryListener() (<-chan ethtypes.Log, <-chan error) {
	ch := make(chan *factory.IDAOfiV1FactoryPairCreated)
	var start *uint64 // keep track of starting block number to prevent gaps during resubscribes
	// however I don't think it's supported on most RPCs

	sub := event.Resubscribe(100*time.Millisecond, func(ctx context.Context) (event.Subscription, error) {
		// Connect to rpc
		client, err := u.ClientPool.RandClient(true)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to find available client")
		}

		filterer, err := factory.NewIDAOfiV1FactoryFilterer(u.Config.FactoryAddress, client)
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}

		// Watch pair created from `start`
		opts := &bind.WatchOpts{Start: start, Context: ctx}
		sub, err := filterer.WatchPairCreated(opts, ch, []ethcommon.Address{}, []ethcommon.Address{})
		if err != nil {
			log.Error().Err(err).
				//Uint64("start", *start).
				Msg("WatchPairCreated failed to subscribe")
		}

		log.Info().Msg("FactoryListener live")
		return sub, err
	})

	logc := make(chan ethtypes.Log)
	errc := make(chan error, 1)

	go func() {
		for {
			select {
			case err := <-sub.Err():
				errc <- err
				return
			case CreatedPair := <-ch:
				log.Info().
					Str("Index", CreatedPair.Length.String()).
					Str("token0", CreatedPair.BaseToken.Hex()).
					Str("token1", CreatedPair.QuoteToken.Hex()).
					Str("pair", CreatedPair.Pair.Hex()).
					Msg("new pair detected")

				u.InsertPairAddress(CreatedPair)
				// Add new pair to pairCache
				start = &CreatedPair.Raw.BlockNumber

				// Subscribe to new pair's logs
				query := geth.FilterQuery{
					Addresses: []ethcommon.Address{CreatedPair.Pair},
				}

				sub, err := u.ClientPool.SubscribeFilterLogs(context.Background(), query, logc)
				if err != nil {
					log.Fatal().Err(err).
						Interface("query", query).
						Msg("Subscribe pairs failed to subscribe")
				}
				errc = common.MergeErrChans(errc, sub.Err())
			}
		}
	}()

	return logc, errc
}

// Subscribe to all pairs
func (u *UniswapV2) PairListener() (<-chan ethtypes.Log, <-chan error, error) {
	// Subscribe to Uniswap Pair logs on ethereum
	pairs, err := u.pairCache.AllPairAddresses()
	if err != nil {
		log.Fatal().Err(err).Msg("Could not start PairListener")
	}

	query := geth.FilterQuery{
		Addresses: pairs,
	}

	client, err := u.ClientPool.RandClient(true)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to find available client")
	}
	return client.ChunkedSubscribeFilterLogs(query, 0)
}

// Historical
// Get all pairs history. Similar to PairListener() but grabs historical data
// from chain and writes to chan instead of kafka.
// filterQuery gets addresses based on current AllPairAddresses.
func (u *UniswapV2) PairHistoryPipeline(filterQuery geth.FilterQuery, blockChunkSize, maxConnections int64) <-chan *kafkautils.Message {
	// Get all Uniswap Pair addresses for the query
	pairs, err := u.pairCache.AllPairAddresses()
	if err != nil {
		log.Fatal().Err(err).Msg("Could not start PairListener")
	}
	filterQuery.Addresses = pairs

	// Pipeline
	client, err := u.ClientPool.RandClient(false)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to find available client")
	}
	logc, _ := client.ChunkedFilterLogs(filterQuery, 0, blockChunkSize, maxConnections)
	return u.ProcessPairLogs(logc, true)
}

// Processors

// ProcessPairLogs pineline
func (u *UniswapV2) ProcessPairLogs(logc <-chan ethtypes.Log, isBackfillJob bool) <-chan *kafkautils.Message {
	out := make(chan *kafkautils.Message, 25000)

	// Initialize uniswap pair ABI for reading logs
	contractAbi, err := abi.JSON(strings.NewReader(IDAOfiV1PairABI))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read uniswapv2 pair abi")
	}

	var wg sync.WaitGroup
	go func() {
		for vLog := range logc {
			wg.Add(1)
			go func(vLog ethtypes.Log) {
				defer wg.Done()
				if !vLog.Removed {
					// Add timestamp from block since logs don't include timestamp
					client, err := u.ClientPool.RandClient(false)
					if err != nil {
						log.Warn().Err(err).Msg("failed to find available client")
					}

					ts, err := ethclient.GetLogTimestamp(client.Client, vLog, u.blockCache)
					if err != nil {
						log.Error().Err(err).
							Interface("blockNumber", vLog.BlockNumber).
							Msg("GetLogTimetsamp error")
					}

					// Writes to out chan
					err = u.ProcessPairLog(contractAbi, vLog, ts, out, isBackfillJob)
					if err != nil {
						log.Error().Err(err).
							Interface("vLog", vLog).
							Msg("ProcessPairLog error")
						return
					}
				}
			}(vLog)
		}
		wg.Wait()
		close(out)
	}()

	return out
}

// Parse the event and generate LiquidityPool, Trade, and LPEvent kafka messages
func (u *UniswapV2) ProcessPairLog(contractAbi abi.ABI, vLog ethtypes.Log, ts uint64, out chan<- *kafkautils.Message, isBackfillJob bool) error {
	ev, err := contractAbi.EventByID(vLog.Topics[0])
	if err != nil {
		log.Warn().Err(err).Msg("EventByID error, skipping")
		return err
	}

	if ev == nil {
		return nil // fmt.Errorf("ignore if event id isn't defined in a partial ABI")
	}

	switch ev.Name {
	// state-changing functions
	case "Swap": //, "Mint", "Burn", "Skim", "Sync":
		event := new(IDAOfiV1PairSwap)
		if err := ethclient.UnpackLog(contractAbi, event, "Swap", vLog); err != nil {
			log.Error().Err(err).Msg("Unpack swap event error")
			return err
		}

		lpe := uniswapv2.LiquidityPoolChange{
			Ts:     common.UnixToTimestampPb(int64(ts * 1000)),
			Block:  vLog.BlockNumber,
			Idx:    uint64(vLog.Index),
			Tx:     vLog.TxHash.Bytes(),
			Sender: event.Sender.Bytes(),
			To:     event.To.Bytes(),
		}
		var isBuy bool
		var isSell bool

		// translate from daofi to uniswap style
		pair, ok := u.pairCache.Get(vLog.Address)
		if !ok {
			return fmt.Errorf("could not get daofi pair")
		}
		if bytes.Compare(event.TokenIn.Bytes(), pair.Token0) == 0 {
			// equal, so base token is the token sold
			lpe.Amount0In = event.AmountIn.Bytes()
			lpe.Amount1Out = event.AmountOut.Bytes()
			isSell = true
		} else {
			lpe.Amount1In = event.AmountIn.Bytes()
			lpe.Amount0Out = event.AmountOut.Bytes()
			isBuy = true
		}

		out <- &kafkautils.Message{
			Topic:    u.LPEventTopic,
			Key:      kafkautils.NewKey(u.Config.Namespace, vLog.Address.Hex()),
			ProtoMsg: &lpe,
		}

		// Backfills don't need trade and lp since they can be derived from LPEventTopic
		if isBackfillJob {
			return nil
		}

		// Construct Liquidity Pool message
		lp, err := u.LP(vLog.Address)
		if err != nil {
			log.Error().
				Err(err).
				Str("pair", vLog.Address.Hex()).
				Msg("Skipping this pair")
		} else {
			lpProto := lp.Proto()
			lpProto.Ts = common.UnixToTimestampPb(int64(ts * 1000))
			out <- &kafkautils.Message{
				Topic:    u.LPTopic,
				Key:      kafkautils.NewKey(u.Config.Namespace, vLog.Address.Hex()),
				ProtoMsg: &lpProto,
			}

			// Construct Trade message
			in0d, _ := setDecimals(big.NewInt(0).SetBytes(lpe.Amount0In), lp.token0.Decimals).Float64()
			out0d, _ := setDecimals(big.NewInt(0).SetBytes(lpe.Amount0Out), lp.token0.Decimals).Float64()
			in1d, _ := setDecimals(big.NewInt(0).SetBytes(lpe.Amount1In), lp.token1.Decimals).Float64()
			out1d, _ := setDecimals(big.NewInt(0).SetBytes(lpe.Amount1Out), lp.token1.Decimals).Float64()

			tradeProto := &common.Trade{
				Ts: common.UnixToTimestampPb(int64(ts * 1000)),
				//Price
				Amount0In:  in0d,  // sum to get sell volume
				Amount0Out: out0d, // sum to get buy volume
				Amount1In:  in1d,
				Amount1Out: out1d,
				//Side
				EventType: common.Trade_TYPE_MARKET, // all uniswap trades are market orders
			}

			switch {
			case isBuy:
				// 0out, 1in (eg eth <- usdt)
				tradeProto.Price = tradeProto.Amount1In / tradeProto.Amount0Out
				tradeProto.Side = common.Trade_SIDE_BUY
			case isSell:
				// 0in, 1out (eg eth -> usdt)
				tradeProto.Price = tradeProto.Amount1Out / tradeProto.Amount0In
				tradeProto.Side = common.Trade_SIDE_SELL
			default:
				log.Error().
					Interface("tradeProto", tradeProto).
					Msg("Amounts are inconsistent: unhandled")
			}

			if isBuy || isSell {
				out <- &kafkautils.Message{
					Topic:    u.TradeTopic,
					Key:      kafkautils.NewKey(u.Config.Namespace, vLog.Address.Hex()),
					ProtoMsg: tradeProto,
				}
			}
		}

		return nil
	default:
		return nil
	}
}

// Helpers

// returns LP{token0, token1}. No contract calls usually
func (u *UniswapV2) getTokens(pairAddress ethcommon.Address) (*LP, error) {
	pair, ok := u.pairCache.Get(pairAddress)
	if !ok {
		return nil, fmt.Errorf("pair address %s not found in PairMap", pairAddress.Hex())
	}
	token0, ok := u.tokenCache.Get(ethcommon.BytesToAddress(pair.Token0))
	if !ok {
		return nil, fmt.Errorf("could not get token0 %s", ethcommon.BytesToAddress(pair.Token0).Hex())
	}
	token1, ok := u.tokenCache.Get(ethcommon.BytesToAddress(pair.Token1))
	if !ok {
		return nil, fmt.Errorf("could not get token1 %s", ethcommon.BytesToAddress(pair.Token1).Hex())
	}

	return &LP{token0: token0, token1: token1}, nil
}

func (u *UniswapV2) LP(pairAddress ethcommon.Address) (*LP, error) {
	pair, ok := u.pairCache.Get(pairAddress)
	if !ok {
		return nil, fmt.Errorf("pair address %s not found in PairMap", pairAddress.Hex())
	}
	token0, ok := u.tokenCache.Get(ethcommon.BytesToAddress(pair.Token0))
	if !ok {
		return nil, fmt.Errorf("could not get token0 %s", ethcommon.BytesToAddress(pair.Token0).Hex())
	}
	token1, ok := u.tokenCache.Get(ethcommon.BytesToAddress(pair.Token1))
	if !ok {
		return nil, fmt.Errorf("could not get token1 %s", ethcommon.BytesToAddress(pair.Token1).Hex())
	}

	// Initialize contract instance to call contract functions
	client, err := u.ClientPool.RandClient(true)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to find available client")
	}

	contract, err := NewIDAOfiV1Pair(pairAddress, client)
	if err != nil {
		log.Fatal().Err(err)
	}

	reserves, err := contract.GetReserves(nil)
	if err != nil {
		log.Fatal().Err(err)
	}

	return &LP{Reserves{
		Reserve0: reserves.ReserveBase,
		Reserve1: reserves.ReserveQuote,
	}, token0, token1}, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func setDecimals(input *big.Int, decimals uint32) *big.Rat {
	dec := big.NewInt(int64(decimals))
	exp := new(big.Int).Exp(big.NewInt(10), dec, nil)

	return new(big.Rat).SetFrac(input, exp)
}

// Block with Daofi's Factory contract creation
var factoryContractBlock = big.NewInt(12105674)

// eth_filter and eth_call have a limit of 10000 blocks range
const batchSize = 10000

func (c *UniswapV2) InsertPairAddresses() {
	lastBlock := c.Db.GetLastBlockQueried(c.TokenNamespace, c.Namespace, factoryContractBlock)
	client, _ := c.ClientPool.RandClient(true)
	latestBlock, _ := client.BlockNumber(context.Background())

	fromBlock := lastBlock.Uint64()
	toBlock := lastBlock.Uint64() + batchSize

	// We limit the filter call here to a range of 10000 blocks so we do not hit limits instead
	// of doing 1 big filter call.
	for fromBlock < latestBlock {
		filterer, err := factory.NewIDAOfiV1FactoryFilterer(c.FactoryAddress, client)
		if err != nil {
			log.Fatal().Err(err).Msg("FactoryFilterer failed")
		}

		opts := &bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: context.Background()}
		log.Info().Uint64("fromBlock", fromBlock).Uint64("toBlock", toBlock).Msg("Filtering PairCreated Logs from FactoryAddress")
		pairCreatedIterator, err := filterer.FilterPairCreated(opts, []ethcommon.Address{}, []ethcommon.Address{})
		if err != nil {
			log.Fatal().Err(err).Msg("FilterPairCreated failed")
		}

		if err != nil {
			log.Error().Err(err).Msg("Failed filtering for PairCreated events")
		}

		for pairCreatedIterator.Next() {
			event := pairCreatedIterator.Event

			log.Info().
				Str("token0", event.BaseToken.Hex()).
				Str("token1", event.QuoteToken.Hex()).
				Msg("Inserting pair from PairCreated event")

			c.InsertPairAddress(event)
		}

		c.Db.UpdatelastBlock(c.TokenNamespace, c.Namespace, toBlock)
		fromBlock = toBlock
		toBlock += batchSize
	}
}

func (c *UniswapV2) InsertPairAddress(event *factory.IDAOfiV1FactoryPairCreated) {
	retries := 0

	for retries <= 2 {
		t0, _ := c.tokenCache.Get(event.BaseToken)
		t1, _ := c.tokenCache.Get(event.QuoteToken)

		if t0 != nil && t1 != nil {
			pair := &common.Pair{
				Ns:     c.TokenNamespace,
				App:    c.Namespace,
				Id:     event.Pair.Bytes(),
				Added:  timestamppb.Now(),
				I:      int64(event.Fee),
				Token0: event.BaseToken.Bytes(),
				Token1: event.QuoteToken.Bytes(),
				Idh:    event.Pair.Hex(),
				D0:     t0.Decimals,
				N0:     t0.Name,
				S0:     t0.Symbol,
				D1:     t1.Decimals,
				N1:     t1.Name,
				S1:     t1.Symbol,
			}

			c.pairCache.Add(event.Pair, pair)
			break
		}

		time.Sleep(time.Duration(1) * time.Second)
		retries++
	}
}
