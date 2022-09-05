package aave

import (
	"github.com/nakji-network/connectors/aave/lendingpool"
	"google.golang.org/protobuf/proto"
)

const (
	LendingPoolContractAddr = "0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9"
)

var TopicTypes = map[string]proto.Message{
	"nakji.aave.0_0_0.aave_lpborrow":              &lendingpool.LendingPoolBorrow{},
	"nakji.aave.0_0_0.aave_lpdeposit":             &lendingpool.LendingPoolDeposit{},
	"nakji.aave.0_0_0.aave_lpflashloan":           &lendingpool.LendingPoolFlashLoan{},
	"nakji.aave.0_0_0.aave_lpliquidationcall":     &lendingpool.LendingPoolLiquidationCall{},
	"nakji.aave.0_0_0.aave_lprebalance":           &lendingpool.LendingPoolRebalance{},
	"nakji.aave.0_0_0.aave_lprepay":               &lendingpool.LendingPoolRepay{},
	"nakji.aave.0_0_0.aave_lpreserveupdated":      &lendingpool.LendingPoolReserveUpdated{},
	"nakji.aave.0_0_0.aave_lpreserveascollateral": &lendingpool.LendingPoolReserveAsCollateral{},
	"nakji.aave.0_0_0.aave_lpswap":                &lendingpool.LendingPoolSwap{},
	"nakji.aave.0_0_0.aave_lpwithdraw":            &lendingpool.LendingPoolWithdraw{},
}
