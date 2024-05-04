package floats

import "github.com/blackchip-org/zc/v6"

func add(e *zc.OpEnv) {
	x := e.Args[0].(float64)
	y := e.Args[1].(float64)
	e.Returns = []any{x + y}
}
