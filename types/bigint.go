package types

import (
	"fmt"
	"math/big"
)

var (
	zeroBigInt big.Int
)

type gBigInt struct {
	val *big.Int
}

func (g gBigInt) Type() Type     { return BigInt }
func (g gBigInt) Format() string { return g.val.String() }
func (g gBigInt) String() string { return fmt.Sprintf("%v(%v)", g.Type().String(), g.Format()) }
func (g gBigInt) Value() any     { return g.val }

type BigIntType struct{}

func (t BigIntType) String() string { return "BigInt" }

func (t BigIntType) Parse(s string) (*big.Int, bool) {
	r := new(big.Int)
	return r.SetString(s, 0)
}

func (t BigIntType) ParseGeneric(s string) (Generic, bool) {
	v, ok := t.Parse(s)
	if !ok {
		return Nil, false
	}
	return t.Generic(v), true
}

func (t BigIntType) Generic(i *big.Int) Generic {
	return gBigInt{val: i}
}

func (t BigIntType) Value(v Generic) *big.Int {
	return v.Value().(*big.Int)
}

func op1BigInt(fn func(*big.Int, *big.Int) error) OpFn {
	return func(args []Generic) ([]Generic, error) {
		x := BigInt.Value(args[0])
		z := new(big.Int)
		err := fn(z, x)
		return []Generic{BigInt.Generic(z)}, err
	}
}

func op2BigInt(fn func(*big.Int, *big.Int, *big.Int) error) OpFn {
	return func(args []Generic) ([]Generic, error) {
		x := BigInt.Value(args[0])
		y := BigInt.Value(args[1])
		z := new(big.Int)
		err := fn(z, x, y)
		return []Generic{BigInt.Generic(z)}, err
	}
}

func divBigIntFn(z *big.Int, x *big.Int, y *big.Int) error {
	if y.Cmp(&zeroBigInt) == 0 {
		return ErrDivisionByZero
	}
	z.Div(x, y)
	return nil
}

var (
	absBigInt = op1BigInt(func(z *big.Int, x *big.Int) error { z.Abs(x); return nil })
	addBigInt = op2BigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.Add(x, y); return nil })
	divBigInt = op2BigInt(divBigIntFn)
	mulBigInt = op2BigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.Mul(x, y); return nil })
	subBigInt = op2BigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.Sub(x, y); return nil })
)
