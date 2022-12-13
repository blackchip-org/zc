package zc

import (
	"fmt"
	"strings"

	"github.com/blackchip-org/zc/lang/ast"
	"github.com/blackchip-org/zc/lang/parser"
)

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
	c.trace(node, "alias %v %v", node.To, node.From)
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
			vb, err := c.ParseBool(v)
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
	c.Info = "ok"
	return nil
}

func (c *Calc) evalNativeNode(node *ast.NativeNode) error {
	c.trace(node, "native %v", strings.Join([]string{node.Name, node.Export}, " "))
	export := node.Export
	if export == "" {
		export = node.Name
	}
	fn, ok := c.Natives[node.Name]
	if !ok {
		return c.err(node, fmt.Errorf("no such native: %v", node.Name))
	}
	c.Funcs[node.Name] = fn
	c.Exports[export] = fn
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
		c.Stack.Push(c.FormatBool(false))
	} else {
		c.Stack.Push(c.FormatBool(true))
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
		c.Stack.Push(c.FormatValue(interp))
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
		Out:      c.Out,
		name:     name,
		config:   c.config,
		main:     NewStack("main"),
		global:   make(map[string]*Stack),
		Funcs:    make(map[string]CalcFunc),
		Exports:  make(map[string]CalcFunc),
		Natives:  make(map[string]CalcFunc),
		defs:     c.defs,
		Modules:  c.Modules,
		Settings: c.Settings,
	}
	dc.global["main"] = dc.main
	dc.Stack = dc.main
	dc.local = dc.global
	return dc
}

func functionContext(c *Calc, node *ast.FuncNode) *Calc {
	dc := &Calc{
		Out:      c.Out,
		name:     c.name + "." + node.Name,
		config:   c.config,
		main:     NewStack("main"),
		global:   c.global,
		local:    make(map[string]*Stack),
		Funcs:    c.Funcs,
		Exports:  c.Exports,
		Natives:  c.Natives,
		defs:     c.defs,
		Modules:  c.Modules,
		Settings: c.Settings,
	}
	dc.local["main"] = dc.main
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
		// TODO: Continue here on error for the case when bootstrapping
		// the prelude itself. Might be a better way to handle this.
		if !ok {
			continue
		}
		for name, fn := range mod.Exports {
			dc.Funcs[name] = fn
		}
	}

	for name, fn := range def.Natives {
		dc.Natives[name] = fn
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
