package zc

import (
	"fmt"
	"log"
	"strings"

	"github.com/blackchip-org/zc/internal"
	"github.com/blackchip-org/zc/lang/ast"
	"github.com/blackchip-org/zc/lang/parser"
	"github.com/blackchip-org/zc/lang/token"
	"github.com/shopspring/decimal"
)

var (
	Places    int32  = 16
	RoundMode string = "half-up"
)

type Config struct {
	ModuleDefs []ModuleDef
	Prelude    []string
	Trace      bool
}

type ModuleDef struct {
	Name       string
	ScriptPath string
	Natives    map[string]CalcFunc
}

type Frame struct {
	Pos  token.Pos
	Func string
}

func (f Frame) String() string {
	return fmt.Sprintf("[%v] %v", f.Pos, f.Func)
}

type CalcError struct {
	Message string
	Frames  []Frame
}

func (c CalcError) Error() string {
	return c.Message
}

type CalcFunc func(*Calc) error

type Calc struct {
	Out     *strings.Builder
	Stack   *Stack
	name    string
	parent  *Calc
	frame   Frame
	config  Config
	main    *Stack
	global  map[string]*Stack
	local   map[string]*Stack
	funcs   map[string]CalcFunc
	defs    map[string]ModuleDef
	modules map[string]*Calc
}

func NewCalc(config Config) (*Calc, error) {
	c := &Calc{
		Out:     &strings.Builder{},
		name:    "<cli>",
		config:  config,
		main:    NewStack("main"),
		global:  make(map[string]*Stack),
		defs:    make(map[string]ModuleDef),
		modules: make(map[string]*Calc),
		funcs:   make(map[string]CalcFunc),
	}
	c.Stack = c.main
	c.local = c.global

	for _, def := range config.ModuleDefs {
		c.Install(def)
	}
	for name, fn := range builtin {
		c.funcs[name] = fn
	}
	for _, prelude := range config.Prelude {
		if err := c.Include(prelude); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *Calc) Eval(src []byte) (err error) {
	c.Out.Reset()
	ast, err := parser.Parse("", src)
	if err != nil {
		return
	}
	err = c.evalNode(ast)

	return
}

func (c *Calc) EvalString(src string) error {
	return c.Eval([]byte(src))
}

func (c *Calc) Define(name string) *Stack {
	stack, ok := c.local[name]
	if !ok {
		stack, ok = c.global[name]
		if !ok {
			stack = NewStack(name)
			c.local[name] = stack
		}
	}
	return stack
}

func (c *Calc) Import(modName string) error {
	dc, err := c.load(modName)
	if err != nil {
		return err
	}
	for funcName, fn := range dc.funcs {
		qName := modName + "." + funcName
		c.funcs[qName] = fn
	}
	return nil
}

func (c *Calc) Include(modName string) error {
	dc, err := c.load(modName)
	if err != nil {
		return err
	}
	for funcName, fn := range dc.funcs {
		c.funcs[funcName] = fn
	}
	return nil
}

func (c *Calc) Install(def ModuleDef) {
	if def.Name == "" {
		panic(fmt.Sprintf("unable to install a module with no name: %+v", def))
	}
	c.defs[def.Name] = def
}

func (c Calc) Interpolate(v string) (string, error) {
	var result, name strings.Builder

	inQuote := false
	inEscape := false

	for _, ch := range v {
		if ch == '`' && !inQuote && !inEscape {
			inQuote = true
		} else if ch == '`' && !inEscape {
			inQuote = false
			stack, err := c.StackFor(name.String())
			if err != nil {
				return "", fmt.Errorf("no such stack: %v", name.String())
			}
			for i, item := range stack.Items() {
				if i != 0 {
					result.WriteString("  ")
				}
				result.WriteString(item)
			}
			name.Reset()
		} else if ch == '\\' {
			inEscape = true
		} else {
			inEscape = false
			if inQuote {
				name.WriteRune(ch)
			} else {
				result.WriteRune(ch)
			}
		}
	}
	if name.Len() > 0 {
		return "", fmt.Errorf("expected`")
	}
	return result.String(), nil
}

func (c *Calc) PopBool() (bool, error) {
	v, err := c.Stack.Pop()
	if err != nil {
		return false, err
	}
	b, err := ParseBool(v)
	if err != nil {
		return false, err
	}
	return b, nil
}

func (c *Calc) PopDecimal() (decimal.Decimal, error) {
	v, err := c.Stack.Pop()
	if err != nil {
		return decimal.Zero, err
	}
	d, err := ParseDecimal(v)
	if err != nil {
		return decimal.Zero, err
	}
	return d, err
}

func (c *Calc) PopInt() (int, error) {
	v, err := c.Stack.Pop()
	if err != nil {
		return 0, err
	}
	i, err := ParseInt(v)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (c *Calc) PopInt32() (int32, error) {
	v, err := c.Stack.Pop()
	if err != nil {
		return 0, err
	}
	i, err := ParseInt32(v)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (c *Calc) Printf(format string, a ...any) {
	c.Out.WriteString(fmt.Sprintf(format, a...))
}

func (c *Calc) Print(a any) {
	c.Out.WriteString(fmt.Sprint(a))
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

func (c *Calc) evalBlock(nodes []ast.Node) error {
	for _, node := range nodes {
		if err := c.evalNode(node); err != nil {
			return err
		}
	}
	return nil
}

func (c *Calc) evalNode(node ast.Node) error {
	switch n := node.(type) {
	case *ast.ExprNode:
		return c.evalExprNode(n)
	case *ast.IfNode:
		return c.evalIfNode(n)
	case *ast.FileNode:
		return c.evalFileNode(n)
	case *ast.FuncNode:
		return c.evalFuncNode(n)
	case *ast.ImportNode:
		return c.evalImportNode(n)
	case *ast.IncludeNode:
		return c.evalIncludeNode(n)
	case *ast.InvokeNode:
		return c.evalInvokeNode(n)
	case *ast.MacroNode:
		return c.evalMacroNode(n)
	case *ast.RefNode:
		return c.evalRefNode(n)
	case *ast.StackNode:
		return c.evalStackNode(n)
	case *ast.ValueNode:
		return c.evalValueNode(n)
	case *ast.WhileNode:
		return c.evalWhileNode(n)
	}
	panic(fmt.Sprintf("unknown node: %+v", node))
}

func (c *Calc) evalExprNode(expr *ast.ExprNode) error {
	for _, node := range expr.Expr {
		if err := c.evalNode(node); err != nil {
			return err
		}
	}
	c.Stack = c.main
	return nil
}

func (c *Calc) evalIfNode(ifNode *ast.IfNode) error {
	for _, caseNode := range ifNode.Cases {
		// Final "else" condition will have no case expression
		if caseNode.Cond == nil {
			return c.evalBlock(caseNode.Block)
		} else {
			err := c.evalExprNode(caseNode.Cond)
			if err != nil {
				return err
			}
			v, err := c.Stack.Pop()
			if err != nil {
				return err
			}
			vb, err := ParseBool(v)
			if err != nil {
				return err
			}
			if vb {
				return c.evalBlock(caseNode.Block)
			}
		}
	}
	return nil
}

func (c *Calc) evalFileNode(file *ast.FileNode) error {
	for _, line := range file.Block {
		if err := c.evalNode(line); err != nil {
			return err
		}
	}
	return nil
}

func (c *Calc) evalFuncNode(fn *ast.FuncNode) error {
	c.trace(fn, "define func: %v", fn.Name)
	c.funcs[fn.Name] = func(ic *Calc) error {
		return ic.invokeFunction(c, fn)
	}
	return nil
}

func (c *Calc) evalImportNode(importNode *ast.ImportNode) error {
	for _, name := range importNode.Names {
		c.trace(importNode, "import %v", name)
		if err := c.Import(name); err != nil {
			return err
		}
	}
	return nil
}

func (c *Calc) evalIncludeNode(include *ast.IncludeNode) error {
	for _, name := range include.Names {
		c.trace(include, "include %v", name)
		if err := c.Include(name); err != nil {
			return err
		}
	}
	return nil
}

func (c *Calc) evalInvokeNode(invoke *ast.InvokeNode) error {
	c.trace(invoke, "invoke %v", invoke.Name)
	fn, ok := c.funcs[invoke.Name]
	if !ok {
		return fmt.Errorf("no such function: %v", invoke.Name)
	}
	c.frame = Frame{
		Pos:  invoke.Pos(),
		Func: invoke.Name,
	}
	if err := fn(c); err != nil {
		return err
	}
	c.frame = Frame{}
	return nil
}

func (c *Calc) evalMacroNode(mac *ast.MacroNode) error {
	c.trace(mac, "define macro: %v", mac.Name)
	c.funcs[mac.Name] = func(_ *Calc) error {
		return c.invokeMacro(mac)
	}
	return nil
}

func (c *Calc) evalRefNode(ref *ast.RefNode) error {
	c.trace(ref, "ref %v%v", ref.Type, ref.Name)
	stack, err := c.StackFor(ref.Name)
	if err != nil {
		return err
	}

	switch ref.Type {
	case ast.AllRef:
		for _, item := range stack.Items() {
			c.Stack.Push(item)
		}
	case ast.TopRef:
		top, err := stack.Get()
		if err != nil {
			return err
		}
		c.Stack.Push(top)
	}
	return nil
}

func (c *Calc) evalStackNode(node *ast.StackNode) error {
	c.trace(node, "stack %v", node.Name)
	stack := c.Define(node.Name)
	c.Stack = stack
	return nil
}

func (c *Calc) evalValueNode(value *ast.ValueNode) error {
	c.trace(value, "value %v", value.Value)
	interp, err := c.Interpolate(value.Value)
	if err != nil {
		return c.err(value, err.Error())
	}
	if interp != value.Value {
		c.trace(value, "interpolate %v", interp)
	}
	c.Stack.Push(interp)
	return nil
}

func (c *Calc) evalWhileNode(while *ast.WhileNode) error {
	c.trace(while, "while")
	for {
		if err := c.evalExprNode(while.Cond); err != nil {
			return err
		}
		result, err := c.PopBool()
		if err != nil {
			return err
		}
		if !result {
			break
		}
		if err := c.evalBlock(while.Block); err != nil {
			return err
		}
	}
	return nil
}

func (c *Calc) moduleContext(name string) *Calc {
	dc := &Calc{
		Out:     c.Out,
		name:    name,
		parent:  c,
		config:  c.config,
		main:    NewStack("main"),
		global:  make(map[string]*Stack),
		funcs:   make(map[string]CalcFunc),
		defs:    c.defs,
		modules: c.modules,
	}
	dc.Stack = dc.main
	dc.local = dc.global
	return dc
}

func functionContext(c *Calc, name string) *Calc {
	dc := &Calc{
		Out:     c.Out,
		name:    c.name + "." + name,
		parent:  c,
		config:  c.config,
		main:    NewStack("main"),
		global:  c.global,
		local:   make(map[string]*Stack),
		funcs:   c.funcs,
		defs:    c.defs,
		modules: c.modules,
	}
	dc.Stack = dc.main
	return dc
}

func (c *Calc) invokeFunction(mod *Calc, fn *ast.FuncNode) error {
	dc := functionContext(mod, fn.Name)
	for _, param := range fn.Params {
		if param.Type == ast.TopRef {
			val, err := c.Stack.Pop()
			if err != nil {
				return fmt.Errorf("not enough arguments, missing '%v'", param.Name)
			}
			c.trace(fn, "func(%v) param %v=%v\n", fn.Name, param.Name, val)
			dc.Define(param.Name).Set(val)
		} else {
			c.trace(fn, "func(%v) param %v=%v", fn.Name, param.Name, c.Stack.Items())
			target := dc.Define(param.Name)
			for c.Stack.Len() > 0 {
				val := c.Stack.MustPop()
				target.Push(val)
			}
		}
	}
	if err := dc.evalBlock(fn.Block); err != nil {
		return err
	}
	for dc.main.Len() > 0 {
		val := dc.main.MustPop()
		c.trace(fn, "func(%v) return %v", fn.Name, val)
		c.Stack.Push(val)
	}
	return nil
}

func (c *Calc) invokeMacro(mac *ast.MacroNode) error {
	if err := c.evalBlock(mac.Expr.Expr); err != nil {
		return err
	}
	return nil
}

func (c *Calc) load(name string) (*Calc, error) {
	def, ok := c.defs[name]
	if !ok {
		return nil, fmt.Errorf("no such module: %v", name)
	}
	mod, ok := c.modules[def.Name]
	if ok {
		return mod, nil
	}

	dc := c.moduleContext(name)
	for name, fn := range def.Natives {
		dc.funcs[name] = fn
	}

	if def.ScriptPath != "" {
		src, err := internal.Scripts.ReadFile(def.ScriptPath)
		if err != nil {
			return nil, err
		}
		ast, err := parser.Parse(def.ScriptPath, src)
		if err != nil {
			return nil, err
		}
		if err := dc.evalNode(ast); err != nil {
			return nil, err
		}
	}

	for name, fn := range builtin {
		dc.funcs[name] = fn
	}

	c.modules[def.Name] = dc
	return dc, nil
}

func (c *Calc) err(node ast.Node, format string, a ...any) error {
	var frames []Frame
	frames = append(frames, Frame{
		Pos:  node.Pos(),
		Func: c.frame.Func,
	})
	for f := c; f != nil; f = f.parent {
		frames = append(frames, Frame{
			Pos:  f.frame.Pos,
			Func: f.frame.Func,
		})
	}
	return CalcError{
		Message: fmt.Sprintf(format, a...),
		Frames:  frames,
	}
}

func (c *Calc) trace(node ast.Node, format string, a ...any) {
	if c.config.Trace {
		msg := fmt.Sprintf(format, a...)
		log.Printf("[%v] eval: %v", node.Pos(), msg)
	}
}
