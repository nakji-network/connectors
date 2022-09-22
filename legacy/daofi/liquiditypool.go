package daofi

import (
	"math/big"

	. "blep.ai/data/connectors/source/evm"
	"blep.ai/data/connectors/source/uniswapv2"
	"github.com/ethereum/go-ethereum/common"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

type LP struct {
	Reserves
	token0 *ERC20
	token1 *ERC20
}

// From the abi.go file
type Reserves struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}

func (l LP) Proto() uniswapv2.LiquidityPoolReserve {
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

	return uniswapv2.LiquidityPoolReserve{
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
