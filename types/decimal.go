package types

import (
	"strings"

	"github.com/shopspring/decimal"
)

type vDecimal struct {
	val decimal.Decimal
}

func (v vDecimal) Type() Type     { return Decimal }
func (v vDecimal) Format() string { return Decimal.Format(v.val) }
func (v vDecimal) String() string { return stringV(v) }
func (v vDecimal) Native() any    { return v.val }

type DecimalType struct{}

func (t DecimalType) String() string { return "Decimal" }

func (t DecimalType) Parse(s string) (decimal.Decimal, error) {
	s = cleanNumber(s)
	d, err := decimal.NewFromString(s)
	if err != nil {
		return decimal.Zero, parseErr(t, s)
	}
	return d, nil
}

func (t DecimalType) ParseValue(s string) (Value, error) {
	sl := strings.ToLower(s)
	if !strings.HasSuffix(s, "d") {
		if strings.Contains(sl, "e") {
			return Nil, parseErr(t, s)
		}
	}
	s = strings.TrimSuffix(s, "d")
	v, err := t.Parse(s)
	if err != nil {
		return Nil, parseErr(t, s)
	}
	return t.Value(v), nil
}

func (t DecimalType) Format(d decimal.Decimal) string {
	return d.String()
}

func (t DecimalType) Value(d decimal.Decimal) Value {
	return vDecimal{val: d}
}

func (t DecimalType) Native(v Value) decimal.Decimal {
	return v.Native().(decimal.Decimal)
}
