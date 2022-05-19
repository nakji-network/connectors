package GUNIV3DAIUSDC2

import (
	blepEthclient "blep.ai/data/chain/ethereum/ethclient"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

type SmartContract struct{}

func (sc *SmartContract) Message(eventName string, contractAbi *abi.ABI, vLog types.Log, timestamp *timestamp.Timestamp) protoreflect.ProtoMessage {
	switch eventName {
	case "ProxyAdminTransferred":
		e := new(GUNIV3DAIUSDC2ProxyAdminTransferred)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
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
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
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
