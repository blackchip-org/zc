package ops

import t "github.com/blackchip-org/zc/types"

func op1Bool(fn func(bool) bool) Func {
	return func(args []t.Generic) ([]t.Generic, error) {
		x := t.Bool.Value(args[0])
		z := fn(x)
		return []t.Generic{t.Bool.Generic(z)}, nil
	}
}

func op2Bool(fn func(bool, bool) bool) Func {
	return func(args []t.Generic) ([]t.Generic, error) {
		x := t.Bool.Value(args[0])
		y := t.Bool.Value(args[1])
		z := fn(x, y)
		return []t.Generic{t.Bool.Generic(z)}, nil
	}
}

var (
	andBool = op2Bool(func(x bool, y bool) bool { return x && y })
	eqBool  = op2Bool(func(x bool, y bool) bool { return x == y })
	orBool  = op2Bool(func(x bool, y bool) bool { return x || y })
	notBool = op1Bool(func(x bool) bool { return !x })
)
