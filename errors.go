package zc

import (
	"fmt"
	"reflect"
	"strings"
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
