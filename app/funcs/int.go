package funcs

import "github.com/blackchip-org/zc/v6"

func IntToData(c zc.Calc) {
	x := zc.BigInt.Pop(c)
	d := zc.Data.New()
	d.Write(x.Bytes())
	zc.Data.Push(c, d)
}
