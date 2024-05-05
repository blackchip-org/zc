package ints

import (
	"github.com/blackchip-org/zc/v6"
)

func addIntArch(e *zc.OpEnv) {
	x := e.Args[0].(int)
	y := e.Args[1].(int)
	e.Returns = []any{x + y}
}
