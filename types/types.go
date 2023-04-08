package types

import "fmt"

type Type interface {
	String() string
	ParseGeneric(string) (Generic, bool)
}

type Generic interface {
	String() string
	Format() string
	Type() Type
	Value() any
}

var (
	BigInt   = BigIntType{}
	Bool     = BoolType{}
	Complex  = ComplexType{}
	Decimal  = DecimalType{}
	Float    = FloatType{}
	None     = noneType{}
	Rational = RationalType{}
)

var Nil Generic = gNone{}

func Is(v string, t Type) bool {
	_, ok := t.ParseGeneric(v)
	return ok
}

func To(v Generic, t Type) (Generic, bool) {
	if v.Type() == t {
		return v, true
	}
	return t.ParseGeneric(v.Format())
}

func MustParseGeneric(s string, t Type) Generic {
	v, ok := t.ParseGeneric(s)
	if !ok {
		panic("unable to parse " + t.String())
	}
	return v
}

var NumberTypes = []Type{
	BigInt,
	Decimal,
	Float,
	Rational,
	Complex,
}

func ParseNumber(s string) (Generic, bool) {
	for _, t := range NumberTypes {
		v, ok := t.ParseGeneric(s)
		if ok {
			return v, true
		}
	}
	return Nil, false
}

func ParseNumbers(ss []string) ([]Generic, error) {
	var r []Generic
	for _, s := range ss {
		n, ok := ParseNumber(s)
		if !ok {
			return []Generic{}, fmt.Errorf("not a number: %v", s)
		}
		r = append(r, n)
	}
	return r, nil
}

func MustParseNumbers(ss []string) []Generic {
	r, err := ParseNumbers(ss)
	if err != nil {
		panic(err.Error())
	}
	return r
}

func FormatGenerics(vs []Generic) []string {
	var r []string
	for _, v := range vs {
		r = append(r, v.Format())
	}
	return r
}
