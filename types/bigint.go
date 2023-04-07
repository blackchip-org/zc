package types

import (
	"fmt"
	"math/big"
)

var (
	zeroBigInt big.Int
)

type bigIntVal struct {
	val *big.Int
}

func (v bigIntVal) Type() Type     { return BigInt }
func (v bigIntVal) Format() string { return v.val.String() }
func (v bigIntVal) String() string { return fmt.Sprintf("%v(%v)", v.Type().String(), v.Format()) }

type BigIntType struct{}

func (t BigIntType) String() string { return "BigInt" }

func (t BigIntType) Parse(s string) (*big.Int, bool) {
	r := new(big.Int)
	return r.SetString(s, 0)
}

func (t BigIntType) ParseValue(s string) (Value, bool) {
	v, ok := t.Parse(s)
	if !ok {
		return Nil, false
	}
	return t.Value(v), true
}

func (t BigIntType) Value(i *big.Int) Value {
	return bigIntVal{val: i}
}

func (t BigIntType) Unwrap(v Value) *big.Int {
	return v.(bigIntVal).val
}

type op2BigIntFn func(*big.Int, *big.Int, *big.Int) error

func op2BigInt(fn op2BigIntFn) OpFn {
	return func(args []Value) ([]Value, error) {
		x := BigInt.Unwrap(args[0])
		y := BigInt.Unwrap(args[1])
		z := new(big.Int)
		err := fn(z, x, y)
		return []Value{BigInt.Value(z)}, err
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
	addBigInt = op2BigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.Add(x, y); return nil })
	divBigInt = op2BigInt(divBigIntFn)
	mulBigInt = op2BigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.Mul(x, y); return nil })
	subBigInt = op2BigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.Sub(x, y); return nil })
)
