package types

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/blackchip-org/zc/scanner"
)

type gRational struct {
	val *big.Rat
}

func formatRational(r *big.Rat) string {
	n := r.Num().Int64()
	d := r.Denom().Int64()

	if n > d {
		w := n / d
		n := n % d
		return fmt.Sprintf("%v %v/%v", w, n, d)
	}
	return r.RatString()
}

func (g gRational) Type() Type     { return Rational }
func (g gRational) Format() string { return formatRational(g.val) }
func (g gRational) String() string { return fmt.Sprintf("%v(%v)", g.Type().String(), g.Format()) }
func (g gRational) Native() any    { return g.val }

type RationalType struct{}

func (t RationalType) String() string { return "Rational" }

func (t RationalType) Parse(str string) (*big.Rat, error) {
	i, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return big.NewRat(i, 1), nil
	}
	f, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return new(big.Rat).SetFloat64(f), nil
	}

	s := scanner.NewString("", str)

	var w, n, d int64

	s1 := s.Scan(scanner.Int)
	i1, err := strconv.ParseInt(s1, 10, 64)
	if err != nil {
		return &big.Rat{}, parseErr(t, str)
	}
	switch s.Ch {
	case '_', '-', ' ':
		w = i1
	case '/':
		n = i1
	default:
		return &big.Rat{}, parseErr(t, str)
	}
	s.Next()

	s2 := s.Scan(scanner.UInt)
	i2, err := strconv.ParseInt(s2, 10, 64)
	if err != nil {
		return &big.Rat{}, parseErr(t, str)
	}
	if w != 0 {
		n = i2
		if s.Ch != '/' {
			return &big.Rat{}, parseErr(t, str)
		}
		s.Next()
		s3 := s.Scan(scanner.UInt)
		i3, err := strconv.ParseInt(s3, 10, 64)
		if err != nil {
			return &big.Rat{}, parseErr(t, str)
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

func (t RationalType) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return Nil, err
	}
	return t.Value(v), nil
}

func (t RationalType) Value(i *big.Rat) Value {
	return gRational{val: i}
}

func (t RationalType) Native(v Value) *big.Rat {
	return v.Native().(*big.Rat)
}
