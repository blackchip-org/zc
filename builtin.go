package zc

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var builtin = map[string]CalcFunc{
	"abort":       abort,
	"exit":        exit,
	"places":      places,
	"places=":     placesGet,
	"n":           n,
	"nothing":     nothing,
	"round":       round,
	"round-mode":  roundMode,
	"round-mode=": roundModeGet,
	"trace":       trace,
	"trace-off":   traceOff,
	"undef":       undef,
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
	code, err := calc.ParseInt(a)
	if err != nil {
		return err
	}
	os.Exit(code)
	return nil
}

func places(calc *Calc) error {
	places, err := calc.PopInt32()
	if err != nil {
		return err
	}
	if places < 0 {
		return fmt.Errorf("invalid number of places: %v", places)
	}
	Places = places
	calc.Printf("ok")
	return nil
}

func placesGet(calc *Calc) error {
	calc.Stack.Push(calc.FormatInt(int(Places)))
	return nil
}

func n(calc *Calc) error {
	calc.Stack.Push(calc.FormatInt(calc.Stack.Len()))
	return nil
}

func nothing(calc *Calc) error {
	return nil
}

func round(calc *Calc) error {
	places, err := calc.PopInt32()
	if err != nil {
		return err
	}
	value, err := calc.PopFix()
	if err != nil {
		return err
	}
	fn := roundModes[RoundMode]
	r := fn(value, places)
	calc.Stack.Push(calc.FormatDecimal(r))
	return nil
}

func roundMode(calc *Calc) error {
	mode, err := calc.Stack.Pop()
	if err != nil {
		return err
	}
	_, ok := roundModes[mode]
	if !ok {
		return fmt.Errorf("invalid rounding mode: %v", mode)
	}
	RoundMode = mode
	calc.Print("ok")
	return err
}

func roundModeGet(calc *Calc) error {
	calc.Stack.Push(RoundMode)
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
