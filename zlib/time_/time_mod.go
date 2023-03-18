package time_

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "time",
		Include:    true,
		ScriptPath: "zc:zlib/time_/time.zc",
		Init:       InitTime,
		Natives: map[string]zc.CalcFunc{
			"add-duration":      AddDuration,
			"date":              Date,
			"date-layout":       DateLayout,
			"date-layout=":      DateLayoutGet,
			"date-time":         DateTime,
			"date-time-layout":  DateTimeLayout,
			"date-time-layout=": DateTimeLayoutGet,
			"day-year":          DayYear,
			"hours":             Hours,
			"local":             Local,
			"local=":            LocalGet,
			"minutes":           Minutes,
			"now":               Now,
			"seconds":           Seconds,
			"subtract-time":     SubtractTime,
			"time":              Time_,
			"time-layout":       TimeLayout,
			"time-layout=":      TimeLayoutGet,
			"time-zone":         TimeZone,
			"travel":            Travel,
			"travel-end":        TravelEnd,
		},
	}
)
