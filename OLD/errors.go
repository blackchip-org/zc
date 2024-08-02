package zc

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func ErrArgMismatch(name string) error {
	return fmt.Errorf("argument mismatch for operation: %v", name)
}

func ErrDivisionByZero(e *OpEnv) error {
	return fmt.Errorf("%v: division by zero", e.Op.Name)
}

func ErrInfinity(e *OpEnv, sign int) error {
	var inf string
	switch {
	case sign < 0:
		inf = "-infinity"
	case sign > 0:
		inf = "+infinity"
	default:
		inf = "infinity"
	}
	return fmt.Errorf("%v: %v", e.Op.Name, inf)
}

func ErrInvalidArg(e *OpEnv, format string, args ...any) error {
	return fmt.Errorf("%v: invalid argument, %v", e.Op.Name, fmt.Sprintf(format, args...))
}

func ErrNoSuchOp(name string) error {
	return fmt.Errorf("no such operation: %v", name)
}

func ErrNotANumber(e *OpEnv) error {
	return fmt.Errorf("%v: not a number", e.Op.Name)
}

func ErrNotEnoughArgs(have, want int) error {
	return fmt.Errorf("expected %v argument(s), got %v", want, have)
}

func ErrOp(env *OpEnv, err error) error {
	return fmt.Errorf("%v: %v", env.Op.Name, err)
}

func ErrReturnMismatch(name string) error {
	return fmt.Errorf("return mismatch for operation: %v", name)
}

var ErrStackEmpty = errors.New("stack empty")
var ErrStackUnderflow = errors.New("stack underflow")

func ErrUndefined(e *OpEnv) error {
	return fmt.Errorf("%v: undefined", e.Op.Name)
}

func ErrWrongGoType(want string, have any) error {
	return fmt.Errorf("expected type %v, got %v", want, goName(have))
}

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
