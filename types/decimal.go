package types

import "github.com/blackchip-org/zc/number"

type vDecimal struct {
	val number.Decimal
}

func (v vDecimal) Type() Type     { return Decimal }
func (v vDecimal) String() string { return Decimal.Format(v.val) }
func (v vDecimal) Native() any    { return v.val }

type DecimalType struct{}

func (t DecimalType) String() string { return "Decimal" }

func (t DecimalType) Parse(s string) (number.Decimal, error) {
	return number.ParseDecimal(cleanNumber(s))
}

func (t DecimalType) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return nil, err
	}
	return t.Value(v), nil
}

func (t DecimalType) Format(d number.Decimal) string {
	return d.String()
}

func (t DecimalType) Value(d number.Decimal) Value {
	return vDecimal{val: d}
}

func (t DecimalType) Native(v Value) number.Decimal {
	return v.Native().(number.Decimal)
}
