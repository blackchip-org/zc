package zlib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/blackchip-org/zc"
)

var (
	ordRegexp      = regexp.MustCompile(`(\d{4})-(\d+)$`)
	weekdayFormats = []string{
		"Mon",
		"Monday",
	}
	monthDayFormats = []string{
		"Jan 2 2006",
		"Jan 2",
		"01/02",
	}
	dayMonthFormats = []string{
		"2 Jan 2006",
		"2 Jan",
		"02/01",
	}
	hour12Formats = []string{
		"3:04:05PM",
		"3:04:05pm",
		"3:04PM",
		"3:04pm",
	}
	hour24Formats = []string{
		"15:04:05",
		"15:04",
		"1504",
	}
	zoneFormats = []string{
		"MST -0700",
		"-0700",
		"MST",
	}
	otherFormats = []string{
		"2006-01-02",
	}
)

type timeState struct {
	local              *time.Location
	localZoneName      string
	hour24             bool
	dayMonth           bool
	zoneFormatOverride string
	dateFormatOverride string
	timeFormatOverride string
	formats            []string
	travel             time.Time
}

func (t timeState) zoneFormat() string {
	if t.zoneFormatOverride != "" {
		return t.zoneFormatOverride
	}
	return zoneFormats[0]
}

func (t timeState) dateFormat() string {
	if t.dateFormatOverride != "" {
		return t.dateFormatOverride
	}
	if t.dayMonth {
		return dayMonthFormats[0]
	}
	return monthDayFormats[0]
}

func (t timeState) timeFormat() string {
	if t.timeFormatOverride != "" {
		return t.timeFormatOverride
	}
	if t.hour24 {
		return hour24Formats[0]
	}
	return hour12Formats[0]
}

func (t timeState) dateTimeFormat() string {
	return weekdayFormats[0] + " " + t.dateFormat() + " " + t.timeFormat() + " " + t.zoneFormat()
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
		local:         loc,
		localZoneName: tz,
	}
	rebuildFormats(env)
	return nil
}

func parseTime(env *zc.Env, v string) (time.Time, timeAttrs, error) {
	s := getTimeState(env)
	loc := s.local

	if matches := ordRegexp.FindStringSubmatch(v); matches != nil {
		year, _ := strconv.Atoi(matches[1])
		days, _ := strconv.Atoi(matches[2])
		t := time.Date(year, 1, days, 0, 0, 0, 0, loc)
		return t, timeAttrs{}, nil
	}

	for _, layout := range s.formats {
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
	return time.Time{}, timeAttrs{}, fmt.Errorf("expecting Date, Time, or DateTime but got %v", v)
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

	if attrs.requireZone {
		foundZone := false
		for _, zf := range zoneFormats {
			if strings.Contains(layout, zf) {
				foundZone = true
				break
			}
		}
		if !foundZone {
			layout += " " + zoneFormats[0]
		}
	}
	if _, err := strconv.ParseInt(name, 10, 8); err == nil {
		layout = strings.Replace(layout, " MST", "", 1)
	}
	if name == "UTC" {
		layout = strings.Replace(layout, " -0700", "", 1)
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

func rebuildFormats(env *zc.Env) {
	s := getTimeState(env)
	var formats []string

	dateFormats := monthDayFormats
	if s.dayMonth {
		dateFormats = dayMonthFormats
	}
	timeFormats := hour12Formats
	if s.hour24 {
		timeFormats = hour24Formats
	}

	for _, df := range dateFormats {
		for _, wf := range weekdayFormats {
			dateFormats = append(dateFormats, wf+" "+df)
		}
	}
	for _, tf := range timeFormats {
		for _, zf := range zoneFormats {
			timeFormats = append(timeFormats, tf+" "+zf)
		}
	}

	var dateTimeFormats []string
	for _, df := range dateFormats {
		for _, tf := range timeFormats {
			dateTimeFormats = append(dateTimeFormats, df+" "+tf)
		}
	}

	allFormats := [][]string{
		dateTimeFormats,
		dateFormats,
		timeFormats,
		otherFormats,
	}

	for _, f := range allFormats {
		formats = append(formats, f...)
	}
	s.formats = formats
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
	s := getTimeState(env)
	t, attrs, err := popTime(env)
	if err != nil {
		return err
	}
	attrs.layout = s.dateTimeFormat()
	pushTime(env, t, attrs)
	return nil
}

func FormatsGet(env *zc.Env) error {
	s := getTimeState(env)
	for _, layout := range s.formats {
		env.Stack.Push(layout)
	}
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
	s := getTimeState(env)
	attrs := timeAttrs{layout: s.dateTimeFormat()}

	if !s.travel.IsZero() {
		pushTime(env, s.travel, attrs)
		return nil
	}

	loc := getTimeState(env).local
	t := time.Now().In(loc)
	pushTime(env, t, attrs)
	return nil
}

func Ord(env *zc.Env) error {
	t, _, err := popTime(env)
	if err != nil {
		return err
	}
	r := fmt.Sprintf("%04d-%03d", t.Year(), t.YearDay())
	env.Stack.Push(r)
	return nil
}

func Travel(env *zc.Env) error {
	s := getTimeState(env)
	t, _, err := popTime(env)
	if err != nil {
		return err
	}
	s.travel = t
	return nil
}

func TravelEnd(env *zc.Env) error {
	s := getTimeState(env)
	s.travel = time.Time{}
	return nil
}

func Local(env *zc.Env) error {
	zone, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	loc, err := time.LoadLocation(zone)
	if err != nil {
		return fmt.Errorf("unknown time zone: %v", zone)
	}
	getTimeState(env).local = loc
	return nil
}

func LocalGet(env *zc.Env) error {
	env.Stack.Push(getTimeState(env).localZoneName)
	return nil
}
