
package MCD_ADM

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
	case "Etch":
		e := new(MCDADMEtch)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Etch{
				Ts:   timestamp,
				Slate:  e.Slate[:],
		}
	case "LogSetAuthority":
		e := new(MCDADMLogSetAuthority)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LogSetAuthority{
				Ts:   timestamp,
				Authority:  e.Authority.Bytes(),
		}
	case "LogSetOwner":
		e := new(MCDADMLogSetOwner)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LogSetOwner{
				Ts:   timestamp,
				Owner:  e.Owner.Bytes(),
		}
	}
	return nil
}
