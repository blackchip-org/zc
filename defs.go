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
	Types    []string `yaml:"types"`
	Ops      []OpDef  `yaml:"ops"`
	Overview string
}
