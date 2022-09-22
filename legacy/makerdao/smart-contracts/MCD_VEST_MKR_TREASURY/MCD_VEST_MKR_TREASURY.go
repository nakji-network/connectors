package MCD_VEST_MKR_TREASURY

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
		e := new(MCDVESTMKRTREASURYDeny)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
			Ts:  timestamp,
			Usr: e.Usr.Bytes(),
		}
	case "File":
		e := new(MCDVESTMKRTREASURYFile)
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
		e := new(MCDVESTMKRTREASURYInit)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Init{
			Ts:  timestamp,
			Id:  e.Id.Bytes(),
			Usr: e.Usr.Bytes(),
		}
	case "Move":
		e := new(MCDVESTMKRTREASURYMove)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Move{
			Ts:  timestamp,
			Id:  e.Id.Bytes(),
			Dst: e.Dst.Bytes(),
		}
	case "Rely":
		e := new(MCDVESTMKRTREASURYRely)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Rely{
			Ts:  timestamp,
			Usr: e.Usr.Bytes(),
		}
	case "Restrict":
		e := new(MCDVESTMKRTREASURYRestrict)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Restrict{
			Ts: timestamp,
			Id: e.Id.Bytes(),
		}
	case "Unrestrict":
		e := new(MCDVESTMKRTREASURYUnrestrict)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Unrestrict{
			Ts: timestamp,
			Id: e.Id.Bytes(),
		}
	case "Vest":
		e := new(MCDVESTMKRTREASURYVest)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Vest{
			Ts:  timestamp,
			Id:  e.Id.Bytes(),
			Amt: e.Amt.Bytes(),
		}
	case "Yank":
		e := new(MCDVESTMKRTREASURYYank)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Yank{
			Ts:  timestamp,
			Id:  e.Id.Bytes(),
			End: e.End.Bytes(),
		}
	}
	return nil
}
