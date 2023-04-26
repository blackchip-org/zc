package zc

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/blackchip-org/zc/pkg/scanner"
	"github.com/shopspring/decimal"
)

// ---

type DecimalType struct{}

func (t DecimalType) String() string { return "Decimal" }

func (t DecimalType) Parse(s string) (decimal.Decimal, bool) {
	ls := strings.ToLower(s)
	if !strings.HasSuffix(ls, "d") {
		// If scientific notation is being used, let this be parsed
		// by the float type instead
		if strings.Contains(ls, "e") {
			return decimal.Zero, false
		}
	}
	s = strings.TrimSuffix(s, "d")
	s = cleanNumber(s)
	d, err := decimal.NewFromString(s)
	if err != nil {
		return decimal.Zero, false
	}
	return d, true
}

func (t DecimalType) MustParse(s string) decimal.Decimal {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t DecimalType) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t DecimalType) Format(v decimal.Decimal) string {
	return v.String()
}

func PopDecimal(c Calc) decimal.Decimal     { return Decimal.MustParse(c.MustPop()) }
func PushDecimal(c Calc, r decimal.Decimal) { c.Push(Decimal.Format(r)) }

// ---

type FloatType struct{}

func (t FloatType) String() string { return "Float" }

func (t FloatType) Parse(s string) (float64, bool) {
	s = cleanNumber(s)
	s = strings.TrimSuffix(s, "f")
	s = strings.Replace(s, "×10", "e", 1)
	s = strings.Replace(s, "x10", "e", 1)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, false
	}
	return f, true
}

func (t FloatType) MustParse(s string) float64 {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t FloatType) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t FloatType) Format(v float64) string {
	s := fmt.Sprintf("%v", v)
	return strings.Replace(s, "e+", "e", 1)
}

func PopFloat(c Calc) float64 { return Float.MustParse(c.MustPop()) }
func PushFloat(c Calc, r float64) {
	if math.IsNaN(r) {
		ErrNotANumber(c)
	} else if math.IsInf(r, 1) {
		ErrInfinity(c, 1)
	} else if math.IsInf(r, -1) {
		ErrInfinity(c, -1)
	} else {
		c.Push(Float.Format(r))
	}
}

// ---

type RationalType struct{}

func (t RationalType) String() string { return "Rational" }

func (t RationalType) Parse(s string) (*big.Rat, bool) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return big.NewRat(i, 1), true
	}
	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return new(big.Rat).SetFloat64(f), true
	}

	sc := scanner.NewString(s)

	var w, n, d int64

	s1 := sc.Scan(scanner.Int)
	i1, err := strconv.ParseInt(s1, 10, 64)
	if err != nil {
		return nil, false
	}
	switch sc.Ch {
	case '_', '-', ' ':
		w = i1
	case '/':
		n = i1
	default:
		return nil, false
	}
	sc.Next()

	s2 := sc.Scan(scanner.UInt)
	i2, err := strconv.ParseInt(s2, 10, 64)
	if err != nil {
		return nil, false
	}
	if w != 0 {
		n = i2
		if sc.Ch != '/' {
			return nil, false
		}
		sc.Next()
		s3 := sc.Scan(scanner.UInt)
		i3, err := strconv.ParseInt(s3, 10, 64)
		if err != nil {
			return nil, false
		}
		d = i3
	} else {
		d = i2
	}

	if w != 0 {
		n = n + (d * w)
	}

	return big.NewRat(n, d), true
}

func (t RationalType) MustParse(s string) *big.Rat {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t RationalType) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
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

func PopRational(c Calc) *big.Rat     { return Rational.MustParse(c.MustPop()) }
func PushRational(c Calc, r *big.Rat) { c.Push(Rational.Format(r)) }
