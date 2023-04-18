package types

import "github.com/blackchip-org/zc/number"

type vBigInt struct {
	val number.BigInt
}

func (v vBigInt) Type() Type     { return BigInt }
func (v vBigInt) String() string { return BigInt.Format(v.val) }
func (v vBigInt) Native() any    { return v.val }

type BigIntType struct{}

func (t BigIntType) String() string { return "BigInt" }

func (t BigIntType) Parse(s string) (number.BigInt, error) {
	return number.ParseBigInt(cleanNumber(s))
}

func (t BigIntType) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return nil, err
	}
	return t.Value(v), nil
}

func (t BigIntType) Format(i number.BigInt) string {
	return i.String()
}

func (t BigIntType) Value(i number.BigInt) Value {
	return vBigInt{val: i}
}

func (t BigIntType) Native(v Value) number.BigInt {
	return v.Native().(number.BigInt)
}
