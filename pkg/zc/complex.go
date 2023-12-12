package zc

import (
	"strconv"
	"strings"
)

type ComplexType struct{}

func (t ComplexType) String() string { return "Complex" }

func (t ComplexType) Parse(s string) (complex128, bool) {
	if !strings.HasSuffix(s, "i") {
		return 0, false
	}
	c, err := strconv.ParseComplex(s, 128)
	if err != nil {
		return 0, false
	}
	return c, true
}

func (t ComplexType) MustParse(s string) complex128 {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t ComplexType) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t ComplexType) Format(v complex128) string {
	s := strconv.FormatComplex(v, 'g', 16, 128)
	// For some reason, the complex number is surrounded by parens.
	// Remove them.
	return s[1 : len(s)-1]
}

func PopComplex(c Calc) complex128     { return Complex.MustParse(c.MustPop()) }
func PushComplex(c Calc, r complex128) { c.Push(Complex.Format(r)) }
