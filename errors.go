package zc

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func ErrDivisionByZero(e *OpEnv) error {
	return fmt.Errorf("%v: division by zero", e.Op.Name)
}

func ErrInvalidArg(e *OpEnv, format string, args ...any) error {
	return fmt.Errorf("%v: invalid argument, %v", e.Op.Name, fmt.Sprintf(format, args...))
}

func ErrNoMatchForOp(name string) error {
	return fmt.Errorf("no match for operation: %v", name)
}

func ErrNoSuchOp(name string) error {
	return fmt.Errorf("no such operation: %v", name)
}

func ErrNotEnoughArgs(have, want int) error {
	return fmt.Errorf("expected %v argument(s), got %v", want, have)
}

func ErrOp(env *OpEnv, err error) error {
	return fmt.Errorf("%v: %v", env.Op.Name, err)
}

var ErrStackEmpty = errors.New("stack empty")
var ErrStackUnderflow = errors.New("stack underflow")

func ErrWrongGoType(want string, have any) error {
	return fmt.Errorf("expected type %v, got %v", want, goName(have))
}

func goName(v any) string {
	var name strings.Builder
	t := reflect.TypeOf(v)
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
		name.WriteRune('*')
	}
	name.WriteString(t.Name())
	return name.String()
}
