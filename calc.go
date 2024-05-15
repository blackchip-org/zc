package zc

import (
	"fmt"
	"slices"
	"strings"

	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v6/pkg/stack"
)

type Calc struct {
	Stack    stack.Stack[Item]
	Err      error
	Info     string
	Listener Listener
	Catalog  *Catalog
}

func NewCalc(cat *Catalog) *Calc {
	c := &Calc{
		Catalog: cat,
		Stack:   stack.NewSlice[Item](),
	}
	return c
}

func (c *Calc) push(item Item) {
	c.Stack.Push(item)
}

func (c *Calc) pop() (Item, error) {
	item, ok := c.Stack.Pop()
	if !ok {
		return item, fmt.Errorf("stack empty")
	}
	return item, nil
}

func (c *Calc) Push(a any) {
	if c.Err != nil {
		return
	}
	k, ok := c.Catalog.KindByType(a)
	if !ok {
		panic(fmt.Errorf("unregistered type: %v", TypeName(a)))
	}
	c.push(Item{Value: a, Kind: k})
	if c.Listener != nil {
		c.Listener(NewStackEvent("push", c.Stack))
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

	destKind, ok := c.Catalog.KindByType(dest)
	if !ok {
		panic(fmt.Errorf("unregistered type: %v", TypeName(dest)))
	}
	if !destKind.Is(item.Value) {
		valConv, ok := destKind.To(item.Value)
		if !ok {
			c.Err = fmt.Errorf("cannot convert %v from %v to %v", item.Value, item.Kind.Name(), TypeName(dest))
			return c.Err
		}
		destKind.Copy(valConv, dest)
	} else {
		destKind.Copy(item.Value, dest)
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
	ops, ok := c.Catalog.OpByName(name)
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
	if op.VarReturn != "" {
		c.Stack.Clear()
	} else {
		stack.PopN(c.Stack, len(env.Args))
	}

	// Check the return values from the operation and ensure the match
	// up with the expected kinds. If there is a mismatch here, this is
	// an implementation bug with the operation and a panic is raised.
	idx := 0
	for _, kindName := range op.Returns {
		if !c.assembleRet(kindName, env, idx) {
			panic(invalidRetKinds(name, op.Returns, op.VarReturn))
		}
		idx++
	}

	// If the return can contain a variable number of values, then each
	// additional value must be the same kind as the last found in the
	// return signature.
	if op.VarReturn != "" {
		for idx < len(env.Returns) {
			if !c.assembleRet(env.Op.VarReturn, env, idx) {
				panic(invalidRetKinds(name, op.Returns, op.VarReturn))
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
	if op.VarParam != "" {
		idx = 0
	}

	// Check that the kind of each argument matches up with the expected
	// parameter kind.
	for _, kindName := range op.Params {
		if !c.assembleArg(kindName, &env, idx) {
			return nil, invalidArgKinds(name, op.Params, op.VarParam)
		}
		idx++
	}

	// If the arguments can contain a variable number of values, then each
	// additional value must be the same kind
	if op.VarParam != "" {
		for idx < stackLen {
			if !c.assembleArg(op.VarParam, &env, idx) {
				return nil, invalidArgKinds(name, op.Params, op.VarParam)
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

func (c *Calc) assembleArg(kindName string, env *OpEnv, idx int) bool {
	paramKind, ok := c.Catalog.KindByName(kindName)
	if !ok {
		panic(fmt.Errorf("unknown kind: %v", kindName))
	}
	item := stack.At(c.Stack, idx)
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
	retKind, ok := c.Catalog.KindByName(kindName)
	if !ok {
		panic(fmt.Errorf("unknown kind: %v", kindName))
	}
	ret := env.Returns[idx]
	if !retKind.Is(ret) {
		return false
	}
	c.push(Item{Value: ret, Kind: retKind})
	return true
}

var isBeginQuote = scan.Rune('\'', '"', '[')

func scanWord(s *scan.Scanner) string {
	var inQuote rune
	for s.HasMore() {
		switch {
		case inQuote == 0 && isBeginQuote(s.This):
			inQuote = s.This
			s.Skip()
		case inQuote == '\'' && s.This == '\'':
			inQuote = 0
			s.Skip()
		case inQuote == '"' && s.This == '"':
			inQuote = 0
			s.Skip()
		case inQuote == '[' && s.This == ']':
			inQuote = 0
			s.Skip()
		case scan.IsSpace(s.This):
			return s.Emit().Val
		default:
			s.Keep()
		}
	}
	return s.Emit().Val
}

func (c *Calc) Eval(line string) error {
	c.Err = nil
	s := scan.NewScannerFromString("", line)
	var words []string
	for s.HasMore() {
		scan.While(s, scan.IsSpace, s.Skip)
		word := scanWord(s)
		if word != "" {
			words = append(words, word)
		}
	}
	for _, w := range words {
		if c.Err != nil {
			return c.Err
		}
		if IsValue(w) {
			c.Push(w)
		} else {
			c.Do(w)
		}
	}
	return c.Err
}

func invalidArgKinds(opName string, kindNames []string, varArgs string) error {
	ks := slices.Clone(kindNames)
	if varArgs != "" {
		ks = append(ks, varArgs+"*")
	}
	return fmt.Errorf("%v: invalid arguments, expected %v", opName, strings.Join(ks, " | "))
}

func invalidRetKinds(opName string, kindNames []string, varRets string) error {
	ks := slices.Clone(kindNames)
	if varRets != "" {
		ks = append(ks, varRets+"*")
	}
	return fmt.Errorf("%v: invalid returns, expected %v", opName, strings.Join(ks, " | "))
}
