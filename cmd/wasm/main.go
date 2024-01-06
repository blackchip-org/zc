package main

import (
	"syscall/js"

	"github.com/blackchip-org/zc/v5/pkg/calc"
	"github.com/blackchip-org/zc/v5/pkg/zc"
)

var c zc.Calc

func zcEval() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		in := args[0].String()
		c.Eval(in)
		var err string
		var stack []interface{}

		if c.Error() != nil {
			err = c.Error().Error()
		}
		for i := 0; i < c.StackLen(); i++ {
			stack = append(stack, c.Stack()[i])
		}
		return map[string]any{
			"stack": stack,
			"info":  c.Info(),
			"error": err,
		}
	})
}

func main() {
	c = calc.New()
	js.Global().Set("zcEval", zcEval())
	<-make(chan struct{})
}
