package calc

import (
	"fmt"

	"github.com/blackchip-org/zc/pkg/ops"
	"github.com/blackchip-org/zc/pkg/zc"
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
	zc.Macro("π", "pi"),

	// a
	zc.Macro("a", "add"),
	zc.GenOp("abs",
		zc.Func(ops.AbsBigInt, zc.BigInt),
		zc.Func(ops.AbsDecimal, zc.Decimal),
		zc.Func(ops.AbsFloat, zc.Float),
		zc.Func(ops.AbsRational, zc.Rational),
		zc.Func(ops.AbsComplex, zc.Complex),
	),
	zc.Op("acos", ops.AcosFloat, zc.Float),
	zc.Op("acosh", ops.AcoshFloat, zc.Float),
	zc.Op("asin", ops.AsinFloat, zc.Float),
	zc.Op("asinh", ops.AsinhFloat, zc.Float),
	zc.Op("atan", ops.AtanFloat, zc.Float),
	zc.Op("atanh", ops.AtanhFloat, zc.Float),
	zc.Macro("atto", ops.Atto),
	zc.GenOp("add",
		zc.Func(ops.AddBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.AddDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.AddFloat, zc.Float, zc.Float),
		zc.Func(ops.AddRational, zc.Rational, zc.Rational),
		zc.Func(ops.AddComplex, zc.Complex, zc.Complex),
		zc.Func(ops.AddDurationDateTime, zc.Duration, zc.DateTime),
		zc.Func(ops.AddDateTimeDuration, zc.Time, zc.Duration),
	),
	zc.GenOp("and",
		zc.Func(ops.AndBool, zc.Bool, zc.Bool),
		zc.Func(ops.AndBigInt, zc.BigInt),
	),

	// b
	zc.Op("bin", ops.Bin, zc.BigInt),
	zc.Op("bit", ops.Bit, zc.BigInt, zc.Int),
	zc.Op("bits", ops.Bits, zc.BigInt),
	zc.Op("bytes", ops.Bytes, zc.BigInt),

	// c
	zc.Macro("c", "clear"),
	zc.Macro("c-f", ops.CToF),
	zc.Macro("c-k", ops.CToK),
	zc.GenOp("ceil",
		zc.Func(zc.NoOp, zc.BigInt),
		zc.Func(ops.CeilDecimal, zc.Decimal),
		zc.Func(ops.CeilFloat, zc.Float),
	),
	zc.Macro("centi", ops.Centi),
	zc.Op("char-codepoint", ops.CharToCodePoint, zc.Rune),
	zc.Macro("char-cp", "char-codepoint"),
	zc.Op("clear", ops.Clear),
	zc.Op("cmyk-rgb", ops.CMYKToRGB, zc.Uint8, zc.Uint8, zc.Uint8, zc.Uint8),
	zc.Op("codepoint-char", ops.CodePointToChar, zc.Int32),
	zc.Op("color-sample", ops.ColorSample, zc.Uint8, zc.Uint8, zc.Uint8),
	zc.Op("complex", ops.Complex, zc.Float, zc.Float),
	zc.Op("cos", ops.CosFloat, zc.Float),
	zc.Op("cosh", ops.CoshFloat, zc.Float),
	zc.Macro("cp-char", "codepoint-char"),

	// d
	zc.Macro("d", "div"),
	zc.Op("date", ops.Date, zc.DateTime),
	zc.Op("datetime", ops.DateTime, zc.DateTime),
	zc.Op("day-year", ops.DayYear, zc.DateTime),
	zc.Op("dec", ops.Dec, zc.BigInt),
	zc.Macro("deca", ops.Deca),
	zc.Macro("deci", ops.Deci),
	zc.GenOp("div",
		zc.Func(ops.DivDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.DivFloat, zc.Float, zc.Float),
		zc.Func(ops.DivRational, zc.Rational, zc.Rational),
		zc.Func(ops.DivComplex, zc.Complex, zc.Complex),
	),
	zc.Op("down", ops.Down),
	zc.Op("drop", ops.Drop),
	zc.Macro("dt", "datetime"),
	zc.Op("dup", ops.Dup, zc.String),

	// e
	zc.Macro("e", ops.E),
	zc.GenOp("eq",
		zc.Func(ops.EqBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.EqDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.EqFloat, zc.Float, zc.Float),
		zc.Func(ops.EqRational, zc.Rational, zc.Rational),
		zc.Func(ops.EqComplex, zc.Complex, zc.Complex),
		zc.Func(ops.Is, zc.String, zc.String),
	),
	zc.Op("eval", ops.Eval, zc.String),
	zc.Macro("exa", ops.Exa),
	zc.Op("exp", ops.ExpFloat, zc.Float),

	// f
	zc.Macro("f-c", ops.FToC),
	zc.Macro("false", ops.False),
	zc.Macro("femto", ops.Femto),
	zc.Op("filter", ops.Filter, zc.String),
	zc.GenOp("floor",
		zc.Func(zc.NoOp, zc.BigInt),
		zc.Func(ops.FloorDecimal, zc.Decimal),
		zc.Func(ops.FloorFloat, zc.Float),
	),
	zc.Op("fold", ops.Fold),

	// g
	zc.Macro("giga", ops.Giga),
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

	// h
	zc.Macro("hecto", ops.Hecto),
	zc.Op("hex", ops.HexBigInt, zc.BigInt),
	zc.Op("hours", ops.Hours, zc.Duration),
	zc.Op("hsl-rgb", ops.HSLtoRGB, zc.Float, zc.Float, zc.Float),

	// i
	zc.Op("is", ops.Is, zc.String, zc.String),

	// j
	zc.Op("join", ops.Join, zc.String),

	// k
	zc.Macro("k-c", ops.KToC),
	zc.Macro("kilo", ops.Kilo),
	zc.Macro("km-mi", ops.KmToMi),
	zc.Macro("km-nmi", ops.KmToNmi),

	// l
	zc.Op("left", ops.Left, zc.String, zc.Int),
	zc.Op("len", ops.Len, zc.String),
	zc.Op("local-zone", ops.LocalZone, zc.String),
	zc.Op("local-zone=", ops.LocalZoneGet),
	zc.Op("log", ops.LogFloat, zc.Float),
	zc.Op("log10", ops.Log10Float, zc.Float),
	zc.Op("log2", ops.Log2Float, zc.Float),
	zc.Op("lower", ops.Lower, zc.String),
	zc.Op("lsh", ops.Lsh, zc.BigInt, zc.Uint),
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

	// m
	zc.Macro("m", "mul"),
	zc.Macro("m-nmi", ops.MToNmi),
	zc.Op("map", ops.Map, zc.String),
	zc.Macro("mega", ops.Mega),
	zc.Macro("mi-km", ops.MiToKm),
	zc.Macro("mi-nmi", ops.MiToNmi),
	zc.Macro("micro", ops.Micro),
	zc.Macro("milli", ops.Milli),
	zc.Op("minutes", ops.Minutes, zc.Duration),
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

	// n
	zc.Op("n", ops.N, zc.String),
	zc.Macro("nano", ops.Nano),
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
	zc.Macro("nmi-km", ops.NmiToKm),
	zc.Macro("nmi-m", ops.NmiToM),
	zc.Macro("nmi-mi", ops.NmiToMi),
	zc.GenOp("not",
		zc.Func(ops.NotBool, zc.Bool),
		zc.Func(ops.NotBigInt, zc.BigInt),
	),
	zc.Op("now", ops.Now),
	zc.Op("now-set", ops.NowSet, zc.DateTime),
	zc.Op("now-restore", ops.NowRestore),

	// o
	zc.Op("oct", ops.Oct, zc.BigInt),
	zc.GenOp("or",
		zc.Func(ops.OrBool, zc.Bool, zc.Bool),
		zc.Func(ops.OrBigInt, zc.BigInt, zc.BigInt),
	),

	// p
	zc.Macro("peta", ops.Peta),
	zc.Macro("pi", ops.Pi),
	zc.Macro("pico", ops.Pico),
	zc.GenOp("pow",
		zc.Func(ops.PowBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.PowFloat, zc.Float, zc.Float),
		zc.Func(ops.PowComplex, zc.Complex, zc.Complex),
	),

	// q
	zc.Macro("quetta", ops.Quetta),
	zc.Macro("quecto", ops.Quecto),

	// r
	zc.Macro("r", "round"),
	zc.Op("rand", ops.Rand),
	zc.Op("rand-choice", ops.RandChoice),
	zc.Op("rand-int", ops.RandInt, zc.Int),
	zc.Op("rand-seed", ops.RandSeed, zc.Int),
	zc.Op("rand-seed=", ops.RandSeedGet),
	zc.Macro("reduce", "fold"),
	zc.GenOp("rem",
		zc.Func(ops.RemBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.RemFloat, zc.Float, zc.Float),
	),
	zc.Op("repeat", ops.Repeat, zc.String, zc.Int),
	zc.Op("reverse", ops.Reverse),
	zc.Op("rgb-cmyk", ops.RBGToCMYK, zc.Uint8, zc.Uint8, zc.Uint8),
	zc.Op("rgb-hsl", ops.RGBToHSL, zc.Uint8, zc.Uint8, zc.Uint8),
	zc.Op("right", ops.Right, zc.String, zc.Int),
	zc.Op("roll", ops.Roll, zc.String),
	zc.Macro("ronna", ops.Ronna),
	zc.Macro("ronto", ops.Ronto),
	zc.Op("rot-13", ops.Rot13, zc.String),
	zc.Op("round", ops.Round, zc.Decimal, zc.Int),
	zc.Op("rounding-mode", ops.RoundingMode, zc.String),
	zc.Op("rounding-mode=", ops.RoundingModeGet),
	zc.Op("rsh", ops.Rsh, zc.BigInt, zc.Uint),

	// s
	zc.Macro("s", "sub"),
	zc.Op("scientific-notation", ops.ScientificNotation, zc.Float),
	zc.Op("seconds", ops.Seconds, zc.Duration),
	zc.Op("shuffle", ops.Shuffle),
	zc.GenOp("sign",
		zc.Func(ops.SignBigInt, zc.BigInt),
		zc.Func(ops.SignDecimal, zc.Decimal),
		zc.Func(ops.SignFloat, zc.Float),
		zc.Func(ops.SignRational, zc.Rational),
	),
	zc.Op("sin", ops.SinFloat, zc.Float),
	zc.Op("sinh", ops.SinhFloat, zc.Float),
	zc.Macro("sn", "scientific-notation"),
	zc.Op("split", ops.Split, zc.String, zc.String),
	zc.GenOp("sqrt",
		zc.Func(ops.SqrtFloat, zc.Float),
	),
	zc.GenOp("sub",
		zc.Func(ops.SubBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.SubDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.SubFloat, zc.Float, zc.Float),
		zc.Func(ops.SubRational, zc.Rational, zc.Rational),
		zc.Func(ops.SubComplex, zc.Complex, zc.Complex),
		zc.Func(ops.SubDateTime, zc.DateTime, zc.DateTime),
	),
	zc.Macro("sum", ops.Sum),
	zc.Op("swap", ops.Swap, zc.String, zc.String),

	// t
	zc.Op("take", ops.Take, zc.Int),
	zc.Op("tan", ops.TanFloat, zc.Float),
	zc.Op("tanh", ops.TanhFloat, zc.Float),
	zc.Macro("tera", ops.Tera),
	zc.Op("time", ops.Time, zc.DateTime),
	zc.Op("time-zone", ops.TimeZone, zc.DateTime, zc.String),
	zc.Macro("top", ops.Top),
	zc.Macro("true", ops.True),
	zc.Macro("tz", "time-zone"),

	// u
	zc.Macro("u8en", "utf-8-encode"),
	zc.Macro("u8de", "utf-8-decode"),
	zc.Op("up", ops.Up),
	zc.Op("upper", ops.Upper, zc.String),
	zc.Op("utf-8-decode", ops.UTF8Decode, zc.BigInt),
	zc.Op("utf-8-encode", ops.UTF8Encode, zc.String),

	// v
	zc.Op("version", ops.Version),

	// x
	zc.Op("xor", ops.Xor, zc.BigInt, zc.BigInt),

	// y
	zc.Macro("yotta", ops.Yotta),
	zc.Macro("yocto", ops.Yocto),

	// z
	zc.Macro("zepto", ops.Zepto),
	zc.Macro("zetta", ops.Zetta),
}

var opsTable map[string]zc.CalcFunc

func addMacros(table map[string]string) {
	for k, v := range table {
		opsTable[k] = evalOp(zc.Macro(k, v))
	}
}

func init() {
	opsTable = make(map[string]zc.CalcFunc)
	for _, v := range opsList {
		k := v.Name
		if _, exists := opsTable[k]; exists {
			panic(fmt.Sprintf("duplicate operation: %v", k))
		}
		opsTable[k] = evalOp(v)
	}
	addMacros(ops.TimeZones)
}
