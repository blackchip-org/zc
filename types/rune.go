package types

import (
	"unicode/utf8"
)

type RuneType struct{}

func (t RuneType) String() string { return "Rune" }

func (t RuneType) Parse(s string) (rune, error) {
	if utf8.RuneCountInString(s) != 1 {
		return 0, parseErr(t, s)
	}
	r, _ := utf8.DecodeRuneInString(s)
	return r, nil
}
