package types

import (
	"fmt"
	"strings"
)

type boolVal struct {
	val bool
}

func (v boolVal) Type() Type { return Bool }
func (v boolVal) Format() string {
	if v.val {
		return "true"
	}
	return "false"
}
func (v boolVal) String() string { return fmt.Sprintf("%v(%v)", v.Type().String(), v.Format()) }

type BoolType struct{}

func (t BoolType) String() string { return "Bool" }

func (t BoolType) Parse(s string) (bool, bool) {
	sl := strings.ToLower(s)
	switch sl {
	case "true":
		return true, true
	case "false":
		return false, true
	}
	return false, false
}

func (t BoolType) ParseValue(s string) (Value, bool) {
	b, ok := t.Parse(s)
	if !ok {
		return Nil, false
	}
	return t.Value(b), true
}

func (t BoolType) Value(b bool) Value {
	return boolVal{val: b}
}

func (t BoolType) Unwrap(v Value) bool {
	return v.(boolVal).val
}
