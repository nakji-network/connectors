
package TUSD

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
	case "NewPendingOwner":
		e := new(TUSDNewPendingOwner)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &NewPendingOwner{
				Ts:   timestamp,
				CurrentOwner:  e.CurrentOwner.Bytes(),
				PendingOwner:  e.PendingOwner.Bytes(),
		}
	case "ProxyOwnershipTransferred":
		e := new(TUSDProxyOwnershipTransferred)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &ProxyOwnershipTransferred{
				Ts:   timestamp,
				PreviousOwner:  e.PreviousOwner.Bytes(),
				NewOwner:  e.NewOwner.Bytes(),
		}
	case "Upgraded":
		e := new(TUSDUpgraded)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Upgraded{
				Ts:   timestamp,
				Implementation:  e.Implementation.Bytes(),
		}
	}
	return nil
}
