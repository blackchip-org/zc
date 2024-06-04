package funcs

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
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
