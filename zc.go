package zc

import (
	"fmt"
	"reflect"
	"unicode"

	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v6/app/state"
)

const ProgName = "zc"

type Item struct {
	Val any
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

type Catalog struct {
	ops map[string][]Op
}

func NewCatalog() *Catalog {
	return &Catalog{ops: make(map[string][]Op)}
}

func (c *Catalog) addOp(name string, op Op) {
	var ops []Op
	var ok bool

	if ops, ok = c.ops[name]; ok {
		// Make sure that there isn't already an operation with the same
		// parameter signature
		for _, other := range ops {
			if reflect.DeepEqual(other.Params, op.Params) {
				var params []string
				for _, param := range op.Params {
					params = append(params, param.Name())
				}
				stack := append(params, name)
				panic(fmt.Errorf("duplicate op: %v", FormatList(
					stack)))
			}
			ops = append(ops, op)
			c.ops[name] = ops
		}
	} else {
		// No operation with this name has been defined yet.
		ops = []Op{op}
	}

	if op.Func == nil {
		panic(fmt.Errorf("no function for op: %v", op.Name))
	}
	c.ops[name] = ops
}

func (c *Catalog) AddOp(ops ...Op) {
	for _, op := range ops {
		if op.Func != nil {
			c.addOp(op.Name, op)
		}
		if op.Overloads != "" {
			c.addOp(op.Overloads, op)
		}
	}
}

func (c *Catalog) AddMacro(name string, mac string) {
	if _, ok := c.ops[name]; ok {
		panic(fmt.Errorf("duplicate op: %v", name))
	}
	toks := ScanWords(mac)
	op := Op{
		Name:  name,
		Macro: toks,
	}
	c.ops[name] = []Op{op}
}

func (c *Catalog) AddVolume(vols ...Vol) {
	for _, vol := range vols {
		c.AddOp(vol.Ops...)
		for _, mac := range vol.Macros {
			c.AddMacro(mac.Name, mac.Expr)
		}
	}
}

func (c *Catalog) OpFor(name string) ([]Op, bool) {
	ops, ok := c.ops[name]
	return ops, ok
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
