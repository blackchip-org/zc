package zlib

import (
	"github.com/blackchip-org/zc"
)

var All = []zc.ModuleDef{
	Assert,
	Bool,
	BoolBigInt,
	BoolFixed,
	Conf,
	Dev,
	Math,
	MathBigInt,
	MathFixed,
	MathFloat,
	Io,
	Prog,
	Runtime,
	Stack,
	Str,
	StrBool,
	Test,
	Time,
	Tz,
}

var (
	PreludeUser = []string{
		"conf",
		"math",
		"stack",
		"str", // TODO: This should only be in dev
	}
	PreludeDev = []string{
		"dev",
		"stack",
		"bool",
		"conf",
		"math",
		"str",
	}
)

var (
	Assert = zc.ModuleDef{
		Name:       "assert",
		ScriptPath: "zc:zlib/assert.zc",
	}
	Bool = zc.ModuleDef{
		Name:       "bool",
		Include:    true,
		ScriptPath: "zc:zlib/bool.zc",
		Natives: map[string]zc.CalcFunc{
			"and": And,
			"eq":  Eq,
			"neq": Neq,
			"gt":  Gt,
			"gte": Gte,
			"lt":  Lt,
			"lte": Lte,
			"not": Not,
			"or":  Or,
		},
	}
	BoolBigInt = zc.ModuleDef{
		Name:       "bool.bigint",
		ScriptPath: "zc:zlib/bool-bigint.zc",
		Natives: map[string]zc.CalcFunc{
			"eq":  EqBigInt,
			"gt":  GtBigInt,
			"gte": GteBigInt,
			"neq": NeqBigInt,
			"lt":  LtBigInt,
			"lte": LteBigInt,
		},
	}
	BoolFixed = zc.ModuleDef{
		Name:       "bool.fixed",
		ScriptPath: "zc:zlib/bool-fixed.zc",
		Natives: map[string]zc.CalcFunc{
			"eq":  EqFixed,
			"gt":  GtFixed,
			"gte": GteFixed,
			"neq": NeqFixed,
			"lt":  LtFixed,
			"lte": LteFixed,
		},
	}
	Conf = zc.ModuleDef{
		Name:       "conf",
		ScriptPath: "zc:zlib/conf.zc",
		Natives: map[string]zc.CalcFunc{
			"auto-currency":  AutoCurrency,
			"auto-currency=": AutoCurrencyGet,
			"int-format":     IntFormat,
			"int-format=":    IntFormatGet,
			"min-digits":     MinDigits,
			"min-digits=":    MinDigitsGet,
			"precision":      Precision,
			"precision=":     PrecisionGet,
			"point":          Point,
			"point=":         PointGet,
			"rounding-mode":  RoundingMode,
			"rounding-mode=": RoundingModeGet,
		},
	}
	Dev = zc.ModuleDef{
		Name:       "dev",
		ScriptPath: "zc:zlib/dev.zc",
		Natives: map[string]zc.CalcFunc{
			"abort":     Abort,
			"eval":      Eval,
			"exit":      Exit,
			"nothing":   Nothing,
			"trace":     Trace,
			"trace-off": TraceOff,
			"undef":     Undef,
		},
	}
	Math = zc.ModuleDef{
		Name:       "math",
		Include:    true,
		ScriptPath: "zc:zlib/math.zc",
		Natives: map[string]zc.CalcFunc{
			"abs":   Abs,
			"add":   Add,
			"ceil":  Ceil,
			"div":   DivFixed,
			"floor": Floor,
			"mod":   Mod,
			"mul":   Mul,
			"neg":   Neg,
			"pow":   Pow,
			"rem":   Rem,
			"round": Round,
			"sign":  Sign,
			"sqrt":  Sqrt,
			"sub":   Sub,
		},
	}
	MathBigInt = zc.ModuleDef{
		Name:       "math.bigint",
		ScriptPath: "zc:zlib/math-bigint.zc",
		Natives: map[string]zc.CalcFunc{
			"abs":   AbsBigInt,
			"add":   AddBigInt,
			"ceil":  CeilBigInt,
			"div":   DivBigInt,
			"floor": FloorBigInt,
			"mod":   ModBigInt,
			"mul":   MulBigInt,
			"neg":   NegBigInt,
			"pow":   PowBigInt,
			"rem":   RemBigInt,
			"sign":  SignBigInt,
			"sub":   SubBigInt,
		},
	}
	MathFixed = zc.ModuleDef{
		Name:       "math.fixed",
		ScriptPath: "zc:zlib/math-fixed.zc",
		Natives: map[string]zc.CalcFunc{
			"abs":   AbsFixed,
			"add":   AddFixed,
			"ceil":  CeilFixed,
			"div":   DivFixed,
			"floor": FloorFixed,
			"mod":   ModFixed,
			"mul":   MulFixed,
			"neg":   NegFixed,
			"pow":   PowFixed,
			"rem":   RemFixed,
			"sign":  SignFixed,
			"sub":   SubFixed,
		},
	}
	MathFloat = zc.ModuleDef{
		Name:       "math.float",
		ScriptPath: "zc:zlib/math-float.zc",
		Natives: map[string]zc.CalcFunc{
			"sqrt": SqrtFloat,
		},
	}
	Io = zc.ModuleDef{
		Name:       "io",
		ScriptPath: "zc:zlib/io.zc",
		Natives: map[string]zc.CalcFunc{
			"print": Print,
		},
	}
	Prog = zc.ModuleDef{
		Name:       "prog",
		Include:    true,
		ScriptPath: "zc:zlib/prog.zc",
		Natives: map[string]zc.CalcFunc{
			"and": AndBitwise,
			"bin": Bin,
			"bit": Bit,
			"dec": Dec,
			"hex": Hex,
			"len": LenBitwise,
			"lsh": Lsh,
			"not": NotBitwise,
			"oct": Oct,
			"or":  OrBitwise,
			"rsh": Rsh,
			"xor": Xor,
		},
	}
	Runtime = zc.ModuleDef{
		Name:       "runtime",
		ScriptPath: "zc:zlib/runtime.zc",
		Natives: map[string]zc.CalcFunc{
			"exports": Exports,
			"funcs":   Funcs,
		},
	}
	Stack = zc.ModuleDef{
		Name:       "stack",
		Include:    true,
		ScriptPath: "zc:zlib/stack.zc",
		Natives: map[string]zc.CalcFunc{
			"n": N,
		},
	}
	StrBool = zc.ModuleDef{
		Name:       "str-bool",
		ScriptPath: "zc:zlib/str-bool.zc",
		Natives: map[string]zc.CalcFunc{
			"eq":  EqStr,
			"gt":  GtStr,
			"gte": GteStr,
			"neq": NeqStr,
			"lt":  LtStr,
			"lte": LteStr,
		},
	}
	Str = zc.ModuleDef{
		Name:       "str",
		Include:    true,
		ScriptPath: "zc:zlib/str.zc",
		Natives: map[string]zc.CalcFunc{
			"len":         Len,
			"starts-with": StartsWith,
		},
	}
	Test = zc.ModuleDef{
		Name:       "test",
		ScriptPath: "zc:zlib/test.zc",
	}
	Time = zc.ModuleDef{
		Name:       "time",
		Include:    true,
		ScriptPath: "zc:zlib/time.zc",
		Natives: map[string]zc.CalcFunc{
			"after":      After,
			"date-time":  DateTime,
			"formats=":   FormatsGet,
			"local":      Local,
			"local=":     LocalGet,
			"now":        Now,
			"ord":        Ord,
			"in":         In,
			"travel":     Travel,
			"travel-end": TravelEnd,
		},
		Init: InitTime,
	}
	Tz = zc.ModuleDef{
		Name:       "tz",
		ScriptPath: "zc:zlib/tz.zc",
	}
)
