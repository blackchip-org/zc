package calc

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/ops"
)

var opsList = []zc.OpDecl{
	zc.Macro("!=", "neq"),
	zc.Macro("%", "mod"),
	zc.Macro("*", "mul"),
	zc.Macro("**", "pow"),
	zc.Macro("+", "add"),
	zc.Macro("-", "sub"),
	zc.Macro("/", "div"),
	zc.Macro("<", "lt"),
	zc.Macro("<=", "lte"),
	zc.Macro("==", "eq"),
	zc.Macro(">", "gt"),
	zc.Macro(">=", "gte"),
	zc.Macro("^", "pow"),
	zc.Macro("a", "add"),
	zc.GenOp("abs",
		zc.Func(ops.AbsBigInt, zc.BigInt),
		zc.Func(ops.AbsDecimal, zc.Decimal),
		zc.Func(ops.AbsFloat, zc.Float),
		zc.Func(ops.AbsRational, zc.Rational),
		zc.Func(ops.AbsComplex, zc.Complex),
	),
	zc.GenOp("add",
		zc.Func(ops.AddBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.AddDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.AddFloat, zc.Float, zc.Float),
		zc.Func(ops.AddRational, zc.Rational, zc.Rational),
		zc.Func(ops.AddComplex, zc.Complex, zc.Complex),
	),
	zc.GenOp("and",
		zc.Func(ops.AndBigInt, zc.BigInt),
	),
	zc.Macro("c", "clear"),
	zc.GenOp("ceil",
		zc.Func(zc.NoOp, zc.BigInt),
		zc.Func(ops.CeilDecimal, zc.Decimal),
		zc.Func(ops.CeilFloat, zc.Float),
	),
	zc.Op("clear", ops.Clear),
	zc.Macro("d", "div"),
	zc.GenOp("div",
		zc.Func(ops.DivDecimal, zc.Decimal),
		zc.Func(ops.DivFloat, zc.Float, zc.Float),
		zc.Func(ops.DivRational, zc.Rational, zc.Rational),
		zc.Func(ops.DivComplex, zc.Complex, zc.Complex),
	),
	zc.GenOp("eq",
		zc.Func(ops.EqBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.EqDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.EqFloat, zc.Float, zc.Float),
		zc.Func(ops.EqRational, zc.Rational, zc.Rational),
		zc.Func(ops.EqComplex, zc.Complex, zc.Complex),
	),
	zc.GenOp("floor",
		zc.Func(zc.NoOp, zc.BigInt),
		zc.Func(ops.FloorDecimal, zc.Decimal),
		zc.Func(ops.FloorFloat, zc.Float),
	),
	zc.Op("fold", ops.Fold),
	zc.GenOp("gt",
		zc.Func(ops.GtBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.GtDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.GtRational, zc.Rational, zc.Rational),
		zc.Func(ops.GtFloat, zc.Float, zc.Float),
	),
	zc.GenOp("gte",
		zc.Func(ops.GteBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.GteDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.GteRational, zc.Rational, zc.Rational),
		zc.Func(ops.GteFloat, zc.Float, zc.Float),
	),
	zc.Op("hex", ops.HexBigInt, zc.BigInt),
	zc.GenOp("lt",
		zc.Func(ops.LtBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.LtDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.LtRational, zc.Rational, zc.Rational),
		zc.Func(ops.LtFloat, zc.Float, zc.Float),
	),
	zc.GenOp("lte",
		zc.Func(ops.LteBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.LteDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.LteRational, zc.Rational, zc.Rational),
		zc.Func(ops.LteFloat, zc.Float, zc.Float),
	),
	zc.Macro("m", "mul"),
	zc.Op("map", ops.Map, zc.String),
	zc.GenOp("mod",
		zc.Func(ops.ModBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.ModDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.ModFloat, zc.Float, zc.Float),
	),
	zc.GenOp("mul",
		zc.Func(ops.MulBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.MulDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.MulFloat, zc.Float, zc.Float),
		zc.Func(ops.MulRational, zc.Rational, zc.Rational),
		zc.Func(ops.MulComplex, zc.Complex, zc.Complex),
	),
	zc.GenOp("neg",
		zc.Func(ops.NegBigInt, zc.BigInt),
		zc.Func(ops.NegDecimal, zc.Decimal),
		zc.Func(ops.NegFloat, zc.Float),
		zc.Func(ops.NegRational, zc.Rational),
	),
	zc.GenOp("neq",
		zc.Func(ops.NeqBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.NeqDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.NeqFloat, zc.Float, zc.Float),
		zc.Func(ops.NeqRational, zc.Rational, zc.Rational),
		zc.Func(ops.NeqComplex, zc.Complex, zc.Complex),
	),
	zc.GenOp("pow",
		zc.Func(ops.PowBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.PowFloat, zc.Float, zc.Float),
		zc.Func(ops.PowComplex, zc.Complex, zc.Complex),
	),
	zc.GenOp("rem",
		zc.Func(ops.RemBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.RemFloat, zc.Float, zc.Float),
	),
	zc.Macro("s", "sub"),
	zc.GenOp("sign",
		zc.Func(ops.SignBigInt, zc.BigInt),
		zc.Func(ops.SignDecimal, zc.Decimal),
		zc.Func(ops.SignFloat, zc.Float),
		zc.Func(ops.SignRational, zc.Rational),
	),
	zc.GenOp("sub",
		zc.Func(ops.SubBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.SubDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.SubFloat, zc.Float, zc.Float),
		zc.Func(ops.SubRational, zc.Rational, zc.Rational),
		zc.Func(ops.SubComplex, zc.Complex, zc.Complex),
	),
	zc.Macro("sum", ops.Sum),
}

var opsTable map[string]zc.CalcFunc

func init() {
	opsTable = make(map[string]zc.CalcFunc)
	for _, v := range opsList {
		k := v.Name
		if _, exists := opsTable[k]; exists {
			panic(zc.ErrDuplicateOp(k))
		}
		opsTable[k] = zc.EvalOp(v)
	}
}
