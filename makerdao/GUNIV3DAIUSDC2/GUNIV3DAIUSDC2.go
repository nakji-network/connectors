package GUniV3DaiUSDC2

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
	case "ProxyAdminTransferred":
		e := new(GUNIV3DAIUSDC2ProxyAdminTransferred)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &ProxyAdminTransferred{
			Ts:            timestamp,
			PreviousAdmin: e.PreviousAdmin.Bytes(),
			NewAdmin:      e.NewAdmin.Bytes(),
		}
	case "ProxyImplementationUpdated":
		e := new(GUNIV3DAIUSDC2ProxyImplementationUpdated)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &ProxyImplementationUpdated{
			Ts:                     timestamp,
			PreviousImplementation: e.PreviousImplementation.Bytes(),
			NewImplementation:      e.NewImplementation.Bytes(),
		}
	}
	return nil
}
