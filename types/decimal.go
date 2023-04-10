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
func (g gDecimal) Value() any     { return g.val }

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

func (t DecimalType) ParseGeneric(s string) (Generic, error) {
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
	return t.Generic(v), nil
}

func (t DecimalType) Format(d decimal.Decimal) string {
	return d.String()
}

func (t DecimalType) Generic(d decimal.Decimal) Generic {
	return gDecimal{val: d}
}

func (t DecimalType) Value(v Generic) decimal.Decimal {
	return v.Value().(decimal.Decimal)
}
