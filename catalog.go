package zc

import (
	"cmp"
	"fmt"
	"maps"
	"reflect"
	"slices"

	"github.com/blackchip-org/zc/v6/types"
)

type CatalogBuilder struct {
	types map[string]Type
	ops   map[string][]Op
}

func NewCatalogBuilder() *CatalogBuilder {
	return &CatalogBuilder{
		types: make(map[string]Type),
		ops:   make(map[string][]Op),
	}
}

func (c *CatalogBuilder) AddType(ts ...Type) {
	for _, t := range ts {
		name := t.Name()
		if _, ok := c.types[name]; ok {
			panic(fmt.Errorf("duplicate type: %v", name))
		}
		c.types[name] = t
	}
}

func (c *CatalogBuilder) addOp(name string, op Op) {
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
				panic(fmt.Errorf("duplicate op: %v", FormatList(stack)))
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

func (c *CatalogBuilder) AddOp(ops ...Op) {
	for _, op := range ops {
		if op.Func != nil {
			c.addOp(op.Name, op)
		}
		for _, alias := range op.Aliases {
			//c.AddMacro(alias, op.Name)
			c.addOp(alias, op)
		}
	}
}

func (c *CatalogBuilder) AddMacro(name string, mac string) {
	if _, ok := c.ops[name]; ok {
		panic(fmt.Errorf("duplicate op: %v", name))
	}
	op := Op{
		Name:  name,
		Macro: mac,
	}
	c.ops[name] = []Op{op}
}

func (c *CatalogBuilder) AddVolume(vols ...Vol) {
	for _, vol := range vols {
		c.AddType(vol.Types...)
		c.AddOp(vol.Ops...)
	}
}

func (c *CatalogBuilder) Build() *Catalog {
	cat := &Catalog{
		types: maps.Clone(c.types),
		ops:   make(map[string][]Op),
	}
	for k, v := range c.ops {
		cat.ops[k] = slices.Clone(v)
	}
	//val := types.Val
	//cat.types[val.Name()] = val
	return cat
}

// ----------------------------------------------------------------------------

type Catalog struct {
	types map[string]Type
	ops   map[string][]Op
}

func (c *Catalog) TypeOf(v any) (Type, bool) {
	for _, t := range c.types {
		if t.Is(v) {
			return t, true
		}
	}
	return nil, false
}

func (c *Catalog) OpFor(name string) ([]Op, bool) {
	op, ok := c.ops[name]
	return op, ok
}

func (c *Catalog) Ops() []Op {
	var ops []Op
	for _, v := range c.ops {
		ops = append(ops, v...)
	}
	slices.SortStableFunc(ops, func(a, b Op) int {
		return cmp.Compare(a.Name, b.Name)
	})
	return ops
}

func (c *Catalog) OpNames() []string {
	var names []string
	for _, ops := range c.ops {
		for _, op := range ops {
			names = append(names, op.Name)
			names = append(names, op.Aliases...)
		}
	}
	slices.Sort(names)
	return slices.Compact(names)
}

func (c *Catalog) Dup(v any) any {
	t, ok := c.TypeOf(v)
	if !ok || t == types.Val {
		panic(fmt.Errorf("unregistered type: %v", types.GoName(v)))
	}
	return t.Dup(v)
}
