package funcs

import "github.com/blackchip-org/zc/v6"

func AndBool(c zc.Calc) {
	y := zc.Bool.Pop(c)
	x := zc.Bool.Pop(c)
	zc.Bool.Push(c, x && y)
}

func NotBool(c zc.Calc) {
	x := zc.Bool.Pop(c)
	zc.Bool.Push(c, !x)
}

func OrBool(c zc.Calc) {
	y := zc.Bool.Pop(c)
	x := zc.Bool.Pop(c)
	zc.Bool.Push(c, x || y)
}
