package types

import (
	"errors"
	"fmt"
	"math/cmplx"
	"strconv"
)

type complexVal struct {
	val complex128
}

func formatComplex(c complex128) string {
	s := strconv.FormatComplex(c, 'g', 16, 128)
	// For some reason, the complex number is surrounded by parens.
	// Remove them.
	return s[1 : len(s)-1]
}

func (v complexVal) Type() Type     { return Complex }
func (v complexVal) Format() string { return formatComplex(v.val) }
func (v complexVal) String() string { return fmt.Sprintf("%v(%v)", v.Type().String(), v.Format()) }

type ComplexType struct{}

func (t ComplexType) String() string { return "Complex" }

func (t ComplexType) Parse(s string) (complex128, bool) {
	c, err := strconv.ParseComplex(s, 128)
	return c, err == nil
}

func (t ComplexType) ParseValue(s string) (Value, bool) {
	v, ok := t.Parse(s)
	if !ok {
		return Nil, false
	}
	return t.Value(v), true
}

func (t ComplexType) Value(c complex128) Value {
	return complexVal{val: c}
}

func (t ComplexType) Unwrap(v Value) complex128 {
	return v.(complexVal).val
}

type op2ComplexFn func(complex128, complex128) (complex128, error)

func op2Complex(fn op2ComplexFn) OpFn {
	return func(args []Value) ([]Value, error) {
		x := Complex.Unwrap(args[0])
		y := Complex.Unwrap(args[1])
		z, err := fn(x, y)
		if err != nil {
			return []Value{}, err
		}
		if cmplx.IsInf(z) || cmplx.IsNaN(z) {
			return []Value{}, errors.New(formatComplex(z))
		}
		return []Value{Complex.Value(z)}, nil
	}
}

var (
	addComplex = op2Complex(func(x complex128, y complex128) (complex128, error) { return x + y, nil })
	divComplex = op2Complex(func(x complex128, y complex128) (complex128, error) { return x / y, nil })
	mulComplex = op2Complex(func(x complex128, y complex128) (complex128, error) { return x * y, nil })
	subComplex = op2Complex(func(x complex128, y complex128) (complex128, error) { return x - y, nil })
)
