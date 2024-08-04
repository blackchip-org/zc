package funcs

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/vars"
)

func RoundingModeSet(c zc.Calc) {
	conf := vars.ForConf(c)
	mode := zc.String.Pop(c)
	if err := conf.SetRoundingMode(mode); err != nil {
		c.Raise(zc.ErrInvalidArg(err.Error()))
		return
	}
	c.Notify("rounding mode set to %v", mode)
}

func RoundingModeGet(c zc.Calc) {
	conf := vars.ForConf(c)
	zc.String.Push(c, conf.GetRoundingMode())
}
