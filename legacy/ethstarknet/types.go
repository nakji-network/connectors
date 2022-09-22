package starknet

import (
	"blep.ai/data/connectors/source/ethstarknet/smart-contracts/Starknet"

	"google.golang.org/protobuf/proto"
)

var TopicTypes = map[string]proto.Message{
	"nakji.starknet.0_0_0.starknet_consumedmessagetol1":    &Starknet.ConsumedMessageToL1{},
	"nakji.starknet.0_0_0.starknet_consumedmessagetol2":    &Starknet.ConsumedMessageToL2{},
	"nakji.starknet.0_0_0.starknet_logmessagetol1":         &Starknet.LogMessageToL1{},
	"nakji.starknet.0_0_0.starknet_logmessagetol2":         &Starknet.LogMessageToL2{},
	"nakji.starknet.0_0_0.starknet_lognewgovernoraccepted": &Starknet.LogNewGovernorAccepted{},
	"nakji.starknet.0_0_0.starknet_lognominatedgovernor":   &Starknet.LogNominatedGovernor{},
	"nakji.starknet.0_0_0.starknet_lognominationcancelled": &Starknet.LogNominationCancelled{},
	"nakji.starknet.0_0_0.starknet_logoperatoradded":       &Starknet.LogOperatorAdded{},
	"nakji.starknet.0_0_0.starknet_logoperatorremoved":     &Starknet.LogOperatorRemoved{},
	"nakji.starknet.0_0_0.starknet_logremovedgovernor":     &Starknet.LogRemovedGovernor{},
	"nakji.starknet.0_0_0.starknet_logstatetransitionfact": &Starknet.LogStateTransitionFact{},
	"nakji.starknet.0_0_0.starknet_logstateupdate":         &Starknet.LogStateUpdate{},
}

var ABIs = map[string]string{
	"Starknet": Starknet.StarknetABI,
}
