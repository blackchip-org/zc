package doc

import (
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

type Vol struct {
	Name     string
	Title    string
	Overview string
	Ops      []Op
}

type Op struct {
	Name    string   `yaml:"name"`
	Title   string   `yaml:"title"`
	Funcs   []Func   `yaml:"funcs,omitempty"`
	Macro   string   `yaml:"macro,omitempty"`
	Aliases []string `yaml:"aliases,omitempty"`
	Desc    string   `yaml:"desc"`
	Example []Expect `yaml:"example"`
}

func (o Op) AllNames() []string {
	return append([]string{o.Name}, o.Aliases...)
}

type Func struct {
	Name string   `yaml:"name"`
	I    []string `yaml:"i"`
	O    []string `yaml:"o"`
	Var  bool     `yaml:"var"`
	Id   []string `yaml:"id"`
}

func (f Func) Params() string {
	var p []string
	for idx, in := range f.I {
		prefix := ""
		if idx < len(f.Id) {
			prefix = f.Id[idx] + ":"
		}
		p = append(p, prefix+in)
	}
	return strings.Join(p, " ")
}

func (f Func) Returns() string {
	var r []string
	for idx, out := range f.O {
		prefix := ""
		if idx < len(f.Id) {
			prefix = f.Id[idx] + ":"
		}
		r = append(r, prefix+out)
	}
	return strings.Join(r, " ")
}

type Expect struct {
	I string   `yaml:"i"`
	O []string `yaml:"o"`
}

func LoadDir(dir string) ([]*Vol, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var vols []*Vol
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".yaml") {
			continue
		}
		data, err := os.ReadFile(path.Join(dir, f.Name()))
		if err != nil {
			return vols, err
		}
		var vol Vol
		err = yaml.Unmarshal(data, &vol)
		if err != nil {
			return vols, err
		}
		vols = append(vols, &vol)
	}
	return vols, err
}
