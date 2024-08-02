package funcs

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/state"
	"github.com/cockroachdb/apd/v3"
)

func AbsDecimal(e *zc.OpEnv) {
	x := zc.Decimal.Pop(e)
	x.Abs(x)
	zc.Decimal.Push(e, x)
}

func CbrtDecimal(e *zc.OpEnv) {
	var zero apd.Decimal
	s := state.ForDec(e.State)
	x := zc.Decimal.Pop(e)

	if x.Cmp(&zero) < 0 {
		e.Err = zc.ErrInvalidArg(e, "%v < 0", x)
		return
	}
	_, err := s.Context.Cbrt(x, x)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	zc.Decimal.Push(e, x)
}

func CeilDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	x := zc.Decimal.Pop(e)
	_, err := s.Context.Ceil(x, x)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	zc.Decimal.Push(e, x)
}

func ExpDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	x := zc.Decimal.Pop(e)
	_, err := s.Context.Exp(x, x)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	zc.Decimal.Push(e, x)
}

func FloorDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	x := zc.Decimal.Pop(e)
	_, err := s.Context.Floor(x, x)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	zc.Decimal.Push(e, x)
}

func LnDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	x := zc.Decimal.Pop(e)
	_, err := s.Context.Ln(x, x)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	zc.Decimal.Push(e, x)
}

func Log10Decimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	x := zc.Decimal.Pop(e)
	_, err := s.Context.Log10(x, x)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	zc.Decimal.Push(e, x)
}

func SqrtDecimal(e *zc.OpEnv) {
	var zero apd.Decimal
	s := state.ForDec(e.State)
	x := zc.Decimal.Pop(e)

	if x.Cmp(&zero) < 0 {
		e.Err = zc.ErrInvalidArg(e, "%v < 0", x)
		return
	}
	_, err := s.Context.Sqrt(x, x)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	zc.Decimal.Push(e, x)
}
