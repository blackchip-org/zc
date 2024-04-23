package zcalc

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/kinds"
)

var Down = zc.Op{
	Name:       "down",
	Params:     []string{kinds.Val},
	VarParams:  true,
	Returns:    []string{kinds.Val},
	VarReturns: true,
	Func:       down,
}

func down(e *zc.OpEnv) {
	top := len(e.Args) - 1
	e.Returns = append(
		[]any{e.Args[top]},
		e.Args[:top]...,
	)
}

var Dup = zc.Op{
	Name:    "dup",
	Params:  []string{kinds.Val},
	Returns: []string{kinds.Val, kinds.Val},
	Func:    dup,
}

func dup(e *zc.OpEnv) {
	x := e.Catalog.Copy(e.Args[0])
	e.Returns = []any{x, e.Args[0]}
}
