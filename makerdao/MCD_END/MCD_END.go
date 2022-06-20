
package MCD_END

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
		e := new(MCDENDCage)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Cage{
				Ts:   timestamp,
		}
	case "Cage0":
		e := new(MCDENDCage0)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Cage0{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
		}
	case "Cash":
		e := new(MCDENDCash)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Cash{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
				Usr:  e.Usr.Bytes(),
				Wad:  e.Wad.Bytes(),
		}
	case "Deny":
		e := new(MCDENDDeny)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	case "File":
		e := new(MCDENDFile)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &File{
				Ts:   timestamp,
				What:  e.What[:],
				Data:  e.Data.Bytes(),
		}
	case "File0":
		e := new(MCDENDFile0)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &File0{
				Ts:   timestamp,
				What:  e.What[:],
				Data:  e.Data.Bytes(),
		}
	case "Flow":
		e := new(MCDENDFlow)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Flow{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
		}
	case "Free":
		e := new(MCDENDFree)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Free{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
				Usr:  e.Usr.Bytes(),
				Ink:  e.Ink.Bytes(),
		}
	case "Pack":
		e := new(MCDENDPack)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Pack{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
				Wad:  e.Wad.Bytes(),
		}
	case "Rely":
		e := new(MCDENDRely)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Rely{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	case "Skim":
		e := new(MCDENDSkim)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Skim{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
				Urn:  e.Urn.Bytes(),
				Wad:  e.Wad.Bytes(),
				Art:  e.Art.Bytes(),
		}
	case "Skip":
		e := new(MCDENDSkip)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Skip{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
				Id:  e.Id.Bytes(),
				Usr:  e.Usr.Bytes(),
				Tab:  e.Tab.Bytes(),
				Lot:  e.Lot.Bytes(),
				Art:  e.Art.Bytes(),
		}
	case "Snip":
		e := new(MCDENDSnip)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Snip{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
				Id:  e.Id.Bytes(),
				Usr:  e.Usr.Bytes(),
				Tab:  e.Tab.Bytes(),
				Lot:  e.Lot.Bytes(),
				Art:  e.Art.Bytes(),
		}
	case "Thaw":
		e := new(MCDENDThaw)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Thaw{
				Ts:   timestamp,
		}
	}
	return nil
}
