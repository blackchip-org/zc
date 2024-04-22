package zc

import "github.com/blackchip-org/zc/v6/errors"

type Catalog struct {
	kinds []Kind
}

func NewCatalog() *Catalog {
	return &Catalog{}
}

func (c *Catalog) AddKind(ks ...Kind) {
	c.kinds = append(c.kinds, ks...)
}

func (c *Catalog) KindOf(v any) (Kind, bool) {
	for _, t := range c.kinds {
		if t.Is(v) {
			return t, true
		}
	}
	return nil, false
}

func (c *Catalog) Copy(v any) any {
	k, ok := c.KindOf(v)
	if !ok {
		panic(errors.NewUnregisteredType(v))
	}
	return k.Dup(v)
}
