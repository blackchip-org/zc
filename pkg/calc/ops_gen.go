// Code generated by "gen-ops"; DO NOT EDIT
package calc

import (
	"github.com/blackchip-org/zc/pkg/ops"
	"github.com/blackchip-org/zc/pkg/zc"
)

var opsList = []zc.OpDecl{
	zc.Macro("*", "mul"),
	zc.Macro("**", "pow"),
	zc.Macro("+", "add"),
	zc.Macro("-", "sub"),
	zc.Macro("/", "div"),
	zc.Macro("^", "pow"),
	zc.Macro("a", "add"),
	zc.GenOp("abs",
		zc.Func(ops.AbsBigInt, zc.BigInt),
		zc.Func(ops.AbsDecimal, zc.Decimal),
		zc.Func(ops.AbsFloat, zc.Float),
		zc.Func(ops.AbsRational, zc.Rational),
		zc.Func(ops.AbsComplex, zc.Complex),
	),
	zc.GenOp("acos",
		zc.Func(ops.AcosFloat, zc.Float),
		zc.Func(ops.AcosComplex, zc.Complex),
	),
	zc.GenOp("acosh",
		zc.Func(ops.AcoshFloat, zc.Float),
		zc.Func(ops.AcoshComplex, zc.Complex),
	),
	zc.GenOp("add",
		zc.Func(ops.AddBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.AddDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.AddFloat, zc.Float, zc.Float),
		zc.Func(ops.AddRational, zc.Rational, zc.Rational),
		zc.Func(ops.AddComplex, zc.Complex, zc.Complex),
		zc.Func(ops.AddDuration, zc.Duration, zc.Duration),
		zc.Func(ops.AddDurationDateTime, zc.Duration, zc.DateTime),
		zc.Func(ops.AddDateTimeDuration, zc.DateTime, zc.Duration),
	),
	zc.GenOp("and",
		zc.Func(ops.AndBool, zc.Bool, zc.Bool),
		zc.Func(ops.AndBigInt, zc.BigInt, zc.BigInt),
	),
	zc.GenOp("asin",
		zc.Func(ops.AsinFloat, zc.Float),
		zc.Func(ops.AsinComplex, zc.Complex),
	),
	zc.GenOp("asinh",
		zc.Func(ops.AsinhFloat, zc.Float),
		zc.Func(ops.AsinComplex, zc.Complex),
	),
	zc.GenOp("atan",
		zc.Func(ops.AtanFloat, zc.Float),
		zc.Func(ops.AtanComplex, zc.Complex),
	),
	zc.GenOp("atanh",
		zc.Func(ops.AtanhFloat, zc.Float),
		zc.Func(ops.AtanhComplex, zc.Float),
	),
	zc.GenOp("average",
		zc.Func(ops.Average),
	),
	zc.Macro("avg", "average"),
	zc.Op("bin", ops.Bin, zc.BigInt),
	zc.Op("bit", ops.Bit, zc.BigInt, zc.Int),
	zc.Op("bits", ops.Bits, zc.BigInt),
	zc.Op("bytes", ops.Bytes, zc.BigInt),
	zc.Macro("c", "clear"),
	zc.Macro("c-f", "9 5 div mul 32 add"),
	zc.Macro("c-k", "273.15 add"),
	zc.GenOp("ceil",
		zc.Func(zc.NoOp, zc.BigInt),
		zc.Func(ops.CeilDecimal, zc.Decimal),
		zc.Func(ops.CeilFloat, zc.Float),
	),
	zc.Op("char-codepoint", ops.CharToCodePoint, zc.Char),
	zc.Macro("char-cp", "char-codepoint"),
	zc.Op("clear", ops.Clear),
	zc.Op("cmyk-rgb", ops.CMYKToRGB, zc.Uint8, zc.Uint8, zc.Uint8, zc.Uint8),
	zc.Op("codepoint-char", ops.CodePointToChar, zc.Int32),
	zc.Op("color-sample", ops.ColorSample, zc.Uint8, zc.Uint8, zc.Uint8),
	zc.Op("complex", ops.Complex, zc.Float, zc.Float),
	zc.Op("conj", ops.Conj, zc.Complex),
	zc.GenOp("cos",
		zc.Func(ops.CosFloat, zc.Float),
		zc.Func(ops.CosComplex, zc.Complex),
	),
	zc.GenOp("cosh",
		zc.Func(ops.CoshFloat, zc.Float),
		zc.Func(ops.CoshComplex, zc.Float),
	),
	zc.Op("cot", ops.CotComplex, zc.Complex),
	zc.Macro("cp-char", "codepoint-char"),
	zc.Macro("d", "div"),
	zc.Op("date", ops.Date, zc.DateTime),
	zc.Op("datetime", ops.DateTime, zc.DateTime),
	zc.Op("day-year", ops.DayYear, zc.DateTime),
	zc.GenOp("dec",
		zc.Func(ops.DecDMS, zc.DMS),
		zc.Func(ops.Dec, zc.Decimal),
		zc.Func(ops.DecFloat, zc.Float),
		zc.Func(ops.DecRational, zc.Rational),
		zc.Func(ops.DecBigInt, zc.BigInt),
	),
	zc.Op("deg-min", ops.DM, zc.DMS),
	zc.Op("deg-min-round", ops.DMRound, zc.DMS, zc.Int32),
	zc.Op("deg-min-sec", ops.DMS, zc.DMS),
	zc.Op("deg-min-sec-round", ops.DMSRound, zc.DMS, zc.Int32),
	zc.Macro("deg-rad", "pi 180 div mul"),
	zc.Op("denom", ops.Denom, zc.Rational),
	zc.GenOp("div",
		zc.Func(ops.DivDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.DivFloat, zc.Float, zc.Float),
		zc.Func(ops.DivRational, zc.Rational, zc.Rational),
		zc.Func(ops.DivComplex, zc.Complex, zc.Complex),
	),
	zc.Macro("dm", "deg-min"),
	zc.Macro("dmr", "deg-min-round"),
	zc.Macro("dms", "deg-min-sec"),
	zc.Macro("dmsr", "deg-min-sec-round"),
	zc.Macro("dn", "down"),
	zc.Op("down", ops.Down),
	zc.Macro("doy", "day-year"),
	zc.Op("drop", ops.Drop, zc.Str),
	zc.Macro("dt", "datetime"),
	zc.Op("dup", ops.Dup, zc.Str),
	zc.Macro("e", "2.718281828459045"),
	zc.Op("epsg.utm", ops.UTM, zc.Str),
	zc.Macro("epsg.web-mercator", "'EPSG:3857'"),
	zc.Macro("epsg.wgs-84", "'EPSG:4326'"),
	zc.GenOp("eq",
		zc.Func(ops.EqBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.EqDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.EqFloat, zc.Float, zc.Float),
		zc.Func(ops.EqRational, zc.Rational, zc.Rational),
		zc.Func(ops.EqComplex, zc.Complex, zc.Complex),
		zc.Func(ops.Is, zc.Str, zc.Str),
	),
	zc.Op("eval", ops.Eval),
	zc.GenOp("exp",
		zc.Func(ops.ExpFloat, zc.Float),
		zc.Func(ops.ExpComplex, zc.Complex),
	),
	zc.Macro("f-c", "32 sub 5 9 div mul"),
	zc.Macro("false", "[false]"),
	zc.Op("filter", ops.Filter),
	zc.GenOp("floor",
		zc.Func(zc.NoOp, zc.BigInt),
		zc.Func(ops.FloorDecimal, zc.Decimal),
		zc.Func(ops.FloorFloat, zc.Float),
	),
	zc.Op("fold", ops.Fold),
	zc.Macro("greater-than", "gt"),
	zc.Macro("greater-than-or-equal", "gte"),
	zc.GenOp("gt",
		zc.Func(ops.GtBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.GtDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.GtRational, zc.Rational, zc.Rational),
		zc.Func(ops.GtFloat, zc.Float, zc.Float),
		zc.Func(ops.GtStr, zc.Str, zc.Str),
	),
	zc.GenOp("gte",
		zc.Func(ops.GteBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.GteDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.GteRational, zc.Rational, zc.Rational),
		zc.Func(ops.GteFloat, zc.Float, zc.Float),
		zc.Func(ops.GteStr, zc.Str, zc.Str),
	),
	zc.Op("haversine", ops.Haversine, zc.DMS, zc.DMS, zc.DMS, zc.DMS),
	zc.Op("hex", ops.HexBigInt, zc.BigInt),
	zc.Op("hours", ops.Hours, zc.Duration),
	zc.Op("hsl-rgb", ops.HSLToRGB, zc.Float, zc.Float, zc.Float),
	zc.Op("imag", ops.Imag, zc.Complex),
	zc.Op("is", ops.Is, zc.Str, zc.Str),
	zc.Macro("is-dec-min-sec", "is-dms"),
	zc.Op("is-dms", ops.IsDMS, zc.Str),
	zc.Op("is-not", ops.IsNot, zc.Str, zc.Str),
	zc.Op("join", ops.Join),
	zc.Macro("k-c", "273.15 sub"),
	zc.Macro("km-mi", "0.62137119 mul"),
	zc.Macro("km-nmi", "0.539957 mul"),
	zc.Op("left", ops.Left, zc.Str, zc.Int),
	zc.Macro("left-shift", "lsh"),
	zc.Op("len", ops.Len, zc.Str),
	zc.Macro("less-than", "lt"),
	zc.Macro("less-than-or-equal", "lte"),
	zc.Op("local-zone", ops.LocalZone, zc.Str),
	zc.Op("local-zone=", ops.LocalZoneGet),
	zc.GenOp("log",
		zc.Func(ops.LogFloat, zc.Float),
		zc.Func(ops.LogFloat, zc.Complex),
	),
	zc.GenOp("log10",
		zc.Func(ops.Log10Float, zc.Float),
		zc.Func(ops.Log10Complex, zc.Complex),
	),
	zc.Op("log2", ops.Log2Float, zc.Float),
	zc.Op("lower", ops.Lower, zc.Str),
	zc.Op("lsh", ops.Lsh, zc.BigInt, zc.Uint),
	zc.GenOp("lt",
		zc.Func(ops.LtBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.LtDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.LtRational, zc.Rational, zc.Rational),
		zc.Func(ops.LtFloat, zc.Float, zc.Float),
		zc.Func(ops.LtStr, zc.Str, zc.Str),
	),
	zc.GenOp("lte",
		zc.Func(ops.LteBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.LteDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.LteRational, zc.Rational, zc.Rational),
		zc.Func(ops.LteFloat, zc.Float, zc.Float),
		zc.Func(ops.LteStr, zc.Str, zc.Str),
	),
	zc.Macro("m", "mul"),
	zc.Macro("m-nmi", "0.000539957 mul"),
	zc.Op("map", ops.Map),
	zc.Macro("mi-km", "1.609344 mul"),
	zc.Macro("mi-nmi", "0.868976 mul"),
	zc.GenOp("minutes",
		zc.Func(ops.MinutesDMS, zc.DMS),
		zc.Func(ops.Minutes, zc.Duration),
	),
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
	zc.Op("n", ops.N),
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
		zc.Func(ops.IsNot, zc.Str, zc.Str),
	),
	zc.Macro("nmi-km", "1.852 mul"),
	zc.Macro("nmi-m", "1852 mul"),
	zc.Macro("nmi-mi", "1.15078 mul"),
	zc.GenOp("not",
		zc.Func(ops.NotBool, zc.Bool),
		zc.Func(ops.NotBigInt, zc.BigInt),
	),
	zc.Op("now", ops.Now),
	zc.Op("now-restore", ops.NowRestore),
	zc.Op("now-set", ops.NowSet, zc.DateTime),
	zc.Op("num", ops.Num, zc.Rational),
	zc.Op("oct", ops.Oct, zc.BigInt),
	zc.GenOp("or",
		zc.Func(ops.OrBool, zc.Bool, zc.Bool),
		zc.Func(ops.OrBigInt, zc.BigInt, zc.BigInt),
	),
	zc.Op("phase", ops.PhaseComplex, zc.Complex),
	zc.Macro("pi", "3.14159265358979323"),
	zc.Op("polar", ops.PolarComplex, zc.Complex),
	zc.GenOp("pow",
		zc.Func(ops.PowBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.PowFloat, zc.Float, zc.Float),
		zc.Func(ops.PowComplex, zc.Complex, zc.Complex),
	),
	zc.Op("proj", ops.Proj, zc.Float, zc.Float, zc.Str, zc.Str),
	zc.Macro("r", "round"),
	zc.Macro("rad-deg", "180 pi div mul"),
	zc.Op("rand", ops.Rand),
	zc.Op("rand-choice", ops.RandChoice, zc.Str),
	zc.Op("rand-int", ops.RandInt, zc.Int),
	zc.Op("rand-seed", ops.RandSeed, zc.Int64),
	zc.Op("rand-seed=", ops.RandSeedGet),
	zc.Op("real", ops.Real, zc.Complex),
	zc.Op("rect", ops.RectComplex, zc.Float, zc.Float),
	zc.Macro("reduce", "fold"),
	zc.GenOp("rem",
		zc.Func(ops.RemBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.RemFloat, zc.Float, zc.Float),
	),
	zc.Op("repeat", ops.Repeat),
	zc.Macro("rev", "reverse"),
	zc.Op("reverse", ops.Reverse),
	zc.Op("rgb-cmyk", ops.RGBToCMYK, zc.Uint8, zc.Uint8, zc.Uint8),
	zc.Op("rgb-hsl", ops.RGBToHSL, zc.Uint8, zc.Uint8, zc.Uint8),
	zc.Op("right", ops.Right, zc.Str, zc.Int),
	zc.Macro("right-shift", "rsh"),
	zc.Op("roll", ops.Roll, zc.Str),
	zc.Macro("rot13", "rotate-13"),
	zc.Op("rotate-13", ops.Rot13, zc.Str),
	zc.GenOp("round",
		zc.Func(ops.RoundDecimal, zc.Decimal, zc.Int),
		zc.Func(ops.RoundFloat, zc.Float, zc.Int),
	),
	zc.Op("rounding-mode", ops.RoundingMode),
	zc.Op("rounding-mode=", ops.RoundingModeGet),
	zc.Op("rsh", ops.Rsh, zc.BigInt, zc.Uint),
	zc.Macro("s", "sub"),
	zc.Op("scientific-notation", ops.ScientificNotation, zc.Float),
	zc.GenOp("seconds",
		zc.Func(ops.SecondsDMS, zc.DMS),
		zc.Func(ops.Seconds, zc.Duration),
	),
	zc.Op("shuffle", ops.Shuffle),
	zc.Macro("si.atto", "1e-18"),
	zc.Macro("si.centi", "1e-02"),
	zc.Macro("si.deca", "1e01"),
	zc.Macro("si.deci", "1e-01"),
	zc.Macro("si.exa", "1e18"),
	zc.Macro("si.femto", "1e-15"),
	zc.Macro("si.giga", "1e09"),
	zc.Macro("si.hecto", "1e02"),
	zc.Macro("si.kilo", "1e03"),
	zc.Macro("si.mega", "1e06"),
	zc.Macro("si.micro", "1e-06"),
	zc.Macro("si.milli", "1e-03"),
	zc.Macro("si.nano", "1e-09"),
	zc.Macro("si.peta", "1e15"),
	zc.Macro("si.pico", "1e-12"),
	zc.Macro("si.quecto", "1e-30"),
	zc.Macro("si.quetta", "1e30"),
	zc.Macro("si.ronna", "1e27"),
	zc.Macro("si.ronto", "1e-27"),
	zc.Macro("si.terra", "1e12"),
	zc.Macro("si.yocto", "1e-24"),
	zc.Macro("si.yotta", "1e24"),
	zc.Macro("si.zepto", "1e-21"),
	zc.Macro("si.zetta", "1e21"),
	zc.GenOp("sign",
		zc.Func(ops.SignBigInt, zc.BigInt),
		zc.Func(ops.SignDecimal, zc.Decimal),
		zc.Func(ops.SignFloat, zc.Float),
		zc.Func(ops.SignRational, zc.Rational),
	),
	zc.GenOp("sin",
		zc.Func(ops.SinFloat, zc.Float),
		zc.Func(ops.SinComplex, zc.Complex),
	),
	zc.GenOp("sinh",
		zc.Func(ops.SinhFloat, zc.Float),
		zc.Func(ops.SinhComplex, zc.Complex),
	),
	zc.Macro("sn", "scientific-notation"),
	zc.Op("split", ops.Split, zc.Str, zc.Str),
	zc.GenOp("sqrt",
		zc.Func(ops.SqrtFloat, zc.Float),
		zc.Func(ops.SqrtComplex, zc.Complex),
	),
	zc.Macro("square-root", "sqrt"),
	zc.GenOp("sub",
		zc.Func(ops.SubBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.SubDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.SubFloat, zc.Float, zc.Float),
		zc.Func(ops.SubRational, zc.Rational, zc.Rational),
		zc.Func(ops.SubComplex, zc.Complex, zc.Complex),
		zc.Func(ops.SubDuration, zc.Duration, zc.Duration),
		zc.Func(ops.SubDateTimeDuration, zc.DateTime, zc.Duration),
	),
	zc.Macro("sum", "[add] fold"),
	zc.Op("swap", ops.Swap, zc.Str, zc.Str),
	zc.Op("take", ops.Take, zc.Int),
	zc.GenOp("tan",
		zc.Func(ops.TanFloat, zc.Float),
		zc.Func(ops.TanComplex, zc.Complex),
	),
	zc.GenOp("tanh",
		zc.Func(ops.TanhFloat, zc.Float),
		zc.Func(ops.TanhComplex, zc.Complex),
	),
	zc.Op("time", ops.Time, zc.DateTime),
	zc.Op("timezone", ops.TimeZone, zc.DateTime, zc.Str),
	zc.Macro("top", "1 take"),
	zc.Macro("true", "[true]"),
	zc.Macro("tz", "timezone"),
	zc.Macro("u8de", "utf8-decode"),
	zc.Macro("u8en", "utf8-encode"),
	zc.Op("up", ops.Up),
	zc.Op("upper", ops.Upper, zc.Str),
	zc.Op("utf8-decode", ops.UTF8Decode, zc.BigInt),
	zc.Op("utf8-encode", ops.UTF8Encode, zc.Str),
	zc.Op("version", ops.Version),
	zc.Op("xor", ops.Xor, zc.BigInt, zc.BigInt),
	zc.Macro("π", "pi"),
}