package zc

import (
	"fmt"
	"strings"
)

type Env struct {
	parent  *Env
	Calc    *Calc
	Stack   *Stack
	Main    *Stack
	stacks  map[string]*Stack
	Funcs   map[string]CalcFunc
	Exports []string
	Module  string
}

func NewEnv(calc *Calc) *Env {
	e := &Env{
		Calc:   calc,
		stacks: make(map[string]*Stack),
		Funcs:  make(map[string]CalcFunc),
	}
	e.Main = NewStack(calc, "main")
	e.stacks["main"] = e.Main
	e.Stack = e.Main
	return e
}

func (e *Env) Derive() *Env {
	de := NewEnv(e.Calc)
	de.parent = e
	return de
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
	s := NewStack(e.Calc, name)
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

func (e *Env) Interpolate(v string) (string, error) {
	var result, name strings.Builder

	inQuote := false
	inEscape := false

	for _, ch := range v {
		if ch == '`' && !inQuote && !inEscape {
			inQuote = true
		} else if ch == '`' && !inEscape {
			inQuote = false
			stack, ok := e.StackFor(name.String())
			if !ok {
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
