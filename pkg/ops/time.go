package ops

import (
	"time"

	"github.com/blackchip-org/zc/pkg/ptime/locale"
	"github.com/blackchip-org/zc/pkg/zc"
)

type timeState struct {
	locale   *locale.Locale
	zone     *time.Location
	zoneName string
}

func getTimeState(c zc.Calc) *timeState {
	s, ok := c.State("time")
	if !ok {
		loc := time.Now().Location()
		tz, _ := time.Now().Zone()

		ts := &timeState{
			locale:   locale.EnUS,
			zone:     loc,
			zoneName: tz,
		}
		zc.Now = func() time.Time { return time.Now().In(ts.zone) }
		s = ts
		c.NewState("time", s)
	}
	return s.(*timeState)
}

func AddDurationDateTime(c zc.Calc) {
	a1 := zc.PopDateTime(c)
	a0 := zc.PopDuration(c)
	r0 := a1.Add(a0)
	zc.PushDateTime(c, r0)
}

func AddDateTimeDuration(c zc.Calc) {
	a1 := zc.PopDuration(c)
	a0 := zc.PopDateTime(c)
	r0 := a0.Add(a1)
	zc.PushDateTime(c, r0)
}

func Date(c zc.Calc) {
	a0 := zc.PopDateTime(c)
	zc.PushDate(c, a0)
}

func DateTime(c zc.Calc) {
	a0 := zc.PopDateTime(c)
	zc.PushDateTime(c, a0)
}

func DayYear(c zc.Calc) {
	a0 := zc.PopDateTime(c)
	r0 := a0.YearDay()
	zc.PushInt(c, r0)
}

func Hours(c zc.Calc) {
	a0 := zc.PopDuration(c)
	r0 := a0.Hours()
	zc.PushFloat(c, r0)
}

func LocalZone(c zc.Calc) {
	s := getTimeState(c)
	zone := zc.PopString(c)

	var loc *time.Location
	var err error
	offset, ok := s.locale.Offsets[s.locale.Key(zone)]
	if ok {
		zone = s.locale.DisplayNames[s.locale.Key(zone)]
		loc = time.FixedZone(zone, offset)
	} else {
		loc, err = time.LoadLocation(zone)
		if err != nil {
			zc.ErrInvalidArgument(c, c.Op(), zone)
			return
		}
	}
	s.zone = loc
	s.zoneName = zone
	c.SetInfo("local time zone is now %v", zc.Quote(s.zoneName))
}

func LocalZoneGet(c zc.Calc) {
	s := getTimeState(c)
	zc.PushString(c, s.zoneName)
}

func Minutes(c zc.Calc) {
	a0 := zc.PopDuration(c)
	r0 := a0.Minutes()
	zc.PushFloat(c, r0)
}

func Now(c zc.Calc) {
	r0 := zc.Now()
	zc.PushDateTime(c, r0)
}

func NowSet(c zc.Calc) {
	s := getTimeState(c)
	a0 := zc.PopDateTime(c)
	zc.Now = func() time.Time { return a0.In(s.zone) }
	c.SetInfo("now set to %v", zc.Quote(zc.DateTime.Format(a0)))
}

func NowRestore(c zc.Calc) {
	s := getTimeState(c)
	zc.Now = func() time.Time { return time.Now().In(s.zone) }
	c.SetInfo("now restored")
}

func Seconds(c zc.Calc) {
	a0 := zc.PopDuration(c)
	r0 := a0.Seconds()
	zc.PushFloat(c, r0)
}

func SubDateTime(c zc.Calc) {
	a1 := zc.PopDateTime(c)
	a0 := zc.PopDateTime(c)
	r0 := a0.Sub(a1)
	zc.PushDuration(c, r0)
}

func Time(c zc.Calc) {
	a0 := zc.PopDateTime(c)
	zc.PushTime(c, a0)
}

func TimeZone(c zc.Calc) {
	s := getTimeState(c)
	zone := zc.PopString(c)
	dt := zc.PopDateTime(c)

	var loc *time.Location
	var err error
	offset, ok := s.locale.Offsets[s.locale.Key(zone)]
	if ok {
		loc = time.FixedZone(zone, offset)
	} else {
		loc, err = time.LoadLocation(zone)
		if err != nil {
			zc.ErrInvalidArgument(c, c.Op(), zone)
			return
		}
	}

	r0 := dt.In(loc)
	zc.PushDateTime(c, r0)
}
