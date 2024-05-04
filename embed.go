package zc

import (
	"cmp"
	"embed"
	"fmt"
	"log"
	"path"
	"slices"
	"strings"

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
		fmt.Println(filename)
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

		slices.SortFunc(doc.Ops, func(a OpDoc, b OpDoc) int {
			return cmp.Compare(a.Name, b.Name)
		})
		return doc, err
	}
	return doc, fmt.Errorf("no yaml in %v", dir)
}
