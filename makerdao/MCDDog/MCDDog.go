
package MCDDog

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
	case "Bark":
		e := new(MCDDOGBark)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Bark{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
				Urn:  e.Urn.Bytes(),
				Ink:  e.Ink.Bytes(),
				Art:  e.Art.Bytes(),
				Due:  e.Due.Bytes(),
				Clip:  e.Clip.Bytes(),
				Id:  e.Id.Bytes(),
		}
	case "Cage":
		e := new(MCDDOGCage)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Cage{
				Ts:   timestamp,
		}
	case "Deny":
		e := new(MCDDOGDeny)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	case "Digs":
		e := new(MCDDOGDigs)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Digs{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
				Rad:  e.Rad.Bytes(),
		}
	case "File":
		e := new(MCDDOGFile)
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
		e := new(MCDDOGFile0)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &File0{
				Ts:   timestamp,
				What:  e.What[:],
				Data:  e.Data.Bytes(),
		}
	case "File1":
		e := new(MCDDOGFile1)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &File1{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
				What:  e.What[:],
				Data:  e.Data.Bytes(),
		}
	case "File2":
		e := new(MCDDOGFile2)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &File2{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
				What:  e.What[:],
				Clip:  e.Clip.Bytes(),
		}
	case "Rely":
		e := new(MCDDOGRely)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Rely{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	}
	return nil
}
