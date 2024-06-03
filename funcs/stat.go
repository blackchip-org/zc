package funcs

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
)

func Factorial(e *zc.OpEnv) {
	n := e.Args[0].(uint)
	if n == 0 {
		e.Returns = []any{big.NewInt(1)}
		return
	}

	c := e.Calc()
	c.Eval("1 1 1")
	for i := uint(0); i < n; i++ {
		c.Eval("mul up 1 add dup down")
	}
	c.Eval("drop")
	var r big.Int
	e.Err = c.Pop(&r)
	e.Returns = []any{&r}
}
