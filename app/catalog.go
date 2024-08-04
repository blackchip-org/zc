package app

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/vols"
)

var mainCatalog *zc.Catalog

func init() {
	mainCatalog = zc.NewCatalog()
	mainCatalog.AddVol(
		vols.Anno,
		vols.Basic,
		vols.Conf,
		vols.Format,
		vols.Stack,
		vols.Stat,
	)
}
