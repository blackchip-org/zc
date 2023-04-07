package types

import "fmt"

type Type interface {
	String() string
	ParseValue(string) (Value, bool)
}

type Value interface {
	String() string
	Format() string
	Type() Type
}

var (
	BigInt  = BigIntType{}
	Bool    = BoolType{}
	Complex = ComplexType{}
	Decimal = DecimalType{}
	Float   = FloatType{}
	None    = noneType{}
)

var Nil Value = noneVal{}

func Is(v string, t Type) bool {
	_, ok := t.ParseValue(v)
	return ok
}

func To(v Value, t Type) (Value, bool) {
	if v.Type() == t {
		return v, true
	}
	return t.ParseValue(v.Format())
}

func MustParseValue(s string, t Type) Value {
	v, ok := t.ParseValue(s)
	if !ok {
		panic("unable to parse " + t.String())
	}
	return v
}

var NumberTypes = []Type{
	BigInt,
	Decimal,
	Float,
	Complex,
}

func ParseNumber(s string) (Value, bool) {
	for _, t := range NumberTypes {
		v, ok := t.ParseValue(s)
		if ok {
			return v, true
		}
	}
	return Nil, false
}

func ParseNumbers(ss []string) ([]Value, error) {
	var r []Value
	for _, s := range ss {
		n, ok := ParseNumber(s)
		if !ok {
			return []Value{}, fmt.Errorf("not a number: %v", s)
		}
		r = append(r, n)
	}
	return r, nil
}

func MustParseNumbers(ss []string) []Value {
	r, err := ParseNumbers(ss)
	if err != nil {
		panic(err.Error())
	}
	return r
}

func FormatValues(vs []Value) []string {
	var r []string
	for _, v := range vs {
		r = append(r, v.Format())
	}
	return r
}
