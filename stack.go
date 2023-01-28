package zc

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/shopspring/decimal"
)

type Stack struct {
	Name string // FIXME: should remove this
	calc *Calc
	data []string
}

func NewStack(calc *Calc, name string) *Stack {
	return &Stack{Name: name, calc: calc}
}

func (s *Stack) Items() []string {
	items := make([]string, len(s.data))
	copy(items, s.data)
	return items
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Push(v string) {
	s.data = append(s.data, v)
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
	if len(s.data) == 0 {
		return "", fmt.Errorf("%v: stack empty", s.Name)
	}
	return s.data[0], nil
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
	r, err := s.calc.ParseBigInt(v)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *Stack) PopBigIntWithAttrs() (*big.Int, FormatAttrs, error) {
	v, err := s.Pop()
	if err != nil {
		return nil, FormatAttrs{}, err
	}
	r, err := s.calc.ParseBigInt(v)
	if err != nil {
		return nil, FormatAttrs{}, err
	}
	return r, ParseFormatAttrs(v), nil
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
	b, err := s.calc.ParseBool(v)
	if err != nil {
		return false, err
	}
	return b, nil
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

func (s *Stack) PopFixed() (decimal.Decimal, error) {
	v, err := s.Pop()
	if err != nil {
		return decimal.Zero, err
	}
	d, err := s.calc.ParseFixed(v)
	if err != nil {
		return decimal.Zero, err
	}
	return d, err
}

func (s *Stack) PopFixed2() (decimal.Decimal, decimal.Decimal, error) {
	b, err := s.PopFixed()
	if err != nil {
		return decimal.Zero, decimal.Zero, err
	}
	a, err := s.PopFixed()
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
	f, err := s.calc.ParseFloat(v)
	if err != nil {
		return 0, err
	}
	return f, err
}

func (s *Stack) PopInt() (int, error) {
	v, err := s.Pop()
	if err != nil {
		return 0, err
	}
	i, err := s.calc.ParseInt(v)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (s *Stack) PopInt32() (int32, error) {
	v, err := s.Pop()
	if err != nil {
		return 0, err
	}
	i, err := s.calc.ParseInt32(v)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (s *Stack) PopUint() (uint, error) {
	v, err := s.Pop()
	if err != nil {
		return 0, err
	}
	i, err := s.calc.ParseUint(v)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (s *Stack) PushBigInt(v *big.Int) {
	s.Push(s.calc.FormatBigInt(v))
}

func (s *Stack) PushBigIntWithAttrs(v *big.Int, attrs FormatAttrs) {
	s.Push(s.calc.FormatBigIntWithAttrs(v, attrs))
}

func (s *Stack) PushBool(v bool) {
	s.Push(s.calc.FormatBool(v))
}

func (s *Stack) PushFixed(v decimal.Decimal) {
	s.Push(s.calc.FormatFixed(v))
}

func (s *Stack) PushFixedWithAttrs(v decimal.Decimal, attrs FormatAttrs) {
	s.Push(s.calc.FormatFixedWithAttrs(v, attrs))
}

func (s *Stack) PushFloat(v float64) {
	s.Push(s.calc.FormatFloat(v))
}

func (s *Stack) PushInt(v int) {
	s.Push(s.calc.FormatInt(v))
}

func (s *Stack) PushInt32(v int32) {
	s.Push(s.calc.FormatInt32(v))
}

func (s *Stack) PushInt64(v int64) {
	s.Push(s.calc.FormatInt64(v))
}

func (s *Stack) PushUint(v uint) {
	s.Push(s.calc.FormatUint(v))
}

func (s *Stack) PushUint32(v uint32) {
	s.Push(s.calc.FormatUint32(v))
}

func (s *Stack) PushUint64(v uint64) {
	s.Push(s.calc.FormatUint64(v))
}

func (s *Stack) PushValue(v string) {
	s.Push(s.calc.FormatValue(v))
}

func (s *Stack) MustPop() string {
	val, err := s.Pop()
	if err != nil {
		panic(err)
	}
	return val
}
