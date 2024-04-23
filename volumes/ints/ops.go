package ints

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/kinds"
)

var Add = zc.Op{
	Name:    "add.intarch",
	Aliases: []string{"addia"},
	Params:  []string{kinds.IntArch, kinds.IntArch},
	Returns: []string{kinds.IntArch},
	Func:    add,
}

func add(e *zc.OpEnv) {
	x := e.Args[0].(int)
	y := e.Args[1].(int)
	e.Returns = []any{x + y}
}
