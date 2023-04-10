package types

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

type gDecimal struct {
	val decimal.Decimal
}

func (g gDecimal) Type() Type     { return Decimal }
func (g gDecimal) Format() string { return Decimal.Format(g.val) }
func (g gDecimal) String() string { return fmt.Sprintf("%v(%v)", g.Type().String(), g.Format()) }
func (g gDecimal) Native() any    { return g.val }

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
	return gDecimal{val: d}
}

func (t DecimalType) Native(v Value) decimal.Decimal {
	return v.Native().(decimal.Decimal)
}
