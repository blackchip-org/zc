package types

import (
	"fmt"
	"strconv"
	"strings"
)

type gFloat struct {
	val float64
}

func (g gFloat) Type() Type     { return Float }
func (g gFloat) Format() string { return Float.Format(g.val) }
func (g gFloat) String() string { return fmt.Sprintf("%v(%v)", g.Type().String(), g.Format()) }
func (g gFloat) Native() any    { return g.val }

type FloatType struct{}

func (t FloatType) String() string { return "Float" }

func (t FloatType) Parse(s string) (float64, error) {
	s = cleanNumber(s)
	s = strings.TrimSuffix(s, "f")
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, parseErr(t, s)
	}
	return f, nil
}

func (t FloatType) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return Nil, err
	}
	return t.Value(v), nil
}

func (t FloatType) Format(f float64) string {
	return strconv.FormatFloat(f, 'g', 16, 64)
}

func (t FloatType) Value(f float64) Value {
	return gFloat{val: f}
}

func (t FloatType) Native(v Value) float64 {
	return v.Native().(float64)
}
