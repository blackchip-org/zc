package funcs

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/shopspring/decimal"
)

func AddDec(e *zc.OpEnv) {
	x := e.Args[0].(decimal.Decimal)
	y := e.Args[1].(decimal.Decimal)
	x = x.Add(y)
	e.Returns = []any{x}
}

func DivDec(e *zc.OpEnv) {
	x := e.Args[0].(decimal.Decimal)
	y := e.Args[1].(decimal.Decimal)
	if y.IsZero() {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	x = x.Div(y)
	e.Returns = []any{x}
}

func MulDec(e *zc.OpEnv) {
	x := e.Args[0].(decimal.Decimal)
	y := e.Args[1].(decimal.Decimal)
	x = x.Mul(y)
	e.Returns = []any{x}
}

func NegDec(e *zc.OpEnv) {
	x := e.Args[0].(decimal.Decimal)
	x = x.Neg()
	e.Returns = []any{x}
}

func SubDec(e *zc.OpEnv) {
	x := e.Args[0].(decimal.Decimal)
	y := e.Args[1].(decimal.Decimal)
	x = x.Sub(y)
	e.Returns = []any{x}
}
