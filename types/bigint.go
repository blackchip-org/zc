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

func opDivBigInt(fn func(*big.Int, *big.Int, *big.Int) error) OpFn {
	return func(args []Generic) ([]Generic, error) {
		x := BigInt.Value(args[0])
		y := BigInt.Value(args[1])
		if y.Cmp(&zeroBigInt) == 0 {
			return []Generic{}, ErrDivisionByZero
		}
		z := new(big.Int)
		err := fn(z, x, y)
		return []Generic{BigInt.Generic(z)}, err
	}
}

var (
	absBigInt   = op1BigInt(func(z *big.Int, x *big.Int) error { z.Abs(x); return nil })
	addBigInt   = op2BigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.Add(x, y); return nil })
	ceilBigInt  = op1BigInt(func(z *big.Int, x *big.Int) error { z.Set(x); return nil })
	divBigInt   = opDivBigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.Div(x, y); return nil })
	floorBigInt = op1BigInt(func(z *big.Int, x *big.Int) error { z.Set(x); return nil })
	modBigInt   = opDivBigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.Mod(x, y); return nil })
	mulBigInt   = op2BigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.Mul(x, y); return nil })
	negBigInt   = op1BigInt(func(z *big.Int, x *big.Int) error { z.Neg(x); return nil })
	powBigInt   = op2BigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.Exp(x, y, nil); return nil })
	remBigInt   = opDivBigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.Rem(x, y); return nil })
	signBigInt  = op1BigInt(func(z *big.Int, x *big.Int) error { z.SetInt64(int64(x.Sign())); return nil })
	subBigInt   = op2BigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.Sub(x, y); return nil })
)
