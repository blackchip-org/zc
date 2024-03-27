package zc

import (
	"cmp"
	"time"

	"github.com/blackchip-org/zc/v6/pkg/ptime"
	"github.com/blackchip-org/zc/v6/pkg/ptime/locale"
	"github.com/blackchip-org/zc/v6/pkg/types"
)

const (
	dateLayout     = "[weekday/abbr] [month/abbr] [day] [year]"
	timeLayout     = "[hour/12]:[minute]:[second][period/alt] [offset-zone]"
	dateTimeLayout = dateLayout + " " + timeLayout
)

var pt = ptime.For(locale.EnUS)
var Now = func() time.Time { return time.Now() }

// ---

type DateType struct{}

func (t DateType) String() string { return "Date" }

func (t DateType) Parse(s string) (time.Time, bool) {
	parsed, err := pt.ParseDate(s)
	if err != nil {
		return time.Time{}, false
	}
	tm, err := pt.Time(parsed, Now())
	if err != nil {
		return time.Time{}, false
	}
	return tm, true
}

func (t DateType) MustParse(s string) time.Time {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t DateType) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t DateType) Compare(x1 string, x2 string) (int, bool) {
	t1, ok := t.Parse(x1)
	if !ok {
		return 0, false
	}
	t2, ok := t.Parse(x2)
	if !ok {
		return 0, false
	}
	return t1.Compare(t2), true
}

func (t DateType) Format(tm time.Time) string {
	return pt.Format(dateLayout, tm)
}

func PopDate(c Calc) time.Time     { return Date.MustParse(c.MustPop()) }
func PushDate(c Calc, r time.Time) { c.Push(Date.Format(r)) }

// ---

type DateTimeType struct{}

func (t DateTimeType) String() string { return "DateTime" }

func (t DateTimeType) Parse(str string) (time.Time, bool) {
	parsed, err := pt.Parse(str)
	if err != nil {
		return time.Time{}, false
	}
	tm, err := pt.Time(parsed, Now())
	if err != nil {
		return time.Time{}, false
	}
	return tm, true
}

func (t DateTimeType) MustParse(s string) time.Time {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t DateTimeType) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t DateTimeType) Compare(x1 string, x2 string) (int, bool) {
	t1, ok := t.Parse(x1)
	if !ok {
		return 0, false
	}
	t2, ok := t.Parse(x2)
	if !ok {
		return 0, false
	}
	return t1.Compare(t2), true
}

func (t DateTimeType) Format(tm time.Time) string {
	return pt.Format(dateTimeLayout, tm)
}

func PopDateTime(c Calc) time.Time     { return DateTime.MustParse(c.MustPop()) }
func PushDateTime(c Calc, r time.Time) { c.Push(DateTime.Format(r)) }

// ---

type DurationType struct{}

func (t DurationType) String() string { return "Duration" }

func (t DurationType) Parse(str string) (time.Duration, bool) {
	return types.ParseDuration(str)
}

func (t DurationType) MustParse(s string) time.Duration {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t DurationType) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t DurationType) Compare(x1 string, x2 string) (int, bool) {
	d1, ok := t.Parse(x1)
	if !ok {
		return 0, false
	}
	d2, ok := t.Parse(x2)
	if !ok {
		return 0, false
	}
	return cmp.Compare(d1, d2), true
}

func (t DurationType) Format(v time.Duration) string {
	return types.FormatDuration(v)
}

func PopDuration(c Calc) time.Duration     { return Duration.MustParse(c.MustPop()) }
func PushDuration(c Calc, r time.Duration) { c.Push(Duration.Format(r)) }

// ---

type TimeType struct{}

func (t TimeType) String() string { return "Time" }

func (t TimeType) Parse(s string) (time.Time, bool) {
	parsed, err := pt.ParseTime(s)
	if err != nil {
		return time.Time{}, false
	}
	tm, err := pt.Time(parsed, Now())
	if err != nil {
		return time.Time{}, false
	}
	return tm, true
}

func (t TimeType) MustParse(s string) time.Time {
	r, ok := t.Parse(s)
	if !ok {
		PanicExpectedType(t, s)
	}
	return r
}

func (t TimeType) Is(s string) bool {
	_, ok := t.Parse(s)
	return ok
}

func (t TimeType) Compare(x1 string, x2 string) (int, bool) {
	t1, ok := t.Parse(x1)
	if !ok {
		return 0, false
	}
	t2, ok := t.Parse(x2)
	if !ok {
		return 0, false
	}
	return t1.Compare(t2), true
}

func (t TimeType) Format(tm time.Time) string {
	return pt.Format(timeLayout, tm)
}

func PopTime(c Calc) time.Time     { return Time.MustParse(c.MustPop()) }
func PushTime(c Calc, r time.Time) { c.Push(Time.Format(r)) }
