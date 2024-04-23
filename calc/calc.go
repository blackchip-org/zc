package calc

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/volumes/intb"
	"github.com/blackchip-org/zc/v6/volumes/ints"
	"github.com/blackchip-org/zc/v6/volumes/text"
	"github.com/blackchip-org/zc/v6/volumes/zcalc"
)

func NewStandard() *zc.Calc {
	cat := zc.NewCatalog()
	cat.AddVolume(
		intb.Volume,
		ints.Volume,
		text.Volume,
		zcalc.Volume,
	)
	return zc.NewCalc(cat)
}
