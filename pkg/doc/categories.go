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
		"stack",
		"stat",
		"zc",
	}},
	{"Units of Measure", []string{
		"len",
		"temp",
	}},
	{"Library", []string{
		"angle",
		"anno",
		"bool",
		"color",
		"complex",
		"cmp",
		"crypto",
		"decimal",
		"format",
		"geo",
		"hof",
		"rand",
		"rational",
		"text",
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
