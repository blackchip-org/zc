package zc

import (
	"fmt"

	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v6/types"
)

type Calc struct {
	Stack    Stack
	Err      error
	Info     string
	Listener Listener
	Catalog  *Catalog
	scanner  scan.Scanner
}

func NewCalc(cat *Catalog) *Calc {
	return &Calc{Catalog: cat}
}

func (c *Calc) Push(a any) {
	if c.Err != nil {
		return
	}
	t, ok := c.Catalog.TypeOf(a)
	if !ok {
		panic(fmt.Errorf("unregistered type: %v", types.GoName(a)))
	}
	c.Stack.Push(Item{Value: a, Type: t})
	if c.Listener != nil {
		c.Listener(NewStackEvent("push", c.Stack))
	}
}

func (c *Calc) Pop(dest any) error {
	if c.Err != nil {
		panic(c.Err)
	}
	item, ok := c.Stack.Pop()
	if !ok {
		return ErrStackEmpty
	}
	destType, ok := c.Catalog.TypeOf(dest)
	if !ok {
		panic(fmt.Errorf("unregistered type: %v", types.GoName(dest)))
	}
	if !destType.Is(item.Value) {
		valConv, ok := destType.To(item.Value)
		if !ok {
			c.Err = fmt.Errorf("cannot convert %v from %v to %v", item.Value, item.Type.Name(), types.GoName(dest))
			return c.Err
		}
		destType.Copy(valConv, dest)
	} else {
		destType.Copy(item.Value, dest)
	}

	if c.Listener != nil {
		c.Listener(NewStackEvent("pop", c.Stack))
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

func (c *Calc) do(name string) {
	if c.Err != nil {
		return
	}
	c.Info = ""

	// An name may have more than one operation mapped to it. If there is
	// nothing at all matching by this name, set an error and return
	ops, ok := c.Catalog.OpFor(name)
	if !ok {
		c.Err = fmt.Errorf("no such operation: %v", name)
		return
	}

	// If this is a macro, simply evaluate it and return
	if len(ops) == 1 && ops[0].Macro != "" {
		c.Eval(ops[0].Macro)
		return
	}

	// Check each op to see if the kinds of the arguments match up with
	// the kinds of the parameters
	var op Op
	var err error
	var env *OpEnv
	for _, op = range ops {
		env, err = c.checkOp(name, op)
		if err == nil {
			break
		}
	}

	// No match found. When there is a single operation, return the actual
	// error that was returned. Otherwise, indicate that no matches were found
	if err != nil {
		if len(ops) > 1 {
			err = fmt.Errorf("no match for operation: %v", name)
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
	if op.VarReturn != nil {
		c.Stack.Clear()
	} else {
		c.Stack.PopN(len(env.Args))
	}

	// Check the return values from the operation and ensure the match
	// up with the expected kinds. If there is a mismatch here, this is
	// an implementation bug with the operation and a panic is raised.
	idx := 0
	for _, typ := range op.Returns {
		if !c.assembleRet(typ, env, idx) {
			panic(invalidRetTypes(name, op.Returns, op.VarReturn))
		}
		idx++
	}

	// If the return can contain a variable number of values, then each
	// additional value must be the same kind as the last found in the
	// return signature.
	if op.VarReturn != nil {
		for idx < len(env.Returns) {
			if !c.assembleRet(env.Op.VarReturn, env, idx) {
				panic(invalidRetTypes(name, op.Returns, op.VarReturn))
			}
			idx++
		}
		if c.Err != nil {
			return
		}
	}

	if c.Listener != nil {
		c.Listener(NewStackEvent("eval", c.Stack))
	}
}

func (c *Calc) checkOp(name string, op Op) (*OpEnv, error) {
	// First, make sure there are enough arguments for this operation
	stackLen := c.Stack.Len()
	if len(op.Params) > stackLen {
		return nil, fmt.Errorf("%v: not enough arguments, expected %v", name, len(op.Params))
	}

	// Create a call environment. Arguments are processed by starting
	// at the top of the stack and working towards index 0.
	var env OpEnv
	env.Catalog = c.Catalog

	// If there are a variable number of parameters, then consume
	// the entire stack.
	idx := stackLen - len(op.Params)
	if op.VarParam != nil {
		idx = 0
	}

	// Check that the kind of each argument matches up with the expected
	// parameter kind.
	for _, type_ := range op.Params {
		if !c.assembleArg(type_, &env, idx) {
			return nil, invalidArgTypes(name, op.Params, op.VarParam)
		}
		idx++
	}

	// If the arguments can contain a variable number of values, then each
	// additional value must be the same kind
	if op.VarParam != nil {
		for idx < stackLen {
			if !c.assembleArg(op.VarParam, &env, idx) {
				return nil, invalidArgTypes(name, op.Params, op.VarParam)
			}
			idx++
		}
	}
	return &env, nil
}

func (c *Calc) Do(names ...string) {
	c.Info = ""
	for _, name := range names {
		if c.Err != nil {
			return
		}
		c.do(name)
	}
}

func (c *Calc) assembleArg(typ Type, env *OpEnv, idx int) bool {
	item := c.Stack.At(idx)
	arg := item.Value
	if !typ.Is(arg) {
		convArg, ok := typ.To(arg)
		if !ok {
			return false
		}
		env.Args = append(env.Args, convArg)
	} else {
		env.Args = append(env.Args, arg)
	}
	return true
}

func (c *Calc) assembleRet(typ Type, env *OpEnv, idx int) bool {
	ret := env.Returns[idx]
	if !typ.Is(ret) {
		return false
	}
	c.Stack.Push(Item{Value: ret, Type: typ})
	return true
}

func invalidArgTypes(opName string, types []Type, varArgs Type) error {
	var names []string
	for _, t := range types {
		names = append(names, t.Name())
	}
	if varArgs != nil {
		names = append(names, varArgs.Name()+"*")
	}
	return fmt.Errorf("%v: invalid arguments, expected %v", opName, FormatList(names))
}

func invalidRetTypes(opName string, types []Type, varRets Type) error {
	var names []string
	for _, t := range types {
		names = append(names, t.Name())
	}
	if varRets != nil {
		names = append(names, varRets.Name()+"*")
	}
	return fmt.Errorf("%v: invalid returns, expected %v", opName, FormatList(names))
}

func (c *Calc) Eval(line string) error {
	c.scanner.InitFromString("", line)
	r := scan.NewRunner(&c.scanner, rules)
	toks := r.All()
	for _, tok := range toks {
		switch tok.Type {
		case TokenValue:
			c.Push(tok.Val)
		case TokenName:
			c.Do(tok.Val)
		default:
			panic(fmt.Errorf("unexpected token type: %v", tok.Type))
		}
		if c.Err != nil {
			return c.Err
		}
	}
	return nil
}
