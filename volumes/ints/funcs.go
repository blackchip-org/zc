package ints

import (
	"github.com/blackchip-org/zc/v6"
)

func addIntArch(e *zc.OpEnv) {
	x := e.Args[0].(int)
	y := e.Args[1].(int)
	e.Returns = []any{x + y}
}

func int8_(e *zc.OpEnv) {
	x := e.Args[0].(int64)
	e.Returns = []any{int8(x)}
}
