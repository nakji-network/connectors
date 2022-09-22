package MCD_PSM_GUSD_A

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
	case "BuyGem":
		e := new(MCDPSMGUSDABuyGem)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &BuyGem{
			Ts:    timestamp,
			Owner: e.Owner.Bytes(),
			Value: e.Value.Bytes(),
			Fee:   e.Fee.Bytes(),
		}
	case "Deny":
		e := new(MCDPSMGUSDADeny)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
			Ts:  timestamp,
			Usr: e.Usr.Bytes(),
		}
	case "File":
		e := new(MCDPSMGUSDAFile)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &File{
			Ts:   timestamp,
			What: e.What[:],
			Data: e.Data.Bytes(),
		}
	case "Rely":
		e := new(MCDPSMGUSDARely)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Rely{
			Ts:  timestamp,
			Usr: e.Usr.Bytes(),
		}
	case "SellGem":
		e := new(MCDPSMGUSDASellGem)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &SellGem{
			Ts:    timestamp,
			Owner: e.Owner.Bytes(),
			Value: e.Value.Bytes(),
			Fee:   e.Fee.Bytes(),
		}
	}
	return nil
}
