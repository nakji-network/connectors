package woofi

import (
	"github.com/ethereum/go-ethereum/core/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ISmartContract interface {
	Network() string
	Address() string
	Message(vLog types.Log, ts *timestamppb.Timestamp) proto.Message
}
