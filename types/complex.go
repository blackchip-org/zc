package types

import (
	"strconv"
)

type vComplex struct {
	val complex128
}

func (v vComplex) Type() Type     { return Complex }
func (v vComplex) Format() string { return Complex.Format(v.val) }
func (v vComplex) String() string { return v.String() }
func (v vComplex) Native() any    { return v.val }

type ComplexType struct{}

func (t ComplexType) String() string { return "Complex" }

func (t ComplexType) Parse(s string) (complex128, error) {
	c, err := strconv.ParseComplex(s, 128)
	if err != nil {
		return 0, parseErr(t, s)
	}
	return c, nil
}

func (t ComplexType) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return Nil, err
	}
	return t.Value(v), nil
}

func (t ComplexType) Format(c complex128) string {
	s := strconv.FormatComplex(c, 'g', 16, 128)
	// For some reason, the complex number is surrounded by parens.
	// Remove them.
	return s[1 : len(s)-1]
}

func (t ComplexType) Value(c complex128) Value {
	return vComplex{val: c}
}

func (t ComplexType) Native(v Value) complex128 {
	return v.Native().(complex128)
}
