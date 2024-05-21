package intb

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
)

func add(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	y := e.Args[1].(*big.Int)
	x.Add(x, y)
	e.Returns = []any{x}
}

func pow(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	y := e.Args[1].(*big.Int)
	x.Exp(x, y, nil)
	e.Returns = []any{x}
}

func sub(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	y := e.Args[1].(*big.Int)
	x.Sub(x, y)
	e.Returns = []any{x}
}
