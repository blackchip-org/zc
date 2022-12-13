package zlib

import (
	"github.com/blackchip-org/zc"
)

var All = []zc.ModuleDef{
	Assert,
	Bool,
	FixBool,
	FixMath,
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
	FixBool = zc.ModuleDef{
		Name:       "dec-bool",
		ScriptPath: "zc:zlib/fix-bool.zc",
		Natives: map[string]zc.CalcFunc{
			"eq":  EqFix,
			"gt":  GtFix,
			"gte": GteFix,
			"neq": NeqFix,
			"lt":  LtFix,
			"lte": LteFix,
		},
	}
	FixMath = zc.ModuleDef{
		Name:       "dec-math",
		ScriptPath: "zc:zlib/fix-math.zc",
		Natives: map[string]zc.CalcFunc{
			"abs":   AbsFix,
			"add":   AddFix,
			"ceil":  CeilFix,
			"div":   DivFix,
			"floor": FloorFix,
			"mod":   ModFix,
			"mul":   MulFix,
			"neg":   NegFix,
			"pow":   PowFix,
			"rem":   RemFix,
			"sign":  SignFix,
			"sub":   SubFix,
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
			"div":   DivFix,
			"floor": Floor,
			"mod":   Mod,
			"mul":   Mul,
			"neg":   Neg,
			"pow":   Pow,
			"rem":   Rem,
			"sign":  Sign,
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
