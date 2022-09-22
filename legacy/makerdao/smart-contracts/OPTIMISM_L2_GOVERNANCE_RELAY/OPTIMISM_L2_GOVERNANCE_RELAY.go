package OPTIMISM_L2_GOVERNANCE_RELAY

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
	case "Closed":
		e := new(OPTIMISML2GOVERNANCERELAYClosed)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Closed{
			Ts: timestamp,
		}
	case "Deny":
		e := new(OPTIMISML2GOVERNANCERELAYDeny)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
			Ts:  timestamp,
			Usr: e.Usr.Bytes(),
		}
	case "ERC20DepositInitiated":
		e := new(OPTIMISML2GOVERNANCERELAYERC20DepositInitiated)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &ERC20DepositInitiated{
			Ts:      timestamp,
			L1Token: e.L1Token.Bytes(),
			L2Token: e.L2Token.Bytes(),
			From:    e.From.Bytes(),
			To:      e.To.Bytes(),
			Amount:  e.Amount.Bytes(),
			Data:    e.Data[:],
		}
	case "ERC20WithdrawalFinalized":
		e := new(OPTIMISML2GOVERNANCERELAYERC20WithdrawalFinalized)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &ERC20WithdrawalFinalized{
			Ts:      timestamp,
			L1Token: e.L1Token.Bytes(),
			L2Token: e.L2Token.Bytes(),
			From:    e.From.Bytes(),
			To:      e.To.Bytes(),
			Amount:  e.Amount.Bytes(),
			Data:    e.Data[:],
		}
	case "Rely":
		e := new(OPTIMISML2GOVERNANCERELAYRely)
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
