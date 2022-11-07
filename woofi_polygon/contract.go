package woofi_polygon

import (
	"github.com/ethereum/go-ethereum/core/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ISmartContract interface {
	Address() string
	Events() []proto.Message
	Message(vLog types.Log, timestamp *timestamppb.Timestamp) proto.Message
}
