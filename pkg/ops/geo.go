package ops

import (
	"math"

	"github.com/blackchip-org/zc/v5/pkg/ext"
	"github.com/blackchip-org/zc/v5/pkg/zc"
)

const EarthRadius = 6371000

/*
oper	earth-equatorial-radius
func	- -- Int
macro	6378137 'meters' anno
title	Equatorial radius of the Earth

desc
The equatorial radius of the Earth in meters.

Source:
https://nssdc.gsfc.nasa.gov/planetary/factsheet/earthfact.html
end

example
earth-equatorial-radius -- 6378137 # meters
end
*/

/*
oper	earth-polar-radius
func	- -- Int
macro	6356752 'meters' anno
title	Polar radius of the Earth

desc
The polar radius,of the Earth in meters.

Source:
https://nssdc.gsfc.nasa.gov/planetary/factsheet/earthfact.html
end

example
earth-equatorial-radius -- 6378137 # meters
end
*/

/*
oper	earth-radius
func	EarthRadiusFn -- Int
title	Average radius of the Earth

desc
The globally-average value of the Earth in meters.

Source:
https://nssdc.gsfc.nasa.gov/planetary/factsheet/earthfact.html
end

example
earth-radius -- 6371000 # meters
end
*/
func EarthRadiusFn(c zc.Calc) {
	zc.PushInt(c, EarthRadius)
	zc.Annotate(c, "meters")
}

/*
oper	haversine
func	Haversine lat0:DMS lon0:DMS lat1:DMS lon1:DMS -- Float
title	Great circle distance between two points

desc
Calculates the great circle distance between (*lat0*, *lon0*) and
(*lat1*, *lon1*) using the haversine formula.

Source:
https://community.esri.com/t5/coordinate-reference-systems-blog/distance-on-a-sphere-the-haversine-formula/ba-p/902128
end

example
51.510357 -0.116773 -- 51.510357 | -0.116773
38.889931 -77.009003 -- 51.510357 | -0.116773 | 38.889931 | -77.009003
haversine dec 3 round -- 5897658.289
end
*/
func Haversine(c zc.Calc) {
	lon2 := zc.PopDMS(c)
	lat2 := zc.PopDMS(c)
	lon1 := zc.PopDMS(c)
	lat1 := zc.PopDMS(c)

	phi1 := lat1.Radians()
	phi2 := lat2.Radians()

	deltaPhi := lat2.Sub(lat1).Radians()
	deltaLambda := lon2.Sub(lon1).Radians()

	a := math.Pow(math.Sin(deltaPhi/2), 2) + math.Cos(phi1)*math.Cos(phi2)*math.Pow(math.Sin(deltaLambda/2.0), 2)
	c0 := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	r0 := EarthRadius * c0

	zc.PushFloat(c, r0)
	zc.Annotate(c, "meters")
}

/*
oper	proj
func	Proj p0:Float p1:Float s:Str t:Str -- Float Float
title	Transform coordinate

desc
Transform coordinate (*p0*, *p1*) in coordinate system *s* to a coordinate
in system *t*. The order of the coordinates is defined by the coordinate system
and it may be (lat, lon) or (x, y).
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
