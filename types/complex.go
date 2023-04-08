package types

import (
	"errors"
	"fmt"
	"math/cmplx"
	"strconv"
)

type gComplex struct {
	val complex128
}

func formatComplex(c complex128) string {
	s := strconv.FormatComplex(c, 'g', 16, 128)
	// For some reason, the complex number is surrounded by parens.
	// Remove them.
	return s[1 : len(s)-1]
}

func (g gComplex) Type() Type     { return Complex }
func (g gComplex) Format() string { return formatComplex(g.val) }
func (g gComplex) String() string { return fmt.Sprintf("%v(%v)", g.Type().String(), g.Format()) }
func (g gComplex) Value() any     { return g.val }

type ComplexType struct{}

func (t ComplexType) String() string { return "Complex" }

func (t ComplexType) Parse(s string) (complex128, bool) {
	c, err := strconv.ParseComplex(s, 128)
	return c, err == nil
}

func (t ComplexType) ParseGeneric(s string) (Generic, bool) {
	v, ok := t.Parse(s)
	if !ok {
		return Nil, false
	}
	return t.Generic(v), true
}

func (t ComplexType) Generic(c complex128) Generic {
	return gComplex{val: c}
}

func (t ComplexType) Value(v Generic) complex128 {
	return v.Value().(complex128)
}

func op1Complex(fn func(complex128) (complex128, error)) OpFn {
	return func(args []Generic) ([]Generic, error) {
		x := Complex.Value(args[0])
		z, err := fn(x)
		if err != nil {
			return []Generic{}, err
		}
		if cmplx.IsInf(z) || cmplx.IsNaN(z) {
			return []Generic{}, errors.New(formatComplex(z))
		}
		return []Generic{Complex.Generic(z)}, nil
	}
}

func op2Complex(fn func(complex128, complex128) (complex128, error)) OpFn {
	return func(args []Generic) ([]Generic, error) {
		x := Complex.Value(args[0])
		y := Complex.Value(args[1])
		z, err := fn(x, y)
		if err != nil {
			return []Generic{}, err
		}
		if cmplx.IsInf(z) || cmplx.IsNaN(z) {
			return []Generic{}, errors.New(formatComplex(z))
		}
		return []Generic{Complex.Generic(z)}, nil
	}
}

func absComplexFn(args []Generic) ([]Generic, error) {
	x := Complex.Value(args[0])
	z := cmplx.Abs(x)
	return []Generic{Float.Generic(z)}, nil
}

var (
	absComplex = absComplexFn
	addComplex = op2Complex(func(x complex128, y complex128) (complex128, error) { return x + y, nil })
	divComplex = op2Complex(func(x complex128, y complex128) (complex128, error) { return x / y, nil })
	mulComplex = op2Complex(func(x complex128, y complex128) (complex128, error) { return x * y, nil })
	subComplex = op2Complex(func(x complex128, y complex128) (complex128, error) { return x - y, nil })
)
