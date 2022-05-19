package abracadabra

import (
	"blep.ai/data/connectors/source/abracadabra/magiccrv"
	"google.golang.org/protobuf/proto"
)

var TopicTypes = map[string]proto.Message{
	"nakji.abracadabra.0_0_0.magiccrv_transfer":       &magiccrv.Transfer{},
	"nakji.abracadabra.0_0_0.magiccrv_initializemint": &magiccrv.Approval{},
}
