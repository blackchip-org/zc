package zlib

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/native"
)

var All = []zc.ModuleDef{
	Assert,
	Basic,
	Bool,
	Runtime,
	Str,
	Test,
}

var (
	PreludeCLI = []string{
		"basic",
		"str",
	}
	PreludeDev = []string{
		"basic",
		"bool",
		"str",
	}
)

var (
	Assert = zc.ModuleDef{
		Name:       "assert",
		ScriptPath: "zc:zlib/assert.zc",
	}
	Basic = zc.ModuleDef{
		Name:       "basic",
		ScriptPath: "zc:zlib/basic.zc",
		Natives: map[string]zc.CalcFunc{
			"+":     native.Add,
			"/":     native.DivDec,
			"*":     native.Mul,
			"-":     native.Sub,
			"a":     native.Add,
			"d":     native.DivDec,
			"m":     native.Mul,
			"s":     native.Sub,
			"add":   native.Add,
			"add-d": native.AddDec,
			"add-i": native.AddBigInt,
			"div":   native.DivDec,
			"div-d": native.DivDec,
			"div-i": native.DivBigInt,
			"mul":   native.Mul,
			"mul-d": native.MulDec,
			"mul-i": native.MulBigInt,
			"pow":   native.Pow,
			"pow-d": native.PowDec,
			"pow-i": native.PowBigInt,
			"sub":   native.Sub,
			"sub-d": native.SubDec,
			"sub-i": native.SubBigInt,
		},
	}
	Bool = zc.ModuleDef{
		Name:       "bool",
		ScriptPath: "zc:zlib/bool.zc",
		Natives: map[string]zc.CalcFunc{
			"eq":  native.Eq,
			"neq": native.Neq,
			"gt":  native.Gt,
			"gte": native.Gte,
			"lt":  native.Lt,
			"lte": native.Lte,
			"not": native.Not,
		},
	}
	Runtime = zc.ModuleDef{
		Name: "runtime",
		Natives: map[string]zc.CalcFunc{
			"exports": native.Exports,
			"funcs":   native.Funcs,
		},
	}
	Str = zc.ModuleDef{
		Name: "str",
		Natives: map[string]zc.CalcFunc{
			"len":         native.Len,
			"starts-with": native.StartsWith,
		},
	}
	Test = zc.ModuleDef{
		Name:       "test",
		ScriptPath: "zc:zlib/test.zc",
	}
)
