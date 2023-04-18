package types

import "github.com/blackchip-org/zc/number"

type vComplex struct {
	val number.Complex
}

func (v vComplex) Type() Type     { return Complex }
func (v vComplex) String() string { return Complex.Format(v.val) }
func (v vComplex) Native() any    { return v.val }

type ComplexType struct{}

func (t ComplexType) String() string { return "Complex" }

func (t ComplexType) Parse(s string) (number.Complex, error) {
	return number.ParseComplex(s)
}

func (t ComplexType) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return nil, err
	}
	return t.Value(v), nil
}

func (t ComplexType) Format(c number.Complex) string {
	return c.String()
}

func (t ComplexType) Value(c number.Complex) Value {
	return vComplex{val: c}
}

func (t ComplexType) Native(v Value) number.Complex {
	return v.Native().(number.Complex)
}
