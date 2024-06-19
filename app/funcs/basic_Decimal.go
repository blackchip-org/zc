package funcs

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/state"
	"github.com/cockroachdb/apd/v3"
)

func AddDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	y := zc.Decimal.As(e.Pop())
	x := zc.Decimal.As(e.Pop())
	defer zc.Decimal.Recycle(y)

	_, err := s.Context.Add(x, x, y)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	e.PushVal(x)
}

func CbrtDecimal(e *zc.OpEnv) {
	var zero apd.Decimal
	s := state.ForDec(e.State)
	x := zc.Decimal.As(e.Pop())

	if x.Cmp(&zero) < 0 {
		e.Err = zc.ErrInvalidArg(e, "%v < 0", x)
		return
	}
	_, err := s.Context.Cbrt(x, x)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	e.PushVal(x)
}

func DivDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	y := zc.Decimal.As(e.Pop())
	x := zc.Decimal.As(e.Pop())
	defer zc.Decimal.Recycle(y)

	cond, err := s.Context.Quo(x, x, y)
	if cond.DivisionByZero() {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	e.PushVal(x)
}

func MulDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	y := zc.Decimal.As(e.Pop())
	x := zc.Decimal.As(e.Pop())
	defer zc.Decimal.Recycle(y)

	_, err := s.Context.Mul(x, x, y)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	e.PushVal(x)
}

func NegDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	x := zc.Decimal.As(e.Pop())
	_, err := s.Context.Neg(x, x)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	e.PushVal(x)
}

func PowDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	y := zc.Decimal.As(e.Pop())
	x := zc.Decimal.As(e.Pop())
	defer zc.Decimal.Recycle(y)

	_, err := s.Context.Pow(x, x, y)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	e.PushVal(x)
}

func RemDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	y := zc.Decimal.As(e.Pop())
	x := zc.Decimal.As(e.Pop())
	defer zc.Decimal.Recycle(y)

	cond, err := s.Context.Rem(x, x, y)
	if cond.DivisionByZero() {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	e.PushVal(x)
}

func SignDecimal(e *zc.OpEnv) {
	x := zc.Decimal.As(e.Pop())
	e.PushVal(x.Sign())
	zc.Decimal.Recycle(x)
}

func SqrtDecimal(e *zc.OpEnv) {
	var zero apd.Decimal
	s := state.ForDec(e.State)
	x := zc.Decimal.As(e.Pop())

	if x.Cmp(&zero) < 0 {
		e.Err = zc.ErrInvalidArg(e, "%v < 0", x)
		return
	}
	_, err := s.Context.Sqrt(x, x)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	e.PushVal(x)
}

func SubDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	y := zc.Decimal.As(e.Pop())
	x := zc.Decimal.As(e.Pop())
	defer zc.Decimal.Recycle(y)

	_, err := s.Context.Sub(x, x, y)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	e.PushVal(x)
}
