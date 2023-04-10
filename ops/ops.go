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

type Func func([]t.Value) ([]t.Value, error)

type FuncDecl struct {
	Params []t.Type
	Fn     Func
}

type Def struct {
	Name string
	NArg int
}

func (o Def) String() string {
	return fmt.Sprintf("%v/%v", o.Name, o.NArg)
}

var (
	Abs   = Def{Name: "abs", NArg: 1}
	Add   = Def{Name: "add", NArg: 2}
	And   = Def{Name: "and", NArg: 2}
	Ceil  = Def{Name: "ceil", NArg: 1}
	Div   = Def{Name: "div", NArg: 2}
	Eq    = Def{Name: "eq", NArg: 2}
	Floor = Def{Name: "floor", NArg: 1}
	Gt    = Def{Name: "gt", NArg: 2}
	Gte   = Def{Name: "gte", NArg: 2}
	Lt    = Def{Name: "lt", NArg: 2}
	Lte   = Def{Name: "lte", NArg: 2}
	Mod   = Def{Name: "mod", NArg: 2}
	Mul   = Def{Name: "mul", NArg: 2}
	Neg   = Def{Name: "neg", NArg: 1}
	Neq   = Def{Name: "neq", NArg: 2}
	Not   = Def{Name: "not", NArg: 1}
	Or    = Def{Name: "or", NArg: 2}
	Pow   = Def{Name: "pow", NArg: 2}
	Rem   = Def{Name: "rem", NArg: 2}
	Sign  = Def{Name: "sign", NArg: 1}
	Sqrt  = Def{Name: "sqrt", NArg: 1}
	Sub   = Def{Name: "sub", NArg: 2}
)

var OpTable = map[Def][]FuncDecl{
	Abs: {
		{Params: []t.Type{t.BigInt}, Fn: absBigInt},
		{Params: []t.Type{t.Decimal}, Fn: absDecimal},
		{Params: []t.Type{t.Float}, Fn: absFloat},
		{Params: []t.Type{t.Rational}, Fn: absRational},
		{Params: []t.Type{t.Complex}, Fn: absComplex},
	},
	Add: {
		{Params: []t.Type{t.BigInt, t.BigInt}, Fn: addBigInt},
		{Params: []t.Type{t.Decimal, t.Decimal}, Fn: addDecimal},
		{Params: []t.Type{t.Float, t.Float}, Fn: addFloat},
		{Params: []t.Type{t.Rational, t.Rational}, Fn: addRational},
		{Params: []t.Type{t.Complex, t.Complex}, Fn: addComplex},
	},
	And: {
		{Params: []t.Type{t.Bool, t.Bool}, Fn: andBool},
		{Params: []t.Type{t.BigInt, t.BigInt}, Fn: andBigInt},
	},
	Ceil: {
		{Params: []t.Type{t.BigInt}, Fn: ceilBigInt},
		{Params: []t.Type{t.Decimal}, Fn: ceilDecimal},
		{Params: []t.Type{t.Float}, Fn: ceilFloat},
	},
	Div: {
		{Params: []t.Type{t.BigInt, t.BigInt}, Fn: divBigInt},
		{Params: []t.Type{t.Decimal, t.Decimal}, Fn: divDecimal},
		{Params: []t.Type{t.Float, t.Float}, Fn: divFloat},
		{Params: []t.Type{t.Rational, t.Rational}, Fn: divRational},
		{Params: []t.Type{t.Complex, t.Complex}, Fn: divComplex},
	},
	Eq: {
		{Params: []t.Type{t.Bool, t.Bool}, Fn: eqBool},
		{Params: []t.Type{t.BigInt, t.BigInt}, Fn: eqBigInt},
		{Params: []t.Type{t.Decimal, t.Decimal}, Fn: eqDecimal},
		{Params: []t.Type{t.Float, t.Float}, Fn: eqFloat},
		{Params: []t.Type{t.Rational, t.Rational}, Fn: eqRational},
		{Params: []t.Type{t.Complex, t.Complex}, Fn: eqComplex},
		{Params: []t.Type{t.String, t.String}, Fn: eqString},
	},
	Floor: {
		{Params: []t.Type{t.BigInt}, Fn: floorBigInt},
		{Params: []t.Type{t.Decimal}, Fn: floorDecimal},
		{Params: []t.Type{t.Float}, Fn: floorFloat},
	},
	Gt: {
		{Params: []t.Type{t.BigInt, t.BigInt}, Fn: gtBigInt},
		{Params: []t.Type{t.Decimal, t.Decimal}, Fn: gtDecimal},
		{Params: []t.Type{t.Float, t.Float}, Fn: gtFloat},
		{Params: []t.Type{t.Rational, t.Rational}, Fn: gtRational},
		{Params: []t.Type{t.String, t.String}, Fn: gtString},
	},
	Gte: {
		{Params: []t.Type{t.BigInt, t.BigInt}, Fn: gteBigInt},
		{Params: []t.Type{t.Decimal, t.Decimal}, Fn: gteDecimal},
		{Params: []t.Type{t.Float, t.Float}, Fn: gteFloat},
		{Params: []t.Type{t.Rational, t.Rational}, Fn: gteRational},
		{Params: []t.Type{t.String, t.String}, Fn: gteString},
	},
	Lt: {
		{Params: []t.Type{t.BigInt, t.BigInt}, Fn: ltBigInt},
		{Params: []t.Type{t.Decimal, t.Decimal}, Fn: ltDecimal},
		{Params: []t.Type{t.Float, t.Float}, Fn: ltFloat},
		{Params: []t.Type{t.Rational, t.Rational}, Fn: ltRational},
		{Params: []t.Type{t.String, t.String}, Fn: ltString},
	},
	Lte: {
		{Params: []t.Type{t.BigInt, t.BigInt}, Fn: lteBigInt},
		{Params: []t.Type{t.Decimal, t.Decimal}, Fn: lteDecimal},
		{Params: []t.Type{t.Float, t.Float}, Fn: lteFloat},
		{Params: []t.Type{t.Rational, t.Rational}, Fn: lteRational},
		{Params: []t.Type{t.String, t.String}, Fn: lteString},
	},
	Mod: {
		{Params: []t.Type{t.BigInt, t.BigInt}, Fn: modBigInt},
		{Params: []t.Type{t.Decimal, t.Decimal}, Fn: modDecimal},
		{Params: []t.Type{t.Float, t.Float}, Fn: modFloat},
	},
	Mul: {
		{Params: []t.Type{t.BigInt, t.BigInt}, Fn: mulBigInt},
		{Params: []t.Type{t.Decimal, t.Decimal}, Fn: mulDecimal},
		{Params: []t.Type{t.Float, t.Float}, Fn: mulFloat},
		{Params: []t.Type{t.Rational, t.Rational}, Fn: mulRational},
		{Params: []t.Type{t.Complex, t.Complex}, Fn: mulComplex},
	},
	Neg: {
		{Params: []t.Type{t.BigInt}, Fn: negBigInt},
		{Params: []t.Type{t.Decimal}, Fn: negDecimal},
		{Params: []t.Type{t.Float}, Fn: negFloat},
		{Params: []t.Type{t.Rational}, Fn: negRational},
	},
	Neq: {
		{Params: []t.Type{t.BigInt, t.BigInt}, Fn: neqBigInt},
		{Params: []t.Type{t.Decimal, t.Decimal}, Fn: neqDecimal},
		{Params: []t.Type{t.Float, t.Float}, Fn: neqFloat},
		{Params: []t.Type{t.Rational, t.Rational}, Fn: neqRational},
		{Params: []t.Type{t.Complex, t.Complex}, Fn: neqComplex},
		{Params: []t.Type{t.String, t.String}, Fn: neqString},
	},
	Not: {
		{Params: []t.Type{t.Bool}, Fn: notBool},
		{Params: []t.Type{t.BigInt}, Fn: notBigInt},
	},
	Or: {
		{Params: []t.Type{t.Bool, t.Bool}, Fn: orBool},
		{Params: []t.Type{t.BigInt, t.BigInt}, Fn: orBigInt},
	},
	Pow: {
		{Params: []t.Type{t.BigInt, t.BigInt}, Fn: powBigInt},
		{Params: []t.Type{t.Float, t.Float}, Fn: powFloat},
		{Params: []t.Type{t.Complex, t.Complex}, Fn: powComplex},
	},
	Rem: {
		{Params: []t.Type{t.BigInt, t.BigInt}, Fn: remBigInt},
		{Params: []t.Type{t.Float, t.Float}, Fn: remFloat},
	},
	Sign: {
		{Params: []t.Type{t.BigInt}, Fn: signBigInt},
		{Params: []t.Type{t.Decimal}, Fn: signDecimal},
		{Params: []t.Type{t.Float}, Fn: signFloat},
		{Params: []t.Type{t.Rational}, Fn: signRational},
	},
	Sqrt: {
		{Params: []t.Type{t.Float}, Fn: sqrtFloat},
	},
	Sub: {
		{Params: []t.Type{t.BigInt, t.BigInt}, Fn: subBigInt},
		{Params: []t.Type{t.Decimal, t.Decimal}, Fn: subDecimal},
		{Params: []t.Type{t.Float, t.Float}, Fn: subFloat},
		{Params: []t.Type{t.Rational, t.Rational}, Fn: subRational},
		{Params: []t.Type{t.Complex, t.Complex}, Fn: subComplex},
	},
}

func Eval(op Def, args []t.Value) ([]t.Value, error) {
	return eval(op, args, false)
}

func eval(op Def, args []t.Value, exact bool) ([]t.Value, error) {
	ops, ok := OpTable[op]
	if !ok {
		return []t.Value{}, fmt.Errorf("no such operation: %v", op.Name)
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
			var callArgs []t.Value
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
	return []t.Value{}, fmt.Errorf("no %v operation for %v", op.Name, strings.Join(types, ", "))
}
