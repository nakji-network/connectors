package WooPPV3

import (
	"strings"

	"github.com/nakji-network/connector/common"
	"github.com/nakji-network/connectors/woofi_polygon"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SmartContract struct {
	abi       abi.ABI
	addr      string
	addrBytes []byte
}

func NewContract(addr string) *SmartContract {
	contractAbi, err := abi.JSON(strings.NewReader(WooPPV3ABI))
	if err != nil {
		log.Fatal().Err(err).Msg("error reading WooPPV3ABI")
	}
	return &SmartContract{abi: contractAbi, addr: addr, addrBytes: ethcommon.HexToAddress(addr).Bytes()}
}

func (sc *SmartContract) Address() string {
	return sc.addr
}

func (sc *SmartContract) Events() []proto.Message {
	return []proto.Message{
		&FeeManagerUpdated{},
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

func (sc *SmartContract) Message(vLog woofi_polygon.Log) proto.Message {
	ev, err := sc.abi.EventByID(vLog.Topics[0])
	if err != nil {
		log.Warn().Err(err).Msg("EventByID error, skipping")
		return nil
	}
	switch ev.Name {
	case "Paused":
		e := new(WooPPV3Paused)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &Paused{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			Account:         e.Account.Bytes(),
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	case "OwnershipTransferred":
		e := new(WooPPV3OwnershipTransferred)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &OwnershipTransferred{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			PreviousOwner:   e.PreviousOwner.Bytes(),
			NewOwner:        e.NewOwner.Bytes(),
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	case "ParametersUpdated":
		e := new(WooPPV3ParametersUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &ParametersUpdated{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			BaseToken:       e.BaseToken.Bytes(),
			NewThreshold:    e.NewThreshold.Bytes(),
			NewR:            e.NewR.Bytes(),
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	case "FeeManagerUpdated":
		e := new(WooPPV3FeeManagerUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &FeeManagerUpdated{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			NewFeeManager:   e.NewFeeManager.Bytes(),
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	case "RewardManagerUpdated":
		e := new(WooPPV3RewardManagerUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &RewardManagerUpdated{
			Ts:               timestamppb.New(vLog.BlockTime),
			BlockNumber:      vLog.BlockNumber,
			Index:            uint64(vLog.Index),
			TxHash:           vLog.TxHash.Bytes(),
			NewRewardManager: e.NewRewardManager.Bytes(),
			ContractAddress:  sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	case "Unpaused":
		e := new(WooPPV3Unpaused)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &Unpaused{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			Account:         e.Account.Bytes(),
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	case "WooGuardianUpdated":
		e := new(WooPPV3WooGuardianUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &WooGuardianUpdated{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			NewWooGuardian:  e.NewWooGuardian.Bytes(),
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	case "WooSwap":
		e := new(WooPPV3WooSwap)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &WooSwap{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			FromToken:       e.FromToken.Bytes(),
			ToToken:         e.ToToken.Bytes(),
			FromAmount:      e.FromAmount.Bytes(),
			ToAmount:        e.ToAmount.Bytes(),
			From:            e.From.Bytes(),
			To:              e.To.Bytes(),
			RebateTo:        e.RebateTo.Bytes(),
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	case "OwnershipTransferPrepared":
		e := new(WooPPV3OwnershipTransferPrepared)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &OwnershipTransferPrepared{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			PreviousOwner:   e.PreviousOwner.Bytes(),
			NewOwner:        e.NewOwner.Bytes(),
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	case "WooracleUpdated":
		e := new(WooPPV3WooracleUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &WooracleUpdated{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			NewWooracle:     e.NewWooracle.Bytes(),
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	case "StrategistUpdated":
		e := new(WooPPV3StrategistUpdated)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &StrategistUpdated{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			Strategist:      e.Strategist.Bytes(),
			Flag:            e.Flag,
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	case "Withdraw":
		e := new(WooPPV3Withdraw)
		if err := common.UnpackLog(sc.abi, e, ev.Name, vLog.Log); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}
		msg := &Withdraw{
			Ts:              timestamppb.New(vLog.BlockTime),
			BlockNumber:     vLog.BlockNumber,
			Index:           uint64(vLog.Index),
			TxHash:          vLog.TxHash.Bytes(),
			Token:           e.Token.Bytes(),
			To:              e.To.Bytes(),
			Amount:          e.Amount.Bytes(),
			ContractAddress: sc.addrBytes,
		}
		if vLog.SenderAddress != nil {
			msg.SenderAddress = vLog.SenderAddress.Bytes()
		}
		if vLog.ReceiverAddress != nil {
			msg.ReceiverAddress = vLog.ReceiverAddress.Bytes()
		}
		return msg
	default:
		log.Error().Msgf("invalid event: %s", ev.Name)
		return nil
	}
}
