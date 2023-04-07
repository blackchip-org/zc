package types

type noneVal struct{}

func (v noneVal) Type() Type     { return None }
func (v noneVal) Format() string { return "" }
func (v noneVal) String() string { return "None" }

type noneType struct{}

func (t noneType) String() string { return "None" }

func (t noneType) ParseValue(s string) (Value, bool) {
	return noneVal{}, false
}
