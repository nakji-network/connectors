package woofi

import (
	"github.com/nakji-network/connectors/woofi/WOOPP"
	"google.golang.org/protobuf/proto"
)

var TopicTypes = map[string]proto.Message{
	"nakji.woofi.0_0_0.woopp_feemanagerupdated":         &WOOPP.FeeManagerUpdated{},
	"nakji.woofi.0_0_0.woopp_ownershiptransferprepared": &WOOPP.OwnershipTransferPrepared{},
	"nakji.woofi.0_0_0.woopp_ownershiptransferred":      &WOOPP.OwnershipTransferred{},
	"nakji.woofi.0_0_0.woopp_parametersupdated":         &WOOPP.ParametersUpdated{},
	"nakji.woofi.0_0_0.woopp_paused":                    &WOOPP.Paused{},
	"nakji.woofi.0_0_0.woopp_rewardmanagerupdated":      &WOOPP.RewardManagerUpdated{},
	"nakji.woofi.0_0_0.woopp_strategistupdated":         &WOOPP.StrategistUpdated{},
	"nakji.woofi.0_0_0.woopp_unpaused":                  &WOOPP.Unpaused{},
	"nakji.woofi.0_0_0.woopp_withdraw":                  &WOOPP.Withdraw{},
	"nakji.woofi.0_0_0.woopp_wooguardianupdated":        &WOOPP.WooGuardianUpdated{},
	"nakji.woofi.0_0_0.woopp_wooracleupdated":           &WOOPP.WooracleUpdated{},
	"nakji.woofi.0_0_0.woopp_wooswap":                   &WOOPP.WooSwap{},
	
}

var ABIs = map[string]string{
	"WOOPP":                          WOOPP.WOOPPABI,
}