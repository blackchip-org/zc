package app

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/modules"
)

func DefaultConfig() zc.Config {
	return zc.Config{
		ModuleDefs: modules.All,
		Prelude:    modules.PreludeCLI,
	}
}

func NewDefaultCalc() (*zc.Calc, error) {
	return zc.NewCalc(DefaultConfig())
}
