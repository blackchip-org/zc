package types

import (
	"fmt"
	"math/big"
)

var Int = intType{}

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
