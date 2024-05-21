package zc

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
