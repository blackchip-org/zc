package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"path"
	"slices"
	"strings"

	"github.com/blackchip-org/zc/v5/pkg/doc"
	"github.com/blackchip-org/zc/v5/pkg/doc1"
)

//go:generate go run main.go

func main() {
	log.SetFlags(0)

	vols, err := doc.LoadDir("../../docsrc/ops")
	if err != nil {
		log.Fatal(err)
	}

	for _, vol := range vols {
		writeVol(vol)
	}
	writeCategories()
}

func writeVol(vol doc.Vol) {
	out := &strings.Builder{}

	fmt.Fprintf(out, "<!-- Document generated by \"gen-doc\"; DO NOT EDIT -->\n")
	if prelude, ok := doc.Preludes[vol.Name]; ok {
		fmt.Fprintf(out, "%v\n", prelude)
	}
	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "# %v\n\n", vol.Name)
	fmt.Fprintf(out, "%v\n\n", vol.Title)
	if vol.Overview != "" {
		fmt.Fprintf(out, "## Overview\n\n")
		fmt.Fprintf(out, vol.Overview)
		fmt.Fprintf(out, "\n")
	}
	fmt.Fprintf(out, "## Index\n\n")
	fmt.Fprintf(out, "| Operation | Description\n")
	fmt.Fprintf(out, "|-----------|------------\n")

	slices.SortFunc(vol.Ops, func(a doc.Op, b doc.Op) int {
		return cmp.Compare(a.Name, b.Name)
	})
	for _, op := range vol.Ops {
		allNames := op.Name
		if len(op.Aliases) > 0 {
			allNames += ", " + strings.Join(op.Aliases, ", ")
		}
		entry := fmt.Sprintf("[`%v`](#%v)", strings.Join(op.AllNames(), ", "), op.Name)
		fmt.Fprintf(out, "| %v | %v\n", entry, op.Title)
	}
	fmt.Fprintf(out, "\n## Operations\n")

	for _, op := range vol.Ops {
		writeOp(out, op)
	}

	file := path.Join("../../doc/ops", vol.Name+".md")
	if err := os.WriteFile(file, []byte(out.String()), 0o644); err != nil {
		log.Fatalf("unable to write file: %v", err)
	}
}

func writeOp(out *strings.Builder, op doc.Op) {
	fmt.Fprintf(out, "\n### %v\n\n", op.Name)
	fmt.Fprintf(out, "%v\n\n", op.Desc)
	if len(op.Aliases) > 0 {
		if len(op.Aliases) == 1 {
			fmt.Fprintf(out, "Alias: ")
		} else {
			fmt.Fprintf(out, "Aliases: ")
		}
		var fmtAliases []string
		for _, a := range op.Aliases {
			fmtAliases = append(fmtAliases, fmt.Sprintf("`%v`", a))
		}
		fmt.Fprintf(out, "%v\n\n", strings.Join(fmtAliases, ", "))
	}

	if len(op.Funcs) > 0 {
		fmt.Fprintf(out, "```\n")
		for _, fn := range op.Funcs {
			fmt.Fprintf(out, "( ")
			fmt.Fprintf(out, "%v -- ", fn.Params())
			fmt.Fprintf(out, "%v )\n", fn.Returns())
		}
		fmt.Fprintf(out, "```\n")
		fmt.Fprintln(out, "")
	}
	if op.Macro != "" {
		fmt.Fprintf(out, "```\n")
		fmt.Fprintf(out, "def %v %v\n", op.Name, op.Macro)
		fmt.Fprintf(out, "```\n")
		fmt.Fprintln(out, "")
	}

	if len(op.Example) > 0 {
		fmt.Fprintf(out, "Example:\n\n<!-- test: %v -->\n\n", op.Name)
		writeExample(out, op.Example)
	}
}

func writeExample(out *strings.Builder, expected []doc.Expect) {
	fmt.Fprintf(out, "| Input | Stack\n")
	fmt.Fprintf(out, "|-------|------\n")
	for _, ex := range expected {
		o := strings.Join(ex.O, " \\| ")
		quotes := "`"
		if strings.HasPrefix(o, "*") || o == "" {
			quotes = ""
		}
		fmt.Fprintf(out, "| `%v` | %v%v%v\n", ex.I, quotes, o, quotes)
	}
}

func writeCategories() {
	out := &strings.Builder{}
	fmt.Fprintf(out, "<!-- Document generated by \"gen-doc\"; DO NOT EDIT -->\n\n")
	fmt.Fprintf(out, "# Operation Reference\n\n")

	groupsSeen := make(map[string]bool)
	for id := range doc1.GroupTitles {
		groupsSeen[id] = false
	}

	width := 25

	for _, cat := range doc1.Categories {
		fmt.Fprintf(out, "## %v\n\n", cat.Name)
		fmt.Fprintf(out, "| %-[1]*v | Description\n", width, "Volume")
		fmt.Fprintf(out, "|%v|-----------\n", strings.Repeat("-", width+2))
		for _, group := range cat.Groups {
			title, ok := doc1.GroupTitles[group]
			if !ok {
				log.Fatalf("unknown group: %v", group)
			}
			seen := groupsSeen[group]
			if seen {
				log.Fatalf("duplicate group entry: %v", group)
			}
			groupsSeen[group] = true
			link := fmt.Sprintf("[%v](ops/%v.md)", group, group)
			fmt.Fprintf(out, "| %-[1]*v | %v \n", width, link, title)
		}
		fmt.Fprintf(out, "\n\n")
	}

	for group, seen := range groupsSeen {
		if !seen {
			log.Fatalf("group not seen: %v", group)
		}
	}

	if err := os.WriteFile("../../doc/ops.md", []byte(out.String()), 0o644); err != nil {
		log.Fatal(err)
	}
}
