package zc

import (
	"fmt"
	"strings"
)

func ErrDivisionByZero(c Calc) {
	c.SetError(fmt.Errorf("division by zero: %v", c.Op()))
}

func ErrExpectedType(t Type, val string) error {
	return fmt.Errorf("expected %v for %v", t, Quote(val))
}

func ErrInfinity(c Calc, sign int) {
	var inf string
	switch {
	case sign < 1:
		inf = "-infinity"
	case sign > 0:
		inf = "+infinity"
	default:
		inf = "infinity"
	}
	c.SetError(fmt.Errorf("%v: %v", inf, c.Op()))
}

func ErrInvalidArgs(c Calc) {
	c.SetError(fmt.Errorf("invalid arguments: %v", c.Op()))
}

func ErrInvalidFunc(c Calc, fn string, reason string) {
	c.SetError(fmt.Errorf("[%v] invalid function: %v", fn, reason))
}

func ErrNoOpFor(c Calc, op string, types ...Type) {
	var typeNames []string
	for _, t := range types {
		typeNames = append(typeNames, t.String())
	}
	c.SetError(fmt.Errorf("[%v] no operation for %v", op, strings.Join(typeNames, ", ")))
}

func ErrNotANumber(c Calc) {
	c.SetError(fmt.Errorf("not a number: %v", c.Op()))
}

func ErrNotEnoughArgs(c Calc, op string, expected int) {
	c.SetError(fmt.Errorf("[%v] not enough arguments, expected %v", op, expected))
}

func ErrUnknownOp(name string) error {
	return fmt.Errorf("unknown operation: %v", name)
}
