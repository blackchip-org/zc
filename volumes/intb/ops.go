package intb

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/kinds"
)

var Add = zc.Op{
	Name:    "add.int",
	Aliases: []string{"addi"},
	Params:  []string{kinds.Int, kinds.Int},
	Returns: []string{kinds.Int},
	Func:    add,
}

func add(e *zc.OpEnv) {
	x := e.Args[0].(*big.Int)
	y := e.Args[1].(*big.Int)
	x.Add(x, y)
	e.Returns = []any{x}
}
