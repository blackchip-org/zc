package funcs

import (
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

func Time(c zc.Calc) {
	x := zc.DateTime.Pop(c)
	zc.Time.Push(c, x)
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
