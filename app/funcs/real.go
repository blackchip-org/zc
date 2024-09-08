package funcs

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/vars"
)

func Dec(c zc.Calc) {
	x := zc.Decimal.Pop(c)
	zc.Decimal.Push(c, x)
}

func DecRat(c zc.Calc) {
	x := zc.Rat.Pop(c)
	f, exact := x.Float64()
	if zc.IsFloatErr(c, f) {
		return
	}
	d := zc.Decimal.New()
	d.SetFloat64(f)
	if !exact {
		c.Notify(zc.NoticeInexact)
	}
	zc.Decimal.Push(c, d)
}

func RoundingModeSet(c zc.Calc) {
	conf := vars.ForReal(c)
	mode := zc.String.Pop(c)
	if err := conf.SetRoundingMode(mode); err != nil {
		c.Raise(zc.ErrInvalidArg(err.Error()))
		return
	}
	c.Notify("rounding mode set to %v", mode)
}

func RoundingModeGet(c zc.Calc) {
	conf := vars.ForReal(c)
	zc.String.Push(c, conf.GetRoundingMode())
}

func Trunc(c zc.Calc) {
	x := zc.Decimal.Pop(c)
	i, f := zc.Decimal.New(), zc.Decimal.New()
	x.Modf(i, f)
	zc.Decimal.Recycle(f, x)
	zc.Decimal.Push(c, i)
}
