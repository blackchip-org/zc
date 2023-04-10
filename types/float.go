package types

import (
	"strconv"
	"strings"
)

type vFloat struct {
	val float64
}

func (v vFloat) Type() Type     { return Float }
func (v vFloat) Format() string { return Float.Format(v.val) }
func (v vFloat) String() string { return stringV(v) }
func (v vFloat) Native() any    { return v.val }

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
	return vFloat{val: f}
}

func (t FloatType) Native(v Value) float64 {
	return v.Native().(float64)
}
