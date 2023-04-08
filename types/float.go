package types

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type gFloat struct {
	val float64
}

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'g', 16, 64)
}

func (g gFloat) Type() Type     { return Float }
func (g gFloat) Format() string { return formatFloat(g.val) }
func (g gFloat) String() string { return fmt.Sprintf("%v(%v)", g.Type().String(), g.Format()) }
func (g gFloat) Value() any     { return g.val }

type FloatType struct{}

func (t FloatType) String() string { return "Float" }

func (t FloatType) Parse(s string) (float64, bool) {
	s = strings.TrimSuffix(s, "f")
	f, err := strconv.ParseFloat(s, 64)
	return f, err == nil
}

func (t FloatType) ParseGeneric(s string) (Generic, bool) {
	v, ok := t.Parse(s)
	if !ok {
		return Nil, false
	}
	return t.Generic(v), true
}

func (t FloatType) Generic(f float64) Generic {
	return gFloat{val: f}
}

func (t FloatType) Value(v Generic) float64 {
	return v.Value().(float64)
}

func checkFloat(f float64) error {
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return errors.New(formatFloat(f))
	}
	return nil
}

func op1Float(fn func(float64) (float64, error)) OpFn {
	return func(args []Generic) ([]Generic, error) {
		x := Float.Value(args[0])
		z, err := fn(x)
		if err != nil {
			return []Generic{}, err
		}
		if err := checkFloat(z); err != nil {
			return []Generic{}, err
		}
		return []Generic{Float.Generic(z)}, nil
	}
}

func op2Float(fn func(float64, float64) (float64, error)) OpFn {
	return func(args []Generic) ([]Generic, error) {
		x := Float.Value(args[0])
		y := Float.Value(args[1])
		z, err := fn(x, y)
		if err != nil {
			return []Generic{}, err
		}
		if err := checkFloat(z); err != nil {
			return []Generic{}, err
		}
		return []Generic{Float.Generic(z)}, nil
	}
}

var (
	absFloat = op1Float(func(x float64) (float64, error) { return math.Abs(x), nil })
	addFloat = op2Float(func(x float64, y float64) (float64, error) { return x + y, nil })
	divFloat = op2Float(func(x float64, y float64) (float64, error) { return x / y, nil })
	mulFloat = op2Float(func(x float64, y float64) (float64, error) { return x * y, nil })
	subFloat = op2Float(func(x float64, y float64) (float64, error) { return x - y, nil })
)
