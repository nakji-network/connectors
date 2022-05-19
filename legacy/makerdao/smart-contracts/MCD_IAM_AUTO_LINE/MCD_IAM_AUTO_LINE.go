package MCD_IAM_AUTO_LINE

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
		e := new(MCDIAMAUTOLINEDeny)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
			Ts:  timestamp,
			Usr: e.Usr.Bytes(),
		}
	case "Exec":
		e := new(MCDIAMAUTOLINEExec)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Exec{
			Ts:      timestamp,
			Ilk:     e.Ilk[:],
			Line:    e.Line.Bytes(),
			LineNew: e.LineNew.Bytes(),
		}
	case "Rely":
		e := new(MCDIAMAUTOLINERely)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Rely{
			Ts:  timestamp,
			Usr: e.Usr.Bytes(),
		}
	case "Remove":
		e := new(MCDIAMAUTOLINERemove)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Remove{
			Ts:  timestamp,
			Ilk: e.Ilk[:],
		}
	case "Setup":
		e := new(MCDIAMAUTOLINESetup)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Setup{
			Ts:   timestamp,
			Ilk:  e.Ilk[:],
			Line: e.Line.Bytes(),
			Gap:  e.Gap.Bytes(),
			Ttl:  e.Ttl.Bytes(),
		}
	}
	return nil
}
