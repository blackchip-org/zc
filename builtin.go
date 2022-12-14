package zc

import (
	"errors"
	"os"
	"strings"
)

var builtin = map[string]CalcFunc{
	"abort":     abort,
	"exit":      exit,
	"n":         n,
	"nothing":   nothing,
	"trace":     trace,
	"trace-off": traceOff,
	"undef":     undef,
}

func abort(calc *Calc) error {
	a, err := calc.Stack.Pop()
	if err != nil {
		return errors.New("aborted")
	}
	return errors.New(a)
}

func eval(calc *Calc) error {
	node, err := calc.Stack.Pop()
	if err != nil {
		return err
	}
	return calc.EvalString("<eval>", node)
}

func exit(calc *Calc) error {
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

func n(calc *Calc) error {
	calc.Stack.Push(calc.Val.FormatInt(calc.Stack.Len()))
	return nil
}

func nothing(calc *Calc) error {
	return nil
}

func trace(calc *Calc) error {
	calc.config.Trace = true
	return nil
}

func traceOff(calc *Calc) error {
	calc.config.Trace = false
	return nil
}

func undef(calc *Calc) error {
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
