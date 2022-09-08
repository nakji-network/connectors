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
	Addr string
	Abi  abi.ABI
}

func NewContract(addr string) (*SmartContract, error) {
	contractAbi, err := abi.JSON(strings.NewReader(LendingpoolABI))
	if err != nil {
		return nil, fmt.Errorf("error reading aave lending pool abi: %s", err)
	}
	return &SmartContract{Addr: addr, Abi: contractAbi}, nil
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
		event := new(LendingpoolBorrow)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msgf("error unpacking event: %s", ev.Name)
			return nil
		}
		return &Borrow{
			Ts:             ts,
			BlockNumber:    vLog.BlockNumber,
			LogIndex:       uint64(vLog.Index),
			TxHash:         vLog.TxHash.Bytes(),
			Reserve:        event.Reserve.Bytes(),
			User:           event.User.Bytes(),
			OnBehalfOf:     event.OnBehalfOf.Bytes(),
			Amount:         event.Amount.Bytes(),
			BorrowRate:     event.BorrowRate.Bytes(),
			BorrowRateMode: event.BorrowRateMode.Bytes(),
			Referral:       uint32(event.Referral),
		}
	case "Deposit":
		event := new(LendingpoolDeposit)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msgf("error unpacking event: %s", ev.Name)
			return nil
		}
		return &Deposit{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			LogIndex:    uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Reserve:     event.Reserve.Bytes(),
			User:        event.User.Bytes(),
			OnBehalfOf:  event.OnBehalfOf.Bytes(),
			Amount:      event.Amount.Bytes(),
			Referral:    uint32(event.Referral),
		}
	case "FlashLoan":
		event := new(LendingpoolFlashLoan)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msgf("error unpacking event: %s", ev.Name)
			return nil
		}
		return &FlashLoan{
			Ts:           ts,
			BlockNumber:  vLog.BlockNumber,
			LogIndex:     uint64(vLog.Index),
			TxHash:       vLog.TxHash.Bytes(),
			Target:       event.Target.Bytes(),
			Initiator:    event.Initiator.Bytes(),
			Asset:        event.Asset.Bytes(),
			Amount:       event.Amount.Bytes(),
			Premium:      event.Premium.Bytes(),
			ReferralCode: uint32(event.ReferralCode),
		}
	case "LiquidationCall":
		event := new(LendingpoolLiquidationCall)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msgf("error unpacking event: %s", ev.Name)
			return nil
		}
		return &LiquidationCall{
			Ts:                         ts,
			BlockNumber:                vLog.BlockNumber,
			LogIndex:                   uint64(vLog.Index),
			TxHash:                     vLog.TxHash.Bytes(),
			CollateralAsset:            event.CollateralAsset.Bytes(),
			DebtAsset:                  event.DebtAsset.Bytes(),
			User:                       event.User.Bytes(),
			DebtToCover:                event.DebtToCover.Bytes(),
			LiquidatedCollateralAmount: event.LiquidatedCollateralAmount.Bytes(),
			Liquidator:                 event.Liquidator.Bytes(),
			ReceiveAToken:              event.ReceiveAToken,
		}
	case "RebalanceStableBorrowRate":
		event := new(LendingpoolRebalanceStableBorrowRate)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msgf("error unpacking event: %s", ev.Name)
			return nil
		}
		return &RebalanceStableBorrowRate{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			LogIndex:    uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Reserve:     event.Reserve.Bytes(),
			User:        event.User.Bytes(),
		}
	case "Repay":
		event := new(LendingpoolRepay)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msgf("error unpacking event: %s", ev.Name)
			return nil
		}
		return &Repay{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			LogIndex:    uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Reserve:     event.Reserve.Bytes(),
			User:        event.User.Bytes(),
			Repayer:     event.Repayer.Bytes(),
			Amount:      event.Amount.Bytes(),
		}
	case "ReserveDataUpdated":
		event := new(LendingpoolReserveDataUpdated)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msgf("error unpacking event: %s", ev.Name)
			return nil
		}
		return &ReserveDataUpdated{
			Ts:                  ts,
			BlockNumber:         vLog.BlockNumber,
			LogIndex:            uint64(vLog.Index),
			TxHash:              vLog.TxHash.Bytes(),
			LiquidityRate:       event.LiquidityRate.Bytes(),
			LiquidityIndex:      event.LiquidityIndex.Bytes(),
			StableBorrowRate:    event.StableBorrowRate.Bytes(),
			VariableBorrowIndex: event.VariableBorrowIndex.Bytes(),
			VariableBorrowRate:  event.VariableBorrowRate.Bytes(),
		}
	case "ReserveUsedAsCollateralDisabled":
		event := new(LendingpoolReserveUsedAsCollateralDisabled)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msgf("error unpacking event: %s", ev.Name)
			return nil
		}
		return &ReserveUsedAsCollateralDisabled{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			LogIndex:    uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Reserve:     event.Reserve.Bytes(),
			User:        event.User.Bytes(),
		}
	case "ReserveUsedAsCollateralEnabled":
		event := new(LendingpoolReserveUsedAsCollateralEnabled)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msgf("error unpacking event: %s", ev.Name)
			return nil
		}
		return &ReserveUsedAsCollateralEnabled{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			LogIndex:    uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Reserve:     event.Reserve.Bytes(),
			User:        event.User.Bytes(),
		}
	case "Swap":
		event := new(LendingpoolSwap)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msgf("error unpacking event: %s", ev.Name)
			return nil
		}
		return &Swap{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			LogIndex:    uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Reserve:     event.Reserve.Bytes(),
			User:        event.User.Bytes(),
			RateMode:    event.RateMode.Bytes(),
		}
	case "Withdraw":
		event := new(LendingpoolWithdraw)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msgf("error unpacking event: %s", ev.Name)
			return nil
		}
		return &Withdraw{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			LogIndex:    uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Reserve:     event.Reserve.Bytes(),
			User:        event.User.Bytes(),
			To:          event.To.Bytes(),
			Amount:      event.Amount.Bytes(),
		}
	case "Paused":
		event := new(LendingpoolPaused)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msgf("error unpacking event: %s", ev.Name)
			return nil
		}
		return &Paused{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			LogIndex:    uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
		}
	case "Unpaused":
		event := new(LendingpoolUnpaused)
		if err := common.UnpackLog(c.Abi, event, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msgf("error unpacking event: %s", ev.Name)
			return nil
		}
		return &Unpaused{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			LogIndex:    uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
		}
	default:
		log.Error().Msgf("invalid event: %s", ev.Name)
		return nil
	}
}
