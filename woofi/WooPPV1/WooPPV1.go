package WooPPV1

import (
	"github.com/nakji-network/connector/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SmartContract struct{}

func (sc *SmartContract) Message(eventName string, contractAbi *abi.ABI, vLog types.Log, ts *timestamppb.Timestamp) protoreflect.ProtoMessage {
	switch eventName {
	case "ParametersUpdated":
		e := new(WooPPV1ParametersUpdated)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &ParametersUpdated{
			Ts:           ts,
			BlockNumber:  vLog.BlockNumber,
			Index:        uint64(vLog.Index),
			TxHash:       vLog.TxHash.Bytes(),
			BaseToken:    e.BaseToken.Bytes(),
			NewThreshold: e.NewThreshold.Bytes(),
			NewLpFeeRate: e.NewLpFeeRate.Bytes(),
			NewR:         e.NewR.Bytes(),
		}
	case "Paused":
		e := new(WooPPV1Paused)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &Paused{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			Index:       uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Account:     e.Account.Bytes(),
		}
	case "RewardManagerUpdated":
		e := new(WooPPV1RewardManagerUpdated)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &RewardManagerUpdated{
			Ts:               ts,
			BlockNumber:      vLog.BlockNumber,
			Index:            uint64(vLog.Index),
			TxHash:           vLog.TxHash.Bytes(),
			NewRewardManager: e.NewRewardManager.Bytes(),
		}
	case "WooGuardianUpdated":
		e := new(WooPPV1WooGuardianUpdated)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &WooGuardianUpdated{
			Ts:             ts,
			BlockNumber:    vLog.BlockNumber,
			Index:          uint64(vLog.Index),
			TxHash:         vLog.TxHash.Bytes(),
			NewWooGuardian: e.NewWooGuardian.Bytes(),
		}
	case "OwnershipTransferPrepared":
		e := new(WooPPV1OwnershipTransferPrepared)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &OwnershipTransferPrepared{
			Ts:            ts,
			BlockNumber:   vLog.BlockNumber,
			Index:         uint64(vLog.Index),
			TxHash:        vLog.TxHash.Bytes(),
			PreviousOwner: e.PreviousOwner.Bytes(),
			NewOwner:      e.NewOwner.Bytes(),
		}
	case "Withdraw":
		e := new(WooPPV1Withdraw)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &Withdraw{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			Index:       uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Token:       e.Token.Bytes(),
			To:          e.To.Bytes(),
			Amount:      e.Amount.Bytes(),
		}
	case "WooracleUpdated":
		e := new(WooPPV1WooracleUpdated)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &WooracleUpdated{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			Index:       uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			NewWooracle: e.NewWooracle.Bytes(),
		}
	case "StrategistUpdated":
		e := new(WooPPV1StrategistUpdated)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &StrategistUpdated{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			Index:       uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Strategist:  e.Strategist.Bytes(),
			Flag:        e.Flag,
		}
	case "Unpaused":
		e := new(WooPPV1Unpaused)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &Unpaused{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			Index:       uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			Account:     e.Account.Bytes(),
		}
	case "WooSwap":
		e := new(WooPPV1WooSwap)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &WooSwap{
			Ts:          ts,
			BlockNumber: vLog.BlockNumber,
			Index:       uint64(vLog.Index),
			TxHash:      vLog.TxHash.Bytes(),
			FromToken:   e.FromToken.Bytes(),
			ToToken:     e.ToToken.Bytes(),
			FromAmount:  e.FromAmount.Bytes(),
			ToAmount:    e.ToAmount.Bytes(),
			From:        e.FromToken.Bytes(),
			To:          e.To.Bytes(),
		}
	case "OwnershipTransferred":
		e := new(WooPPV1OwnershipTransferred)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &OwnershipTransferred{
			Ts:            ts,
			BlockNumber:   vLog.BlockNumber,
			Index:         uint64(vLog.Index),
			TxHash:        vLog.TxHash.Bytes(),
			PreviousOwner: e.PreviousOwner.Bytes(),
			NewOwner:      e.NewOwner.Bytes(),
		}
	default:
		log.Error().Msgf("invalid event: %s", eventName)
		return nil
	}
}
