package zc

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/blackchip-org/zc/pkg/scanner"
	"github.com/shopspring/decimal"
)

// ---

type DecimalType struct{}

func (t DecimalType) String() string { return "Decimal" }

func (t DecimalType) Parse(s string) (decimal.Decimal, error) {
	ls := strings.ToLower(s)
	if !strings.HasSuffix(ls, "d") {
		// If scientific notation is being used, let this be parsed
		// by the float type instead
		if strings.Contains(ls, "e") {
			return decimal.Zero, ErrExpectedType(t, s)
		}
	}
	s = strings.TrimSuffix(s, "d")
	s = cleanNumber(s)
	d, err := decimal.NewFromString(s)
	if err != nil {
		return decimal.Zero, ErrExpectedType(t, s)
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
		return 0, ErrExpectedType(t, s)
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
		return &big.Rat{}, ErrExpectedType(t, s)
	}
	switch sc.Ch {
	case '_', '-', ' ':
		w = i1
	case '/':
		n = i1
	default:
		return &big.Rat{}, ErrExpectedType(t, s)
	}
	sc.Next()

	s2 := sc.Scan(scanner.UInt)
	i2, err := strconv.ParseInt(s2, 10, 64)
	if err != nil {
		return &big.Rat{}, ErrExpectedType(t, s)
	}
	if w != 0 {
		n = i2
		if sc.Ch != '/' {
			return &big.Rat{}, ErrExpectedType(t, s)
		}
		sc.Next()
		s3 := sc.Scan(scanner.UInt)
		i3, err := strconv.ParseInt(s3, 10, 64)
		if err != nil {
			return &big.Rat{}, ErrExpectedType(t, s)
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
