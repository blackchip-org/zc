package modules

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/native"
)

var All = []zc.ModuleDef{
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
	Basic = zc.ModuleDef{
		Name:    "basic",
		Scripts: []string{"internal/modules/basic.zc"},
		Natives: map[string]zc.CalcFunc{
			"add": native.Add,
		},
	}
	Bool = zc.ModuleDef{
		Name: "bool",
		Natives: map[string]zc.CalcFunc{
			"add": native.Add,
			"gt":  native.Gt,
			"lt":  native.Lt,
		},
	}
)
