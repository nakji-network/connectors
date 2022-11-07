package WooPPV2

import (
	"strings"

	"github.com/nakji-network/connector/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SmartContract struct {
	addr string
	abi  abi.ABI
}

func NewContract(addr string) *SmartContract {
	contractAbi, err := abi.JSON(strings.NewReader(WooPPV2ABI))
	if err != nil {
		log.Fatal().Err(err).Msg("error reading WooPPV2ABI")
	}
	return &SmartContract{addr: addr, abi: contractAbi}
}

func (sc *SmartContract) Address() string {
	return sc.addr
}

func (sc *SmartContract) Events() []proto.Message {
	return []proto.Message{
		&OwnershipTransferPrepared{},
		&OwnershipTransferred{},
		&ParametersUpdated{},
		&Paused{},
		&RewardManagerUpdated{},
		&StrategistUpdated{},
		&Unpaused{},
		&Withdraw{},
		&WooGuardianUpdated{},
		&WooracleUpdated{},
		&WooSwap{},
	}
}

func (sc *SmartContract) Message(vLog types.Log, ts *timestamppb.Timestamp) proto.Message {
	ev, err := sc.abi.EventByID(vLog.Topics[0])
	if err != nil {
		log.Warn().Err(err).Msg("EventByID error, skipping")
		return nil
	}
	switch ev.Name {
	case "ParametersUpdated":
		e := new(WooPPV2ParametersUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &ParametersUpdated{
			Ts:              ts,
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			BaseToken:       e.BaseToken.Bytes(),
			NewThreshold:    e.NewThreshold.Bytes(),
			NewLpFeeRate:    e.NewLpFeeRate.Bytes(),
			NewR:            e.NewR.Bytes(),
			ContractAddress: sc.Address(),
		}
	case "Paused":
		e := new(WooPPV2Paused)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &Paused{
			Ts:              ts,
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			Account:         e.Account.Bytes(),
			ContractAddress: sc.Address(),
		}
	case "RewardManagerUpdated":
		e := new(WooPPV2RewardManagerUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &RewardManagerUpdated{
			Ts:               ts,
			BlockNumber:      vLog.BlockNumber,
			Index:            uint64(vLog.Index),
			TxHash:           vLog.TxHash.Bytes(),
			NewRewardManager: e.NewRewardManager.Bytes(),
			ContractAddress:  sc.Address(),
		}
	case "WooGuardianUpdated":
		e := new(WooPPV2WooGuardianUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &WooGuardianUpdated{
			Ts:              ts,
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			NewWooGuardian:  e.NewWooGuardian.Bytes(),
			ContractAddress: sc.Address(),
		}
	case "OwnershipTransferPrepared":
		e := new(WooPPV2OwnershipTransferPrepared)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &OwnershipTransferPrepared{
			Ts:              ts,
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			PreviousOwner:   e.PreviousOwner.Bytes(),
			NewOwner:        e.NewOwner.Bytes(),
			ContractAddress: sc.Address(),
		}
	case "Withdraw":
		e := new(WooPPV2Withdraw)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &Withdraw{
			Ts:              ts,
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			Token:           e.Token.Bytes(),
			To:              e.To.Bytes(),
			Amount:          e.Amount.Bytes(),
			ContractAddress: sc.Address(),
		}
	case "WooracleUpdated":
		e := new(WooPPV2WooracleUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &WooracleUpdated{
			Ts:              ts,
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			NewWooracle:     e.NewWooracle.Bytes(),
			ContractAddress: sc.Address(),
		}
	case "StrategistUpdated":
		e := new(WooPPV2StrategistUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &StrategistUpdated{
			Ts:              ts,
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			Strategist:      e.Strategist.Bytes(),
			Flag:            e.Flag,
			ContractAddress: sc.Address(),
		}
	case "Unpaused":
		e := new(WooPPV2Unpaused)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &Unpaused{
			Ts:              ts,
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			Account:         e.Account.Bytes(),
			ContractAddress: sc.Address(),
		}
	case "WooSwap":
		e := new(WooPPV2WooSwap)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &WooSwap{
			Ts:              ts,
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			FromToken:       e.FromToken.Bytes(),
			ToToken:         e.ToToken.Bytes(),
			FromAmount:      e.FromAmount.Bytes(),
			ToAmount:        e.ToAmount.Bytes(),
			From:            e.From.Bytes(),
			To:              e.To.Bytes(),
			ContractAddress: sc.Address(),
		}
	case "OwnershipTransferred":
		e := new(WooPPV2OwnershipTransferred)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &OwnershipTransferred{
			Ts:              ts,
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			PreviousOwner:   e.PreviousOwner.Bytes(),
			NewOwner:        e.NewOwner.Bytes(),
			ContractAddress: sc.Address(),
		}
	default:
		log.Error().Msgf("invalid event: %s", ev.Name)
		return nil
	}
}
