package zcalc

import (
	"github.com/blackchip-org/zc/v6"
)

func down(e *zc.OpEnv) {
	top := len(e.Args) - 1
	e.Returns = append(
		[]any{e.Args[top]},
		e.Args[:top]...,
	)
}

func dup(e *zc.OpEnv) {
	x := e.Catalog.Dup(e.Args[0])
	e.Returns = []any{x, e.Args[0]}
}
