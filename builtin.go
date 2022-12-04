package zc

import (
	"errors"
	"fmt"
)

var builtin = map[string]CalcFunc{
	"abort":       abort,
	"c":           clear,
	"clear":       clear,
	"copy":        copy_,
	"cp":          copy_,
	"n":           n,
	"places":      places,
	"places=":     placesGet,
	"pop":         pop,
	"print":       print,
	"round":       round,
	"round-mode":  roundMode,
	"round-mode=": roundModeGet,
	"recv":        recv,
	"send":        send,
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

func copy_(calc *Calc) error {
	a, err := calc.Stack.Pop()
	if err != nil {
		return err
	}
	calc.Stack.Push(a)
	calc.Stack.Push(a)
	return nil
}

func n(calc *Calc) error {
	n := calc.Stack.Len()
	calc.Stack.Push(FormatInt(n))
	return nil
}

func pop(calc *Calc) error {
	_, err := calc.Stack.Pop()
	return err
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
	calc.Stack.Push(FormatInt(int(Places)))
	return nil
}

func print(calc *Calc) error {
	a, err := calc.Stack.Pop()
	if err != nil {
		return err
	}
	calc.Print(a)
	return nil
}

func recv(calc *Calc) error {
	if calc.Stack == calc.main {
		return errors.New("on main stack")
	}
	a, err := calc.main.Pop()
	if err != nil {
		return err
	}
	calc.Stack.Push(a)
	return nil
}

func round(calc *Calc) error {
	places, err := calc.PopInt32()
	if err != nil {
		return err
	}
	value, err := calc.PopDecimal()
	if err != nil {
		return err
	}
	fn := roundModes[RoundMode]
	r := fn(value, places)
	calc.Stack.Push(FormatDecimal(r))
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

func send(calc *Calc) error {
	if calc.Stack == calc.main {
		return errors.New("on main stack")
	}
	a, err := calc.Stack.Pop()
	if err != nil {
		return err
	}
	calc.main.Push(a)
	return nil
}
