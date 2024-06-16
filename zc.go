package zc

type Item struct {
	Val any
}

type Type interface {
	From(any) (any, Type, bool)
	Recycle(any)
}

type OpEnv struct {
	*Stack[*Item]
}

type Op struct {
	Params []Type
	Func   func(OpEnv)
}
