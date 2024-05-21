package basic_Int

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
)

var zero big.Int

func add(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	y := e.Args[1].(*big.Int)
	x.Add(x, y)
	e.Returns = []any{x}
}

func mod(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	y := e.Args[1].(*big.Int)

	if y.Cmp(&zero) == 0 {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	x.Mod(x, y)
	e.Returns = []any{x}
}

func mul(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	y := e.Args[1].(*big.Int)
	x.Mul(x, y)
	e.Returns = []any{x}
}

func neg(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	x.Neg(x)
	e.Returns = []any{x}
}

func pow(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	y := e.Args[1].(*big.Int)
	x.Exp(x, y, nil)
	e.Returns = []any{x}
}

func rem(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	y := e.Args[1].(*big.Int)
	x.Rem(x, y)
	e.Returns = []any{x}
}

func sign(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	r := x.Sign()
	e.Returns = []any{r}
}

func sqrt(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	if x.Cmp(&zero) < 0 {
		e.Err = zc.ErrInvalidArg(e, "%v < 0", x)
		return
	}
	x.Sqrt(x)
	e.Returns = []any{x}
}

func sub(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	y := e.Args[1].(*big.Int)
	x.Sub(x, y)
	e.Returns = []any{x}
}
