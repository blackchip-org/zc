package calc

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/vols"
)

var (
	Min   *zc.Catalog
	Basic *zc.Catalog
	All   *zc.Catalog
)

func init() {
	b := zc.NewCatalogBuilder()
	b.AddVolume(
		vols.BasicInt,
	)
	Min = b.Build()

	Basic = b.Build()
	All = b.Build()
}

func New() *zc.Calc { return zc.NewCalc(All) }
