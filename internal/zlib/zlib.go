package zlib

import (
	"github.com/blackchip-org/zc"
)

var All = []zc.ModuleDef{
	Assert,
	Bool,
	DecBool,
	DecMath,
	Math,
	IntBool,
	IntMath,
	Prog,
	Runtime,
	Stack,
	Str,
	StrBool,
	Test,
}

var (
	PreludeCLI = []string{
		"math",
		"stack",
		"str", // TODO: This should only be in dev
	}
	PreludeDev = []string{
		"bool",
		"math",
		"stack",
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
	DecBool = zc.ModuleDef{
		Name:       "dec-bool",
		ScriptPath: "zc:zlib/dec-bool.zc",
		Natives: map[string]zc.CalcFunc{
			"eq":  EqDec,
			"gt":  GtDec,
			"gte": GteDec,
			"neq": NeqDec,
			"lt":  LtDec,
			"lte": LteDec,
		},
	}
	DecMath = zc.ModuleDef{
		Name:       "dec-math",
		ScriptPath: "zc:zlib/dec-math.zc",
		Natives: map[string]zc.CalcFunc{
			"add": AddDec,
			"div": DivDec,
			"mul": MulDec,
			"pow": PowDec,
			"sub": SubDec,
		},
	}
	Math = zc.ModuleDef{
		Name:       "math",
		Include:    true,
		ScriptPath: "zc:zlib/math.zc",
		Natives: map[string]zc.CalcFunc{
			"+":   Add,
			"/":   DivDec,
			"*":   Mul,
			"-":   Sub,
			"a":   Add,
			"d":   DivDec,
			"m":   Mul,
			"s":   Sub,
			"add": Add,
			"div": DivDec,
			"mul": Mul,
			"pow": Pow,
			"sub": Sub,
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
			"add": AddBigInt,
			"div": DivBigInt,
			"mul": MulBigInt,
			"pow": PowBigInt,
			"sub": SubBigInt,
		},
	}
	Prog = zc.ModuleDef{
		Name:    "prog",
		Include: true,
		Natives: map[string]zc.CalcFunc{
			"bin": Bin,
			"hex": Hex,
			"oct": Oct,
		},
	}
	Runtime = zc.ModuleDef{
		Name: "runtime",
		Natives: map[string]zc.CalcFunc{
			"exports": Exports,
			"funcs":   Funcs,
		},
	}
	Stack = zc.ModuleDef{
		Name:       "stack",
		Include:    true,
		ScriptPath: "zc:zlib/stack.zc",
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
		Name:    "str",
		Include: true,
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
