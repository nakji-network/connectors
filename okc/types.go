package okc

import (
	"github.com/nakji-network/connectors/okc/smart-contracts/cOKC"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var protos = []proto.Message{
	&cOKC.Approval{},
	&cOKC.MintClaimed{},
	&cOKC.RankClaimed{},
	&cOKC.Staked{},
	&cOKC.Transfer{},
	&cOKC.Withdrawn{},
}

type ISmartContract interface {
	Message(eventName string, contractAbi *abi.ABI, vLog types.Log, timestamp *timestamppb.Timestamp) protoreflect.ProtoMessage
}

type Contract struct {
	ABI           *abi.ABI
	Name          string
	MessageParser ISmartContract
}

var ContractAddresses = map[string]string{
	"cOKC": "0x1cC4D981e897A3D2E7785093A648c0a75fAd0453",
}

var ABIs = map[string]string{
	"cOKC": cOKC.COKCABI,
}
