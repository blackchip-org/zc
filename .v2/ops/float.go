package ops

import (
	"errors"
	"math"
	"math/cmplx"

	t "github.com/blackchip-org/zc/types"
)

func checkFloat(f float64) error {
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return errors.New(t.Float.Format(f))
	}
	return nil
}

func op1Float(fn func(float64) (float64, error)) Func {
	return func(args []t.Value) ([]t.Value, error) {
		x := t.Float.Native(args[0])
		z, err := fn(x)
		if err != nil {
			return []t.Value{}, err
		}
		if err := checkFloat(z); err != nil {
			return []t.Value{}, err
		}
		return []t.Value{t.Float.Value(z)}, nil
	}
}

func op2Float(fn func(float64, float64) (float64, error)) Func {
	return func(args []t.Value) ([]t.Value, error) {
		x := t.Float.Native(args[0])
		y := t.Float.Native(args[1])
		z, err := fn(x, y)
		if err != nil {
			return []t.Value{}, err
		}
		if err := checkFloat(z); err != nil {
			return []t.Value{}, err
		}
		return []t.Value{t.Float.Value(z)}, nil
	}
}

func opCmpFloat(fn func(float64, float64) bool) Func {
	return func(args []t.Value) ([]t.Value, error) {
		x := t.Float.Native(args[0])
		y := t.Float.Native(args[1])
		z := fn(x, y)
		return []t.Value{t.Bool.Value(z)}, nil
	}
}

func opDivFloat(fn func(float64, float64) (float64, error)) Func {
	return func(args []t.Value) ([]t.Value, error) {
		x := t.Float.Native(args[0])
		y := t.Float.Native(args[1])
		if y == 0 {
			return []t.Value{}, ErrDivisionByZero
		}
		z, err := fn(x, y)
		if err != nil {
			return []t.Value{}, err
		}
		if err := checkFloat(z); err != nil {
			return []t.Value{}, err
		}
		return []t.Value{t.Float.Value(z)}, nil
	}
}

func sqrtFloat(args []t.Value) ([]t.Value, error) {
	x := t.Float.Native(args[0])
	if x < 0 {
		z := cmplx.Sqrt(complex(x, 0))
		return []t.Value{t.Complex.Value(z)}, nil
	}
	z := math.Sqrt(x)
	return []t.Value{t.Float.Value(z)}, nil
}

func signFloatFn(x float64) (float64, error) {
	if x == 0 {
		return 0, nil
	}
	if x > 0 {
		return 1, nil
	}
	return -1, nil
}

var (
	absFloat   = op1Float(func(x float64) (float64, error) { return math.Abs(x), nil })
	addFloat   = op2Float(func(x float64, y float64) (float64, error) { return x + y, nil })
	ceilFloat  = op1Float(func(x float64) (float64, error) { return math.Ceil(x), nil })
	divFloat   = opDivFloat(func(x float64, y float64) (float64, error) { return x / y, nil })
	eqFloat    = opCmpFloat(func(x float64, y float64) bool { return x == y })
	floorFloat = op1Float(func(x float64) (float64, error) { return math.Floor(x), nil })
	gtFloat    = opCmpFloat(func(x float64, y float64) bool { return x > y })
	gteFloat   = opCmpFloat(func(x float64, y float64) bool { return x >= y })
	ltFloat    = opCmpFloat(func(x float64, y float64) bool { return x < y })
	lteFloat   = opCmpFloat(func(x float64, y float64) bool { return x <= y })
	modFloat   = opDivFloat(func(x float64, y float64) (float64, error) { return math.Mod(x, y), nil })
	mulFloat   = op2Float(func(x float64, y float64) (float64, error) { return x * y, nil })
	negFloat   = op1Float(func(x float64) (float64, error) { return -x, nil })
	neqFloat   = opCmpFloat(func(x float64, y float64) bool { return x != y })
	powFloat   = op2Float(func(x float64, y float64) (float64, error) { return math.Pow(x, y), nil })
	remFloat   = opDivFloat(func(x float64, y float64) (float64, error) { return math.Remainder(x, y), nil })
	signFloat  = op1Float(signFloatFn)
	subFloat   = op2Float(func(x float64, y float64) (float64, error) { return x - y, nil })
)
