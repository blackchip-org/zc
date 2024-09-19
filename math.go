package zc

import (
	"cmp"
	"math"

	"github.com/blackchip-org/zc/v6/msg"
)

func Clamp[T cmp.Ordered](z T, a T, b T) T {
	if z < a {
		return a
	}
	if z > b {
		return b
	}
	return z
}

func IsFloatErr(c Calc, f float64) bool {
	switch {
	case math.IsInf(f, -1):
		c.Raise(msg.ErrInfinity(-1))
		return true
	case math.IsInf(f, 1):
		c.Raise(msg.ErrInfinity(1))
		return true
	case math.IsNaN(f):
		c.Raise(msg.ErrNotANumber())
		return true
	default:
		return false
	}
}
