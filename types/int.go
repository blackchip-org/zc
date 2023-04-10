package types

import (
	"fmt"
	"strconv"
)

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

func (t IntType) Format(v int) string {
	return fmt.Sprintf("%v", v)
}

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

func (t Int8Type) Format(v int8) string {
	return fmt.Sprintf("%v", v)
}

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

func (t Int16Type) Format(v int16) string {
	return fmt.Sprintf("%v", v)
}

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

func (t Int32Type) Format(v int32) string {
	return fmt.Sprintf("%v", v)
}

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

func (t Int64Type) Format(v int64) string {
	return fmt.Sprintf("%v", v)
}

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

func (t UintType) Format(v uint) string {
	return fmt.Sprintf("%v", v)
}

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

func (t Uint8Type) Format(v uint8) string {
	return fmt.Sprintf("%v", v)
}

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

func (t Uint16Type) Format(v uint16) string {
	return fmt.Sprintf("%v", v)
}

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

func (t Uint32Type) Format(v uint32) string {
	return fmt.Sprintf("%v", v)
}

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
