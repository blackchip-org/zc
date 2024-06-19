package funcs

import "github.com/blackchip-org/zc/v6"

func Label(e *zc.OpEnv) {
	x := zc.String.Pop(e)
	e.Label(x)
}

func Unit(e *zc.OpEnv) {
	x := zc.String.Pop(e)
	e.Unit(x)
}
