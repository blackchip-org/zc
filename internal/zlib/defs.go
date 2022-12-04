package zlib

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/native"
)

var All = []zc.ModuleDef{
	Assert,
	Basic,
	Bool,
	Str,
}

var (
	PreludeCLI = []string{
		"basic",
		"bool",
		"str",
	}
)

var (
	Assert = zc.ModuleDef{
		Name:       "assert",
		ScriptPath: "zlib/assert.zc",
	}
	Basic = zc.ModuleDef{
		Name:       "basic",
		ScriptPath: "zlib/basic.zc",
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
		Name: "bool",
		Natives: map[string]zc.CalcFunc{
			"eq":  native.Eq,
			"gt":  native.Gt,
			"lt":  native.Lt,
			"not": native.Not,
		},
	}
	Str = zc.ModuleDef{
		Name: "str",
		Natives: map[string]zc.CalcFunc{
			"len": native.Len,
		},
	}
)
