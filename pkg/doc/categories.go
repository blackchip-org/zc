package doc

import "slices"

type Category struct {
	Name string
	Vols []string
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

func FilterByCategory(src []*Vol, members []string) []*Vol {
	var target []*Vol
	for _, o := range src {
		if slices.Contains(members, o.Name) {
			target = append(target, o)
		}
	}
	return target
}
