package locale

import (
	"fmt"
	"strconv"
	"strings"
)

type String2D [][]string

func (s String2D) Main(index int) string {
	if len(s[index]) == 0 {
		return ""
	}
	return s[index][0]
}

func (s String2D) Alt(index int) string {
	switch len(s[index]) {
	case 0:
		return ""
	case 1:
		return s[index][0]
	default:
		return s[index][1]
	}
}

const (
	AM = iota
	PM
	Noon
	Midnight
)

type Def struct {
	MonthDayOrder     bool
	MonthNamesWide    []string
	MonthNamesAbbr    []string
	DayNamesWide      []string
	DayNamesAbbr      []string
	PeriodNamesAbbr   String2D
	PeriodNamesNarrow String2D
	ZoneNamesShort    map[string]string
	DateSep           []string
	TimeSep           []string
	HourSep           []string
	DecimalSep        string
	DateTimeSep       []string
	UTCFlags          []string
}

type Locale struct {
	Def
	MonthNum     map[string]int
	DayNum       map[string]int
	PeriodNum    map[string]int
	Offsets      map[string]int
	DisplayNames map[string]string
}

func New(def Def) (*Locale, error) {
	l := &Locale{
		Def:          def,
		MonthNum:     make(map[string]int),
		DayNum:       make(map[string]int),
		PeriodNum:    make(map[string]int),
		Offsets:      make(map[string]int),
		DisplayNames: make(map[string]string),
	}

	if len(def.MonthNamesAbbr) != 12 {
		return nil, fmt.Errorf("invalid number of month names (abbreviated)")
	}
	if len(def.MonthNamesWide) != 12 {
		return nil, fmt.Errorf("invalid number of month names (wide)")
	}
	for i := 0; i < 12; i++ {
		abbr := def.MonthNamesAbbr[i]
		wide := def.MonthNamesWide[i]
		abbrKey := l.Key(abbr)
		wideKey := l.Key(wide)

		l.MonthNum[abbrKey] = i + 1
		l.MonthNum[wideKey] = i + 1
		l.DisplayNames[abbrKey] = abbr
		l.DisplayNames[wideKey] = wide
	}

	if len(def.DayNamesAbbr) != 7 {
		return nil, fmt.Errorf("invalid number of day names (abbreviated)")
	}
	if len(def.DayNamesWide) != 7 {
		return nil, fmt.Errorf("invalid number of day names (wide)")
	}
	for i := 0; i < 7; i++ {
		abbr := def.DayNamesAbbr[i]
		wide := def.DayNamesWide[i]
		abbrKey := l.Key(abbr)
		wideKey := l.Key(wide)

		l.DayNum[abbrKey] = i
		l.DayNum[wideKey] = i
		l.DisplayNames[abbrKey] = abbr
		l.DisplayNames[wideKey] = wide
	}

	for i, names := range def.PeriodNamesAbbr {
		for _, name := range names {
			nameKey := l.Key(name)
			l.PeriodNum[nameKey] = i
			l.DisplayNames[nameKey] = name
		}
	}
	for i, names := range def.PeriodNamesNarrow {
		for _, name := range names {
			nameKey := l.Key(name)
			l.PeriodNum[nameKey] = i
			l.DisplayNames[nameKey] = name
		}
	}

	for zone, offset := range def.ZoneNamesShort {
		zoneKey := l.Key(zone)
		runes := []rune(offset)
		if len(runes) != 5 {
			return nil, fmt.Errorf("invalid offset: %v", offset)
		}

		var sign int
		switch runes[0] {
		case '+':
			sign = 1
		case '-':
			sign = -1
		default:
			return nil, fmt.Errorf("invalid offset: %v", offset)
		}

		hrs, err := strconv.Atoi(string(runes[1:3]))
		if err != nil {
			return nil, fmt.Errorf("invalid offset: %v", offset)
		}
		min, err := strconv.Atoi(string(runes[3:5]))
		if err != nil {
			return nil, fmt.Errorf("invalid offset: %v", offset)
		}
		offset := sign*hrs*3600 + min*60
		l.Offsets[zoneKey] = offset
		l.DisplayNames[zoneKey] = zone
	}
	for _, flag := range l.UTCFlags {
		flagKey := l.Key(flag)
		l.Offsets[flagKey] = 0
		l.DisplayNames[flagKey] = flag
	}

	return l, nil
}

func (l *Locale) Key(v string) string {
	v = strings.ReplaceAll(v, ".", "")
	return strings.ToLower(v)
}

func MustNew(def Def) *Locale {
	l, err := New(def)
	if err != nil {
		panic(err)
	}
	return l
}

var table = map[string]*Locale{
	"en-US": EnUS,
	"fr-FR": FrFR,
}

func Lookup(name string) (*Locale, bool) {
	l, ok := table[name]
	return l, ok
}

func MustLookup(name string) *Locale {
	l, ok := Lookup(name)
	if !ok {
		panic(fmt.Sprintf("locale '%v' not found", name))
	}
	return l
}
