package funcs

import (
	"github.com/blackchip-org/zc/v6"
)

func Clear(e *zc.OpEnv) {
	e.Clear()
}

func DupBigInt(e *zc.OpEnv) {
	x := zc.BigInt.As(e.Top())
	x2 := zc.BigInt.New()
	x2.Set(x)
	e.PushVal(x2)
}

func DupDecimal(e *zc.OpEnv) {
	x := zc.Decimal.As(e.Top())
	x2 := zc.Decimal.New()
	x2.Set(x)
	e.PushVal(x2)
}

func Rotate(e *zc.OpEnv) {
	e.Stack.Rotate()
}

func TuckBigInt(e *zc.OpEnv) {
	y := e.Pop()
	x := e.Pop()
	y2 := zc.BigInt.New()
	y2.Set(zc.BigInt.As(y))
	e.PushVal(y2)
	e.Push(x, y)
}
