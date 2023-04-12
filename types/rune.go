package types

import (
	"fmt"
	"unicode/utf8"
)

type vRune struct {
	val rune
}

func (v vRune) Type() Type     { return RuneType{} }
func (v vRune) String() string { return RuneType{}.Format(v.val) }
func (v vRune) Native() any    { return v.val }

type RuneType struct{}

func (t RuneType) String() string { return "Rune" }

func (t RuneType) Parse(s string) (rune, error) {
	if utf8.RuneCountInString(s) != 1 {
		return 0, parseErr(t, s)
	}
	r, _ := utf8.DecodeRuneInString(s)
	return r, nil
}

func (t RuneType) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return nil, err
	}
	return t.Value(v), nil
}

func (t RuneType) Format(r rune) string {
	return fmt.Sprintf("%c", r)
}

func (t RuneType) Value(r rune) Value {
	return vRune{val: r}
}
