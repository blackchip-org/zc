package types

import (
	"fmt"

	"github.com/blackchip-org/zc/scanner"
)

type Type interface {
	String() string
	ParseValue(string) (Value, error)
}

type Value interface {
	String() string
	Format() string
	Type() Type
	Native() any
}

var (
	BigInt   = BigIntType{}
	Bool     = BoolType{}
	Complex  = ComplexType{}
	Decimal  = DecimalType{}
	Float    = FloatType{}
	Int      = IntType{}
	Int8     = Int8Type{}
	Int16    = Int16Type{}
	Int32    = Int32Type{}
	Int64    = Int64Type{}
	None     = noneType{}
	Rational = RationalType{}
	Rune     = RuneType{}
	String   = StringType{}
	Uint     = UintType{}
	Uint8    = Uint8Type{}
	Uint16   = Uint16Type{}
	Uint32   = Uint32Type{}
	Uint64   = Uint64Type{}
)

var Nil Value = vNone{}

func Is(v string, t Type) bool {
	_, err := t.ParseValue(v)
	return err == nil
}

func To(v Value, t Type) (Value, error) {
	if v.Type() == t {
		return v, nil
	}
	r, err := t.ParseValue(v.Format())
	if err != nil {
		return Nil, fmt.Errorf("expecting %v but got %v", t, QuoteFunc(v.Format()))
	}
	return r, nil
}

func MustParse(s string, t Type) Value {
	v, err := t.ParseValue(s)
	if err != nil {
		panic(err)
	}
	return v
}

var GenericTypes = []Type{
	BigInt,
	Decimal,
	Float,
	Rational,
	Complex,
	Bool,
	String,
}

func Parse(s string) Value {
	for _, t := range GenericTypes {
		v, err := t.ParseValue(s)
		if err == nil {
			return v
		}
	}
	panic("unreachable code")
}

func ParseN(ss []string) []Value {
	var r []Value
	for _, s := range ss {
		r = append(r, Parse(s))
	}
	return r
}

func FormatN(gs []Value) []string {
	var r []string
	for _, g := range gs {
		r = append(r, g.Format())
	}
	return r
}

var QuoteFunc = func(s string) string {
	return s
}

func parseErr(t Type, s string) error {
	return fmt.Errorf("expecting %v but got %v", t, QuoteFunc(s))
}

func isFormatting(ch rune) bool {
	if ch == ',' || ch == '_' || ch == ' ' {
		return true
	}
	if scanner.IsCurrency(ch) {
		return true
	}
	return false
}

func cleanNumber(str string) string {
	var s scanner.Scanner
	s.SetString("", str)
	for s.Ok() {
		s.SkipIf(isFormatting)
	}
	return s.Token()
}
