package zc

type Item struct {
	Val any
}

type Type interface {
	From(Recycler, *Item) (any, bool)
}

type Recycler interface {
	Recycle() (*Item, bool)
}

type OpEnv struct {
	*Stack[*Item]
}

type Op struct {
	Params []Type
	Func   func(OpEnv)
}
