package zc

import (
	"errors"
	"fmt"
)

func ErrArgMismatch(name string) error {
	return fmt.Errorf("argument mismatch for operation: %v", name)
}

func ErrNoSuchOp(name string) error {
	return fmt.Errorf("no such operation: %v", name)
}

var ErrStackEmpty = errors.New("stack empty")

func ErrWrongGoType(want Type, have any) error {
	return fmt.Errorf("expected type %v, got: %v", want.GoName(), have)
}

/*
func goName(v any) string {
	if v == nil {
		return "nil"
	}
	var name strings.Builder
	t := reflect.TypeOf(v)
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
		name.WriteRune('*')
	}
	name.WriteString(t.Name())
	return name.String()
}
*/
