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
		vols.Bool,
		vols.Cmp,
		vols.Color,
		vols.Complex,
		vols.Const,
		vols.Crypto,
		vols.Dec,
		vols.Emoji,
		vols.Entity,
		vols.Epsg,
		vols.Float,
		vols.Format,
		vols.Geo,
		vols.Hof,
		vols.Iec,
		vols.Int,
		vols.Len,
		vols.Mass,
		vols.Prog,
		vols.Rand,
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
