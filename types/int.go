package types

import (
	"fmt"
	"strconv"
)

type vInt struct {
	val int
}

func (v vInt) Type() Type     { return IntType{} }
func (v vInt) String() string { return IntType{}.Format(v.val) }
func (v vInt) Native() any    { return v.val }

type IntType struct{}

func (t IntType) String() string { return "Int" }

func (t IntType) Parse(s string) (int, error) {
	s = cleanNumber(s)
	r, err := strconv.ParseInt(s, 0, 0)
	if err != nil {
		return 0, parseErr(t, s)
	}
	return int(r), nil
}

func (t IntType) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return nil, err
	}
	return t.Value(v), nil
}

func (t IntType) Format(v int) string {
	return fmt.Sprintf("%v", v)
}

func (t IntType) Value(i int) Value {
	return vInt{val: i}
}

func (t IntType) Native(v Value) int {
	return v.Native().(int)
}

type vInt8 struct {
	val int8
}

func (v vInt8) Type() Type     { return Int8Type{} }
func (v vInt8) String() string { return Int8Type{}.Format(v.val) }
func (v vInt8) Native() any    { return v.val }

type Int8Type struct{}

func (t Int8Type) String() string { return "Int8" }

func (t Int8Type) Parse(s string) (int8, error) {
	s = cleanNumber(s)
	r, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return 0, parseErr(t, s)
	}
	return int8(r), nil
}

func (t Int8Type) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return nil, err
	}
	return t.Value(v), nil
}

func (t Int8Type) Format(v int8) string {
	return fmt.Sprintf("%v", v)
}

func (t Int8Type) Value(i int8) Value {
	return vInt8{val: i}
}

type vInt16 struct {
	val int16
}

func (v vInt16) Type() Type     { return Int16Type{} }
func (v vInt16) Format() string { return Int16Type{}.Format(v.val) }
func (v vInt16) String() string { return fmt.Sprintf("%v(%v)", v.Type().String(), v.Format()) }
func (v vInt16) Native() any    { return v.val }

type Int16Type struct{}

func (t Int16Type) String() string { return "Int16" }

func (t Int16Type) Parse(s string) (int16, error) {
	s = cleanNumber(s)
	r, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return 0, parseErr(t, s)
	}
	return int16(r), nil
}

func (t Int16Type) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return nil, err
	}
	return t.Value(v), nil
}

func (t Int16Type) Format(v int16) string {
	return fmt.Sprintf("%v", v)
}

func (t Int16Type) Value(i int16) Value {
	return vInt16{val: i}
}

type vInt32 struct {
	val int32
}

func (v vInt32) Type() Type     { return Int32Type{} }
func (v vInt32) Format() string { return Int32Type{}.Format(v.val) }
func (v vInt32) String() string { return fmt.Sprintf("%v(%v)", v.Type().String(), v.Format()) }
func (v vInt32) Native() any    { return v.val }

type Int32Type struct{}

func (t Int32Type) String() string { return "Int32" }

func (t Int32Type) Parse(s string) (int32, error) {
	s = cleanNumber(s)
	r, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return 0, parseErr(t, s)
	}
	return int32(r), nil
}

func (t Int32Type) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return nil, err
	}
	return t.Value(v), nil
}

func (t Int32Type) Format(v int32) string {
	return fmt.Sprintf("%v", v)
}

func (t Int32Type) Value(i int32) Value {
	return vInt32{val: i}
}

type vInt64 struct {
	val int64
}

func (v vInt64) Type() Type     { return Int64Type{} }
func (v vInt64) Format() string { return Int64Type{}.Format(v.val) }
func (v vInt64) String() string { return fmt.Sprintf("%v(%v)", v.Type().String(), v.Format()) }
func (v vInt64) Native() any    { return v.val }

type Int64Type struct{}

func (t Int64Type) String() string { return "Int64" }

func (t Int64Type) Parse(s string) (int64, error) {
	s = cleanNumber(s)
	r, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, parseErr(t, s)
	}
	return int64(r), nil
}

func (t Int64Type) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return nil, err
	}
	return t.Value(v), nil
}

func (t Int64Type) Format(v int64) string {
	return fmt.Sprintf("%v", v)
}

func (t Int64Type) Value(i int64) Value {
	return vInt64{val: i}
}

type vUint struct {
	val uint
}

func (v vUint) Type() Type     { return UintType{} }
func (v vUint) Format() string { return UintType{}.Format(v.val) }
func (v vUint) String() string { return fmt.Sprintf("%v(%v)", v.Type().String(), v.Format()) }
func (v vUint) Native() any    { return v.val }

type UintType struct{}

func (t UintType) String() string { return "Uint" }

func (t UintType) Parse(s string) (uint, error) {
	s = cleanNumber(s)
	r, err := strconv.ParseInt(s, 0, 0)
	if err != nil {
		return 0, parseErr(t, s)
	}
	return uint(r), nil
}

func (t UintType) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return nil, err
	}
	return t.Value(v), nil
}

func (t UintType) Format(v uint) string {
	return fmt.Sprintf("%v", v)
}

func (t UintType) Value(i uint) Value {
	return vUint{val: i}
}

type vUint8 struct {
	val uint8
}

func (v vUint8) Type() Type     { return Uint8Type{} }
func (v vUint8) Format() string { return Uint8Type{}.Format(v.val) }
func (v vUint8) String() string { return fmt.Sprintf("%v(%v)", v.Type().String(), v.Format()) }
func (v vUint8) Native() any    { return v.val }

type Uint8Type struct{}

func (t Uint8Type) String() string { return "Uint8" }

func (t Uint8Type) Parse(s string) (uint8, error) {
	s = cleanNumber(s)
	r, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return 0, parseErr(t, s)
	}
	return uint8(r), nil
}

func (t Uint8Type) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return nil, err
	}
	return t.Value(v), nil
}

func (t Uint8Type) Format(v uint8) string {
	return fmt.Sprintf("%v", v)
}

func (t Uint8Type) Value(i uint8) Value {
	return vUint8{val: i}
}

type vUint16 struct {
	val uint16
}

func (v vUint16) Type() Type     { return Uint16Type{} }
func (v vUint16) Format() string { return Uint16Type{}.Format(v.val) }
func (v vUint16) String() string { return fmt.Sprintf("%v(%v)", v.Type().String(), v.Format()) }
func (v vUint16) Native() any    { return v.val }

type Uint16Type struct{}

func (t Uint16Type) String() string { return "Uint16" }

func (t Uint16Type) Parse(s string) (uint16, error) {
	s = cleanNumber(s)
	r, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return 0, parseErr(t, s)
	}
	return uint16(r), nil
}

func (t Uint16Type) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return nil, err
	}
	return t.Value(v), nil
}

func (t Uint16Type) Format(v uint16) string {
	return fmt.Sprintf("%v", v)
}

func (t Uint16Type) Value(i uint16) Value {
	return vUint16{val: i}
}

type vUint32 struct {
	val uint32
}

func (v vUint32) Type() Type     { return Uint32Type{} }
func (v vUint32) Format() string { return Uint32Type{}.Format(v.val) }
func (v vUint32) String() string { return fmt.Sprintf("%v(%v)", v.Type().String(), v.Format()) }
func (v vUint32) Native() any    { return v.val }

type Uint32Type struct{}

func (t Uint32Type) String() string { return "Uint32" }

func (t Uint32Type) Parse(s string) (uint32, error) {
	s = cleanNumber(s)
	r, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return 0, parseErr(t, s)
	}
	return uint32(r), nil
}

func (t Uint32Type) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return nil, err
	}
	return t.Value(v), nil
}

func (t Uint32Type) Format(v uint32) string {
	return fmt.Sprintf("%v", v)
}

func (t Uint32Type) Value(i uint32) Value {
	return vUint32{val: i}
}

type vUint64 struct {
	val uint64
}

func (v vUint64) Type() Type     { return Uint64Type{} }
func (v vUint64) Format() string { return Uint64Type{}.Format(v.val) }
func (v vUint64) String() string { return fmt.Sprintf("%v(%v)", v.Type().String(), v.Format()) }
func (v vUint64) Native() any    { return v.val }

type Uint64Type struct{}

func (t Uint64Type) String() string { return "Uint64" }

func (t Uint64Type) Parse(s string) (uint64, error) {
	s = cleanNumber(s)
	r, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, parseErr(t, s)
	}
	return uint64(r), nil
}

func (t Uint64Type) Format(v uint64) string {
	return fmt.Sprintf("%v", v)
}

func (t Uint64Type) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return nil, err
	}
	return t.Value(v), nil
}

func (t Uint64Type) Value(i uint64) Value {
	return vUint64{val: i}
}
