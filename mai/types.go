package mai

import (
	"github.com/nakji-network/connectors/mai/smart-contracts/mai"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var protos = []proto.Message{
	&mai.Approval{},
	&mai.OwnershipTransferred{},
	&mai.TransferVault{},
	&mai.DepositCollateral{},
	&mai.WithdrawCollateral{},
	&mai.BuyRiskyVault{},
	&mai.Transfer{},
	&mai.CreateVault{},
	&mai.DestroyVault{},
	&mai.BorrowToken{},
	&mai.PayBackToken{},
}

type ISmartContract interface {
	Message(eventName string, contractAbi *abi.ABI, vLog types.Log, timestamp *timestamppb.Timestamp) protoreflect.ProtoMessage
}

type Contract struct {
	ABI           *abi.ABI
	Name          string
	MessageParser ISmartContract
}

var ABIs = map[string]string{
	"mai": mai.MaiABI,
}

var Parsers = map[string]ISmartContract{
	"mai": &mai.SmartContract{},
}
