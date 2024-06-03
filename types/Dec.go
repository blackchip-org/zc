package types

import (
	"fmt"

	"github.com/blackchip-org/zc/v6/state"
	"github.com/cockroachdb/apd/v3"
)

var Dec = decType{}

type decType struct{}

func (t decType) Name() string { return "Dec" }

func (t decType) Is(a any) bool {
	switch a.(type) {
	case *apd.Decimal:
		return true
	}
	return false
}

func (t decType) As(a any) *apd.Decimal {
	v, ok := a.(*apd.Decimal)
	if !ok {
		panic(fmt.Errorf("expected *apd.Decimal but got %v", GoName(a)))
	}
	return v
}

func (t decType) Dup(a any) any {
	r := new(apd.Decimal)
	v := t.As(a)
	r.Set(v)
	return r
}

func (t decType) Copy(src, dest any) {
	d, ok := dest.(*apd.Decimal)
	if !ok {
		panic(fmt.Errorf("expected *apd.Decimal but got: %v", GoName(dest)))
	}
	s := t.As(src)
	d.Set(s)
}

func (t decType) To(state state.State, a any) (any, bool) {
	switch v := a.(type) {
	case string:
		d, _, err := apd.NewFromString(v)
		if err != nil {
			return nil, false
		}
		return d, true
	}
	return nil, false
}

func (t decType) Format(a any) string {
	d, ok := a.(*apd.Decimal)
	if !ok {
		panic(fmt.Errorf("expected *apd.Decimal but got: %v", GoName(a)))
	}
	return RemoveTrailingZeros(d.Text('g'))
}
