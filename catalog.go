package zc

import (
	"fmt"
	"slices"
)

type Catalog struct {
	ops map[string]Op
}

func NewCatalog() *Catalog {
	return &Catalog{ops: make(map[string]Op)}
}

func (c *Catalog) AddOp(ops ...Op) {
	for _, op := range ops {
		if len(op.Funcs) == 0 {
			panic(fmt.Errorf("no funcs for op: %v", op.Name))
		}
		if prevOp, ok := c.ops[op.Name]; ok {
			prevOp.Funcs = append(prevOp.Funcs, op.Funcs...)
			c.ops[prevOp.Name] = prevOp
		} else {
			c.ops[op.Name] = op
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
	c.ops[name] = op
}

func (c *Catalog) AddVol(vols ...Vol) {
	for _, vol := range vols {
		c.AddOp(vol.Ops...)
		for _, mac := range vol.Macros {
			c.AddMacro(mac.Name, mac.Expr)
		}
	}
}

func (c *Catalog) OpFor(name string) (Op, bool) {
	op, ok := c.ops[name]
	return op, ok
}

func (c *Catalog) OpNames() []string {
	var names []string
	for name := range c.ops {
		names = append(names, name)
	}
	slices.Sort(names)
	return names
}
