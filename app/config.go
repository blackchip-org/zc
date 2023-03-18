package app

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/zlib"
)

func DefaultConfig() zc.Config {
	return zc.Config{
		ModuleDefs:   zlib.All,
		Preload:      zlib.Preload,
		PreludeCLI:   zlib.PreludeUser,
		PreludeDev:   zlib.PreludeDev,
		RoundingMode: zc.RoundingModeHalfUp,
		IntLayout:    ",000",
		Point:        '.',
	}
}

func NewDefaultCalc() *zc.Calc {
	calc, err := zc.NewCalc(DefaultConfig())
	if err != nil {
		panic(err)
	}
	return calc
}
