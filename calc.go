package zc

import (
	"reflect"
	"slices"

	"github.com/blackchip-org/zc/v6/errors"
)

type Kind interface {
	String() string
	Is(any) bool
	Dup(any) any
	Copy(any, any)
	To(any) (any, bool)
}

type OpContext struct {
	Catalog *Catalog
	Args    []any
	Returns []any
	Err     error
}

type OpDecl struct {
	Name       string
	Params     []Kind
	VarParams  bool
	Returns    []Kind
	VarReturns bool
	Func       func(*OpContext)
}

type meta struct {
	Kind Kind
	Anno string
}

type Calc struct {
	Err  error
	cat  *Catalog
	vals []any
	meta []meta
}

func NewCalc(cat *Catalog) *Calc {
	c := &Calc{
		cat: cat,
	}
	return c
}

func (c *Calc) push(a any, k Kind) {
	c.vals = append(c.vals, a)
	c.meta = append(c.meta, meta{Kind: k})
}

func (c *Calc) pop() (v any, m meta) {
	top := len(c.vals) - 1
	v, c.vals = c.vals[top], c.vals[:top]
	m, c.meta = c.meta[top], c.meta[:top]
	return
}

func (c *Calc) Push(a any) {
	if c.Err != nil {
		return
	}
	if a == nil {
		panic("attempt to push a nil value")
	}
	k, ok := c.cat.KindOf(a)
	if !ok {
		panic(errors.NewUnregisteredType(a))
	}
	c.push(a, k)
}

func (c *Calc) Pop(dest any) error {
	if c.Err != nil {
		panic(c.Err)
	}
	if len(c.vals) == 0 {
		return errors.StackEmpty
	}

	destKind, ok := c.cat.KindOf(dest)
	if !ok {
		panic(errors.NewUnregisteredType(dest))
	}
	v, m := c.pop()
	if !destKind.Is(v) {
		cv, ok := destKind.To(v)
		if !ok {
			vt := reflect.TypeOf(v).Name()
			return errors.NewInvalidConversion(m.Kind.String(), vt, v)
		}
		v = cv
	}
	destKind.Copy(v, dest)
	return nil
}

func (c *Calc) PopString() string {
	var r string
	if err := c.Pop(&r); err != nil {
		panic(err)
	}
	return r
}

func (c *Calc) Stack() []any {
	return slices.Clone(c.vals)
}

func (c *Calc) Do(op OpDecl) {
	if c.Err != nil {
		return
	}
	stackLen := len(c.vals)
	if len(op.Params) > stackLen {
		c.Err = errors.StackEmpty
		return
	}

	var ctx OpContext
	ctx.Catalog = c.cat
	argStart := stackLen - len(op.Params)
	if op.VarParams {
		argStart = 0
	}
	for i, k := range op.Params {
		c.assembleArg(k, &ctx, argStart+i)
		if ctx.Err != nil {
			return
		}
	}
	if op.VarParams {
		for i := len(op.Params); i < len(c.vals); i++ {
			k := op.Params[len(op.Params)-1]
			c.assembleArg(k, &ctx, argStart+i)
			if ctx.Err != nil {
				return
			}
		}
	}
	op.Func(&ctx)
	if ctx.Err != nil {
		c.Err = ctx.Err
		return
	}
	c.vals, c.meta = c.vals[:argStart], c.meta[:argStart]

	for i, k := range op.Returns {
		c.assembleRet(k, ctx, i)
		if c.Err != nil {
			return
		}
	}
	if op.VarReturns {
		for i := len(op.Returns); i < len(ctx.Returns); i++ {
			k := op.Returns[len(op.Returns)-1]
			c.assembleRet(k, ctx, i)
		}
		if c.Err != nil {
			return
		}
	}
}

func (c *Calc) assembleArg(paramKind Kind, ctx *OpContext, idx int) {
	arg := c.vals[idx]
	argKind := c.meta[idx].Kind
	if !paramKind.Is(arg) {
		convArg, ok := paramKind.To(arg)
		if !ok {
			c.Err = errors.NewUnexpectedArgKind(paramKind.String(), argKind.String(), arg, idx)
			return
		}
		ctx.Args = append(ctx.Args, convArg)
	} else {
		ctx.Args = append(ctx.Args, arg)
	}
}

func (c *Calc) assembleRet(retKind Kind, ctx OpContext, idx int) {
	ret := ctx.Returns[idx]
	if !retKind.Is(ret) {
		c.Err = errors.NewUnexpectedRetKind(retKind.String(), ret, idx)
		return
	}
	c.push(ret, retKind)
}
