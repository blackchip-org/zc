package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/blackchip-org/zc/v5/pkg/doc"
)

//go:generate go run main.go

func main() {
	log.SetFlags(0)

	ops, err := doc.ParseSourceFiles("../../pkg/ops")
	if err != nil {
		log.Fatal(err)
	}

	groups := make(map[string]struct{})
	for _, op := range ops {
		groups[op.Group] = struct{}{}
	}

	for group := range groups {
		writeGroup(group, doc.FilterByGroup(ops, group))
	}
	writeCategories()
}

func writeGroup(group string, ops []*doc.Op) {
	out := &strings.Builder{}
	table := doc.Table(ops)
	names := doc.SortedNames(table)

	fmt.Fprintf(out, "<!-- Document generated by \"gen-doc\"; DO NOT EDIT -->")
	if prelude, ok := doc.Preludes[group]; ok {
		fmt.Fprintf(out, "%v\n", prelude)
	}
	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "# %v\n\n", group)
	if groupTitle, ok := doc.GroupTitles[group]; ok {
		fmt.Fprintf(out, "%v\n\n", groupTitle)
	} else {
		log.Fatalf("no group title for %v\n", group)
	}

	exampleFile := path.Join("../../doc/examples", group+".md")
	if _, err := os.Stat(exampleFile); err == nil {
		fmt.Fprintf(out, "[Examples](../examples/%v.md)\n\n", group)
	}

	width := 0
	for _, name := range names {
		op := table[name]
		contents := strings.Join(op.AllNames(), ", ")
		w := len(contents) + len(name)
		if w > width {
			width = w
		}
	}

	fmt.Fprintf(out, "| %-[1]*v | Description\n", width, "Operation")
	fmt.Fprintf(out, "|%v|%v\n", strings.Repeat("-", width+2), strings.Repeat("-", 15))
	for _, name := range names {
		op := table[name]
		if op.Name != name {
			continue
		}
		entry := fmt.Sprintf("[`%v`](#%v)", strings.Join(op.AllNames(), ", "), name)
		fmt.Fprintf(out, "| %-[1]*v | %v\n", width, entry, table[name].Title)
	}
	fmt.Fprintln(out, "")

	for _, name := range names {
		op := table[name]
		if name == op.Name {
			writeOp(out, op)
		}
	}

	file := path.Join("../../doc/ops", group+".md")
	if err := os.WriteFile(file, []byte(out.String()), 0o644); err != nil {
		log.Fatalf("unable to write file: %v", err)
	}
}

func writeOp(out *strings.Builder, op *doc.Op) {
	fmt.Fprintf(out, "\n## %v\n\n", op.Name)
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
			var fmtParams []string
			for _, p := range fn.Params {
				fmtParams = append(fmtParams, p.String())
			}
			fmt.Fprintf(out, "%v -- ", strings.Join(fmtParams, " "))
			fmtParams = nil
			for _, p := range fn.Returns {
				fmtParams = append(fmtParams, p.String())
			}
			fmt.Fprintf(out, "%v )\n", strings.Join(fmtParams, " "))
		}
		fmt.Fprintf(out, "```\n")
		fmt.Fprintln(out, "")
	}

	if len(op.Example) > 0 {
		fmt.Fprintf(out, "Example:\n\n<!-- test: %v -->\n\n", op.Name)
		writeExample(out, op.Example)
	}
}

func writeExample(out *strings.Builder, expected []doc.Expect) {
	width := 0
	for _, ex := range expected {
		w := len(ex.In)
		if w > width {
			width = w
		}
	}
	if width < 3 {
		width = 3
	}

	fmt.Fprintf(out, "| %-[1]*v | Stack\n", width+2, "Input")
	fmt.Fprintf(out, "|%v|%v\n", strings.Repeat("-", width+4), strings.Repeat("-", 15))
	for _, ex := range expected {
		escOut := strings.ReplaceAll(ex.Out, "|", "\\|")
		quotes := "`"
		if strings.HasPrefix(escOut, "*") || escOut == "" {
			quotes = ""
		}
		fmt.Fprintf(out, "| `%-[1]*v` | %v%v%v\n", width, ex.In, quotes, escOut, quotes)
	}
}

func writeCategories() {
	out := &strings.Builder{}
	fmt.Fprintf(out, "<!-- Document generated by \"gen-doc\"; DO NOT EDIT -->\n\n")
	fmt.Fprintf(out, "# Operation Reference\n\n")

	groupsSeen := make(map[string]bool)
	for id := range doc.GroupTitles {
		groupsSeen[id] = false
	}

	width := 25

	for _, cat := range doc.Categories {
		fmt.Fprintf(out, "## %v\n\n", cat.Name)
		fmt.Fprintf(out, "| %-[1]*v | Description\n", width, "Volume")
		fmt.Fprintf(out, "|%v|-----------\n", strings.Repeat("-", width+2))
		for _, group := range cat.Groups {
			title, ok := doc.GroupTitles[group]
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
