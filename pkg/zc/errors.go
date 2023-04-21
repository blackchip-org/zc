package zc

import (
	"fmt"
	"strings"
)

func ErrDivisionByZero(c Calc, a0 any, a1 any) {
	c.SetError(fmt.Errorf("division by zero: %v %v", Format(a0), Format(a1)))
}

func ErrExpectedType(t Type, val string) error {
	return fmt.Errorf("expected %v for %v", t, Quote(val))
}

func ErrInvalidArg(c Calc, arg any) {
	c.SetError(fmt.Errorf("[%v] invalid argument: %v", c.Op(), Format(arg)))
}

func ErrInvalidFunc(c Calc, fn string, reason string) {
	c.SetError(fmt.Errorf("[%v] invalid function: %v", fn, reason))
}

func ErrModuloByZero(c Calc, a0 any, a1 any) {
	c.SetError(fmt.Errorf("modulo by zero: %v %v", Format(a0), Format(a1)))
}

func ErrNoOpFor(c Calc, op string, types ...Type) {
	var typeNames []string
	for _, t := range types {
		typeNames = append(typeNames, t.String())
	}
	c.SetError(fmt.Errorf("[%v] no operation for %v", op, strings.Join(typeNames, ", ")))
}

func ErrNotEnoughArgs(c Calc, op string, expected int) {
	c.SetError(fmt.Errorf("[%v] not enough arguments, expected %v", op, expected))
}

func ErrUnknownOp(name string) error {
	return fmt.Errorf("unknown operation: %v", name)
}
