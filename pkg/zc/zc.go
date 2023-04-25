package zc

import (
	"strings"
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
	SetOp(string)
	SetArgs([]string)
	Op() OpCall
	OpNames() []string
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

type OpCall struct {
	Name string
	Args []string
}

func (c OpCall) String() string {
	words := append(c.Args, c.Name)
	return strings.Join(words, " ")
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

func NoOp(c Calc) {}
