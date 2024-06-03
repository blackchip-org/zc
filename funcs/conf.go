package funcs

import (
	"fmt"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/state"
)

func FloatPrecSet(e *zc.OpEnv) {
	s := state.ForConf(e.State)
	p := e.Args[0].(uint)
	s.FloatPrec = p
	e.Info = fmt.Sprintf("precision set to %v", p)
}

func FloatPrecGet(e *zc.OpEnv) {
	s := state.ForConf(e.State)
	e.Returns = []any{s.FloatPrec}
	e.Annos = []string{"precision"}
}
