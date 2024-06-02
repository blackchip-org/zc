package types

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

var (
	Int     = intType{}
	IntArch = intArchType{}
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

func (t intType) To(a any) (any, bool) {
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

type intArchType struct{}

func (t intArchType) Name() string { return "IntArch" }

func (t intArchType) Is(a any) bool {
	switch a.(type) {
	case int, *int:
		return true
	}
	return false
}

func (t intArchType) As(a any) int {
	v, ok := a.(int)
	if !ok {
		panic(fmt.Errorf("expected int but got: %v", GoName(a)))
	}
	return v
}

func (t intArchType) Dup(a any) any {
	return a
}

func (t intArchType) Copy(src, dest any) {
	d, ok := dest.(*int)
	if !ok {
		panic(fmt.Errorf("expected *int but got: %v", GoName(dest)))
	}
	s := t.As(src)
	*d = s
}

func (t intArchType) To(a any) (any, bool) {
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

func (t intArchType) Format(a any) string {
	d, ok := a.(int)
	if !ok {
		panic(fmt.Errorf("expected int but got: %v", GoName(a)))
	}
	return strconv.Itoa(d)
}
