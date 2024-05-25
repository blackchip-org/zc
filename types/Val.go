package types

var Val = valType{}

type valType struct{}

func (t valType) Name() string { return "Val" }

func (t valType) Is(v any) bool {
	if v == nil {
		return false
	}
	return true
}

func (t valType) Dup(any) any {
	panic("type Val cannot be duplicated")
}

func (t valType) Copy(any, any) {
	panic("type Val cannot be copied")
}

func (t valType) To(a any) (any, bool) {
	return a, false
}
