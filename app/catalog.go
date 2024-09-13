package app

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/vols"
)

var mainCatalog *zc.Catalog

func init() {
	mainCatalog = zc.NewCatalog()
	mainCatalog.AddVol(
		vols.About,
		vols.Angle,
		vols.Basic,
		vols.BasicFloat,
		vols.Bool,
		vols.Cmp,
		vols.Color,
		vols.Const,
		vols.Crypto,
		vols.Emoji,
		vols.Entity,
		vols.Epsg,
		vols.Format,
		vols.Geo,
		vols.Hof,
		vols.Iec,
		vols.Int,
		vols.Len,
		vols.Mass,
		vols.Prog,
		vols.Rand,
		vols.Real,
		vols.Seq,
		vols.Si,
		vols.Stack,
		vols.Stat,
		vols.Temp,
		vols.Text,
		vols.Time,
		vols.Trig,
		vols.Tz,
	)
}
