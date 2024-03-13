package zc

import (
	"cmp"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v5/pkg/types"
)

// ---

type DecimalType struct{}

func (t DecimalType) String() string { return "Decimal" }

func (t DecimalType) Parse(s string) (types.Decimal, bool) {
	ls := strings.ToLower(s)
	if !strings.HasSuffix(ls, "d") {
		// If scientific notation is being used, let this be parsed
		// by the float type instead
		if strings.Contains(ls, "e") {
			return types.DecimalZero, false
		}
	}
	s = strings.TrimSuffix(s, "d")
	s = cleanNumber(s)
	d, err := types.NewDecimalFromString(s)
	if err != nil {
		return types.DecimalZero, false
	}
	return d, true
}

func (t DecimalType) MustParse(s string) types.Decimal {
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

func (t DecimalType) Format(v types.Decimal) string {
	return v.String()
}

func (t DecimalType) Compare(x1 string, x2 string) (int, bool) {
	d1, ok := t.Parse(x1)
	if !ok {
		return 0, false
	}
	d2, ok := t.Parse(x2)
	if !ok {
		return 0, false
	}
	return d1.Cmp(d2), true
}

func PopDecimal(c Calc) types.Decimal     { return Decimal.MustParse(c.MustPop()) }
func PushDecimal(c Calc, r types.Decimal) { c.Push(Decimal.Format(r)) }

// ---

type FloatType struct{}

func (t FloatType) String() string { return "Float" }

func (t FloatType) Parse(s string) (float64, bool) {
	s = PreParseFloat(s, "f")
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

func (t FloatType) Compare(x1 string, x2 string) (int, bool) {
	f1, ok := t.Parse(x1)
	if !ok {
		return 0, false
	}
	f2, ok := t.Parse(x2)
	if !ok {
		return 0, false
	}
	return cmp.Compare(f1, f2), true
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

	sc := scan.NewScannerFromString("", s)

	var sign, whole, num, denom int64

	scan.SignedIntRule.Eval(sc)
	s1 := sc.Emit().Val
	i1, err := strconv.ParseInt(s1, 10, 64)
	if err != nil {
		return nil, false
	}
	if i1 < 0 {
		sign = -1
		i1 = i1 * -1
	} else {
		sign = 1
	}
	switch sc.This {
	case '_', '-', ' ':
		whole = i1
	case '/':
		num = i1
	default:
		return nil, false
	}
	sc.Skip()

	scan.IntRule.Eval(sc)
	s2 := sc.Emit().Val
	i2, err := strconv.ParseInt(s2, 10, 64)
	if err != nil {
		return nil, false
	}
	if whole != 0 {
		num = i2
		if sc.This != '/' {
			return nil, false
		}
		sc.Skip()
		scan.IntRule.Eval(sc)
		s3 := sc.Emit().Val
		i3, err := strconv.ParseInt(s3, 10, 64)
		if err != nil {
			return nil, false
		}
		denom = i3
	} else {
		denom = i2
	}

	if whole != 0 {
		num = num + (denom * whole)
	}
	num = num * sign
	return big.NewRat(num, denom), true
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

func (t RationalType) Compare(x1 string, x2 string) (int, bool) {
	r1, ok := t.Parse(x1)
	if !ok {
		return 0, false
	}
	r2, ok := t.Parse(x2)
	if !ok {
		return 0, false
	}
	return r1.Cmp(r2), true
}

func (t RationalType) Format(v *big.Rat) string {
	n := v.Num().Int64()
	d := v.Denom().Int64()

	if n > d {
		w := n / d
		n := n % d
		if n == 0 && d == 1 {
			return fmt.Sprintf("%v", w)
		}
		return fmt.Sprintf("%v %v/%v", w, n, d)
	}
	return v.RatString()
}

func PopRational(c Calc) *big.Rat     { return Rational.MustParse(c.MustPop()) }
func PushRational(c Calc, r *big.Rat) { c.Push(Rational.Format(r)) }

// ---

type BigFloatType struct {
	Precision    uint
	RoundingMode big.RoundingMode
}

func (t BigFloatType) String() string { return "BigFloat" }

func (t BigFloatType) Parse(s string) (*big.Float, bool) {
	s = PreParseFloat(s, "bf")
	r, _, err := big.ParseFloat(s, 0, t.Precision, t.RoundingMode)
	if err != nil {
		return nil, false
	}
	return r, true
}

func (t BigFloatType) MustParse(s string) *big.Float {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t BigFloatType) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t BigFloatType) Format(v *big.Float) string {
	s := v.String()
	return PostFormatFloat(s)
}

func (t BigFloatType) Compare(x1 string, x2 string) (int, bool) {
	f1, ok := t.Parse(x1)
	if !ok {
		return 0, false
	}
	f2, ok := t.Parse(x2)
	if !ok {
		return 0, false
	}
	return f1.Cmp(f2), true
}

func PopBigFloat(c Calc) *big.Float { return BigFloat.MustParse(c.MustPop()) }
func PushBigFloat(c Calc, r *big.Float) {
	c.Push(BigFloat.Format(r))
}
