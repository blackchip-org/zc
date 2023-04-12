package stdlib

import (
	"strconv"

	"github.com/blackchip-org/zc/types"
)

type slComplex struct {
	val complex128
}

func unwrapC(c types.Complex) complex128 {
	return c.(slComplex).val
}

func wrapC(c complex128) slComplex {
	return slComplex{val: c}
}

func (c slComplex) Add(c2 types.Complex) types.Complex {
	return wrapC(c.val + unwrapC(c2))
}

func (c slComplex) Type() types.Type {
	return types.ComplexType{}
}

func (c slComplex) Value() types.Value {
	return c
}

func (c slComplex) Native() any {
	return c
}

func (c slComplex) String() string {
	s := strconv.FormatComplex(c.val, 'g', 16, 128)
	// For some reason, the complex number is surrounded by parens.
	// Remove them.
	return s[1 : len(s)-1]
}

func ParseComplex(s string) (types.Complex, error) {
	c, err := strconv.ParseComplex(s, 128)
	if err != nil {
		return nil, types.ErrParse
	}
	return wrapC(c), nil
}

func NewComplex(r float64, i float64) types.Complex {
	return wrapC(complex(r, i))
}

func UseComplex() {
	types.UseComplex(types.ComplexImpl{
		Parse: ParseComplex,
		New:   NewComplex,
	})
}
