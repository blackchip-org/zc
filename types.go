package zc

import (
	"fmt"
	"math/big"
	"strconv"
)

var (
	Any    = AnyType{}
	BigInt = BigIntType{}
	Int    = IntType{}
	String = StringType{}
)

var (
	poolSize = 8
	//decPool  = newPool[apd.Decimal](poolSize)
	intPool = newPool[big.Int](poolSize)
	//ratPool  = newPool[big.Rat](poolSize)
)

// ----------------------------------------------------------------------------
type AnyType struct{}

func (t AnyType) AppName() string { return "Any" }
func (t AnyType) GoName() string  { return "any" }

func (t AnyType) Parse(str string) (any, bool) {
	return str, true
}

func (t AnyType) Format(a any) string {
	return fmt.Sprint(a)
}

// ----------------------------------------------------------------------------
type BigIntType struct{}

func (t BigIntType) AppName() string { return "Int" }
func (t BigIntType) GoName() string  { return "*big.Int" }

func (t BigIntType) New() *big.Int {
	return intPool.New()
}

func (t BigIntType) Recycle(vals ...*big.Int) {
	for _, val := range vals {
		intPool.Recycle(val)
	}
}

func (t BigIntType) As(a any) *big.Int {
	v, ok := a.(*big.Int)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t BigIntType) Push(c Calc, val *big.Int) {
	c.Push(Item{TypeVal: val, Type: t})
}

func (t BigIntType) Pop(c Calc) *big.Int {
	return t.As(c.Pop().TypeVal)
}

func (t BigIntType) Parse(str string) (any, bool) {
	str = PreParseNumber(str)
	v := t.New()
	v, ok := v.SetString(str, 0)
	if !ok {
		t.Recycle(v)
		return nil, false
	}
	return v, true
}

func (t BigIntType) Format(a any) string {
	v := t.As(a)
	return v.String()
}

// ----------------------------------------------------------------------------
type IntType struct{}

func (t IntType) AppName() string { return "Int/s" }
func (t IntType) GoName() string  { return "int" }

func (t IntType) As(a any) int {
	v, ok := a.(int)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t IntType) Push(c Calc, val int) {
	c.Push(Item{TypeVal: val, Type: t})
}

func (t IntType) Pop(c Calc) int {
	return t.As(c.Pop().TypeVal)
}

func (t IntType) Parse(str string) (any, bool) {
	str = PreParseNumber(str)
	i64, err := strconv.ParseInt(str, 0, 0)
	return int(i64), err == nil
}

func (t IntType) Format(a any) string {
	v := t.As(a)
	return strconv.FormatInt(int64(v), 10)
}

// ----------------------------------------------------------------------------
type StringType struct{}

func (t StringType) AppName() string { return "Text" }
func (t StringType) GoName() string  { return "string" }

func (t StringType) As(a any) string {
	v, ok := a.(string)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t StringType) Push(c Calc, val string) {
	c.Push(Item{TypeVal: val, Type: t})

}

func (t StringType) Pop(c Calc) string {
	return t.As(c.Pop().TypeVal)
}

func (t StringType) Parse(str string) (any, bool) {
	return str, true
}

func (t StringType) Format(a any) string {
	return t.As(a)
}
