package zc

import (
	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v6/app/state"
)

const ProgName = "zc"

type Item struct {
	Val   any
	Repr  string
	Unit  string
	Label string
}

func (i Item) String() string {
	var label, val string
	if i.Label != "" {
		label = i.Label + ": "
	}
	if i.Repr != "" {
		val = i.Repr
	} else {
		val = Format(i.Val)
	}
	return label + val + i.Unit
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

func (e *OpEnv) Label(l string) {
	item := e.Stack.Pop()
	item.Label = l
	e.Stack.Push(item)
}

func (e *OpEnv) Unit(u string) {
	item := e.Stack.Pop()
	item.Unit = u
	e.Stack.Push(item)
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
