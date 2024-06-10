package funcs

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/calc/types"
)

var ZeroBigFloat big.Float

func AddBigFloat(e *zc.OpEnv) {
	x := e.Args[0].(*big.Float)
	y := e.Args[1].(*big.Float)
	x.Add(x, y)
	e.Returns = []any{x}
}

func DivBigFloat(e *zc.OpEnv) {
	x := e.Args[0].(*big.Float)
	y := e.Args[1].(*big.Float)
	if y.Cmp(&ZeroBigFloat) == 0 {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	x.Quo(x, y)
	e.Returns = []any{x}
}

func MulBigFloat(e *zc.OpEnv) {
	x := e.Args[0].(*big.Float)
	y := e.Args[1].(*big.Float)
	x.Mul(x, y)
	e.Returns = []any{x}
}

func NegBigFloat(e *zc.OpEnv) {
	x := e.Args[0].(*big.Float)
	x.Neg(x)
	e.Returns = []any{x}
}

func SignBigFloat(e *zc.OpEnv) {
	x := e.Args[0].(*big.Float)
	s := x.Sign()
	e.Returns = []any{s}
}

func SqrtBigFloat(e *zc.OpEnv) {
	x := e.Args[0].(*big.Float)
	if x.Cmp(types.FloatZero) < 0 {
		e.Err = zc.ErrInvalidArg(e, "%v < 0", x)
	}
	x.Sqrt(x)
	e.Returns = []any{x}
}

func SubBigFloat(e *zc.OpEnv) {
	x := e.Args[0].(*big.Float)
	y := e.Args[1].(*big.Float)
	x.Sub(x, y)
	e.Returns = []any{x}
}
