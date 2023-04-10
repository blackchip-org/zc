package types

import (
	"math/big"
)

type vBigInt struct {
	val *big.Int
}

func (v vBigInt) Type() Type     { return BigInt }
func (v vBigInt) Format() string { return BigInt.Format(v.val) }
func (v vBigInt) String() string { return stringV(v) }
func (v vBigInt) Native() any    { return v.val }

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

func (t BigIntType) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return Nil, err
	}
	return t.Value(v), nil
}

func (t BigIntType) Format(v *big.Int) string {
	return v.String()
}

func (t BigIntType) Value(i *big.Int) Value {
	return vBigInt{val: i}
}

func (t BigIntType) Native(v Value) *big.Int {
	return v.Native().(*big.Int)
}
