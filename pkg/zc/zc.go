package zc

import (
	"unicode"

	"github.com/blackchip-org/zc/pkg/scanner"
)

const ProgName = "zc"

type Calc interface {
	Eval(string) error
	Stack() []string
	StackLen() int
	SetStack([]string)
	Peek(int) (string, bool)
	Pop() (string, bool)
	MustPop() string
	Push(string)
	Info() string
	SetInfo(string, ...any)
	Error() error
	SetError(error)
	Derive() Calc
	NewState(string, any)
	State(string) (any, bool)
	Op() string
	Ops() []string
}

type CalcFunc func(Calc)

type OpDecl struct {
	Name  string
	Macro string
	Funcs []FuncDecl
}

type FuncDecl struct {
	Func   CalcFunc
	Params []Type
}

func Op(name string, fn CalcFunc, param ...Type) OpDecl {
	return OpDecl{
		Name: name,
		Funcs: []FuncDecl{
			{Func: fn, Params: param},
		},
	}
}

func GenOp(name string, funcs ...FuncDecl) OpDecl {
	return OpDecl{Name: name, Funcs: funcs}
}

func Macro(name string, expr string) OpDecl {
	return OpDecl{Name: name, Macro: expr}
}

func Func(fn CalcFunc, params ...Type) FuncDecl {
	return FuncDecl{Func: fn, Params: params}
}

func EvalOp(op OpDecl) CalcFunc {
	return func(c Calc) {
		if op.Macro != "" {
			c.Eval(op.Macro)
			return
		}
		if len(op.Funcs) == 0 {
			panic("no functions for operation")
		}
		for _, decl := range op.Funcs {
			if isOpMatch(c, decl) {
				decl.Func(c)
				return
			}
		}

		// For now, we are going to check the first decl to
		// determine the number of arguments.
		nArgs := len(op.Funcs[0].Params)
		var types []Type
		for i := 0; i < nArgs; i++ {
			v, ok := c.Peek(i)
			if !ok {
				ErrNotEnoughArgs(c, op.Name, nArgs)
				return
			}
			types = append(types, Identify(v))
		}
		ErrNoOpFor(c, op.Name, types...)
	}
}

func isOpMatch(c Calc, decl FuncDecl) bool {
	for i, param := range decl.Params {
		arg, ok := c.Peek(len(decl.Params) - i - 1)
		if !ok {
			return false
		}
		if !param.Is(arg) {
			return false
		}
	}
	return true
}

func NoOp(c Calc) {}

func IsValuePrefix(ch rune, next rune) bool {
	switch {
	case unicode.IsDigit(ch), unicode.Is(unicode.Sc, ch):
		return true
	case (ch == '-' || ch == '+' || ch == '.') && unicode.IsDigit(next):
		return true
	}
	return false
}

func Quote(v string) string {
	var s scanner.Scanner
	s.SetString(v)

	needsQuotes := false
	if !IsValuePrefix(s.Ch, s.Lookahead) {
		needsQuotes = true
	} else {
		s.ScanUntil(unicode.IsSpace)
		if !s.End() {
			needsQuotes = true
		}
	}

	if !needsQuotes {
		return v
	}

	s.SetString(v)
	s.Text.WriteRune('\'')
	for s.Ok() {
		if s.Ch == '\'' {
			s.Text.WriteString("\\'")
		} else {
			s.Keep()
		}
	}
	s.Text.WriteRune('\'')
	return s.Token()
}
