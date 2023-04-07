package types

import (
	"fmt"
	"strconv"
)

type floatVal struct {
	val float64
}

func (v floatVal) Type() Type     { return Float }
func (v floatVal) Format() string { return strconv.FormatFloat(v.val, 'g', 16, 64) }
func (v floatVal) String() string { return fmt.Sprintf("%v(%v)", v.Type().String(), v.Format()) }

type FloatType struct{}

func (t FloatType) String() string { return "Float" }

func (t FloatType) Parse(s string) (float64, bool) {
	f, err := strconv.ParseFloat(s, 64)
	return f, err == nil
}

func (t FloatType) ParseValue(s string) (Value, bool) {
	v, ok := t.Parse(s)
	if !ok {
		return Nil, false
	}
	return t.Value(v), true
}

func (t FloatType) Value(f float64) Value {
	return floatVal{val: f}
}

func (t FloatType) Unwrap(v Value) float64 {
	return v.(floatVal).val
}

/*
func checkFloat(f float64) error {
	if math.IsNaN(f) {
		return ErrNaN
	}
	if math.IsInf(f, 1) {
		return ErrInfPlus
	}
	if math.IsInf(f, -1) {
		return ErrInfMinus
	}
	return nil
}

type op2FloatFn func(float64, float64) (float64, error)

func op2Float(fn op2FloatFn) OpFn {
	return func(args []Value) ([]Value, error) {
		x := toFloat(args[0])
		y := toFloat(args[1])
		z, err := fn(x, y)
		if err != nil {
			return []Value{}, err
		}
		if err := checkFloat(z); err != nil {
			return []Value{}, err
		}
		return []Value{fromFloat(z)}, nil
	}
}

var (
	addFloat = op2Float(func(x float64, y float64) (float64, error) { return x + y, nil })
	divFloat = op2Float(func(x float64, y float64) (float64, error) { return x / y, nil })
	mulFloat = op2Float(func(x float64, y float64) (float64, error) { return x * y, nil })
	subFloat = op2Float(func(x float64, y float64) (float64, error) { return x - y, nil })
)
*/
