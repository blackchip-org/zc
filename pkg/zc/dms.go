package zc

import (
	"github.com/blackchip-org/dms"
)

// --

var (
	dmsParser = dms.NewDefaultParser()
)

type DMSType struct{}

func (t DMSType) String() string { return "DMS" }

func (t DMSType) Parse(str string) (dms.Fields, bool) {
	f, err := dmsParser.ParseFields(str)
	if err != nil {
		return f, false
	}
	return f, true
}

func (t DMSType) MustParse(s string) dms.Fields {
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

func (t DMSType) Format(v dms.Fields) string {
	return v.String()
}

func PopDMS(c Calc) dms.Fields     { return DMS.MustParse(c.MustPop()) }
func PushDMS(c Calc, r dms.Fields) { c.Push(DMS.Format(r)) }

// --
