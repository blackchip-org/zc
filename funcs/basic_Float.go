package funcs

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
)

func AddBigFloat(e *zc.OpEnv) {
	x := e.Args[0].(*big.Float)
	y := e.Args[1].(*big.Float)
	x.Add(x, y)
	e.Returns = []any{x}
}
