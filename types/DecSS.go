package types

import (
	"fmt"

	"github.com/blackchip-org/zc/v6/state"
	"github.com/shopspring/decimal"
)

var DecSS = decSSType{}

type decSSType struct{}

func (t decSSType) Name() string { return "Dec/ss" }

func (t decSSType) Is(a any) bool {
	switch a.(type) {
	case decimal.Decimal:
		return true
	}
	return false
}

func (t decSSType) As(a any) decimal.Decimal {
	v, ok := a.(decimal.Decimal)
	if !ok {
		panic(fmt.Errorf("expected decimal.Decimal but got %v", GoName(a)))
	}
	return v
}

func (t decSSType) Dup(a any) any {
	return a
}

func (t decSSType) Copy(src, dest any) {
	d, ok := dest.(*decimal.Decimal)
	if !ok {
		panic(fmt.Errorf("expected *decimal.Decimal but got %v", GoName(dest)))
	}
	*d = t.As(src)
}

func (t decSSType) To(state state.State, a any) (any, bool) {
	switch v := a.(type) {
	case string:
		d, err := decimal.NewFromString(v)
		if err != nil {
			return nil, false
		}
		return d, true
	}
	return nil, false
}

func (t decSSType) Format(a any) string {
	d, ok := a.(decimal.Decimal)
	if !ok {
		panic(fmt.Errorf("expected decimal.Decimal but got: %v", GoName(a)))
	}
	return d.String()
}
