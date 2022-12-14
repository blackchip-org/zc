package zlib

import (
	"errors"
	"os"
	"strings"

	"github.com/blackchip-org/zc"
)

func Abort(calc *zc.Calc) error {
	a, err := calc.Stack.Pop()
	if err != nil {
		return errors.New("aborted")
	}
	return errors.New(a)
}

func Eval(calc *zc.Calc) error {
	node, err := calc.Stack.Pop()
	if err != nil {
		return err
	}
	return calc.EvalString("<eval>", node)
}

func Exit(calc *zc.Calc) error {
	a, err := calc.Stack.Pop()
	if err != nil {
		return err
	}
	code, err := calc.Val.ParseInt(a)
	if err != nil {
		return err
	}
	os.Exit(code)
	return nil
}

func Nothing(calc *zc.Calc) error {
	return nil
}

func Trace(calc *zc.Calc) error {
	calc.Config.Trace = true
	return nil
}

func TraceOff(calc *zc.Calc) error {
	calc.Config.Trace = false
	return nil
}

// FIXME: This should be handled better. Maybe a statement?
func Undef(calc *zc.Calc) error {
	target, err := calc.Stack.Pop()
	if err != nil {
		return err
	}

	var n = 0
	for name := range calc.Funcs {
		parts := strings.Split(name, ".")
		if parts[0] == target {
			delete(calc.Funcs, name)
			n++
		}
	}
	calc.Printf("%v undefined", n)
	return nil
}
