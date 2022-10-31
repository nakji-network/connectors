package WOOPP

import (
	"github.com/nakji-network/connector/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SmartContract struct{}

func (sc *SmartContract) Message(eventName string, contractAbi *abi.ABI, vLog types.Log, timestamp *timestamppb.Timestamp) protoreflect.ProtoMessage {
	switch eventName {
	case "Paused":
		e := new(WOOPPPaused)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Paused{
				Ts:   timestamp,
				BlockNumber: vLog.BlockNumber,
				Index: uint64(vLog.Index),
				TxHash: vLog.TxHash.Bytes(),
				Account:  e.Account.Bytes(),
		}
	case "OwnershipTransferred":
		e := new(WOOPPOwnershipTransferred)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &OwnershipTransferred{
				Ts:   timestamp,
				BlockNumber: vLog.BlockNumber,
				Index: uint64(vLog.Index),
				TxHash: vLog.TxHash.Bytes(),
				PreviousOwner:  e.PreviousOwner.Bytes(),
				NewOwner:  e.NewOwner.Bytes(),
		}
	case "ParametersUpdated":
		e := new(WOOPPParametersUpdated)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &ParametersUpdated{
				Ts:   timestamp,
				BlockNumber: vLog.BlockNumber,
				Index: uint64(vLog.Index),
				TxHash: vLog.TxHash.Bytes(),
				BaseToken:  e.BaseToken.Bytes(),
				NewThreshold:  e.NewThreshold.Bytes(),
				NewR: e.NewR.Bytes(),
		}
	case "FeeManagerUpdated":
		e := new(WOOPPFeeManagerUpdated)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &FeeManagerUpdated{
				Ts:   timestamp,
				BlockNumber: vLog.BlockNumber,
				Index: uint64(vLog.Index),
				TxHash: vLog.TxHash.Bytes(),
				NewFeeManager:  e.NewFeeManager.Bytes(),
		}
	case "RewardManagerUpdated":
		e := new(WOOPPRewardManagerUpdated)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &RewardManagerUpdated{
				Ts:   timestamp,
				BlockNumber: vLog.BlockNumber,
				Index: uint64(vLog.Index),
				TxHash: vLog.TxHash.Bytes(),
				NewRewardManager:  e.NewRewardManager.Bytes(),
		}
	case "Unpaused":
		e := new(WOOPPUnpaused)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Unpaused{
				Ts:   timestamp,
				BlockNumber: vLog.BlockNumber,
				Index: uint64(vLog.Index),
				TxHash: vLog.TxHash.Bytes(),
				Account:  e.Account.Bytes(),
		}
	case "WooGuardianUpdated":
		e := new(WOOPPWooGuardianUpdated)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &WooGuardianUpdated{
				Ts:   timestamp,
				BlockNumber: vLog.BlockNumber,
				Index: uint64(vLog.Index),
				TxHash: vLog.TxHash.Bytes(),
				NewWooGuardian:  e.NewWooGuardian.Bytes(),
		}
	case "WooSwap":
		e := new(WOOPPWooSwap)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &WooSwap{
				Ts:   timestamp,
				BlockNumber: vLog.BlockNumber,
				Index: uint64(vLog.Index),
				TxHash: vLog.TxHash.Bytes(),
				FromToken:  e.FromToken.Bytes(),
				ToToken:  e.ToToken.Bytes(),
				FromAmount:  e.FromAmount.Bytes(),
				ToAmount:  e.ToAmount.Bytes(),
				From: e.From.Bytes(),
				To: e.To.Bytes(),
				RebateTo: e.RebateTo.Bytes(),
		}
	case "OwnershipTransferPrepared":
		e := new(WOOPPOwnershipTransferPrepared)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &OwnershipTransferPrepared{
				Ts:   timestamp,
				BlockNumber: vLog.BlockNumber,
				Index: uint64(vLog.Index),
				TxHash: vLog.TxHash.Bytes(),
				PreviousOwner: e.PreviousOwner.Bytes(),
				NewOwner: e.NewOwner.Bytes(),
		}
	case "WooracleUpdated":
		e := new(WOOPPWooracleUpdated)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &WooracleUpdated{
				Ts:   timestamp,
				BlockNumber: vLog.BlockNumber,
				Index: uint64(vLog.Index),
				TxHash: vLog.TxHash.Bytes(),
				NewWooracle: e.NewWooracle.Bytes(),
		}
	case "StrategistUpdated":
		e := new(WOOPPStrategistUpdated)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &StrategistUpdated{
				Ts:   timestamp,
				BlockNumber: vLog.BlockNumber,
				Index: uint64(vLog.Index),
				TxHash: vLog.TxHash.Bytes(),
				Strategist: e.Strategist.Bytes(),
				Flag: e.Flag,
		}
	case "Withdraw":
		e := new(WOOPPWithdraw)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Withdraw{
				Ts:   timestamp,
				BlockNumber: vLog.BlockNumber,
				Index: uint64(vLog.Index),
				TxHash: vLog.TxHash.Bytes(),
				Token: e.Token.Bytes(),
				To: e.To.Bytes(),
				Amount: e.Amount.Bytes(),
		}
	}
	return nil
}
