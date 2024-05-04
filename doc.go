package zc

type Expect struct {
	Input  string   `yaml:"i"`
	Output []string `yaml:"o"`
}

type OpDoc struct {
	Name    string   `yaml:"name"`
	Title   string   `yaml:"title"`
	Aliases []string `yaml:"aliases"`
	Func    string   `yaml:"func"`
	Macro   string   `yaml:"macro"`
	Params  []string `yaml:"params"`
	Returns []string `yaml:"returns"`
	Prec    string   `yaml:"prec"`
	Desc    string   `yaml:"desc"`
	Example []Expect `yaml:"example"`
}

type VolDoc struct {
	Name     string `yaml:"name"`
	Title    string `yaml:"title"`
	Package  string `yaml:"package"`
	Overview string
	Kinds    []string `yaml:"kinds"`
	Ops      []OpDoc  `yaml:"ops"`
}
