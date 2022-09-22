package ETH

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
		e := new(ETHApproval)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Approval{
			Ts:  timestamp,
			Src: e.Src.Bytes(),
			Guy: e.Guy.Bytes(),
			Wad: e.Wad.Bytes(),
		}
	case "Deposit":
		e := new(ETHDeposit)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Deposit{
			Ts:  timestamp,
			Dst: e.Dst.Bytes(),
			Wad: e.Wad.Bytes(),
		}
	case "Transfer":
		e := new(ETHTransfer)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Transfer{
			Ts:  timestamp,
			Src: e.Src.Bytes(),
			Dst: e.Dst.Bytes(),
			Wad: e.Wad.Bytes(),
		}
	case "Withdrawal":
		e := new(ETHWithdrawal)
		if err := blepEthclient.UnpackLog(*contractAbi, e, eventName, vLog); err != nil {
			log.Error().Err(err).Msg("Failed to unpack log")
			return nil
		}

		return &Withdrawal{
			Ts:  timestamp,
			Src: e.Src.Bytes(),
			Wad: e.Wad.Bytes(),
		}
	}
	return nil
}
