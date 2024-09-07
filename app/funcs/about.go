package funcs

import (
	"fmt"

	"github.com/blackchip-org/zc/v6"
)

func Version(c zc.Calc) {
	v := fmt.Sprintf("%v %v (%v) git%v", zc.ProgName, zc.Version, zc.BuildDate, zc.Commit)
	c.Notify(v)
}
