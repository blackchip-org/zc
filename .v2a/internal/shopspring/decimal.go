package shopspring

import (
	"github.com/blackchip-org/zc/number"
	ss "github.com/shopspring/decimal"
)

type ssDecimal struct {
	val ss.Decimal
}

func unwrap(d number.Decimal) ss.Decimal {
	return d.(ssDecimal).val
}

func wrap(d ss.Decimal) ssDecimal {
	return ssDecimal{val: d}
}

func (d ssDecimal) Add(d2 number.Decimal) number.Decimal {
	return wrap(d.val.Add(unwrap(d2)))
}

func (d ssDecimal) String() string {
	return d.val.String()
}

func ParseDecimal(s string) (number.Decimal, error) {
	d, err := ss.NewFromString(s)
	return wrap(d), err
}

func NewDecimal(f float64) number.Decimal {
	return wrap(ss.NewFromFloat(f))
}

func UseDecimal() {
	number.UseDecimal(number.DecimalImpl{
		Parse: ParseDecimal,
		New:   NewDecimal,
	})
}
