package zlib

import (
	"github.com/blackchip-org/zc"
)

var All = []zc.ModuleDef{
	Assert,
	Bool,
	Math,
	Prog,
	Runtime,
	Stack,
	Str,
	Test,
}

var (
	PreludeCLI = []string{
		"math",
		"stack",
		"str",
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
	Math = zc.ModuleDef{
		Name:       "math",
		Include:    true,
		ScriptPath: "zc:zlib/math.zc",
		Natives: map[string]zc.CalcFunc{
			"+":     Add,
			"/":     DivDec,
			"*":     Mul,
			"-":     Sub,
			"a":     Add,
			"d":     DivDec,
			"m":     Mul,
			"s":     Sub,
			"add":   Add,
			"add-d": AddDec,
			"add-i": AddBigInt,
			"div":   DivDec,
			"div-d": DivDec,
			"div-i": DivBigInt,
			"mul":   Mul,
			"mul-d": MulDec,
			"mul-i": MulBigInt,
			"pow":   Pow,
			"pow-d": PowDec,
			"pow-i": PowBigInt,
			"sub":   Sub,
			"sub-d": SubDec,
			"sub-i": SubBigInt,
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
	Stack = zc.ModuleDef{
		Name:       "stack",
		Include:    true,
		ScriptPath: "zc:zlib/stack.zc",
	}
	Runtime = zc.ModuleDef{
		Name: "runtime",
		Natives: map[string]zc.CalcFunc{
			"exports": Exports,
			"funcs":   Funcs,
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
