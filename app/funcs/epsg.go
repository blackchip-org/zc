package funcs

import (
	"fmt"
	"strconv"

	"github.com/blackchip-org/zc/v6"
)

func UTM(c zc.Calc) {
	p0 := []rune(zc.String.Pop(c))
	if len(p0) == 0 {
		c.Raise(zc.ErrInvalidArg("hemisphere"))
		return
	}

	rZone, hemi := p0[:len(p0)-1], p0[len(p0)-1]

	var base int
	switch hemi {
	case 'n', 'N':
		base = 32600
	case 's', 'S':
		base = 32700
	default:
		c.Raise(zc.ErrInvalidArg("hemisphere"))
		return
	}

	zone, err := strconv.Atoi(string(rZone))
	if err != nil || zone < 0 || zone > 60 {
		c.Raise(zc.ErrInvalidArg("zone"))
		return
	}
	r0 := fmt.Sprintf("EPSG:%v", base+zone)
	zc.String.Push(c, r0)
}
