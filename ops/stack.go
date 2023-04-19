package ops

import "github.com/blackchip-org/zc"

func Clear(c zc.Calc) {
	c.SetStack([]string{})
}
