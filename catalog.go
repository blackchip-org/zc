package zc

import "github.com/blackchip-org/zc/v6/errors"

type Catalog struct {
	kinds    map[string]Kind
	ordKinds []Kind
	ops      map[string]Op
}

func NewCatalog() *Catalog {
	return &Catalog{
		kinds: make(map[string]Kind),
		ops:   make(map[string]Op),
	}
}

func (c *Catalog) AddKind(ks ...Kind) {
	for _, k := range ks {
		name := k.Name()
		if _, ok := c.kinds[name]; ok {
			panic(errors.DuplicateKind(name))
		}
		c.kinds[name] = k
		c.ordKinds = append(c.ordKinds, k)
	}
}

func (c *Catalog) addOp(name string, op Op) {
	if _, ok := c.ops[name]; ok {
		panic(errors.DuplicateOp(name))
	}
	if op.Func == nil {
		panic(errors.NoFuncForOp(name))
	}
	c.ops[name] = op
}

func (c *Catalog) AddOp(ops ...Op) {
	for _, op := range ops {
		c.addOp(op.Name, op)
		for _, alias := range op.Aliases {
			c.addOp(alias, op)
		}
	}
}

func (c *Catalog) AddVolume(vols ...Volume) {
	for _, vol := range vols {
		c.AddKind(vol.Kinds...)
		c.AddOp(vol.Ops...)
	}
}

func (c *Catalog) KindFor(name string) (Kind, bool) {
	k, ok := c.kinds[name]
	return k, ok
}

func (c *Catalog) KindOf(v any) (Kind, bool) {
	for _, t := range c.ordKinds {
		if t.Is(v) {
			return t, true
		}
	}
	return nil, false
}

func (c *Catalog) OpFor(name string) (Op, bool) {
	op, ok := c.ops[name]
	return op, ok
}

func (c *Catalog) Copy(v any) any {
	k, ok := c.KindOf(v)
	if !ok {
		panic(errors.NewUnregisteredType(v))
	}
	return k.Dup(v)
}
