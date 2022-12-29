package zlib

import (
	"github.com/blackchip-org/zc"
)

var All = []zc.ModuleDef{
	Assert,
	Bool,
	Conf,
	Dev,
	FixedBool,
	FixedMath,
	Math,
	IntBool,
	IntMath,
	Io,
	Prog,
	Runtime,
	Stack,
	Str,
	StrBool,
	Test,
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
	Conf = zc.ModuleDef{
		Name:       "conf",
		ScriptPath: "zc:zlib/conf.zc",
		Natives: map[string]zc.CalcFunc{
			"places":  Places,
			"places=": PlacesGet,
			"round":   RoundMode,
			"round=":  RoundModeGet,
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
	FixedBool = zc.ModuleDef{
		Name:       "fixed-bool",
		ScriptPath: "zc:zlib/fixed-bool.zc",
		Natives: map[string]zc.CalcFunc{
			"eq":  EqFixed,
			"gt":  GtFixed,
			"gte": GteFixed,
			"neq": NeqFixed,
			"lt":  LtFixed,
			"lte": LteFixed,
		},
	}
	FixedMath = zc.ModuleDef{
		Name:       "fixed-math",
		ScriptPath: "zc:zlib/fixed-math.zc",
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
	FloatMath = zc.ModuleDef{
		Name:       "float-math",
		ScriptPath: "zc:zlib/float-math.zc",
		Natives: map[string]zc.CalcFunc{
			"sqrt": SqrtFloat,
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
	IntBool = zc.ModuleDef{
		Name:       "int-bool",
		ScriptPath: "zc:zlib/int-bool.zc",
		Natives: map[string]zc.CalcFunc{
			"eq":  EqBigInt,
			"gt":  GtBigInt,
			"gte": GteBigInt,
			"neq": NeqBigInt,
			"lt":  LtBigInt,
			"lte": LteBigInt,
		},
	}
	IntMath = zc.ModuleDef{
		Name:       "int-math",
		ScriptPath: "zc:zlib/int-math.zc",
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
)
