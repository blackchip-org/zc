package zc

import (
	"fmt"
	"unicode/utf8"
)

// ---

type RuneType struct{}

func (t RuneType) String() string { return "Rune" }

func (t RuneType) Parse(s string) (rune, bool) {
	if utf8.RuneCountInString(s) != 1 {
		return 0, false
	}
	r, _ := utf8.DecodeRuneInString(s)
	return r, true
}

func (t RuneType) MustParse(s string) rune {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t RuneType) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t RuneType) Format(v rune) string {
	return fmt.Sprintf("%c", v)
}

func PopRune(c Calc) rune     { return Rune.MustParse(c.MustPop()) }
func PushRune(c Calc, r rune) { c.Push(Rune.Format(r)) }

// ---

type StringType struct{}

func (t StringType) String() string { return "String" }

func (t StringType) Parse(s string) (string, bool) {
	return s, true
}

func (t StringType) MustParse(s string) string {
	return s
}

func (t StringType) Is(s string) bool {
	return true
}

func (t StringType) Format(v string) string {
	return v
}

func PopString(c Calc) string     { return c.MustPop() }
func PushString(c Calc, r string) { c.Push(r) }
