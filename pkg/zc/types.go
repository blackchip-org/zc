package zc

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"unicode"

	"github.com/blackchip-org/zc/pkg/types"
)

type Type interface {
	String() string
	Is(string) bool
}

var (
	BigInt   = BigIntType{}
	Bool     = BoolType{}
	Char     = CharType{}
	Complex  = ComplexType{}
	Date     = DateType{}
	DateTime = DateTimeType{}
	Decimal  = DecimalType{}
	DMS      = DMSType{}
	Duration = DurationType{}
	Float    = FloatType{}
	Int      = IntType{}
	Int64    = Int64Type{}
	Int32    = Int32Type{}
	Rational = RationalType{}
	Str      = StrType{}
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
	case types.Decimal:
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
		return Str.Format(t)
	case uint:
		return Uint.Format(t)
	}
	return fmt.Sprintf("%v", a)
}

// ---

type BoolType struct{}

func (t BoolType) String() string { return "Bool" }

func (t BoolType) Parse(s string) (bool, bool) {
	ls := strings.TrimSpace(strings.ToLower(s))
	switch ls {
	case "true":
		return true, true
	case "false":
		return false, true
	}
	return false, false
}

func (t BoolType) MustParse(s string) bool {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t BoolType) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t BoolType) Format(v bool) string {
	if v {
		return "true"
	}
	return "false"
}

func PopBool(c Calc) bool     { return Bool.MustParse(c.MustPop()) }
func PushBool(c Calc, r bool) { c.Push(Bool.Format(r)) }

// ---

type ComplexType struct{}

func (t ComplexType) String() string { return "Complex" }

func (t ComplexType) Parse(s string) (complex128, bool) {
	if !strings.HasSuffix(s, "i") {
		return 0, false
	}
	c, err := strconv.ParseComplex(s, 128)
	if err != nil {
		return 0, false
	}
	return c, true
}

func (t ComplexType) MustParse(s string) complex128 {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t ComplexType) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t ComplexType) Format(v complex128) string {
	s := strconv.FormatComplex(v, 'g', 16, 128)
	// For some reason, the complex number is surrounded by parens.
	// Remove them.
	return s[1 : len(s)-1]
}

func PopComplex(c Calc) complex128     { return Complex.MustParse(c.MustPop()) }
func PushComplex(c Calc, r complex128) { c.Push(Complex.Format(r)) }

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
