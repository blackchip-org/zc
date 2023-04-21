package zc

import (
	"fmt"
	"math"
	"time"

	"github.com/blackchip-org/zc/pkg/ptime"
	"github.com/blackchip-org/zc/pkg/ptime/locale"
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

func (t DateType) Parse(s string) (time.Time, error) {
	parsed, err := pt.ParseDate(s)
	if err != nil {
		return time.Time{}, ErrUnexpectedType(t, s)
	}
	tm, err := pt.Time(parsed, Now())
	if err != nil {
		return time.Time{}, ErrUnexpectedType(t, s)
	}
	return tm, nil
}

func (t DateType) MustParse(s string) time.Time {
	r, err := t.Parse(s)
	if err != nil {
		panic(err)
	}
	return r
}

func (t DateType) Is(s string) bool {
	_, err := t.Parse(s)
	return err == nil
}

func (t DateType) Format(tm time.Time) string {
	return pt.Format(dateLayout, tm)
}

func PopDate(c Calc) time.Time     { return Date.MustParse(c.MustPop()) }
func PushDate(c Calc, r time.Time) { c.Push(Date.Format(r)) }

// ---

type DateTimeType struct{}

func (t DateTimeType) String() string { return "DateTime" }

func (t DateTimeType) Parse(s string) (time.Time, error) {
	parsed, err := pt.Parse(s)
	if err != nil {
		return time.Time{}, ErrUnexpectedType(t, s)
	}
	tm, err := pt.Time(parsed, Now())
	if err != nil {
		return time.Time{}, ErrUnexpectedType(t, s)
	}
	return tm, nil
}

func (t DateTimeType) MustParse(s string) time.Time {
	r, err := t.Parse(s)
	if err != nil {
		panic(err)
	}
	return r
}

func (t DateTimeType) Is(s string) bool {
	_, err := t.Parse(s)
	return err == nil
}

func (t DateTimeType) Format(tm time.Time) string {
	return pt.Format(dateTimeLayout, tm)
}

func PopDateTime(c Calc) time.Time     { return DateTime.MustParse(c.MustPop()) }
func PushDateTime(c Calc, r time.Time) { c.Push(DateTime.Format(r)) }

// ---

type DurationType struct{}

func (t DurationType) String() string { return "Duration" }

func (t DurationType) Parse(s string) (time.Duration, error) {
	d, err := time.ParseDuration(s)
	if err != nil {
		return time.Duration(0), ErrUnexpectedType(t, s)
	}
	return d, nil

}

func (t DurationType) MustParse(s string) time.Duration {
	r, err := t.Parse(s)
	if err != nil {
		panic(err)
	}
	return r
}

func (t DurationType) Is(s string) bool {
	_, err := t.Parse(s)
	return err == nil
}

func (t DurationType) Format(v time.Duration) string {
	dsec := int(v.Truncate(time.Second).Seconds())
	sec := math.Abs(float64(dsec % 60))
	min := math.Abs(float64(dsec / 60 % 60))
	hrs := dsec / 3600

	if min == 0 && sec == 0 {
		return fmt.Sprintf("%vh", hrs)
	}
	if sec == 0 {
		return fmt.Sprintf("%vh%vm", hrs, min)
	}
	return fmt.Sprintf("%vh%vm%vs", hrs, min, sec)
}

func PopDuration(c Calc) time.Duration     { return Duration.MustParse(c.MustPop()) }
func PushDuration(c Calc, r time.Duration) { c.Push(Duration.Format(r)) }

// ---

type TimeType struct{}

func (t TimeType) String() string { return "Time" }

func (t TimeType) Parse(s string) (time.Time, error) {
	parsed, err := pt.ParseTime(s)
	if err != nil {
		return time.Time{}, ErrUnexpectedType(t, s)
	}
	tm, err := pt.Time(parsed, Now())
	if err != nil {
		return time.Time{}, ErrUnexpectedType(t, s)
	}
	return tm, nil
}

func (t TimeType) MustParse(s string) time.Time {
	r, err := t.Parse(s)
	if err != nil {
		panic(err)
	}
	return r
}

func (t TimeType) Is(s string) bool {
	_, err := t.Parse(s)
	return err == nil
}

func (t TimeType) Format(tm time.Time) string {
	return pt.Format(timeLayout, tm)
}

func PopTime(c Calc) time.Time     { return Time.MustParse(c.MustPop()) }
func PushTime(c Calc, r time.Time) { c.Push(Time.Format(r)) }
