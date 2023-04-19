package ops

import "github.com/blackchip-org/zc"

const (
	False = "[false]"
	True  = "[true]"
)

func AndBool(c zc.Calc) {
	a1 := zc.PopBool(c)
	a0 := zc.PopBool(c)
	r0 := a0 && a1
	zc.PushBool(c, r0)
}

func NotBool(c zc.Calc) {
	a0 := zc.PopBool(c)
	r0 := !a0
	zc.PushBool(c, r0)
}

func OrBool(c zc.Calc) {
	a1 := zc.PopBool(c)
	a0 := zc.PopBool(c)
	r0 := a0 || a1
	zc.PushBool(c, r0)
}
