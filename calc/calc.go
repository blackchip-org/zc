package calc

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/volumes/basic"
	"github.com/blackchip-org/zc/v6/volumes/boolean"
	"github.com/blackchip-org/zc/v6/volumes/floatb"
	"github.com/blackchip-org/zc/v6/volumes/floats"
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
		basic.Volume,
		boolean.Volume,
		floats.Volume,
		ints.Volume,
		text.Volume,
		zcalc.Volume,
	)
	Min = b.Build()

	b.AddVolume(
		intb.Volume,
		floatb.Volume,
	)
	Basic = b.Build()

	All = b.Build()
}

func New() *zc.Calc { return zc.NewCalc(All) }
