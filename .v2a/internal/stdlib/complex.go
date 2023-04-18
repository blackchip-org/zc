package stdlib

import (
	"errors"
	"strconv"

	"github.com/blackchip-org/zc/number"
)

type slComplex struct {
	val complex128
}

func unwrapC(c number.Complex) complex128 {
	return c.(slComplex).val
}

func wrapC(c complex128) slComplex {
	return slComplex{val: c}
}

func (c slComplex) Add(c2 number.Complex) number.Complex {
	return wrapC(c.val + unwrapC(c2))
}

func (c slComplex) String() string {
	s := strconv.FormatComplex(c.val, 'g', 16, 128)
	// For some reason, the complex number is surrounded by parens.
	// Remove them.
	return s[1 : len(s)-1]
}

func ParseComplex(s string) (number.Complex, error) {
	c, err := strconv.ParseComplex(s, 128)
	if err != nil {
		return nil, errors.New("unable to parse")
	}
	return wrapC(c), nil
}

func NewComplex(r float64, i float64) number.Complex {
	return wrapC(complex(r, i))
}

func UseComplex() {
	number.UseComplex(number.ComplexImpl{
		Parse: ParseComplex,
		New:   NewComplex,
	})
}
