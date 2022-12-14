package app

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/zlib"
)

func DefaultConfig() zc.Config {
	return zc.Config{
		ModuleDefs:  zlib.All,
		PreludeCLI:  zlib.PreludeUser,
		PreludeDev:  zlib.PreludeDev,
		ValueConfig: zc.DefaultValueConfig(),
	}
}

func NewDefaultCalc() *zc.Calc {
	calc, err := zc.NewCalcWithConfig(DefaultConfig())
	if err != nil {
		panic(err)
	}
	return calc
}
