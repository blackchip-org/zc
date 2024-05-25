package funcs

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
)

var ZeroBigInt big.Int

func AddBigInt(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	y := e.Args[1].(*big.Int)
	x.Add(x, y)
	e.Returns = []any{x}
}

func DivBigInt(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	y := e.Args[1].(*big.Int)
	if y.Cmp(&ZeroBigInt) == 0 {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	x.Div(x, y)
	e.Returns = []any{x}
}

func ModBigInt(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	y := e.Args[1].(*big.Int)
	x.Mod(x, y)
	e.Returns = []any{x}
}

func MulBigInt(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	y := e.Args[1].(*big.Int)
	x.Mul(x, y)
	e.Returns = []any{x}
}

func NegBigInt(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	x.Neg(x)
	e.Returns = []any{x}
}

func PowBigInt(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	y := e.Args[1].(*big.Int)
	x.Exp(x, y, nil)
	e.Returns = []any{x}
}

func RemBigInt(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	y := e.Args[1].(*big.Int)
	x.Rem(x, y)
	e.Returns = []any{x}
}

func SignBigInt(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	r := x.Sign()
	e.Returns = []any{r}
}

func SqrtBigInt(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	if x.Cmp(&ZeroBigInt) < 0 {
		e.Err = zc.ErrInvalidArg(e, "%v < 0", x)
		return
	}
	x.Sqrt(x)
	e.Returns = []any{x}
}

func SubBigInt(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	y := e.Args[1].(*big.Int)
	x.Sub(x, y)
	e.Returns = []any{x}
}
