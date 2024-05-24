package zc

import (
	"unicode"
	"unicode/utf8"
)

type Type interface {
	Name() string
	Is(any) bool
	Dup(any) any
	Copy(any, any)
	To(any) (any, bool)
}

type OpEnv struct {
	Catalog *Catalog
	Op      Op
	Args    []any
	Returns []any
	Err     error
}

type Op struct {
	Name      string
	Aliases   []string
	Params    []Type
	VarParam  Type
	Returns   []Type
	VarReturn Type
	Func      func(*OpEnv)
	Macro     string
}

type Vol struct {
	Name  string
	Types []Type
	Ops   []Op
}

func IsValuePrefix(ch rune, next rune) bool {
	switch {
	case unicode.IsDigit(ch):
		return true
	case (ch == '-' || ch == '+' || ch == '.') && unicode.IsDigit(next):
		return true
	}
	return false
}

func IsValue(s string) bool {
	var ch, next rune
	var w int
	n := len(s)
	if n > 0 {
		ch, w = utf8.DecodeRuneInString(s)
		if n > 1 {
			next, _ = utf8.DecodeRuneInString(s[w:])
		}
	}
	return IsValuePrefix(ch, next)
}
