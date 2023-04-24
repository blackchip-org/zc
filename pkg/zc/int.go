package zc

import (
	"fmt"
	"math/big"
	"strconv"
)

type BigIntType struct{}

func (t BigIntType) String() string { return "BigInt" }

func (t BigIntType) Parse(s string) (*big.Int, bool) {
	s = cleanNumber(s)
	var r big.Int
	_, ok := r.SetString(s, 0)
	if !ok {
		return nil, false
	}
	return &r, true
}

func (t BigIntType) MustParse(s string) *big.Int {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
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

func PopBigInt(c Calc) *big.Int     { return BigInt.MustParse(c.MustPop()) }
func PushBigInt(c Calc, r *big.Int) { c.Push(BigInt.Format(r)) }

// ---

type IntType struct{}

func (t IntType) String() string { return "Int" }

func (t IntType) Parse(s string) (int, bool) {
	r, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, false
	}
	return int(r), true
}

func (t IntType) MustParse(s string) int {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
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

func PopInt(c Calc) int     { return Int.MustParse(c.MustPop()) }
func PushInt(c Calc, r int) { c.Push(Int.Format(r)) }

// ---

type Int64Type struct{}

func (t Int64Type) String() string { return "Int64" }

func (t Int64Type) Parse(s string) (int64, bool) {
	r, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, false
	}
	return int64(r), true
}

func (t Int64Type) MustParse(s string) int64 {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t Int64Type) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t Int64Type) Format(v int64) string {
	return fmt.Sprintf("%v", v)
}

func PopInt64(c Calc) int64     { return Int64.MustParse(c.MustPop()) }
func PushInt64(c Calc, r int64) { c.Push(Int64.Format(r)) }

// ---

type Int32Type struct{}

func (t Int32Type) String() string { return "Int32" }

func (t Int32Type) Parse(s string) (int32, bool) {
	r, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, false
	}
	return int32(r), true
}

func (t Int32Type) MustParse(s string) int32 {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t Int32Type) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t Int32Type) Format(v int32) string {
	return fmt.Sprintf("%v", v)
}

func PopInt32(c Calc) int32     { return Int32.MustParse(c.MustPop()) }
func PushInt32(c Calc, r int32) { c.Push(Int32.Format(r)) }

// ---

type UintType struct{}

func (t UintType) String() string { return "Uint" }

func (t UintType) Parse(s string) (uint, bool) {
	r, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		return 0, false
	}
	return uint(r), true
}

func (t UintType) MustParse(s string) uint {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t UintType) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t UintType) Format(v uint) string {
	return fmt.Sprintf("%v", v)
}

func PopUint(c Calc) uint     { return Uint.MustParse(c.MustPop()) }
func PushUint(c Calc, r uint) { c.Push(Uint.Format(r)) }

// ---

type Uint8Type struct{}

func (t Uint8Type) String() string { return "Uint8" }

func (t Uint8Type) Parse(s string) (uint8, bool) {
	r, err := strconv.ParseUint(s, 0, 8)
	if err != nil {
		return 0, false
	}
	return uint8(r), true
}

func (t Uint8Type) MustParse(s string) uint8 {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t Uint8Type) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t Uint8Type) Format(v uint8) string {
	return fmt.Sprintf("%v", v)
}

func PopUint8(c Calc) uint8     { return Uint8.MustParse(c.MustPop()) }
func PushUint8(c Calc, r uint8) { c.Push(Uint8.Format(r)) }
