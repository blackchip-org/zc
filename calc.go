package zc

import (
	"embed"
	"fmt"
	"io"
	"os"

	"github.com/blackchip-org/zc/lang"
)

type NativeFn func(*Calc) error

type Module struct {
	Natives map[string]NativeFn
}

//go:embed test/*.zc
var Static embed.FS

type Calc struct {
	Out    io.Writer
	main   *Stack
	stack  *Stack
	global map[string]*Stack
	local  map[string]*Stack
	fn     map[string]NativeFn
}

func NewCalc(prelude []Module) *Calc {
	c := &Calc{
		Out:    os.Stdout,
		main:   NewStack(),
		global: make(map[string]*Stack),
		local:  nil,
		fn:     make(map[string]NativeFn),
	}
	c.stack = c.main

	for _, mod := range prelude {
		if mod.Natives != nil {
			for name, native := range mod.Natives {
				c.fn[name] = native
			}
		}
	}
	return c
}

func (c *Calc) Eval(src []byte) error {
	ast, errs := lang.Parse("", src)
	if errs != nil {
		return errs
	}
	err := c.evalNode(ast)
	return err
}

func (c *Calc) EvalString(src string) error {
	return c.Eval([]byte(src))
}

func (c *Calc) Stack() *Stack {
	return c.stack
}

func (c *Calc) Global(name string) *Stack {
	stack, ok := c.global[name]
	if !ok {
		stack = NewStack()
		c.global[name] = stack
	}
	return stack
}

func (c *Calc) evalNode(node lang.NodeAST) error {
	switch n := node.(type) {
	case *lang.FileNode:
		return c.evalFileNode(n)
	case *lang.InvokeNode:
		return c.evalInvokeNode(n)
	case *lang.ExprNode:
		return c.evalLineNode(n)
	case *lang.ValueNode:
		return c.evalValueNode(n)
	}
	panic(fmt.Sprintf("unknown node: %+v", node))
}

func (c *Calc) evalFileNode(file *lang.FileNode) error {
	for _, line := range file.Nodes {
		if err := c.evalNode(line); err != nil {
			return err
		}
	}
	return nil
}

func (c *Calc) evalInvokeNode(invoke *lang.InvokeNode) error {
	fn, ok := c.fn[invoke.Name]
	if !ok {
		return fmt.Errorf("no such function: %v", invoke.Name)
	}
	return fn(c)
}

func (c *Calc) evalLineNode(line *lang.ExprNode) error {
	for _, node := range line.Nodes {
		if err := c.evalNode(node); err != nil {
			return err
		}
	}
	return nil
}

func (c *Calc) evalValueNode(value *lang.ValueNode) error {
	c.stack.Push(value.Value)
	return nil
}
