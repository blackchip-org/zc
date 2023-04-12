package types

type vString struct {
	val string
}

func (v vString) Type() Type     { return String }
func (v vString) Format() string { return v.val }
func (v vString) String() string { return stringV(v) }
func (v vString) Native() any    { return v.val }

type StringType struct{}

func (t StringType) String() string { return "String" }

func (t StringType) ParseValue(s string) (Value, error) {
	return vString{val: s}, nil
}

func (t StringType) Value(s string) Value {
	return vString{val: s}
}

func (t StringType) Native(v Value) string {
	return v.(vString).val
}
