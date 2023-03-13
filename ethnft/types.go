package ethnft

import (
	"github.com/nakji-network/connectors/ethnft/smart-contracts/erc1155"
	"github.com/nakji-network/connectors/ethnft/smart-contracts/erc721"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	EventApproval = "Approval"
	EventTransfer = "Transfer"
)

var protos = []proto.Message{
	&erc721.ApprovalForAll{},
	&erc721.Approval{},
	&erc721.Transfer{},
	&erc1155.ApprovalForAll{},
	&erc1155.TransferBatch{},
	&erc1155.TransferSingle{},
	&erc1155.URI{},
}

type ISmartContract interface {
	Message(eventName string, contractAbi *abi.ABI, vLog types.Log, timestamp *timestamppb.Timestamp) protoreflect.ProtoMessage
}

type Contract struct {
	ABI           *abi.ABI
	Name          string
	MessageParser ISmartContract
}
