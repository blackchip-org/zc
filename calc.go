package zc

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strings"

	"github.com/blackchip-org/zc/internal"
	"github.com/blackchip-org/zc/lang/ast"
	"github.com/blackchip-org/zc/lang/parser"
	"github.com/blackchip-org/zc/lang/token"
	"github.com/shopspring/decimal"
)

type NumberFormatOptions struct {
	IntPat  string
	Point   rune
	FracPat string
}

func (n NumberFormatOptions) Separators() map[rune]struct{} {
	seps := make(map[rune]struct{})
	for _, pat := range []string{n.IntPat, n.FracPat} {
		for _, ch := range pat {
			if ch != '0' {
				seps[ch] = struct{}{}
			}
		}
	}
	return seps
}

func DefaultNumberFormatOptions() NumberFormatOptions {
	return NumberFormatOptions{
		IntPat: ",000",
		Point:  '.',
	}
}

var (
	Places       int32  = 16
	RoundMode    string = "half-up"
	NumberFormat        = DefaultNumberFormatOptions()
)

type Config struct {
	ModuleDefs []ModuleDef
	PreludeCLI []string
	PreludeDev []string
	Trace      bool
}

type ModuleDef struct {
	Name       string
	Include    bool
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
	Info    string
	Stack   *Stack
	name    string
	config  Config
	main    *Stack
	global  map[string]*Stack
	local   map[string]*Stack
	Funcs   map[string]CalcFunc
	Exports map[string]CalcFunc
	defs    map[string]ModuleDef
	Modules map[string]*Calc
}

func NewCalc(config Config) (*Calc, error) {
	c := &Calc{
		Out:     &strings.Builder{},
		name:    "<cli>",
		config:  config,
		main:    NewStack("main"),
		global:  make(map[string]*Stack),
		defs:    make(map[string]ModuleDef),
		Modules: make(map[string]*Calc),
		Funcs:   make(map[string]CalcFunc),
		Exports: make(map[string]CalcFunc),
	}
	c.Stack = c.main
	c.local = c.global

	for _, def := range config.ModuleDefs {
		c.Install(def)
	}
	for name, fn := range builtin {
		c.Funcs[name] = fn
	}
	c.Funcs["eval"] = eval

	for _, prelude := range config.PreludeCLI {
		if err := c.Include(prelude); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *Calc) Eval(name string, src []byte) (err error) {
	c.Out.Reset()
	c.Info = ""
	ast, err := parser.Parse(name, src)
	if err != nil {
		return
	}
	err = c.evalNode(ast)

	return
}

func (c *Calc) EvalString(name string, src string) error {
	return c.Eval(name, []byte(src))
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

func (c *Calc) import_(def ModuleDef, prefix string) error {
	dc, err := c.load(def)
	if err != nil {
		return err
	}
	for funcName, fn := range dc.Exports {
		var qName string
		if prefix != "" {
			qName = prefix + "." + funcName
		} else {
			qName = funcName
		}
		c.Funcs[qName] = fn
	}
	return nil
}

func (c *Calc) Import(modName string, alias string) error {
	def, ok := c.defs[modName]
	if !ok {
		return fmt.Errorf("no such module: %v", modName)
	}

	prefix := modName
	if alias != "" {
		prefix = alias
	}
	return c.import_(def, prefix)
}

func (c *Calc) ImportFile(file string, name string) error {
	def := ModuleDef{Name: name, ScriptPath: file}
	return c.import_(def, name)
}

func (c *Calc) Include(modName string) error {
	def, ok := c.defs[modName]
	if !ok {
		return fmt.Errorf("no such module: %v", modName)
	}
	dc, err := c.load(def)
	if err != nil {
		return err
	}
	for funcName, fn := range dc.Exports {
		c.Funcs[funcName] = fn
	}
	return nil
}

func (c *Calc) IncludeFile(file string) error {
	src, err := c.LoadFile(file)
	if err != nil {
		return err
	}

	ast, err := parser.Parse(file, src)
	if err != nil {
		return nil
	}

	return c.evalNode(ast)
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

func (c *Calc) Peek2() (string, string, error) {
	items := c.Stack.Items()
	n := len(items)
	if n < 2 {
		return "", "", fmt.Errorf("%v: stack empty", c.Stack.Name)
	}
	return items[n-2], items[n-1], nil
}

func (c *Calc) Pop2() (string, string, error) {
	b, err := c.Stack.Pop()
	if err != nil {
		return "", "", err
	}
	a, err := c.Stack.Pop()
	if err != nil {
		return "", "", err
	}
	return a, b, nil
}

func (c *Calc) PopBigInt() (*big.Int, error) {
	v, err := c.Stack.Pop()
	if err != nil {
		return nil, err
	}
	r, err := ParseBigInt(v)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Calc) PopBigInt2() (*big.Int, *big.Int, error) {
	b, err := c.PopBigInt()
	if err != nil {
		return nil, nil, err
	}
	a, err := c.PopBigInt()
	if err != nil {
		return nil, nil, err
	}
	return a, b, nil
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

func (c *Calc) PopBool2() (bool, bool, error) {
	b, err := c.PopBool()
	if err != nil {
		return false, false, err
	}
	a, err := c.PopBool()
	if err != nil {
		return false, false, err
	}
	return a, b, nil
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

func (c *Calc) PopDecimal2() (decimal.Decimal, decimal.Decimal, error) {
	b, err := c.PopDecimal()
	if err != nil {
		return decimal.Zero, decimal.Zero, err
	}
	a, err := c.PopDecimal()
	if err != nil {
		return decimal.Zero, decimal.Zero, err
	}
	return a, b, nil
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

func (c *Calc) LoadFile(p string) ([]byte, error) {
	if strings.HasPrefix(p, "zc:") {
		p = p[3:]
		return internal.Files.ReadFile(p)
	}
	return ioutil.ReadFile(p)
}

func (c *Calc) evalBlock(nodes []ast.Node) error {
	for _, node := range nodes {
		if err := c.evalNode(node); err != nil {
			return c.err(node, err)
		}
	}
	return nil
}

func (c *Calc) evalNode(node ast.Node) error {
	switch n := node.(type) {
	case *ast.AliasNode:
		return c.evalAliasNode(n)
	case *ast.ExprNode:
		return c.evalExprNode(n)
	case *ast.IfNode:
		return c.evalIfNode(n)
	case *ast.FileNode:
		return c.evalFileNode(n)
	case *ast.ForNode:
		return c.evalForNode(n)
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
	case *ast.NativeNode:
		return c.evalNativeNode(n)
	case *ast.RefNode:
		return c.evalRefNode(n)
	case *ast.StackNode:
		return c.evalStackNode(n)
	case *ast.TryNode:
		return c.evalTryNode(n)
	case *ast.UseNode:
		return c.evalUseNode(n)
	case *ast.ValueNode:
		return c.evalValueNode(n)
	case *ast.WhileNode:
		return c.evalWhileNode(n)
	}
	panic(fmt.Sprintf("unknown node: %+v", node))
}

func (c *Calc) evalAliasNode(node *ast.AliasNode) error {
	c.trace(node, "alias %v %v", node.From, node.To)
	fn, ok := c.Funcs[node.From]
	if !ok {
		return c.err(node, fmt.Errorf("no such function or macro: %v", node.From))
	}
	c.Funcs[node.To] = fn
	c.Exports[node.To] = fn
	c.Info = "ok"
	return nil
}

func (c *Calc) evalExprNode(expr *ast.ExprNode) error {
	for _, node := range expr.Expr {
		if err := c.evalNode(node); err != nil {
			return c.err(node, err)
		}
	}
	c.Stack = c.main
	return nil
}

func (c *Calc) evalIfNode(ifNode *ast.IfNode) error {
	for _, caseNode := range ifNode.Cases {
		// Final "else" condition will have no case expression
		if caseNode.Cond == nil {
			if err := c.evalBlock(caseNode.Block); err != nil {
				return c.err(caseNode, err)
			}
		} else {
			if err := c.evalExprNode(caseNode.Cond); err != nil {
				return c.err(caseNode.Cond, err)
			}
			v, err := c.Stack.Pop()
			if err != nil {
				return c.err(caseNode.Cond, err)
			}
			vb, err := ParseBool(v)
			if err != nil {
				return c.err(caseNode.Cond, err)
			}
			if vb {
				if err := c.evalBlock(caseNode.Block); err != nil {
					return err
				}
				break
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

func (c *Calc) evalForNode(node *ast.ForNode) error {
	c.trace(node, "for(%v) start", node.Stack.Name)

	expr := NewStack("")
	c.Stack = expr
	if err := c.evalExprNode(node.Expr); err != nil {
		return c.err(node.Expr, err)
	}
	c.Stack = c.main

	i := c.Define(node.Stack.Name)
	for _, item := range expr.Items() {
		c.trace(node, "for(%v) iter: %v", node.Stack.Name, item)
		i.Set(item)
		if err := c.evalBlock(node.Block); err != nil {
			return err
		}
	}
	c.trace(node, "for(%v) end", node.Stack.Name)
	return nil
}

func (c *Calc) evalFuncNode(fn *ast.FuncNode) error {
	c.trace(fn, "define func: %v", fn.Name)
	c.Funcs[fn.Name] = func(ic *Calc) error {
		return ic.invokeFunction(c, fn)
	}
	c.Exports[fn.Name] = c.Funcs[fn.Name]
	return nil
}

func (c *Calc) evalImportNode(node *ast.ImportNode) error {
	mod := node.Module

	if mod.Alias != "" {
		c.trace(node, "import %v %v", mod.Name, mod.Alias)
	} else {
		c.trace(node, "import %v", mod.Name)
	}

	if mod.Zlib {
		if err := c.Import(mod.Name, mod.Alias); err != nil {
			return c.err(node, err)
		}
	} else {
		if err := c.ImportFile(mod.Name, mod.Alias); err != nil {
			return c.err(node, err)
		}
	}
	c.Info = "ok"
	return nil
}

func (c *Calc) evalIncludeNode(node *ast.IncludeNode) error {
	mod := node.Module
	c.trace(node, "include %v", mod.Name)
	if mod.Zlib {
		if err := c.Include(mod.Name); err != nil {
			return c.err(node, err)
		}
	} else {
		if err := c.IncludeFile(mod.Name); err != nil {
			return c.err(node, err)
		}
	}
	c.Info = "ok"
	return nil
}

func (c *Calc) evalInvokeNode(node *ast.InvokeNode) error {
	c.trace(node, "invoke %v", node.Name)
	fn, ok := c.Funcs[node.Name]
	if !ok {
		return c.err(node, fmt.Errorf("no such function: %v", node.Name))
	}
	if err := fn(c); err != nil {
		return c.chain(node, err)
	}
	return nil
}

func (c *Calc) evalMacroNode(mac *ast.MacroNode) error {
	c.trace(mac, "define macro: %v", mac.Name)
	c.Funcs[mac.Name] = func(caller *Calc) error {
		return caller.invokeMacro(mac)
	}
	c.Exports[mac.Name] = c.Funcs[mac.Name]
	return nil
}

func (c *Calc) evalNativeNode(node *ast.NativeNode) error {
	return nil
}

func (c *Calc) evalRefNode(ref *ast.RefNode) error {
	c.trace(ref, "ref %v%v", ref.Type, ref.Name)
	stack, err := c.StackFor(ref.Name)
	if err != nil {
		return c.err(ref, err)
	}

	switch ref.Type {
	case ast.AllRef:
		for _, item := range stack.Items() {
			c.Stack.Push(item)
		}
	case ast.TopRef:
		top, err := stack.Get()
		if err != nil {
			return c.err(ref, err)
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

func (c *Calc) evalTryNode(node *ast.TryNode) error {
	c.trace(node, "try")
	if err := c.evalExprNode(node.Expr); err != nil {
		c.Stack.Push(err.Error())
		c.Stack.Push(FormatBool(false))
	} else {
		c.Stack.Push(FormatBool(true))
	}
	return nil
}

func (c *Calc) evalUseNode(node *ast.UseNode) error {
	c.trace(node, "use %v", node.Name)
	def, ok := c.defs[node.Name]
	if !ok {
		return c.err(node, fmt.Errorf("no such module: %v", node.Name))
	}
	if def.Include {
		if err := c.Include(node.Name); err != nil {
			return c.err(node, err)
		}
		c.Info = "ok, included"
	} else {
		if err := c.Import(node.Name, ""); err != nil {
			return c.err(node, err)
		}
		c.Info = "ok, imported"
	}
	return nil
}

func (c *Calc) evalValueNode(value *ast.ValueNode) error {
	c.trace(value, "value %v", value.Value)
	interp, err := c.Interpolate(value.Value)
	if err != nil {
		return c.err(value, err)
	}
	if interp != value.Value {
		c.trace(value, "interpolate %v", interp)
	}

	if value.IsString {
		c.Stack.Push(interp)
	} else {
		c.Stack.Push(FormatValue(interp))
	}
	return nil
}

func (c *Calc) evalWhileNode(while *ast.WhileNode) error {
	c.trace(while, "while")
	for {
		if err := c.evalExprNode(while.Cond); err != nil {
			return c.err(while.Cond, err)
		}
		result, err := c.PopBool()
		if err != nil {
			return c.err(while.Cond, err)
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
		config:  c.config,
		main:    NewStack("main"),
		global:  make(map[string]*Stack),
		Funcs:   make(map[string]CalcFunc),
		Exports: make(map[string]CalcFunc),
		defs:    c.defs,
		Modules: c.Modules,
	}
	dc.Stack = dc.main
	dc.local = dc.global
	return dc
}

func functionContext(c *Calc, node *ast.FuncNode) *Calc {
	dc := &Calc{
		Out:     c.Out,
		name:    c.name + "." + node.Name,
		config:  c.config,
		main:    NewStack("main"),
		global:  c.global,
		local:   make(map[string]*Stack),
		Funcs:   c.Funcs,
		Exports: c.Exports,
		defs:    c.defs,
		Modules: c.Modules,
	}
	dc.Stack = dc.main
	return dc
}

func (c *Calc) invokeFunction(mod *Calc, fn *ast.FuncNode) error {
	dc := functionContext(mod, fn)
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
	c.trace(fn, "func(%v) end", fn.Name)
	return nil
}

func (c *Calc) invokeMacro(mac *ast.MacroNode) error {
	if err := c.evalBlock(mac.Expr.Expr); err != nil {
		return err
	}
	return nil
}

func (c *Calc) load(def ModuleDef) (*Calc, error) {
	dc := c.moduleContext(def.Name)

	for name, fn := range builtin {
		dc.Funcs[name] = fn
	}
	dc.Funcs["eval"] = eval

	for _, prelude := range c.config.PreludeDev {
		mod, ok := c.Modules[prelude]
		if !ok {
			continue
		}
		for name, fn := range mod.Exports {
			dc.Funcs[name] = fn
		}
	}

	for name, fn := range def.Natives {
		dc.Funcs[name] = fn
		dc.Exports[name] = fn
	}

	if def.ScriptPath != "" {
		src, err := dc.LoadFile(def.ScriptPath)
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

	c.Modules[def.Name] = dc
	return dc, nil
}

func (c *Calc) chain(node *ast.InvokeNode, err error) error {
	frame := Frame{Pos: node.Pos()}

	errCalc, ok := err.(CalcError)
	if ok {
		if len(errCalc.Frames) > 0 {
			errCalc.Frames[len(errCalc.Frames)-1].Func = node.Name
		}
		errCalc.Frames = append(errCalc.Frames, frame)
		return errCalc
	}
	return CalcError{
		Message: err.Error(),
		Frames:  []Frame{frame},
	}
}

func (c *Calc) err(node ast.Node, err error) error {
	errCalc, ok := err.(CalcError)
	if ok {
		return errCalc
	}
	frame := Frame{Pos: node.Pos()}
	return CalcError{
		Message: err.Error(),
		Frames:  []Frame{frame},
	}
}

func (c *Calc) trace(node ast.Node, format string, a ...any) {
	if c.config.Trace {
		msg := fmt.Sprintf(format, a...)
		if c.Stack.Len() > 0 {
			log.Printf("eval: %v(%v)", c.Stack.Name, c.Stack)
		}
		log.Printf("eval:     %v @ %v", msg, node.Pos())
		//log.Println()
	}
}
