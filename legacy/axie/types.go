// Code generated by connectorgen - Edit as necessary.
package axie

import (
	"blep.ai/data/connectors/source/axie/axienft"
	"blep.ai/data/connectors/source/axie/axs"
	"blep.ai/data/connectors/source/axie/bridge"
	"blep.ai/data/connectors/source/axie/slp"
	"google.golang.org/protobuf/proto"
)

var TopicTypes = map[string]proto.Message{
	"nakji.axie.0_0_0.axienft_approval":       &axienft.Approval{},
	"nakji.axie.0_0_0.axienft_approvalforall": &axienft.ApprovalForAll{},
	"nakji.axie.0_0_0.axienft_axieevolved":    &axienft.AxieEvolved{},
	"nakji.axie.0_0_0.axienft_axierebirthed":  &axienft.AxieRebirthed{},
	"nakji.axie.0_0_0.axienft_axieretired":    &axienft.AxieRetired{},
	"nakji.axie.0_0_0.axienft_axiespawned":    &axienft.AxieSpawned{},
	"nakji.axie.0_0_0.axienft_transfer":       &axienft.Transfer{},
	"nakji.axie.0_0_0.axs_approval":           &axs.Approval{},
	"nakji.axie.0_0_0.axs_transfer":           &axs.Transfer{},
	"nakji.axie.0_0_0.bridge_adminchanged":    &bridge.AdminChanged{},
	"nakji.axie.0_0_0.bridge_adminremoved":    &bridge.AdminRemoved{},
	"nakji.axie.0_0_0.bridge_paused":          &bridge.Paused{},
	"nakji.axie.0_0_0.bridge_proxyupdated":    &bridge.ProxyUpdated{},
	"nakji.axie.0_0_0.bridge_tokendeposited":  &bridge.TokenDeposited{},
	"nakji.axie.0_0_0.bridge_tokenwithdrew":   &bridge.TokenWithdrew{},
	"nakji.axie.0_0_0.bridge_unpaused":        &bridge.Unpaused{},
	"nakji.axie.0_0_0.slp_adminchanged":       &slp.AdminChanged{},
	"nakji.axie.0_0_0.slp_adminremoved":       &slp.AdminRemoved{},
	"nakji.axie.0_0_0.slp_approval":           &slp.Approval{},
	"nakji.axie.0_0_0.slp_minteradded":        &slp.MinterAdded{},
	"nakji.axie.0_0_0.slp_minterremoved":      &slp.MinterRemoved{},
	"nakji.axie.0_0_0.slp_transfer":           &slp.Transfer{},
}
