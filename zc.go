package zc

type Item struct {
	Val any
}

func (i Item) String() string {
	return Format(i.Val)
}

type Type interface {
	From(any) (any, Type, bool)
	Recycle(any)
}

type OpEnv struct {
	*Stack[Item]
}

func (e *OpEnv) PushVal(vals ...any) {
	for _, val := range vals {
		e.Push(Item{Val: val})
	}
}

type Op struct {
	Params []Type
	Func   func(OpEnv)
}
