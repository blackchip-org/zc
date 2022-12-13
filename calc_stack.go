package zc

import (
	"fmt"
	"math/big"

	"github.com/shopspring/decimal"
)

func (c *Calc) Peek2() (string, string, error) {
	items := c.Stack.Items()
	n := len(items)
	if n < 2 {
		return "", "", fmt.Errorf("%v: stack empty", c.Stack.Name)
	}
	return items[n-2], items[n-1], nil
}

func (c *Calc) Pop2() (string, string, error) {
	b, err := c.Stack.Pop()
	if err != nil {
		return "", "", err
	}
	a, err := c.Stack.Pop()
	if err != nil {
		return "", "", err
	}
	return a, b, nil
}

func (c *Calc) PopBigInt() (*big.Int, error) {
	v, err := c.Stack.Pop()
	if err != nil {
		return nil, err
	}
	r, err := c.ParseBigInt(v)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Calc) PopBigInt2() (*big.Int, *big.Int, error) {
	b, err := c.PopBigInt()
	if err != nil {
		return nil, nil, err
	}
	a, err := c.PopBigInt()
	if err != nil {
		return nil, nil, err
	}
	return a, b, nil
}

func (c *Calc) PopBool() (bool, error) {
	v, err := c.Stack.Pop()
	if err != nil {
		return false, err
	}
	b, err := c.ParseBool(v)
	if err != nil {
		return false, err
	}
	return b, nil
}

func (c *Calc) PopBool2() (bool, bool, error) {
	b, err := c.PopBool()
	if err != nil {
		return false, false, err
	}
	a, err := c.PopBool()
	if err != nil {
		return false, false, err
	}
	return a, b, nil
}

func (c *Calc) PopFix() (decimal.Decimal, error) {
	v, err := c.Stack.Pop()
	if err != nil {
		return decimal.Zero, err
	}
	d, err := c.ParseDecimal(v)
	if err != nil {
		return decimal.Zero, err
	}
	return d, err
}

func (c *Calc) PopFix2() (decimal.Decimal, decimal.Decimal, error) {
	b, err := c.PopFix()
	if err != nil {
		return decimal.Zero, decimal.Zero, err
	}
	a, err := c.PopFix()
	if err != nil {
		return decimal.Zero, decimal.Zero, err
	}
	return a, b, nil
}

func (c *Calc) PopInt() (int, error) {
	v, err := c.Stack.Pop()
	if err != nil {
		return 0, err
	}
	i, err := c.ParseInt(v)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (c *Calc) PopInt32() (int32, error) {
	v, err := c.Stack.Pop()
	if err != nil {
		return 0, err
	}
	i, err := c.ParseInt32(v)
	if err != nil {
		return 0, err
	}
	return i, nil
}
