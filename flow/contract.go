package flow

import "google.golang.org/protobuf/proto"

type ISmartContract interface {
	Address() string
	Name() string
	Events() []string
	Message(Log) proto.Message
}
