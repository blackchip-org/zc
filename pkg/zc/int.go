package zc

import (
	"cmp"
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

func (t BigIntType) Compare(x1 string, x2 string) (int, bool) {
	i1, ok := t.Parse(x1)
	if !ok {
		return 0, false
	}
	i2, ok := t.Parse(x2)
	if !ok {
		return 0, false
	}
	return i1.Cmp(i2), true
}

func PopBigInt(c Calc) *big.Int     { return BigInt.MustParse(c.MustPop()) }
func PushBigInt(c Calc, r *big.Int) { c.Push(BigInt.Format(r)) }

// ---

type IntType struct{}

func (t IntType) String() string { return "Int" }

func (t IntType) Parse(s string) (int, bool) {
	s = cleanNumber(s)
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

func (t IntType) Compare(x1 string, x2 string) (int, bool) {
	i1, ok := t.Parse(x1)
	if !ok {
		return 0, false
	}
	i2, ok := t.Parse(x2)
	if !ok {
		return 0, false
	}
	return cmp.Compare(i1, i2), true
}

func PopInt(c Calc) int     { return Int.MustParse(c.MustPop()) }
func PushInt(c Calc, r int) { c.Push(Int.Format(r)) }

// ---

type Int64Type struct{}

func (t Int64Type) String() string { return "Int64" }

func (t Int64Type) Parse(s string) (int64, bool) {
	s = cleanNumber(s)
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

func (t Int64Type) Compare(x1 string, x2 string) (int, bool) {
	i1, ok := t.Parse(x1)
	if !ok {
		return 0, false
	}
	i2, ok := t.Parse(x2)
	if !ok {
		return 0, false
	}
	return cmp.Compare(i1, i2), true
}

func PopInt64(c Calc) int64     { return Int64.MustParse(c.MustPop()) }
func PushInt64(c Calc, r int64) { c.Push(Int64.Format(r)) }

// ---

type Int32Type struct{}

func (t Int32Type) String() string { return "Int32" }

func (t Int32Type) Parse(s string) (int32, bool) {
	s = cleanNumber(s)
	r, err := strconv.ParseInt(s, 0, 32)
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

func (t Int32Type) Compare(x1 string, x2 string) (int, bool) {
	i1, ok := t.Parse(x1)
	if !ok {
		return 0, false
	}
	i2, ok := t.Parse(x2)
	if !ok {
		return 0, false
	}
	return cmp.Compare(i1, i2), true
}

func PopInt32(c Calc) int32     { return Int32.MustParse(c.MustPop()) }
func PushInt32(c Calc, r int32) { c.Push(Int32.Format(r)) }

// ---

type Int16Type struct{}

func (t Int16Type) String() string { return "Int16" }

func (t Int16Type) Parse(s string) (int16, bool) {
	s = cleanNumber(s)
	r, err := strconv.ParseInt(s, 0, 16)
	if err != nil {
		return 0, false
	}
	return int16(r), true
}

func (t Int16Type) MustParse(s string) int16 {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t Int16Type) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t Int16Type) Format(v int16) string {
	return fmt.Sprintf("%v", v)
}

func (t Int16Type) Compare(x1 string, x2 string) (int, bool) {
	i1, ok := t.Parse(x1)
	if !ok {
		return 0, false
	}
	i2, ok := t.Parse(x2)
	if !ok {
		return 0, false
	}
	return cmp.Compare(i1, i2), true
}

func PopInt16(c Calc) int16     { return Int16.MustParse(c.MustPop()) }
func PushInt16(c Calc, r int16) { c.Push(Int16.Format(r)) }

// ---

type Int8Type struct{}

func (t Int8Type) String() string { return "Int8" }

func (t Int8Type) Parse(s string) (int8, bool) {
	s = cleanNumber(s)
	r, err := strconv.ParseInt(s, 0, 8)
	if err != nil {
		return 0, false
	}
	return int8(r), true
}

func (t Int8Type) MustParse(s string) int8 {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t Int8Type) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t Int8Type) Format(v int8) string {
	return fmt.Sprintf("%v", v)
}

func (t Int8Type) Compare(x1 string, x2 string) (int, bool) {
	i1, ok := t.Parse(x1)
	if !ok {
		return 0, false
	}
	i2, ok := t.Parse(x2)
	if !ok {
		return 0, false
	}
	return cmp.Compare(i1, i2), true
}

func PopInt8(c Calc) int8     { return Int8.MustParse(c.MustPop()) }
func PushInt8(c Calc, r int8) { c.Push(Int8.Format(r)) }

// ---

type UintType struct{}

func (t UintType) String() string { return "Uint" }

func (t UintType) Parse(s string) (uint, bool) {
	s = cleanNumber(s)
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

func (t UintType) Compare(x1 string, x2 string) (int, bool) {
	i1, ok := t.Parse(x1)
	if !ok {
		return 0, false
	}
	i2, ok := t.Parse(x2)
	if !ok {
		return 0, false
	}
	return cmp.Compare(i1, i2), true
}

func PopUint(c Calc) uint     { return Uint.MustParse(c.MustPop()) }
func PushUint(c Calc, r uint) { c.Push(Uint.Format(r)) }

// ---

type Uint64Type struct{}

func (t Uint64Type) String() string { return "Uint64" }

func (t Uint64Type) Parse(s string) (uint64, bool) {
	s = cleanNumber(s)
	r, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		return 0, false
	}
	return uint64(r), true
}

func (t Uint64Type) MustParse(s string) uint64 {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t Uint64Type) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t Uint64Type) Format(v uint64) string {
	return fmt.Sprintf("%v", v)
}

func (t Uint64Type) Compare(x1 string, x2 string) (int, bool) {
	i1, ok := t.Parse(x1)
	if !ok {
		return 0, false
	}
	i2, ok := t.Parse(x2)
	if !ok {
		return 0, false
	}
	return cmp.Compare(i1, i2), true
}

func PopUint64(c Calc) uint64     { return Uint64.MustParse(c.MustPop()) }
func PushUint64(c Calc, r uint64) { c.Push(Uint64.Format(r)) }

// ---

type Uint32Type struct{}

func (t Uint32Type) String() string { return "Uint32" }

func (t Uint32Type) Parse(s string) (uint32, bool) {
	s = cleanNumber(s)
	r, err := strconv.ParseUint(s, 0, 32)
	if err != nil {
		return 0, false
	}
	return uint32(r), true
}

func (t Uint32Type) MustParse(s string) uint32 {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t Uint32Type) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t Uint32Type) Format(v uint32) string {
	return fmt.Sprintf("%v", v)
}

func (t Uint32Type) Compare(x1 string, x2 string) (int, bool) {
	i1, ok := t.Parse(x1)
	if !ok {
		return 0, false
	}
	i2, ok := t.Parse(x2)
	if !ok {
		return 0, false
	}
	return cmp.Compare(i1, i2), true
}

func PopUint32(c Calc) uint32     { return Uint32.MustParse(c.MustPop()) }
func PushUint32(c Calc, r uint32) { c.Push(Uint32.Format(r)) }

// ---

type Uint16Type struct{}

func (t Uint16Type) String() string { return "Uint16" }

func (t Uint16Type) Parse(s string) (uint16, bool) {
	s = cleanNumber(s)
	r, err := strconv.ParseUint(s, 0, 16)
	if err != nil {
		return 0, false
	}
	return uint16(r), true
}

func (t Uint16Type) MustParse(s string) uint16 {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t Uint16Type) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t Uint16Type) Format(v uint16) string {
	return fmt.Sprintf("%v", v)
}

func (t Uint16Type) Compare(x1 string, x2 string) (int, bool) {
	i1, ok := t.Parse(x1)
	if !ok {
		return 0, false
	}
	i2, ok := t.Parse(x2)
	if !ok {
		return 0, false
	}
	return cmp.Compare(i1, i2), true
}

func PopUint16(c Calc) uint16     { return Uint16.MustParse(c.MustPop()) }
func PushUint16(c Calc, r uint16) { c.Push(Uint16.Format(r)) }

// ---

type Uint8Type struct{}

func (t Uint8Type) String() string { return "Uint8" }

func (t Uint8Type) Parse(s string) (uint8, bool) {
	s = cleanNumber(s)
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

func (t Uint8Type) Compare(x1 string, x2 string) (int, bool) {
	i1, ok := t.Parse(x1)
	if !ok {
		return 0, false
	}
	i2, ok := t.Parse(x2)
	if !ok {
		return 0, false
	}
	return cmp.Compare(i1, i2), true
}

func PopUint8(c Calc) uint8     { return Uint8.MustParse(c.MustPop()) }
func PushUint8(c Calc, r uint8) { c.Push(Uint8.Format(r)) }
