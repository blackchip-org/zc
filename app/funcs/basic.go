package funcs

import "github.com/blackchip-org/zc/v6"

func AddBigInt(c zc.Calc) {
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	x.Add(x, y)
	zc.BigInt.Push(c, x)
	zc.BigInt.Recycle(y)
}
