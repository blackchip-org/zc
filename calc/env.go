package calc

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/coll"
	"github.com/blackchip-org/zc/lang/parser"
)

type invokeFunc func(*env) error

type env struct {
	parent    *env
	name      string
	calc      *calc
	stack     coll.Deque[string]
	stackName string
	stacks    map[string]coll.Deque[string]
	funcs     map[string]invokeFunc
	mod       *zc.ModuleDef
	exports   []zc.FuncDecl
}

func newEnv(calc *calc, name string) *env {
	e := &env{
		calc:   calc,
		name:   name,
		stacks: make(map[string]coll.Deque[string]),
	}
	e.NewStack(zc.MainStack)
	e.UseStack(zc.MainStack)
	return e
}

func (e *env) Calc() zc.Calc {
	return e.calc
}

func (e *env) Stack() coll.Deque[string] {
	return e.stack
}

func (e *env) NewStack(name string) bool {
	_, exists := e.stacks[name]
	if exists {
		return false
	}
	e.stacks[name] = coll.NewDequeSlice[string]()
	return true
}

func (e *env) UseStack(name string) bool {
	s, ok := e.stacks[name]
	if !ok {
		return false
	}
	e.stack = s
	e.stackName = name
	return true
}

func (e *env) Eval(name string, src []byte) error {
	root, err := parser.Parse(name, src)
	if err != nil {
		return err
	}
	coll.Push(e.calc.frames, frame{
		pos:      root.Pos(),
		funcName: "<eval>",
		env:      e,
	})
	err = evalFile(e, root)
	coll.Pop(e.calc.frames)
	return err
}

func (e *env) derive(name string) *env {
	de := newEnv(e.calc, e.name+"."+name)
	de.parent = e
	return de
}

func (e *env) lookupFunc(name string) (invokeFunc, bool) {
	fn, ok := e.funcs[name]
	if ok {
		return fn, true
	}
	if e.parent == nil {
		return nil, false
	}
	return e.parent.lookupFunc(name)
}
