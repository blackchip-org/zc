//go:build !proj

package zlib

import (
	"github.com/blackchip-org/zc"
)

func Proj(env *zc.Env) error {
	return zc.ErrFunctionNotAvailable
}
