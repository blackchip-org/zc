package types

import "github.com/blackchip-org/zc/number"

type vRational struct {
	val number.Rational
}

func (v vRational) Type() Type     { return RationalType{} }
func (v vRational) String() string { return RationalType{}.Format(v.val) }
func (v vRational) Native() any    { return v.val }

type RationalType struct{}

func (t RationalType) String() string { return "Rational" }

func (t RationalType) Parse(s string) (number.Rational, error) {
	return number.ParseRational(s)
}

func (t RationalType) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return nil, err
	}
	return t.Value(v), nil
}

func (t RationalType) Format(r number.Rational) string {
	return r.String()
}

func (t RationalType) Value(r number.Rational) Value {
	return vRational{val: r}
}

func (t RationalType) Native(v Value) number.Rational {
	return v.Native().(number.Rational)
}
