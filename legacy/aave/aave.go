// TODO: upon start, backfill x amount
package aave

import (
	"context"
	"strings"
	"sync"

	"blep.ai/data/common"
	"github.com/ethereum/go-ethereum"

	"os"
	"os/signal"

	"blep.ai/data/chain/ethereum/ethclient"
	. "blep.ai/data/connectors/source/aave/lendingpool"
	"blep.ai/data/database"
	"github.com/nakji-network/connector/kafkautils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

type Config struct {
	KP             *kafkautils.Producer
	Topics         map[string]kafkautils.Topic
	RPCURLs        []string
	Db             *database.Database
	FactoryAddress ethcommon.Address
	Namespace      string //uniswapv2, honeyswap, etc
	//EnableBlockCache bool // load all historical blocks into memory. useful for backfill only I think
}

type Aave struct {
	Config
	*ethclient.ClientPool
	*Status

	backfillNumBlocks uint64
}

// For kubernetes healthcheck
type Status struct {
	Live  bool
	Ready bool
	Data  map[string]interface{}
}

// Connect to eth rpc and initialize pair+token cache
func New(config Config) *Aave {
	if config.Namespace == "" {
		log.Fatal().Msg("config.Namespace must be set.")
	}

	log.Info().Strs("url", config.RPCURLs).Msg("connecting to Ethereum RPC")
	ethClientPool, err := ethclient.DialPoolContext(context.Background(), config.RPCURLs)
	if err != nil {
		log.Fatal().Err(err).Msg("Ethereum RPC connection error")
	}

	if err != nil {
		log.Fatal().Err(err).Msg("token cache failed to create")
	}

	return &Aave{
		Config:            config,
		ClientPool:        ethClientPool,
		Status:            new(Status),
		backfillNumBlocks: 100,
	}
}

//
//
// Listeners

// Start all listeners
func (c *Aave) Start() {
	c.Status.Live = true

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	defer c.ClientPool.Close()

	// Keepalive by subscribing to new heads
	headers := make(chan *ethtypes.Header)
	defer close(headers)
	sub, err := c.ClientPool.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal().Err(err).Msg("SubscribeNewHead error")
	}

	// All Listeners
	log1, errc1 := c.subscribe()

	c.KP.EnableTransactions()
	c.Status.Ready = true

	// Initialize lending pool ABI for reading logs
	contractAbi, err := abi.JSON(strings.NewReader(AaveLendingPoolABI))

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read aave lending pool abi")
	}

	sink := make(chan *kafkautils.Message) // 0 length for real time streams

	// out chan written to kafka
	go c.KP.WriteAndCommitSink(sink)

	for {
		select {
		case err := <-mergeErrChans(errc1, sub.Err()):
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

			ethclient.CacheBlockTimestamp(header.Hash(), header.Time)
		case vLog := <-log1:
			ts, err := c.ClientPool.GetLogTimestamp(vLog, nil)
			if err != nil {
				log.Error().Err(err).
					Interface("blockNumber", vLog.BlockNumber).
					Msg("GetLogTimetsamp error")
			}
			c.processLogs(vLog, sink, ts, contractAbi)

		}
	}
}

func (c *Aave) subscribe() (<-chan ethtypes.Log, <-chan error) {
	query := ethereum.FilterQuery{
		Addresses: []ethcommon.Address{c.Config.FactoryAddress},
	}
	logc := make(chan ethtypes.Log)

	sub, err := c.ClientPool.SubscribeFilterLogs(context.Background(), query, logc)
	if err != nil {
		log.Error().Err(err).Msg("Error Subscribing")
	}

	return logc, sub.Err()
}

func (c *Aave) processLogs(vLog ethtypes.Log, out chan<- *kafkautils.Message, ts uint64, contractAbi abi.ABI) {
	ev, err := contractAbi.EventByID(vLog.Topics[0])
	if err != nil {
		log.Warn().Err(err).Msg("EventByID error, skipping")
		return
	}
	log.Info().Str("event", ev.Name).Msg("Debug")
	switch ev.Name {
	case "Borrow":
		event := new(AaveLendingPoolBorrow)
		if err := ethclient.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack borrow event error")
			return
		}
		out <- &kafkautils.Message{
			Topic: c.Topics["Borrow"],
			Key:   kafkautils.NewKey(c.Config.Namespace, vLog.Address.Hex()),
			ProtoMsg: &LendingPoolBorrow{
				Ts:         common.UnixToTimestampPb(int64(ts * 1000)),
				Block:      vLog.BlockNumber,
				Idx:        uint64(vLog.Index),
				Tx:         vLog.TxHash.Bytes(),
				Reserve:    event.Reserve.Bytes(),
				User:       event.User.Bytes(),
				BehalfOf:   event.OnBehalfOf.Bytes(),
				Amount:     event.Amount.Bytes(),
				BorrowRate: event.BorrowRate.Bytes(),
				RateMode:   event.BorrowRateMode.Bytes(),
				Referral:   uint64(event.Referral),
			},
		}
	case "Deposit":
		event := new(AaveLendingPoolDeposit)
		if err := ethclient.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack deposit event error")
			return
		}
		out <- &kafkautils.Message{
			Topic: c.Topics["Deposit"],
			Key:   kafkautils.NewKey(c.Config.Namespace, vLog.Address.Hex()),
			ProtoMsg: &LendingPoolDeposit{
				Ts:       common.UnixToTimestampPb(int64(ts * 1000)),
				Block:    vLog.BlockNumber,
				Idx:      uint64(vLog.Index),
				Tx:       vLog.TxHash.Bytes(),
				Reserve:  event.Reserve.Bytes(),
				User:     event.User.Bytes(),
				BehalfOf: event.OnBehalfOf.Bytes(),
				Amount:   event.Amount.Bytes(),
				Referral: uint64(event.Referral),
			},
		}
	case "FlashLoan":
		event := new(AaveLendingPoolFlashLoan)
		if err := ethclient.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack flash loan event error")
			return
		}
		out <- &kafkautils.Message{
			Topic: c.Topics["FlashLoan"],
			Key:   kafkautils.NewKey(c.Config.Namespace, vLog.Address.Hex()),
			ProtoMsg: &LendingPoolFlashLoan{
				Ts:        common.UnixToTimestampPb(int64(ts * 1000)),
				Block:     vLog.BlockNumber,
				Idx:       uint64(vLog.Index),
				Tx:        vLog.TxHash.Bytes(),
				Target:    event.Target.Bytes(),
				Initiator: event.Initiator.Bytes(),
				Asset:     event.Asset.Bytes(),
				Amount:    event.Amount.Bytes(),
				Premium:   event.Premium.Bytes(),
				Referral:  uint64(event.ReferralCode),
			},
		}
	case "LiquidationCall":
		event := new(AaveLendingPoolLiquidationCall)
		if err := ethclient.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack liquidation event error")
			return
		}
		out <- &kafkautils.Message{
			Topic: c.Topics["LiquidationCall"],
			Key:   kafkautils.NewKey(c.Config.Namespace, vLog.Address.Hex()),
			ProtoMsg: &LendingPoolLiquidationCall{
				Ts:               common.UnixToTimestampPb(int64(ts * 1000)),
				Block:            vLog.BlockNumber,
				Idx:              uint64(vLog.Index),
				Tx:               vLog.TxHash.Bytes(),
				CollateralAsset:  event.CollateralAsset.Bytes(),
				DebtAsset:        event.DebtAsset.Bytes(),
				User:             event.User.Bytes(),
				DebtToCover:      event.DebtToCover.Bytes(),
				CollateralAmount: event.LiquidatedCollateralAmount.Bytes(),
				Liquidator:       event.Liquidator.Bytes(),
				ReceiveAToken:    event.ReceiveAToken,
			},
		}
	case "RebalanceStableBorrowRate":
		event := new(AaveLendingPoolRebalanceStableBorrowRate)
		if err := ethclient.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack rebalance event error")
			return
		}
		out <- &kafkautils.Message{
			Topic: c.Topics["RebalanceStableBorrowRate"],
			Key:   kafkautils.NewKey(c.Config.Namespace, vLog.Address.Hex()),
			ProtoMsg: &LendingPoolRebalance{
				Ts:      common.UnixToTimestampPb(int64(ts * 1000)),
				Block:   vLog.BlockNumber,
				Idx:     uint64(vLog.Index),
				Tx:      vLog.TxHash.Bytes(),
				Reserve: event.Reserve.Bytes(),
				User:    event.User.Bytes(),
			},
		}
	case "Repay":
		event := new(AaveLendingPoolRepay)
		if err := ethclient.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack repay event error")
			return
		}
		out <- &kafkautils.Message{
			Topic: c.Topics["Repay"],
			Key:   kafkautils.NewKey(c.Config.Namespace, vLog.Address.Hex()),
			ProtoMsg: &LendingPoolRepay{
				Ts:      common.UnixToTimestampPb(int64(ts * 1000)),
				Block:   vLog.BlockNumber,
				Idx:     uint64(vLog.Index),
				Tx:      vLog.TxHash.Bytes(),
				Reserve: event.Reserve.Bytes(),
				User:    event.User.Bytes(),
				Repayer: event.Repayer.Bytes(),
				Amount:  event.Amount.Bytes(),
			},
		}
	case "ReserveDataUpdated":
		event := new(AaveLendingPoolReserveDataUpdated)
		if err := ethclient.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack reserve data updated event error")
			return
		}
		out <- &kafkautils.Message{
			Topic: c.Topics["ReserveDataUpdated"],
			Key:   kafkautils.NewKey(c.Config.Namespace, vLog.Address.Hex()),
			ProtoMsg: &LendingPoolReserveUpdated{
				Ts:               common.UnixToTimestampPb(int64(ts * 1000)),
				Block:            vLog.BlockNumber,
				Idx:              uint64(vLog.Index),
				Tx:               vLog.TxHash.Bytes(),
				LiquidityRate:    event.LiquidityRate.Bytes(),
				LiquidityIndex:   event.LiquidityIndex.Bytes(),
				StableBorrowRate: event.StableBorrowRate.Bytes(),
				VarBorrowIndex:   event.VariableBorrowIndex.Bytes(),
				VarBorrowRate:    event.VariableBorrowRate.Bytes(),
			},
		}
	case "ReserveUsedAsCollateralDisabled":
		event := new(AaveLendingPoolReserveUsedAsCollateralDisabled)
		if err := ethclient.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack reserve used as collateral disabled event error")
			return
		}
		out <- &kafkautils.Message{
			Topic: c.Topics["ReserveUsedAsCollateral"],
			Key:   kafkautils.NewKey(c.Config.Namespace, vLog.Address.Hex()),
			ProtoMsg: &LendingPoolReserveAsCollateral{
				Ts:         common.UnixToTimestampPb(int64(ts * 1000)),
				Block:      vLog.BlockNumber,
				Idx:        uint64(vLog.Index),
				Tx:         vLog.TxHash.Bytes(),
				Reserve:    event.Reserve.Bytes(),
				User:       event.User.Bytes(),
				Collateral: false,
			},
		}
	case "ReserveUsedAsCollateralEnabled":
		event := new(AaveLendingPoolReserveUsedAsCollateralEnabled)
		if err := ethclient.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack reserve used as collateral disabled event error")
			return
		}
		out <- &kafkautils.Message{
			Topic: c.Topics["ReserveUsedAsCollateral"],
			Key:   kafkautils.NewKey(c.Config.Namespace, vLog.Address.Hex()),
			ProtoMsg: &LendingPoolReserveAsCollateral{
				Ts:         common.UnixToTimestampPb(int64(ts * 1000)),
				Block:      vLog.BlockNumber,
				Idx:        uint64(vLog.Index),
				Tx:         vLog.TxHash.Bytes(),
				Reserve:    event.Reserve.Bytes(),
				User:       event.User.Bytes(),
				Collateral: true,
			},
		}
	case "Swap":
		event := new(AaveLendingPoolSwap)
		if err := ethclient.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack swap event error")
			return
		}
		out <- &kafkautils.Message{
			Topic: c.Topics["Swap"],
			Key:   kafkautils.NewKey(c.Config.Namespace, vLog.Address.Hex()),
			ProtoMsg: &LendingPoolSwap{
				Ts:       common.UnixToTimestampPb(int64(ts * 1000)),
				Block:    vLog.BlockNumber,
				Idx:      uint64(vLog.Index),
				Tx:       vLog.TxHash.Bytes(),
				Reserve:  event.Reserve.Bytes(),
				User:     event.User.Bytes(),
				RateMode: event.RateMode.Bytes(),
			},
		}
	case "Withdraw":
		event := new(AaveLendingPoolWithdraw)
		if err := ethclient.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack withdraw event error")
			return
		}
		out <- &kafkautils.Message{
			Topic: c.Topics["Withdraw"],
			Key:   kafkautils.NewKey(c.Config.Namespace, vLog.Address.Hex()),
			ProtoMsg: &LendingPoolWithdraw{
				Ts:      common.UnixToTimestampPb(int64(ts * 1000)),
				Block:   vLog.BlockNumber,
				Idx:     uint64(vLog.Index),
				Tx:      vLog.TxHash.Bytes(),
				Reserve: event.Reserve.Bytes(),
				User:    event.User.Bytes(),
				To:      event.To.Bytes(),
				Amount:  event.Amount.Bytes(),
			},
		}
	}
}

// Chan utils

func mergeErrChans(cs ...<-chan error) chan error {
	var wg sync.WaitGroup
	out := make(chan error)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan error) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
