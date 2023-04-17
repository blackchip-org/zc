package calc

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/coll"
	"github.com/blackchip-org/zc/lang/ast"
	"github.com/blackchip-org/zc/types"
)

var (
	// Not an actual error, used to return from a function
	errFuncReturn = errors.New("function return")
)

func evalFile(e *env, file *ast.File) error {
	for _, line := range file.Stmts {
		if err := evalStmt(e, line); err != nil {
			return err
		}
	}
	return nil
}

func evalStmt(e *env, stmt ast.Stmt) error {
	switch s := stmt.(type) {
	// case *ast.AliasStmt:
	// 	return e.evalAliasStmt(s)
	case *ast.ExprStmt:
		return evalExprStmt(e, s)
	// case *ast.IfStmt:
	// 	return e.evalIfStmt(s)
	// case *ast.ForStmt:
	// 	return e.evalForStmt(s)
	// case *ast.FuncStmt:
	// 	return e.evalFuncStmt(s)
	// case *ast.ImportStmt:
	// 	return e.evalImportStmt(s)
	case *ast.IncludeStmt:
		return evalIncludeStmt(e, s)
		// case *ast.MacroStmt:
		// 	return e.evalMacroStmt(s)
	case *ast.NativeStmt:
		return evalNativeStmt(e, s)
		// case *ast.ReturnStmt:
		// 	return e.evalReturnStmt(s)
		// case *ast.TryStmt:
		// 	return e.evalTryStmt(s)
		// case *ast.UseStmt:
		// 	return e.evalUseStmt(s)
		// case *ast.WhileStmt:
		// 	return e.evalWhileStmt(s)
	}
	panic(fmt.Sprintf("unknown stmt: %+v", stmt))
}

func evalNativeStmt(e *env, node *ast.NativeStmt) error {
	evalTrace(e, node, "native %v", strings.Join([]string{node.Name, node.Export}, " "))
	export := node.Export
	if export == "" {
		export = node.Name
	}
	decl, ok := e.mod.Natives[node.Name]
	if !ok {
		return evalErr(e, node, fmt.Errorf("no such native: %v", node.Name))
	}
	e.funcs[export] = func(caller *env) error {
		return invoke(e, caller, decl, node)
	}
	if e.mod != nil {
		e.exports = append(e.exports, decl)
	}
	return nil
}

func evalExprStmt(e *env, stmt *ast.ExprStmt) error {
	err := evalExpr(e, stmt.Expr)
	e.UseStack(zc.MainStack)
	//e.traceStack()
	return err
}

func evalIncludeStmt(e *env, node *ast.IncludeStmt) error {
	var names []string
	for _, ref := range node.Modules {
		evalTrace(e, node, "include %v", ref.Name)
		mod, err := load(e.calc, ref.Name)
		if err != nil {
			return err
		}
		for _, decl := range mod.exports {

		}
	}
	e.calc.SetInfo("included %s", strings.Join(names, ", "))
	return nil
}

func evalValueAtom(e *env, node *ast.ValueAtom) error {
	evalTrace(e, node, "value %v", node.Value)
	val := node.Value
	// if !node.IsPlain {
	// 	interp, err := e.Interpolate(val)
	// 	if err != nil {
	// 		return e.err(node, err)
	// 	}
	// 	if interp != val {
	// 		e.trace(node, "interpolate %v", interp)
	// 	}
	// 	val = interp
	// }

	// if node.IsString {
	// 	e.Stack.Push(val)
	// } else {
	// 	e.Stack.PushValue(val)
	// }
	coll.Push(e.stack, val)
	//e.traceStack()
	return nil
}

// -----

func load(c *calc, name string) (*env, error) {
	if env, exists := c.modules[name]; exists {
		return env, nil
	}
	root, ok := c.lib.Parse(name)
	if !ok {
		return nil, fmt.Errorf("no such module: %v", name)
	}
	env := newEnv(c, name)
	if err := evalFile(env, root); err != nil {
		return nil, err
	}
	c.modules[name] = env
	return env, nil
}

func invoke(callee *env, caller *env, decl zc.FuncDecl, node *ast.NativeStmt) error {
	de := callee.derive(node.Name)
	var args []types.Value
	for _, param := range decl.Params {
		str, ok := coll.Pop(de.stack)
		if !ok {
			return errStackEmpty(de.stackName)
		}
		arg, err := param.ParseValue(str)
		if err != nil {
			return err
		}
		args = append(args, arg)
	}
	results, err := decl.Func(de, args)
	if err != nil {
		return err
	}
	for _, result := range results {
		coll.Push(callee.stack, result.String())
	}
	return nil
}

func (e *env) evalAtom(atom ast.Atom) error {
	switch a := atom.(type) {
	case *ast.InvokeAtom:
		return evalInvokeAtom(e, a)
		/*
			case *ast.RefAtom:
					return e.evalRefAtom(a)
			case *ast.SelectAtom:
					return e.evalSelectAtom(a)
		*/
	case *ast.ValueAtom:
		return e.evalValueAtom(a)
	}
	panic(fmt.Sprintf("unknown atom: %+v", atom))
}

func evalExpr(e *env, expr *ast.Expr) error {
	for _, atom := range expr.Atoms {
		if err := e.evalAtom(atom); err != nil {
			return evalErr(e, atom, err)
		}
	}
	return nil
}

func evalInvokeAtom(e *env, node *ast.InvokeAtom) error {
	evalTrace(e, node, "invoke %v", node.Name)
	fn, ok := e.lookupFunc(node.Name)
	if !ok {
		return evalErr(e, node, fmt.Errorf("no such function: %v", node.Name))
	}

	if err := fn(e); err != nil {
		if errors.Is(err, errFuncReturn) {
			return err
		}
		return evalErr(e, node, err)
		//return e.chain(node.Name, node.Pos(), err)
	}
	return nil
}

/*
func evalImport(e *env, name string, alias string) error {

	if zlib {
		def, ok = e.Calc.ModuleDef(name)
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
*/

// ----

func evalErr(e *env, node ast.Node, err error) error {
	if errors.Is(err, errFuncReturn) {
		return err
	}
	errCalc, ok := err.(zc.CalcError)
	if ok {
		return errCalc
	}
	frame := frame{pos: node.Pos()}
	return zc.CalcError{
		Message: err.Error(),
		Frames:  []zc.Frame{frame},
	}
}

func errStackEmpty(stackName string) error {
	return fmt.Errorf("%v: stack empty", stackName)
}

/*
func chain(e *env, name string, pos scanner.Pos, err error) error {
	frame := frame{pos: pos}

	errCalc, ok := err.(zc.CalcError)
	if ok {
		if len(errCalc.Frames) > 0 {
			errCalc.Frames[len(errCalc.Frames)-1].Func = name
		}
		errCalc.Frames = append(errCalc.Frames, frame)
		return errCalc
	}
	return zc.CalcError{
		Message: err.Error(),
		Frames:  []Frame{frame},
	}
}
*/

func evalTrace(e *env, node ast.Node, format string, a ...any) {
	if e.Calc().Trace() {
		msg := fmt.Sprintf(format, a...)
		log.Printf("eval:     %v @ %v", msg, node.Pos())
	}
}
