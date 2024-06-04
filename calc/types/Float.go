package types

import (
	"fmt"
	"math/big"

	"github.com/blackchip-org/zc/v6/calc/state"
)

var (
	Float     = floatType{}
	FloatZero = big.NewFloat(0)
)

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

func (t floatType) To(states state.State, a any) (any, bool) {
	switch v := a.(type) {
	case float64:
		return big.NewFloat(v), true
	case string:
		s := state.ForConf(states)
		f, _, err := big.ParseFloat(v, 0, s.FloatPrec, s.RoundingMode)
		return f, err == nil
	}
	return nil, false
}

func (t floatType) Format(a any) string {
	d, ok := a.(*big.Float)
	if !ok {
		panic(fmt.Errorf("expected *big.Float but got: %v", GoName(a)))
	}
	return fmt.Sprintf("%v", d)
}
