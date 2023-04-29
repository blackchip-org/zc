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
		"bool",
		"conv",
		"cmp",
		"format",
		"hof",
		"stack",
		"text",
		"zc",
	}},
	{"Library", []string{
		"angle",
		"color",
		"crypto",
		"geo",
		"rand",
		"time",
	}},
	{"Tables", []string{
		"epsg",
		"si",
		"tz",
	}},
}
