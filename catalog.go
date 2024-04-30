package zc

import (
	"maps"
	"slices"

	"github.com/blackchip-org/zc/v6/errors"
)

type CatalogBuilder struct {
	kinds    map[string]Kind
	ordKinds []Kind
	ops      map[string]Op
	ordOps   []Op
}

func NewCatalogBuilder() *CatalogBuilder {
	return &CatalogBuilder{
		kinds: make(map[string]Kind),
		ops:   make(map[string]Op),
	}
}

func (c *CatalogBuilder) AddKind(ks ...Kind) {
	for _, k := range ks {
		name := k.Name()
		if _, ok := c.kinds[name]; ok {
			panic(errors.DuplicateKind(name))
		}
		c.kinds[name] = k
		c.ordKinds = append(c.ordKinds, k)
	}
}

func (c *CatalogBuilder) addOp(name string, op Op) {
	if _, ok := c.ops[name]; ok {
		panic(errors.DuplicateOp(name))
	}
	if op.Func == nil {
		panic(errors.NoFuncForOp(name))
	}
	c.ops[name] = op
}

func (c *CatalogBuilder) AddOp(ops ...Op) {
	for _, op := range ops {
		c.addOp(op.Name, op)
		for _, alias := range op.Aliases {
			c.addOp(alias, op)
		}
		c.ordOps = append(c.ordOps, op)
	}
}

func (c *CatalogBuilder) AddVolume(vols ...Volume) {
	for _, vol := range vols {
		c.AddKind(vol.Kinds...)
		c.AddOp(vol.Ops...)
	}
}

func (c *CatalogBuilder) Build() *Catalog {
	cat := &Catalog{
		kinds:    maps.Clone(c.kinds),
		ordKinds: slices.Clone(c.ordKinds),
		ops:      maps.Clone(c.ops),
		ordOps:   slices.Clone(c.ordOps),
	}
	val := valKind{}
	cat.kinds[val.Name()] = val
	return cat
}

type Catalog struct {
	kinds    map[string]Kind
	ordKinds []Kind
	ops      map[string]Op
	ordOps   []Op
}

func (c *Catalog) KindByName(name string) (Kind, bool) {
	k, ok := c.kinds[name]
	return k, ok
}

func (c *Catalog) KindByType(v any) (Kind, bool) {
	for _, t := range c.ordKinds {
		if t.Is(v) {
			return t, true
		}
	}
	return nil, false
}

func (c *Catalog) OpByName(name string) (Op, bool) {
	op, ok := c.ops[name]
	return op, ok
}

func (c *Catalog) Ops() []Op {
	return slices.Clone(c.ordOps)
}

func (c *Catalog) Copy(v any) any {
	k, ok := c.KindByType(v)
	if !ok {
		panic(errors.UnregisteredType(v))
	}
	return k.Dup(v)
}

type valKind struct{}

func (k valKind) Name() string { return "Val" }

func (k valKind) Is(v any) bool {
	if v == nil {
		return false
	}
	return true
}

func (k valKind) Dup(any) any {
	panic("Dup undefined")
}

func (k valKind) Copy(any, any) {
	panic("Copy undefined")
}

func (k valKind) To(a any) (any, bool) {
	return a, false
}
