package zlib

import (
	"github.com/blackchip-org/zc"
)

var All = []zc.ModuleDef{
	Bool,
	BoolBigInt,
	BoolFixed,
	Conf,
	Dev,
	Dict,
	Fn,
	Math,
	MathBigInt,
	MathFixed,
	MathFloat,
	Io,
	Prog,
	Rand,
	Runtime,
	Sci,
	SI,
	Stack,
	Str,
	StrBool,
	Test,
	Time,
	Tz,
	Unit,
	Zc,
}

var (
	Preload = []string{
		"zc",
		"dev",
		"stack",
		"bool",
		"conf",
		"math",
		"str",
		"dict",
		"fn",
	}

	PreludeUser = []string{
		"bool",
		"conf",
		"math",
		"stack",
		"str", // TODO: This should only be in dev
		"fn",
		"zc",
	}
	// Order here needs to be sorted based on dependencies. Do not put in
	// alphabetical order. Modules to the top of the list do not have access
	// to functions found at the bottom of the list.
	//
	// Also, update dev mode when this changes
	PreludeDev = []string{
		"bool",
		"dev",
		"stack",
		"conf",
		"math",
		"str",
		"dict",
		"fn",
	}
)

var (
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
			"locale":         Locale,
			"locale=":        LocaleGet,
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
		Include:    true,
		Natives: map[string]zc.CalcFunc{
			"abort":     Abort,
			"exit":      Exit,
			"nothing":   Nothing,
			"trace":     Trace,
			"trace-off": TraceOff,
			"undef":     Undef,
		},
	}
	Dict = zc.ModuleDef{
		Name:       "dict",
		ScriptPath: "zc:zlib/dict.zc",
	}
	Fn = zc.ModuleDef{
		Name:       "fn",
		ScriptPath: "zc:zlib/fn.zc",
		Include:    true,
		Natives: map[string]zc.CalcFunc{
			"eval": Eval,
		},
	}
	Geom = zc.ModuleDef{
		Name:       "geom",
		ScriptPath: "zc:zlib/geom.zc",
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
	Rand = zc.ModuleDef{
		Name:       "rand",
		ScriptPath: "zc:zlib/rand.zc",
		Init:       InitRand,
		Natives: map[string]zc.CalcFunc{
			"float":   FloatRand,
			"int":     IntRand,
			"seed":    Seed,
			"seed=":   SeedGet,
			"shuffle": Shuffle,
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
	Sci = zc.ModuleDef{
		Name:       "sci",
		Include:    true,
		ScriptPath: "zc:zlib/sci.zc",
		Natives: map[string]zc.CalcFunc{
			"scientific-notation": ScientificNotation,
		},
	}
	SI = zc.ModuleDef{
		Name:       "si",
		Include:    true,
		ScriptPath: "zc:zlib/si.zc",
	}
	Stack = zc.ModuleDef{
		Name:       "stack",
		Include:    true,
		ScriptPath: "zc:zlib/stack.zc",
		Natives: map[string]zc.CalcFunc{
			"at": At,
			"n":  N,
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
		Natives: map[string]zc.CalcFunc{
			"file": TestFile,
		},
	}
	Time = zc.ModuleDef{
		Name:       "time",
		Include:    true,
		ScriptPath: "zc:zlib/time.zc",
		Init:       InitTime,
		Natives: map[string]zc.CalcFunc{
			"add-duration":      AddDuration,
			"date":              Date,
			"date-layout":       DateLayout,
			"date-layout=":      DateLayoutGet,
			"date-time":         DateTime,
			"date-time-layout":  DateTimeLayout,
			"date-time-layout=": DateTimeLayoutGet,
			"day-year":          DayYear,
			"local":             Local,
			"local=":            LocalGet,
			"now":               Now,
			"subtract-time":     SubtractTime,
			"time":              Time_,
			"time-layout":       TimeLayout,
			"time-layout=":      TimeLayoutGet,
			"time-zone":         TimeZone,
			"travel":            Travel,
			"travel-end":        TravelEnd,
		},
	}
	Tz = zc.ModuleDef{
		Name:       "tz",
		ScriptPath: "zc:zlib/tz.zc",
	}
	Unit = zc.ModuleDef{
		Name:       "unit",
		ScriptPath: "zc:zlib/unit.zc",
	}
	Zc = zc.ModuleDef{
		Name:       "zc",
		Include:    true,
		ScriptPath: "zc:zlib/zc.zc",
		Natives: map[string]zc.CalcFunc{
			"version": Version,
		},
	}
)
