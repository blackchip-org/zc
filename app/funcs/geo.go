package funcs

import (
	"math"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/ext"
)

const EarthRadius = 6371000

func Haversine(c zc.Calc) {
	lon2, err := zc.DMS.Pop(c).Degrees().Float64()
	if err != nil {
		c.Raise(err)
		return
	}

	lat2, err := zc.DMS.Pop(c).Degrees().Float64()
	if err != nil {
		c.Raise(err)
		return
	}

	lon1, err := zc.DMS.Pop(c).Degrees().Float64()
	if err != nil {
		c.Raise(err)
		return
	}

	lat1, err := zc.DMS.Pop(c).Degrees().Float64()
	if err != nil {
		c.Raise(err)
		return
	}

	piOver180 := math.Pi / 180

	phi1 := lat1 * piOver180
	phi2 := lat2 * piOver180

	deltaPhi := (lat2 - lat1) * piOver180
	deltaLambda := (lon2 - lon1) * piOver180

	a := math.Pow(math.Sin(deltaPhi/2), 2) + math.Cos(phi1)*math.Cos(phi2)*math.Pow(math.Sin(deltaLambda/2.0), 2)
	c0 := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	r0 := EarthRadius * c0

	zc.Float64.Push(c, r0)
	c.SetUnit("m")
}

func Proj(c zc.Calc) {
	tcrs := zc.String.Pop(c)
	scrs := zc.String.Pop(c)
	p1 := zc.Float64.Pop(c)
	p0 := zc.Float64.Pop(c)

	r0, r1, err := ext.ProjTransform(p0, p1, scrs, tcrs)
	if err != nil {
		c.Raise(err)
		return
	}
	zc.Float64.Push(c, r0)
	zc.Float64.Push(c, r1)
}
