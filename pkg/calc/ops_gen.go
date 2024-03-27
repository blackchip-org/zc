// Code generated by "gen-ops"; DO NOT EDIT
package calc

import (
	"github.com/blackchip-org/zc/v5/pkg/ops"
	"github.com/blackchip-org/zc/v5/pkg/zc"
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
		zc.Func(ops.AbsComplex, zc.Complex),
		zc.Func(ops.AbsBigInt, zc.BigInt),
		zc.Func(ops.AbsDecimal, zc.Decimal),
		zc.Func(ops.AbsFloat, zc.Float),
		zc.Func(ops.AbsRational, zc.Rational),
	),
	zc.GenOp("acos",
		zc.Func(ops.AcosComplex, zc.Complex),
		zc.Func(ops.AcosFloat, zc.Float),
	),
	zc.GenOp("acosh",
		zc.Func(ops.AcoshComplex, zc.Complex),
		zc.Func(ops.AcoshFloat, zc.Float),
	),
	zc.GenOp("add",
		zc.Func(ops.AddBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.AddDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.AddBigFloat, zc.BigFloat, zc.BigFloat),
		zc.Func(ops.AddFloat, zc.Float, zc.Float),
		zc.Func(ops.AddRational, zc.Rational, zc.Rational),
		zc.Func(ops.AddComplex, zc.Complex, zc.Complex),
		zc.Func(ops.AddDuration, zc.Duration, zc.Duration),
		zc.Func(ops.AddDurationDateTime, zc.Duration, zc.DateTime),
		zc.Func(ops.AddDateTimeDuration, zc.DateTime, zc.Duration),
	),
	zc.Op("add-i16", ops.AddInt16, zc.Int16, zc.Int16),
	zc.Op("add-i32", ops.AddInt32, zc.Int32, zc.Int32),
	zc.Op("add-i64", ops.AddInt64, zc.Int64, zc.Int64),
	zc.Op("add-i8", ops.AddInt8, zc.Int8, zc.Int8),
	zc.Op("add-ia", ops.AddIntArch, zc.Int, zc.Int),
	zc.Op("add-u16", ops.AddUint16, zc.Uint16, zc.Uint16),
	zc.Op("add-u32", ops.AddUint32, zc.Uint32, zc.Uint32),
	zc.Op("add-u64", ops.AddUint64, zc.Uint64, zc.Uint64),
	zc.Op("add-u8", ops.AddUint8, zc.Uint8, zc.Uint8),
	zc.Op("add-ua", ops.AddUintArch, zc.Uint, zc.Uint),
	zc.GenOp("and",
		zc.Func(ops.AndBool, zc.Bool, zc.Bool),
		zc.Func(ops.AndBigInt, zc.BigInt, zc.BigInt),
	),
	zc.Op("anno", ops.Anno, zc.Val, zc.Str),
	zc.Macro("annotate", "anno"),
	zc.Op("apply", ops.Apply),
	zc.GenOp("asin",
		zc.Func(ops.AsinComplex, zc.Complex),
		zc.Func(ops.AsinFloat, zc.Float),
	),
	zc.GenOp("asinh",
		zc.Func(ops.AsinComplex, zc.Complex),
		zc.Func(ops.AsinhFloat, zc.Float),
	),
	zc.GenOp("atan",
		zc.Func(ops.AtanComplex, zc.Complex),
		zc.Func(ops.AtanFloat, zc.Float),
	),
	zc.Op("atan2", ops.Atan2Float, zc.Float, zc.Float),
	zc.GenOp("atanh",
		zc.Func(ops.AtanhComplex, zc.Complex),
		zc.Func(ops.AtanhFloat, zc.Float),
	),
	zc.Macro("atto", "1e-18"),
	zc.GenOp("average",
		zc.Func(ops.Average),
	),
	zc.Macro("avg", "average"),
	zc.Op("bin", ops.Bin, zc.BigInt),
	zc.Op("bit", ops.Bit, zc.BigInt, zc.Int),
	zc.Op("bits", ops.Bits, zc.BigInt),
	zc.Op("bytes", ops.Bytes, zc.BigInt),
	zc.Macro("c", "clear"),
	zc.Macro("c-f", "9 5 div mul 32 add /°F anno"),
	zc.Macro("c-k", "273.15 add /K anno"),
	zc.Macro("ca", "clear-all"),
	zc.GenOp("ceil",
		zc.Func(zc.NoOp, zc.BigInt),
		zc.Func(ops.CeilDecimal, zc.Decimal),
		zc.Func(ops.CeilFloat, zc.Float),
	),
	zc.Macro("centi", "1e-02"),
	zc.Op("char-codepoint", ops.CharToCodePoint, zc.Char),
	zc.Macro("char-cp", "char-codepoint"),
	zc.Op("clear", ops.Clear),
	zc.Op("clear-all", ops.ClearAll),
	zc.Op("cmyk-rgb", ops.CMYKToRGB, zc.Uint8, zc.Uint8, zc.Uint8, zc.Uint8),
	zc.Op("codepoint-char", ops.CodePointToChar, zc.Int32),
	zc.Op("color-sample", ops.ColorSample, zc.Uint8, zc.Uint8, zc.Uint8),
	zc.Op("complex", ops.Complex, zc.Float, zc.Float),
	zc.Op("conj", ops.Conj, zc.Complex),
	zc.GenOp("cos",
		zc.Func(ops.CosComplex, zc.Complex),
		zc.Func(ops.CosFloat, zc.Float),
	),
	zc.GenOp("cosh",
		zc.Func(ops.CoshComplex, zc.Complex),
		zc.Func(ops.CoshFloat, zc.Float),
	),
	zc.Op("cot", ops.CotComplex, zc.Complex),
	zc.Macro("cp-char", "codepoint-char"),
	zc.Macro("d", "div"),
	zc.Macro("d-int", "div-int"),
	zc.Op("date", ops.Date, zc.DateTime),
	zc.Op("datetime", ops.DateTime, zc.DateTime),
	zc.Op("datetime?", ops.DateTimeIs, zc.Str),
	zc.Op("day-year", ops.DayYear, zc.DateTime),
	zc.GenOp("dec",
		zc.Func(ops.DecDMS, zc.DMS),
		zc.Func(ops.Dec, zc.Decimal),
		zc.Func(ops.DecFloat, zc.Float),
		zc.Func(ops.DecBigInt, zc.BigInt),
		zc.Func(ops.DecRational, zc.Rational),
	),
	zc.Macro("dec-min-sec?", "dms?"),
	zc.Op("dec-prec", ops.DecPrecGet),
	zc.Op("dec-prec=", ops.DecPrecSet, zc.Int),
	zc.Op("dec?", ops.DecimalIs, zc.Str),
	zc.Macro("deca", "1e01"),
	zc.Macro("deci", "1e-01"),
	zc.Op("deg-min", ops.DM, zc.DMS),
	zc.Op("deg-min-round", ops.DMRound, zc.DMS, zc.Int),
	zc.Op("deg-min-sec", ops.DMS, zc.DMS),
	zc.Op("deg-min-sec-round", ops.DMSRound, zc.DMS, zc.Int),
	zc.Macro("deg-rad", "pi 180 div mul"),
	zc.Op("denom", ops.Denom, zc.Rational),
	zc.GenOp("div",
		zc.Func(ops.DivDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.DivBigFloat, zc.BigFloat, zc.BigFloat),
		zc.Func(ops.DivFloat, zc.Float, zc.Float),
		zc.Func(ops.DivRational, zc.Rational, zc.Rational),
		zc.Func(ops.DivComplex, zc.Complex, zc.Complex),
	),
	zc.Op("div-int", ops.DivBigInt, zc.BigInt, zc.BigInt),
	zc.Op("div-mod-int", ops.DivModBigInt, zc.BigInt, zc.BigInt),
	zc.Op("div-rem", ops.DivRemDec, zc.Decimal, zc.Decimal, zc.Int32),
	zc.Macro("dm", "deg-min"),
	zc.Macro("dm-int", "div-mod-int"),
	zc.Macro("dmr", "deg-min-round"),
	zc.Macro("dms", "deg-min-sec"),
	zc.Op("dms?", ops.DMSIs, zc.Str),
	zc.Macro("dmsr", "deg-min-sec-round"),
	zc.Macro("dn", "down"),
	zc.Op("down", ops.Down),
	zc.Macro("doy", "day-year"),
	zc.Macro("dr", "div-rem"),
	zc.Op("drop", ops.Drop, zc.Val),
	zc.Macro("dt", "datetime"),
	zc.Macro("dt?", "datetime?"),
	zc.Op("dup", ops.Dup, zc.Val),
	zc.Macro("e", "2.71828182845904523536028747135266249775724709369995957496696763"),
	zc.Macro("earth-equatorial-radius", "6378137 /m anno"),
	zc.Macro("earth-polar-radius", "6356752 /m anno"),
	zc.Op("earth-radius", ops.EarthRadiusFn),
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
	zc.Op("eval", ops.Eval, zc.Str),
	zc.Macro("exa", "1e18"),
	zc.Macro("exbi", "2 60 pow"),
	zc.GenOp("exp",
		zc.Func(ops.ExpComplex, zc.Complex),
		zc.Func(ops.ExpFloat, zc.Float),
	),
	zc.Macro("f-c", "32 sub 5 9 div mul /°C anno"),
	zc.Macro("fact", "factorial"),
	zc.Op("factorial", ops.Factorial, zc.BigInt),
	zc.Macro("false", "[false]"),
	zc.Macro("femto", "1e-15"),
	zc.Macro("fib", "fibonacci"),
	zc.Op("fibonacci", ops.Fibonacci, zc.Int),
	zc.Op("filter", ops.Filter),
	zc.GenOp("floor",
		zc.Func(zc.NoOp, zc.BigInt),
		zc.Func(ops.FloorDecimal, zc.Decimal),
		zc.Func(ops.FloorFloat, zc.Float),
	),
	zc.Op("fold", ops.Fold),
	zc.Macro("ft-m", "0.3048 mul /m anno"),
	zc.Macro("ft-mi", "5280 div /mi anno"),
	zc.Macro("ft-yd", "3 div /yd anno"),
	zc.Macro("g-kg", "1000 div /kg anno"),
	zc.Macro("g-oz", "1 28.349523125 div mul /oz anno"),
	zc.Macro("g-ozt", "1 31.1034768 div mul [oz t] anno"),
	zc.Op("get", ops.Get),
	zc.Macro("gibi", "2 30 pow"),
	zc.Macro("giga", "1e09"),
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
	zc.Macro("hecto", "1e02"),
	zc.Op("hex", ops.HexBigInt, zc.BigInt),
	zc.Op("hours", ops.Hours, zc.Duration),
	zc.Op("hsl-rgb", ops.HSLToRGB, zc.Float, zc.Float, zc.Float),
	zc.Macro("i16-max", "32767"),
	zc.Macro("i16-min", "-32768"),
	zc.Op("i16?", ops.IsInt16, zc.Val),
	zc.Macro("i32-max", "2147483647"),
	zc.Macro("i32-min", "-2147483648"),
	zc.Op("i32?", ops.IsInt32, zc.Val),
	zc.Macro("i64-max", "9223372036854775807"),
	zc.Macro("i64-min", "-9223372036854775808"),
	zc.Op("i64?", ops.IsInt64, zc.Val),
	zc.Macro("i8-max", "127"),
	zc.Macro("i8-min", "-128"),
	zc.Op("i8?", ops.IsInt8, zc.Val),
	zc.Op("ia-max", ops.MaxIntArch),
	zc.Op("ia-min", ops.MinIntArch),
	zc.Op("ia?", ops.IsIntArch, zc.Val),
	zc.Op("imag", ops.Imag, zc.Complex),
	zc.Macro("in-mm", "25.4 mul /mm anno"),
	zc.GenOp("int",
		zc.Func(ops.IntBigInt, zc.BigInt),
		zc.Func(ops.IntBigFloat, zc.BigFloat),
		zc.Func(ops.IntDecimal, zc.Decimal),
		zc.Func(ops.IntRational, zc.Rational),
	),
	zc.Op("int?", ops.IsInt, zc.Val),
	zc.Op("inv", ops.Inv, zc.Rational),
	zc.Op("is", ops.Is, zc.Str, zc.Str),
	zc.Op("join", ops.Join),
	zc.Op("join-bits", ops.JoinBits),
	zc.Macro("joinb", "join-bits"),
	zc.Macro("k-c", "273.15 sub /°C anno"),
	zc.Macro("kg-g", "1000 mul /kg anno"),
	zc.Macro("kg-lb", "1 0.45359237 div mul /lb anno"),
	zc.Macro("kibi", "2 10 pow"),
	zc.Macro("kilo", "1e03"),
	zc.Macro("km-m", "1000 mul /m anno"),
	zc.Macro("km-mi", "0.62137119 mul /mi anno"),
	zc.Macro("km-nmi", "0.539957 mul /nmi anno"),
	zc.Macro("lb-kg", "0.45359237 mul /kg anno"),
	zc.Op("left", ops.Left, zc.Str, zc.Int),
	zc.Macro("left-shift", "lsh"),
	zc.Op("len", ops.Len, zc.Str),
	zc.Macro("less-than", "lt"),
	zc.Macro("less-than-or-equal", "lte"),
	zc.Op("local-zone", ops.LocalZone, zc.Str),
	zc.Op("local-zone=", ops.LocalZoneGet),
	zc.GenOp("log",
		zc.Func(ops.LogComplex, zc.Complex),
		zc.Func(ops.LogFloat, zc.Float),
	),
	zc.GenOp("log10",
		zc.Func(ops.Log10Complex, zc.Complex),
		zc.Func(ops.Log10Float, zc.Float),
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
	zc.Macro("m-ft", "0.3048 div /ft anno"),
	zc.Macro("m-km", "1000 div /km anno"),
	zc.Macro("m-nmi", "0.000539957 mul /nmi anno"),
	zc.Macro("m-yd", "0.9144 div /yd anno"),
	zc.Op("map", ops.Map),
	zc.Op("md5", ops.Md5, zc.Str),
	zc.Macro("mebi", "2 20 pow"),
	zc.Macro("mega", "1e06"),
	zc.Macro("mi-ft", "5280 mul /ft anno"),
	zc.Macro("mi-km", "1.609344 mul /km anno"),
	zc.Macro("mi-nmi", "0.868976 mul /nmi anno"),
	zc.Macro("micro", "1e-06"),
	zc.Macro("milli", "1e-03"),
	zc.GenOp("minutes",
		zc.Func(ops.MinutesDMS, zc.DMS),
		zc.Func(ops.Minutes, zc.Duration),
	),
	zc.Macro("mm-in", "25.4 div /in anno"),
	zc.GenOp("mod",
		zc.Func(ops.ModBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.ModDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.ModFloat, zc.Float, zc.Float),
	),
	zc.GenOp("mul",
		zc.Func(ops.MulBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.MulDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.MulBigFloat, zc.BigFloat, zc.BigFloat),
		zc.Func(ops.MulFloat, zc.Float, zc.Float),
		zc.Func(ops.MulRational, zc.Rational, zc.Rational),
		zc.Func(ops.MulComplex, zc.Complex, zc.Complex),
	),
	zc.Op("mul-i16", ops.MulInt16, zc.Int16, zc.Int16),
	zc.Op("mul-i32", ops.MulInt32, zc.Int32, zc.Int32),
	zc.Op("mul-i64", ops.MulInt64, zc.Int64, zc.Int64),
	zc.Op("mul-i8", ops.MulInt8, zc.Int8, zc.Int8),
	zc.Op("mul-ia", ops.MulIntArch, zc.Int, zc.Int),
	zc.Op("mul-u16", ops.MulUint16, zc.Uint16, zc.Uint16),
	zc.Op("mul-u32", ops.MulUint32, zc.Uint32, zc.Uint32),
	zc.Op("mul-u64", ops.MulUint64, zc.Uint64, zc.Uint64),
	zc.Op("mul-u8", ops.MulUint8, zc.Uint8, zc.Uint8),
	zc.Op("mul-ua", ops.MulUintArch, zc.Uint, zc.Uint),
	zc.Macro("n", "size"),
	zc.Macro("nano", "1e-09"),
	zc.GenOp("neg",
		zc.Func(ops.NegBigInt, zc.BigInt),
		zc.Func(ops.NegDecimal, zc.Decimal),
		zc.Func(ops.NegBigFloat, zc.BigFloat),
		zc.Func(ops.NegFloat, zc.Float),
		zc.Func(ops.NegRational, zc.Rational),
	),
	zc.Macro("nmi-km", "1.852 mul /km anno"),
	zc.Macro("nmi-m", "1852 mul /m anno"),
	zc.Macro("nmi-mi", "1.15078 mul /mi anno"),
	zc.Op("no-anno", ops.NoAnno, zc.Val),
	zc.Macro("no-annotation", "no-anno"),
	zc.Macro("noa", "no-anno"),
	zc.GenOp("not",
		zc.Func(ops.NotBool, zc.Bool),
		zc.Func(ops.NotBigInt, zc.BigInt),
	),
	zc.Op("now", ops.Now),
	zc.Op("now-", ops.NowRestore),
	zc.Op("now=", ops.NowSet, zc.DateTime),
	zc.Op("num", ops.Num, zc.Rational),
	zc.Op("oct", ops.Oct, zc.BigInt),
	zc.GenOp("or",
		zc.Func(ops.OrBool, zc.Bool, zc.Bool),
		zc.Func(ops.OrBigInt, zc.BigInt, zc.BigInt),
	),
	zc.Macro("oz-g", "28.349523125 mul /g anno"),
	zc.Macro("ozt-g", "31.1034768 mul /g anno"),
	zc.Macro("pebi", "2 50 pow"),
	zc.Macro("peta", "1e15"),
	zc.Op("phase", ops.PhaseComplex, zc.Complex),
	zc.Macro("pi", "3.14159265358979323846264338327950288419716939937510582097494459"),
	zc.Macro("pico", "1e-12"),
	zc.Op("polar", ops.PolarComplex, zc.Complex),
	zc.GenOp("pow",
		zc.Func(ops.PowBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.PowFloat, zc.Float, zc.Float),
		zc.Func(ops.PowComplex, zc.Complex, zc.Complex),
	),
	zc.Macro("prod", "[mul] fold"),
	zc.Op("proj", ops.Proj, zc.Float, zc.Float, zc.Str, zc.Str),
	zc.Macro("q-int", "quo-int"),
	zc.Macro("qr-int", "quo-rem-int"),
	zc.Macro("quecto", "1e-30"),
	zc.Macro("quetta", "1e30"),
	zc.Op("quo-int", ops.QuoBigInt, zc.BigInt, zc.BigInt),
	zc.Op("quo-rem-int", ops.QuoRemBigInt, zc.BigInt, zc.BigInt),
	zc.Macro("r", "round"),
	zc.Macro("rad-deg", "180 pi div mul"),
	zc.Op("rand", ops.Rand),
	zc.Op("rand-choice", ops.RandChoice, zc.Val),
	zc.Op("rand-int", ops.RandInt, zc.Int),
	zc.Op("rand-seed", ops.RandSeed, zc.Int64),
	zc.Op("rand-seed=", ops.RandSeedGet),
	zc.Macro("rat", "rational"),
	zc.GenOp("rational",
		zc.Func(ops.RationalBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.RationalFloat, zc.Float),
	),
	zc.Op("rational?", ops.RationalIs, zc.Str),
	zc.Macro("re", "recall"),
	zc.Op("real", ops.Real, zc.Complex),
	zc.Op("recall", ops.Recall),
	zc.Op("rect", ops.RectComplex, zc.Float, zc.Float),
	zc.Macro("reduce", "fold"),
	zc.GenOp("rem",
		zc.Func(ops.RemBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.RemFloat, zc.Float, zc.Float),
	),
	zc.Op("repeat", ops.Repeat, zc.Val, zc.Int),
	zc.Macro("rev", "reverse"),
	zc.Op("reverse", ops.Reverse),
	zc.Op("rgb-cmyk", ops.RGBToCMYK, zc.Uint8, zc.Uint8, zc.Uint8),
	zc.Op("rgb-hsl", ops.RGBToHSL, zc.Uint8, zc.Uint8, zc.Uint8),
	zc.Op("right", ops.Right, zc.Str, zc.Int),
	zc.Macro("right-shift", "rsh"),
	zc.Op("roll", ops.Roll, zc.Str),
	zc.Macro("ronna", "1e27"),
	zc.Macro("ronto", "1e-27"),
	zc.Macro("rot13", "rotate-13"),
	zc.Op("rotate-13", ops.Rot13, zc.Str),
	zc.GenOp("round",
		zc.Func(ops.RoundComplex, zc.Complex, zc.Int),
		zc.Func(ops.RoundDecimal, zc.Decimal, zc.Int),
		zc.Func(ops.RoundFloat, zc.Float, zc.Int),
	),
	zc.Op("rounding-mode", ops.RoundingMode),
	zc.Op("rounding-mode=", ops.RoundingModeGet),
	zc.Op("rsh", ops.Rsh, zc.BigInt, zc.Uint),
	zc.Macro("s", "sub"),
	zc.GenOp("scientific-notation",
		zc.Func(ops.ScientificNotation, zc.Float),
		zc.Func(ops.ScientificNotationBigInt, zc.BigInt),
	),
	zc.GenOp("seconds",
		zc.Func(ops.SecondsDMS, zc.DMS),
		zc.Func(ops.Seconds, zc.Duration),
	),
	zc.Macro("seq", "sequence"),
	zc.Op("sequence", ops.Sequence, zc.BigInt, zc.BigInt),
	zc.Op("set", ops.Set),
	zc.Op("sha1", ops.Sha1, zc.Str),
	zc.Op("sha1hmac", ops.Sha1Hmac, zc.Str, zc.Str),
	zc.Op("sha224", ops.Sha224, zc.Str),
	zc.Op("sha224hmac", ops.Sha224Hmac, zc.Str, zc.Str),
	zc.Op("sha256", ops.Sha256, zc.Str),
	zc.Op("sha256hmac", ops.Sha256Hmac, zc.Str, zc.Str),
	zc.Op("sha384", ops.Sha384, zc.Str),
	zc.Op("sha384hmac", ops.Sha384Hmac, zc.Str, zc.Str),
	zc.Op("sha512", ops.Sha512, zc.Str),
	zc.Op("sha512hmac", ops.Sha512Hmac, zc.Str, zc.Str),
	zc.Op("shuffle", ops.Shuffle),
	zc.GenOp("sign",
		zc.Func(ops.SignBigInt, zc.BigInt),
		zc.Func(ops.SignDecimal, zc.Decimal),
		zc.Func(ops.SignBigFloat, zc.BigFloat),
		zc.Func(ops.SignFloat, zc.Float),
		zc.Func(ops.SignRational, zc.Rational),
	),
	zc.GenOp("sin",
		zc.Func(ops.SinComplex, zc.Complex),
		zc.Func(ops.SinFloat, zc.Float),
	),
	zc.GenOp("sinh",
		zc.Func(ops.SinhComplex, zc.Complex),
		zc.Func(ops.SinhFloat, zc.Float),
	),
	zc.Op("size", ops.Size),
	zc.Macro("sn", "scientific-notation"),
	zc.Op("sort", ops.Sort),
	zc.Op("sort-str", ops.SortStr),
	zc.Op("split", ops.Split, zc.Str, zc.Str),
	zc.Op("split-bits", ops.SplitBits, zc.BigInt, zc.Int),
	zc.Macro("splitb", "split-bits"),
	zc.Macro("sq", "dup mul"),
	zc.GenOp("sqrt",
		zc.Func(ops.SqrtFloat, zc.Float),
		zc.Func(ops.SqrtBigFloat, zc.BigFloat),
		zc.Func(ops.SqrtComplex, zc.Complex),
	),
	zc.Op("sqrt-int", ops.SqrtBigInt, zc.BigInt),
	zc.Macro("square", "sq"),
	zc.Macro("square-root", "sqrt"),
	zc.Macro("st", "store"),
	zc.Macro("standard-deviation-pop", "var-p sqrt"),
	zc.Macro("standard-deviation-samp", "var-s sqrt"),
	zc.Macro("stdev-p", "standard-deviation-pop"),
	zc.Macro("stdev-s", "standard-deviation-samp"),
	zc.Op("store", ops.Store),
	zc.GenOp("sub",
		zc.Func(ops.SubBigInt, zc.BigInt, zc.BigInt),
		zc.Func(ops.SubDecimal, zc.Decimal, zc.Decimal),
		zc.Func(ops.SubBigFloat, zc.BigFloat, zc.BigFloat),
		zc.Func(ops.SubFloat, zc.Float, zc.Float),
		zc.Func(ops.SubRational, zc.Rational, zc.Rational),
		zc.Func(ops.SubComplex, zc.Complex, zc.Complex),
		zc.Func(ops.SubDuration, zc.Duration, zc.Duration),
		zc.Func(ops.SubDateTimeDuration, zc.DateTime, zc.Duration),
		zc.Func(ops.SubDateTime, zc.DateTime, zc.DateTime),
	),
	zc.Op("sub-i16", ops.SubInt16, zc.Int16, zc.Int16),
	zc.Op("sub-i32", ops.SubInt32, zc.Int32, zc.Int32),
	zc.Op("sub-i64", ops.SubInt64, zc.Int64, zc.Int64),
	zc.Op("sub-i8", ops.SubInt8, zc.Int8, zc.Int8),
	zc.Op("sub-ia", ops.SubIntArch, zc.Int, zc.Int),
	zc.Op("sub-u16", ops.SubUint16, zc.Uint16, zc.Uint16),
	zc.Op("sub-u32", ops.SubUint32, zc.Uint32, zc.Uint32),
	zc.Op("sub-u64", ops.SubUint64, zc.Uint64, zc.Uint64),
	zc.Op("sub-u8", ops.SubUint8, zc.Uint8, zc.Uint8),
	zc.Op("sub-ua", ops.SubUintArch, zc.Uint, zc.Uint),
	zc.Macro("sum", "[add] fold"),
	zc.Macro("sw", "swap"),
	zc.Op("swap", ops.Swap, zc.Val, zc.Val),
	zc.Op("take", ops.Take),
	zc.GenOp("tan",
		zc.Func(ops.TanComplex, zc.Complex),
		zc.Func(ops.TanFloat, zc.Float),
	),
	zc.GenOp("tanh",
		zc.Func(ops.TanhComplex, zc.Complex),
		zc.Func(ops.TanhFloat, zc.Float),
	),
	zc.Macro("tebi", "2 40 pow"),
	zc.Macro("terra", "1e12"),
	zc.Op("time", ops.Time, zc.DateTime),
	zc.Op("timezone", ops.TimeZone, zc.DateTime, zc.Str),
	zc.Op("tone", ops.Tone, zc.Str, zc.Int),
	zc.Macro("top", "1 take"),
	zc.Macro("true", "[true]"),
	zc.Macro("tz", "timezone"),
	zc.Macro("u16-max", "65535"),
	zc.Op("u16?", ops.IsUint16, zc.Val),
	zc.Macro("u32-max", "4294967295"),
	zc.Op("u32?", ops.IsUint32, zc.Val),
	zc.Macro("u64-max", "18446744073709551615"),
	zc.Op("u64?", ops.IsUint64, zc.Val),
	zc.Macro("u8-max", "255"),
	zc.Op("u8?", ops.IsUint8, zc.Val),
	zc.Macro("u8de", "utf8-decode"),
	zc.Macro("u8en", "utf8-encode"),
	zc.Op("ua-max", ops.MaxUintArch),
	zc.Op("ua?", ops.IsUintArch, zc.Val),
	zc.Macro("unesc", "unescape"),
	zc.Op("unescape", ops.Unescape, zc.Str),
	zc.Op("up", ops.Up),
	zc.Op("upper", ops.Upper, zc.Str),
	zc.Op("utf8-decode", ops.UTF8Decode, zc.BigInt),
	zc.Op("utf8-encode", ops.UTF8Encode, zc.Str),
	zc.Macro("var-p", "variance-pop"),
	zc.Macro("var-s", "variance-samp"),
	zc.GenOp("variance-pop",
		zc.Func(ops.VariancePop),
	),
	zc.GenOp("variance-samp",
		zc.Func(ops.VarianceSamp),
	),
	zc.Op("version", ops.Version),
	zc.Op("xor", ops.Xor, zc.BigInt, zc.BigInt),
	zc.Macro("yd-ft", "3 mul /ft anno"),
	zc.Macro("yd-m", "0.9144 mul /m anno"),
	zc.Macro("yocto", "1e-24"),
	zc.Macro("yotta", "1e24"),
	zc.Macro("zepto", "1e-21"),
	zc.Macro("zetta", "1e21"),
	zc.Macro("π", "pi"),
}