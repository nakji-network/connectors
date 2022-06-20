package CHANGELOG

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
	case "Deny":
		e := new(CHANGELOGDeny)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
			Ts:  timestamp,
			Usr: e.Usr.Bytes(),
		}
	case "Rely":
		e := new(CHANGELOGRely)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Rely{
			Ts:  timestamp,
			Usr: e.Usr.Bytes(),
		}
	case "RemoveAddress":
		e := new(CHANGELOGRemoveAddress)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &RemoveAddress{
			Ts:  timestamp,
			Key: e.Key[:],
		}
	case "UpdateAddress":
		e := new(CHANGELOGUpdateAddress)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &UpdateAddress{
			Ts:   timestamp,
			Key:  e.Key[:],
			Addr: e.Addr.Bytes(),
		}
	case "UpdateIPFS":
		e := new(CHANGELOGUpdateIPFS)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &UpdateIPFS{
			Ts: timestamp,
		}
	case "UpdateSha256sum":
		e := new(CHANGELOGUpdateSha256sum)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &UpdateSha256Sum{
			Ts: timestamp,
		}
	case "UpdateVersion":
		e := new(CHANGELOGUpdateVersion)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &UpdateVersion{
			Ts: timestamp,
		}
	}
	return nil
}
