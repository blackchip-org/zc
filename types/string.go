package types

import "fmt"

type gString struct {
	val string
}

func (g gString) Type() Type     { return String }
func (g gString) Format() string { return g.val }
func (g gString) String() string { return fmt.Sprintf("%v(%v)", g.Type().String(), g.Format()) }
func (g gString) Value() any     { return g.val }

type StringType struct{}

func (t StringType) String() string { return "String" }

func (t StringType) ParseGeneric(s string) (Generic, error) {
	return gString{val: s}, nil
}

func (t StringType) Generic(s string) Generic {
	return gString{val: s}
}

func (t StringType) Value(v Generic) string {
	return v.(gString).val
}
