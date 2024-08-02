package funcs

import (
	"fmt"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/state"
)

func AddDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	y := zc.Decimal.Pop(e)
	x := zc.Decimal.Pop(e)
	defer zc.Decimal.Recycle(y)

	c, err := s.Context.Add(x, x, y)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	fmt.Printf("*** COND: %v\n", c)
	zc.Decimal.Push(e, x)
}

func DivDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	y := zc.Decimal.Pop(e)
	x := zc.Decimal.Pop(e)
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
	zc.Decimal.Push(e, x)
}

func MulDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	y := zc.Decimal.Pop(e)
	x := zc.Decimal.Pop(e)
	defer zc.Decimal.Recycle(y)

	_, err := s.Context.Mul(x, x, y)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	zc.Decimal.Push(e, x)
}

func NegDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	x := zc.Decimal.Pop(e)
	_, err := s.Context.Neg(x, x)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	zc.Decimal.Push(e, x)
}

func PowDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	y := zc.Decimal.Pop(e)
	x := zc.Decimal.Pop(e)
	defer zc.Decimal.Recycle(y)

	_, err := s.Context.Pow(x, x, y)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	zc.Decimal.Push(e, x)
}

func RemDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	y := zc.Decimal.Pop(e)
	x := zc.Decimal.Pop(e)
	defer zc.Decimal.Recycle(y)

	if y.IsZero() {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	_, err := s.Context.Rem(x, x, y)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	zc.Decimal.Push(e, x)
}

func SignDecimal(e *zc.OpEnv) {
	x := zc.Decimal.Pop(e)
	zc.Int.Push(e, x.Sign())
	zc.Decimal.Recycle(x)
}

func SubDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	y := zc.Decimal.Pop(e)
	x := zc.Decimal.Pop(e)
	defer zc.Decimal.Recycle(y)

	_, err := s.Context.Sub(x, x, y)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	zc.Decimal.Push(e, x)
}
