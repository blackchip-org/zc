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
	Trace      bool
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
	Stack   *Stack
	config  Config
	main    *Stack
	global  map[string]*Stack
	local   map[string]*Stack
	funcs   map[string]CalcFn
	defs    map[string]ModuleDef
	modules map[string]*Calc
}

func NewCalc(config Config) *Calc {
	c := &Calc{
		Out:     os.Stdout,
		config:  config,
		main:    NewStack(),
		global:  make(map[string]*Stack),
		defs:    make(map[string]ModuleDef),
		modules: make(map[string]*Calc),
		funcs:   make(map[string]CalcFn),
	}
	c.Stack = c.main
	c.local = c.global
	for name, fn := range builtin {
		c.funcs[name] = native(c, fn)
	}
	for _, def := range config.ModuleDefs {
		c.Install(def)
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

func (c *Calc) StackFor(name string) (*Stack, error) {
	stack, ok := c.local[name]
	if ok {
		return stack, nil
	}
	stack, ok = c.global[name]
	if ok {
		return stack, nil
	}
	return nil, fmt.Errorf("no such stack: %v", name)
}

func (c *Calc) Define(name string) *Stack {
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
	case *lang.ExprNode:
		return c.evalExprNode(n)
	case *lang.FileNode:
		return c.evalFileNode(n)
	case *lang.FuncNode:
		return c.evalFuncNode(n)
	case *lang.IncludeNode:
		return c.evalIncludeNode(n)
	case *lang.InvokeNode:
		return c.evalInvokeNode(n)
	case *lang.RefNode:
		return c.evalRefNode(n)
	case *lang.StackNode:
		return c.evalStackNode(n)
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

func (c *Calc) evalFuncNode(fn *lang.FuncNode) error {
	c.trace("define func: %v", fn.Name)
	c.funcs[fn.Name] = script(c, fn)
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
	fn, ok := c.funcs[invoke.Name]
	if !ok {
		return fmt.Errorf("no such function: %v", invoke.Name)
	}
	return fn()
}

func (c *Calc) evalExprNode(expr *lang.ExprNode) error {
	for _, node := range expr.Nodes {
		if err := c.evalNode(node); err != nil {
			return err
		}
	}
	c.Stack = c.main
	return nil
}

func (c *Calc) evalRefNode(ref *lang.RefNode) error {
	c.trace("ref %v%v", ref.Type, ref.Name)
	stack, err := c.StackFor(ref.Name)
	if err != nil {
		return err
	}

	switch ref.Type {
	case lang.AllRef:
		for _, item := range stack.Items() {
			c.Stack.Push(item)
		}
	case lang.TopRef:
		top, err := stack.Get()
		if err != nil {
			return err
		}
		c.Stack.Push(top)
	}
	return nil
}

func (c *Calc) evalStackNode(node *lang.StackNode) error {
	c.trace("stack %v", node.Name)
	stack := c.Define(node.Name)
	c.Stack = stack
	return nil
}

func (c *Calc) evalValueNode(value *lang.ValueNode) error {
	c.trace("value %v", value.Value)
	c.Stack.Push(value.Value)
	return nil
}

func (c *Calc) functionContext() *Calc {
	dc := &Calc{
		Out:     c.Out,
		config:  c.config,
		main:    NewStack(),
		global:  c.global,
		local:   make(map[string]*Stack),
		funcs:   c.funcs,
		defs:    c.defs,
		modules: c.modules,
	}
	dc.Stack = dc.main
	return dc
}

func (c *Calc) invokeFunction(node lang.NodeAST) error {
	fn := node.(*lang.FuncNode)
	dc := c.functionContext()
	for _, param := range fn.Params {
		val, err := c.Stack.Pop()
		if err != nil {
			return fmt.Errorf("not enough arguments, missing '%v'", param.Name)
		}
		c.trace("\tparam %v=%v\n", param.Name, val)
		dc.Define(param.Name).Set(val)
	}
	if err := dc.evalBody(fn.Body); err != nil {
		return err
	}
	for dc.Stack.Len() > 0 {
		val := dc.Stack.MustPop()
		c.trace("\treturn %v", val)
		c.Stack.Push(val)
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
		config:  c.config,
		main:    NewStack(),
		global:  make(map[string]*Stack),
		local:   nil,
		funcs:   make(map[string]CalcFn),
		defs:    c.defs,
		modules: c.modules,
	}
	dc.Stack = dc.main
	return dc
}

func (c *Calc) Install(def ModuleDef) {
	if def.Name == "" {
		panic(fmt.Sprintf("unable to install a module with no name: %+v", def))
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
		dc.funcs[name] = native(dc, fn)
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
	for name, fn := range dc.funcs {
		c.funcs[name] = fn
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
	if c.config.Trace {
		fmt.Fprint(c.Out, "eval: ")
		fmt.Fprintf(c.Out, format, a...)
		fmt.Fprintln(c.Out)
	}
}
