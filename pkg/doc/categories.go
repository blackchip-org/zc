package doc

type Category struct {
	Name   string
	Groups []string
}

var Categories = []Category{
	{"Calculator Operations", []string{
		"basic",
		"prog",
		"sci",
		"stat",
	}},
	{"Units of Measure", []string{
		"len",
		"temp",
	}},
	{"Standard Operations", []string{
		"anno",
		"bool",
		"cmp",
		"decimal",
		"format",
		"hof",
		"stack",
		"text",
		"zc",
	}},
	{"Library", []string{
		"angle",
		"color",
		"complex",
		"crypto",
		"geo",
		"rand",
		"rational",
		"time",
	}},
	{"Tables", []string{
		"entity",
		"emoji",
		"epsg",
		"si",
		"tz",
	}},
}
