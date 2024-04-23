package calc

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/volumes/intb"
	"github.com/blackchip-org/zc/v6/volumes/ints"
	"github.com/blackchip-org/zc/v6/volumes/text"
	"github.com/blackchip-org/zc/v6/volumes/zcalc"
)

var (
	Min   *zc.Catalog
	Basic *zc.Catalog
	All   *zc.Catalog
)

func init() {
	b := zc.NewCatalogBuilder()
	b.AddVolume(
		ints.Volume,
		text.Volume,
		zcalc.Volume,
	)
	Min = b.Build()

	b.AddVolume(
		intb.Volume,
	)
	Basic = b.Build()

	All = b.Build()
}

func New() *zc.Calc { return zc.NewCalc(All) }
