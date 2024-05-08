package ints

import (
	"fmt"
	"math"
	"strconv"

	"github.com/blackchip-org/zc/v6"
)

var (
	IntArchKind  = intArchKind{}
	Int8Kind     = int8Kind{}
	Int16Kind    = int16Kind{}
	Int32Kind    = int32Kind{}
	Int64Kind    = int64Kind{}
	IntArchUKind = intArchUKind{}
	Int8UKind    = int8UKind{}
	Int16UKind   = int16UKind{}
	Int32UKind   = int32UKind{}
	Int64UKind   = int64UKind{}
)

type intArchKind struct{}

func (k intArchKind) Name() string { return "IntArch" }

func (k intArchKind) Is(a any) bool {
	switch a.(type) {
	case int, *int:
		return true
	}
	return false
}

func (k intArchKind) As(a any) int {
	v, ok := a.(int)
	if !ok {
		panic(fmt.Errorf("expected int but got: %v", zc.TypeName(a)))
	}
	return v
}

func (k intArchKind) Dup(a any) any {
	return a
}

func (k intArchKind) Copy(src, dest any) {
	d, ok := dest.(*int)
	if !ok {
		panic(fmt.Errorf("expected *int but got: %v", zc.TypeName(dest)))
	}
	s := k.As(src)
	*d = s
}

func (k intArchKind) To(a any) (any, bool) {
	switch v := a.(type) {
	case int64:
		if v > math.MaxInt || v < math.MinInt {
			return nil, false
		}
		return int(v), true
	case int32:
		return int(v), true
	case int16:
		return int(v), true
	case int8:
		return int(v), true
	case uint:
		if v < 0 || v > math.MaxUint {
			return nil, false
		}
		return int(v), true
	case uint64:
		if v < 0 {
			return nil, false
		}
		return int(v), true
	case uint32:
		if v < 0 || v > math.MaxUint32 {
			return nil, false
		}
		return int(v), true
	case uint16:
		if v < 0 || v > math.MaxUint16 {
			return nil, false
		}
		return int(v), true
	case uint8:
		if v < 0 || v > math.MaxUint8 {
			return nil, false
		}
		return int(v), true
	case string:
		i, err := strconv.ParseInt(v, 0, 64)
		if err != nil || i < math.MinInt || i > math.MaxInt {
			return nil, false
		}
		return int(i), true
	}
	return nil, false
}

// ----------------------------------------------------------------------------
type int8Kind struct{}

func (k int8Kind) Name() string { return "Int8" }

func (k int8Kind) Is(a any) bool {
	switch a.(type) {
	case int8, *int8:
		return true
	}
	return false
}

func (k int8Kind) As(a any) int8 {
	v, ok := a.(int8)
	if !ok {
		panic(fmt.Errorf("expected int8 but got: %v", zc.TypeName(a)))
	}
	return v
}

func (k int8Kind) Dup(a any) any {
	return a
}

func (k int8Kind) Copy(src, dest any) {
	d, ok := dest.(*int8)
	if !ok {
		panic(fmt.Errorf("expected *int8 but got: %v", zc.TypeName(dest)))
	}
	s := k.As(src)
	*d = s
}

func (k int8Kind) To(a any) (any, bool) {
	switch v := a.(type) {
	case string:
		i, ok := strconv.ParseInt(v, 0, 8)
		return int8(i), ok == nil
	}
	return nil, false
}

// ----------------------------------------------------------------------------
type int16Kind struct{}

func (k int16Kind) Name() string { return "Int16" }

func (k int16Kind) Is(a any) bool {
	switch a.(type) {
	case int16, *int16:
		return true
	}
	return false
}

func (k int16Kind) As(a any) int16 {
	v, ok := a.(int16)
	if !ok {
		panic(fmt.Errorf("expected int16 but got: %v", zc.TypeName(a)))
	}
	return v
}

func (k int16Kind) Dup(a any) any {
	return a
}

func (k int16Kind) Copy(src, dest any) {
	d, ok := dest.(*int16)
	if !ok {
		panic(fmt.Errorf("expected *int16 but got: %v", zc.TypeName(dest)))
	}
	s := k.As(src)
	*d = s
}

func (k int16Kind) To(a any) (any, bool) {
	switch v := a.(type) {
	case string:
		i, ok := strconv.ParseInt(v, 0, 16)
		return int16(i), ok == nil
	}
	return nil, false
}

// ----------------------------------------------------------------------------
type int32Kind struct{}

func (k int32Kind) Name() string { return "Int32" }

func (k int32Kind) Is(a any) bool {
	switch a.(type) {
	case int32, *int32:
		return true
	}
	return false
}

func (k int32Kind) As(a any) int32 {
	v, ok := a.(int32)
	if !ok {
		panic(fmt.Errorf("expected int32 but got: %v", zc.TypeName(a)))
	}
	return v
}

func (k int32Kind) Dup(a any) any {
	return a
}

func (k int32Kind) Copy(src, dest any) {
	d, ok := dest.(*int32)
	if !ok {
		panic(fmt.Errorf("expected *int32 but got: %v", zc.TypeName(dest)))
	}
	s := k.As(src)
	*d = s
}

func (k int32Kind) To(a any) (any, bool) {
	switch v := a.(type) {
	case string:
		i, ok := strconv.ParseInt(v, 0, 32)
		return int32(i), ok == nil
	}
	return nil, false
}

// ----------------------------------------------------------------------------
type int64Kind struct{}

func (k int64Kind) Name() string { return "Int64" }

func (k int64Kind) Is(a any) bool {
	switch a.(type) {
	case int64, *int64:
		return true
	}
	return false
}

func (k int64Kind) As(a any) int64 {
	v, ok := a.(int64)
	if !ok {
		panic(fmt.Errorf("expected int64 but got: %v", zc.TypeName(a)))
	}
	return v
}

func (k int64Kind) Dup(a any) any {
	return a
}

func (k int64Kind) Copy(src, dest any) {
	d, ok := dest.(*int64)
	if !ok {
		panic(fmt.Errorf("expected *int64 but got: %v", zc.TypeName(dest)))
	}
	s := k.As(src)
	*d = s
}

func (k int64Kind) To(a any) (any, bool) {
	switch v := a.(type) {
	case string:
		i, ok := strconv.ParseInt(v, 0, 64)
		return i, ok == nil
	}
	return nil, false
}

// ----------------------------------------------------------------------------

type intArchUKind struct{}

func (k intArchUKind) Name() string { return "IntArchU" }

func (k intArchUKind) Is(a any) bool {
	switch a.(type) {
	case uint, *uint:
		return true
	}
	return false
}

func (k intArchUKind) As(a any) uint {
	v, ok := a.(uint)
	if !ok {
		panic(fmt.Errorf("expected uint but got: %v", zc.TypeName(a)))
	}
	return v
}

func (k intArchUKind) Dup(a any) any {
	return a
}

func (k intArchUKind) Copy(src, dest any) {
	d, ok := dest.(*uint)
	if !ok {
		panic(fmt.Errorf("expected *uint but got: %v", zc.TypeName(dest)))
	}
	s := k.As(src)
	*d = s
}

func (k intArchUKind) To(a any) (any, bool) {
	return nil, false
}

// ----------------------------------------------------------------------------
type int8UKind struct{}

func (k int8UKind) Name() string { return "Int8U" }

func (k int8UKind) Is(a any) bool {
	switch a.(type) {
	case uint8, *uint8:
		return true
	}
	return false
}

func (k int8UKind) As(a any) uint8 {
	v, ok := a.(uint8)
	if !ok {
		panic(fmt.Errorf("expected uint8 but got: %v", zc.TypeName(a)))
	}
	return v
}

func (k int8UKind) Dup(a any) any {
	return a
}

func (k int8UKind) Copy(src, dest any) {
	d, ok := dest.(*uint8)
	if !ok {
		panic(fmt.Errorf("expected *uint8 but got: %v", zc.TypeName(dest)))
	}
	s := k.As(src)
	*d = s
}

func (k int8UKind) To(a any) (any, bool) {
	return nil, false
}

// ----------------------------------------------------------------------------
type int16UKind struct{}

func (k int16UKind) Name() string { return "Int16U" }

func (k int16UKind) Is(a any) bool {
	switch a.(type) {
	case uint16, *uint16:
		return true
	}
	return false
}

func (k int16UKind) As(a any) uint16 {
	v, ok := a.(uint16)
	if !ok {
		panic(fmt.Errorf("expected uint16 but got: %v", zc.TypeName(a)))
	}
	return v
}

func (k int16UKind) Dup(a any) any {
	return a
}

func (k int16UKind) Copy(src, dest any) {
	d, ok := dest.(*uint16)
	if !ok {
		panic(fmt.Errorf("expected *uint16 but got: %v", zc.TypeName(dest)))
	}
	s := k.As(src)
	*d = s
}

func (k int16UKind) To(a any) (any, bool) {
	return nil, false
}

// ----------------------------------------------------------------------------
type int32UKind struct{}

func (k int32UKind) Name() string { return "Int32U" }

func (k int32UKind) Is(a any) bool {
	switch a.(type) {
	case uint32, *uint32:
		return true
	}
	return false
}

func (k int32UKind) As(a any) uint32 {
	v, ok := a.(uint32)
	if !ok {
		panic(fmt.Errorf("expected uint32 but got: %v", zc.TypeName(a)))
	}
	return v
}

func (k int32UKind) Dup(a any) any {
	return a
}

func (k int32UKind) Copy(src, dest any) {
	d, ok := dest.(*uint32)
	if !ok {
		panic(fmt.Errorf("expected *uint32 but got: %v", zc.TypeName(dest)))
	}
	s := k.As(src)
	*d = s
}

func (k int32UKind) To(a any) (any, bool) {
	return nil, false
}

// ----------------------------------------------------------------------------
type int64UKind struct{}

func (k int64UKind) Name() string { return "Int64U" }

func (k int64UKind) Is(a any) bool {
	switch a.(type) {
	case uint64, *uint64:
		return true
	}
	return false
}

func (k int64UKind) As(a any) uint64 {
	v, ok := a.(uint64)
	if !ok {
		panic(fmt.Errorf("expected uint64 but got: %v", zc.TypeName(a)))
	}
	return v
}

func (k int64UKind) Dup(a any) any {
	return a
}

func (k int64UKind) Copy(src, dest any) {
	d, ok := dest.(*uint64)
	if !ok {
		panic(fmt.Errorf("expected *uint64 but got: %v", zc.TypeName(dest)))
	}
	s := k.As(src)
	*d = s
}

func (k int64UKind) To(a any) (any, bool) {
	return nil, false
}
