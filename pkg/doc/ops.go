package doc

import (
	"os"
	"path"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	"gopkg.in/yaml.v3"
)

type Vol struct {
	Name     string
	Title    string
	Overview string
	Ops      []Op
}

type Op struct {
	Prefix  string   `yaml:"prefix"`
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

func QName(prefix string, name string) string {
	if prefix != "" {
		return prefix + "." + name
	}
	return name
}

type opByName []Op

func (b opByName) Len() int      { return len(b) }
func (b opByName) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

// Anything that isn't a letter is sorted before anything else
func (b opByName) Less(i, j int) bool {
	ni, nj := b[i].Name, b[j].Name
	ci, _ := utf8.DecodeRuneInString(ni)
	cj, _ := utf8.DecodeRuneInString(nj)
	switch {
	case unicode.IsLetter(ci) && !unicode.IsLetter(cj):
		return false
	case !unicode.IsLetter(ci) && unicode.IsLetter(cj):
		return true
	default:
		return ni < nj
	}
}

func SortOps(ops []Op) {
	sort.Sort(opByName(ops))
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

func LoadDir(dir string) ([]Vol, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var vols []Vol
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
		vols = append(vols, vol)
	}
	return vols, err
}
