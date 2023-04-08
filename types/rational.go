package types

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/blackchip-org/zc/scanner"
)

var ratZero big.Rat

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
func (g gRational) Value() any     { return g.val }

type RationalType struct{}

func (t RationalType) String() string { return "Rational" }

func (t RationalType) Parse(str string) (*big.Rat, bool) {
	i, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return big.NewRat(i, 1), true
	}

	s := scanner.NewString("", str)

	var w, n, d int64

	s1 := s.Scan(scanner.Int)
	i1, err := strconv.ParseInt(s1, 10, 64)
	if err != nil {
		return &big.Rat{}, false
	}
	switch s.Ch {
	case '_', '-', ' ':
		w = i1
	case '/':
		n = i1
	default:
		return &big.Rat{}, false
	}
	s.Next()

	s2 := s.Scan(scanner.UInt)
	i2, err := strconv.ParseInt(s2, 10, 64)
	if err != nil {
		return &big.Rat{}, false
	}
	if w != 0 {
		n = i2
		if s.Ch != '/' {
			return &big.Rat{}, false
		}
		s.Next()
		s3 := s.Scan(scanner.UInt)
		i3, err := strconv.ParseInt(s3, 10, 64)
		if err != nil {
			return &big.Rat{}, false
		}
		d = i3
	} else {
		d = i2
	}

	fmt.Printf("--- w %v n %v d %v\n", w, n, d)
	if w != 0 {
		n = n + (d * w)
	}

	return big.NewRat(n, d), true
}

func (t RationalType) ParseGeneric(s string) (Generic, bool) {
	v, ok := t.Parse(s)
	if !ok {
		return Nil, false
	}
	return t.Generic(v), true
}

func (t RationalType) Generic(i *big.Rat) Generic {
	return gRational{val: i}
}

func (t RationalType) Value(v Generic) *big.Rat {
	return v.Value().(*big.Rat)
}

func op1Rational(fn func(*big.Rat, *big.Rat) error) OpFn {
	return func(args []Generic) ([]Generic, error) {
		x := Rational.Value(args[0])
		z := new(big.Rat)
		err := fn(z, x)
		return []Generic{Rational.Generic(z)}, err
	}
}

func op2Rational(fn func(*big.Rat, *big.Rat, *big.Rat) error) OpFn {
	return func(args []Generic) ([]Generic, error) {
		x := Rational.Value(args[0])
		y := Rational.Value(args[1])
		z := new(big.Rat)
		err := fn(z, x, y)
		return []Generic{Rational.Generic(z)}, err
	}
}

func divRationalFn(z *big.Rat, x *big.Rat, y *big.Rat) error {
	if y.Cmp(&ratZero) == 0 {
		return ErrDivisionByZero
	}
	z.Quo(x, y)
	return nil
}

var (
	absRational = op1Rational(func(z *big.Rat, x *big.Rat) error { z.Abs(x); return nil })
	addRational = op2Rational(func(z *big.Rat, x *big.Rat, y *big.Rat) error { z.Add(x, y); return nil })
	divRational = op2Rational(divRationalFn)
	mulRational = op2Rational(func(z *big.Rat, x *big.Rat, y *big.Rat) error { z.Mul(x, y); return nil })
	subRational = op2Rational(func(z *big.Rat, x *big.Rat, y *big.Rat) error { z.Sub(x, y); return nil })
)
