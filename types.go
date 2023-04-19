package zc

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"unicode"

	"github.com/blackchip-org/zc/scanner"
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
	Decimal  = DecimalType{}
	Float    = FloatType{}
	Int      = IntType{}
	Rational = RationalType{}
	String   = StringType{}
)

type BigIntType struct{}

func (t BigIntType) String() string { return "BigInt" }

func (t BigIntType) Parse(s string) (*big.Int, error) {
	s = cleanNumber(s)
	var r big.Int
	_, ok := r.SetString(s, 0)
	if !ok {
		return nil, ErrUnexpectedType(t, s)
	}
	return &r, nil
}

func (t BigIntType) MustParse(s string) *big.Int {
	r, err := t.Parse(s)
	if err != nil {
		panic(err)
	}
	return r
}

func (t BigIntType) Is(s string) bool {
	_, err := t.Parse(s)
	return err == nil
}

func (t BigIntType) Format(v *big.Int) string {
	return v.String()
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
	return false, ErrUnexpectedType(t, s)
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
		return 0, ErrUnexpectedType(t, s)
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

// ---

type DecimalType struct{}

func (t DecimalType) String() string { return "Decimal" }

func (t DecimalType) Parse(s string) (decimal.Decimal, error) {
	ls := strings.ToLower(s)
	if !strings.HasSuffix(ls, "d") {
		// If scientific notation is being used, let this be parsed
		// by the float type instead
		if strings.Contains(ls, "e") {
			return decimal.Zero, ErrUnexpectedType(t, s)
		}
	}
	s = strings.TrimSuffix(s, "d")
	s = cleanNumber(s)
	d, err := decimal.NewFromString(s)
	if err != nil {
		return decimal.Zero, ErrUnexpectedType(t, s)
	}
	return d, nil
}

func (t DecimalType) MustParse(s string) decimal.Decimal {
	z, err := t.Parse(s)
	if err != nil {
		panic(err)
	}
	return z
}

func (t DecimalType) Is(s string) bool {
	_, err := t.Parse(s)
	return err == nil
}

func (t DecimalType) Format(v decimal.Decimal) string {
	return v.String()
}

// ---

type FloatType struct{}

func (t FloatType) String() string { return "Float" }

func (t FloatType) Parse(s string) (float64, error) {
	s = cleanNumber(s)
	s = strings.TrimSuffix(s, "f")
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, ErrUnexpectedType(t, s)
	}
	return f, nil
}

func (t FloatType) MustParse(s string) float64 {
	r, err := t.Parse(s)
	if err != nil {
		panic(err)
	}
	return r
}

func (t FloatType) Is(s string) bool {
	_, err := t.Parse(s)
	return err == nil
}

func (t FloatType) Format(v float64) string {
	return strconv.FormatFloat(v, 'g', 16, 64)
}

// ---

type IntType struct{}

func (t IntType) String() string { return "Int" }

func (t IntType) Parse(s string) (int, error) {
	r, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, ErrUnexpectedType(t, s)
	}
	return int(r), nil
}

func (t IntType) MustParse(s string) int {
	r, err := t.Parse(s)
	if err != nil {
		panic(err)
	}
	return r
}

func (t IntType) Is(s string) bool {
	_, err := t.Parse(s)
	return err == nil
}

func (t IntType) Format(v int) string {
	return fmt.Sprintf("%v", v)
}

// ---

type RationalType struct{}

func (t RationalType) String() string { return "Rational" }

func (t RationalType) Parse(s string) (*big.Rat, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return big.NewRat(i, 1), nil
	}
	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return new(big.Rat).SetFloat64(f), nil
	}

	sc := scanner.NewString(s)

	var w, n, d int64

	s1 := sc.Scan(scanner.Int)
	i1, err := strconv.ParseInt(s1, 10, 64)
	if err != nil {
		return &big.Rat{}, ErrUnexpectedType(t, s)
	}
	switch sc.Ch {
	case '_', '-', ' ':
		w = i1
	case '/':
		n = i1
	default:
		return &big.Rat{}, ErrUnexpectedType(t, s)
	}
	sc.Next()

	s2 := sc.Scan(scanner.UInt)
	i2, err := strconv.ParseInt(s2, 10, 64)
	if err != nil {
		return &big.Rat{}, ErrUnexpectedType(t, s)
	}
	if w != 0 {
		n = i2
		if sc.Ch != '/' {
			return &big.Rat{}, ErrUnexpectedType(t, s)
		}
		sc.Next()
		s3 := sc.Scan(scanner.UInt)
		i3, err := strconv.ParseInt(s3, 10, 64)
		if err != nil {
			return &big.Rat{}, ErrUnexpectedType(t, s)
		}
		d = i3
	} else {
		d = i2
	}

	if w != 0 {
		n = n + (d * w)
	}

	return big.NewRat(n, d), nil
}

func (t RationalType) MustParse(s string) *big.Rat {
	r, err := t.Parse(s)
	if err != nil {
		panic(err)
	}
	return r
}

func (t RationalType) Is(s string) bool {
	_, err := t.Parse(s)
	return err == nil
}

func (t RationalType) Format(v *big.Rat) string {
	n := v.Num().Int64()
	d := v.Denom().Int64()

	if n > d {
		w := n / d
		n := n % d
		return fmt.Sprintf("%v %v/%v", w, n, d)
	}
	return v.RatString()
}

// ---

type StringType struct{}

func (t StringType) String() string { return "String" }

func (t StringType) Parse(s string) (string, error) {
	return s, nil
}

func (t StringType) MustParse(s string) string {
	return s
}

func (t StringType) Is(s string) bool {
	return true
}

func (t StringType) Format(v string) string {
	return v
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
