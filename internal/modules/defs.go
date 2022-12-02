package modules

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/native"
)

var All = []zc.ModuleDef{
	Basic,
}

var Basic = zc.ModuleDef{
	Name:    "basic",
	Scripts: []string{"internal/modules/basic.zc"},
	Natives: map[string]zc.NativeFn{
		"add": native.Add,
	},
}
