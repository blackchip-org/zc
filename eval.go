package zc

import (
	"fmt"
	"log"
	"strings"

	"github.com/blackchip-org/zc/lang/ast"
)

func (e *Env) evalBlock(nodes []ast.Node) error {
	for _, node := range nodes {
		if err := e.evalNode(node); err != nil {
			return e.err(node, err)
		}
	}
	return nil
}

func (e *Env) evalNode(node ast.Node) error {
	switch n := node.(type) {
	case *ast.AliasNode:
		return e.evalAliasNode(n)
	case *ast.ExprNode:
		return e.evalExprNode(n)
	case *ast.IfNode:
		return e.evalIfNode(n)
	case *ast.FileNode:
		return e.evalFileNode(n)
	case *ast.ForNode:
		return e.evalForNode(n)
	case *ast.FuncNode:
		return e.evalFuncNode(n)
	case *ast.ImportNode:
		return e.evalImportNode(n)
	case *ast.IncludeNode:
		return e.evalIncludeNode(n)
	case *ast.InvokeNode:
		return e.evalInvokeNode(n)
	case *ast.MacroNode:
		return e.evalMacroNode(n)
	case *ast.NativeNode:
		return e.evalNativeNode(n)
	case *ast.RefNode:
		return e.evalRefNode(n)
	case *ast.StackNode:
		return e.evalStackNode(n)
	case *ast.TryNode:
		return e.evalTryNode(n)
	case *ast.UseNode:
		return e.evalUseNode(n)
	case *ast.ValueNode:
		return e.evalValueNode(n)
	case *ast.WhileNode:
		return e.evalWhileNode(n)
	}
	panic(fmt.Sprintf("unknown node: %+v", node))
}

func (e *Env) evalAliasNode(node *ast.AliasNode) error {
	e.trace(node, "alias %v %v", node.To, node.From)
	fn, ok := e.Func(node.From)
	if !ok {
		return e.err(node, fmt.Errorf("no such function or macro: %v", node.From))
	}
	e.Funcs[node.To] = fn
	if e.Module != "" {
		e.Exports = append(e.Exports, node.To)
	}
	e.Calc.Info = "ok"
	return nil
}

func (e *Env) evalExprNode(expr *ast.ExprNode) error {
	for _, node := range expr.Expr {
		if err := e.evalNode(node); err != nil {
			return e.err(node, err)
		}
	}
	e.Stack = e.Main
	return nil
}

func (e *Env) evalIfNode(ifNode *ast.IfNode) error {
	for _, caseNode := range ifNode.Cases {
		// Final "else" condition will have no case expression
		if caseNode.Cond == nil {
			if err := e.evalBlock(caseNode.Block); err != nil {
				return e.err(caseNode, err)
			}
		} else {
			if err := e.evalExprNode(caseNode.Cond); err != nil {
				return e.err(caseNode.Cond, err)
			}
			v, err := e.Stack.Pop()
			if err != nil {
				return e.err(caseNode.Cond, err)
			}
			vb, err := e.Calc.ParseBool(v)
			if err != nil {
				return e.err(caseNode.Cond, err)
			}
			if vb {
				if err := e.evalBlock(caseNode.Block); err != nil {
					return err
				}
				break
			}
		}
	}
	return nil
}

func (e *Env) evalFileNode(file *ast.FileNode) error {
	for _, line := range file.Block {
		if err := e.evalNode(line); err != nil {
			return err
		}
	}
	return nil
}

func (e *Env) evalForNode(node *ast.ForNode) error {
	e.trace(node, "for(%v) start", node.Stack.Name)

	expr := NewStack(e.Calc, "")
	e.Stack = expr
	err := e.evalExprNode(node.Expr)
	e.Stack = e.Main
	if err != nil {
		return e.err(node.Expr, err)
	}

	inner := e.Derive()
	i := inner.NewStack(node.Stack.Name)
	for _, item := range expr.Items() {
		e.trace(node, "for(%v) iter: %v", node.Stack.Name, item)
		i.Clear().Push(item)
		if err := inner.evalBlock(node.Block); err != nil {
			return err
		}
	}
	e.trace(node, "for(%v) end", node.Stack.Name)
	return nil
}

func (e *Env) evalFuncNode(fn *ast.FuncNode) error {
	e.trace(fn, "define func: %v", fn.Name)
	e.Funcs[fn.Name] = func(caller *Env) error {
		return e.invokeFunction(caller, fn)
	}
	if e.Module != "" {
		e.Exports = append(e.Exports, fn.Name)
	}
	return nil
}

func (e *Env) evalImport(name string, alias string, zlib bool) error {
	var def ModuleDef
	var ok bool
	if zlib {
		def, ok = e.Calc.ModuleDefs[name]
		if !ok {
			return fmt.Errorf("no such module: %v", name)
		}
	} else {
		def = ModuleDef{Name: name, ScriptPath: name}
	}
	if _, err := e.Import(def, alias); err != nil {
		return err
	}
	e.Calc.Info = "ok"
	return nil

}

func (e *Env) evalImportNode(node *ast.ImportNode) error {
	mod := node.Module
	alias := mod.Alias
	if alias != "" {
		e.trace(node, "import %v %v", mod.Name, alias)
	} else {
		e.trace(node, "import %v", mod.Name)
		alias = mod.Name
	}
	if err := e.evalImport(mod.Name, alias, mod.Zlib); err != nil {
		return e.err(node, err)
	}
	return nil
}

func (e *Env) evalIncludeNode(node *ast.IncludeNode) error {
	mod := node.Module
	e.trace(node, "include %v", mod.Name)
	if err := e.evalImport(mod.Name, "", mod.Zlib); err != nil {
		return e.err(node, err)
	}
	return nil
}

func (e *Env) evalInvokeNode(node *ast.InvokeNode) error {
	e.trace(node, "invoke %v", node.Name)
	fn, ok := e.Func(node.Name)
	if !ok {
		return e.err(node, fmt.Errorf("no such function: %v", node.Name))
	}
	if err := fn(e); err != nil {
		return e.chain(node, err)
	}
	return nil
}

func (e *Env) evalMacroNode(mac *ast.MacroNode) error {
	e.trace(mac, "define macro: %v", mac.Name)
	e.Funcs[mac.Name] = func(caller *Env) error {
		return caller.invokeMacro(mac)
	}
	if e.Module != "" {
		e.Exports = append(e.Exports, mac.Name)
	}
	e.Calc.Info = "ok"
	return nil
}

func (e *Env) evalNativeNode(node *ast.NativeNode) error {
	e.trace(node, "native %v", strings.Join([]string{node.Name, node.Export}, " "))
	export := node.Export
	if export == "" {
		export = node.Name
	}
	fn, ok := e.Calc.Natives[node.Name]
	if !ok {
		return e.err(node, fmt.Errorf("no such native: %v", node.Name))
	}
	e.Funcs[export] = fn
	if e.Module != "" {
		e.Exports = append(e.Exports, node.Name)
	}
	return nil
}

func (e *Env) evalRefNode(ref *ast.RefNode) error {
	e.trace(ref, "ref %v%v", ref.Type, ref.Name)
	stack, ok := e.StackFor(ref.Name)
	if !ok {
		return e.err(ref, fmt.Errorf("no such stack: %v", ref.Name))
	}

	switch ref.Type {
	case ast.AllRef:
		for _, item := range stack.Items() {
			e.Stack.Push(item)
		}
	case ast.TopRef:
		top, err := stack.Peek()
		if err != nil {
			return e.err(ref, err)
		}
		e.Stack.Push(top)
	}
	return nil
}

func (e *Env) evalStackNode(node *ast.StackNode) error {
	e.trace(node, "stack %v", node.Name)
	stack, ok := e.StackFor(node.Name)
	if !ok {
		stack = e.NewStack(node.Name)
	}
	e.Stack = stack
	return nil
}

func (e *Env) evalTryNode(node *ast.TryNode) error {
	e.trace(node, "try")
	if err := e.evalExprNode(node.Expr); err != nil {
		e.Stack.Push(err.Error())
		e.Stack.PushBool(false)
	} else {
		e.Stack.PushBool(true)
	}
	return nil
}

func (e *Env) evalUseNode(node *ast.UseNode) error {
	e.trace(node, "use %v", node.Name)
	def, ok := e.Calc.ModuleDefs[node.Name]
	if !ok {
		return e.err(node, fmt.Errorf("no such module: %v", node.Name))
	}
	alias := ""
	msg := "ok, imported"
	if !def.Include {
		alias = def.Name
		msg = "ok, included"
	}
	if _, err := e.Import(def, alias); err != nil {
		return e.err(node, err)
	}
	e.Calc.Info = msg
	return nil
}

func (e *Env) evalValueNode(value *ast.ValueNode) error {
	e.trace(value, "value %v", value.Value)
	interp, err := e.Interpolate(value.Value)
	if err != nil {
		return e.err(value, err)
	}
	if interp != value.Value {
		e.trace(value, "interpolate %v", interp)
	}

	if value.IsString {
		e.Stack.Push(interp)
	} else {
		e.Stack.PushValue(interp)
	}
	return nil
}

func (e *Env) evalWhileNode(while *ast.WhileNode) error {
	e.trace(while, "while")
	for {
		if err := e.evalExprNode(while.Cond); err != nil {
			return e.err(while.Cond, err)
		}
		result, err := e.Stack.PopBool()
		if err != nil {
			return e.err(while.Cond, err)
		}
		if !result {
			break
		}
		if err := e.evalBlock(while.Block); err != nil {
			return err
		}
	}
	return nil
}

func (e *Env) invokeFunction(caller *Env, fn *ast.FuncNode) error {
	callee := e
	for _, param := range fn.Params {
		if param.Type == ast.TopRef {
			val, err := caller.Stack.Pop()
			if err != nil {
				return fmt.Errorf("not enough arguments, missing '%v'", param.Name)
			}
			e.trace(fn, "func(%v) param %v=%v\n", fn.Name, param.Name, val)
			callee.NewStack(param.Name).Push(val)
		} else {
			e.trace(fn, "func(%v) param %v=%v", fn.Name, param.Name, e.Stack.Items())
			target := callee.NewStack(param.Name)
			for caller.Stack.Len() > 0 {
				val := caller.Stack.MustPop()
				target.Push(val)
			}
		}
	}
	if err := callee.evalBlock(fn.Block); err != nil {
		return err
	}
	for callee.Main.Len() > 0 {
		val := callee.Main.MustPop()
		e.trace(fn, "func(%v) return %v", fn.Name, val)
		caller.Stack.Push(val)
	}
	e.trace(fn, "func(%v) end", fn.Name)
	return nil
}

func (e *Env) invokeMacro(mac *ast.MacroNode) error {
	if err := e.evalBlock(mac.Expr.Expr); err != nil {
		return err
	}
	return nil
}

func (e *Env) chain(node *ast.InvokeNode, err error) error {
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

func (e *Env) err(node ast.Node, err error) error {
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

func (e *Env) trace(node ast.Node, format string, a ...any) {
	if e.Calc.Trace {
		msg := fmt.Sprintf(format, a...)
		if e.Stack.Len() > 0 {
			log.Printf("eval: %v(%v)", e.Stack.Name, e.Stack)
		}
		log.Printf("eval:     %v @ %v", msg, node.Pos())
		//log.Println()
	}
}
