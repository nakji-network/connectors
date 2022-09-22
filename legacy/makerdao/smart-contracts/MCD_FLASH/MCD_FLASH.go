package MCD_FLASH

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
		e := new(MCDFLASHDeny)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
			Ts:  timestamp,
			Usr: e.Usr.Bytes(),
		}
	case "File":
		e := new(MCDFLASHFile)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &File{
			Ts:   timestamp,
			What: e.What[:],
			Data: e.Data.Bytes(),
		}
	case "FlashLoan":
		e := new(MCDFLASHFlashLoan)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &FlashLoan{
			Ts:       timestamp,
			Receiver: e.Receiver.Bytes(),
			Token:    e.Token.Bytes(),
			Amount:   e.Amount.Bytes(),
			Fee:      e.Fee.Bytes(),
		}
	case "Rely":
		e := new(MCDFLASHRely)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Rely{
			Ts:  timestamp,
			Usr: e.Usr.Bytes(),
		}
	case "VatDaiFlashLoan":
		e := new(MCDFLASHVatDaiFlashLoan)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &VatDaiFlashLoan{
			Ts:       timestamp,
			Receiver: e.Receiver.Bytes(),
			Amount:   e.Amount.Bytes(),
			Fee:      e.Fee.Bytes(),
		}
	}
	return nil
}
