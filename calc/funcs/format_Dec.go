package funcs

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/calc/state"
	"github.com/blackchip-org/zc/v6/calc/types"
	"github.com/cockroachdb/apd/v3"
)

func RoundDecimal(e *zc.OpEnv) {
	s := state.ForDec(e.State)
	x := e.Args[0].(*apd.Decimal)
	p := e.Args[1].(int32)

	// For Quantize this needs to be the opposite. -3 is to round to three
	// places after the decimal point
	p = -p

	_, err := s.Context.Quantize(x, x, p)
	if err != nil {
		e.Err = zc.ErrOp(e, err)
		return
	}
	e.Returns = []any{x}
}

func ScientificNotationDecimal(e *zc.OpEnv) {
	x := e.Args[0].(*apd.Decimal)
	sn := types.Dec.FormatWith('e', x)
	e.Returns = []any{sn}
}
