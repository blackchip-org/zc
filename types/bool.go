package types

import (
	"strings"
)

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

func (t BoolType) Format(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
