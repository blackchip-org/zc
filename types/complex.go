package types

import (
	"fmt"
	"strconv"
)

type gComplex struct {
	val complex128
}

func (g gComplex) Type() Type     { return Complex }
func (g gComplex) Format() string { return Complex.Format(g.val) }
func (g gComplex) String() string { return fmt.Sprintf("%v(%v)", g.Type().String(), g.Format()) }
func (g gComplex) Value() any     { return g.val }

type ComplexType struct{}

func (t ComplexType) String() string { return "Complex" }

func (t ComplexType) Parse(s string) (complex128, error) {
	c, err := strconv.ParseComplex(s, 128)
	if err != nil {
		return 0, parseErr(t, s)
	}
	return c, nil
}

func (t ComplexType) ParseGeneric(s string) (Generic, error) {
	v, err := t.Parse(s)
	if err != nil {
		return Nil, err
	}
	return t.Generic(v), nil
}

func (t ComplexType) Format(c complex128) string {
	s := strconv.FormatComplex(c, 'g', 16, 128)
	// For some reason, the complex number is surrounded by parens.
	// Remove them.
	return s[1 : len(s)-1]
}

func (t ComplexType) Generic(c complex128) Generic {
	return gComplex{val: c}
}

func (t ComplexType) Value(v Generic) complex128 {
	return v.Value().(complex128)
}
