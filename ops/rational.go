package ops

import (
	"math/big"

	t "github.com/blackchip-org/zc/types"
)

var ratZero big.Rat

func op1Rational(fn func(*big.Rat, *big.Rat) error) Func {
	return func(args []t.Generic) ([]t.Generic, error) {
		x := t.Rational.Value(args[0])
		z := new(big.Rat)
		err := fn(z, x)
		return []t.Generic{t.Rational.Generic(z)}, err
	}
}

func op2Rational(fn func(*big.Rat, *big.Rat, *big.Rat) error) Func {
	return func(args []t.Generic) ([]t.Generic, error) {
		x := t.Rational.Value(args[0])
		y := t.Rational.Value(args[1])
		z := new(big.Rat)
		err := fn(z, x, y)
		return []t.Generic{t.Rational.Generic(z)}, err
	}
}

func opCmpRational(fn func(*big.Rat, *big.Rat) bool) Func {
	return func(args []t.Generic) ([]t.Generic, error) {
		x := t.Rational.Value(args[0])
		y := t.Rational.Value(args[1])
		z := fn(x, y)
		return []t.Generic{t.Bool.Generic(z)}, nil
	}
}

func divRationalFn(z *big.Rat, x *big.Rat, y *big.Rat) error {
	if y.Cmp(&ratZero) == 0 {
		return ErrDivisionByZero
	}
	z.Quo(x, y)
	return nil
}

var (
	absRational  = op1Rational(func(z *big.Rat, x *big.Rat) error { z.Abs(x); return nil })
	addRational  = op2Rational(func(z *big.Rat, x *big.Rat, y *big.Rat) error { z.Add(x, y); return nil })
	divRational  = op2Rational(divRationalFn)
	eqRational   = opCmpRational(func(x *big.Rat, y *big.Rat) bool { return x.Cmp(y) == 0 })
	gtRational   = opCmpRational(func(x *big.Rat, y *big.Rat) bool { return x.Cmp(y) > 0 })
	gteRational  = opCmpRational(func(x *big.Rat, y *big.Rat) bool { return x.Cmp(y) >= 0 })
	ltRational   = opCmpRational(func(x *big.Rat, y *big.Rat) bool { return x.Cmp(y) < 0 })
	lteRational  = opCmpRational(func(x *big.Rat, y *big.Rat) bool { return x.Cmp(y) <= 0 })
	mulRational  = op2Rational(func(z *big.Rat, x *big.Rat, y *big.Rat) error { z.Mul(x, y); return nil })
	negRational  = op1Rational(func(z *big.Rat, x *big.Rat) error { z.Neg(x); return nil })
	neqRational  = opCmpRational(func(x *big.Rat, y *big.Rat) bool { return x.Cmp(y) != 0 })
	signRational = op1Rational(func(z *big.Rat, x *big.Rat) error { z.SetInt64(int64(x.Sign())); return nil })
	subRational  = op2Rational(func(z *big.Rat, x *big.Rat, y *big.Rat) error { z.Sub(x, y); return nil })
)
