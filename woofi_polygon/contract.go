package woofi_polygon

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"google.golang.org/protobuf/proto"
)

type Log struct {
	types.Log
	BlockTime       time.Time
	SenderAddress   *common.Address
	ReceiverAddress *common.Address
}

type ISmartContract interface {
	Address() string
	Events() []proto.Message
	Message(Log) proto.Message
}
