package main

//go:generate go run main.go ../../doc/index.md

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/blackchip-org/zc/v5/pkg/doc"
)

func main() {
	log.SetFlags(0)

	vols, err := doc.LoadDir("../../pkg/ops")
	if err != nil {
		log.Fatal(err)
	}
	index := doc.Index(vols)

	out := &strings.Builder{}
	fmt.Fprintf(out, "<!-- Document generated by \"gen-index\"; DO NOT EDIT -->\n\n")
	fmt.Fprint(out, "# index\n\n")

	thisHeading := rune(0)
	for _, entry := range index {
		heading := entry.Heading()
		if heading != 0 && heading != thisHeading {
			fmt.Fprintf(out, "\n## %v\n\n", string(heading))
			thisHeading = heading
		}
		name := entry.Name
		if name == "+" || name == "-" {
			name = "\\" + name
		}
		if len(entry.Ops) == 1 {
			opIdx := entry.Ops[0]
			op := opIdx.Op
			fmt.Fprintf(out, "- [%v](ops/%v.md#%v): %v\n", name, opIdx.VolName, op.Name, op.Title)
		} else {
			for _, opIdx := range entry.Ops {
				op := opIdx.Op
				fmt.Fprintf(out, "- [%v](ops/%v.md#%v) (%v): %v\n", name, opIdx.VolName, op.Name, opIdx.VolName, op.Title)
			}
		}
	}

	if err := os.WriteFile("../../doc/index.md", []byte(out.String()), 0o644); err != nil {
		log.Fatal(err)
	}
}
