//go:build !proj

package geo

import (
	"github.com/blackchip-org/zc"
)

func Transform(env *zc.Env) error {
	return zc.ErrFunctionNotAvailable
}
