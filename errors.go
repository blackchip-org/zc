package zc

import (
	"fmt"
)

func ErrUnsupported(name string) error {
	return fmt.Errorf("unsupported operation: %v", name)
}

type EmptyStackError struct {
	Name string
}

func (e EmptyStackError) Error() string {
	return e.Name + ": empty stack"
}
