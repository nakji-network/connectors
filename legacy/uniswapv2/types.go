package uniswapv2

import (
	reflect "reflect"

	"blep.ai/data/database"
	"google.golang.org/protobuf/proto"
)

var TopicTypes = map[string]proto.Message{
	"nakji.common.0_0_0.liquiditypool_reserve":   &LiquidityPoolReserve{},
	"nakji.common.0_0_0.liquiditypool_reserve-*": &LiquidityPoolReserve{},
	"nakji.common.0_0_0.liquiditypool_change":    &LiquidityPoolChange{},
}

func (l *LiquidityPoolReserve) Value() float64 {
	return l.Reserve1 / l.Reserve0
}

// TODO - Refactor this
func (lp *LiquidityPoolReserve) Viewer() map[reflect.Type]database.View {
	return map[reflect.Type]database.View{
		reflect.TypeOf(LiquidityPoolReserve{}): {
			Src: "timescaledb",
			Cmd: `
				time_bucket('%s', "ts")  AS "ts",
				LAST ("Reserve0", "ts") AS "reserve0",
				LAST ("Reserve1", "ts") AS "reserve1"
			`,
		},
	}
}
