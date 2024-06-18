package pretty

import (
	"strings"

	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func RepeatString(n int, s string) string {
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteString(s)
	}
	return b.String()
}
