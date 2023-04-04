package time_

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/ptime"
	"github.com/blackchip-org/zc/ptime/locale"
)

const (
	defaultDateLayout     = "[weekday/abbr] [month/abbr] [day] [year]"
	defaultTimeLayout     = "[hour/12]:[minute]:[second][period/alt] [offset-zone]"
	defaultDateTimeLayout = defaultDateLayout + " " + defaultTimeLayout
)

type timeState struct {
	locale         *locale.Locale
	p              *ptime.P
	dateLayout     string
	timeLayout     string
	dateTimeLayout string
	local          *time.Location
	localZone      string
	travel         time.Time
}

func (s timeState) formatDate(t time.Time) string {
	return ptime.Format(s.locale, s.dateLayout, t)
}

func (s timeState) formatTime(t time.Time) string {
	return ptime.Format(s.locale, s.timeLayout, t)
}

func (s timeState) formatDateTime(t time.Time) string {
	return ptime.Format(s.locale, s.dateTimeLayout, t)
}

func (t timeState) now() time.Time {
	if t.travel.IsZero() {
		return time.Now().In(t.local)
	}
	return t.travel.In(t.local)
}

func getTimeState(env *zc.Env) *timeState {
	return env.Calc.States["time"].(*timeState)
}

func InitTime(env *zc.Env) error {
	loc := time.Now().Location()
	tz, _ := time.Now().Zone()
	env.Calc.States["time"] = &timeState{
		locale:         locale.EnUS,
		p:              ptime.For(locale.EnUS),
		local:          loc,
		localZone:      tz,
		dateLayout:     defaultDateLayout,
		timeLayout:     defaultTimeLayout,
		dateTimeLayout: defaultDateTimeLayout,
	}
	return nil
}

func parseDateTime(env *zc.Env, v string) (time.Time, error) {
	s := getTimeState(env)

	parsed, err := s.p.Parse(v)
	if err != nil {
		return time.Time{}, fmt.Errorf("expected DateTime, Time, or Date but got: %v", v)
	}
	t, err := s.p.Time(parsed, s.now())
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to parse: %v", v)
	}
	return t, nil
}

func parseDate(env *zc.Env, v string) (time.Time, error) {
	s := getTimeState(env)

	parsed, err := s.p.ParseDate(v)
	if err != nil {
		return time.Time{}, fmt.Errorf("expected DateTime, Time, or Date but got: %v", v)
	}
	t, err := s.p.Time(parsed, s.now())
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to parse: %v", v)
	}
	return t, nil
}

func parseTime(env *zc.Env, v string) (time.Time, error) {
	s := getTimeState(env)

	parsed, err := s.p.Parse(v)
	if err != nil {
		return time.Time{}, fmt.Errorf("expected DateTime, Time, or Date but got: %v", v)
	}
	t, err := s.p.Time(parsed, s.now())
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to parse: %v", v)
	}
	return t, nil
}

func isTime(env *zc.Env, v string) bool {
	_, err := parseTime(env, v)
	return err == nil
}

func popDate(env *zc.Env) (time.Time, error) {
	s, err := env.Stack.Pop()
	if err != nil {
		return time.Time{}, err
	}
	return parseDate(env, s)
}

func popDateTime(env *zc.Env) (time.Time, error) {
	s, err := env.Stack.Pop()
	if err != nil {
		return time.Time{}, err
	}
	return parseDateTime(env, s)
}

func popTime(env *zc.Env) (time.Time, error) {
	s, err := env.Stack.Pop()
	if err != nil {
		return time.Time{}, err
	}
	return parseTime(env, s)
}

func formatDate(env *zc.Env, t time.Time) string {
	s := getTimeState(env)
	return s.formatDate(t)
}

func formatTime(env *zc.Env, t time.Time) string {
	s := getTimeState(env)
	return s.formatTime(t)
}

func formatDateTime(env *zc.Env, t time.Time) string {
	s := getTimeState(env)
	return s.formatDateTime(t)
}

func pushDate(env *zc.Env, t time.Time) {
	env.Stack.Push(formatDate(env, t))
}

func pushTime(env *zc.Env, t time.Time) {
	env.Stack.Push(formatTime(env, t))
}

func pushDateTime(env *zc.Env, t time.Time) {
	env.Stack.Push(formatDateTime(env, t))
}

func parseDuration(v string) (time.Duration, error) {
	d, err := time.ParseDuration(v)
	if err != nil {
		return time.Duration(0), fmt.Errorf("expecting Duration but got %v", v)
	}
	return d, nil
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
	dsec := int(d.Truncate(time.Second).Seconds())
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

func pushDuration(env *zc.Env, d time.Duration) {
	env.Stack.Push(formatDuration((d)))
}

func AddDuration(env *zc.Env) error {
	var (
		t   time.Time
		d   time.Duration
		err error
	)

	a, err := env.Stack.Pop()
	if err != nil {
		return err
	}

	b, err := env.Stack.Pop()
	if err != nil {
		return err
	}

	t, err = parseDateTime(env, a)
	if err != nil {
		t, err = parseDateTime(env, b)
		if err != nil {
			return fmt.Errorf("expecting DateTime but got '%v' and '%v'", a, b)
		}
	}

	d, err = parseDuration(a)
	if err != nil {
		d, err = parseDuration(b)
		if err != nil {
			return fmt.Errorf("expecting Duration but got '%v' and '%v'", a, b)
		}
	}

	z := t.Add(d)
	pushDateTime(env, z)
	return nil
}

func Date(env *zc.Env) error {
	t, err := popDateTime(env)
	if err != nil {
		return err
	}
	pushDate(env, t)
	return nil
}

func DateLayout(env *zc.Env) error {
	s := getTimeState(env)
	layout, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	s.dateLayout = layout
	return nil
}

func DateLayoutGet(env *zc.Env) error {
	s := getTimeState(env)
	env.Stack.Push(s.dateLayout)
	return nil
}

func DateTime(env *zc.Env) error {
	t, err := popDateTime(env)
	if err != nil {
		return err
	}
	pushDateTime(env, t)
	return nil
}

func DateTimeLayout(env *zc.Env) error {
	s := getTimeState(env)
	layout, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	s.dateTimeLayout = layout
	return nil
}

func DateTimeLayoutGet(env *zc.Env) error {
	s := getTimeState(env)
	env.Stack.Push(s.dateTimeLayout)
	return nil
}

func DayYear(env *zc.Env) error {
	t, err := popTime(env)
	if err != nil {
		return err
	}
	r := strconv.Itoa(t.YearDay())
	env.Stack.Push(r)
	return nil
}

func Hours(env *zc.Env) error {
	d, err := popDuration(env)
	if err != nil {
		return err
	}
	z := d.Hours()
	env.Stack.PushFloat(z)
	return nil
}

func Local(env *zc.Env) error {
	s := getTimeState(env)

	zone, err := env.Stack.Pop()
	if err != nil {
		return err
	}

	var loc *time.Location
	offset, ok := s.locale.Offsets[s.locale.Key(zone)]
	if ok {
		zone = s.locale.DisplayNames[s.locale.Key(zone)]
		loc = time.FixedZone(zone, offset)
	} else {
		loc, err = time.LoadLocation(zone)
		if err != nil {
			return fmt.Errorf("unknown time zone: %v", zone)
		}
	}
	s.local = loc
	s.localZone = zone
	env.Calc.Info = "local time zone is now " + zc.Quote(s.localZone)
	return nil
}

func LocalGet(env *zc.Env) error {
	env.Stack.Push(getTimeState(env).localZone)
	return nil
}

func Minutes(env *zc.Env) error {
	d, err := popDuration(env)
	if err != nil {
		return err
	}
	z := d.Minutes()
	env.Stack.PushFloat(z)
	return nil
}

func Now(env *zc.Env) error {
	s := getTimeState(env)
	t := s.now()
	pushDateTime(env, t)
	return nil
}

func Seconds(env *zc.Env) error {
	d, err := popDuration(env)
	if err != nil {
		return err
	}
	z := d.Seconds()
	env.Stack.PushFloat(z)
	return nil
}

func SubtractTime(env *zc.Env) error {
	b, err := popDateTime(env)
	if err != nil {
		return err
	}

	a, err := popDateTime(env)
	if err != nil {
		return err
	}

	z := a.Sub(b)
	pushDuration(env, z)
	return nil
}

func Time_(env *zc.Env) error {
	t, err := popDateTime(env)
	if err != nil {
		return err
	}
	pushTime(env, t)
	return nil
}

func TimeLayout(env *zc.Env) error {
	s := getTimeState(env)
	layout, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	s.timeLayout = layout
	return nil
}

func TimeLayoutGet(env *zc.Env) error {
	s := getTimeState(env)
	env.Stack.Push(s.timeLayout)
	return nil
}

func TimeZone(env *zc.Env) error {
	s := getTimeState(env)

	name, err := env.Stack.Pop()
	if err != nil {
		return err
	}

	t, err := popTime(env)
	if err != nil {
		return err
	}

	var loc *time.Location
	offset, ok := s.locale.Offsets[s.locale.Key(name)]
	if ok {
		loc = time.FixedZone(name, offset)
	} else {
		loc, err = time.LoadLocation(name)
		if err != nil {
			return fmt.Errorf("unknown time zone: %v", name)
		}
	}

	z := t.In(loc)
	pushDateTime(env, z)
	return nil
}

func Travel(env *zc.Env) error {
	s := getTimeState(env)
	t, err := popDateTime(env)
	if err != nil {
		return err
	}
	s.travel = t
	env.Calc.Info = "now set to " + zc.Quote(s.formatDateTime(t))
	return nil
}

func TravelEnd(env *zc.Env) error {
	s := getTimeState(env)
	s.travel = time.Time{}
	return nil
}
