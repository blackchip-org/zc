package funcs

import (
	"github.com/blackchip-org/zc/v6"
)

func AddDecimalSS(e *zc.OpEnv) {
	y := zc.DecimalSS.Pop(e)
	x := zc.DecimalSS.Pop(e)
	x = x.Add(y)
	zc.DecimalSS.Push(e, x)
}

func DivDecimalSS(e *zc.OpEnv) {
	y := zc.DecimalSS.Pop(e)
	x := zc.DecimalSS.Pop(e)
	if y.IsZero() {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	x = x.Div(y)
	zc.DecimalSS.Push(e, x)
}

func MulDecimalSS(e *zc.OpEnv) {
	y := zc.DecimalSS.Pop(e)
	x := zc.DecimalSS.Pop(e)
	x = x.Mul(y)
	zc.DecimalSS.Push(e, x)
}

func NegDecimalSS(e *zc.OpEnv) {
	x := zc.DecimalSS.Pop(e)
	x = x.Neg()
	zc.DecimalSS.Push(e, x)
}

func SubDecimalSS(e *zc.OpEnv) {
	y := zc.DecimalSS.Pop(e)
	x := zc.DecimalSS.Pop(e)
	x = x.Sub(y)
	zc.DecimalSS.Push(e, x)
}
