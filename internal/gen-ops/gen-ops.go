//go:generate go run gen-ops.go

package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"unicode"

	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v6/calc"
)

var (
	OpsDir = path.Join("..", "..", "ops")
)

func main() {
	cat := calc.All
	var names []string
	for _, op := range cat.Ops() {
		names = append(names, op.Name)
	}
	sort.Strings(names)

	os.Mkdir(OpsDir, 0755)
	f, err := os.Create(path.Join(OpsDir, "ops.go"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "package ops\n\nvar (\n")
	for _, name := range names {
		var s scan.Scanner
		s.InitFromString("", name)

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
		fmt.Fprintf(f, "\t%v = \"%v\"\n", s.Emit().Val, name)
	}
	fmt.Fprintf(f, ")\n")
}
