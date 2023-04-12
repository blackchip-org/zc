package types

type vNone struct{}

func (v vNone) Type() Type     { return None }
func (v vNone) Format() string { return "" }
func (v vNone) String() string { return "None" }
func (v vNone) Native() any    { return nil }

type noneType struct{}

func (t noneType) String() string { return "None" }

func (t noneType) ParseValue(s string) (Value, error) {
	return vNone{}, nil
}
