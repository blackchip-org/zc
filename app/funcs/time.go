package funcs

import (
	"runtime"
	"time"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/vars"
)

func AddDuration(c zc.Calc) {
	y := zc.Duration.Pop(c)
	x := zc.Duration.Pop(c)
	zc.Duration.Push(c, x+y)
}

func AddDurationDateTime(c zc.Calc) {
	y := zc.DateTime.Pop(c)
	x := zc.Duration.Pop(c)
	zc.DateTime.Push(c, y.Add(x))
}

func AddDateTimeDuration(c zc.Calc) {
	y := zc.Duration.Pop(c)
	x := zc.DateTime.Pop(c)
	zc.DateTime.Push(c, x.Add(y))
}

func Date(c zc.Calc) {
	x := zc.DateTime.Pop(c)
	zc.Date.Push(c, x)
}

func DateIs(c zc.Calc) {
	x := zc.String.Pop(c)
	_, ok, _ := zc.Date.Parse(c, x)
	zc.Bool.Push(c, ok)
}

func DateTime(c zc.Calc) {
	x := zc.DateTime.Pop(c)
	zc.DateTime.Push(c, x)
}

func DateTimeIs(c zc.Calc) {
	x := zc.String.Pop(c)
	_, ok, _ := zc.DateTime.Parse(c, x)
	zc.Bool.Push(c, ok)
}

func DayYear(c zc.Calc) {
	x := zc.DateTime.Pop(c)
	zc.Int.Push(c, x.YearDay())
}

func Hours(c zc.Calc) {
	x := zc.Duration.Pop(c)
	zc.Float64.Push(c, x.Hours())
	c.SetUnit("h")
}

func LocalZone(c zc.Calc) {
	v := vars.ForTime(c)
	zc.String.Push(c, v.ZoneName)
	c.SetLabel("time zone")
}

func LocalZoneSet(c zc.Calc) {
	v := vars.ForTime(c)
	zone := zc.String.Pop(c)

	var loc *time.Location
	var err error
	offset, ok := v.Locale.Offsets[v.Locale.Key(zone)]
	if ok {
		zone = v.Locale.DisplayNames[v.Locale.Key(zone)]
		loc = time.FixedZone(zone, offset)
	} else {
		loc, err = time.LoadLocation(zone)
		if err != nil {
			c.Raise(zc.ErrInvalidArg("unknown time zone"))
			return
		}
	}
	v.Zone = loc
	v.ZoneName = zone
	c.Notify("local time zone is now %v", zc.Quote(v.ZoneName))
}

func MinutesTime(c zc.Calc) {
	x := zc.Duration.Pop(c)
	zc.Float64.Push(c, x.Minutes())
	c.SetUnit("m")
}

func Now(c zc.Calc) {
	v := vars.ForTime(c)
	zc.DateTime.Push(c, v.Now())
}

func NowSet(c zc.Calc) {
	v := vars.ForTime(c)
	x := zc.DateTime.Pop(c)
	v.Now = func() time.Time { return x.In(v.Zone) }
	c.Notify("now set to %v", zc.Quote(zc.DateTime.Format(c, x)))
}

func NowReset(c zc.Calc) {
	v := vars.ForTime(c)
	v.Now = func() time.Time { return time.Now() }
	c.Notify("now reset")
}

func SecondsTime(c zc.Calc) {
	x := zc.Duration.Pop(c)
	zc.Float64.Push(c, x.Seconds())
	c.SetUnit("s")
}

func SubDuration(c zc.Calc) {
	y := zc.Duration.Pop(c)
	x := zc.Duration.Pop(c)
	zc.Duration.Push(c, x-y)
}

func SubDateTime(c zc.Calc) {
	y := zc.DateTime.Pop(c)
	x := zc.DateTime.Pop(c)
	zc.Duration.Push(c, x.Sub(y))
}

func SubDateTimeDuration(c zc.Calc) {
	y := zc.Duration.Pop(c)
	x := zc.DateTime.Pop(c)
	zc.DateTime.Push(c, x.Add(-y))
}

func Time(c zc.Calc) {
	x := zc.DateTime.Pop(c)
	zc.Time.Push(c, x)
}

func TimeIs(c zc.Calc) {
	x := zc.String.Pop(c)
	_, ok, _ := zc.Time.Parse(c, x)
	zc.Bool.Push(c, ok)
}

func TimeZone(c zc.Calc) {
	if runtime.GOARCH == "wasm" {
		c.Raise(zc.ErrFeatureNotSupported("tz"))
		return
	}

	v := vars.ForTime(c)
	zone := zc.String.Pop(c)
	dt := zc.DateTime.Pop(c)

	var loc *time.Location
	var err error
	offset, ok := v.Locale.Offsets[v.Locale.Key(zone)]
	if ok {
		loc = time.FixedZone(zone, offset)
	} else {
		loc, err = time.LoadLocation(zone)
		if err != nil {
			c.Raise(zc.ErrInvalidArg("unknown time zone: '%v'", zone))
			return
		}
	}
	zc.DateTime.Push(c, dt.In(loc))
}
