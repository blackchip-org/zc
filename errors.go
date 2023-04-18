package zc

import (
	"errors"
	"fmt"
)

var (
	ErrDivisionByZero       = errors.New("division by zero")
	ErrDuplicateOp          = func(name string) error { return fmt.Errorf("duplicate operation: %v", name) }
	ErrStackEmpty           = errors.New("stack empty")
	ErrInvalidArgumentTypes = func(name string) error { return fmt.Errorf("%v: invalid argument types", name) }
	ErrUnexpectedType       = func(t Type, val string) error { return fmt.Errorf("expected %v: %v", t, val) }
	ErrUnknownOp            = func(name string) error { return fmt.Errorf("unknown operation: %v", name) }
)
