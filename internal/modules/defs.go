package modules

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/native"
)

var All = []zc.ModuleDef{
	Assert,
	Basic,
	Bool,
}

var (
	PreludeCLI = []string{
		"basic",
		"bool",
	}
)

var (
	Assert = zc.ModuleDef{
		Name:    "assert",
		Scripts: []string{"internal/modules/assert.zc"},
	}
	Basic = zc.ModuleDef{
		Name:    "basic",
		Scripts: []string{"internal/modules/basic.zc"},
		Natives: map[string]zc.CalcFunc{
			"a":     native.Add,
			"m":     native.Mul,
			"s":     native.Sub,
			"add":   native.Add,
			"add-d": native.AddDec,
			"add-i": native.AddBigInt,
			"mul":   native.Mul,
			"mul-d": native.MulDec,
			"mul-i": native.MulBigInt,
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
)
