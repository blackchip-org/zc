package ops

import t "github.com/blackchip-org/zc/types"

func op1Bool(fn func(bool) bool) Func {
	return func(args []t.Value) ([]t.Value, error) {
		x := t.Bool.Native(args[0])
		z := fn(x)
		return []t.Value{t.Bool.Value(z)}, nil
	}
}

func op2Bool(fn func(bool, bool) bool) Func {
	return func(args []t.Value) ([]t.Value, error) {
		x := t.Bool.Native(args[0])
		y := t.Bool.Native(args[1])
		z := fn(x, y)
		return []t.Value{t.Bool.Value(z)}, nil
	}
}

var (
	andBool = op2Bool(func(x bool, y bool) bool { return x && y })
	eqBool  = op2Bool(func(x bool, y bool) bool { return x == y })
	orBool  = op2Bool(func(x bool, y bool) bool { return x || y })
	notBool = op1Bool(func(x bool) bool { return !x })
)
