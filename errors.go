package zc

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func ErrNotEnoughArgs(have, want int) error {
	return fmt.Errorf("expected %v argument(s), got %v", have, want)
}

var (
	ErrStackEmpty     = errors.New("stack empty")
	ErrStackUnderflow = errors.New("stack underflow")
)

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
