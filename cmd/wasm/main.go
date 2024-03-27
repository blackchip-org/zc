//go:build js

package main

import (
	"syscall/js"

	"github.com/blackchip-org/zc/v6/pkg/calc"
	"github.com/blackchip-org/zc/v6/pkg/repl"
	"github.com/blackchip-org/zc/v6/pkg/zc"
)

var (
	c zc.Calc
	r *repl.REPL
)

func zcCommonPrefix() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			panic("zcCommonPrefix: invalid number of arguments")
		}
		jsValues := args[0]
		var outValues []string
		for i := 0; i < jsValues.Length(); i++ {
			outValues = append(outValues, jsValues.Index(i).String())
		}
		common := repl.CommonPrefix(outValues)
		return common
	})
}

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

func zcQuoteEnd() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		return r.EndQuote
	})
}

func zcWordCompleter() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 2 {
			panic("zcWordCompleter: invalid number of arguments")
		}
		line := args[0].String()
		pos := args[1].Int()
		prefix, candidates, suffix := r.WordCompleter(line, pos)
		var jsCandidates []any
		for _, c := range candidates {
			jsCandidates = append(jsCandidates, c)
		}
		return map[string]any{
			"prefix":     prefix,
			"candidates": jsCandidates,
			"suffix":     suffix,
		}
	})
}

func main() {
	c = calc.New()
	r = repl.New(c)

	js.Global().Set("zcCommonPrefix", zcCommonPrefix())
	js.Global().Set("zcEval", zcEval())
	js.Global().Set("zcStack", zcStack())
	js.Global().Set("zcStackLen", zcStackLen())
	js.Global().Set("zcOpNames", zcOpNames())
	js.Global().Set("zcSetStack", zcSetStack())
	js.Global().Set("zcQuoteEnd", zcQuoteEnd())
	js.Global().Set("zcWordCompleter", zcWordCompleter())

	<-make(chan struct{})
}
