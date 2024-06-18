package zc

type Item struct {
	Val any
}

func (i Item) String() string {
	return Format(i.Val)
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
	Name      string
	Overloads string
	Params    []Type
	Returns   []Type
	Func      func(OpEnv)
}

type Macro struct {
	Name string
	Expr string
}

type Vol struct {
	Name   string
	Types  []Type
	Ops    []Op
	Macros []Macro
}
