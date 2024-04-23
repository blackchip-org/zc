package zcalc

var ValKind = valKind{}

type valKind struct{}

func (k valKind) Name() string { return "Val" }

func (k valKind) Is(v any) bool {
	if v == nil {
		return false
	}
	return true
}

func (k valKind) Dup(any) any {
	panic("Dup undefined")
}

func (k valKind) Copy(any, any) {
	panic("Copy undefined")
}

func (k valKind) To(a any) (any, bool) {
	return a, true
}
