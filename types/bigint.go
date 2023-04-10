package types

import (
	"fmt"
	"math/big"
)

type gBigInt struct {
	val *big.Int
}

func (g gBigInt) Type() Type     { return BigInt }
func (g gBigInt) Format() string { return BigInt.Format(g.val) }
func (g gBigInt) String() string { return fmt.Sprintf("%v(%v)", g.Type().String(), g.Format()) }
func (g gBigInt) Value() any     { return g.val }

type BigIntType struct{}

func (t BigIntType) String() string { return "BigInt" }

func (t BigIntType) Parse(s string) (*big.Int, error) {
	r := new(big.Int)
	s = cleanNumber(s)
	_, ok := r.SetString(s, 0)
	if !ok {
		return nil, parseErr(t, s)
	}
	return r, nil
}

func (t BigIntType) ParseGeneric(s string) (Generic, error) {
	v, err := t.Parse(s)
	if err != nil {
		return Nil, err
	}
	return t.Generic(v), nil
}

func (t BigIntType) Format(v *big.Int) string {
	return v.String()
}

func (t BigIntType) Generic(i *big.Int) Generic {
	return gBigInt{val: i}
}

func (t BigIntType) Value(v Generic) *big.Int {
	return v.Value().(*big.Int)
}
