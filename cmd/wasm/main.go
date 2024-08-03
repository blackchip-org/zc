//go:build js

package main

import (
	"syscall/js"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app"
	"github.com/blackchip-org/zc/v6/app/repl"
)

var (
	c *app.Calc
	r *repl.Repl
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

		items := c.Items()
		for _, item := range items {
			stack = append(stack, map[string]any{
				"value": valueOf(item),
				"label": item.Label,
				"unit":  item.Unit,
			})
		}
		return map[string]any{
			"stack":  stack,
			"notice": r.Notice(),
			"error":  err,
		}
	})
}

func zcStackLen() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		return c.Len()
	})
}

func zcStack() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		var stack []any
		for _, item := range c.Items() {
			stack = append(stack, map[string]any{
				"value": valueOf(item),
				"label": item.Label,
				"unit":  item.Unit,
			})
		}
		return stack
	})
}

func zcOpNames() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		var names []any
		for _, n := range c.Catalog.OpNames() {
			names = append(names, n)
		}
		return names
	})
}

func zcSetStack() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			panic("zcSetStack: invalid number of arguments")
		}
		jsStack := args[0]
		c.SetItems([]zc.Item{})
		for i := 0; i < jsStack.Length(); i++ {
			item := jsStack.Index(i)
			val := item.Get("value")
			zc.String.Push(c, val.String())
		}
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

func valueOf(item zc.Item) string {
	return item.Val()
}

func main() {
	c = app.NewCalc()
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
