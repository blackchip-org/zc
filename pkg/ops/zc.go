package ops

import (
	"fmt"

	"github.com/blackchip-org/zc/pkg/zc"
)

func Version(c zc.Calc) {
	r0 := fmt.Sprintf("%v: %v (%v)", zc.ProgName, zc.Version, zc.BuildDate)
	zc.PushString(c, r0)
}
