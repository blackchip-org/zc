package modules

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/native"
)

var Basic = zc.Module{
	Natives: map[string]zc.NativeFn{
		"add": native.Add,
	},
}
