package zcalc

import (
	"github.com/blackchip-org/zc/v6"
)

func clear(e *zc.OpEnv) {
}

func down(e *zc.OpEnv) {
	n := len(e.Args)
	e.Returns = append(
		[]any{e.Args[n-1]},
		e.Args[:n-1]...,
	)
}

func dup(e *zc.OpEnv) {
	x := e.Catalog.Dup(e.Args[0])
	e.Returns = []any{x, e.Args[0]}
}
