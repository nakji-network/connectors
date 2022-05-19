package MIP21_LIQUIDATION_ORACLE

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
	case "Bump":
		e := new(MIP21LIQUIDATIONORACLEBump)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Bump{
			Ts:  timestamp,
			Ilk: e.Ilk[:],
			Val: e.Val.Bytes(),
		}
	case "Cull":
		e := new(MIP21LIQUIDATIONORACLECull)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Cull{
			Ts:  timestamp,
			Ilk: e.Ilk[:],
			Urn: e.Urn.Bytes(),
		}
	case "Cure":
		e := new(MIP21LIQUIDATIONORACLECure)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Cure{
			Ts:  timestamp,
			Ilk: e.Ilk[:],
		}
	case "Deny":
		e := new(MIP21LIQUIDATIONORACLEDeny)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
			Ts:  timestamp,
			Usr: e.Usr.Bytes(),
		}
	case "File":
		e := new(MIP21LIQUIDATIONORACLEFile)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &File{
			Ts:   timestamp,
			What: e.What[:],
			Data: e.Data.Bytes(),
		}
	case "Init":
		e := new(MIP21LIQUIDATIONORACLEInit)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Init{
			Ts:  timestamp,
			Ilk: e.Ilk[:],
			Val: e.Val.Bytes(),
			Tau: e.Tau.Bytes(),
		}
	case "Rely":
		e := new(MIP21LIQUIDATIONORACLERely)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Rely{
			Ts:  timestamp,
			Usr: e.Usr.Bytes(),
		}
	case "Tell":
		e := new(MIP21LIQUIDATIONORACLETell)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Tell{
			Ts:  timestamp,
			Ilk: e.Ilk[:],
		}
	}
	return nil
}
