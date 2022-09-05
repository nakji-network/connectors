package lendingpool

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nakji-network/connector/common"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SmartContract struct {
	Addr      string
	Abi       abi.ABI
	Namespace string
}

func NewContract(addr string, ns string) (*SmartContract, error) {
	contractAbi, err := abi.JSON(strings.NewReader(AaveLendingPoolABI))
	if err != nil {
		return nil, fmt.Errorf("error reading aave lending pool abi: %s", err)
	}
	return &SmartContract{Addr: addr, Abi: contractAbi, Namespace: ns}, nil
}

func (c *SmartContract) Address() string {
	return c.Addr
}

func (c *SmartContract) Message(vLog types.Log, ts *timestamppb.Timestamp) proto.Message {
	ev, err := c.Abi.EventByID(vLog.Topics[0])
	if err != nil {
		log.Warn().Err(err).Msg("EventByID error, skipping")
		return nil
	}
	switch ev.Name {
	case "Borrow":
		event := new(AaveLendingPoolBorrow)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
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
			Ns:         c.Namespace,
			S:          vLog.Address.Hex(),
		}
	case "Deposit":
		event := new(AaveLendingPoolDeposit)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
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
			Ns:       c.Namespace,
			S:        vLog.Address.Hex(),
		}
	case "FlashLoan":
		event := new(AaveLendingPoolFlashLoan)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
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
			Ns:        c.Namespace,
			S:         vLog.Address.Hex(),
		}
	case "LiquidationCall":
		event := new(AaveLendingPoolLiquidationCall)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
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
			Ns:               c.Namespace,
			S:                vLog.Address.Hex(),
		}
	case "RebalanceStableBorrowRate":
		event := new(AaveLendingPoolRebalanceStableBorrowRate)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
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
			Ns:      c.Namespace,
			S:       vLog.Address.Hex(),
		}
	case "Repay":
		event := new(AaveLendingPoolRepay)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
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
			Ns:      c.Namespace,
			S:       vLog.Address.Hex(),
		}
	case "ReserveDataUpdated":
		event := new(AaveLendingPoolReserveDataUpdated)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
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
			Ns:               c.Namespace,
			S:                vLog.Address.Hex(),
		}
	case "ReserveUsedAsCollateralDisabled":
		event := new(AaveLendingPoolReserveUsedAsCollateralDisabled)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
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
			Ns:         c.Namespace,
			S:          vLog.Address.Hex(),
		}
	case "ReserveUsedAsCollateralEnabled":
		event := new(AaveLendingPoolReserveUsedAsCollateralEnabled)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
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
			Ns:         c.Namespace,
			S:          vLog.Address.Hex(),
		}
	case "Swap":
		event := new(AaveLendingPoolSwap)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
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
			Ns:       c.Namespace,
			S:        vLog.Address.Hex(),
		}
	case "Withdraw":
		event := new(AaveLendingPoolWithdraw)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
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
			Ns:      c.Namespace,
			S:       vLog.Address.Hex(),
		}
	default:
		log.Error().Msgf("invalid event: %s", ev.Name)
		return nil
	}
}
