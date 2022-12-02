package zc

var builtin = map[string]NativeFn{
	"clear": clear,
	"pop":   pop,
	"z":     clear,
}

func clear(calc *Calc) error {
	calc.Stack.Clear()
	return nil
}

func pop(calc *Calc) error {
	_, err := calc.Stack.Pop()
	return err
}
