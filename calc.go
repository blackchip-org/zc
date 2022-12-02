package zc

import (
	"embed"
	"fmt"
	"io"
	"os"

	"github.com/blackchip-org/zc/lang"
)

type Config struct {
	ModuleDefs []ModuleDef
}

type CalcFn func() error
type NativeFn func(*Calc) error

func script(c *Calc, ast lang.NodeAST) CalcFn {
	return func() error { return c.invokeFunction(ast) }
}

func native(c *Calc, fn NativeFn) CalcFn {
	return func() error { return fn(c) }
}

//go:embed internal/modules/*.zc
var scripts embed.FS

type Calc struct {
	Out     io.Writer
	Trace   bool
	main    *Stack
	stack   *Stack
	global  map[string]*Stack
	local   map[string]*Stack
	fn      map[string]CalcFn
	defs    map[string]ModuleDef
	modules map[string]*Calc
}

func NewCalc(config Config) *Calc {
	c := &Calc{
		Out:     os.Stdout,
		main:    NewStack(),
		global:  make(map[string]*Stack),
		defs:    make(map[string]ModuleDef),
		modules: make(map[string]*Calc),
		fn:      make(map[string]CalcFn),
	}
	c.stack = c.main
	c.local = c.global
	for name, fn := range builtin {
		c.fn[name] = native(c, fn)
	}
	for _, def := range config.ModuleDefs {
		c.Define(def)
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

func (c *Calc) Local(name string) *Stack {
	stack, ok := c.local[name]
	if !ok {
		stack = NewStack()
		c.local[name] = stack
	}
	return stack
}

func (c *Calc) evalBody(nodes []lang.NodeAST) error {
	for _, node := range nodes {
		if err := c.evalNode(node); err != nil {
			return err
		}
	}
	return nil
}

func (c *Calc) evalNode(node lang.NodeAST) error {
	switch n := node.(type) {
	case *lang.FileNode:
		return c.evalFileNode(n)
	case *lang.FuncNode:
		return c.evalFnNode(n)
	case *lang.IncludeNode:
		return c.evalIncludeNode(n)
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

func (c *Calc) evalFnNode(fn *lang.FuncNode) error {
	c.trace("define func: %v", fn.Name)
	c.fn[fn.Name] = script(c, fn)
	return nil
}

func (c *Calc) evalIncludeNode(include *lang.IncludeNode) error {
	for _, name := range include.Names {
		c.trace("include %v", name)
		if err := c.Include(name); err != nil {
			return err
		}
	}
	return nil
}

func (c *Calc) evalInvokeNode(invoke *lang.InvokeNode) error {
	c.trace("invoke %v", invoke.Name)
	fn, ok := c.fn[invoke.Name]
	if !ok {
		return fmt.Errorf("no such function: %v", invoke.Name)
	}
	return fn()
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
	c.trace("value %v", value.Value)
	c.stack.Push(value.Value)
	return nil
}

func (c *Calc) functionContext() *Calc {
	dc := &Calc{
		Out:     c.Out,
		main:    NewStack(),
		global:  c.global,
		local:   make(map[string]*Stack),
		fn:      c.fn,
		defs:    c.defs,
		modules: c.modules,
	}
	dc.stack = dc.main
	return dc
}

func (c *Calc) invokeFunction(node lang.NodeAST) error {
	fn := node.(*lang.FuncNode)
	c.trace("***** MADE IT HERE %v", fn.Name)
	dc := c.functionContext()
	for _, param := range fn.Params {
		val, err := c.Stack().Pop()
		if err != nil {
			return err
		}
		c.trace("\tparam %v=%v\n", param.Name, val)
		dc.Local(param.Name).Set(val)
	}
	if err := dc.evalBody(fn.Body); err != nil {
		return err
	}
	for dc.Stack().Len() > 0 {
		val := dc.Stack().MustPop()
		c.trace("\treturn %v", val)
		c.Stack().Push(val)
	}
	return nil
}

type ModuleDef struct {
	Name    string
	Scripts []string
	Natives map[string]NativeFn
}

func (c *Calc) moduleContext() *Calc {
	dc := &Calc{
		Out:     c.Out,
		main:    NewStack(),
		global:  make(map[string]*Stack),
		local:   nil,
		fn:      make(map[string]CalcFn),
		defs:    c.defs,
		modules: c.modules,
	}
	dc.stack = dc.main
	return dc
}

func (c *Calc) Define(def ModuleDef) {
	if def.Name == "" {
		panic(fmt.Sprintf("unable to define a module with no name: %+v", def))
	}
	c.defs[def.Name] = def
}

func (c *Calc) load(name string) (*Calc, error) {
	def, ok := c.defs[name]
	if !ok {
		return nil, fmt.Errorf("no such module: %v", name)
	}

	dc := c.moduleContext()
	for name, fn := range def.Natives {
		dc.fn[name] = native(dc, fn)
	}

	for _, path := range def.Scripts {
		src, err := scripts.ReadFile(path)
		if err != nil {
			return nil, err
		}
		ast, err := lang.Parse(path, src)
		if err != nil {
			return nil, err
		}
		if err := dc.evalNode(ast); err != nil {
			return nil, err
		}
	}
	c.modules[def.Name] = dc
	return dc, nil
}

func (c *Calc) Include(name string) error {
	dc, err := c.load(name)
	if err != nil {
		return err
	}
	for name, fn := range dc.fn {
		c.fn[name] = fn
	}
	return nil
}

func (c *Calc) Printf(format string, a ...any) {
	fmt.Fprintf(c.Out, format, a...)
}

func (c *Calc) Println(a any) {
	fmt.Fprintln(c.Out, a)
}

func (c *Calc) Print(a any) {
	fmt.Fprint(c.Out, a)
}

func (c *Calc) trace(format string, a ...any) {
	if c.Trace {
		fmt.Fprint(c.Out, "eval: ")
		fmt.Fprintf(c.Out, format, a...)
		fmt.Fprintln(c.Out)
	}
}
