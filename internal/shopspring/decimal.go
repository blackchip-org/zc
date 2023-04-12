package shopspring

import (
	"github.com/blackchip-org/zc/types"
	ss "github.com/shopspring/decimal"
)

type ssDecimal struct {
	val ss.Decimal
}

func unwrap(d types.Decimal) ss.Decimal {
	return d.(ssDecimal).val
}

func wrap(d ss.Decimal) ssDecimal {
	return ssDecimal{val: d}
}

func (d ssDecimal) Add(d2 types.Decimal) types.Decimal {
	return wrap(d.val.Add(unwrap(d2)))
}

func (d ssDecimal) Type() types.Type {
	return types.DecimalType{}
}

func (d ssDecimal) Value() types.Value {
	return d
}

func (d ssDecimal) Native() any {
	return d
}

func (d ssDecimal) String() string {
	return d.val.String()
}

func ParseDecimal(s string) (types.Decimal, error) {
	d, err := ss.NewFromString(s)
	return wrap(d), err
}

func NewDecimal(f float64) types.Decimal {
	return wrap(ss.NewFromFloat(f))
}

func UseDecimal() {
	types.UseDecimal(types.DecimalImpl{
		Parse: ParseDecimal,
		New:   NewDecimal,
	})
}
