package zc

import (
	"fmt"
	"strings"
)

func ErrDivisionByZero(c Calc) {
	c.SetError(fmt.Errorf("division by zero: %v", c.Op()))
}

func ErrExpectedType(c Calc, t Type, val string) {
	c.SetError(fmt.Errorf("%v: expected %v for %v", c.Op().Name, t, Quote(val)))
}

func PanicExpectedType(t Type, val string) {
	panic(fmt.Errorf("expected %v for %v", t, Quote(val)))
}

func ErrFeatureNotSupported(name string) error {
	return fmt.Errorf("feature not supported: %v", name)
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
	c.SetError(fmt.Errorf("function %s is invalid: %v", Quote(fn), reason))
}

func ErrNoOpFor(c Calc, name string, args []string) {
	var qArgs []string
	for _, arg := range args {
		qArgs = append(qArgs, Quote(arg))
	}
	c.SetError(fmt.Errorf("no operation for: %v %v", strings.Join(qArgs, " "), name))
}

func ErrNotANumber(c Calc) {
	c.SetError(fmt.Errorf("not a number: %v", c.Op()))
}

func ErrNotEnoughArgs(c Calc, op string, expected int) {
	c.SetError(fmt.Errorf("%v: not enough arguments, expected %v", op, expected))
}

func ErrUnknownOp(name string) error {
	return fmt.Errorf("unknown operation: %v", name)
}
