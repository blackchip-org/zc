package types

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type decimalVal struct {
	val decimal.Decimal
}

func (v decimalVal) Type() Type     { return Decimal }
func (v decimalVal) Format() string { return v.val.String() }
func (v decimalVal) String() string { return fmt.Sprintf("%v(%v)", v.Type().String(), v.Format()) }

type DecimalType struct{}

func (t DecimalType) String() string { return "Decimal" }

func (t DecimalType) Parse(s string) (decimal.Decimal, bool) {
	d, err := decimal.NewFromString(s)
	return d, err == nil
}

func (t DecimalType) ParseValue(s string) (Value, bool) {
	v, ok := t.Parse(s)
	if !ok {
		return Nil, false
	}
	return t.Value(v), true
}

func (t DecimalType) Value(d decimal.Decimal) Value {
	return decimalVal{val: d}
}

func (t DecimalType) Unwrap(v Value) decimal.Decimal {
	return v.(decimalVal).val
}

type op2DecimalFn func(decimal.Decimal, decimal.Decimal) (decimal.Decimal, error)

func op2Decimal(fn op2DecimalFn) OpFn {
	return func(args []Value) ([]Value, error) {
		x := Decimal.Unwrap(args[0])
		y := Decimal.Unwrap(args[1])
		z, err := fn(x, y)
		if err != nil {
			return []Value{}, err
		}
		return []Value{Decimal.Value(z)}, nil
	}
}

func divDecimalFn(x decimal.Decimal, y decimal.Decimal) (decimal.Decimal, error) {
	if y.IsZero() {
		return decimal.Zero, ErrDivisionByZero
		/*
			switch {
			case x.IsPositive():
				return decimal.Zero, ErrInfPlus
			case x.IsNegative():
				return decimal.Zero, ErrInfMinus
			}
			return decimal.Zero, ErrDivisionByZero
		*/
	}
	return x.Div(y), nil
}

var (
	addDecimal = op2Decimal(func(x decimal.Decimal, y decimal.Decimal) (decimal.Decimal, error) { return x.Add(y), nil })
	divDecimal = op2Decimal(divDecimalFn)
	mulDecimal = op2Decimal(func(x decimal.Decimal, y decimal.Decimal) (decimal.Decimal, error) { return x.Mul(y), nil })
	subDecimal = op2Decimal(func(x decimal.Decimal, y decimal.Decimal) (decimal.Decimal, error) { return x.Sub(y), nil })
)
