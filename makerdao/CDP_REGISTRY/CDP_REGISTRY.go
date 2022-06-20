
package CDP_REGISTRY

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
	if eventName == "NewCdpRegistered"{
		e := new(CDPREGISTRYNewCdpRegistered)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &NewCdpRegistered{
				Ts:   timestamp,
				Sender:  e.Sender.Bytes(),
				Owner:  e.Owner.Bytes(),
				Cdp:  e.Cdp.Bytes(),
		}
	}
	return nil
}
