package funcs

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/state"
	"github.com/cockroachdb/apd/v3"
)

func AddDec(e *zc.OpEnv) {
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

func DivDec(e *zc.OpEnv) {
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

func MulDec(e *zc.OpEnv) {
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

func NegDec(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	x := e.Args[0].(*apd.Decimal)
	_, err := s.Context.Neg(x, x)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	e.Returns = []any{x}
}

func SubDec(e *zc.OpEnv) {
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
