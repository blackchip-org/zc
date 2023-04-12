package app

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/zlib"
)

func DefaultConfig() *zc.Config {
	return &zc.Config{
		ModuleDefs:   zlib.All,
		Preload:      zlib.Preload,
		PreludeCLI:   zlib.PreludeUser,
		PreludeDev:   zlib.PreludeDev,
		RoundingMode: zc.RoundingModeHalfUp,
	}
}

func NewDefaultCalc() *zc.CalcImpl {
	calc, err := zc.NewCalc(DefaultConfig())
	if err != nil {
		panic(err)
	}
	return calc
}
