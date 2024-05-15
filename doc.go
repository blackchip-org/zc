package zc

type Expect struct {
	Input  string   `yaml:"i"`
	Output []string `yaml:"o"`
	Error  string   `yaml:"error"`
	Info   string   `yaml:"info"`
}

type OpDoc struct {
	Name      string     `yaml:"name"`
	Ident     string     `yaml:"ident"`
	Overloads string     `yaml:"overloads"`
	Aliases   []string   `yaml:"aliases"`
	Title     string     `yaml:"title"`
	Funcs     []*FuncDoc `yaml:"funcs"`
	Macro     string     `yaml:"macro"`
	Desc      string     `yaml:"desc"`
	Example   []Expect   `yaml:"example"`
	Tests     []TestDoc  `yaml:"tests"`
}

type FuncDoc struct {
	Name    string   `yaml:"name"`
	Ident   string   `yaml:"ident"`
	Params  []string `yaml:"params"`
	Returns []string `yaml:"returns"`
	Prec    string   `yaml:"prec"`
}

type VolDoc struct {
	Name     string `yaml:"name"`
	Ident    string `yaml:"ident"`
	Title    string `yaml:"title"`
	Package  string `yaml:"package"`
	Overview string
	Kinds    []string `yaml:"kinds"`
	Ops      []*OpDoc `yaml:"ops"`
}

type TestDoc struct {
	Name string   `yaml:"name"`
	Test []Expect `yaml:"test"`
}
