package zc

// var builtin = map[string]CalcFunc{
// 	"abort":     abort,
// 	"exit":      exit,
// 	"nothing":   nothing,
// 	"trace":     trace,
// 	"trace-off": traceOff,
// 	"undef":     undef,
// }

// func undef(calc *Calc) error {
// 	target, err := calc.Stack.Pop()
// 	if err != nil {
// 		return err
// 	}

// 	var n = 0
// 	for name := range calc.Funcs {
// 		parts := strings.Split(name, ".")
// 		if parts[0] == target {
// 			delete(calc.Funcs, name)
// 			n++
// 		}
// 	}
// 	calc.Printf("%v undefined", n)
// 	return nil
// }
