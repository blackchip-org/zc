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

/*
oper	add
func	AddDuration p0:Duration p1:Duration -- Duration
func	AddDurationDateTime p0:Duration p1:DateTime -- Duration
func	AddDateTimeDuration p0:DateTime p1:Duration -- Duration
alias	a
alias	+
title	Time or duration addition

desc
Adds a duration and time or adds two durations.
end

example
c 3:30pm -- 3:30pm
2h add -- Mon Jan 2 2006 5:30:00pm -0700 MST
c 2h30m -- 2h30m
45m add -- 3h 15m
end
*/
func AddDuration(c zc.Calc) {
	a1 := zc.PopDuration(c)
	a0 := zc.PopDuration(c)
	r0 := a1 + a0
	zc.PushDuration(c, r0)
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

/*
oper	date
func	Date p0:DateTime -- Date
title	Formats to a common date layout

desc
Formats date/time *p0* to a common date layout. Time information, if any, is
removed.
end

example
'2006-01-02T15:04:05 UTC -- 2006-01-02T15:04:05 UTC
date -- Mon Jan 2 2006
end
*/
func Date(c zc.Calc) {
	a0 := zc.PopDateTime(c)
	zc.PushDate(c, a0)
}

/*
oper	datetime
func	DateTime p0:DateTime -- DateTime
alias	dt
title	Formats to a common date/time layout

desc
Formats a date/time to a common layout.
end

example
'2006-01-02T15:04:05 UTC -- 2006-01-02T15:04:05 UTC
datetime -- Mon Jan 2 2006 3:04:05pm UTC
end
*/
func DateTime(c zc.Calc) {
	a0 := zc.PopDateTime(c)
	zc.PushDateTime(c, a0)
}

/*
oper	day-year
func	DayYear p0:DateTime -- Int
alias	doy
title 	Day of year

desc
Day of year for a given date *p0*.
end

example
2006-03-15 -- 2006-03-15
day-year -- 74
end
*/
func DayYear(c zc.Calc) {
	a0 := zc.PopDateTime(c)
	r0 := a0.YearDay()
	zc.PushInt(c, r0)
}

/*
oper	hours
func	Hours p0:Duration -- Float
title	Hours in duration

desc
Converts the duration *p0* into hours.
end

example
10h20m30s -- 10h20m30s
hours 2 round -- 10.34
end
*/
func Hours(c zc.Calc) {
	a0 := zc.PopDuration(c)
	r0 := a0.Hours()
	zc.PushFloat(c, r0)
}

/*
oper	is-datetime
func	IsDateTime p0:Str -- Bool
alias	is-dt
title 	Checks value can be parsed as a `DateTime`

desc
Returns `true` if the value *p0* can be parsed as a `DateTime“.
end

example
c [2 May 2023] is-datetime -- true
c [2 Nay 2023] is-datetime -- false
end
*/
func IsDateTime(c zc.Calc) {
	p0 := zc.PopString(c)
	r0 := zc.DateTime.Is(p0)
	zc.PushBool(c, r0)
}

/*
oper	local-zone
func	LocalZone p0:Str --
title	Sets the local time zone

desc
Sets the local time zone to *p0*.
end

exammple
now time -- 3:04:05pm -0700 MST
c [est] local-zone -- *local time zone is now 'EST'*
now time -- 5:04:05pm -0500 EST
c [Asia/Jakarta] local-zone -- *local time zone is now 'Asia/Jakarta'*
now time -- 5:04:05am +0700
end
*/
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
			zc.ErrInvalidArgs(c, "unknown time zone")
			return
		}
	}
	s.zone = loc
	s.zoneName = zone
	c.SetInfo("local time zone is now %v", zc.Quote(s.zoneName))
}

/*
oper	local-zone=
func	LocalZoneGet -- Str
title	Gets the local time zone

desc
Gets the local time zone.
end

example
local-zone= -- MST
end
*/
func LocalZoneGet(c zc.Calc) {
	s := getTimeState(c)
	zc.PushString(c, s.zoneName)
}

/*
oper	minutes
func	Minutes p0:Duration -- Float
title	Minutes in duration

desc
Converts the duration *p0* into minutes.
end

example
10h20m30s -- 10h20m30s
minutes 2 round -- 620.5
end
*/
func Minutes(c zc.Calc) {
	a0 := zc.PopDuration(c)
	r0 := a0.Minutes()
	zc.PushFloat(c, r0)
}

/*
oper	now
func	Now -- DateTime
title	Current date and time

desc
The current date and time. If `now-set` has been called, that date and
time will be returned instead.
end

example
now -- Mon Jan 2 2006 3:04:05pm -0700 MST
end
*/
func Now(c zc.Calc) {
	r0 := zc.Now()
	zc.PushDateTime(c, r0)
}

/*
oper	now-set
func	NowSet p0:DateTime --
title 	Override now value

desc
Override the value returned by `now`. Useful for to mock current time while
testing.
end

example
'Nov 5 1955 01:22 -- Nov 5 1955 01:22
now-set -- *now set to 'Sat Nov 5 1955 1:22:00am -0700 MST'*
now -- Sat Nov 5 1955 1:22:00am -0700 MST
end
*/
func NowSet(c zc.Calc) {
	s := getTimeState(c)
	a0 := zc.PopDateTime(c)
	zc.Now = func() time.Time { return a0.In(s.zone) }
	c.SetInfo("now set to %v", zc.Quote(zc.DateTime.Format(a0)))
}

/*
oper	now-restore
func	NowRestore --
title	Cancel now override

desc
Cancel override of the value returned by now.
end
*/
func NowRestore(c zc.Calc) {
	s := getTimeState(c)
	zc.Now = func() time.Time { return time.Now().In(s.zone) }
	c.SetInfo("now restored")
}

/*
oper	seconds
func	Seconds p0:Duration -- Float
title	Seconds in duration

desc
Converts the duration *p0* into seconds.
end

example
10h20m30s -- 10h20m30s
seconds -- 37230
end
*/
func Seconds(c zc.Calc) {
	a0 := zc.PopDuration(c)
	r0 := a0.Seconds()
	zc.PushFloat(c, r0)
}

/*
oper	sub
func	SubDuration p0:Duration p1:Duration -- Duration
func	SubDateTimeDuration p0:DateTime p1:Duration -- Duration
func 	SubDateTime p0:DateTime p1:DateTime -- Duration
alias	s
alias	-
title	Time or duration subtraction
desc
Subtracts a duration from a time or subtracts two durations.
end

example
c 3:30pm 2h sub -- Mon Jan 2 2006 1:30:00pm -0700 MST
c 2h30m 45m sub -- 1h 45m
c 3:30pm 1:30pm sub -- 2h
end
*/
func SubDuration(c zc.Calc) {
	a1 := zc.PopDuration(c)
	a0 := zc.PopDuration(c)
	r0 := a0 - a1
	zc.PushDuration(c, r0)
}

func SubDateTime(c zc.Calc) {
	a1 := zc.PopDateTime(c)
	a0 := zc.PopDateTime(c)
	r0 := a0.Sub(a1)
	zc.PushDuration(c, r0)
}

func SubDateTimeDuration(c zc.Calc) {
	a1 := zc.PopDuration(c)
	a0 := zc.PopDateTime(c)
	r0 := a0.Add(-a1)
	zc.PushDateTime(c, r0)
}

/*
oper	time
func	Time p0:DateTime -- Time
title	Formats to a common time layout

desc
Formats a date/time with the common time layout. Date information, if any,
is discarded.
end

example
'2006-01-02T15:04:05 UTC -- 2006-01-02T15:04:05 UTC
time -- 3:04:05pm UTC
end
*/
func Time(c zc.Calc) {
	a0 := zc.PopDateTime(c)
	zc.PushTime(c, a0)
}

/*
oper	timezone
func	TimeZone p0:DateTime zone:Str -- DateTime
alias	tz
title	Convert time to another time zone

desc
Converts time *p0* to a given time *zone*.
end

example
now -- Mon Jan 2 2006 3:04:05pm -0700 MST
[PST] tz -- Mon Jan 2 2006 2:04:05pm -0800 PST
[Asia/Jakarta] tz -- Tue Jan 3 2006 5:04:05am +0700 WIB
end
*/
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
			zc.ErrInvalidArgs(c, "unknown time zone")
			return
		}
	}

	r0 := dt.In(loc)
	zc.PushDateTime(c, r0)
}
