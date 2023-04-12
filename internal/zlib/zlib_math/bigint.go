package zlib_math

/*
func op2BigInt(fn func(t.BigInt, t.BigInt) t.BigInt) zc.CalcFunc {
	return func(_ zc.Env, args []t.Value) ([]t.Value, error) {
		x := zc.BigInt.Native(args[0])
		y := zc.BigInt.Native(args[1])
		z := fn(x, y)
		return []t.Value{zc.BigInt.Value(z)}, nil
	}
}

func InitBigInt(l *zc.Library) {
	l.RegisterGenOp("add", addBigInt, zc.BigInt, zc.BigInt)
}

var (
	addBigInt = op2BigInt(func(x t.BigInt, y t.BigInt) t.BigInt { return x.Add(y) })
)
*/
