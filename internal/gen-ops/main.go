package main

//go:generate go run main.go

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/blackchip-org/zc/pkg/doc"
)

var prelude = strings.TrimSpace(`
// Code generated by "gen-ops"; DO NOT EDIT
package calc

import (
	"github.com/blackchip-org/zc/pkg/ops"
	"github.com/blackchip-org/zc/pkg/zc"
)

var opsList = []zc.OpDecl{
`)

var epilog = strings.TrimSpace(`
}
`)

func main() {
	log.SetFlags(0)

	ops, err := doc.ParseSourceFiles("../../pkg/ops")
	if err != nil {
		log.Fatal(err)
	}

	table := doc.Table(ops)
	names := doc.SortedNames(table)

	out := &strings.Builder{}
	fmt.Fprintf(out, "%v\n", prelude)
	for _, name := range names {
		op := table[name]
		if name != op.Name {
			fmt.Fprintf(out, "\tzc.Macro(\"%v\", \"%v\"),\n", name, op.Name)
			continue
		}
		if op.Macro != "" {
			fmt.Fprintf(out, "\tzc.Macro(\"%v\", \"%v\"),\n", name, op.Macro)
			continue
		}

		if len(op.Funcs) > 1 {
			fmt.Fprintf(out, "\tzc.GenOp(\"%v\",\n", name)
			for _, fn := range op.Funcs {
				if fn.Name == "-" {
					continue
				}
				fmt.Fprintf(out, "\t\tzc.Func(%v%v),\n", qnameFn(fn.Name), paramTypes(fn.Params))
			}
			fmt.Fprintf(out, "\t),\n")
		} else {
			fn := op.Funcs[0]
			fmt.Fprintf(out, "\tzc.Op(\"%v\", %v%v),\n", name, qnameFn(fn.Name), paramTypes(fn.Params))
		}
	}
	out.WriteString(epilog)
	if err := os.WriteFile("../../pkg/calc/ops_gen.go", []byte(out.String()), 0644); err != nil {
		log.Fatal(err)
	}
}

func paramTypes(params []doc.Param) string {
	if len(params) == 0 {
		return ""
	}
	var inTypes []string
	for _, p := range params {
		if p.All {
			break
		}
		inTypes = append(inTypes, "zc."+p.Type)
	}
	if len(inTypes) == 0 {
		return ""
	}
	return fmt.Sprintf(", %v", strings.Join(inTypes, ", "))
}

func qnameFn(name string) string {
	if name == "NoOp" {
		return "zc.NoOp"
	}
	return "ops." + name
}
