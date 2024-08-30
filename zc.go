package zc

import (
	"strings"

	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v6/pkg/coll"
)

const ProgName = "zc"

type Type interface {
	AppName() string
	GoName() string
	Parse(coll.State, string) (any, bool, error)
	Format(any) string
	Dup(any) any
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

func (i Item) Dup() Item {
	return Item{
		TypeVal: i.Type.Dup(i.TypeVal),
		Type:    i.Type,
		Unit:    i.Unit,
		Label:   i.Label,
	}
}

func (i Item) String() string {
	var s strings.Builder
	s.WriteString(EscapeString(i.Val()))
	if i.Unit != "" {
		s.WriteRune(' ')
		s.WriteString(i.Unit)
	}
	if i.Label != "" {
		s.WriteString(" :")
		s.WriteString(i.Label)
	}
	return s.String()
}

func DupItems(items []Item) []Item {
	items2 := make([]Item, len(items))
	for i, item := range items {
		items2[i] = item.Dup()
	}
	return items2
}

type State interface {
	State(string) (any, bool)
	NewState(string, any)
}

type Calc interface {
	Push(...Item)
	Pop() Item
	Stack() []Item
	SetStack([]Item)
	Len() int
	String() string
	Var(string) (any, bool)
	NewVar(string, any)
	Notify(string, ...any)
	Raise(error)
	Label() string
	Unit() string
	SetLabel(string)
	SetUnit(string)
	Eval(string) error
	New() Calc
	Temp() []Item
	SetTemp([]Item)
	Store(string)
	Load(string)
	Reset()
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
