package funcs

import (
	"fmt"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/vars"
	"github.com/cockroachdb/apd/v3"
)

func RoundComplex(c zc.Calc) {
	var zero apd.Decimal
	d := vars.ForDec(c).Math
	p := zc.Int32.Pop(c)
	x := zc.Complex.Pop(c)

	if p < 0 {
		c.Raise(zc.ErrInvalidArg("%v < 0", p))
		return
	}

	// For Quantize, p needs to be the opposite. -3 is to round to three
	// places after the decimal point
	r := zc.Decimal.New()
	r.SetFloat64(real(x))
	_, err := d.Quantize(r, r, -p)
	if err != nil {
		c.Raise(err)
		return
	}

	i := zc.Decimal.New()
	i.SetFloat64(imag(x))
	_, err = d.Quantize(i, i, -p)
	if err != nil {
		c.Raise(err)
		return
	}

	posSign := ""
	if i.Cmp(&zero) > 0 {
		posSign = "+"
	}
	z := fmt.Sprintf("%v%v%vi", r, posSign, i)
	zc.String.Push(c, z)
}

func RoundDecimal(c zc.Calc) {
	d := vars.ForDec(c).Math
	p := zc.Int32.Pop(c)
	u := c.Unit()
	x := zc.Decimal.Pop(c)

	// For Quantize this needs to be the opposite. -3 is to round to three
	// places after the decimal point
	p = -p

	_, err := d.Quantize(x, x, p)
	if err != nil {
		c.Raise(err)
		return
	}
	zc.Decimal.Push(c, x)
	c.SetUnit(u)
}

func SciDecimal(c zc.Calc) {
	unit := c.Unit()
	x := zc.Decimal.Pop(c)
	str := zc.FormatExponent(x.Text('e'))
	zc.String.Push(c, str)
	c.SetUnit(unit)
}

func SciBigFloat(c zc.Calc) {
	unit := c.Unit()
	x := zc.BigFloat.Pop(c)
	str := zc.FormatExponent(x.Text('e', -1))
	zc.String.Push(c, str)
	c.SetUnit(unit)
}
