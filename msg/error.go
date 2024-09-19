package msg

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func ErrDivisionByZero() error {
	return errors.New("division by zero")
}

func ErrDoesNotReduce() error {
	return fmt.Errorf("does not reduce")
}

func ErrFeatureNotSupported(name string) error {
	return fmt.Errorf("feature not supported: %v", name)
}

func ErrIndexOutOfRange(i int) error {
	return fmt.Errorf("index out of range: %v", i)
}

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

func ErrInvalidRoundingMode(m string) error {
	return fmt.Errorf("invalid rounding mode: %v", m)
}

func ErrInvalidArg(format string, args ...any) error {
	return fmt.Errorf("invalid argument, %v", fmt.Sprintf(format, args...))
}

func ErrMemoryEmpty(name string) error {
	return fmt.Errorf("memory empty: %v", name)
}

func ErrNoReturnValues() error {
	return errors.New("no return values")
}

func ErrNoSuchOp(name string) error {
	return fmt.Errorf("no such operation: %v", name)
}

func ErrNotANumber() error {
	return errors.New("not a number")
}

func ErrNotEnoughArgs() error {
	return errors.New("not enough arguments")
}

func ErrOp(name string, err error) error {
	return fmt.Errorf("%v: %v", name, err)
}

func ErrOverflow(val string) error {
	return fmt.Errorf("overflow: %v", Abbr(val))
}

func ErrStackEmpty() error {
	return errors.New("stack empty")
}

func ErrUnderflow(val string) error {
	return fmt.Errorf("underflow: %v", Abbr(val))
}

func ErrUnexpectedType(val string) error {
	return fmt.Errorf("unexpected type: %v", Quote(Abbr(val)))
}

func ErrUnknownTimeZone(zone string) error {
	return fmt.Errorf("unknown time zone: %v", Quote(zone))
}

func ErrWrongGoType(want string, have any) error {
	return fmt.Errorf("expected type %v, got: %v of type %v", want, Quote(fmt.Sprint(have)), goName(have))
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
