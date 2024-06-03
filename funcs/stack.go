package funcs

import (
	"github.com/blackchip-org/zc/v6"
)

func Clear(e *zc.OpEnv) {
}

func Drop(e *zc.OpEnv) {
}

func Down(e *zc.OpEnv) {
	n := len(e.Args)
	e.Returns = append(
		[]any{e.Args[n-1]},
		e.Args[:n-1]...,
	)
}

func Dup(e *zc.OpEnv) {
	x := e.Catalog.Dup(e.Args[0])
	e.Returns = []any{x, e.Args[0]}
}

func Up(e *zc.OpEnv) {
	e.Returns = append(
		e.Args[1:],
		e.Args[0],
	)
}
