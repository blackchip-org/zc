package ops

import (
	"errors"
	"math/cmplx"

	t "github.com/blackchip-org/zc/types"
)

func op1Complex(fn func(complex128) (complex128, error)) Func {
	return func(args []t.Value) ([]t.Value, error) {
		x := t.Complex.Native(args[0])
		z, err := fn(x)
		if err != nil {
			return []t.Value{}, err
		}
		if cmplx.IsInf(z) || cmplx.IsNaN(z) {
			return []t.Value{}, errors.New(t.Complex.Format(z))
		}
		return []t.Value{t.Complex.Value(z)}, nil
	}
}

func op2Complex(fn func(complex128, complex128) (complex128, error)) Func {
	return func(args []t.Value) ([]t.Value, error) {
		x := t.Complex.Native(args[0])
		y := t.Complex.Native(args[1])
		z, err := fn(x, y)
		if err != nil {
			return []t.Value{}, err
		}
		if cmplx.IsInf(z) || cmplx.IsNaN(z) {
			return []t.Value{}, errors.New(t.Complex.Format(z))
		}
		return []t.Value{t.Complex.Value(z)}, nil
	}
}

func opCmpComplex(fn func(complex128, complex128) bool) Func {
	return func(args []t.Value) ([]t.Value, error) {
		x := t.Complex.Native(args[0])
		y := t.Complex.Native(args[1])
		z := fn(x, y)
		return []t.Value{t.Bool.Value(z)}, nil
	}
}

func divComplexFn(x complex128, y complex128) (complex128, error) {
	if real(y) == 0 && imag(y) == 0 {
		return 0, ErrDivisionByZero
	}
	return x / y, nil
}

func absComplex(args []t.Value) ([]t.Value, error) {
	x := t.Complex.Native(args[0])
	z := cmplx.Abs(x)
	return []t.Value{t.Float.Value(z)}, nil
}

var (
	addComplex = op2Complex(func(x complex128, y complex128) (complex128, error) { return x + y, nil })
	divComplex = op2Complex(divComplexFn)
	eqComplex  = opCmpComplex(func(x complex128, y complex128) bool { return x == y })
	mulComplex = op2Complex(func(x complex128, y complex128) (complex128, error) { return x * y, nil })
	neqComplex = opCmpComplex(func(x complex128, y complex128) bool { return x != y })
	powComplex = op2Complex(func(x complex128, y complex128) (complex128, error) { return cmplx.Pow(x, y), nil })
	subComplex = op2Complex(func(x complex128, y complex128) (complex128, error) { return x - y, nil })
)
