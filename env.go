package zc

import (
	"fmt"
	"strings"

	"github.com/blackchip-org/zc/types"
)

type Env struct {
	parent    *Env
	Calc      *Calc
	Name      string
	Stack     *Stack
	Main      *Stack
	stacks    map[string]*Stack
	Funcs     map[string]CalcFunc
	Exports   []string
	Module    string
	lastStack *Stack // for tracing
}

func NewEnv(calc *Calc, name string) *Env {
	e := &Env{
		Calc:   calc,
		Name:   name,
		stacks: make(map[string]*Stack),
		Funcs:  make(map[string]CalcFunc),
	}
	e.SetMain(NewStack("main"))
	return e
}

func (e *Env) Derive(name string) *Env {
	de := NewEnv(e.Calc, e.Name+"."+name)
	de.parent = e
	return de
}

func (e *Env) DeriveBlock(name string) *Env {
	de := &Env{
		Calc:   e.Calc,
		Name:   e.Name + "." + name,
		stacks: e.stacks,
		Funcs:  e.Funcs,
	}
	de.Main = e.Main
	de.Stack = de.Main
	de.parent = e.parent
	return de
}

func (e *Env) SetMain(s *Stack) {
	e.Main = s
	e.Stack = s
	e.stacks["main"] = s
}

func (e *Env) StackFor(name string) (*Stack, bool) {
	s, ok := e.stacks[name]
	if ok {
		return s, true
	}
	if e.parent == nil {
		return nil, false
	}
	return e.parent.StackFor(name)
}

func (e *Env) NewStack(name string) *Stack {
	s := NewStack(name)
	e.stacks[name] = s
	return s
}

func (e *Env) Func(name string) (CalcFunc, bool) {
	fn, ok := e.Funcs[name]
	if ok {
		return fn, true
	}
	if e.parent == nil {
		return nil, false
	}
	return e.parent.Func(name)
}

func (e *Env) AllFuncs() map[string]CalcFunc {
	r := make(map[string]CalcFunc)
	var do func(e *Env)
	do = func(e *Env) {
		for name, fn := range e.Funcs {
			r[name] = fn
		}
		if e.parent == nil {
			return
		}
		do(e.parent)
	}
	do(e)
	return r
}

func (e *Env) Interpolate(v string) (string, error) {
	var result, name strings.Builder

	inQuote := false
	inEscape := false

	for _, ch := range v {
		if ch == '[' && !inQuote && !inEscape {
			inQuote = true
		} else if ch == ']' && !inEscape {
			inQuote = false
			de := e.Derive("<interp>")
			if err := Eval(de, "<interp>", []byte(name.String())); err != nil {
				return "", err
			}
			result.WriteString(de.Stack.String())
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

func (e *Env) Import(def ModuleDef, prefix string) (*Env, error) {
	mod, err := e.Calc.Load(def)
	if err != nil {
		return nil, err
	}

	for _, name := range mod.Exports {
		qName := name
		if prefix != "" {
			qName = prefix + "." + name
		}
		e.Funcs[qName] = mod.Funcs[name]
	}
	return mod, nil
}

func (e *Env) Get(name string) string {
	s, ok := e.StackFor(name)
	if !ok {
		return ""
	}
	v, err := s.Peek()
	if err != nil {
		return ""
	}
	return v
}

func (e *Env) GetBool(name string) bool {
	s := e.Get(name)
	v, err := types.Bool.Parse(s)
	if err != nil {
		return false
	}
	return v
}

func (e *Env) GetInt(name string) int {
	s := e.Get(name)
	v, err := types.Int.Parse(s)
	if err != nil {
		return 0
	}
	return v
}

func (e *Env) Set(name string, val string) {
	s, ok := e.StackFor(name)
	if !ok {
		s = e.NewStack(name)
	}
	s.Clear().Push(val)
}

func (e *Env) SetBool(name string, val bool) {
	e.Set(name, types.Bool.Format(val))
}

func (e *Env) SetInt(name string, val int) {
	e.Set(name, types.Int.Format(val))
}
