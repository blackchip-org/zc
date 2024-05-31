package types

import (
	"fmt"

	"github.com/shopspring/decimal"
)

var Dec = decType{}

type decType struct{}

func (t decType) Name() string { return "Dec" }

func (t decType) Is(a any) bool {
	switch a.(type) {
	case decimal.Decimal:
		return true
	}
	return false
}

func (t decType) As(a any) decimal.Decimal {
	v, ok := a.(decimal.Decimal)
	if !ok {
		panic(fmt.Errorf("expected decimal.Decimal but got %v", GoName(a)))
	}
	return v
}

func (t decType) Dup(a any) any {
	return a
}

func (t decType) Copy(src, dest any) {
	d, ok := dest.(*decimal.Decimal)
	if !ok {
		panic(fmt.Errorf("expected *decimal.Decimal but got %v", GoName(dest)))
	}
	*d = t.As(src)
}

func (t decType) To(a any) (any, bool) {
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
