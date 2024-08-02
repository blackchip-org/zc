package zc

import (
	"strings"

	"github.com/blackchip-org/scan"
)

type Type interface {
	AppName() string
	GoName() string
	Parse(string) (any, bool)
	Format(any) string
}

type Item struct {
	TypeVal any
	Type    Type
	val     string
	Unit    string
	Label   string
}

func (i Item) Val() string {
	var str string
	if i.val == "" {
		str = i.Type.Format(i.TypeVal)
	} else {
		str = i.val
	}
	return str
}

func (i Item) String() string {
	var s strings.Builder
	s.WriteString(EscapeString(i.Val()))
	s.WriteString(i.Unit)
	if i.Label != "" {
		s.WriteString(" :")
		s.WriteString(i.Label)
	}
	return s.String()
}

type Calc interface {
	Push(Item)
	Pop() Item
	Items() []Item
	SetItems([]Item)
	Len() int
	String() string
	State(string) (any, bool)
	NewState(string, any)
	Notify(string)
	Raise(error)
}

type Op struct {
	Name  string
	Funcs []Func
	Macro []scan.Token
}

type Func struct {
	Params    []Type
	VarParam  Type
	Returns   []Type
	VarReturn Type
	Eval      func(Calc)
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
