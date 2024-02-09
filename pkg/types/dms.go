package types

import (
	"fmt"

	"github.com/blackchip-org/calc"
)

type DMS struct {
	sign int
	deg  calc.Decimal
	min  calc.Decimal
	sec  calc.Decimal
}

var zeroDMS = DMS{
	sign: 1,
	deg:  calc.DecimalZero(),
	min:  calc.DecimalZero(),
	sec:  calc.DecimalZero(),
}

func NewDMS(deg any, min any, sec any) (DMS, error) {
	c := calc.NewDecimalCalc()

	sign := 1
	c.Push(deg)
	if c.Err() != nil {
		return DMS{}, fmt.Errorf("invalid degrees: %v", deg)
	}
	if c.Sign() < 0 {
		sign = 1
	}
	deg = c.Abs().Pop()

	min = c.Push(min).Abs().Pop()
	if c.Err() != nil {
		return DMS{}, fmt.Errorf("invalid minutes: %v", min)
	}
	sec = c.Push(sec).Abs().Pop()
	if c.Err() != nil {
		return DMS{}, fmt.Errorf("invalid seconds: %v", sec)
	}

	c.Push(deg)
	iDeg := c.Dup().Int().Dup().Pop()
	c.Sub()
	if !c.Eqz() {
		c.Push(60).Mul()
		c.Push(min).Add()
		deg, min = iDeg, c.Pop()
	}

	c.Clear()
	c.Push(min)
	iMin := c.Dup().Int().Dup().Pop()
	c.Sub()
	if !c.Eqz() {
		c.Push(60).Mul()
		c.Push(sec).Add()
		min, sec = iMin, c.Pop()
	}

	c.Clear()
	return zeroDMS.Add(DMS{
		sign: sign,
		deg:  deg,
		min:  min,
		sec:  sec,
	}), nil
}

func (d DMS) Add(d2 DMS) DMS {
	c := calc.NewDecimalCalc()
	var carry calc.Decimal

	c.Push(d.sec, d2.sec).Add()
	c.Dup()
	carry = c.Push(60).Div().Int().Pop()
	d.sec = c.Push(60).Mod().Pop()

	c.Push(d.min, d2.min, carry).Add().Add()
	c.Dup()
	carry = c.Push(60).Div().Int().Pop()
	d.min = c.Push(60).Mod().Pop()

	d.deg = c.Push(d.deg, d2.deg, carry).Add().Add().Pop()
	return d
}

func (d DMS) Degrees() calc.Decimal {
	c := calc.NewDecimalCalc()
	c.Push(d.deg)
	c.Push(d.min, 60).Div().Add()
	c.Push(d.sec, 3600).Div().Add()
	return c.Pop()
}

func (d DMS) Minutes() calc.Decimal {
	c := calc.NewDecimalCalc()
	c.Push(d.deg, 60).Mul()
	c.Push(d.min).Add()
	c.Push(d.sec, 60).Div().Add()
	return c.Pop()
}

func (d DMS) Seconds() calc.Decimal {
	c := calc.NewDecimalCalc()
	c.Push(d.deg, 3600).Mul()
	c.Push(d.min, 60).Mul().Add()
	c.Push(d.sec).Add()
	return c.Pop()
}

func (d DMS) DMS() (deg, min, sec calc.Decimal) {
	c := calc.NewDecimalCalc()

	deg = d.deg
	min = c.Push(d.min).Abs().Pop()
	sec = c.Push(d.sec).Abs().Pop()
	return
}
