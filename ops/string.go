package ops

import t "github.com/blackchip-org/zc/types"

func opCmpString(fn func(string, string) bool) OpFn {
	return func(args []t.Generic) ([]t.Generic, error) {
		x := t.String.Value(args[0])
		y := t.String.Value(args[1])
		z := fn(x, y)
		return []t.Generic{t.Bool.Generic(z)}, nil
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
