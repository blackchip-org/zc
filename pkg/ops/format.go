package ops

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/blackchip-org/zc/v5/pkg/types"
	"github.com/blackchip-org/zc/v5/pkg/zc"
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

func round(d types.Decimal, places int32, mode string) types.Decimal {
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

/*
oper	round
func	RoundDecimal d:Decimal n:Int -- Decimal
func	RoundFloat   d:Float   n:Int -- Float
alias	r
title	Round to a given precision

desc
Rounds the number *n* to *d* digits using the current rounding mode.
end

example
2 3 div -- 0.6666666666666666666
2 round -- 0.67
end
*/
func RoundDecimal(c zc.Calc) {
	s := getFormatState(c)
	places := zc.PopInt32(c)
	anno := c.TopAnnotation()
	a0 := zc.PopDecimal(c)
	r0 := round(a0, places, s.roundingMode)
	zc.PushDecimal(c, r0)
	zc.Annotate(c, anno)
}

func RoundFloat(c zc.Calc) {
	places := zc.PopInt(c)
	anno := c.TopAnnotation()
	a0 := zc.PopBigFloat(c)
	r0 := a0.Text('e', places)
	r0 = zc.PostFormatFloat(r0)
	zc.PushString(c, r0)
	zc.Annotate(c, anno)
}

/*
oper	rounding-mode
func	RoundingMode -- Str
title	Set method to use in rounding

desc
Sets the mode to be used when rounding. Valid modes are:

- `half-up`
- `ceil`
- `down`
- `floor`
- `half-even`
- `up`

end

example
1.01 0.05 mul -- 0.0505
2 round -- 0.05
'up' rounding-mode -- *rounding-mode set to 'up'*
c 1.01 0.05 mul -- 0.0505
2 round -- 0.06
end
*/
func RoundingMode(c zc.Calc) {
	s := getFormatState(c)
	a0 := zc.PopString(c)
	if _, ok := roundingModes[a0]; !ok {
		zc.ErrInvalidArgs(c, "rounding mode")
		return
	}
	s.roundingMode = a0
	c.SetInfo("rounding-mode set to %v", zc.Quote(a0))
}

/*
oper rounding-mode=
func RoundingModeGet -- Str
title Method to use in rounding

desc
Gets the current rounding mode
end

example
rounding-mode= -- half-up
end
*/
func RoundingModeGet(c zc.Calc) {
	s := getFormatState(c)
	zc.PushString(c, s.roundingMode)
}

/*
oper	scientific-notation
func	ScientificNotation p0:Float -- Float
func 	ScientificNotationBigInt p0:BigInt -- Float
alias	sn
title	Scientific notation

desc
Formats the value *p0* using scientific notation.
end

example
1234 sn -- 1.234e03
end
*/
func ScientificNotation(c zc.Calc) {
	anno := c.TopAnnotation()
	a0 := zc.PopFloat(c)
	t0 := strconv.FormatFloat(a0, 'e', -1, 64)
	r0 := strings.Replace(t0, "e+", "e", 1)
	zc.PushString(c, r0)
	zc.Annotate(c, anno)
}

func ScientificNotationBigInt(c zc.Calc) {
	anno := c.TopAnnotation()
	a0 := zc.PopBigInt(c)
	f := new(big.Float).SetInt(a0)
	zc.PushBigFloat(c, f)
	zc.Annotate(c, anno)
}
