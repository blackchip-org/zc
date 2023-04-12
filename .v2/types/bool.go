package types

import (
	"fmt"
	"strings"
)

type vBool struct {
	val bool
}

func (v vBool) Type() Type     { return Bool }
func (v vBool) Format() string { return Bool.Format(v.val) }
func (v vBool) String() string { return fmt.Sprintf("%v(%v)", v.Type().String(), v.Format()) }
func (v vBool) Native() any    { return v.val }

type BoolType struct{}

func (t BoolType) String() string { return "Bool" }

func (t BoolType) Parse(s string) (bool, error) {
	sl := strings.ToLower(s)
	switch sl {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}
	return false, parseErr(t, s)
}

func (t BoolType) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return Nil, err
	}
	return t.Value(v), nil
}

func (t BoolType) Format(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func (t BoolType) Value(b bool) Value {
	return vBool{val: b}
}

func (t BoolType) Native(v Value) bool {
	return v.Native().(bool)
}
