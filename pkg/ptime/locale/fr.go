package locale

var FrMonthNamesWide = []string{
	"janvier",
	"février",
	"mars",
	"avril",
	"mai",
	"juin",
	"juillet",
	"août",
	"septembre",
	"octobre",
	"novembre",
	"décembre",
}

var FrMonthNamesAbbr = []string{
	"janv.",
	"févr.",
	"mars",
	"avr.",
	"mai",
	"juin",
	"juil.",
	"août",
	"sept.",
	"oct.",
	"nov.",
	"déc.",
}

var FrDayNamesWide = []string{
	"dimanche",
	"lundi",
	"mardi",
	"mercredi",
	"jeudi",
	"vendredi",
	"samedi",
}

var FrDayNamesAbbr = []string{
	"dim.",
	"lun.",
	"mar.",
	"mer.",
	"jeu.",
	"ven.",
	"sam.",
}

var FrZonesShort = map[string]string{
	"UTC": "+0000",
}

var FrFR = MustNew(Def{
	MonthNamesWide: FrMonthNamesWide,
	MonthNamesAbbr: FrMonthNamesAbbr,
	DayNamesWide:   FrDayNamesWide,
	DayNamesAbbr:   FrDayNamesAbbr,
	ZoneNamesShort: FrZonesShort,
	DateSep:        []string{"-", "/"},
	TimeSep:        []string{":"},
	HourSep:        []string{"h"},
	DecimalSep:     ",",
	DateTimeSep:    []string{"T"},
	UTCFlags:       []string{"Z"},
})
