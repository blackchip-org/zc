package zc

import (
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
	Value any
	Kind  Kind
	Anno  string
}

func Values(items []Item) []any {
	var vals []any
	for _, item := range items {
		vals = append(vals, item.Value)
	}
	return vals
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
	c.stack.Push(Item{Value: a, Kind: k})
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
	k, ok := c.cat.KindByType(a)
	if !ok {
		panic(errors.UnregisteredType(a))
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

	destKind, ok := c.cat.KindByType(dest)
	if !ok {
		panic(errors.UnregisteredType(dest))
	}
	if !destKind.Is(item.Value) {
		valConv, ok := destKind.To(item.Value)
		if !ok {
			c.Err = errors.CannotConvert(item.Kind.Name(), dest, item.Value)
			return c.Err
		}
		destKind.Copy(valConv, dest)
	} else {
		destKind.Copy(item.Value, dest)
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

func (c *Calc) do(name string) {
	if c.Err != nil {
		return
	}

	op, ok := c.cat.OpByName(name)
	if !ok {
		c.Err = errors.UnknownOp(name)
		return
	}

	stackLen := c.stack.Len()
	if len(op.Params) > stackLen {
		c.Err = errors.InvalidArgCount(len(op.Params))
		return
	}

	var env OpEnv
	env.Catalog = c.cat
	argStart := stackLen - len(op.Params)
	if op.VarParams {
		argStart = 0
	}

	for i, kindName := range op.Params {
		if !c.assembleArg(kindName, &env, argStart+i) {
			c.Err = errors.InvalidArgKinds(op.Params, op.VarParams)
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
		if !c.assembleRet(kindName, env, i) {
			panic(errors.InvalidRetKinds(op.Returns, op.VarReturns))
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

func (c *Calc) Do(names ...string) {
	for _, name := range names {
		if c.Err != nil {
			return
		}
		c.do(name)
	}
}

func (c *Calc) assembleArg(kindName string, env *OpEnv, idx int) bool {
	paramKind, ok := c.cat.KindByName(kindName)
	if !ok {
		panic(errors.UnknownKind(kindName))
	}
	item := stack.At(c.stack, idx)
	arg := item.Value
	if !paramKind.Is(arg) {
		convArg, ok := paramKind.To(arg)
		if !ok {
			return false
		}
		env.Args = append(env.Args, convArg)
	} else {
		env.Args = append(env.Args, arg)
	}
	return true
}

func (c *Calc) assembleRet(kindName string, ctx OpEnv, idx int) bool {
	retKind, ok := c.cat.KindByName(kindName)
	if !ok {
		panic(errors.UnknownKind(kindName))
	}
	ret := ctx.Returns[idx]
	if !retKind.Is(ret) {
		return false
	}
	c.push(ret, retKind)
	return true
}
