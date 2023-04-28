package ops

import (
	"fmt"
	"strconv"

	"github.com/blackchip-org/zc/pkg/zc"
)

//tab epsg.web-mercator	-- 'EPSG:3857' -- Web Mercator, EPSG:3857
//tab epsg.wgs-84		-- 'EPSG:4326' -- World Geodetic System of 1984, EPSG:4326

/*
oper	epsg.utm
func	UTM p0:Str -- Str
title	Universal Transverse Mercator (WGS-84)

desc
The EPGS code for the given UTM zone *p0*. The zone should be a number
between 1 and 60 inclusive and is followed by a hemisphere designator of
'n' or 's'.
end

example
17n epsg.utm -- EPSG:32617
end
*/
func UTM(c zc.Calc) {
	p0 := []rune(zc.PopString(c))
	rZone, hemi := p0[:len(p0)-1], p0[len(p0)-1]

	var base int
	switch hemi {
	case 'n', 'N':
		base = 32600
	case 's', 'S':
		base = 32700
	default:
		zc.ErrInvalidArgs(c)
		return
	}

	zone, err := strconv.Atoi(string(rZone))
	if err != nil || zone < 0 || zone > 60 {
		zc.ErrInvalidArgs(c)
		return
	}
	r0 := fmt.Sprintf("EPSG:%v", base+zone)
	zc.PushString(c, r0)
}
