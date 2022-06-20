
package OPTIMISM_DAI_BRIDGE

import (
	"github.com/nakji-network/connector/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SmartContract struct{}

func (sc *SmartContract) Message(eventName string, contractAbi *abi.ABI, vLog types.Log, timestamp *timestamppb.Timestamp) protoreflect.ProtoMessage {
	switch eventName {
	case "Closed":
		e := new(OPTIMISMDAIBRIDGEClosed)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Closed{
				Ts:   timestamp,
		}
	case "Deny":
		e := new(OPTIMISMDAIBRIDGEDeny)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	case "ERC20DepositInitiated":
		e := new(OPTIMISMDAIBRIDGEERC20DepositInitiated)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &ERC20DepositInitiated{
				Ts:   timestamp,
				L1Token:  e.L1Token.Bytes(),
				L2Token:  e.L2Token.Bytes(),
				From:  e.From.Bytes(),
				To:  e.To.Bytes(),
				Amount:  e.Amount.Bytes(),
				Data:  e.Data[:],
		}
	case "ERC20WithdrawalFinalized":
		e := new(OPTIMISMDAIBRIDGEERC20WithdrawalFinalized)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &ERC20WithdrawalFinalized{
				Ts:   timestamp,
				L1Token:  e.L1Token.Bytes(),
				L2Token:  e.L2Token.Bytes(),
				From:  e.From.Bytes(),
				To:  e.To.Bytes(),
				Amount:  e.Amount.Bytes(),
				Data:  e.Data[:],
		}
	case "Rely":
		e := new(OPTIMISMDAIBRIDGERely)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Rely{
				Ts:   timestamp,
				Usr:  e.Usr.Bytes(),
		}
	}
	return nil
}
