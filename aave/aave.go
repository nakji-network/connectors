package aave

import (
	"context"
	"fmt"
	"strings"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/nakji-network/connector"
	"github.com/nakji-network/connector/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	. "github.com/nakji-network/connectors/aave/lendingpool"
)

type Config struct {
	ConnectorName string
	NetworkName   string
	FromBlock     uint64
	NumBlocks     uint64
	Namespace     string
}

type Connector struct {
	*connector.Connector
	*Config
	sub connector.ISubscription
}

func New(c *connector.Connector, config *Config) *Connector {
	return &Connector{
		Connector: c,
		Config:    config,
	}
}

func (c *Connector) Start() {
	addresses := []ethcommon.Address{ethcommon.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9")}

	ctx := context.Background()

	if sub, err := connector.NewSubscription(ctx, c.Connector, c.NetworkName, addresses, c.FromBlock, c.NumBlocks); err == nil {
		c.sub = sub
	} else {
		log.Fatal().Err(err).Msg(fmt.Sprintf("%s connection error", c.NetworkName))
	}

	// Initialize lending pool ABI for reading logs
	contractAbi, err := abi.JSON(strings.NewReader(AaveLendingPoolABI))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read aave lending pool abi")
	}

	for {
		select {
		case <-c.sub.Done():
			log.Info().Msg("connector shutdown")
			return

		//	Listen to error channel
		case err := <-c.sub.Err():
			log.Error().Err(err).Str("network", c.NetworkName).Msg("subscription failed")
			return

		//	Listen to event logs
		case vLog := <-c.sub.Logs():
			if msg := c.parse(vLog, contractAbi); msg != nil {
				log.Debug().Msg(protojson.Format(msg))
				c.EventSink <- msg
			}
		}
	}
}

func (c *Connector) parse(vLog types.Log, contractAbi abi.ABI) proto.Message {
	ev, err := contractAbi.EventByID(vLog.Topics[0])
	if err != nil {
		log.Warn().Err(err).Msg("EventByID error, skipping")
		return nil
	}
	t, err := c.sub.GetBlockTime(vLog)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve timestamp")
	}
	ts := common.UnixToTimestampPb(int64(t * 1000))
	switch ev.Name {
	case "Borrow":
		event := new(AaveLendingPoolBorrow)
		if err := common.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack borrow event error")
			return nil
		}
		return &LendingPoolBorrow{
			Ts:         ts,
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
			Ns:         c.Config.Namespace,
			S:          vLog.Address.Hex(),
		}
	case "Deposit":
		event := new(AaveLendingPoolDeposit)
		if err := common.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack deposit event error")
			return nil
		}
		return &LendingPoolDeposit{
			Ts:       ts,
			Block:    vLog.BlockNumber,
			Idx:      uint64(vLog.Index),
			Tx:       vLog.TxHash.Bytes(),
			Reserve:  event.Reserve.Bytes(),
			User:     event.User.Bytes(),
			BehalfOf: event.OnBehalfOf.Bytes(),
			Amount:   event.Amount.Bytes(),
			Referral: uint64(event.Referral),
			Ns:       c.Config.Namespace,
			S:        vLog.Address.Hex(),
		}
	case "FlashLoan":
		event := new(AaveLendingPoolFlashLoan)
		if err := common.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack flash loan event error")
			return nil
		}
		return &LendingPoolFlashLoan{
			Ts:        ts,
			Block:     vLog.BlockNumber,
			Idx:       uint64(vLog.Index),
			Tx:        vLog.TxHash.Bytes(),
			Target:    event.Target.Bytes(),
			Initiator: event.Initiator.Bytes(),
			Asset:     event.Asset.Bytes(),
			Amount:    event.Amount.Bytes(),
			Premium:   event.Premium.Bytes(),
			Referral:  uint64(event.ReferralCode),
			Ns:        c.Config.Namespace,
			S:         vLog.Address.Hex(),
		}
	case "LiquidationCall":
		event := new(AaveLendingPoolLiquidationCall)
		if err := common.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack liquidation event error")
			return nil
		}
		return &LendingPoolLiquidationCall{
			Ts:               ts,
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
			Ns:               c.Config.Namespace,
			S:                vLog.Address.Hex(),
		}
	case "RebalanceStableBorrowRate":
		event := new(AaveLendingPoolRebalanceStableBorrowRate)
		if err := common.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack rebalance event error")
			return nil
		}
		return &LendingPoolRebalance{
			Ts:      ts,
			Block:   vLog.BlockNumber,
			Idx:     uint64(vLog.Index),
			Tx:      vLog.TxHash.Bytes(),
			Reserve: event.Reserve.Bytes(),
			User:    event.User.Bytes(),
			Ns:      c.Config.Namespace,
			S:       vLog.Address.Hex(),
		}
	case "Repay":
		event := new(AaveLendingPoolRepay)
		if err := common.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack repay event error")
			return nil
		}
		return &LendingPoolRepay{
			Ts:      ts,
			Block:   vLog.BlockNumber,
			Idx:     uint64(vLog.Index),
			Tx:      vLog.TxHash.Bytes(),
			Reserve: event.Reserve.Bytes(),
			User:    event.User.Bytes(),
			Repayer: event.Repayer.Bytes(),
			Amount:  event.Amount.Bytes(),
			Ns:      c.Config.Namespace,
			S:       vLog.Address.Hex(),
		}
	case "ReserveDataUpdated":
		event := new(AaveLendingPoolReserveDataUpdated)
		if err := common.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack reserve data updated event error")
			return nil
		}
		return &LendingPoolReserveUpdated{
			Ts:               ts,
			Block:            vLog.BlockNumber,
			Idx:              uint64(vLog.Index),
			Tx:               vLog.TxHash.Bytes(),
			LiquidityRate:    event.LiquidityRate.Bytes(),
			LiquidityIndex:   event.LiquidityIndex.Bytes(),
			StableBorrowRate: event.StableBorrowRate.Bytes(),
			VarBorrowIndex:   event.VariableBorrowIndex.Bytes(),
			VarBorrowRate:    event.VariableBorrowRate.Bytes(),
			Ns:               c.Config.Namespace,
			S:                vLog.Address.Hex(),
		}
	case "ReserveUsedAsCollateralDisabled":
		event := new(AaveLendingPoolReserveUsedAsCollateralDisabled)
		if err := common.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack reserve used as collateral disabled event error")
			return nil
		}
		return &LendingPoolReserveAsCollateral{
			Ts:         ts,
			Block:      vLog.BlockNumber,
			Idx:        uint64(vLog.Index),
			Tx:         vLog.TxHash.Bytes(),
			Reserve:    event.Reserve.Bytes(),
			User:       event.User.Bytes(),
			Collateral: false,
			Ns:         c.Config.Namespace,
			S:          vLog.Address.Hex(),
		}
	case "ReserveUsedAsCollateralEnabled":
		event := new(AaveLendingPoolReserveUsedAsCollateralEnabled)
		if err := common.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack reserve used as collateral disabled event error")
			return nil
		}
		return &LendingPoolReserveAsCollateral{
			Ts:         ts,
			Block:      vLog.BlockNumber,
			Idx:        uint64(vLog.Index),
			Tx:         vLog.TxHash.Bytes(),
			Reserve:    event.Reserve.Bytes(),
			User:       event.User.Bytes(),
			Collateral: true,
			Ns:         c.Config.Namespace,
			S:          vLog.Address.Hex(),
		}
	case "Swap":
		event := new(AaveLendingPoolSwap)
		if err := common.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack swap event error")
			return nil
		}
		return &LendingPoolSwap{
			Ts:       ts,
			Block:    vLog.BlockNumber,
			Idx:      uint64(vLog.Index),
			Tx:       vLog.TxHash.Bytes(),
			Reserve:  event.Reserve.Bytes(),
			User:     event.User.Bytes(),
			RateMode: event.RateMode.Bytes(),
			Ns:       c.Config.Namespace,
			S:        vLog.Address.Hex(),
		}
	case "Withdraw":
		event := new(AaveLendingPoolWithdraw)
		if err := common.UnpackLog(contractAbi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Unpack withdraw event error")
			return nil
		}
		return &LendingPoolWithdraw{
			Ts:      ts,
			Block:   vLog.BlockNumber,
			Idx:     uint64(vLog.Index),
			Tx:      vLog.TxHash.Bytes(),
			Reserve: event.Reserve.Bytes(),
			User:    event.User.Bytes(),
			To:      event.To.Bytes(),
			Amount:  event.Amount.Bytes(),
			Ns:      c.Config.Namespace,
			S:       vLog.Address.Hex(),
		}
	default:
		log.Error().Err(err).Msgf("invalid event: %s", ev.Name)
		return nil
	}
}
