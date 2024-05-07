package errors

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"
)

func CannotConvert(fromKind string, dest any, val any) error {
	toType := nameOf(dest)
	return fmt.Errorf("cannot convert %v from %v to %v", val, fromKind, toType)
}

func DuplicateKind(name string) error {
	return fmt.Errorf("duplicate kind: %v", name)
}

func DuplicateOp(name string) error {
	return fmt.Errorf("duplicate op: %v", name)
}

func InvalidArgKinds(kindNames []string, varArgs string) error {
	ks := slices.Clone(kindNames)
	if varArgs != "" {
		ks = append(ks, varArgs+"*")
	}
	return fmt.Errorf("invalid arguments, expected %v", strings.Join(ks, " | "))
}

func InvalidArgCount(expected int) error {
	return fmt.Errorf("not enough arguments, expected %v", expected)
}

func InvalidRetKinds(kindNames []string, varRets string) error {
	ks := slices.Clone(kindNames)
	if varRets != "" {
		ks = append(ks, varRets+"*")
	}
	return fmt.Errorf("invalid returns, expected %v", strings.Join(ks, " | "))
}

func NoFuncForOp(name string) error {
	return fmt.Errorf("no function for op: %v", name)
}

func NoSuchOp(name string) error {
	return fmt.Errorf("no such operation: %v", name)
}

func NoMatchOp(name string) error {
	return fmt.Errorf("no match for operation: %v", name)
}

func NotOverloaded(name string) error {
	return fmt.Errorf("operation cannot be overloaded: %v", name)
}

func StackEmpty() error {
	return errors.New("stack empty")
}

func UnknownKind(name string) error {
	return fmt.Errorf("unknown kind: %v", name)
}

func UnexpectedType(expected string, val any) error {
	actual := nameOf(val)
	return fmt.Errorf("expected type %v but got %v with value %v", expected, actual, val)
}

func UnregisteredType(v any) error {
	return fmt.Errorf("unregistered type: %v", nameOf(v))
}

func nameOf(v any) string {
	var name strings.Builder
	t := reflect.TypeOf(v)
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
		name.WriteRune('*')
	}
	name.WriteString(t.Name())
	return name.String()
}
