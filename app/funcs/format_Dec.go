package funcs

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/state"
)

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
