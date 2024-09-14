package zc

import (
	"cmp"
	"embed"
	"fmt"
	"io/fs"
	"path"
	"slices"
	"strings"
	"unicode"

	"github.com/blackchip-org/scan"
	"gopkg.in/yaml.v3"
)

const DefsDir = "app/defs"

//go:embed app/defs/*
var Defs embed.FS

func LoadDefs() ([]VolDef, error) {
	var defs []VolDef
	fs.WalkDir(Defs, DefsDir, func(p string, d fs.DirEntry, e error) error {
		if !strings.HasSuffix(p, ".yaml") {
			return nil
		}
		def, err := loadDef(path.Dir(p), d)
		if err != nil {
			panic(err)
		}
		defs = append(defs, def)
		return nil
	})
	return defs, nil
}

func loadDef(dir string, f fs.DirEntry) (VolDef, error) {
	var def VolDef

	filename := path.Join(dir, f.Name())
	data, err := Defs.ReadFile(filename)
	if err != nil {
		return def, err
	}

	err = yaml.Unmarshal(data, &def)
	if err != nil {
		return def, fmt.Errorf("%v: %v", filename, err)
	}

	overview := strings.TrimSuffix(f.Name(), ".yaml") + ".md"
	overviewFile := path.Join(dir, overview)
	data, err = Defs.ReadFile(overviewFile)
	if err == nil {
		def.Overview = string(data)
	}

	if def.Name == "" {
		panic("no volume name in file: " + filename)
	}

	if def.Ident == "" {
		def.Ident = identFor(def.Name)
		if def.Ident == "" {
			panic("no identifier for: " + def.Name)
		}
	}
	for i, op := range def.Ops {
		if op.Ident == "" {
			op.Ident = identFor(op.Name)
			if op.Ident == "" {
				panic("no identifier for: " + op.Name)
			}
			def.Ops[i] = op
		}
		for i, fn := range op.Funcs {
			if fn.Name == "" && fn.Ident == "" {
				fn.Ident = identFor(op.Name)
			} else if fn.Ident == "" {
				fn.Ident = identFor(fn.Name)
			}
			op.Funcs[i] = fn
		}
	}

	slices.SortFunc(def.Ops, func(a OpDef, b OpDef) int {
		return cmp.Compare(a.Name, b.Name)
	})
	return def, nil
}

func identFor(v string) string {
	if v == "" {
		return ""
	}

	s := scan.NewScannerFromString("", v)
	if !scan.IsLetter(s.This) {
		return ""
	}
	s.Val.WriteRune(unicode.ToUpper(s.This))
	s.Skip()

	for s.HasMore() {
		switch {
		case (s.This == '.' || s.This == '/' || s.This == '-'):
			s.Skip()
			s.Val.WriteRune(unicode.ToUpper(s.This))
			s.Skip()
		case s.This == '?':
			s.Skip()
			s.Val.WriteString("Q")
		case s.This == '=':
			s.Skip()
			s.Val.WriteString("E")
		default:
			s.Keep()
		}
	}
	return s.Emit().Val
}
