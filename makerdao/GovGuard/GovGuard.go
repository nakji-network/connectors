
package GovGuard

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
	case "LogDeny":
		e := new(GOVGUARDLogDeny)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LogDeny{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	case "LogRely":
		e := new(GOVGUARDLogRely)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LogRely{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	case "LogSetRoot":
		e := new(GOVGUARDLogSetRoot)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LogSetRoot{
				Ts:   timestamp,
				NewRoot:  e.NewRoot.Bytes(),
		}
	}
	return nil
}
