package ops

import (
	"github.com/blackchip-org/zc/pkg/zc"
)

func Version(c zc.Calc) {
	c.SetInfo("%v %v (%v)", zc.ProgName, zc.Version, zc.BuildDate)
}
