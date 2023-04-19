package calc

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/zlib"
)

var ops = []zc.OpDecl{
	zc.GenOp("abs",
		zc.Func(zlib.AbsBigInt, zc.BigInt),
		zc.Func(zlib.AbsDecimal, zc.Decimal),
		zc.Func(zlib.AbsFloat, zc.Float),
	),
	zc.GenOp("add",
		zc.Func(zlib.AddBigInt, zc.BigInt, zc.BigInt),
		zc.Func(zlib.AddDecimal, zc.Decimal, zc.Decimal),
		zc.Func(zlib.AddFloat, zc.Float, zc.Float),
	),
	zc.GenOp("and",
		zc.Func(zlib.AndBigInt, zc.BigInt),
	),
	zc.GenOp("ceil",
		zc.Func(zc.NoOp, zc.BigInt),
		zc.Func(zlib.CeilDecimal, zc.Decimal),
		zc.Func(zlib.CeilFloat, zc.Float),
	),
	zc.GenOp("div",
		zc.Func(zlib.DivDecimal, zc.Decimal),
		zc.Func(zlib.DivFloat, zc.Float, zc.Float),
	),
	zc.GenOp("eq",
		zc.Func(zlib.EqBigInt, zc.BigInt),
		zc.Func(zlib.EqDecimal, zc.Decimal),
		zc.Func(zlib.EqFloat, zc.Float, zc.Float),
	),
	zc.GenOp("floor",
		zc.Func(zc.NoOp, zc.BigInt),
		zc.Func(zlib.FloorDecimal, zc.Decimal),
		zc.Func(zlib.FloorFloat, zc.Float),
	),
	zc.GenOp("gt",
		zc.Func(zlib.GtBigInt, zc.BigInt, zc.BigInt),
		zc.Func(zlib.GtDecimal, zc.Decimal, zc.Decimal),
		zc.Func(zlib.GtFloat, zc.Float, zc.Float),
	),
	zc.GenOp("gte",
		zc.Func(zlib.GteBigInt, zc.BigInt, zc.BigInt),
		zc.Func(zlib.GteDecimal, zc.Decimal, zc.Decimal),
		zc.Func(zlib.GteFloat, zc.Float, zc.Float),
	),
	zc.Op("hex", zlib.HexBigInt, zc.BigInt),
	zc.GenOp("lt",
		zc.Func(zlib.LtBigInt, zc.BigInt, zc.BigInt),
		zc.Func(zlib.LtDecimal, zc.Decimal, zc.Decimal),
		zc.Func(zlib.LtFloat, zc.Float, zc.Float),
	),
	zc.GenOp("lte",
		zc.Func(zlib.LteBigInt, zc.BigInt, zc.BigInt),
		zc.Func(zlib.LteDecimal, zc.Decimal, zc.Decimal),
		zc.Func(zlib.LteFloat, zc.Float, zc.Float),
	),
	zc.GenOp("mod",
		zc.Func(zlib.ModBigInt, zc.BigInt, zc.BigInt),
		zc.Func(zlib.ModDecimal, zc.Decimal, zc.Decimal),
		zc.Func(zlib.ModFloat, zc.Float, zc.Float),
	),
	zc.GenOp("sub",
		zc.Func(zlib.SubBigInt, zc.BigInt, zc.BigInt),
		zc.Func(zlib.SubDecimal, zc.Decimal, zc.Decimal),
		zc.Func(zlib.SubFloat, zc.Float, zc.Float),
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
