package zc

type OpDef struct {
	Name     string    `yaml:"name"`
	Ident    string    `yaml:"ident"`
	Aliases  []string  `yaml:"aliases"`
	Title    string    `yaml:"title"`
	Subtitle string    `yaml:"subtitle"`
	Stub     bool      `yaml:"stub"`
	Funcs    []FuncDef `yaml:"funcs"`
	Macro    string    `yaml:"macro"`
	Desc     string    `yaml:"desc"`
	Example  []Expect  `yaml:"example"`
	Tests    []Test    `yaml:"tests"`
}

type FuncDef struct {
	Name    string   `yaml:"name"`
	Ident   string   `yaml:"ident"`
	Params  []string `yaml:"params"`
	Returns []string `yaml:"returns"`
}

type VolDef struct {
	Name        string     `yaml:"name"`
	Ident       string     `yaml:"ident"`
	Title       string     `yaml:"title"`
	Subtitle    string     `yaml:"subtitle"`
	Category    string     `yaml:"category"`
	NoIndex     bool       `yaml:"no-index"`
	RemoveSlash bool       `yaml:"remove-slash"`
	Setup       []string   `yaml:"setup"`
	Ops         []OpDef    `yaml:"ops"`
	Table       [][]string `yaml:"table"`
	Overview    string
}

type Category struct {
	ID    string
	Title string
}

type Expect struct {
	Input  string   `yaml:"i"`
	Output []string `yaml:"o"`
	Error  string   `yaml:"error"`
	Notice string   `yaml:"notice"`
}

type Test struct {
	Name string   `yaml:"name"`
	Test []Expect `yaml:"test"`
}

var Categories = []Category{
	{ID: "calc", Title: "Calculator Operations"},
	{ID: "unit", Title: "Units of Measure"},
	{ID: "lib", Title: "Library"},
	{ID: "tab", Title: "Tables"},
}
