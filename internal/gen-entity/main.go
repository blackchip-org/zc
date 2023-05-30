package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

//go:generate go run main.go

var ignore = map[string]struct{}{
	"&ApplyFunction;":         {},
	"&InvisibleComma;":        {},
	"&InvisibleTimes;":        {},
	"&NewLine;":               {},
	"&MediumSpace;":           {},
	"&NegativeMediumSpace;":   {},
	"&NegativeThickSpace;":    {},
	"&NegativeThinSpace;":     {},
	"&NegativeVeryThinSpace;": {},
	"&NoBreak;":               {},
	"&NonBreakingSpace;":      {},
	"&Tab;":                   {},
	"&ThickSpace;":            {},
	"&ThinSpace;":             {},
	"&VerticalLine;":          {},
	"&VeryThinSpace;":         {},
	"&ZeroWidthSpace;":        {},
	"&af;":                    {},
	"&emsp13;":                {},
	"&emsp14;":                {},
	"&emsp;":                  {},
	"&ensp;":                  {},
	"&hairsp;":                {},
	"&ic;":                    {},
	"&it;":                    {},
	"&lrm;":                   {},
	"&nbsp;":                  {},
	"&numsp;":                 {},
	"&puncsp;":                {},
	"&rlm;":                   {},
	"&shy;":                   {},
	"&thinsp;":                {},
	"&verbar;":                {},
	"&vert;":                  {},
	"&zwj;":                   {},
	"&zwnj;":                  {},
}

const (
	EntityGo = "../../pkg/ops/entity.go"
	EntityMd = "../../doc/ops/entity.md"
)

type Entity struct {
	Codepoints []int  `json:"codepoints"`
	Characters string `json:"characters"`
}

func main() {
	var names []string
	entities := make(map[string]Entity)
	data, err := os.ReadFile("entities.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &entities)

	for name := range entities {
		if _, ok := ignore[name]; ok {
			continue
		}
		if !strings.HasSuffix(name, ";") {
			continue
		}
		names = append(names, name)
	}
	sort.Strings(names)

	fent, err := os.Create(EntityGo)
	if err != nil {
		panic(err)
	}
	defer fent.Close()

	fmt.Fprintf(fent, "package ops\n\n")
	fmt.Fprintf(fent, "var Entities = map[string]string {")
	for _, name := range names {
		entity := entities[name]
		if entity.Characters == "\"" {
			entity.Characters = "\\\"s"
		}
		if entity.Characters == "\\" {
			entity.Characters = "\\\\"
		}
		fmt.Fprintf(fent, "\n\t\"%v\": \"[%v]\",", name, entity.Characters)
	}
	fmt.Fprintf(fent, "\n}\n")

	fdoc, err := os.Create(EntityMd)
	if err != nil {
		log.Panic(err)
	}
	defer fdoc.Close()

	fmt.Fprintf(fdoc, `
# entity

HTML entities.

| Operation | Description
|-----------|------------
`)
	for _, name := range names {
		fmt.Fprintf(fdoc, "| `%v` | %v\n", name, entities[name].Characters)
	}

}
