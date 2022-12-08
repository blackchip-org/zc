package native

import (
	"fmt"
	"sort"

	"github.com/blackchip-org/zc"
)

func Exports(calc *zc.Calc) error {
	var mod *zc.Calc
	var ok bool

	name, err := calc.Stack.Pop()
	if err != nil {
		return fmt.Errorf("expecting module name")
	}

	mod, ok = calc.Modules[name]
	if !ok {
		return fmt.Errorf("no such module: %v", name)
	}

	var funcs []string
	for f := range mod.Exports {
		funcs = append(funcs, f)
	}
	sort.Strings(funcs)
	for _, f := range funcs {
		calc.Stack.Push(f)
	}

	return nil
}

func Funcs(calc *zc.Calc) error {
	var funcs []string
	for f := range calc.Funcs {
		funcs = append(funcs, f)
	}
	sort.Strings(funcs)
	for _, f := range funcs {
		calc.Stack.Push(f)
	}

	return nil
}
