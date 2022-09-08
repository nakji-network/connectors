package aave

import (
	"github.com/nakji-network/connectors/aave/lendingpool"
	"google.golang.org/protobuf/proto"
)

const (
	LendingPoolContractAddr = "0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9"
)

var TopicTypes = []proto.Message{
	&lendingpool.Borrow{},
	&lendingpool.Deposit{},
	&lendingpool.FlashLoan{},
	&lendingpool.LiquidationCall{},
	&lendingpool.RebalanceStableBorrowRate{},
	&lendingpool.Repay{},
	&lendingpool.ReserveDataUpdated{},
	&lendingpool.ReserveUsedAsCollateralDisabled{},
	&lendingpool.ReserveUsedAsCollateralEnabled{},
	&lendingpool.Swap{},
	&lendingpool.Withdraw{},
	&lendingpool.Paused{},
	&lendingpool.Unpaused{},
}
