package calc

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/zlib"
)

var ops = []zc.OpDecl{
	zc.GenOp("abs",
		zc.Func(zlib.AbsBigInt, zc.BigInt),
		zc.Func(zlib.AbsDecimal, zc.Decimal),
	),
	zc.GenOp("add",
		zc.Func(zlib.AddBigInt, zc.BigInt, zc.BigInt),
		zc.Func(zlib.AddDecimal, zc.Decimal, zc.Decimal),
	),
	zc.GenOp("and",
		zc.Func(zlib.AndBigInt, zc.BigInt),
	),
	zc.GenOp("ceil",
		zc.Func(zc.NoOp, zc.BigInt),
		zc.Func(zlib.CeilDecimal, zc.Decimal),
	),
	zc.GenOp("div",
		zc.Func(zlib.DivDecimal, zc.Decimal),
	),
	zc.GenOp("eq",
		zc.Func(zlib.EqBigInt, zc.BigInt),
		zc.Func(zlib.EqDecimal, zc.Decimal),
	),
	zc.GenOp("floor",
		zc.Func(zc.NoOp, zc.BigInt),
		zc.Func(zlib.FloorDecimal, zc.Decimal),
	),
	zc.GenOp("gt",
		zc.Func(zlib.GtBigInt, zc.BigInt, zc.BigInt),
		zc.Func(zlib.GtDecimal, zc.Decimal, zc.Decimal),
	),
	zc.GenOp("gte",
		zc.Func(zlib.GteBigInt, zc.BigInt, zc.BigInt),
		zc.Func(zlib.GteDecimal, zc.Decimal, zc.Decimal),
	),
	zc.Op("hex", zlib.HexBigInt, zc.BigInt),
	zc.GenOp("lt",
		zc.Func(zlib.LtBigInt, zc.BigInt, zc.BigInt),
		zc.Func(zlib.LtDecimal, zc.Decimal, zc.Decimal),
	),
	zc.GenOp("lte",
		zc.Func(zlib.LteBigInt, zc.BigInt, zc.BigInt),
		zc.Func(zlib.LteDecimal, zc.Decimal, zc.Decimal),
	),
	zc.GenOp("mod",
		zc.Func(zlib.ModBigInt, zc.BigInt, zc.BigInt),
		zc.Func(zlib.ModDecimal, zc.Decimal, zc.Decimal),
	),
	zc.GenOp("sub",
		zc.Func(zlib.SubBigInt, zc.BigInt, zc.BigInt),
		zc.Func(zlib.SubDecimal, zc.Decimal, zc.Decimal),
	),
}

var opsTable map[string]zc.CalcFunc

func init() {
	opsTable = make(map[string]zc.CalcFunc)
	for _, v := range ops {
		k := v.Name
		if _, exists := opsTable[k]; exists {
			panic(zc.ErrDuplicateOp(k))
		}
		opsTable[k] = zc.EvalOp(v)
	}
}
