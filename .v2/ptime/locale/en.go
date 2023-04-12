package locale

var EnMonthNamesWide = []string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

var EnMonthNamesAbbr = []string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

var EnDayNamesWide = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

var EnDayNamesAbbr = []string{
	"Sun",
	"Mon",
	"Tue",
	"Wed",
	"Thr",
	"Fri",
	"Sat",
}

var EnPeriodNamesAbbr = String2D{
	AM:       []string{"AM", "am"},
	PM:       []string{"PM", "pm"},
	Midnight: []string{"midnight"},
	Noon:     []string{"noon"},
}

var EnPeriodNamesNarrow = String2D{
	AM:       []string{"a"},
	PM:       []string{"p"},
	Midnight: []string{"mi"},
	Noon:     []string{"n"},
}

var EnUSZonesShort = map[string]string{
	"EST": "-0500",
	"CST": "-0600",
	"MST": "-0700",
	"PST": "-0800",
	"EDT": "-0400",
	"CDT": "-0500",
	"MDT": "-0600",
	"PDT": "-0700",
	"UTC": "+0000",
}

var EnUS = MustNew(Def{
	MonthDayOrder:     true,
	MonthNamesWide:    EnMonthNamesWide,
	MonthNamesAbbr:    EnMonthNamesAbbr,
	DayNamesWide:      EnDayNamesWide,
	DayNamesAbbr:      EnDayNamesAbbr,
	PeriodNamesAbbr:   EnPeriodNamesAbbr,
	PeriodNamesNarrow: EnPeriodNamesNarrow,
	ZoneNamesShort:    EnUSZonesShort,
	DateSep:           []string{"-", "/"},
	TimeSep:           []string{":"},
	DecimalSep:        ".",
	DateTimeSep:       []string{"T"},
	UTCFlags:          []string{"Z"},
})
