package zc

import (
	"unicode"

	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v6/app/state"
)

const ProgName = "zc"

type Item struct {
	Val   any
	Unit  string
	Label string
}

func (i Item) String() string {
	return Format(i.Val)
}

type OpEnv struct {
	*Stack[Item]
	Op    Op
	State state.State
	Err   error
	Info  string
}

func (e *OpEnv) PushVal(vals ...any) {
	for _, val := range vals {
		e.Push(Item{Val: val})
	}
}

type Op struct {
	Name      string
	Overloads string
	Virtual   bool
	Params    []Type
	VarParam  Type
	Returns   []Type
	VarReturn Type
	Func      func(*OpEnv)
	Macro     []scan.Token
}

type Macro struct {
	Name string
	Expr string
}

type Vol struct {
	Name   string
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
