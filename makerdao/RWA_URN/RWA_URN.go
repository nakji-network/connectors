
package RWA_URN

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
		e := new(RWAURNDeny)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	case "Draw":
		e := new(RWAURNDraw)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Draw{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
				Wad:  e.Wad.Bytes(),
		}
	case "File":
		e := new(RWAURNFile)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &File{
				Ts:   timestamp,
				What:  e.What[:],
				Data:  e.Data.Bytes(),
		}
	case "Free":
		e := new(RWAURNFree)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Free{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
				Wad:  e.Wad.Bytes(),
		}
	case "Hope":
		e := new(RWAURNHope)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Hope{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	case "Lock":
		e := new(RWAURNLock)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Lock{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
				Wad:  e.Wad.Bytes(),
		}
	case "Nope":
		e := new(RWAURNNope)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Nope{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	case "Quit":
		e := new(RWAURNQuit)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Quit{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
				Wad:  e.Wad.Bytes(),
		}
	case "Rely":
		e := new(RWAURNRely)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Rely{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	case "Wipe":
		e := new(RWAURNWipe)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Wipe{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
				Wad:  e.Wad.Bytes(),
		}
	}
	return nil
}
