package LERP_FAB

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
	case "Deny":
		e := new(LERPFABDeny)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
			Ts:  timestamp,
			Usr: e.Usr.Bytes(),
		}
	case "LerpFinished":
		e := new(LERPFABLerpFinished)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &LerpFinished{
			Ts:   timestamp,
			Lerp: e.Lerp.Bytes(),
		}
	case "NewIlkLerp":
		e := new(LERPFABNewIlkLerp)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &NewIlkLerp{
			Ts:        timestamp,
			Name:      e.Name[:],
			Target:    e.Target.Bytes(),
			Ilk:       e.Ilk[:],
			What:      e.What[:],
			StartTime: e.StartTime.Bytes(),
			Start:     e.Start.Bytes(),
			End:       e.End.Bytes(),
			Duration:  e.Duration.Bytes(),
		}
	case "NewLerp":
		e := new(LERPFABNewLerp)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &NewLerp{
			Ts:        timestamp,
			Name:      e.Name[:],
			Target:    e.Target.Bytes(),
			What:      e.What[:],
			StartTime: e.StartTime.Bytes(),
			Start:     e.Start.Bytes(),
			End:       e.End.Bytes(),
			Duration:  e.Duration.Bytes(),
		}
	case "Rely":
		e := new(LERPFABRely)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Rely{
			Ts:  timestamp,
			Usr: e.Usr.Bytes(),
		}
	}
	return nil
}
