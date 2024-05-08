package zc

import (
	"unicode"
	"unicode/utf8"
)

const (
	ProgName   = "zc"
	AnnoMarker = "#"
)

type Kind interface {
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
	Overloads string
	Aliases   []string
	Params    []string
	VarParam  string
	Returns   []string
	VarReturn string
	Prec      int
	Func      func(*OpEnv)
	Macro     string
}

type Vol struct {
	Name  string
	Kinds []Kind
	Ops   []Op
}

type Item struct {
	Value any
	Kind  Kind
	Anno  string
}

func (i Item) String() string {
	return String(i.Value)
}

func (i Item) StringWithAnno() string {
	if i.Anno == "" {
		return String(i.Value)
	}
	return String(i.Value) + " # " + i.Anno
}

func ItemValues(items []Item) []any {
	var vals []any
	for _, item := range items {
		vals = append(vals, item.Value)
	}
	return vals
}

func ItemStrings(items []Item) []string {
	var strs []string
	for _, item := range items {
		strs = append(strs, String(item.Value))
	}
	return strs
}

func IsValuePrefix(ch rune, next rune) bool {
	switch {
	case unicode.IsDigit(ch), unicode.Is(unicode.Sc, ch):
		return true
	case (ch == '-' || ch == '+' || ch == '.') && unicode.IsDigit(next):
		return true
	case ch == '/':
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
