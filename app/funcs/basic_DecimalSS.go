package funcs

import (
	"strings"

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
	if y.IsZero() {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
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
		msg := err.Error()
		switch {
		case strings.HasPrefix(msg, "cannot represent imaginary"):
			e.Err = zc.ErrInvalidArg(e, "root of negative number: %v", x)
		case strings.HasPrefix(msg, "cannot represent undefined"):
			e.Err = zc.ErrUndefined(e)
		case strings.HasPrefix(msg, "cannot represent infinity"):
			e.Err = zc.ErrInfinity(e, 0)
		default:
			e.Err = zc.ErrOp(e, err)
		}
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
