
package MCDCropperImp

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
	case "Hope":
		e := new(MCDCROPPERIMPHope)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Hope{
				Ts:   timestamp,
				From:  e.From.Bytes(),
				To:  e.To.Bytes(),
		}
	case "NewProxy":
		e := new(MCDCROPPERIMPNewProxy)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &NewProxy{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
				Urp:  e.Urp.Bytes(),
		}
	case "Nope":
		e := new(MCDCROPPERIMPNope)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Nope{
				Ts:   timestamp,
				From:  e.From.Bytes(),
				To:  e.To.Bytes(),
		}
	}
	return nil
}
