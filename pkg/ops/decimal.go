package ops

import (
	"github.com/blackchip-org/zc/pkg/zc"
	"github.com/shopspring/decimal"
)

func Dec(c zc.Calc) {
	a0 := zc.PopDecimal(c)
	zc.PushDecimal(c, a0)
}

func DecFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := decimal.NewFromFloat(a0)
	zc.PushDecimal(c, r0)
}
