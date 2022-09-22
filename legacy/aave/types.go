package aave

import (
	. "blep.ai/data/connectors/source/aave/lendingpool"
	"google.golang.org/protobuf/proto"
)

var TopicTypes = map[string]proto.Message{
	"nakji.aave.0_0_0.aave_lpborrow":              &LendingPoolBorrow{},
	"nakji.aave.0_0_0.aave_lpdeposit":             &LendingPoolDeposit{},
	"nakji.aave.0_0_0.aave_lpflashloan":           &LendingPoolFlashLoan{},
	"nakji.aave.0_0_0.aave_lpliquidationcall":     &LendingPoolLiquidationCall{},
	"nakji.aave.0_0_0.aave_lprebalance":           &LendingPoolRebalance{},
	"nakji.aave.0_0_0.aave_lprepay":               &LendingPoolRepay{},
	"nakji.aave.0_0_0.aave_lpreserveupdated":      &LendingPoolReserveUpdated{},
	"nakji.aave.0_0_0.aave_lpreserveascollateral": &LendingPoolReserveAsCollateral{},
	"nakji.aave.0_0_0.aave_lpswap":                &LendingPoolSwap{},
	"nakji.aave.0_0_0.aave_lpwithdraw":            &LendingPoolWithdraw{},
}

var TopicMap = map[string]string{
	"Borrow":                    ".fct.nakji.aave.0_0_0.aave_lpborrow",
	"Deposit":                   ".fct.nakji.aave.0_0_0.aave_lpdeposit",
	"FlashLoan":                 ".fct.nakji.aave.0_0_0.aave_lpflashloan",
	"LiquidationCall":           ".fct.nakji.aave.0_0_0.aave_lpliquidationcall",
	"RebalanceStableBorrowRate": ".fct.nakji.aave.0_0_0.aave_lprebalance",
	"Repay":                     ".fct.nakji.aave.0_0_0.aave_lprepay",
	"ReserveDataUpdated":        ".fct.nakji.aave.0_0_0.aave_lpreserveupdated",
	"ReserveUsedAsCollateral":   ".fct.nakji.aave.0_0_0.aave_lpreserveascollateral",
	"Swap":                      ".fct.nakji.aave.0_0_0.aave_lpswap",
	"Withdraw":                  ".fct.nakji.aave.0_0_0.aave_lpwithdraw",
}
