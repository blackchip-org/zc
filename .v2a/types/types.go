package types

import (
	"errors"
	"fmt"

	"github.com/blackchip-org/zc/scanner"
)

var (
	ErrNotSupported = errors.New("not supported")
)

type Type interface {
	String() string
	ParseValue(string) (Value, error)
}

type Value interface {
	String() string
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
	Rational = RationalType{}
	Rune     = RuneType{}
	String   = StringType{}
	Uint     = UintType{}
	Uint8    = Uint8Type{}
	Uint16   = Uint16Type{}
	Uint32   = Uint32Type{}
	Uint64   = Uint64Type{}
)

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
	s.SetString(str)
	for s.Ok() {
		s.SkipIf(isFormatting)
	}
	return s.Token()
}

var QuoteFunc = func(s string) string { return "'" + s + "'" }

func parseErr(t Type, s string) error {
	return fmt.Errorf("expecting %v but got %v", t, QuoteFunc(s))
}
