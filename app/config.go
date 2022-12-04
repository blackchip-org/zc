package app

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/zlib"
)

func DefaultConfig() zc.Config {
	return zc.Config{
		ModuleDefs: zlib.All,
		Prelude:    zlib.PreludeCLI,
	}
}

func NewDefaultCalc() (*zc.Calc, error) {
	return zc.NewCalc(DefaultConfig())
}
