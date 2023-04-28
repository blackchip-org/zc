package ops

import (
	"github.com/blackchip-org/zc/pkg/ext"
	"github.com/blackchip-org/zc/pkg/zc"
)

/*
oper	proj
func	Proj p0:Float p1:Float s:Str t:Str -- Float Float
title	Transform coordinate

desc
Transform coordinate (*p0*, *p01) in coordinate system *s* to a coordinate
in coordinate system *t*. The order of the coordinates is defined by the
coordinate system and it may be (lat, lon) or (x, y).
end

example
39.203611 -76.856944 -- 39.203611 | -76.856944
epsg.wgs-84 18n epsg.utm --  39.203611 | -76.856944 | EPSG:4326 | EPSG:32618
proj -- 339660.12559342897 | 4.341014551927999e06
end
*/
func Proj(c zc.Calc) {
	tCRS := zc.PopString(c)
	sCRS := zc.PopString(c)
	p1 := zc.PopFloat(c)
	p0 := zc.PopFloat(c)

	r0, r1, err := ext.ProjTransform(p0, p1, sCRS, tCRS)
	if err != nil {
		c.SetError(err)
		return
	}
	zc.PushFloat(c, r0)
	zc.PushFloat(c, r1)
}
