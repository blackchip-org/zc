package zc

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/blackchip-org/zc/lang/ast"
	"github.com/blackchip-org/zc/lang/token"
)

var (
	// Not an actual error, used to return from a function
	errFuncReturn = errors.New("function return")
)

func (e *Env) evalAliasStmt(node *ast.AliasStmt) error {
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

func (e *Env) evalAtom(atom ast.Atom) error {
	switch a := atom.(type) {
	case *ast.InvokeAtom:
		return e.evalInvokeAtom(a)
	case *ast.NumberAtom:
		return e.evalNumberAtom(a)
	case *ast.RefAtom:
		return e.evalRefAtom(a)
	case *ast.SelectAtom:
		return e.evalSelectAtom(a)
	case *ast.ValueAtom:
		return e.evalValueAtom(a)
	}
	panic(fmt.Sprintf("unknown atom: %+v", atom))
}

func (e *Env) evalExpr(expr *ast.Expr) error {
	for _, atom := range expr.Atoms {
		if err := e.evalAtom(atom); err != nil {
			return e.err(atom, err)
		}
	}
	return nil
}

func (e *Env) evalExprStmt(stmt *ast.ExprStmt) error {
	err := e.evalExpr(stmt.Expr)
	e.Stack = e.Main
	e.traceStack()
	return err
}

func (e *Env) evalFile(file *ast.File) error {
	for _, line := range file.Stmts {
		if err := e.evalStmt(line); err != nil {
			return err
		}
	}
	return nil
}

func (e *Env) evalForStmt(node *ast.ForStmt) error {
	e.trace(node, "for-begin(%v)", node.Stack.Name)

	expr := NewStack(e.Calc, "")
	e.Stack = expr
	err := e.evalExpr(node.Expr)
	e.Stack = e.Main
	if err != nil {
		return e.err(node.Expr, err)
	}

	for _, item := range expr.Items() {
		e.trace(node, "for(%v) iter: %v", node.Stack.Name, item)
		inner := e.DeriveBlock("for")
		i := inner.NewStack(node.Stack.Name)
		i.Clear().Push(item)
		if err := inner.evalStmts(node.Stmts); err != nil {
			return err
		}
		e.trace(node, "for-next(%v)", node.Stack.Name)
	}
	e.trace(node, "for-end(%v)", node.Stack.Name)
	return nil
}

func (e *Env) evalFuncStmt(fn *ast.FuncStmt) error {
	e.trace(fn, "define func: %v", fn.Name)
	e.Funcs[fn.Name] = func(caller *Env) error {
		return e.invokeFunction(caller, fn)
	}
	if e.Module != "" {
		e.Exports = append(e.Exports, fn.Name)
	}
	return nil
}

func (e *Env) evalIfStmt(ifNode *ast.IfStmt) error {
	e.trace(ifNode, "if")
	for _, caseNode := range ifNode.Cases {
		// Final "else" condition will have no case expression
		if caseNode.Cond == nil {
			e.trace(caseNode, "else")
			if err := e.evalStmts(caseNode.Stmts); err != nil {
				return e.err(caseNode, err)
			}
		} else {
			if err := e.evalExpr(caseNode.Cond); err != nil {
				return e.err(caseNode.Cond, err)
			}
			e.traceStack()
			v, err := e.Stack.PopBool()
			e.trace(caseNode, "case == %v", v)
			if err != nil {
				return e.err(caseNode.Cond, err)
			}
			if v {
				if err := e.evalStmts(caseNode.Stmts); err != nil {
					return err
				}
				break
			}
		}
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
	return nil
}

func (e *Env) evalImportStmt(node *ast.ImportStmt) error {
	var names []string
	for _, mod := range node.Modules {
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
		names = append(names, mod.Name)
	}
	e.Calc.Info = "imported " + strings.Join(names, ", ")
	return nil
}

func (e *Env) evalInvokeAtom(node *ast.InvokeAtom) error {
	e.trace(node, "invoke %v", node.Name)
	fn, ok := e.Func(node.Name)
	if !ok {
		return e.err(node, fmt.Errorf("no such function: %v", node.Name))
	}
	if err := fn(e); err != nil {
		if errors.Is(err, errFuncReturn) {
			return err
		}
		return e.chain(node.Name, node.Pos(), err)
	}
	return nil
}

func (e *Env) evalMacroStmt(mac *ast.MacroStmt) error {
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

func (e *Env) evalNativeStmt(node *ast.NativeStmt) error {
	e.trace(node, "native %v", strings.Join([]string{node.Name, node.Export}, " "))
	export := node.Export
	if export == "" {
		export = node.Name
	}
	fn, ok := e.Calc.Natives[node.Name]
	if !ok {
		return e.err(node, fmt.Errorf("no such native: %v", node.Name))
	}
	e.Funcs[export] = func(caller *Env) error {
		return e.invokeNative(caller, fn, node)
	}
	if e.Module != "" {
		e.Exports = append(e.Exports, node.Name)
	}
	return nil
}

func (e *Env) evalNumberAtom(number *ast.NumberAtom) error {
	e.trace(number, "number %v", number.Value)
	e.Stack.PushValue(e.Calc.LocalizeNumber(number.Value))
	e.traceStack()
	return nil
}

func (e *Env) evalRefAtom(ref *ast.RefAtom) error {
	e.trace(ref, "ref %v%v", ref.Type, ref.Name)
	stack, ok := e.StackFor(ref.Name)
	if !ok {
		return e.err(ref, fmt.Errorf("no such stack: %v", ref.Name))
	}

	switch ref.Type {
	case ast.AllRef:
		for _, item := range stack.ItemsReversed() {
			e.Stack.Push(item)
		}
	case ast.TopRef:
		top, err := stack.Peek()
		if err != nil {
			return e.err(ref, err)
		}
		e.Stack.Push(top)
	case ast.PopRef:
		top, err := stack.Pop()
		if err != nil {
			return e.err(ref, err)
		}
		e.Stack.Push(top)
	}
	e.traceStack()
	return nil
}

func (e *Env) evalReturnStmt(node *ast.ReturnStmt) error {
	e.trace(node, "return")
	return errFuncReturn
}

func (e *Env) evalSelectAtom(node *ast.SelectAtom) error {
	e.trace(node, "select %v", node.Name)
	stack, ok := e.StackFor(node.Name)
	if !ok {
		stack = e.NewStack(node.Name)
	}
	e.Stack = stack
	e.traceStack()
	return nil
}

func (e *Env) evalStmt(stmt ast.Stmt) error {
	switch s := stmt.(type) {
	case *ast.AliasStmt:
		return e.evalAliasStmt(s)
	case *ast.ExprStmt:
		return e.evalExprStmt(s)
	case *ast.IfStmt:
		return e.evalIfStmt(s)
	case *ast.ForStmt:
		return e.evalForStmt(s)
	case *ast.FuncStmt:
		return e.evalFuncStmt(s)
	case *ast.ImportStmt:
		return e.evalImportStmt(s)
	case *ast.MacroStmt:
		return e.evalMacroStmt(s)
	case *ast.NativeStmt:
		return e.evalNativeStmt(s)
	case *ast.ReturnStmt:
		return e.evalReturnStmt(s)
	case *ast.TryStmt:
		return e.evalTryStmt(s)
	case *ast.UseStmt:
		return e.evalUseStmt(s)
	case *ast.WhileStmt:
		return e.evalWhileStmt(s)
	}
	panic(fmt.Sprintf("unknown stmt: %+v", stmt))
}

func (e *Env) evalStmts(stmts []ast.Stmt) error {
	for _, stmt := range stmts {
		if err := e.evalStmt(stmt); err != nil {
			return e.err(stmt, err)
		}
	}
	return nil
}

func (e *Env) evalTryStmt(node *ast.TryStmt) error {
	e.trace(node, "try")
	if err := e.evalExpr(node.Expr); err != nil {
		e.Stack.Push(err.Error())
		e.Stack.PushBool(false)
	} else {
		e.Stack.PushBool(true)
	}
	e.traceStack()
	return nil
}

func (e *Env) evalUseStmt(node *ast.UseStmt) error {
	var names []string
	for _, mod := range node.Modules {
		e.trace(node, "use %v", mod.Name)
		if err := e.evalImport(mod.Name, "", mod.Zlib); err != nil {
			return e.err(node, err)
		}
		names = append(names, mod.Name)
	}
	e.Calc.Info = "using " + strings.Join(names, ", ")
	return nil
}

func (e *Env) evalValueAtom(value *ast.ValueAtom) error {
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
	e.traceStack()
	return nil
}

func (e *Env) evalWhileStmt(while *ast.WhileStmt) error {
	e.trace(while, "while-begin")
	for {

		de := e.DeriveBlock("while")
		if err := de.evalExpr(while.Cond); err != nil {
			return e.err(while.Cond, err)
		}
		result, err := de.Stack.PopBool()
		if err != nil {
			return e.err(while.Cond, err)
		}
		if !result {
			break
		}
		if err := de.evalStmts(while.Stmts); err != nil {
			return err
		}
		e.trace(while, "while-next")
	}
	e.trace(while, "while-end")
	return nil
}

func (e *Env) invokeFunction(caller *Env, fn *ast.FuncStmt) error {
	// callee := e
	callee := e.Derive(fn.Name)
	for _, param := range fn.Params {
		if param.Type == ast.TopRef {
			val, err := caller.Stack.Pop()
			if err != nil {
				return fmt.Errorf("not enough arguments for '%v', missing '%v'", fn.Name, param.Name)
			}
			e.trace(fn, "func(%v) param %v=%v\n", fn.Name, param.Name, val)
			callee.NewStack(param.Name).Push(val)
		} else if param.Type == ast.AllRef {
			e.trace(fn, "func(%v) param %v=(%v)", fn.Name, param.Name, caller.Stack.String())
			target := callee.NewStack(param.Name)
			for caller.Stack.Len() > 0 {
				val := caller.Stack.MustPop()
				//target.Push(val)
				target.Enqueue(val)
			}
		} else {
			return fmt.Errorf("stack reference %v not allowed as parameter", param.Type)
		}
	}
	if err := callee.evalStmts(fn.Stmts); err != nil {
		if !errors.Is(err, errFuncReturn) {
			return err
		}
	}
	returns := NewStack(e.Calc, "<return>")
	for callee.Main.Len() > 0 {
		val := callee.Main.MustPop()
		e.trace(fn, "func(%v) return %v", fn.Name, val)
		returns.Push(val)
	}
	for returns.Len() != 0 {
		v := returns.MustPop()
		caller.Stack.Push(v)
	}
	e.trace(fn, "func(%v) end", fn.Name)
	e.traceStack()
	return nil
}

func (e *Env) invokeNative(caller *Env, fn CalcFunc, stmt *ast.NativeStmt) error {
	callee := e
	callee.Stack = caller.Stack
	if err := fn(callee); err != nil {
		if !errors.Is(err, errFuncReturn) {
			return err
		}
	}
	e.trace(stmt, "func(%v) end", stmt.Name)
	return nil
}

func (e *Env) invokeMacro(mac *ast.MacroStmt) error {
	if err := e.evalExpr(mac.Expr); err != nil {
		return err
	}
	return nil
}

func (e *Env) chain(name string, pos token.Pos, err error) error {
	frame := Frame{Pos: pos}

	errCalc, ok := err.(CalcError)
	if ok {
		if len(errCalc.Frames) > 0 {
			errCalc.Frames[len(errCalc.Frames)-1].Func = name
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
	if errors.Is(err, errFuncReturn) {
		return err
	}
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

func (e *Env) traceStack() {
	if e.Calc.Trace {
		if !e.Stack.Equal(e.lastStack) && e.Stack.Len() > 0 {
			log.Printf("eval: %v(%v)  %v", e.Stack.Name, e.Stack, e.Name)
		}
		e.lastStack = e.Stack.Copy()
	}
}

func (e *Env) trace(node ast.Node, format string, a ...any) {
	if e.Calc.Trace {
		msg := fmt.Sprintf(format, a...)
		log.Printf("eval:     %v @ %v", msg, node.Pos())
	}
}
