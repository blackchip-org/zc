package funcs

import (
	"github.com/blackchip-org/zc/v6"
)

func Clear(e *zc.OpEnv) {
	e.Clear()
}

func Drop(e *zc.OpEnv) {
	e.Pop()
}

func DupBigInt(e *zc.OpEnv) {
	x := zc.BigInt.Pop(e)
	x2 := zc.BigInt.New()
	x2.Set(x)
	zc.BigInt.Push(e, x)
	zc.BigInt.Push(e, x2)
}

func DupDecimal(e *zc.OpEnv) {
	x := zc.Decimal.Pop(e)
	x2 := zc.Decimal.New()
	x2.Set(x)
	zc.Decimal.Push(e, x)
	zc.Decimal.Push(e, x2)
}

func Rotate(e *zc.OpEnv) {
	e.Stack.Rotate()
}

func TuckBigInt(e *zc.OpEnv) {
	y := e.Pop()
	x := e.Pop()
	y2 := zc.BigInt.New()
	y2.Set(zc.BigInt.As(y.Val))
	e.PushVal(y2)
	e.Push(x, y)
}

func TuckDecimal(e *zc.OpEnv) {
	y := e.Pop()
	x := e.Pop()
	y2 := zc.Decimal.New()
	y2.Set(zc.Decimal.As(y.Val))
	e.PushVal(y2)
	e.Push(x, y)
}
