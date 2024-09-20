package msg

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	ErrDivisionByZero      = errors.New("division by zero")
	ErrDoesNotReduce       = errors.New("does not reduce")
	ErrFeatureNotSupported = func(name string) error {
		return fmt.Errorf("feature not supported: %v", name)
	}
	ErrIndexOutOfRange = func(i int) error {
		return fmt.Errorf("index out of range: %v", i)
	}
	ErrInfinity = func(sign int) error {
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
	ErrInvalidRoundingMode = func(m string) error {
		return fmt.Errorf("invalid rounding mode: %v", m)
	}
	ErrInvalidArg = func(format string, args ...any) error {
		return fmt.Errorf("invalid argument, %v", fmt.Sprintf(format, args...))
	}
	ErrMemoryEmpty = func(name string) error {
		return fmt.Errorf("memory empty: %v", name)
	}
	ErrNoReturnValues = errors.New("no return values")
	ErrNoSuchOp       = func(name string) error {
		return fmt.Errorf("no such operation: %v", name)
	}
	ErrNotANumber    = errors.New("not a number")
	ErrNotEnoughArgs = errors.New("not enough arguments")
	ErrOp            = func(name string, err error) error {
		return fmt.Errorf("%v: %v", name, err)
	}
	ErrOverflow = func(val string) error {
		return fmt.Errorf("overflow: %v", Abbr(val))
	}
	ErrStackEmpty = errors.New("stack empty")
	ErrUnderflow  = func(val string) error {
		return fmt.Errorf("underflow: %v", Abbr(val))
	}
	ErrUnexpectedType = func(val string) error {
		return fmt.Errorf("unexpected type: %v", Quote(Abbr(val)))
	}
	ErrUnknownTimeZone = func(zone string) error {
		return fmt.Errorf("unknown time zone: %v", Quote(zone))
	}
	ErrWrongGoType = func(want string, have any) error {
		return fmt.Errorf("expected type %v, got: %v of type %v", want, Quote(fmt.Sprint(have)), goName(have))
	}
)

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
