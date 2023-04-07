package types

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrDivisionByZero = errors.New("division by zero")
	ErrNaN            = errors.New("not a number")
)

type OpFn func([]Value) ([]Value, error)

type Op struct {
	Params []Type
	Fn     OpFn
}

var OpTable map[string][]Op = map[string][]Op{
	"add": {
		{Params: []Type{BigInt, BigInt}, Fn: addBigInt},
		{Params: []Type{Decimal, Decimal}, Fn: addDecimal},
		{Params: []Type{Complex, Complex}, Fn: addComplex},
	},
	"div": {
		{Params: []Type{BigInt, BigInt}, Fn: divBigInt},
		{Params: []Type{Decimal, Decimal}, Fn: divDecimal},
		{Params: []Type{Complex, Complex}, Fn: divComplex},
	},
	"mul": {
		{Params: []Type{BigInt, BigInt}, Fn: mulBigInt},
		{Params: []Type{Decimal, Decimal}, Fn: mulDecimal},
		{Params: []Type{Complex, Complex}, Fn: mulComplex},
	},
	"sub": {
		{Params: []Type{BigInt, BigInt}, Fn: subBigInt},
		{Params: []Type{Decimal, Decimal}, Fn: subDecimal},
		{Params: []Type{Complex, Complex}, Fn: subComplex},
	},
}

func Eval(opName string, args []Value) ([]Value, error) {
	return eval(opName, args, false)
}

func eval(name string, args []Value, exact bool) ([]Value, error) {
	ops, ok := OpTable[name]
	if !ok {
		return []Value{}, fmt.Errorf("no such operation: %v", name)
	}
	for _, op := range ops {
		if len(op.Params) != len(args) {
			continue
		}
		valid := true
		for i, arg := range args {
			if arg.Type() != op.Params[i] {
				valid = false
				break
			}
		}
		if valid {
			vs, err := op.Fn(args)
			return vs, err
		}
	}

	if !exact {
		for _, op := range ops {
			if len(op.Params) != len(args) {
				continue
			}
			valid := true
			var callArgs []Value
			for i, arg := range args {
				callArg, ok := To(arg, op.Params[i])
				if !ok {
					valid = false
					break
				}
				callArgs = append(callArgs, callArg)
			}
			if valid {
				vs, err := op.Fn(callArgs)
				return vs, err
			}
		}
	}

	var types []string
	for _, arg := range args {
		types = append(types, arg.Type().String())
	}
	return []Value{}, fmt.Errorf("no %v operation for %v", name, strings.Join(types, ", "))
}
