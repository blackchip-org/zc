package zc

import (
	"fmt"
	"math/big"
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
	Val      = StrType{}
)

type SortInterface []string

func (s SortInterface) Len() int {
	return len(s)
}

func (s SortInterface) Less(i, j int) bool {
	x1 := s[i]
	x2 := s[j]
	if v, ok := BigInt.Compare(x1, x2); ok {
		return v < 0
	}
	if v, ok := Decimal.Compare(x1, x2); ok {
		return v < 0
	}
	if v, ok := Float.Compare(x1, x2); ok {
		return v < 0
	}
	if v, ok := Rational.Compare(x1, x2); ok {
		return v < 0
	}
	if v, ok := DateTime.Compare(x1, x2); ok {
		return v < 0
	}
	if v, ok := Duration.Compare(x1, x2); ok {
		return v < 0
	}
	if v, ok := Str.Compare(x1, x2); ok {
		return v < 0
	}
	panic("unreachable -- Str.Compare should always be valid")
}

func (s SortInterface) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

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
