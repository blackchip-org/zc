package funcs

import (
	"fmt"
	"strconv"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/msg"
)

func UTM(c zc.Calc) {
	a := zc.String.Pop(c)
	if len(a) == 0 {
		c.Raise(msg.ErrInvalidArg(a))
		return
	}

	rZone, hemi := a[:len(a)-1], a[len(a)-1]

	var base int
	switch hemi {
	case 'n', 'N':
		base = 32600
	case 's', 'S':
		base = 32700
	default:
		c.Raise(msg.ErrInvalidArg(a))
		return
	}

	zone, err := strconv.Atoi(string(rZone))
	if err != nil || zone < 0 || zone > 60 {
		c.Raise(msg.ErrInvalidArg(a))
		return
	}
	r0 := fmt.Sprintf("EPSG:%v", base+zone)
	zc.String.Push(c, r0)
}
