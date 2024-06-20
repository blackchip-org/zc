package funcs

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/state"
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

func ModDecimalSS(e *zc.OpEnv) {
	y := zc.DecimalSS.Pop(e)
	x := zc.DecimalSS.Pop(e)
	x = x.Mod(y)
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

func PowDecimalSS(e *zc.OpEnv) {
	conf := state.ForConf(e.State)
	y := zc.DecimalSS.Pop(e)
	x := zc.DecimalSS.Pop(e)
	z, err := x.PowWithPrecision(y, conf.DecPrecAsInt32())
	if err != nil {
		e.Err = err
		return
	}
	zc.DecimalSS.Push(e, z)
}

func RemDecimalSS(e *zc.OpEnv) {
	p := zc.Int32.Pop(e)
	y := zc.DecimalSS.Pop(e)
	x := zc.DecimalSS.Pop(e)
	if y.IsZero() {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	_, r := x.QuoRem(y, p)
	zc.DecimalSS.Push(e, r)
}

func SubDecimalSS(e *zc.OpEnv) {
	y := zc.DecimalSS.Pop(e)
	x := zc.DecimalSS.Pop(e)
	x = x.Sub(y)
	zc.DecimalSS.Push(e, x)
}
