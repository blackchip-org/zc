package zc

type OpDef struct {
	Name      string    `yaml:"name"`
	Ident     string    `yaml:"ident"`
	Overloads string    `yaml:"overloads"`
	Virtual   bool      `yaml:"virtual"`
	Aliases   []string  `yaml:"aliases"`
	Title     string    `yaml:"title"`
	Funcs     []FuncDef `yaml:"funcs"`
	Macro     string    `yaml:"macro"`
	Desc      string    `yaml:"desc"`
	Example   []Expect  `yaml:"example"`
	Tests     []Test    `yaml:"tests"`
}

type FuncDef struct {
	Name    string   `yaml:"name"`
	Ident   string   `yaml:"ident"`
	Params  []string `yaml:"params"`
	Returns []string `yaml:"returns"`
}

type VolDef struct {
	Name     string   `yaml:"name"`
	Ident    string   `yaml:"ident"`
	Title    string   `yaml:"title"`
	Subtitle string   `yaml:"subtitle"`
	Category string   `yaml:"category"`
	Types    []string `yaml:"types"`
	Ops      []OpDef  `yaml:"ops"`
	Overview string
}

type Category struct {
	ID    string
	Order int
	Title string
}

type Expect struct {
	Input  string   `yaml:"i"`
	Output []string `yaml:"o"`
	Error  string   `yaml:"error"`
	Info   string   `yaml:"info"`
}

type Test struct {
	Name string   `yaml:"name"`
	Test []Expect `yaml:"test"`
}

var Categories = map[string]Category{
	"calc": {Order: 1, ID: "calc", Title: "Calculator Operations"},
	"unit": {Order: 2, ID: "unit", Title: "Units of Measure"},
	"lib":  {Order: 3, ID: "lib", Title: "Library"},
	"tab":  {Order: 4, ID: "tab", Title: "Tables"},
}
