package calc

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/kinds"
)

func NewStandard() *zc.Calc {
	cat := zc.NewCatalog()
	cat.AddKind(kinds.Int, kinds.IntArch)
	cat.AddKind(kinds.Text)
	return zc.NewCalc(cat)
}
