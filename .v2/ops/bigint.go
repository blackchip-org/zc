package ops

import (
	"math/big"

	t "github.com/blackchip-org/zc/types"
)

var zeroBigInt big.Int

func op1BigInt(fn func(*big.Int, *big.Int) error) Func {
	return func(args []t.Value) ([]t.Value, error) {
		x := t.BigInt.Native(args[0])
		z := new(big.Int)
		err := fn(z, x)
		return []t.Value{t.BigInt.Value(z)}, err
	}
}

func op2BigInt(fn func(*big.Int, *big.Int, *big.Int) error) Func {
	return func(args []t.Value) ([]t.Value, error) {
		x := t.BigInt.Native(args[0])
		y := t.BigInt.Native(args[1])
		z := new(big.Int)
		err := fn(z, x, y)
		return []t.Value{t.BigInt.Value(z)}, err
	}
}

func opCmpBigInt(fn func(*big.Int, *big.Int) bool) Func {
	return func(args []t.Value) ([]t.Value, error) {
		x := t.BigInt.Native(args[0])
		y := t.BigInt.Native(args[1])
		z := fn(x, y)
		return []t.Value{t.Bool.Value(z)}, nil
	}
}

// FIXME: this is awkward
func divBigInt(args []t.Value) ([]t.Value, error) {
	x := t.BigInt.Native(args[0])
	y := t.BigInt.Native(args[1])
	if y.Cmp(&zeroBigInt) == 0 {
		return []t.Value{}, ErrDivisionByZero
	}
	z := new(big.Int)
	m := new(big.Int)
	z.DivMod(x, y, m)
	if m.Cmp(&zeroBigInt) != 0 {
		xf, err := t.To(args[0], t.Float)
		if err != nil {
			return []t.Value{}, err
		}
		yf, err := t.To(args[1], t.Float)
		if err != nil {
			return []t.Value{}, err
		}
		zf := t.Float.Native(xf) / t.Float.Native(yf)
		return []t.Value{t.Float.Value(zf)}, nil
	}
	return []t.Value{t.BigInt.Value(z)}, nil
}

func opDivBigInt(fn func(*big.Int, *big.Int, *big.Int) error) Func {
	return func(args []t.Value) ([]t.Value, error) {
		x := t.BigInt.Native(args[0])
		y := t.BigInt.Native(args[1])
		if y.Cmp(&zeroBigInt) == 0 {
			return []t.Value{}, ErrDivisionByZero
		}
		z := new(big.Int)
		err := fn(z, x, y)
		return []t.Value{t.BigInt.Value(z)}, err
	}
}

var (
	absBigInt   = op1BigInt(func(z *big.Int, x *big.Int) error { z.Abs(x); return nil })
	addBigInt   = op2BigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.Add(x, y); return nil })
	andBigInt   = op2BigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.And(x, y); return nil })
	ceilBigInt  = op1BigInt(func(z *big.Int, x *big.Int) error { z.Set(x); return nil })
	eqBigInt    = opCmpBigInt(func(x *big.Int, y *big.Int) bool { return x.Cmp(y) == 0 })
	floorBigInt = op1BigInt(func(z *big.Int, x *big.Int) error { z.Set(x); return nil })
	gtBigInt    = opCmpBigInt(func(x *big.Int, y *big.Int) bool { return x.Cmp(y) > 0 })
	gteBigInt   = opCmpBigInt(func(x *big.Int, y *big.Int) bool { return x.Cmp(y) >= 0 })
	ltBigInt    = opCmpBigInt(func(x *big.Int, y *big.Int) bool { return x.Cmp(y) < 0 })
	lteBigInt   = opCmpBigInt(func(x *big.Int, y *big.Int) bool { return x.Cmp(y) <= 0 })
	modBigInt   = opDivBigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.Mod(x, y); return nil })
	mulBigInt   = op2BigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.Mul(x, y); return nil })
	neqBigInt   = opCmpBigInt(func(x *big.Int, y *big.Int) bool { return x.Cmp(y) != 0 })
	negBigInt   = op1BigInt(func(z *big.Int, x *big.Int) error { z.Neg(x); return nil })
	notBigInt   = op1BigInt(func(z *big.Int, x *big.Int) error { z.Not(x); return nil })
	orBigInt    = op2BigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.Or(x, y); return nil })
	powBigInt   = op2BigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.Exp(x, y, nil); return nil })
	remBigInt   = opDivBigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.Rem(x, y); return nil })
	signBigInt  = op1BigInt(func(z *big.Int, x *big.Int) error { z.SetInt64(int64(x.Sign())); return nil })
	subBigInt   = op2BigInt(func(z *big.Int, x *big.Int, y *big.Int) error { z.Sub(x, y); return nil })
)
