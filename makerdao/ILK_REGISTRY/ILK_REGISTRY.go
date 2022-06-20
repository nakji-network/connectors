
package ILK_REGISTRY

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
	case "AddIlk":
		e := new(ILKREGISTRYAddIlk)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &AddIlk{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
		}
	case "Deny":
		e := new(ILKREGISTRYDeny)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	case "File":
		e := new(ILKREGISTRYFile)
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
		e := new(ILKREGISTRYFile0)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &File0{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
				What:  e.What[:],
				Data:  e.Data.Bytes(),
		}
	case "File1":
		e := new(ILKREGISTRYFile1)
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
		e := new(ILKREGISTRYFile2)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &File2{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
				What:  e.What[:],
		}
	case "NameError":
		e := new(ILKREGISTRYNameError)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &NameError{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
		}
	case "Rely":
		e := new(ILKREGISTRYRely)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Rely{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	case "RemoveIlk":
		e := new(ILKREGISTRYRemoveIlk)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &RemoveIlk{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
		}
	case "SymbolError":
		e := new(ILKREGISTRYSymbolError)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &SymbolError{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
		}
	case "UpdateIlk":
		e := new(ILKREGISTRYUpdateIlk)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &UpdateIlk{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
		}
	}
	return nil
}
