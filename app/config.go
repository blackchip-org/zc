package app

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/zlib"
)

func DefaultConfig() zc.Config {
	return zc.Config{
		ModuleDefs:   zlib.All,
		PreludeCLI:   zlib.PreludeUser,
		PreludeDev:   zlib.PreludeDev,
		Places:       16,
		RoundingMode: zc.RoundingModeHalfUp,
		IntFormat:    ",000",
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
