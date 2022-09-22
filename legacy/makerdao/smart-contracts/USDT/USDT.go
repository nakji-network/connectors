package USDT

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
	case "AddedBlackList":
		e := new(USDTAddedBlackList)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &AddedBlackList{
			Ts:   timestamp,
			User: e.User.Bytes(),
		}
	case "Approval":
		e := new(USDTApproval)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Approval{
			Ts:      timestamp,
			Owner:   e.Owner.Bytes(),
			Spender: e.Spender.Bytes(),
			Value:   e.Value.Bytes(),
		}
	case "Deprecate":
		e := new(USDTDeprecate)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deprecate{
			Ts:         timestamp,
			NewAddress: e.NewAddress.Bytes(),
		}
	case "DestroyedBlackFunds":
		e := new(USDTDestroyedBlackFunds)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &DestroyedBlackFunds{
			Ts:              timestamp,
			BlackListedUser: e.BlackListedUser.Bytes(),
			Balance:         e.Balance.Bytes(),
		}
	case "Issue":
		e := new(USDTIssue)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Issue{
			Ts:     timestamp,
			Amount: e.Amount.Bytes(),
		}
	case "Params":
		e := new(USDTParams)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Params{
			Ts:             timestamp,
			FeeBasisPoints: e.FeeBasisPoints.Bytes(),
			MaxFee:         e.MaxFee.Bytes(),
		}
	case "Pause":
		e := new(USDTPause)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Pause{
			Ts: timestamp,
		}
	case "Redeem":
		e := new(USDTRedeem)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Redeem{
			Ts:     timestamp,
			Amount: e.Amount.Bytes(),
		}
	case "RemovedBlackList":
		e := new(USDTRemovedBlackList)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &RemovedBlackList{
			Ts:   timestamp,
			User: e.User.Bytes(),
		}
	case "Transfer":
		e := new(USDTTransfer)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Transfer{
			Ts:    timestamp,
			From:  e.From.Bytes(),
			To:    e.To.Bytes(),
			Value: e.Value.Bytes(),
		}
	case "Unpause":
		e := new(USDTUnpause)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Unpause{
			Ts: timestamp,
		}
	}
	return nil
}
