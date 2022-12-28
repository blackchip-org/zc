package zlib

import (
	"fmt"
	"sort"

	"github.com/blackchip-org/zc"
)

func Exports(env *zc.Env) error {
	var mod *zc.Env
	var ok bool

	name, err := env.Stack.Pop()
	if err != nil {
		return fmt.Errorf("expecting module name")
	}

	mod, ok = env.Calc.Modules[name]
	if !ok {
		return fmt.Errorf("no such module: %v", name)
	}
	funcs := make([]string, len(mod.Exports))
	copy(funcs, mod.Exports)
	sort.Strings(funcs)
	for _, f := range funcs {
		env.Stack.Push(f)
	}

	return nil
}

func Funcs(env *zc.Env) error {
	var funcs []string
	for f := range env.Funcs {
		funcs = append(funcs, f)
	}
	sort.Strings(funcs)
	for _, f := range funcs {
		env.Stack.Push(f)
	}
	return nil
}
