package zc

import (
	"github.com/blackchip-org/zc/v5/pkg/types"
)

// --

type DMSType struct{}

func (t DMSType) String() string { return "DMS" }

func (t DMSType) Parse(str string) (types.DMS, bool) {
	return types.ParseDMS(str)
}

func (t DMSType) MustParse(s string) types.DMS {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t DMSType) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t DMSType) Format(v types.DMS) string {
	return v.String()
}

func PopDMS(c Calc) types.DMS     { return DMS.MustParse(c.MustPop()) }
func PushDMS(c Calc, r types.DMS) { c.Push(DMS.Format(r)) }

// --
