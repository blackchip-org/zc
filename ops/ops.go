package ops

import (
	"errors"
	"fmt"
	"strings"

	t "github.com/blackchip-org/zc/types"
)

var (
	ErrDivisionByZero = errors.New("division by zero")
)

type OpFn func([]t.Generic) ([]t.Generic, error)

type Op struct {
	Params []t.GenericType
	Fn     OpFn
}

const (
	Abs   = "abs"
	Add   = "add"
	And   = "and"
	Ceil  = "ceil"
	Div   = "div"
	Eq    = "eq"
	Floor = "floor"
	Gt    = "gt"
	Gte   = "gte"
	Lt    = "lt"
	Lte   = "lte"
	Mod   = "mod"
	Mul   = "mul"
	Neg   = "neg"
	Neq   = "neq"
	Not   = "not"
	Or    = "or"
	Pow   = "pow"
	Rem   = "rem"
	Sign  = "sign"
	Sqrt  = "sqrt"
	Sub   = "sub"
)

var OpTable map[string][]Op = map[string][]Op{
	Abs: {
		{Params: []t.GenericType{t.BigInt}, Fn: absBigInt},
		{Params: []t.GenericType{t.Decimal}, Fn: absDecimal},
		{Params: []t.GenericType{t.Float}, Fn: absFloat},
		{Params: []t.GenericType{t.Rational}, Fn: absRational},
		{Params: []t.GenericType{t.Complex}, Fn: absComplex},
	},
	Add: {
		{Params: []t.GenericType{t.BigInt, t.BigInt}, Fn: addBigInt},
		{Params: []t.GenericType{t.Decimal, t.Decimal}, Fn: addDecimal},
		{Params: []t.GenericType{t.Float, t.Float}, Fn: addFloat},
		{Params: []t.GenericType{t.Rational, t.Rational}, Fn: addRational},
		{Params: []t.GenericType{t.Complex, t.Complex}, Fn: addComplex},
	},
	And: {
		{Params: []t.GenericType{t.Bool, t.Bool}, Fn: andBool},
		{Params: []t.GenericType{t.BigInt, t.BigInt}, Fn: andBigInt},
	},
	Ceil: {
		{Params: []t.GenericType{t.BigInt}, Fn: ceilBigInt},
		{Params: []t.GenericType{t.Decimal}, Fn: ceilDecimal},
		{Params: []t.GenericType{t.Float}, Fn: ceilFloat},
	},
	Div: {
		{Params: []t.GenericType{t.BigInt, t.BigInt}, Fn: divBigInt},
		{Params: []t.GenericType{t.Decimal, t.Decimal}, Fn: divDecimal},
		{Params: []t.GenericType{t.Float, t.Float}, Fn: divFloat},
		{Params: []t.GenericType{t.Rational, t.Rational}, Fn: divRational},
		{Params: []t.GenericType{t.Complex, t.Complex}, Fn: divComplex},
	},
	Eq: {
		{Params: []t.GenericType{t.Bool, t.Bool}, Fn: eqBool},
		{Params: []t.GenericType{t.BigInt, t.BigInt}, Fn: eqBigInt},
		{Params: []t.GenericType{t.Decimal, t.Decimal}, Fn: eqDecimal},
		{Params: []t.GenericType{t.Float, t.Float}, Fn: eqFloat},
		{Params: []t.GenericType{t.Rational, t.Rational}, Fn: eqRational},
		{Params: []t.GenericType{t.Complex, t.Complex}, Fn: eqComplex},
		{Params: []t.GenericType{t.String, t.String}, Fn: eqString},
	},
	Floor: {
		{Params: []t.GenericType{t.BigInt}, Fn: floorBigInt},
		{Params: []t.GenericType{t.Decimal}, Fn: floorDecimal},
		{Params: []t.GenericType{t.Float}, Fn: floorFloat},
	},
	Gt: {
		{Params: []t.GenericType{t.BigInt, t.BigInt}, Fn: gtBigInt},
		{Params: []t.GenericType{t.Decimal, t.Decimal}, Fn: gtDecimal},
		{Params: []t.GenericType{t.Float, t.Float}, Fn: gtFloat},
		{Params: []t.GenericType{t.Rational, t.Rational}, Fn: gtRational},
		{Params: []t.GenericType{t.String, t.String}, Fn: gtString},
	},
	Gte: {
		{Params: []t.GenericType{t.BigInt, t.BigInt}, Fn: gteBigInt},
		{Params: []t.GenericType{t.Decimal, t.Decimal}, Fn: gteDecimal},
		{Params: []t.GenericType{t.Float, t.Float}, Fn: gteFloat},
		{Params: []t.GenericType{t.Rational, t.Rational}, Fn: gteRational},
		{Params: []t.GenericType{t.String, t.String}, Fn: gteString},
	},
	Lt: {
		{Params: []t.GenericType{t.BigInt, t.BigInt}, Fn: ltBigInt},
		{Params: []t.GenericType{t.Decimal, t.Decimal}, Fn: ltDecimal},
		{Params: []t.GenericType{t.Float, t.Float}, Fn: ltFloat},
		{Params: []t.GenericType{t.Rational, t.Rational}, Fn: ltRational},
		{Params: []t.GenericType{t.String, t.String}, Fn: ltString},
	},
	Lte: {
		{Params: []t.GenericType{t.BigInt, t.BigInt}, Fn: lteBigInt},
		{Params: []t.GenericType{t.Decimal, t.Decimal}, Fn: lteDecimal},
		{Params: []t.GenericType{t.Float, t.Float}, Fn: lteFloat},
		{Params: []t.GenericType{t.Rational, t.Rational}, Fn: lteRational},
		{Params: []t.GenericType{t.String, t.String}, Fn: lteString},
	},
	Mod: {
		{Params: []t.GenericType{t.BigInt, t.BigInt}, Fn: modBigInt},
		{Params: []t.GenericType{t.Decimal, t.Decimal}, Fn: modDecimal},
		{Params: []t.GenericType{t.Float, t.Float}, Fn: modFloat},
	},
	Mul: {
		{Params: []t.GenericType{t.BigInt, t.BigInt}, Fn: mulBigInt},
		{Params: []t.GenericType{t.Decimal, t.Decimal}, Fn: mulDecimal},
		{Params: []t.GenericType{t.Float, t.Float}, Fn: mulFloat},
		{Params: []t.GenericType{t.Rational, t.Rational}, Fn: mulRational},
		{Params: []t.GenericType{t.Complex, t.Complex}, Fn: mulComplex},
	},
	Neg: {
		{Params: []t.GenericType{t.BigInt}, Fn: negBigInt},
		{Params: []t.GenericType{t.Decimal}, Fn: negDecimal},
		{Params: []t.GenericType{t.Float}, Fn: negFloat},
		{Params: []t.GenericType{t.Rational}, Fn: negRational},
	},
	Neq: {
		{Params: []t.GenericType{t.BigInt, t.BigInt}, Fn: neqBigInt},
		{Params: []t.GenericType{t.Decimal, t.Decimal}, Fn: neqDecimal},
		{Params: []t.GenericType{t.Float, t.Float}, Fn: neqFloat},
		{Params: []t.GenericType{t.Rational, t.Rational}, Fn: neqRational},
		{Params: []t.GenericType{t.Complex, t.Complex}, Fn: neqComplex},
		{Params: []t.GenericType{t.String, t.String}, Fn: neqString},
	},
	Not: {
		{Params: []t.GenericType{t.Bool}, Fn: notBool},
		{Params: []t.GenericType{t.BigInt}, Fn: notBigInt},
	},
	Or: {
		{Params: []t.GenericType{t.Bool, t.Bool}, Fn: orBool},
		{Params: []t.GenericType{t.BigInt, t.BigInt}, Fn: orBigInt},
	},
	Pow: {
		{Params: []t.GenericType{t.BigInt, t.BigInt}, Fn: powBigInt},
		{Params: []t.GenericType{t.Float, t.Float}, Fn: powFloat},
		{Params: []t.GenericType{t.Complex, t.Complex}, Fn: powComplex},
	},
	Rem: {
		{Params: []t.GenericType{t.BigInt, t.BigInt}, Fn: remBigInt},
		{Params: []t.GenericType{t.Float, t.Float}, Fn: remFloat},
	},
	Sign: {
		{Params: []t.GenericType{t.BigInt}, Fn: signBigInt},
		{Params: []t.GenericType{t.Decimal}, Fn: signDecimal},
		{Params: []t.GenericType{t.Float}, Fn: signFloat},
		{Params: []t.GenericType{t.Rational}, Fn: signRational},
	},
	Sqrt: {
		{Params: []t.GenericType{t.Float}, Fn: sqrtFloat},
	},
	Sub: {
		{Params: []t.GenericType{t.BigInt, t.BigInt}, Fn: subBigInt},
		{Params: []t.GenericType{t.Decimal, t.Decimal}, Fn: subDecimal},
		{Params: []t.GenericType{t.Float, t.Float}, Fn: subFloat},
		{Params: []t.GenericType{t.Rational, t.Rational}, Fn: subRational},
		{Params: []t.GenericType{t.Complex, t.Complex}, Fn: subComplex},
	},
}

func Eval(opName string, args []t.Generic) ([]t.Generic, error) {
	return eval(opName, args, false)
}

func eval(name string, args []t.Generic, exact bool) ([]t.Generic, error) {
	ops, ok := OpTable[name]
	if !ok {
		return []t.Generic{}, fmt.Errorf("no such operation: %v", name)
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
			var callArgs []t.Generic
			for i, arg := range args {
				callArg, err := t.To(arg, op.Params[i])
				if err != nil {
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
	return []t.Generic{}, fmt.Errorf("no %v operation for %v", name, strings.Join(types, ", "))
}
