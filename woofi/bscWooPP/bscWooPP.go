package bscWooPP

import (
	"fmt"
	"strings"

	"github.com/nakji-network/connector/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SmartContract struct {
	network string
	address string
	abi     abi.ABI
}

func NewContract(network string, address string) (*SmartContract, error) {
	contractAbi, err := abi.JSON(strings.NewReader(BscWooPPABI))
	if err != nil {
		return nil, fmt.Errorf("error reading BscWooPPABI: %s", err)
	}
	return &SmartContract{network: network, address: address, abi: contractAbi}, nil
}

func (sc *SmartContract) Network() string {
	return sc.network
}

func (sc *SmartContract) Address() string {
	return sc.address
}

func (sc *SmartContract) Message(vLog types.Log, ts *timestamppb.Timestamp) proto.Message {
	ev, err := sc.abi.EventByID(vLog.Topics[0])
	if err != nil {
		log.Warn().Err(err).Msg("EventByID error, skipping")
		return nil
	}
	switch ev.Name {
	case "Paused":
		e := new(BscWooPPPaused)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
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
	case "OwnershipTransferred":
		e := new(BscWooPPOwnershipTransferred)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
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
	case "ParametersUpdated":
		e := new(BscWooPPParametersUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
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
			NewR:         e.NewR.Bytes(),
		}
	case "FeeManagerUpdated":
		e := new(BscWooPPFeeManagerUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		return &FeeManagerUpdated{
			Ts:            ts,
			BlockNumber:   vLog.BlockNumber,
			Index:         uint64(vLog.Index),
			TxHash:        vLog.TxHash.Bytes(),
			NewFeeManager: e.NewFeeManager.Bytes(),
		}
	case "RewardManagerUpdated":
		e := new(BscWooPPRewardManagerUpdated)
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
		}
	case "Unpaused":
		e := new(BscWooPPUnpaused)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
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
	case "WooGuardianUpdated":
		e := new(BscWooPPWooGuardianUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
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
	case "WooSwap":
		e := new(BscWooPPWooSwap)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
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
			From:        e.From.Bytes(),
			To:          e.To.Bytes(),
			RebateTo:    e.RebateTo.Bytes(),
		}
	case "OwnershipTransferPrepared":
		e := new(BscWooPPOwnershipTransferPrepared)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
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
	case "WooracleUpdated":
		e := new(BscWooPPWooracleUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
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
		e := new(BscWooPPStrategistUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
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
	case "Withdraw":
		e := new(BscWooPPWithdraw)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog); err != nil {
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
	default:
		log.Error().Msgf("invalid event: %s", ev.Name)
		return nil
	}
}
