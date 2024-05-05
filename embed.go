package zc

import (
	"cmp"
	"embed"
	"fmt"
	"log"
	"path"
	"slices"
	"strings"
	"unicode"

	"github.com/blackchip-org/scan"
	"gopkg.in/yaml.v3"
)

//go:embed volumes/*/*.yaml volumes/*/*.md
var Docs embed.FS

func LoadDocs() ([]VolDoc, error) {
	files, err := Docs.ReadDir("volumes")
	if err != nil {
		return nil, err
	}

	var docs []VolDoc
	for _, f := range files {
		if f.IsDir() {
			doc, err := loadDoc(path.Join("volumes", f.Name()))
			if err != nil {
				return nil, err
			}
			docs = append(docs, doc)
		}
	}
	return docs, nil
}

func loadDoc(dir string) (VolDoc, error) {
	var doc VolDoc
	files, err := Docs.ReadDir(dir)
	if err != nil {
		return doc, err
	}

	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".yaml") {
			continue
		}
		filename := path.Join(dir, f.Name())
		data, err := Docs.ReadFile(filename)
		if err != nil {
			return doc, err
		}

		err = yaml.Unmarshal(data, &doc)
		if err != nil {
			return doc, fmt.Errorf("%v: %v", filename, err)
		}

		overview := strings.TrimSuffix(f.Name(), ".yaml") + ".md"
		overviewFile := path.Join(dir, overview)
		data, err = Docs.ReadFile(overviewFile)
		if err != nil {
			return doc, err
		}
		doc.Overview = string(data)

		if doc.Name == "" {
			log.Panicf("no volume name in %v", filename)
		}
		if doc.Package == "" {
			doc.Package = doc.Name
		}
		if doc.Ident == "" {
			doc.Ident = identFor(doc.Name)
			if doc.Ident == "" {
				panic("no identifier for: " + doc.Name)
			}
		}
		for _, op := range doc.Ops {
			if op.Ident == "" {
				op.Ident = identFor(op.Name)
				if op.Ident == "" {
					panic("no identifier for: " + op.Name)
				}
			}
		}

		slices.SortFunc(doc.Ops, func(a *OpDoc, b *OpDoc) int {
			return cmp.Compare(a.Name, b.Name)
		})
		return doc, err
	}
	return doc, fmt.Errorf("no yaml in %v", dir)
}

func opIdent(name string) (string, bool) {
	var s scan.Scanner
	s.InitFromString("", name)

	// Skip any names that are only symbols
	if !scan.IsLetter(s.This) {
		return "", false
	}
	s.Val.WriteRune(unicode.ToUpper(s.This))
	s.Skip()

	for s.HasMore() {
		if s.Next == '.' || s.Next == '-' {
			s.Keep()
			s.Skip()
			s.Val.WriteRune(unicode.ToUpper(s.This))
			s.Skip()
		} else {
			s.Keep()
		}
	}
	return s.Emit().Val, true
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
		if s.This == '.' && unicode.IsLetter(s.Next) {
			s.Skip()
			s.Val.WriteRune(unicode.ToUpper(s.This))
			s.Skip()
		} else {
			s.Keep()
		}
	}
	return s.Emit().Val
}
