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
		"si",
		"temp",
	}},
	{"Standard Operations", []string{
		"bool",
		"cmp",
		"format",
		"hof",
		"stack",
		"text",
		"types",
		"zc",
	}},
	{"Library", []string{
		"angle",
		"color",
		"crypto",
		"rand",
		"time",
	}},
	{"Tables", []string{
		"tz",
	}},
}
