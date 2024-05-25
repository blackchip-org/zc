package types

import (
	"fmt"
	"math/big"
)

var Float = floatType{}

type floatType struct{}

func (t floatType) Name() string { return "Float" }

func (t floatType) Is(a any) bool {
	switch a.(type) {
	case *big.Float:
		return true
	}
	return false
}

func (t floatType) As(a any) *big.Float {
	v, ok := a.(*big.Float)
	if !ok {
		panic(fmt.Errorf("expected *big.Float but got %v", GoName(a)))
	}
	return v
}

func (t floatType) Dup(a any) any {
	var r big.Float
	v := t.As(a)
	r.Set(v)
	return &r
}

func (t floatType) Copy(src, dest any) {
	d, ok := dest.(*big.Float)
	if !ok {
		panic(fmt.Errorf("expected *big.Float but got %v", GoName(dest)))
	}
	s := t.As(src)
	d.Set(s)
}

func (t floatType) To(a any) (any, bool) {
	switch v := a.(type) {
	case float64:
		return big.NewFloat(v), true
	case string:
		var f big.Float
		_, ok := f.SetString(v)
		return &f, ok
	}
	return nil, false
}
