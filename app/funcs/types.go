package funcs

import (
	"math"

	"github.com/blackchip-org/zc/v6"
)

func MaxUint(e *zc.OpEnv) {
	zc.Uint.Push(e, math.MaxUint)
}

func Types(e *zc.OpEnv) {
	for i := 0; i < e.Len(); i++ {
		item := e.Stack.Get(i)
		item.Label = zc.TypeOf(item.Val).AppName()
		e.Stack.Set(i, item)
	}
}
