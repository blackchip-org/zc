package types

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrDivisionByZero = errors.New("division by zero")
)

type OpFn func([]Generic) ([]Generic, error)

type Op struct {
	Params []Type
	Fn     OpFn
}

var OpTable map[string][]Op = map[string][]Op{
	"abs": {
		{Params: []Type{BigInt}, Fn: absBigInt},
		{Params: []Type{Decimal}, Fn: absDecimal},
		{Params: []Type{Float}, Fn: absFloat},
		{Params: []Type{Rational}, Fn: absRational},
		{Params: []Type{Complex}, Fn: absComplex},
	},
	"add": {
		{Params: []Type{BigInt, BigInt}, Fn: addBigInt},
		{Params: []Type{Decimal, Decimal}, Fn: addDecimal},
		{Params: []Type{Float, Float}, Fn: addFloat},
		{Params: []Type{Rational, Rational}, Fn: addRational},
		{Params: []Type{Complex, Complex}, Fn: addComplex},
	},
	"div": {
		{Params: []Type{BigInt, BigInt}, Fn: divBigInt},
		{Params: []Type{Decimal, Decimal}, Fn: divDecimal},
		{Params: []Type{Float, Float}, Fn: divFloat},
		{Params: []Type{Rational, Rational}, Fn: divRational},
		{Params: []Type{Complex, Complex}, Fn: divComplex},
	},
	"mul": {
		{Params: []Type{BigInt, BigInt}, Fn: mulBigInt},
		{Params: []Type{Decimal, Decimal}, Fn: mulDecimal},
		{Params: []Type{Float, Float}, Fn: mulFloat},
		{Params: []Type{Rational, Rational}, Fn: mulRational},
		{Params: []Type{Complex, Complex}, Fn: mulComplex},
	},
	"sub": {
		{Params: []Type{BigInt, BigInt}, Fn: subBigInt},
		{Params: []Type{Decimal, Decimal}, Fn: subDecimal},
		{Params: []Type{Float, Float}, Fn: subFloat},
		{Params: []Type{Rational, Rational}, Fn: subRational},
		{Params: []Type{Complex, Complex}, Fn: subComplex},
	},
}

func Eval(opName string, args []Generic) ([]Generic, error) {
	return eval(opName, args, false)
}

func eval(name string, args []Generic, exact bool) ([]Generic, error) {
	ops, ok := OpTable[name]
	if !ok {
		return []Generic{}, fmt.Errorf("no such operation: %v", name)
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
			var callArgs []Generic
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
	return []Generic{}, fmt.Errorf("no %v operation for %v", name, strings.Join(types, ", "))
}
