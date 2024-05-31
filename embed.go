package zc

import (
	"cmp"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"path"
	"slices"
	"strings"
	"unicode"

	"github.com/blackchip-org/scan"
	"gopkg.in/yaml.v3"
)

const DefsDir = "defs"

//go:embed defs/*
var Defs embed.FS

func LoadDefs() ([]VolDef, error) {
	files, err := Defs.ReadDir(DefsDir)
	if err != nil {
		return nil, err
	}

	var defs []VolDef
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".yaml") {
			continue
		}
		def, err := loadDef(DefsDir, f)
		if err != nil {
			return nil, err
		}
		defs = append(defs, def)
	}
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
	if err != nil {
		return def, err
	}
	def.Overview = string(data)

	if def.Name == "" {
		log.Panicf("no volume name in %v", filename)
	}
	if def.Package == "" {
		def.Package = identFor(def.Name)
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
		if op.Overloads == "" {
			parts := strings.SplitN(op.Name, "/", 2)
			if len(parts) > 1 {
				op.Overloads = parts[0]
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
	return def, err
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

	prefix := ""
	for s.HasMore() {
		switch {
		case (s.This == '.' || s.This == '/') && unicode.IsLetter(s.Next):
			s.Skip()
			s.Val.WriteRune(unicode.ToUpper(s.This))
			s.Skip()
		case s.This == '?' && s.Next == scan.EndOfText:
			s.Skip()
			prefix = "Is"
		default:
			s.Keep()
		}
	}
	return prefix + s.Emit().Val
}
