package zc

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// func ErrArgMismatch(name string) error {
// 	return fmt.Errorf("argument mismatch for operation: %v", name)
// }

var ErrDivisionByZero = errors.New("division by zero")

func ErrInfinity(sign int) error {
	var inf string
	switch {
	case sign < 0:
		inf = "-infinity"
	case sign > 0:
		inf = "+infinity"
	default:
		inf = "infinity"
	}
	return errors.New(inf)
}

func ErrInvalidArg(format string, args ...any) error {
	return fmt.Errorf("invalid argument, %v", fmt.Sprintf(format, args...))
}

func ErrNoSuchOp(name string) error {
	return fmt.Errorf("no such operation: %v", name)
}

var ErrNotANumber = errors.New("not a number")

func ErrOp(name string, err error) error {
	return fmt.Errorf("%v: %v", name, err)
}

func ErrOverflow(val string) error {
	return fmt.Errorf("overflow: %v", Abbr(val))
}

var ErrStackEmpty = errors.New("stack empty")

func ErrUnderflow(val string) error {
	return fmt.Errorf("underflow: %v", Abbr(val))
}

func ErrUnexpectedType(val string) error {
	return fmt.Errorf("unexpected type: %v", Quote(Abbr(val)))
}

func ErrWrongGoType(want Type, have any) error {
	return fmt.Errorf("expected type %v, got: %v of type %v", want.GoName(), Quote(fmt.Sprint(have)), goName(have))
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
