package uniswapv2

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"

	"blep.ai/data/connectors/source/evm"
)

type LP struct {
	Reserves
	token0 *evm.ERC20
	token1 *evm.ERC20
}

// From the abi.go file
type Reserves struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}

func (l LP) Proto() LiquidityPoolReserve {
	dec0 := big.NewInt(int64(l.token0.Decimals))
	dec1 := big.NewInt(int64(l.token1.Decimals))
	exp0 := new(big.Int).Exp(big.NewInt(10), dec0, nil)
	exp1 := new(big.Int).Exp(big.NewInt(10), dec1, nil)

	r0 := new(big.Rat).SetFrac(l.Reserves.Reserve0, exp0)
	r1 := new(big.Rat).SetFrac(l.Reserves.Reserve1, exp1)

	// ignore exact because it definitely won't be exact
	r0f, _ := r0.Float64()
	r1f, _ := r1.Float64()
	//log.WithFields(log.Fields{
	//"r0": r0.FloatString(12),
	//"r0float": r0f,
	//}).Debug()
	//log.WithFields(log.Fields{
	//"r0": r1.FloatString(12),
	//"r0float": r1f,
	//}).Debug()

	return LiquidityPoolReserve{
		Ts:       &timestamp.Timestamp{Seconds: int64(l.BlockTimestampLast)},
		Reserve0: r0f,
		Reserve1: r1f,
	}
}

func (p LP) Rat() *big.Rat {
	dec0 := big.NewInt(int64(p.token0.Decimals))
	dec1 := big.NewInt(int64(p.token1.Decimals))
	exp0 := new(big.Int).Exp(big.NewInt(10), dec0, nil)
	exp1 := new(big.Int).Exp(big.NewInt(10), dec1, nil)

	num := new(big.Int).Mul(p.Reserves.Reserve0, exp1)
	den := new(big.Int).Mul(p.Reserves.Reserve1, exp0)

	return new(big.Rat).SetFrac(num, den)
}

func (p LP) InvRat() *big.Rat {
	return new(big.Rat).Inv(p.Rat())
}

// Deprecated. Using Pair ID as ticker
func (l LP) Ticker(delimiter string) string {
	return common.Bytes2Hex(l.token0.Id) + delimiter + common.Bytes2Hex(l.token1.Id)
}

// TODO: make LiquidityPoolEvents viewable from /data endpoint
//func (x *LiquidityPoolEvent) Viewer() map[reflect.Type]database.View{
//return map[reflect.Type]database.View{
//reflect.TypeOf(LiquidityPoolEvent{}): {
//Src: "timescaledb",
//Cmd: `
//time_bucket('%s', "ts")  AS "t",
//FIRST("AssetType", "ts") AS "asset_type",
//(SUM("Amount0In")+SUM("Amount0Out"))/(SUM("Amount1In")+SUM("Amount1Out")) AS "price",
//SUM ("Amount0In")               AS "amount0in",
//SUM ("Amount0Out")              AS "amount0out"
//SUM ("Amount1In")               AS "amount1in",
//SUM ("Amount1Out")              AS "amount1out"
//`,
//},
//reflect.TypeOf(Kline{}): {
//Src: "timescaledb",
//Cmd: `
//time_bucket('%s', "ts")        AS "t",
//FIRST("Price", "ts")           AS "open",
//MAX  ("Price")                        AS "high",
//MIN  ("Price")                        AS "low",
//LAST ("Price", "ts")           AS "close",
//SUM("Amount0Out")+SUM("Amount0In")    AS "volume",
//SUM ("Amount0Out")                    AS "buy_volume",
//SUM ("Amount0In")                     AS "sell_volume"
//`,
//},
//}
//}
