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
	Params []GenericType
	Fn     OpFn
}

const (
	OpAbs   = "abs"
	OpAdd   = "add"
	OpCeil  = "ceil"
	OpDiv   = "div"
	OpFloor = "floor"
	OpMod   = "mod"
	OpMul   = "mul"
	OpNeg   = "neg"
	OpPow   = "pow"
	OpRem   = "rem"
	OpSign  = "sign"
	OpSub   = "sub"
)

var OpTable map[string][]Op = map[string][]Op{
	OpAbs: {
		{Params: []GenericType{BigInt}, Fn: absBigInt},
		{Params: []GenericType{Decimal}, Fn: absDecimal},
		{Params: []GenericType{Float}, Fn: absFloat},
		{Params: []GenericType{Rational}, Fn: absRational},
		{Params: []GenericType{Complex}, Fn: absComplex},
	},
	OpAdd: {
		{Params: []GenericType{BigInt, BigInt}, Fn: addBigInt},
		{Params: []GenericType{Decimal, Decimal}, Fn: addDecimal},
		{Params: []GenericType{Float, Float}, Fn: addFloat},
		{Params: []GenericType{Rational, Rational}, Fn: addRational},
		{Params: []GenericType{Complex, Complex}, Fn: addComplex},
	},
	OpCeil: {
		{Params: []GenericType{BigInt}, Fn: ceilBigInt},
		{Params: []GenericType{Decimal}, Fn: ceilDecimal},
		{Params: []GenericType{Float}, Fn: ceilFloat},
	},
	OpDiv: {
		{Params: []GenericType{BigInt, BigInt}, Fn: divBigInt},
		{Params: []GenericType{Decimal, Decimal}, Fn: divDecimal},
		{Params: []GenericType{Float, Float}, Fn: divFloat},
		{Params: []GenericType{Rational, Rational}, Fn: divRational},
		{Params: []GenericType{Complex, Complex}, Fn: divComplex},
	},
	OpFloor: {
		{Params: []GenericType{BigInt}, Fn: floorBigInt},
		{Params: []GenericType{Decimal}, Fn: floorDecimal},
		{Params: []GenericType{Float}, Fn: floorFloat},
	},
	OpMod: {
		{Params: []GenericType{BigInt, BigInt}, Fn: modBigInt},
		{Params: []GenericType{Decimal, Decimal}, Fn: modDecimal},
		{Params: []GenericType{Float, Float}, Fn: modFloat},
	},
	OpMul: {
		{Params: []GenericType{BigInt, BigInt}, Fn: mulBigInt},
		{Params: []GenericType{Decimal, Decimal}, Fn: mulDecimal},
		{Params: []GenericType{Float, Float}, Fn: mulFloat},
		{Params: []GenericType{Rational, Rational}, Fn: mulRational},
		{Params: []GenericType{Complex, Complex}, Fn: mulComplex},
	},
	OpNeg: {
		{Params: []GenericType{BigInt}, Fn: negBigInt},
		{Params: []GenericType{Decimal}, Fn: negDecimal},
		{Params: []GenericType{Float}, Fn: negFloat},
		{Params: []GenericType{Rational}, Fn: negRational},
	},
	OpPow: {
		{Params: []GenericType{BigInt, BigInt}, Fn: powBigInt},
		{Params: []GenericType{Float, Float}, Fn: powFloat},
		{Params: []GenericType{Complex, Complex}, Fn: powComplex},
	},
	OpRem: {
		{Params: []GenericType{BigInt, BigInt}, Fn: remBigInt},
		{Params: []GenericType{Float, Float}, Fn: remFloat},
	},
	OpSign: {
		{Params: []GenericType{BigInt}, Fn: signBigInt},
		{Params: []GenericType{Decimal}, Fn: signDecimal},
		{Params: []GenericType{Float}, Fn: signFloat},
		{Params: []GenericType{Rational}, Fn: signRational},
	},
	OpSub: {
		{Params: []GenericType{BigInt, BigInt}, Fn: subBigInt},
		{Params: []GenericType{Decimal, Decimal}, Fn: subDecimal},
		{Params: []GenericType{Float, Float}, Fn: subFloat},
		{Params: []GenericType{Rational, Rational}, Fn: subRational},
		{Params: []GenericType{Complex, Complex}, Fn: subComplex},
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
