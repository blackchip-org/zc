package doc

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
	"unicode"

	"github.com/blackchip-org/zc/v5/pkg/scanner"
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
	Macro   string
}

func (o Op) AllNames() []string {
	return append([]string{o.Name}, o.Aliases...)
}

type OpByGroup []*Op

func (b OpByGroup) Len() int           { return len(b) }
func (b OpByGroup) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b OpByGroup) Less(i, j int) bool { return b[i].Group < b[j].Group }

func ParseSourceFile(name string) ([]*Op, error) {
	var ops []*Op
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %v", err)
	}
	defer f.Close()

	group := strings.TrimSuffix(path.Base(name), ".go")
	s := scanner.New(f)
	s.SetName(name)
	//s.Debug = true
	for s.Ok() {
		if s.Ch == '/' && s.Lookahead == '*' {
			s.Scan(scanner.Line)
			op, err := parseOp(s, group)
			if err != nil {
				return nil, err
			}
			ops = append(ops, op)
		} else if s.Ch == '/' && s.Lookahead == '/' {
			s.Next()
			s.Next()
			word := s.Scan(scanner.Word)
			if word == "tab" {
				op, err := parseTableOp(s, group)
				if err != nil {
					return nil, err
				}
				ops = append(ops, op)
			} else {
				s.Scan(scanner.Line)
			}
		} else {
			s.Scan(scanner.Line)
		}
	}
	return ops, s.Error
}

func ParseSourceFiles(dir string) ([]*Op, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var ops []*Op
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

func parseOp(s *scanner.Scanner, group string) (*Op, error) {
	op := &Op{Group: group}
loop:
	for s.Ok() {
		var err error
		word := s.Scan(scanner.Word)
		switch word {
		case "oper":
			op.Name = s.Scan(scanner.LineTrimSpace)
		case "func":
			var fn FuncDecl
			fn, err = parseFn(s)
			op.Funcs = append(op.Funcs, fn)
		case "macro":
			op.Macro = s.Scan(scanner.LineTrimSpace)
		case "alias":
			op.Aliases = append(op.Aliases, s.Scan(scanner.LineTrimSpace))
		case "title":
			op.Title = s.Scan(scanner.LineTrimSpace)
		case "desc":
			op.Desc = parseDesc(s)
		case "example":
			op.Example, err = parseExample(s)
		case "*/":
			break loop
		}
		if err != nil {
			return op, err
		}
	}
	if op.Name == "" {
		return op, scanErr(s, "no name")
	}
	if op.Title == "" {
		return op, scanErr(s, "no title for %v", op.Name)
	}
	return op, nil
}

func parseFn(s *scanner.Scanner) (FuncDecl, error) {
	s.ScanWhile(unicode.IsSpace)
	fn := FuncDecl{
		Name: s.Scan(scanner.Word),
	}
	var err error
	fn.Params, err = parseParams(s)
	if err != nil {
		return fn, err
	}
	fn.Returns, err = parseParams(s)
	if err != nil {
		return fn, err
	}
	return fn, nil
}

func parseParams(s *scanner.Scanner) ([]Param, error) {
	var params []Param
	for s.Ok() {
		s.ScanWhile(scanner.Rune2(' ', '\t'))
		if s.Ch == '\n' {
			s.Next()
			return params, nil
		}
		if s.Ch == '-' && s.Lookahead == '-' {
			s.Next()
			s.Next()
			return params, nil
		}
		var all bool
		var name, pType string
		t := s.ScanWhile(scanner.Or(
			unicode.IsLetter,
			unicode.IsDigit,
			scanner.Rune('.'),
		))
		if t == "" {
			return nil, scanErr(s, "did not find a parameter name or type")
		}
		if t == "..." {
			continue
		}
		if s.Ch == ':' {
			name = t
			s.Next()
			t = s.ScanWhile(scanner.Or(
				unicode.IsLetter,
				unicode.IsDigit,
				scanner.Rune('.'),
			))
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
	return params, nil
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

func parseExample(s *scanner.Scanner) ([]Expect, error) {
	var example []Expect
	s.Scan(scanner.Line)
	for s.Ok() {
		line := s.Scan(scanner.LineTrimSpace)
		if line == "end" {
			break
		}
		if line == "*/" {
			return nil, scanErr(s, "unexpected end of comment")
		}
		parts := strings.Split(line, "--")
		if len(parts) != 2 {
			return nil, scanErr(s, "invalid example line: %v", line)
		}
		in := strings.TrimSpace(parts[0])
		out := strings.TrimSpace(parts[1])
		example = append(example, Expect{In: in, Out: out})
	}
	return example, nil
}

func parseTableOp(s *scanner.Scanner, group string) (*Op, error) {
	line := s.Scan(scanner.LineTrimSpace)
	parts := strings.Split(line, "--")
	if len(parts) != 3 {
		return nil, scanErr(s, "invalid table line: %v", line)
	}
	return &Op{
		Group: group,
		Name:  strings.TrimSpace(parts[0]),
		Macro: strings.TrimSpace(parts[1]),
		Title: strings.TrimSpace(parts[2]),
		Desc:  strings.TrimSpace(parts[2]),
	}, nil
}

func FilterByGroup(src []*Op, group string) []*Op {
	var target []*Op
	for _, o := range src {
		if o.Group == group {
			target = append(target, o)
		}
	}
	return target
}

func SortedNames(table map[string]*Op) []string {
	var names []string
	for name := range table {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func Table(ops []*Op) map[string]*Op {
	table := make(map[string]*Op)
	for _, op := range ops {
		other, ok := table[op.Name]
		if ok {
			other.Aliases = append(other.Aliases, op.Aliases...)
			other.Funcs = append(other.Funcs, op.Funcs...)
		} else {
			table[op.Name] = op
		}
		for _, a := range op.Aliases {
			if _, exists := table[a]; !exists {
				table[a] = op
			}
		}
	}
	return table
}

func Group(ops []*Op) map[string][]*Op {
	table := make(map[string][]*Op)
	for _, op := range ops {
		names := append([]string{op.Name}, op.Aliases...)
		for _, name := range names {
			if other, ok := table[name]; ok {
				other = append(other, op)
				sort.Sort(OpByGroup(other))
				table[name] = other
			} else {
				table[name] = []*Op{op}
			}
		}
	}
	return table
}

func scanErr(s *scanner.Scanner, format string, a ...any) error {
	msg := fmt.Sprintf(format, a...)
	return fmt.Errorf("[%v] %v", s.TokenPos, msg)
}
