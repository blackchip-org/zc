package calc

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/vols"
)

var (
	Min      *zc.Catalog
	Standard *zc.Catalog
	All      *zc.Catalog
)

func init() {
	b := zc.NewCatalogBuilder()
	b.AddVolume(
		vols.Basic,
		vols.Zcalc,
	)
	Min = b.Build()

	b.AddVolume(
		vols.BasicInt,
		vols.BasicDec,
		vols.BasicFloat,
	)
	Standard = b.Build()
	All = b.Build()
}

func New() *zc.Calc { return zc.NewCalc(All) }
