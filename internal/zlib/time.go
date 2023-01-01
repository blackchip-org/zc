package zlib

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/blackchip-org/zc"
)

const (
	zoneLayout       = "MST -0700"
	timeFormatLayout = "Mon Jan 2 2006 3:04PM " + zoneLayout
)

var timeParseLayouts = []string{
	timeFormatLayout,
	time.Kitchen,
	"3:04pm",
	"3pm",
}

type timeState struct {
	loc *time.Location
	tz  string
}

func getTimeState(env *zc.Env) *timeState {
	return env.Calc.States["time"].(*timeState)
}

type timeAttrs struct {
	layout      string
	requireZone bool
}

func InitTime(env *zc.Env) error {
	loc := time.Now().Location()
	tz, _ := time.Now().Zone()
	env.Calc.States["time"] = &timeState{
		loc: loc,
		tz:  tz,
	}
	return nil
}

func parseTime(env *zc.Env, v string) (time.Time, timeAttrs, error) {
	loc := getTimeState(env).loc
	for _, layout := range timeParseLayouts {
		t, err := time.ParseInLocation(layout, v, loc)
		if err != nil {
			continue
		}
		if t.Year() == 0 {
			now := time.Now().In(loc)
			t = time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), loc)
		}
		return t, timeAttrs{layout: layout}, nil
	}
	return time.Time{}, timeAttrs{}, fmt.Errorf("expecting Time but got %v", v)
}

func isTime(env *zc.Env, v string) bool {
	_, _, err := parseTime(env, v)
	return err == nil
}

func popTime(env *zc.Env) (time.Time, timeAttrs, error) {
	s, err := env.Stack.Pop()
	if err != nil {
		return time.Time{}, timeAttrs{}, err
	}
	return parseTime(env, s)
}

func formatTime(t time.Time, attrs timeAttrs) string {
	layout := attrs.layout
	name, _ := t.Zone()

	if attrs.requireZone && !strings.Contains(layout, zoneLayout) {
		layout += " " + zoneLayout
	}
	if _, err := strconv.ParseInt(name, 10, 8); err == nil {
		layout = strings.Replace(layout, " MST", "", 1)
	}

	return t.Format(layout)
}

func pushTime(env *zc.Env, t time.Time, attrs timeAttrs) {
	env.Stack.Push(formatTime(t, attrs))
}

func parseDuration(v string) (time.Duration, error) {
	t, err := time.ParseDuration(v)
	if err != nil {
		return time.Duration(0), fmt.Errorf("expecting Duration but got %v", v)
	}
	return t, nil
}

func isDuration(v string) bool {
	_, err := parseDuration(v)
	return err == nil
}

func popDuration(env *zc.Env) (time.Duration, error) {
	a, err := env.Stack.Pop()
	if err != nil {
		return time.Duration(0), err
	}

	d, err := parseDuration(a)
	if err != nil {
		return time.Duration(0), err
	}

	return d, nil
}

func formatDuration(d time.Duration) string {
	return d.String()
}

func pushDuration(env *zc.Env, d time.Duration) {
	env.Stack.Push(formatDuration((d)))
}

func After(env *zc.Env) error {
	sb, err := env.Stack.Pop()
	if err != nil {
		return err
	}

	sa, err := env.Stack.Pop()
	if err != nil {
		return err
	}

	if isTime(env, sa) {
		return afterTime(sa, sb, env)
	} else if isDuration(sa) {
		return afterDuration(sa, sb, env)
	}
	return fmt.Errorf("expecting Time or Duration but got: %v", sa)
}

func afterTime(sa string, sb string, env *zc.Env) error {
	a, attrs, _ := parseTime(env, sa)
	b, _ := parseDuration(sb)

	z := a.Add(b)
	pushTime(env, z, attrs)
	return nil
}

func afterDuration(sa string, sb string, env *zc.Env) error {
	a, _ := parseDuration(sa)
	b, _ := parseDuration(sb)

	z := a + b
	pushDuration(env, z)
	return nil
}

func DateTime(env *zc.Env) error {
	t, attrs, err := popTime(env)
	if err != nil {
		return err
	}
	attrs.layout = timeFormatLayout
	pushTime(env, t, attrs)
	return nil
}

func In(env *zc.Env) error {
	name, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	loc, err := time.LoadLocation(name)
	if err != nil {
		return fmt.Errorf("unknown time zone: %v", name)
	}
	t, attrs, err := popTime(env)
	if err != nil {
		return err
	}
	zt := t.In(loc)
	attrs.requireZone = true
	pushTime(env, zt, attrs)
	return nil
}

func Now(env *zc.Env) error {
	loc := getTimeState(env).loc
	t := time.Now().In(loc)
	pushTime(env, t, timeAttrs{layout: timeFormatLayout})
	return nil
}

// func Offset(env *zc.Env) error {
// 	t := time.Now()
// 	_, offset := t.Zone()
// 	dur := time.Duration(offset) * time.Second
// 	str := fmt.Sprintf("%02d:%02d", int(dur.Hours()), int(dur.Minutes())%60)
// 	env.Stack.Push(str)
// 	return nil
// }

func TzSet(env *zc.Env) error {
	zone, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	loc, err := time.LoadLocation(zone)
	if err != nil {
		return fmt.Errorf("unknown time zone: %v", zone)
	}
	getTimeState(env).loc = loc
	return nil
}

func TzGet(env *zc.Env) error {
	env.Stack.Push(getTimeState(env).tz)
	return nil
}
