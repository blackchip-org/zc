package zlib

import (
	"fmt"
	"time"

	"github.com/blackchip-org/zc"
)

var (
	defaultDateLayouts = []string{
		"Mon Jan 2 2006",
		"Jan 2 2006",
		"Jan 2",
		"01/02",
		"2 Jan 2006",
		"2006-01-02",
		"2006-002",
	}

	defaultTimeLayouts = []string{
		"3:04:05pm",
		"3:04:05PM",
		"3:04pm",
		"3:04PM",
		"15:04:05",
		"15:04",
	}

	defaultZoneLayouts = []string{
		"MST -0700",
		"-0700",
		"MST",
	}

	// defaultDateTimeLayouts = []string{
	// 	"Mon Jan 2 2006 3:04:05pm",
	// }
)

type chronoState struct {
	dateLayouts     []string
	timeLayouts     []string
	zoneLayouts     []string
	dateTimeLayouts []string
	local           *time.Location
	localZoneName   string
	travel          time.Time
}

func (t chronoState) dateFormat() string {
	return t.dateLayouts[0]
}

func (t chronoState) timeFormat() string {
	return t.timeLayouts[0]
}

func (t chronoState) zoneFormat() string {
	return t.zoneLayouts[0]
}

func (t chronoState) dateTimeFormat() string {
	return t.dateFormat() + " " + t.timeFormat() + " " + t.zoneFormat()
}

func (t chronoState) now() time.Time {
	if t.travel.IsZero() {
		return time.Now()
	}
	return t.travel
}

func getChronoState(env *zc.Env) *chronoState {
	return env.Calc.States["chrono"].(*chronoState)
}

type timeAttrs struct {
	layout      string
	requireZone bool
}

func InitChrono(env *zc.Env) error {
	loc := time.Now().Location()
	tz, _ := time.Now().Zone()
	env.Calc.States["chrono"] = &chronoState{
		local:         loc,
		localZoneName: tz,
		dateLayouts:   defaultDateLayouts,
		timeLayouts:   defaultTimeLayouts,
		zoneLayouts:   defaultZoneLayouts,
		//dateTimeLayouts: defaultDateTimeLayouts,
	}
	//rebuildFormats(env)
	return nil
}

func parseDateTime(env *zc.Env, v string) (time.Time, timeAttrs, error) {
	s := getChronoState(env)
	loc := s.local

	var layouts []string
	for _, dLayout := range s.dateLayouts {
		for _, tLayout := range s.timeLayouts {
			layouts = append(layouts, dLayout+" "+tLayout)
			for _, zLayout := range s.zoneLayouts {
				layouts = append(layouts, dLayout+" "+tLayout+" "+zLayout)
			}
		}
	}

	for _, layout := range layouts {
		t, err := time.ParseInLocation(layout, v, loc)
		if err != nil {
			continue
		}
		return t, timeAttrs{layout: layout}, nil
	}

	t, attrs, err := parseDate(env, v)
	if err == nil {
		return t, attrs, nil
	}

	t, attrs, err = parseTime(env, v)
	if err == nil {
		return t, attrs, nil
	}

	return time.Time{}, timeAttrs{}, fmt.Errorf("expecting Date, Time, or DateTime but got %v", v)
}

func parseDate(env *zc.Env, v string) (time.Time, timeAttrs, error) {
	s := getChronoState(env)
	loc := s.local

	for _, layout := range s.dateLayouts {
		t, err := time.ParseInLocation(layout, v, loc)
		if err != nil {
			continue
		}
		if t.Year() == 0 {
			now := s.now().In(loc)
			t = time.Date(now.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), loc)
		}
		return t, timeAttrs{layout: layout}, nil
	}
	return time.Time{}, timeAttrs{}, fmt.Errorf("expecting Date but got %v", v)
}

func parseTime(env *zc.Env, v string) (time.Time, timeAttrs, error) {
	s := getChronoState(env)
	loc := s.local

	var layouts []string
	for _, layout := range s.timeLayouts {
		layouts = append(layouts, layout)
		for _, zlayout := range s.zoneLayouts {
			layouts = append(layouts, layout+" "+zlayout)
		}
	}

	for _, layout := range layouts {
		t, err := time.ParseInLocation(layout, v, loc)
		if err != nil {
			continue
		}
		now := s.now().In(loc)
		nowT := time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), loc)
		return nowT, timeAttrs{layout: layout}, nil
	}
	return time.Time{}, timeAttrs{}, fmt.Errorf("expecting Time but got %v", v)
}

func isTime(env *zc.Env, v string) bool {
	_, _, err := parseTime(env, v)
	return err == nil
}

func popDate(env *zc.Env) (time.Time, timeAttrs, error) {
	s, err := env.Stack.Pop()
	if err != nil {
		return time.Time{}, timeAttrs{}, err
	}
	return parseDate(env, s)
}

func popDateTime(env *zc.Env) (time.Time, timeAttrs, error) {
	s, err := env.Stack.Pop()
	if err != nil {
		return time.Time{}, timeAttrs{}, err
	}
	return parseDateTime(env, s)
}

func popTime(env *zc.Env) (time.Time, timeAttrs, error) {
	s, err := env.Stack.Pop()
	if err != nil {
		return time.Time{}, timeAttrs{}, err
	}
	return parseTime(env, s)
}

func formatDate(env *zc.Env, t time.Time, attr timeAttrs) string {
	s := getChronoState(env)
	return t.Format(s.dateFormat())
}

func formatTime(env *zc.Env, t time.Time, attr timeAttrs) string {
	s := getChronoState(env)
	return t.Format(s.timeFormat() + " " + s.zoneFormat())
}

func formatDateTime(env *zc.Env, t time.Time, attrs timeAttrs) string {
	s := getChronoState(env)
	return t.Format(s.dateTimeFormat())

	// layout := attrs.layout
	// name, _ := t.Zone()

	// if attrs.requireZone {
	// 	foundZone := false
	// 	for _, zf := range zoneFormats {
	// 		if strings.Contains(layout, zf) {
	// 			foundZone = true
	// 			break
	// 		}
	// 	}
	// 	if !foundZone {
	// 		layout += " " + zoneFormats[0]
	// 	}
	// }

	// zoneIsOffset := false
	// if _, err := strconv.ParseInt(name, 10, 8); err == nil {
	// 	zoneIsOffset = true
	// }
	// if name == "" || zoneIsOffset {
	// 	layout = strings.Replace(layout, " MST", "", 1)
	// }
	// if name == "UTC" {
	// 	layout = strings.Replace(layout, " -0700", "", 1)
	// }

	// return t.Format(layout)
}

func pushDate(env *zc.Env, t time.Time, attrs timeAttrs) {
	env.Stack.Push(formatDate(env, t, attrs))
}

func pushTime(env *zc.Env, t time.Time, attrs timeAttrs) {
	env.Stack.Push(formatTime(env, t, attrs))
}

func pushDateTime(env *zc.Env, t time.Time, attrs timeAttrs) {
	env.Stack.Push(formatDateTime(env, t, attrs))
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

// func rebuildFormats(env *zc.Env) {
// 	s := getTimeState(env)
// 	var formats []string

// 	dateFormats := monthDayFormats
// 	if s.dayMonth {
// 		dateFormats = dayMonthFormats
// 	}
// 	timeFormats := hour12Formats
// 	if s.hour24 {
// 		timeFormats = hour24Formats
// 	}

// 	for _, df := range dateFormats {
// 		for _, wf := range weekdayFormats {
// 			dateFormats = append(dateFormats, wf+" "+df)
// 		}
// 	}
// 	for _, tf := range timeFormats {
// 		for _, zf := range zoneFormats {
// 			timeFormats = append(timeFormats, tf+" "+zf)
// 		}
// 	}

// 	var dateTimeFormats []string
// 	for _, df := range dateFormats {
// 		for _, tf := range timeFormats {
// 			dateTimeFormats = append(dateTimeFormats, df+" "+tf)
// 		}
// 	}

// 	allFormats := [][]string{
// 		dateTimeFormats,
// 		dateFormats,
// 		timeFormats,
// 		otherFormats,
// 	}

// 	for _, f := range allFormats {
// 		formats = append(formats, f...)
// 	}
// 	s.formats = formats
// }

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
	pushDateTime(env, z, attrs)
	return nil
}

func afterDuration(sa string, sb string, env *zc.Env) error {
	a, _ := parseDuration(sa)
	b, _ := parseDuration(sb)

	z := a + b
	pushDuration(env, z)
	return nil
}

// func DateLayouts(env *zc.Env) error {
// 	for env.Stack.Len() > 0 {
// 		l := env.Stack.MustPop()
// 		_, err := time.Parse(l, l)
// 		if err != nil {
// 			return fmt.Errorf("invalid format %v: %v", err)
// 		}
// 	}
// }

func Date(env *zc.Env) error {
	s := getChronoState(env)
	t, attrs, err := popDateTime(env)
	if err != nil {
		return err
	}
	attrs.layout = s.dateFormat()
	pushDate(env, t, attrs)
	return nil
}

func DateTime(env *zc.Env) error {
	s := getChronoState(env)
	t, attrs, err := popDateTime(env)
	if err != nil {
		return err
	}
	attrs.layout = s.dateTimeFormat()
	pushDateTime(env, t, attrs)
	return nil
}

// func FormatsGet(env *zc.Env) error {
// 	s := getTimeState(env)
// 	for _, layout := range s.formats {
// 		env.Stack.Push(layout)
// 	}
// 	return nil
// }

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
	pushDateTime(env, zt, attrs)
	return nil
}

func Now(env *zc.Env) error {
	s := getChronoState(env)
	attrs := timeAttrs{layout: s.dateTimeFormat()}

	if !s.travel.IsZero() {
		pushDateTime(env, s.travel, attrs)
		return nil
	}

	loc := getChronoState(env).local
	t := time.Now().In(loc)
	pushDateTime(env, t, attrs)
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

func Time(env *zc.Env) error {
	s := getChronoState(env)
	t, attrs, err := popDateTime(env)
	if err != nil {
		return err
	}
	attrs.layout = s.timeFormat()
	pushTime(env, t, attrs)
	return nil
}

func Travel(env *zc.Env) error {
	s := getChronoState(env)
	t, _, err := popDateTime(env)
	if err != nil {
		return err
	}
	s.travel = t
	return nil
}

func TravelEnd(env *zc.Env) error {
	s := getChronoState(env)
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
	getChronoState(env).local = loc
	return nil
}

func LocalGet(env *zc.Env) error {
	env.Stack.Push(getChronoState(env).localZoneName)
	return nil
}
