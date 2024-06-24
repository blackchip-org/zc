package funcs

import (
	"fmt"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/state"
	"github.com/cockroachdb/apd/v3"
)

func RoundComplex128(e *zc.OpEnv) {
	var zero apd.Decimal
	s := state.ForDec(e.State)
	p := zc.Int32.Pop(e)
	x := zc.Complex128.Pop(e)

	if p < 0 {
		e.Err = zc.ErrInvalidArg(e, "%v < 0", p)
		return
	}

	// For Quantize, p needs to be the opposite. -3 is to round to three
	// places after the decimal point
	r := zc.Decimal.New()
	r.SetFloat64(real(x))
	_, err := s.Context.Quantize(r, r, -p)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}

	i := zc.Decimal.New()
	i.SetFloat64(imag(x))
	_, err = s.Context.Quantize(i, i, -p)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}

	posSign := ""
	if i.Cmp(&zero) > 0 {
		posSign = "+"
	}
	z := fmt.Sprintf("%v%v%vi", r, posSign, i)
	zc.String.Push(e, z)
}

func RoundDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	p := zc.Int32.Pop(e)
	x := zc.Decimal.Pop(e)

	// For Quantize this needs to be the opposite. -3 is to round to three
	// places after the decimal point
	p = -p

	_, err := s.Context.Quantize(x, x, p)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	zc.Decimal.Push(e, x)
}

func ScientificNotationDecimal(e *zc.OpEnv) {
	x := zc.Decimal.Pop(e)
	sn := zc.FormatExponent(x.Text('e'))
	zc.String.Push(e, sn)
	zc.Decimal.Recycle(x)
}
