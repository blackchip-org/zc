package types

import (
	"fmt"
	"math"
	"math/big"
	"strconv"

	"github.com/blackchip-org/zc/v6/state"
)

var (
	Int   = intType{}
	IntA  = intAType{}
	Int32 = int32Type{}
	IntU  = intUType{}
)

type intType struct{}

func (t intType) Name() string { return "Int" }

func (t intType) Is(a any) bool {
	switch a.(type) {
	case *big.Int:
		return true
	}
	return false
}

func (t intType) As(a any) *big.Int {
	v, ok := a.(*big.Int)
	if !ok {
		panic(fmt.Errorf("expected *big.Int but got: %v", GoName(a)))
	}
	return v
}

func (t intType) Dup(a any) any {
	r := new(big.Int)
	v := t.As(a)
	r.Set(v)
	return r
}

func (t intType) Copy(src, dest any) {
	d, ok := dest.(*big.Int)
	if !ok {
		panic(fmt.Errorf("expected *big.Int but got: %v", GoName(dest)))
	}
	s := t.As(src)
	d.Set(s)
}

func (t intType) To(state state.State, a any) (any, bool) {
	switch v := a.(type) {
	case int:
		return big.NewInt(int64(v)), true
	case string:
		var i big.Int
		_, ok := i.SetString(v, 0)
		return &i, ok
	}
	return nil, false
}

func (t intType) Format(a any) string {
	d, ok := a.(*big.Int)
	if !ok {
		panic(fmt.Errorf("expected *big.Int but got: %v", GoName(a)))
	}
	return d.String()
}

// ----------------------------------------------------------------------------

type intAType struct{}

func (t intAType) Name() string { return "Int/a" }

func (t intAType) Is(a any) bool {
	switch a.(type) {
	case int, *int:
		return true
	}
	return false
}

func (t intAType) As(a any) int {
	v, ok := a.(int)
	if !ok {
		panic(fmt.Errorf("expected int but got: %v", GoName(a)))
	}
	return v
}

func (t intAType) Dup(a any) any {
	return a
}

func (t intAType) Copy(src, dest any) {
	d, ok := dest.(*int)
	if !ok {
		panic(fmt.Errorf("expected *int but got: %v", GoName(dest)))
	}
	s := t.As(src)
	*d = s
}

func (t intAType) To(state state.State, a any) (any, bool) {
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
		if v > math.MaxInt {
			return nil, false
		}
		return int(v), true
	case uint64:
		if v > math.MaxInt {
			return 0, false
		}
		return int(v), true
	case uint32:
		return int(v), true
	case uint16:
		return int(v), true
	case uint8:
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

func (t intAType) Format(a any) string {
	d, ok := a.(int)
	if !ok {
		panic(fmt.Errorf("expected int but got: %v", GoName(a)))
	}
	return strconv.Itoa(d)
}

// ----------------------------------------------------------------------------

type int32Type struct{}

func (t int32Type) Name() string { return "Int/32" }

func (t int32Type) Is(a any) bool {
	switch a.(type) {
	case int32, *int32:
		return true
	}
	return false
}

func (t int32Type) As(a any) int32 {
	v, ok := a.(int32)
	if !ok {
		panic(fmt.Errorf("expected int32 but got: %v", GoName(a)))
	}
	return v
}

func (t int32Type) Dup(a any) any {
	return a
}

func (t int32Type) Copy(src, dest any) {
	d, ok := dest.(*int32)
	if !ok {
		panic(fmt.Errorf("expected *int32 but got: %v", GoName(dest)))
	}
	s := t.As(src)
	*d = s
}

func (t int32Type) To(state state.State, a any) (any, bool) {
	switch v := a.(type) {
	// case int64:
	// 	if v > math.MaxInt || v < math.MinInt {
	// 		return nil, false
	// 	}
	// 	return int(v), true
	// case int32:
	// 	return int(v), true
	// case int16:
	// 	return int(v), true
	// case int8:
	// 	return int(v), true
	// case uint:
	// 	if v > math.MaxInt {
	// 		return nil, false
	// 	}
	// 	return int(v), true
	// case uint64:
	// 	if v > math.MaxInt {
	// 		return 0, false
	// 	}
	// 	return int(v), true
	// case uint32:
	// 	return int(v), true
	// case uint16:
	// 	return int(v), true
	// case uint8:
	// 	return int(v), true
	case string:
		i, err := strconv.ParseInt(v, 0, 32)
		if err != nil {
			return 0, false
		}
		return int32(i), true
	}
	return nil, false
}

func (t int32Type) Format(a any) string {
	d, ok := a.(int32)
	if !ok {
		panic(fmt.Errorf("expected int32 but got: %v", GoName(a)))
	}
	return strconv.FormatInt(int64(d), 10)
}

// ----------------------------------------------------------------------------

type intUType struct{}

func (t intUType) Name() string { return "Int/u" }

func (t intUType) Is(a any) bool {
	switch a.(type) {
	case uint, *uint:
		return true
	}
	return false
}

func (t intUType) As(a any) uint {
	v, ok := a.(uint)
	if !ok {
		panic(fmt.Errorf("expected uint but got: %v", GoName(a)))
	}
	return v
}

func (t intUType) Dup(a any) any {
	return a
}

func (t intUType) Copy(src, dest any) {
	d, ok := dest.(*uint)
	if !ok {
		panic(fmt.Errorf("expected *uint but got: %v", GoName(dest)))
	}
	s := t.As(src)
	*d = s
}

func (t intUType) To(state state.State, a any) (any, bool) {
	switch v := a.(type) {
	// case int64:
	// 	if v > math.MaxInt || v < math.MinInt {
	// 		return nil, false
	// 	}
	// 	return int(v), true
	// case int32:
	// 	return int(v), true
	// case int16:
	// 	return int(v), true
	// case int8:
	// 	return int(v), true
	// case uint:
	// 	if v > math.MaxInt {
	// 		return nil, false
	// 	}
	// 	return int(v), true
	// case uint64:
	// 	if v > math.MaxInt {
	// 		return 0, false
	// 	}
	// 	return int(v), true
	// case uint32:
	// 	return int(v), true
	// case uint16:
	// 	return int(v), true
	// case uint8:
	// 	return int(v), true
	case string:
		i, err := strconv.ParseUint(v, 0, 0)
		if err != nil {
			return 0, false
		}
		return uint(i), true
	}
	return nil, false
}

func (t intUType) Format(a any) string {
	d, ok := a.(uint)
	if !ok {
		panic(fmt.Errorf("expected int but got: %v", GoName(a)))
	}
	return strconv.FormatUint(uint64(d), 10)
}
