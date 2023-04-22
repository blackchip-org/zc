package ops

import (
	"github.com/blackchip-org/zc/pkg/zc"
	"github.com/shopspring/decimal"
)

type formatState struct {
	roundingMode string
}

func getFormatState(c zc.Calc) *formatState {
	s, ok := c.State("format")
	if !ok {
		s = &formatState{roundingMode: "half-up"}
		c.NewState("format", s)
	}
	return s.(*formatState)
}

var roundingModes = map[string]struct{}{
	"ceil":      {},
	"down":      {},
	"floor":     {},
	"half-up":   {},
	"half-even": {},
	"up":        {},
}

func round(d decimal.Decimal, places int32, mode string) decimal.Decimal {
	switch mode {
	case "ceil":
		return d.RoundCeil(places)
	case "down":
		return d.RoundDown(places)
	case "floor":
		return d.RoundFloor(places)
	case "half-up":
		return d.Round(places)
	case "half-even":
		return d.RoundBank(places)
	case "up":
		return d.RoundUp(places)
	}
	panic("invalid rounding mode")
}

func Round(c zc.Calc) {
	s := getFormatState(c)
	places := zc.PopInt32(c)
	a0 := zc.PopDecimal(c)
	r0 := round(a0, places, s.roundingMode)
	zc.PushDecimal(c, r0)
}

func RoundingMode(c zc.Calc) {
	s := getFormatState(c)
	a0 := zc.PopString(c)
	if _, ok := roundingModes[a0]; !ok {
		zc.ErrInvalidArgs(c)
		return
	}
	s.roundingMode = a0
	c.SetInfo("rounding-mode set to %v", zc.Quote(a0))
}

func RoundingModeGet(c zc.Calc) {
	s := getFormatState(c)
	zc.PushString(c, s.roundingMode)
}
