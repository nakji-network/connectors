
package MCD_PSM_USDC_A

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
	case "BuyGem":
		e := new(MCDPSMUSDCABuyGem)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &BuyGem{
				Ts:   timestamp,
				Owner:  e.Owner.Bytes(),
				Value:  e.Value.Bytes(),
				Fee:  e.Fee.Bytes(),
		}
	case "Deny":
		e := new(MCDPSMUSDCADeny)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deny{
				Ts:   timestamp,
				User:  e.User.Bytes(),
		}
	case "File":
		e := new(MCDPSMUSDCAFile)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &File{
				Ts:   timestamp,
				What:  e.What[:],
				Data:  e.Data.Bytes(),
		}
	case "Rely":
		e := new(MCDPSMUSDCARely)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Rely{
				Ts:   timestamp,
				User:  e.User.Bytes(),
		}
	case "SellGem":
		e := new(MCDPSMUSDCASellGem)
		if err := common.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &SellGem{
				Ts:   timestamp,
				Owner:  e.Owner.Bytes(),
				Value:  e.Value.Bytes(),
				Fee:  e.Fee.Bytes(),
		}
	}
	return nil
}
