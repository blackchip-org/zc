package zc

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"unicode"

	"github.com/shopspring/decimal"
)

type Type interface {
	String() string
	Is(string) bool
}

var (
	BigInt   = BigIntType{}
	Bool     = BoolType{}
	Complex  = ComplexType{}
	Date     = DateType{}
	DateTime = DateTimeType{}
	Decimal  = DecimalType{}
	Duration = DurationType{}
	Float    = FloatType{}
	Int      = IntType{}
	Int64    = Int64Type{}
	Int32    = Int32Type{}
	Rational = RationalType{}
	Rune     = RuneType{}
	String   = StringType{}
	Time     = TimeType{}
	Uint     = UintType{}
	Uint8    = Uint8Type{}
)

func Format(a any) string {
	switch t := a.(type) {
	case *big.Int:
		return BigInt.Format(t)
	case bool:
		return Bool.Format(t)
	case complex128:
		return Complex.Format(t)
	case decimal.Decimal:
		return Decimal.Format(t)
	case float64:
		return Float.Format(t)
	case int:
		return Int.Format(t)
	case int64:
		return Int64.Format(t)
	case *big.Rat:
		return Rational.Format(t)
	case string:
		return String.Format(t)
	case uint:
		return Uint.Format(t)
	}
	return fmt.Sprintf("%v", a)
}

func Identify(v string) Type {
	if BigInt.Is(v) {
		return BigInt
	}
	if Decimal.Is(v) {
		return Decimal
	}
	if Float.Is(v) {
		return Float
	}
	if Rational.Is(v) {
		return Rational
	}
	if Complex.Is(v) {
		return Complex
	}
	if Bool.Is(v) {
		return Bool
	}
	return String
}

// ---

type BoolType struct{}

func (t BoolType) String() string { return "Bool" }

func (t BoolType) Parse(s string) (bool, error) {
	ls := strings.ToLower(s)
	switch ls {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}
	return false, ErrExpectedType(t, s)
}

func (t BoolType) MustParse(s string) bool {
	r, err := t.Parse(s)
	if err != nil {
		panic(err)
	}
	return r
}

func (t BoolType) Is(s string) bool {
	_, err := t.Parse(s)
	return err == nil
}

func (t BoolType) Format(v bool) string {
	if v {
		return "true"
	}
	return "false"
}

// ---

type ComplexType struct{}

func (t ComplexType) String() string { return "Complex" }

func (t ComplexType) Parse(s string) (complex128, error) {
	c, err := strconv.ParseComplex(s, 128)
	if err != nil {
		return 0, ErrExpectedType(t, s)
	}
	return c, nil
}

func (t ComplexType) MustParse(s string) complex128 {
	r, err := t.Parse(s)
	if err != nil {
		panic(err)
	}
	return r
}

func (t ComplexType) Is(s string) bool {
	_, err := t.Parse(s)
	return err == nil
}

func (t ComplexType) Format(v complex128) string {
	s := strconv.FormatComplex(v, 'g', 16, 128)
	// For some reason, the complex number is surrounded by parens.
	// Remove them.
	return s[1 : len(s)-1]
}

// ===

func isFormatting(ch rune) bool {
	if ch == ',' || ch == '_' || ch == ' ' {
		return true
	}
	if unicode.Is(unicode.Sc, ch) {
		return true
	}
	return false
}

func cleanNumber(str string) string {
	var res strings.Builder
	for _, ch := range str {
		if !isFormatting(ch) {
			res.WriteRune(ch)
		}
	}
	return res.String()
}
