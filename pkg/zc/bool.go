package zc

import "strings"

type BoolType struct{}

func (t BoolType) String() string { return "Bool" }

func (t BoolType) Parse(s string) (bool, bool) {
	ls := strings.TrimSpace(strings.ToLower(s))
	switch ls {
	case "true":
		return true, true
	case "false":
		return false, true
	}
	return false, false
}

func (t BoolType) MustParse(s string) bool {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t BoolType) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t BoolType) Format(v bool) string {
	if v {
		return "true"
	}
	return "false"
}

func PopBool(c Calc) bool     { return Bool.MustParse(c.MustPop()) }
func PushBool(c Calc, r bool) { c.Push(Bool.Format(r)) }
