
package RWA_CONDUIT

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
	case "Cage":
		e := new(RWACONDUITCage)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Cage{
				Ts:   timestamp,
		}
	case "Cull":
		e := new(RWACONDUITCull)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Cull{
				Ts:   timestamp,
				Tab:  e.Tab.Bytes(),
		}
	case "Deny":
		e := new(RWACONDUITDeny)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	case "Draw":
		e := new(RWACONDUITDraw)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Draw{
				Ts:   timestamp,
				Wad:  e.Wad.Bytes(),
		}
	case "Exit":
		e := new(RWACONDUITExit)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Exit{
				Ts:   timestamp,
				Wad:  e.Wad.Bytes(),
		}
	case "File":
		e := new(RWACONDUITFile)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &File{
				Ts:   timestamp,
				What:  e.What[:],
				Data:  e.Data.Bytes(),
		}
	case "Join":
		e := new(RWACONDUITJoin)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Join{
				Ts:   timestamp,
				Wad:  e.Wad.Bytes(),
		}
	case "Migrate":
		e := new(RWACONDUITMigrate)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Migrate{
				Ts:   timestamp,
				Dst:  e.Dst.Bytes(),
		}
	case "Recover":
		e := new(RWACONDUITRecover)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Recover{
				Ts:   timestamp,
				Recovered:  e.Recovered.Bytes(),
				PayBack:  e.PayBack.Bytes(),
		}
	case "Rely":
		e := new(RWACONDUITRely)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Rely{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	case "Tell":
		e := new(RWACONDUITTell)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Tell{
				Ts:   timestamp,
				Wad:  e.Wad.Bytes(),
		}
	case "Unwind":
		e := new(RWACONDUITUnwind)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Unwind{
				Ts:   timestamp,
				PayBack:  e.PayBack.Bytes(),
		}
	case "Wipe":
		e := new(RWACONDUITWipe)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Wipe{
				Ts:   timestamp,
				Wad:  e.Wad.Bytes(),
		}
	}
	return nil
}
