package funcs

import (
	"github.com/blackchip-org/zc/v6"
)

func AddDecimalSS(e *zc.OpEnv) {
	y := zc.DecimalSS.As(e.Pop())
	x := zc.DecimalSS.As(e.Pop())
	e.PushVal(x.Add(y))
}

func DivDecimalSS(e *zc.OpEnv) {
	y := zc.DecimalSS.As(e.Pop())
	x := zc.DecimalSS.As(e.Pop())
	if y.IsZero() {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	e.PushVal(x.Div(y))
}

func MulDecimalSS(e *zc.OpEnv) {
	y := zc.DecimalSS.As(e.Pop())
	x := zc.DecimalSS.As(e.Pop())
	e.PushVal(x.Mul(y))
}

func NegDecimalSS(e *zc.OpEnv) {
	x := zc.DecimalSS.As(e.Pop())
	e.PushVal(x.Neg())
}

func SubDecimalSS(e *zc.OpEnv) {
	y := zc.DecimalSS.As(e.Pop())
	x := zc.DecimalSS.As(e.Pop())
	e.PushVal(x.Sub(y))
}
