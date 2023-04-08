package types

import (
	"fmt"
	"strings"
)

type gBool struct {
	val bool
}

func formatBool(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func (g gBool) Type() Type     { return Bool }
func (g gBool) Format() string { return formatBool(g.val) }
func (g gBool) String() string { return fmt.Sprintf("%v(%v)", g.Type().String(), g.Format()) }
func (g gBool) Value() any     { return g.val }

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

func (t BoolType) ParseGeneric(s string) (Generic, bool) {
	b, ok := t.Parse(s)
	if !ok {
		return Nil, false
	}
	return t.Generic(b), true
}

func (t BoolType) Generic(b bool) Generic {
	return gBool{val: b}
}

func (t BoolType) Value(v Generic) bool {
	return v.Value().(bool)
}
