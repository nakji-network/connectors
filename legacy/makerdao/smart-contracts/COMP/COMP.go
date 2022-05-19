package COMP

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
	case "Approval":
		e := new(COMPApproval)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Approval{
			Ts:      timestamp,
			Owner:   e.Owner.Bytes(),
			Spender: e.Spender.Bytes(),
			Amount:  e.Amount.Bytes(),
		}
	case "DelegateChanged":
		e := new(COMPDelegateChanged)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &DelegateChanged{
			Ts:           timestamp,
			Delegator:    e.Delegator.Bytes(),
			FromDelegate: e.FromDelegate.Bytes(),
			ToDelegate:   e.ToDelegate.Bytes(),
		}
	case "DelegateVotesChanged":
		e := new(COMPDelegateVotesChanged)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &DelegateVotesChanged{
			Ts:              timestamp,
			Delegate:        e.Delegate.Bytes(),
			PreviousBalance: e.PreviousBalance.Bytes(),
			NewBalance:      e.NewBalance.Bytes(),
		}
	case "Transfer":
		e := new(COMPTransfer)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Transfer{
			Ts:     timestamp,
			From:   e.From.Bytes(),
			To:     e.To.Bytes(),
			Amount: e.Amount.Bytes(),
		}
	}
	return nil
}
