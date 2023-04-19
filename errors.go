package zc

import (
	"errors"
	"fmt"
	"strings"
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

	ErrNoOpForTypes = func(c Calc, op string, types ...Type) {
		var typeNames []string
		for _, t := range types {
			typeNames = append(typeNames, t.String())
		}
		c.SetError(fmt.Errorf("%v: no operation for %v", op, strings.Join(typeNames, ", ")))
	}

	ErrNotEnoughArguments = func(c Calc, op string, expected int) {
		c.SetError(fmt.Errorf("%v: not enough arguments, expected %v", op, expected))
	}

	ErrStackEmpty = errors.New("stack empty")

	ErrUnexpectedType = func(t Type, val string) error {
		return fmt.Errorf("expected %v: %v", t, val)
	}

	ErrUnknownOp = func(name string) error {
		return fmt.Errorf("unknown operation: %v", name)
	}
)
