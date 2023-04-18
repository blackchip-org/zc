package zc

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

type Type interface {
	String() string
	Is(string) bool
}

var (
	BigInt  = BigIntType{}
	Bool    = BoolType{}
	Decimal = DecimalType{}
	Int     = IntType{}
)

type BigIntType struct{}

func (t BigIntType) String() string { return "BigInt" }

func (t BigIntType) Parse(s string) (*big.Int, bool) {
	var r big.Int
	_, ok := r.SetString(s, 0)
	return &r, ok
}

func (t BigIntType) MustParse(s string) *big.Int {
	r, ok := t.Parse(s)
	if !ok {
		panic(ErrUnexpectedType(t, s))
	}
	return r
}

func (t BigIntType) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t BigIntType) Format(v *big.Int) string {
	return v.String()
}

type BoolType struct{}

func (t BoolType) String() string { return "Bool" }

func (t BoolType) Parse(s string) (bool, bool) {
	ls := strings.ToLower(s)
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
		panic(ErrUnexpectedType(t, s))
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

type DecimalType struct{}

func (t DecimalType) String() string { return "Decimal" }

func (t DecimalType) Parse(s string) (decimal.Decimal, bool) {
	d, err := decimal.NewFromString(s)
	return d, err == nil
}

func (t DecimalType) MustParse(s string) decimal.Decimal {
	z, ok := t.Parse(s)
	if !ok {
		panic(ErrUnexpectedType(t, s))
	}
	return z
}

func (t DecimalType) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t DecimalType) Format(v decimal.Decimal) string {
	return v.String()
}

type IntType struct{}

func (t IntType) String() string { return "Int" }

func (t IntType) Parse(s string) (int, bool) {
	r, err := strconv.ParseInt(s, 0, 64)
	return int(r), err == nil
}

func (t IntType) MustParse(s string) int {
	r, ok := t.Parse(s)
	if !ok {
		panic(ErrUnexpectedType(t, s))
	}
	return r
}

func (t IntType) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t IntType) Format(v int) string {
	return fmt.Sprintf("%v", v)
}
