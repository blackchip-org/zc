package types

import "fmt"

type gString struct {
	val string
}

func (g gString) Type() Type     { return String }
func (g gString) Format() string { return g.val }
func (g gString) String() string { return fmt.Sprintf("%v(%v)", g.Type().String(), g.Format()) }
func (g gString) Native() any    { return g.val }

type StringType struct{}

func (t StringType) String() string { return "String" }

func (t StringType) ParseValue(s string) (Value, error) {
	return gString{val: s}, nil
}

func (t StringType) Value(s string) Value {
	return gString{val: s}
}

func (t StringType) Native(v Value) string {
	return v.(gString).val
}
