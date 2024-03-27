package ops

import (
	"github.com/blackchip-org/zc/v6/pkg/zc"
)

/*
oper	version
func	Version --
title	Version number

desc
Version number of the calculator.
end
*/
func Version(c zc.Calc) {
	c.SetInfo("%v %v (%v)", zc.ProgName, zc.Version, zc.BuildDate)
}
