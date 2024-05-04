package zc

import (
	"fmt"
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
	Op      Op
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
	Prec       int
	Func       func(*OpEnv)
}

type Vol struct {
	Name  string
	Kinds []Kind
	Ops   []Op
}

type Item struct {
	Value any
	Kind  Kind
	Anno  string
}

func (i Item) String() string {
	return fmt.Sprintf("%v", i.Value)
}

func Values(items []Item) []any {
	var vals []any
	for _, item := range items {
		vals = append(vals, item.Value)
	}
	return vals
}

type Calc struct {
	Err      error
	Listener Listener
	cat      *Catalog
	stack    stack.Stack[Item]
}

func NewCalc(cat *Catalog) *Calc {
	c := &Calc{
		cat:   cat,
		stack: stack.NewSlice[Item](),
	}
	return c
}

func (c *Calc) push(item Item) {
	c.stack.Push(item)
}

func (c *Calc) pop() (Item, error) {
	item, ok := c.stack.Pop()
	if !ok {
		return item, errors.StackEmpty()
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
	c.push(Item{Value: a, Kind: k})
	if c.Listener != nil {
		c.Listener(NewStackEvent("push", c.stack.Items()))
	}
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

	if c.Listener != nil {
		c.Listener(NewStackEvent("pop", c.stack.Items()))
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

	// An name may have more than one operation mapped to it. If there is
	// nothing at all matching by this name, set an error and return
	ops, ok := c.cat.OpByName(name)
	if !ok {
		c.Err = errors.NoSuchOp(name)
		return
	}

	// Check each op to see if the kinds of the arguments match up with
	// the kinds of the parameters
	var op Op
	var err error
	var env *OpEnv
	for _, op = range ops {
		env, err = c.checkOp(op)
		if err == nil {
			break
		}
	}

	// No match found. When there is a single operation, return the actual
	// error that was returned. Otherwise, indicate that no matches were found
	if err != nil {
		if len(ops) > 1 {
			err = errors.NoMatchOp(name)
		}
		c.Err = err
		return
	}

	// Execute the operation and return now if there was an error
	env.Op = op
	if c.Listener != nil {
		c.Listener(NewOpEvent(env))
	}
	op.Func(env)

	if env.Err != nil {
		c.Err = env.Err
		return
	}

	// Arguments that were used for the operation now have to be removed
	// from the stack. If there are variable returns, remove all items.
	if op.VarReturns {
		c.stack.Clear()
	} else {
		stack.PopN(c.stack, len(env.Args))
	}

	// Check the return values from the operation and ensure the match
	// up with the expected kinds. If there is a mismatch here, this is
	// an implementation bug with the operation and a panic is raised.
	for i, kindName := range op.Returns {
		if !c.assembleRet(kindName, env, i) {
			panic(errors.InvalidRetKinds(op.Returns, op.VarReturns))
		}
	}

	// If the return can contain a variable number of values, then each
	// additional value must be the same kind as the last found in the
	// return signature.
	if op.VarReturns {
		for i := len(op.Returns); i < len(env.Returns); i++ {
			k := op.Returns[len(op.Returns)-1]
			if !c.assembleRet(k, env, i) {
				panic(errors.InvalidRetKinds(op.Returns, op.VarReturns))
			}
		}
		if c.Err != nil {
			return
		}
	}

	if c.Listener != nil {
		c.Listener(NewStackEvent("eval", c.stack.Items()))
	}
}

func (c *Calc) checkOp(op Op) (*OpEnv, error) {
	// First, make sure there are enough arguments for this operation
	stackLen := c.stack.Len()
	if len(op.Params) > stackLen {
		return nil, errors.InvalidArgCount(len(op.Params))
	}

	// Create a call environment and find which value on the stack is the
	// first argument. If the operation is defined to take a variable
	// number of arguments, then the whole stack is the argument set.
	var env OpEnv
	env.Catalog = c.cat
	argStart := stackLen - len(op.Params)
	if op.VarParams {
		argStart = 0
	}

	// Check that the kind of each argument matches up with the expected
	// parameter kind.
	for i, kindName := range op.Params {
		if !c.assembleArg(kindName, &env, argStart+i) {
			return nil, errors.InvalidArgKinds(op.Params, op.VarParams)
		}
	}

	// If the arguments can contain a variable number of values, then each
	// additional value must be the same kind as the last found in the
	// parameter signature.
	if op.VarParams {
		for i := len(op.Params); i < stackLen; i++ {
			kindName := op.Params[len(op.Params)-1]
			if !c.assembleArg(kindName, &env, argStart+i) {
				return nil, errors.InvalidArgKinds(op.Params, op.VarParams)
			}
		}
	}
	return &env, nil
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

func (c *Calc) assembleRet(kindName string, env *OpEnv, idx int) bool {
	retKind, ok := c.cat.KindByName(kindName)
	if !ok {
		panic(errors.UnknownKind(kindName))
	}
	ret := env.Returns[idx]
	if !retKind.Is(ret) {
		return false
	}
	c.push(Item{Value: ret, Kind: retKind})
	return true
}
