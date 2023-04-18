```go

var ops map[string]Something {
    "add": GenOp(genAdd)
}

genAdd = []Something{
    op(addBigInt, ts.BigInt, ts.BigInt)
}

func addBigInt(e zc.Env) {
    x := zc.BigIntArg(e)
    y := zc.BigIntArg(e)

    z := x.Add(y)

    e.Res.BigInt(z)
}

```