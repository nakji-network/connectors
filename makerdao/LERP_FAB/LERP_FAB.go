
package LERP_FAB

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
		e := new(LERPFABDeny)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	case "LerpFinished":
		e := new(LERPFABLerpFinished)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LerpFinished{
				Ts:   timestamp,
				Lerp:  e.Lerp.Bytes(),
		}
	case "NewIlkLerp":
		e := new(LERPFABNewIlkLerp)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &NewIlkLerp{
				Ts:   timestamp,
				Name:  e.Name[:],
				Target:  e.Target.Bytes(),
				Ilk:  e.Ilk[:],
				What:  e.What[:],
				StartTime:  e.StartTime.Bytes(),
				Start:  e.Start.Bytes(),
				End:  e.End.Bytes(),
				Duration:  e.Duration.Bytes(),
		}
	case "NewLerp":
		e := new(LERPFABNewLerp)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &NewLerp{
				Ts:   timestamp,
				Name:  e.Name[:],
				Target:  e.Target.Bytes(),
				What:  e.What[:],
				StartTime:  e.StartTime.Bytes(),
				Start:  e.Start.Bytes(),
				End:  e.End.Bytes(),
				Duration:  e.Duration.Bytes(),
		}
	case "Rely":
		e := new(LERPFABRely)
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
