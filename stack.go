package zc

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/blackchip-org/zc/ops"
	"github.com/blackchip-org/zc/types"
	"github.com/shopspring/decimal"
)

type Stack struct {
	Name string // FIXME: should remove this
	data []string
}

func NewStack(calc *Calc, name string) *Stack {
	return &Stack{Name: name}
}

func (s *Stack) Copy() *Stack {
	ns := &Stack{
		Name: s.Name,
		data: make([]string, len(s.data)),
	}
	copy(ns.data, s.data)
	return ns
}

func (s *Stack) Equal(os *Stack) bool {
	if os == nil {
		return false
	}
	if s.Name != os.Name {
		return false
	}
	if len(s.data) != len(os.data) {
		return false
	}
	for i := 0; i < len(s.data); i++ {
		if s.data[i] != os.data[i] {
			return false
		}
	}
	return true
}

func (s *Stack) Items() []string {
	items := make([]string, len(s.data))
	copy(items, s.data)
	return items
}

func (s *Stack) ItemsReversed() []string {
	items := make([]string, len(s.data))
	for i := 0; i < len(s.data); i++ {
		items[i] = s.data[len(s.data)-i-1]
	}
	return items
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Push(v string) {
	s.data = append(s.data, v)
}

func (s *Stack) PushAll(xs []string) {
	s.data = append(s.data, xs...)
}

func (s *Stack) Enqueue(v string) {
	s.data = append([]string{v}, s.data...)
}

func (s *Stack) Clear() *Stack {
	s.data = nil
	return s
}

func (s *Stack) String() string {
	var sb strings.Builder
	for i, item := range s.data {
		if i != 0 {
			sb.WriteString(" | ")
		}
		sb.WriteString(item)
	}
	return sb.String()
}

func (s *Stack) Peek() (string, error) {
	n := len(s.data)
	if n == 0 {
		return "", fmt.Errorf("%v: stack empty", s.Name)
	}
	return s.data[n-1], nil
}

func (s *Stack) Peek2() (string, string, error) {
	n := len(s.data)
	if n < 2 {
		return "", "", fmt.Errorf("%v: stack empty", s.Name)
	}
	return s.data[n-2], s.data[n-1], nil
}

func (s *Stack) Pop() (string, error) {
	n := len(s.data)
	if n == 0 {
		return "", fmt.Errorf("%v: stack empty", s.Name)
	}
	var top string
	top, s.data = s.data[n-1], s.data[:n-1]
	return top, nil
}

func (s *Stack) Pop2() (string, string, error) {
	b, err := s.Pop()
	if err != nil {
		return "", "", err
	}
	a, err := s.Pop()
	if err != nil {
		return "", "", err
	}
	return a, b, nil
}

func (s *Stack) PopBigInt() (*big.Int, error) {
	v, err := s.Pop()
	if err != nil {
		return nil, err
	}
	return types.BigInt.Parse(v)
}

func (s *Stack) PopBigInt2() (*big.Int, *big.Int, error) {
	b, err := s.PopBigInt()
	if err != nil {
		return nil, nil, err
	}
	a, err := s.PopBigInt()
	if err != nil {
		return nil, nil, err
	}
	return a, b, nil
}

func (s *Stack) PopBool() (bool, error) {
	v, err := s.Pop()
	if err != nil {
		return false, err
	}
	return types.Bool.Parse(v)
}

func (s *Stack) PopBool2() (bool, bool, error) {
	b, err := s.PopBool()
	if err != nil {
		return false, false, err
	}
	a, err := s.PopBool()
	if err != nil {
		return false, false, err
	}
	return a, b, nil
}

func (s *Stack) PopDecimal() (decimal.Decimal, error) {
	v, err := s.Pop()
	if err != nil {
		return decimal.Zero, err
	}
	return types.Decimal.Parse(v)
}

func (s *Stack) PopDecimal2() (decimal.Decimal, decimal.Decimal, error) {
	b, err := s.PopDecimal()
	if err != nil {
		return decimal.Zero, decimal.Zero, err
	}
	a, err := s.PopDecimal()
	if err != nil {
		return decimal.Zero, decimal.Zero, err
	}
	return a, b, nil
}

func (s *Stack) PopFloat() (float64, error) {
	v, err := s.Pop()
	if err != nil {
		return 0, err
	}
	return types.Float.Parse(v)
}

func (s *Stack) PopInt() (int, error) {
	v, err := s.Pop()
	if err != nil {
		return 0, err
	}
	return types.Int.Parse(v)
}

func (s *Stack) PopInt32() (int32, error) {
	v, err := s.Pop()
	if err != nil {
		return 0, err
	}
	return types.Int32.Parse(v)
}

func (s *Stack) PopInt64() (int64, error) {
	v, err := s.Pop()
	if err != nil {
		return 0, err
	}
	return types.Int64.Parse(v)
}

func (s *Stack) PopRune() (rune, error) {
	v, err := s.Pop()
	if err != nil {
		return 0, err
	}
	return types.Rune.Parse(v)
}

func (s *Stack) PopUint() (uint, error) {
	v, err := s.Pop()
	if err != nil {
		return 0, err
	}
	i, err := types.Uint.Parse(v)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (s *Stack) PopUint8() (uint8, error) {
	v, err := s.Pop()
	if err != nil {
		return 0, err
	}
	i, err := types.Uint8.Parse(v)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (s *Stack) PushBigInt(v *big.Int) {
	s.Push(types.BigInt.Format(v))
}

func (s *Stack) PushBool(v bool) {
	s.Push(types.Bool.Format(v))
}

func (s *Stack) PushDecimal(v decimal.Decimal) {
	s.Push(types.Decimal.Format(v))
}

func (s *Stack) PushFloat(v float64) {
	s.Push(types.Float.Format(v))
}

func (s *Stack) PushInt(v int) {
	s.Push(types.Int.Format(v))
}

func (s *Stack) PushInt32(v int32) {
	s.Push(types.Int32.Format(v))
}

func (s *Stack) PushInt64(v int64) {
	s.Push(types.Int64.Format(v))
}

func (s *Stack) PushRune(r rune) {
	s.Push(string(r))
}

func (s *Stack) PushUint(v uint) {
	s.Push(types.Uint.Format(v))
}

func (s *Stack) PushUint8(v uint8) {
	s.Push(types.Uint8.Format(v))
}

func (s *Stack) PushUint32(v uint32) {
	s.Push(types.Uint32.Format(v))
}

func (s *Stack) PushUint64(v uint64) {
	s.Push(types.Uint64.Format(v))
}

func (s *Stack) PushValue(v string) {
	s.Push(v)
}

func (s *Stack) MustPop() string {
	val, err := s.Pop()
	if err != nil {
		panic(err)
	}
	return val
}

// FIXME this is a hack for now
func (s *Stack) Eval(op ops.Def) error {
	var args []types.Generic
	for i := 0; i < op.NArg; i++ {
		s, err := s.Pop()
		if err != nil {
			return err
		}
		args = append([]types.Generic{types.Parse(s)}, args...)
	}
	result, err := ops.Eval(op, args)
	if err != nil {
		return err
	}
	for _, r := range result {
		s.Push(r.Format())
	}
	return nil
}
