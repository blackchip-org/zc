package funcs

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/state"
	"github.com/cockroachdb/apd/v3"
)

func AddDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	y := zc.Decimal.As(e.Pop())
	x := zc.Decimal.As(*e.PopTop())
	defer zc.Decimal.Recycle(y)

	_, err := s.Context.Add(x, x, y)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
}

func CbrtDecimal(e *zc.OpEnv) {
	var zero apd.Decimal
	s := state.ForDec(e.State)
	x := zc.Decimal.As(*e.PopTop())

	if x.Cmp(&zero) < 0 {
		e.Err = zc.ErrInvalidArg(e, "%v < 0", x)
		return
	}
	_, err := s.Context.Cbrt(x, x)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
}

func DivDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	y := zc.Decimal.As(e.Pop())
	x := zc.Decimal.As(*e.PopTop())
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
}

func MulDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	y := zc.Decimal.As(e.Pop())
	x := zc.Decimal.As(*e.PopTop())
	defer zc.Decimal.Recycle(y)

	_, err := s.Context.Mul(x, x, y)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
}

func NegDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	x := zc.Decimal.As(*e.PopTop())
	_, err := s.Context.Neg(x, x)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
}

func PowDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	y := zc.Decimal.As(e.Pop())
	x := zc.Decimal.As(*e.PopTop())
	defer zc.Decimal.Recycle(y)

	_, err := s.Context.Pow(x, x, y)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
}

func RemDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	y := zc.Decimal.As(e.Pop())
	x := zc.Decimal.As(*e.PopTop())
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
}

func SignDecimal(e *zc.OpEnv) {
	x := e.PopTop()
	x.Val = zc.Decimal.As(*x).Sign()
}

func SqrtDecimal(e *zc.OpEnv) {
	var zero apd.Decimal
	s := state.ForDec(e.State)
	x := zc.Decimal.As(*e.PopTop())

	if x.Cmp(&zero) < 0 {
		e.Err = zc.ErrInvalidArg(e, "%v < 0", x)
		return
	}
	_, err := s.Context.Sqrt(x, x)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
}

func SubDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	y := zc.Decimal.As(e.Pop())
	x := zc.Decimal.As(*e.PopTop())
	defer zc.Decimal.Recycle(y)

	_, err := s.Context.Sub(x, x, y)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
}
