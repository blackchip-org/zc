package zc

import (
	"fmt"
	"unicode/utf8"
)

// ---

type RuneType struct{}

func (t RuneType) String() string { return "Rune" }

func (t RuneType) Parse(s string) (rune, error) {
	if utf8.RuneCountInString(s) != 1 {
		return 0, ErrExpectedType(t, s)
	}
	r, _ := utf8.DecodeRuneInString(s)
	return r, nil
}

func (t RuneType) MustParse(s string) rune {
	r, err := t.Parse(s)
	if err != nil {
		panic(err)
	}
	return r
}

func (t RuneType) Is(s string) bool {
	_, err := t.Parse(s)
	return err == nil
}

func (t RuneType) Format(v rune) string {
	return fmt.Sprintf("%c", v)
}

// ---

type StringType struct{}

func (t StringType) String() string { return "String" }

func (t StringType) Parse(s string) (string, error) {
	return s, nil
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
