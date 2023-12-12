package zc

import (
	"cmp"
	"fmt"
	"unicode/utf8"
)

// ---

type CharType struct{}

func (t CharType) String() string { return "Char" }

func (t CharType) Parse(s string) (rune, bool) {
	if utf8.RuneCountInString(s) != 1 {
		return 0, false
	}
	r, _ := utf8.DecodeRuneInString(s)
	return r, true
}

func (t CharType) MustParse(s string) rune {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t CharType) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t CharType) Format(v rune) string {
	return fmt.Sprintf("%c", v)
}

func PopRune(c Calc) rune     { return Char.MustParse(c.MustPop()) }
func PushRune(c Calc, r rune) { c.Push(Char.Format(r)) }

// ---

type StrType struct{}

func (t StrType) String() string { return "Str" }

func (t StrType) Parse(s string) (string, bool) {
	return s, true
}

func (t StrType) MustParse(s string) string {
	return s
}

func (t StrType) Is(s string) bool {
	return true
}

func (t StrType) Compare(x1 string, x2 string) (int, bool) {
	s1, ok := t.Parse(x1)
	if !ok {
		return 0, false
	}
	s2, ok := t.Parse(x2)
	if !ok {
		return 0, false
	}
	return cmp.Compare(s1, s2), true
}

func (t StrType) Format(v string) string {
	return v
}

func PopString(c Calc) string     { return c.MustPop() }
func PushString(c Calc, r string) { c.Push(r) }
