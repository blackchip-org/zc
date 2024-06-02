package zc

import (
	"unicode"
	"unicode/utf8"
)

const ProgName = "zc"

type Type interface {
	Name() string
	Is(any) bool
	Dup(any) any
	Copy(any, any)
	To(any) (any, bool)
	Format(any) string
}

type OpEnv struct {
	Catalog *Catalog
	State   map[string]any
	Op      Op
	Args    []any
	Returns []any
	Err     error
}

type Op struct {
	Name      string
	Overloads string
	Virtual   bool
	Params    []Type
	VarParam  Type
	Returns   []Type
	VarReturn Type
	Macro     string
	Func      func(*OpEnv)
}

type Macro struct {
	Name string
	Expr string
}

type Vol struct {
	Name   string
	Types  []Type
	Ops    []Op
	Macros []Macro
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
