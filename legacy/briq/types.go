package briq

import "google.golang.org/protobuf/proto"

var TopicTypes = map[string]proto.Message{
	"nakji.briq.0_0_0.storage_transfersingle": &BriqEvent{},
	"nakji.briq.0_0_0.storage_mutate":         &BriqEvent{},
	"nakji.briq.0_0_0.storage_converttoft":    &BriqEvent{},
	"nakji.briq.0_0_0.storage_converttonft":   &BriqEvent{},
	"nakji.briq.0_0_0.storage_transfer":       &BriqEvent{},
	"nakji.briq.0_0_0.storage_uri":            &BriqEvent{},
}
