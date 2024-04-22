package kinds

var Val = ValType{}

type ValType struct{}

func (t ValType) String() string { return "Val" }

func (t ValType) Is(v any) bool {
	if v == nil {
		return false
	}
	return true
}

func (t ValType) Dup(any) any {
	panic("Dup undefined")
}

func (t ValType) Copy(any, any) {
	panic("Copy undefined")
}

func (t ValType) To(a any) (any, bool) {
	return a, true
}
