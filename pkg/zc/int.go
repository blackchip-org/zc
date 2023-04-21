package zc

import (
	"fmt"
	"math/big"
	"strconv"
)

type BigIntType struct{}

func (t BigIntType) String() string { return "BigInt" }

func (t BigIntType) Parse(s string) (*big.Int, error) {
	s = cleanNumber(s)
	var r big.Int
	_, ok := r.SetString(s, 0)
	if !ok {
		return nil, ErrExpectedType(t, s)
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

type IntType struct{}

func (t IntType) String() string { return "Int" }

func (t IntType) Parse(s string) (int, error) {
	r, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, ErrExpectedType(t, s)
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

type Int64Type struct{}

func (t Int64Type) String() string { return "Int64" }

func (t Int64Type) Parse(s string) (int64, error) {
	r, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, ErrExpectedType(t, s)
	}
	return int64(r), nil
}

func (t Int64Type) MustParse(s string) int64 {
	r, err := t.Parse(s)
	if err != nil {
		panic(err)
	}
	return r
}

func (t Int64Type) Is(s string) bool {
	_, err := t.Parse(s)
	return err == nil
}

func (t Int64Type) Format(v int64) string {
	return fmt.Sprintf("%v", v)
}

// ---

type Int32Type struct{}

func (t Int32Type) String() string { return "Int32" }

func (t Int32Type) Parse(s string) (int32, error) {
	r, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, ErrExpectedType(t, s)
	}
	return int32(r), nil
}

func (t Int32Type) MustParse(s string) int32 {
	r, err := t.Parse(s)
	if err != nil {
		panic(err)
	}
	return r
}

func (t Int32Type) Is(s string) bool {
	_, err := t.Parse(s)
	return err == nil
}

func (t Int32Type) Format(v int32) string {
	return fmt.Sprintf("%v", v)
}

// ---

type UintType struct{}

func (t UintType) String() string { return "Uint" }

func (t UintType) Parse(s string) (uint, error) {
	r, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		return 0, ErrExpectedType(t, s)
	}
	return uint(r), nil
}

func (t UintType) MustParse(s string) uint {
	r, err := t.Parse(s)
	if err != nil {
		panic(err)
	}
	return r
}

func (t UintType) Is(s string) bool {
	_, err := t.Parse(s)
	return err == nil
}

func (t UintType) Format(v uint) string {
	return fmt.Sprintf("%v", v)
}

// ---

type Uint8Type struct{}

func (t Uint8Type) String() string { return "Uint8" }

func (t Uint8Type) Parse(s string) (uint8, error) {
	r, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		return 0, ErrExpectedType(t, s)
	}
	return uint8(r), nil
}

func (t Uint8Type) MustParse(s string) uint8 {
	r, err := t.Parse(s)
	if err != nil {
		panic(err)
	}
	return r
}

func (t Uint8Type) Is(s string) bool {
	_, err := t.Parse(s)
	return err == nil
}

func (t Uint8Type) Format(v uint8) string {
	return fmt.Sprintf("%v", v)
}
