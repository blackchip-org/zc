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
		"mass",
		"temp",
	}},
	{"Library", []string{
		"angle",
		"anno",
		"bin",
		"bool",
		"color",
		"complex",
		"cmp",
		"crypto",
		"format",
		"geo",
		"hof",
		"rand",
		"rational",
		"seq",
		"text",
		"time",
	}},
	{"Tables", []string{
		"entity",
		"emoji",
		"epsg",
		"iec",
		"si",
		"tz",
	}},
}
