package zc

import (
	"fmt"
	"reflect"
	"slices"
)

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

func (c *Catalog) OpNames() []string {
	var names []string
	for name := range c.ops {
		names = append(names, name)
	}
	slices.Sort(names)
	return names
}
