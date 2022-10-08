package woofi

import (
	"github.com/nakji-network/connectors/woofi/bscWOOPP"
	"github.com/nakji-network/connectors/woofi/polygonWOOPP"
	"google.golang.org/protobuf/proto"
)

var TopicTypes = map[string]map[string]proto.Message{
	"bsc": {
		"nakji.woofi.0_0_0.bscwoopp_feemanagerupdated":         &bscWOOPP.FeeManagerUpdated{},
		"nakji.woofi.0_0_0.bscwoopp_ownershiptransferprepared": &bscWOOPP.OwnershipTransferPrepared{},
		"nakji.woofi.0_0_0.bscwoopp_ownershiptransferred":      &bscWOOPP.OwnershipTransferred{},
		"nakji.woofi.0_0_0.bscwoopp_parametersupdated":         &bscWOOPP.ParametersUpdated{},
		"nakji.woofi.0_0_0.bscwoopp_paused":                    &bscWOOPP.Paused{},
		"nakji.woofi.0_0_0.bscwoopp_rewardmanagerupdated":      &bscWOOPP.RewardManagerUpdated{},
		"nakji.woofi.0_0_0.bscwoopp_strategistupdated":         &bscWOOPP.StrategistUpdated{},
		"nakji.woofi.0_0_0.bscwoopp_unpaused":                  &bscWOOPP.Unpaused{},
		"nakji.woofi.0_0_0.bscwoopp_withdraw":                  &bscWOOPP.Withdraw{},
		"nakji.woofi.0_0_0.bscwoopp_wooguardianupdated":        &bscWOOPP.WooGuardianUpdated{},
		"nakji.woofi.0_0_0.bscwoopp_wooracleupdated":           &bscWOOPP.WooracleUpdated{},
		"nakji.woofi.0_0_0.bscwoopp_wooswap":                   &bscWOOPP.WooSwap{},
	},
	"polygon": {
		"nakji.woofi.0_0_0.polygonwoopp_feemanagerupdated":         &polygonWOOPP.FeeManagerUpdated{},
		"nakji.woofi.0_0_0.polygonwoopp_ownershiptransferprepared": &polygonWOOPP.OwnershipTransferPrepared{},
		"nakji.woofi.0_0_0.polygonwoopp_ownershiptransferred":      &polygonWOOPP.OwnershipTransferred{},
		"nakji.woofi.0_0_0.polygonwoopp_parametersupdated":         &polygonWOOPP.ParametersUpdated{},
		"nakji.woofi.0_0_0.polygonwoopp_paused":                    &polygonWOOPP.Paused{},
		"nakji.woofi.0_0_0.polygonwoopp_rewardmanagerupdated":      &polygonWOOPP.RewardManagerUpdated{},
		"nakji.woofi.0_0_0.polygonwoopp_strategistupdated":         &polygonWOOPP.StrategistUpdated{},
		"nakji.woofi.0_0_0.polygonwoopp_unpaused":                  &polygonWOOPP.Unpaused{},
		"nakji.woofi.0_0_0.polygonwoopp_withdraw":                  &polygonWOOPP.Withdraw{},
		"nakji.woofi.0_0_0.polygonwoopp_wooguardianupdated":        &polygonWOOPP.WooGuardianUpdated{},
		"nakji.woofi.0_0_0.polygonwoopp_wooracleupdated":           &polygonWOOPP.WooracleUpdated{},
		"nakji.woofi.0_0_0.polygonwoopp_wooswap":                   &polygonWOOPP.WooSwap{},
	},
}

var ABIs = map[string]string{
	"bscWOOPP":     bscWOOPP.BscWOOPPABI,
	"polygonWOOPP": polygonWOOPP.PolygonWOOPPABI,
}
