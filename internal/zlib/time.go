package zlib

import (
	"fmt"
	"time"

	"github.com/blackchip-org/ptime"
	"github.com/blackchip-org/ptime/locale"
	"github.com/blackchip-org/zc"
)

const (
	defaultDateLayout     = "[weekday/abbr] [month/abbr] [day] [year]"
	defaultTimeLayout     = "[hour/12]:[minute]:[second][period/alt] [offset-zone]"
	defaultDateTimeLayout = defaultDateLayout + " " + defaultTimeLayout
)

type timeState struct {
	locale         *locale.Locale
	parser         *ptime.Parser
	dateLayout     string
	timeLayout     string
	dateTimeLayout string
	local          *time.Location
	localZoneName  string
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
		return time.Now()
	}
	return t.travel
}

func getTimeState(env *zc.Env) *timeState {
	return env.Calc.States["time"].(*timeState)
}

func InitTime(env *zc.Env) error {
	loc := time.Now().Location()
	tz, _ := time.Now().Zone()
	env.Calc.States["time"] = &timeState{
		locale:         locale.EnUS,
		parser:         ptime.NewParser(locale.EnUS),
		local:          loc,
		localZoneName:  tz,
		dateLayout:     defaultDateLayout,
		timeLayout:     defaultTimeLayout,
		dateTimeLayout: defaultDateTimeLayout,
	}
	return nil
}

func parseDateTime(env *zc.Env, v string) (time.Time, error) {
	s := getTimeState(env)

	parsed, err := s.parser.Parse(v)
	if err != nil {
		return time.Time{}, fmt.Errorf("expected DateTime, Time, or Date but got: %v", v)
	}
	t, err := ptime.Time(parsed, s.now())
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to parse: %v", v)
	}
	return t, nil
}

func parseDate(env *zc.Env, v string) (time.Time, error) {
	s := getTimeState(env)

	parsed, err := s.parser.ParseDate(v)
	if err != nil {
		return time.Time{}, fmt.Errorf("expected DateTime, Time, or Date but got: %v", v)
	}
	t, err := ptime.Time(parsed, s.now())
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to parse: %v", v)
	}
	return t, nil
}

func parseTime(env *zc.Env, v string) (time.Time, error) {
	s := getTimeState(env)

	parsed, err := s.parser.Parse(v)
	if err != nil {
		return time.Time{}, fmt.Errorf("expected DateTime, Time, or Date but got: %v", v)
	}
	t, err := ptime.Time(parsed, s.now())
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
	sec := dsec % 60
	min := dsec / 60 % 60
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

func Now(env *zc.Env) error {
	s := getTimeState(env)
	t := s.now()
	pushDateTime(env, t)
	return nil
}

func Ord(env *zc.Env) error {
	t, err := popTime(env)
	if err != nil {
		return err
	}
	r := fmt.Sprintf("%04d-%03d", t.Year(), t.YearDay())
	env.Stack.Push(r)
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
	offset, ok := s.locale.Offsets[name]
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
	return nil
}

func TravelEnd(env *zc.Env) error {
	s := getTimeState(env)
	s.travel = time.Time{}
	return nil
}
