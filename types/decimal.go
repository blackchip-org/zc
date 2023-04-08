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
func (g gDecimal) Format() string { return g.val.String() }
func (g gDecimal) String() string { return fmt.Sprintf("%v(%v)", g.Type().String(), g.Format()) }
func (g gDecimal) Value() any     { return g.val }

type DecimalType struct{}

func (t DecimalType) String() string { return "Decimal" }

func (t DecimalType) Parse(s string) (decimal.Decimal, bool) {
	d, err := decimal.NewFromString(s)
	return d, err == nil
}

func (t DecimalType) ParseGeneric(s string) (Generic, bool) {
	sl := strings.ToLower(s)
	if !strings.HasSuffix(s, "d") {
		if strings.Contains(sl, "e") {
			return Nil, false
		}
	}
	s = strings.TrimSuffix(s, "d")
	v, ok := t.Parse(s)
	if !ok {
		return Nil, false
	}
	return t.Generic(v), true
}

func (t DecimalType) Generic(d decimal.Decimal) Generic {
	return gDecimal{val: d}
}

func (t DecimalType) Value(v Generic) decimal.Decimal {
	return v.Value().(decimal.Decimal)
}

func op1Decimal(fn func(decimal.Decimal) (decimal.Decimal, error)) OpFn {
	return func(args []Generic) ([]Generic, error) {
		x := Decimal.Value(args[0])
		z, err := fn(x)
		if err != nil {
			return []Generic{}, err
		}
		return []Generic{Decimal.Generic(z)}, nil
	}
}

func op2Decimal(fn func(decimal.Decimal, decimal.Decimal) (decimal.Decimal, error)) OpFn {
	return func(args []Generic) ([]Generic, error) {
		x := Decimal.Value(args[0])
		y := Decimal.Value(args[1])
		z, err := fn(x, y)
		if err != nil {
			return []Generic{}, err
		}
		return []Generic{Decimal.Generic(z)}, nil
	}
}

func divDecimalFn(x decimal.Decimal, y decimal.Decimal) (decimal.Decimal, error) {
	if y.IsZero() {
		return decimal.Zero, ErrDivisionByZero
	}
	return x.Div(y), nil
}

var (
	absDecimal = op1Decimal(func(x decimal.Decimal) (decimal.Decimal, error) { return x.Abs(), nil })
	addDecimal = op2Decimal(func(x decimal.Decimal, y decimal.Decimal) (decimal.Decimal, error) { return x.Add(y), nil })
	divDecimal = op2Decimal(divDecimalFn)
	mulDecimal = op2Decimal(func(x decimal.Decimal, y decimal.Decimal) (decimal.Decimal, error) { return x.Mul(y), nil })
	subDecimal = op2Decimal(func(x decimal.Decimal, y decimal.Decimal) (decimal.Decimal, error) { return x.Sub(y), nil })
)
