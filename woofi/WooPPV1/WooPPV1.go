package WooPPV1

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
	contractAbi, err := abi.JSON(strings.NewReader(BscWooPPV1ABI))
	if err != nil {
		return nil, fmt.Errorf("error reading BscWooPPV1ABI: %s", err)
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
	case "ParametersUpdated":
		e := new(BscWooPPV1ParametersUpdated)
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
			NewLpFeeRate: e.NewLpFeeRate.Bytes(),
			NewR:         e.NewR.Bytes(),
		}
	case "Paused":
		e := new(BscWooPPV1Paused)
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
	case "RewardManagerUpdated":
		e := new(BscWooPPV1RewardManagerUpdated)
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
	case "WooGuardianUpdated":
		e := new(BscWooPPV1WooGuardianUpdated)
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
	case "OwnershipTransferPrepared":
		e := new(BscWooPPV1OwnershipTransferPrepared)
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
	case "Withdraw":
		e := new(BscWooPPV1Withdraw)
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
	case "WooracleUpdated":
		e := new(BscWooPPV1WooracleUpdated)
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
		e := new(BscWooPPV1StrategistUpdated)
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
	case "Unpaused":
		e := new(BscWooPPV1Unpaused)
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
	case "WooSwap":
		e := new(BscWooPPV1WooSwap)
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
			From:        e.FromToken.Bytes(),
			To:          e.To.Bytes(),
		}
	case "OwnershipTransferred":
		e := new(BscWooPPV1OwnershipTransferred)
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
	default:
		log.Error().Msgf("invalid event: %s", ev.Name)
		return nil
	}
}
