package zc

import (
	"github.com/blackchip-org/dms"
)

// --

var (
	dmsParser    = dms.NewDefaultParser()
	dmsFormatter = dms.NewFormatter(dms.SecType, -1)
)

type DMSType struct{}

func (t DMSType) String() string { return "DMS" }

func (t DMSType) Parse(str string) (dms.Angle, bool) {
	a, err := dmsParser.Parse(str)
	if err != nil {
		return dms.Angle{}, false
	}
	return a, true
}

func (t DMSType) MustParse(s string) dms.Angle {
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

func (t DMSType) Format(v dms.Angle) string {
	return dmsFormatter.Format(v)
}

func PopDMS(c Calc) dms.Angle     { return DMS.MustParse(c.MustPop()) }
func PushDMS(c Calc, r dms.Angle) { c.Push(DMS.Format(r)) }

// --
