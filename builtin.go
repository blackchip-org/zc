package zc

import (
	"errors"
)

var builtin = map[string]CalcFunc{
	"abort": abort,
	"clear": clear,
	"pop":   pop,
	"z":     clear,
}

func abort(calc *Calc) error {
	a, err := calc.Stack.Pop()
	if err != nil {
		return errors.New("aborted")
	}
	return errors.New(a)
}

func clear(calc *Calc) error {
	calc.Stack.Clear()
	return nil
}

func pop(calc *Calc) error {
	_, err := calc.Stack.Pop()
	return err
}
