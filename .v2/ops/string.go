package ops

import t "github.com/blackchip-org/zc/types"

func opCmpString(fn func(string, string) bool) Func {
	return func(args []t.Value) ([]t.Value, error) {
		x := t.String.Native(args[0])
		y := t.String.Native(args[1])
		z := fn(x, y)
		return []t.Value{t.Bool.Value(z)}, nil
	}
}

var (
	eqString  = opCmpString(func(x string, y string) bool { return x == y })
	gtString  = opCmpString(func(x string, y string) bool { return x > y })
	gteString = opCmpString(func(x string, y string) bool { return x >= y })
	ltString  = opCmpString(func(x string, y string) bool { return x < y })
	lteString = opCmpString(func(x string, y string) bool { return x <= y })
	neqString = opCmpString(func(x string, y string) bool { return x != y })
)
