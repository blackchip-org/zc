package types

import (
	"fmt"
	"strings"
)

type gBool struct {
	val bool
}

func (g gBool) Type() Type     { return Bool }
func (g gBool) Format() string { return Bool.Format(g.val) }
func (g gBool) String() string { return fmt.Sprintf("%v(%v)", g.Type().String(), g.Format()) }
func (g gBool) Native() any    { return g.val }

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
	return gBool{val: b}
}

func (t BoolType) Native(v Value) bool {
	return v.Native().(bool)
}
