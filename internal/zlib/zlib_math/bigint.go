package zlib_math

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/number"
	"github.com/blackchip-org/zc/types"
)

func op2BigInt(fn func(number.BigInt, number.BigInt) number.BigInt) zc.CalcFunc {
	return func(_ zc.Env, args []types.Value) ([]types.Value, error) {
		x := types.BigInt.Native(args[0])
		y := types.BigInt.Native(args[1])
		z := fn(x, y)
		return []types.Value{types.BigInt.Value(z)}, nil
	}
}

func InitBigInt(e zc.Env) {
	//e.Calc().RegisterGenOp(zc.OpAdd, addBigInt, types.BigInt, types.BigInt)
}

func addBigIntFn(x number.BigInt, y number.BigInt) number.BigInt { return x.Add(y) }

var (
	addBigInt = op2BigInt(addBigIntFn)
)
