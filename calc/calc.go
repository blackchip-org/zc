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
		vols.Conf,
		vols.Stack,
	)
	Min = b.Build()

	b.AddVolume(
		vols.BasicInt,
		vols.BasicDec,
		vols.BasicFloat,
		vols.Format,
		vols.FormatDec,
	)
	Standard = b.Build()

	b.AddVolume(
		vols.BasicDecSS,
		vols.Stat,
	)
	All = b.Build()
}

func New() *zc.Calc { return zc.NewCalc(All) }
