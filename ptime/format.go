package ptime

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/blackchip-org/zc/ptime/locale"
)

var formatTable = map[string]func(*locale.Locale, string, time.Time) string{
	"weekday":     formatWeekday,
	"year":        formatYear,
	"month":       formatMonth,
	"day":         formatDay,
	"hour":        formatHour,
	"minute":      formatMinute,
	"second":      formatSecond,
	"period":      formatPeriod,
	"zone":        formatZone,
	"offset":      formatOffset,
	"zone-offset": formatZoneOffset,
	"offset-zone": formatOffsetZone,
}

const (
	badField  = "!(BADFIELD)"
	badFormat = "!(BADFORMAT)"
)

func Format(loc *locale.Locale, layout string, t time.Time) string {
	src := []rune(layout)
	i := 0
	var result strings.Builder

	scanFormat := func() string {
		var format strings.Builder
		i++
		for i < len(layout) {
			if src[i] == ']' {
				return format.String()
			}
			format.WriteRune(src[i])
			i++
		}
		return format.String()
	}

	scanField := func() (string, string) {
		var name strings.Builder
		var format string

		i++
		for i < len(layout) {
			if src[i] == '/' {
				format = scanFormat()
				break
			}
			if src[i] == ']' {
				break
			}
			name.WriteRune(src[i])
			i++
		}
		if src[i] != ']' {
			return "", ""
		}
		return name.String(), format
	}

	for i < len(layout) {
		if src[i] == '[' {
			name, format := scanField()
			fn, ok := formatTable[name]
			if !ok {
				result.WriteString(badField)
			} else {
				result.WriteString(fn(loc, format, t))
			}
		} else {
			result.WriteRune(src[i])
		}
		i++
	}
	return strings.TrimSpace(result.String())
}

func formatWeekday(loc *locale.Locale, format string, t time.Time) string {
	switch format {
	case "", "wide":
		return loc.DayNamesWide[t.Weekday()]
	case "abbr":
		return loc.DayNamesAbbr[t.Weekday()]
	}
	return badFormat
}

func formatYear(loc *locale.Locale, format string, t time.Time) string {
	switch format {
	case "":
		return fmt.Sprintf("%04d", t.Year())
	case "2":
		return fmt.Sprintf("%02d", t.Year()%100)
	}
	return badFormat
}

func formatMonth(loc *locale.Locale, format string, t time.Time) string {
	switch format {
	case "":
		return strconv.Itoa(int(t.Month()))
	case "2":
		return fmt.Sprintf("%2d", int(t.Month()))
	case "02":
		return fmt.Sprintf("%02d", int(t.Month()))
	case "abbr":
		return loc.MonthNamesAbbr[int(t.Month())-1]
	case "wide", "name":
		return loc.MonthNamesWide[int(t.Month())-1]
	}
	return badFormat
}

func formatDay(loc *locale.Locale, format string, t time.Time) string {
	switch format {
	case "":
		return strconv.Itoa(t.Day())
	case "2":
		return fmt.Sprintf("%2d", t.Day())
	case "02":
		return fmt.Sprintf("%02d", t.Day())
	case "year":
		return fmt.Sprintf("%03d", t.YearDay())
	}
	return badFormat
}

func formatHour(loc *locale.Locale, format string, t time.Time) string {
	switch format {
	case "", "24":
		return strconv.Itoa(t.Hour())
	case "12":
		h := t.Hour()
		if h > 12 {
			h = h - 12
		}
		return strconv.Itoa(h)
	}
	return badFormat
}

func formatMinute(loc *locale.Locale, format string, t time.Time) string {
	switch format {
	case "":
		return fmt.Sprintf("%02d", t.Minute())
	}
	return badFormat
}

func formatSecond(loc *locale.Locale, format string, t time.Time) string {
	switch format {
	case "", "0":
		return fmt.Sprintf("%02d", t.Second())
	}

	_, err := strconv.Atoi(format)
	if err != nil {
		return badFormat
	}
	s := float64(t.Second()) + (float64(t.Nanosecond()) * 1e-9)
	spec := "%." + format + "f"
	return fmt.Sprintf(spec, s)
}

func formatPeriod(loc *locale.Locale, format string, t time.Time) string {
	period := locale.AM
	if t.Hour() >= 12 {
		period = locale.PM
	}

	switch format {
	case "", "abbr":
		return loc.PeriodNamesAbbr.Main(period)
	case "alt", "abbr-alt":
		return loc.PeriodNamesAbbr.Alt(period)
	case "narrow":
		return loc.PeriodNamesNarrow.Main(period)
	}
	return badFormat
}

func formatZone(loc *locale.Locale, format string, t time.Time) string {
	zone, _ := t.Zone()
	switch format {
	case "":
		return zone
	}
	return badFormat
}

func formatOffset(loc *locale.Locale, format string, t time.Time) string {
	_, offset := t.Zone()

	switch format {
	case "":
		return FormatOffset(offset, "")
	case ":":
		return FormatOffset(offset, ":")
	}
	return badFormat
}

func formatZoneOffset(loc *locale.Locale, format string, t time.Time) string {
	zone := formatZone(loc, "", t)
	offset := formatOffset(loc, format, t)
	if o, ok := loc.Offsets[loc.Key(zone)]; ok && o == 0 {
		return zone
	}
	return zone + " " + offset
}

func formatOffsetZone(loc *locale.Locale, format string, t time.Time) string {
	zone := formatZone(loc, "", t)
	offset := formatOffset(loc, format, t)
	if o, ok := loc.Offsets[loc.Key(zone)]; ok && o == 0 {
		return zone
	}
	return offset + " " + zone
}
