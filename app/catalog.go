package app

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/vols"
)

var mainCatalog *zc.Catalog

func init() {
	mainCatalog = zc.NewCatalog()
	mainCatalog.AddVol(
		vols.Angle,
		vols.Anno,
		vols.Basic,
		vols.BasicFloat,
		vols.Bool,
		vols.Cmp,
		vols.Color,
		vols.Conf,
		vols.Format,
		vols.Hof,
		vols.Len,
		vols.Mass,
		vols.Prog,
		vols.Rand,
		vols.Sci,
		vols.Seq,
		vols.Stack,
		vols.Stat,
		vols.Temp,
		vols.Text,
	)
}
