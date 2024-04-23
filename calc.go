package zc

import (
	"reflect"
	"slices"

	"github.com/blackchip-org/zc/v6/errors"
	"github.com/blackchip-org/zc/v6/pkg/stack"
)

type Kind interface {
	Name() string
	Is(any) bool
	Dup(any) any
	Copy(any, any)
	To(any) (any, bool)
}

type OpEnv struct {
	Catalog *Catalog
	Args    []any
	Returns []any
	Err     error
}

type Op struct {
	Name       string
	Aliases    []string
	Params     []string
	VarParams  bool
	Returns    []string
	VarReturns bool
	Func       func(*OpEnv)
}

type Volume struct {
	Name  string
	Kinds []Kind
	Ops   []Op
}

type Item struct {
	Val  any
	Kind Kind
	Anno string
}

type Calc struct {
	Err   error
	cat   *Catalog
	stack stack.Stack[Item]
}

func NewCalc(cat *Catalog) *Calc {
	c := &Calc{
		cat:   cat,
		stack: stack.NewSlice[Item](),
	}
	return c
}

func (c *Calc) push(a any, k Kind) {
	c.stack.Push(Item{Val: a, Kind: k})
}

func (c *Calc) pop() (Item, error) {
	item, ok := c.stack.Pop()
	if !ok {
		return item, errors.StackEmpty
	}
	return item, nil
}

func (c *Calc) Push(a any) {
	if c.Err != nil {
		return
	}
	if a == nil {
		panic(errors.IllegalNil)
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
	item, err := c.pop()
	if err != nil {
		return err
	}

	destKind, ok := c.cat.KindOf(dest)
	if !ok {
		panic(errors.NewUnregisteredType(dest))
	}
	if !destKind.Is(item.Val) {
		valConv, ok := destKind.To(item.Val)
		if !ok {
			vt := reflect.TypeOf(item.Val).Name()
			return errors.NewInvalidConversion(item.Kind.Name(), vt, valConv)
		}
		destKind.Copy(valConv, dest)
	} else {
		destKind.Copy(item.Val, dest)
	}
	return nil
}

func (c *Calc) PopString() string {
	var r string
	if err := c.Pop(&r); err != nil {
		panic(err)
	}
	return r
}

func (c *Calc) Stack() []Item {
	return slices.Clone(c.stack.Items())
}

func (c *Calc) Do(name string) {
	if c.Err != nil {
		return
	}

	op, ok := c.cat.OpFor(name)
	if !ok {
		c.Err = errors.UnknownOp(name)
		return
	}

	stackLen := c.stack.Len()
	if len(op.Params) > stackLen {
		c.Err = errors.StackEmpty
		return
	}

	var env OpEnv
	env.Catalog = c.cat
	argStart := stackLen - len(op.Params)
	if op.VarParams {
		argStart = 0
	}
	for i, kindName := range op.Params {
		c.assembleArg(kindName, &env, argStart+i)
		if env.Err != nil {
			return
		}
	}
	if op.VarParams {
		for i := len(op.Params); i < stackLen; i++ {
			kindName := op.Params[len(op.Params)-1]
			c.assembleArg(kindName, &env, argStart+i)
			if env.Err != nil {
				return
			}
		}
	}
	op.Func(&env)
	if env.Err != nil {
		c.Err = env.Err
		return
	}
	stack.PopN(c.stack, c.stack.Len()-argStart)

	for i, kindName := range op.Returns {
		c.assembleRet(kindName, env, i)
		if c.Err != nil {
			return
		}
	}
	if op.VarReturns {
		for i := len(op.Returns); i < len(env.Returns); i++ {
			k := op.Returns[len(op.Returns)-1]
			c.assembleRet(k, env, i)
		}
		if c.Err != nil {
			return
		}
	}
}

func (c *Calc) assembleArg(kindName string, ctx *OpEnv, idx int) {
	paramKind, ok := c.cat.KindFor(kindName)
	if !ok {
		panic(errors.UnknownKind(kindName))
	}
	item := stack.At(c.stack, idx)
	arg := item.Val
	argKind := item.Kind
	if !paramKind.Is(arg) {
		convArg, ok := paramKind.To(arg)
		if !ok {
			c.Err = errors.NewUnexpectedArgKind(paramKind.Name(), argKind.Name(), arg, idx)
			return
		}
		ctx.Args = append(ctx.Args, convArg)
	} else {
		ctx.Args = append(ctx.Args, arg)
	}
}

func (c *Calc) assembleRet(kindName string, ctx OpEnv, idx int) {
	retKind, ok := c.cat.KindFor(kindName)
	if !ok {
		panic(errors.UnknownKind(kindName))
	}
	ret := ctx.Returns[idx]
	if !retKind.Is(ret) {
		c.Err = errors.NewUnexpectedRetKind(retKind.Name(), ret, idx)
		return
	}
	c.push(ret, retKind)
}
