package zc

import (
	"errors"
	"fmt"
)

var (
	ErrDivisionByZero = func(c Calc, a0 any, a1 any) {
		c.SetError(fmt.Errorf("division by zero: %v %v", Format(a0), Format(a1)))
	}
	ErrDuplicateOp = func(name string) error {
		return fmt.Errorf("duplicate operation: %v", name)
	}
	ErrInvalidArgumentTypes = func(name string) error {
		return fmt.Errorf("%v: invalid argument types", name)
	}
	ErrModuloByZero = func(c Calc, a0 any, a1 any) {
		c.SetError(fmt.Errorf("modulo by zero: %v %v", Format(a0), Format(a1)))
	}
	ErrStackEmpty     = errors.New("stack empty")
	ErrUnexpectedType = func(t Type, val string) error {
		return fmt.Errorf("expected %v: %v", t, val)
	}
	ErrUnknownOp = func(name string) error {
		return fmt.Errorf("unknown operation: %v", name)
	}
)
