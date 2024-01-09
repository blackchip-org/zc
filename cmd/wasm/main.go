//go:build js

package main

import (
	"syscall/js"

	"github.com/blackchip-org/zc/v5/pkg/calc"
	"github.com/blackchip-org/zc/v5/pkg/repl"
	"github.com/blackchip-org/zc/v5/pkg/zc"
)

var (
	c zc.Calc
	r *repl.REPL
)

func zcEval() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		in := args[0].String()
		r.Eval(in)
		var err string
		var stack []any

		if r.Error() != nil {
			err = r.Error().Error()
		}
		for i := 0; i < c.StackLen(); i++ {
			stack = append(stack, c.Stack()[i])
		}
		return map[string]any{
			"stack": stack,
			"info":  r.Info(),
			"error": err,
		}
	})
}

func zcStackLen() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		return c.StackLen()
	})
}

func zcStack() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		var stack []any
		for _, item := range c.Stack() {
			stack = append(stack, item)
		}
		return stack
	})
}

func zcOpNames() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		var ops []any
		for _, item := range c.OpNames() {
			ops = append(ops, item)
		}
		return ops
	})
}

func zcSetStack() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			panic("zcSetStack: invalid number of arguments")
		}
		jsStack := args[0]
		var outStack []string
		for i := 0; i < jsStack.Length(); i++ {
			outStack = append(outStack, jsStack.Index(i).String())
		}
		c.SetStack(outStack)
		return nil
	})
}

func main() {
	c = calc.New()
	r = repl.New(c)

	js.Global().Set("zcEval", zcEval())
	js.Global().Set("zcStack", zcStack())
	js.Global().Set("zcStackLen", zcStackLen())
	js.Global().Set("zcOpNames", zcOpNames())
	js.Global().Set("zcSetStack", zcSetStack())
	<-make(chan struct{})
}
