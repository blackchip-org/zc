package zc

import "cmp"

func Clamp[T cmp.Ordered](z T, a T, b T) T {
	if z < a {
		return a
	}
	if z > b {
		return b
	}
	return z
}
