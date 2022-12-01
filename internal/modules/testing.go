package modules

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/native"
)

var Testing = zc.Module{
	Natives: map[string]zc.NativeFn{
		"assert":     native.Assert,
		"test":       native.Test,
		"test-suite": native.TestSuite,
	},
}
