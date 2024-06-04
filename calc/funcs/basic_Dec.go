package funcs

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/calc/state"
	"github.com/blackchip-org/zc/v6/calc/types"
	"github.com/cockroachdb/apd/v3"
)

func AddDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	x := e.Args[0].(*apd.Decimal)
	y := e.Args[1].(*apd.Decimal)
	_, err := s.Context.Add(x, x, y)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	e.Returns = []any{x}
}

func CbrtDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	x := e.Args[0].(*apd.Decimal)
	if x.Cmp(types.DecZero) < 0 {
		e.Err = zc.ErrInvalidArg(e, "%v < 0", x)
		return
	}
	_, err := s.Context.Cbrt(x, x)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	e.Returns = []any{x}
}

func DivDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	x := e.Args[0].(*apd.Decimal)
	y := e.Args[1].(*apd.Decimal)
	cond, err := s.Context.Quo(x, x, y)
	if cond.DivisionByZero() {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	e.Returns = []any{x}
}

func MulDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	x := e.Args[0].(*apd.Decimal)
	y := e.Args[1].(*apd.Decimal)
	_, err := s.Context.Mul(x, x, y)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	e.Returns = []any{x}
}

func NegDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	x := e.Args[0].(*apd.Decimal)
	_, err := s.Context.Neg(x, x)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	e.Returns = []any{x}
}

func PowDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	x := e.Args[0].(*apd.Decimal)
	y := e.Args[1].(*apd.Decimal)
	_, err := s.Context.Pow(x, x, y)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	e.Returns = []any{x}
}

func RemDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	x := e.Args[0].(*apd.Decimal)
	y := e.Args[1].(*apd.Decimal)
	_, err := s.Context.Rem(x, x, y)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	e.Returns = []any{x}
}

func SqrtDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	x := e.Args[0].(*apd.Decimal)
	if x.Cmp(types.DecZero) < 0 {
		e.Err = zc.ErrInvalidArg(e, "%v < 0", x)
		return
	}
	_, err := s.Context.Sqrt(x, x)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	e.Returns = []any{x}
}

func SubDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	x := e.Args[0].(*apd.Decimal)
	y := e.Args[1].(*apd.Decimal)
	_, err := s.Context.Sub(x, x, y)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	e.Returns = []any{x}
}
