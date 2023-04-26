package doc

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/blackchip-org/zc/pkg/scanner"
)

type Param struct {
	Name string
	Type string
	All  bool
}

func (p Param) String() string {
	var out strings.Builder
	if p.Name != "" {
		out.WriteString(p.Name)
		out.WriteRune(':')
	}
	out.WriteString(p.Type)
	if p.All {
		out.WriteRune('*')
	}
	return out.String()
}

type FuncDecl struct {
	Name    string
	Params  []Param
	Returns []Param
}

type Expect struct {
	In  string
	Out string
}

type Op struct {
	Group   string
	Name    string
	Funcs   []FuncDecl
	Aliases []string
	Title   string
	Desc    string
	Example []Expect
}

func ParseSourceFile(name string) ([]Op, error) {
	var ops []Op
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %v", err)
	}
	defer f.Close()

	group := strings.TrimSuffix(path.Base(name), ".go")
	s := scanner.New(f)
	//s.Debug = true
	for s.Ok() {
		if s.Ch == '/' && s.Lookahead == '*' {
			s.Scan(scanner.Line)
			ops = append(ops, parseOp(s, group))
		} else {
			s.Scan(scanner.Line)
		}
	}
	return ops, s.Error
}

func ParseSourceFiles(dir string) ([]Op, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var ops []Op
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".go") {
			name := path.Join(dir, f.Name())
			theseOps, err := ParseSourceFile(name)
			if err != nil {
				return nil, err
			}
			ops = append(ops, theseOps...)
		}
	}
	return ops, nil
}

func parseOp(s *scanner.Scanner, group string) Op {
	op := Op{Group: group}
loop:
	for s.Ok() {
		word := s.Scan(scanner.Word)
		switch word {
		case "oper":
			op.Name = s.Scan(scanner.LineTrimSpace)
		case "func":
			op.Funcs = append(op.Funcs, parseFn(s))
		case "alias":
			op.Aliases = append(op.Aliases, s.Scan(scanner.LineTrimSpace))
		case "title":
			op.Title = s.Scan(scanner.LineTrimSpace)
		case "desc":
			op.Desc = parseDesc(s)
		case "example":
			op.Example = parseExample(s)
		case "*/":
			break loop
		}
	}
	return op
}

func parseFn(s *scanner.Scanner) FuncDecl {
	fn := FuncDecl{
		Name: s.Scan(scanner.Word),
	}
	fn.Params = parseParams(s)
	fn.Returns = parseParams(s)
	return fn
}

func parseParams(s *scanner.Scanner) []Param {
	var params []Param
	for s.Ok() {
		s.ScanWhile(scanner.Rune(' '))
		if s.Ch == '\n' {
			s.Next()
			return params
		}
		if s.Ch == '-' && s.Lookahead == '-' {
			s.Next()
			s.Next()
			return params
		}
		var all bool
		var name, pType string
		t := s.ScanWhile(scanner.Or(
			scanner.IsCharAZ,
			scanner.IsDigit09,
		))
		if t == "" {
			panic("no progress")
		}
		if s.Ch == ':' {
			name = t
			s.Next()
			t = s.ScanWhile(scanner.IsCharAZ)
		}
		pType = t
		if s.Ch == '*' {
			s.Next()
			all = true
		}
		params = append(params, Param{
			Name: name,
			Type: pType,
			All:  all,
		})
	}
	return params
}

func parseDesc(s *scanner.Scanner) string {
	var desc []string
	s.Scan(scanner.Line)
	for s.Ok() {
		line := s.Scan(scanner.LineTrimSpace)
		if line == "end" {
			break
		}
		desc = append(desc, line)
	}
	return strings.Join(desc, "\n")
}

func parseExample(s *scanner.Scanner) []Expect {
	var example []Expect
	s.Scan(scanner.Line)
	for s.Ok() {
		line := s.Scan(scanner.LineTrimSpace)
		if line == "end" {
			break
		}
		if line == "*/" {
			panic("unexpected end of comment")
		}
		parts := strings.Split(line, "--")
		in := strings.TrimSpace(parts[0])
		out := strings.TrimSpace(parts[1])
		example = append(example, Expect{In: in, Out: out})
	}
	return example
}

func FilterByGroup(src []Op, group string) []Op {
	var target []Op
	for _, o := range src {
		if o.Group == group {
			target = append(target, o)
		}
	}
	return target
}

func ByName(src []Op) map[string]Op {
	target := make(map[string]Op)
	for _, o := range src {
		if _, exists := target[o.Name]; exists {
			panic(fmt.Sprintf("duplicate op: %v", o.Name))
		}
		target[o.Name] = o
	}
	return target
}

func SortedNames(src []Op) []string {
	var names []string
	for _, o := range src {
		names = append(names, o.Name)
	}
	sort.Strings(names)
	return names
}
