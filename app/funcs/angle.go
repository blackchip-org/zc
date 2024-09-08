package funcs

import (
	"github.com/blackchip-org/dms"
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/types"
)

func DecDMS(c zc.Calc) {
	x := zc.AngleDMS.Pop(c)
	zc.Decimal.Push(c, x.Degrees())
}

func DegMin(c zc.Calc) {
	x := zc.AngleDMS.Pop(c)
	zc.String.Push(c, types.FormatAngleDMS(x, dms.MinUnit, -1))
}

func DegMinRound(c zc.Calc) {
	places := zc.Uint.Pop(c)
	x := zc.AngleDMS.Pop(c)
	zc.String.Push(c, types.FormatAngleDMS(x, dms.MinUnit, int(places)))
}

func DegMinSec(c zc.Calc) {
	x := zc.AngleDMS.Pop(c)
	zc.String.Push(c, types.FormatAngleDMS(x, dms.SecUnit, -1))
}

func DegMinSecRound(c zc.Calc) {
	places := zc.Uint.Pop(c)
	x := zc.AngleDMS.Pop(c)
	zc.String.Push(c, types.FormatAngleDMS(x, dms.SecUnit, int(places)))
}

func DMSIs(c zc.Calc) {
	x := zc.String.Pop(c)
	_, ok, err := zc.AngleDMS.Parse(c, x)
	zc.Bool.Push(c, ok && err == nil)
}

func MinutesDMS(c zc.Calc) {
	x := zc.AngleDMS.Pop(c)
	zc.Decimal.Push(c, x.Minutes())
}

func SecondsDMS(c zc.Calc) {
	x := zc.AngleDMS.Pop(c)
	zc.Decimal.Push(c, x.Seconds())
}
