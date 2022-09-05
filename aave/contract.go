package aave

import (
	"github.com/ethereum/go-ethereum/core/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ISmartContract interface {
	Address() string
	Message(vLog types.Log, ts *timestamppb.Timestamp) proto.Message
}
