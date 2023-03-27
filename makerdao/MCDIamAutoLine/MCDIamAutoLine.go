
package MCDIamAutoLine

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
		e := new(MCDIAMAUTOLINEDeny)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	case "Exec":
		e := new(MCDIAMAUTOLINEExec)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Exec{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
				Line:  e.Line.Bytes(),
				LineNew:  e.LineNew.Bytes(),
		}
	case "Rely":
		e := new(MCDIAMAUTOLINERely)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Rely{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	case "Remove":
		e := new(MCDIAMAUTOLINERemove)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Remove{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
		}
	case "Setup":
		e := new(MCDIAMAUTOLINESetup)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Setup{
				Ts:   timestamp,
				Ilk:  e.Ilk[:],
				Line:  e.Line.Bytes(),
				Gap:  e.Gap.Bytes(),
				Ttl:  e.Ttl.Bytes(),
		}
	}
	return nil
}
