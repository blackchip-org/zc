package modules

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/native"
)

var Bool = zc.Module{
	Natives: map[string]zc.NativeFn{
		"eq": native.Eq,
	},
}
