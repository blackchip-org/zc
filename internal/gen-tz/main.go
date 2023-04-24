package main

//go:generate go run main.go

import (
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strings"
	"unicode"
)

const (
	TzGo = "../../pkg/ops/tz.go"
	TzMd = "../../doc/ops/tz.md"
)

func main() {
	log.SetFlags(0)

	zoneMap := make(map[string]string)
	processDir(zoneMap, []string{}, "/usr/share/zoneinfo")

	var names []string
	for name := range zoneMap {
		names = append(names, name)
	}
	sort.Strings(names)

	ftz, err := os.Create(TzGo)
	if err != nil {
		log.Panic(err)
	}
	defer ftz.Close()

	ftz.WriteString("package ops\n\n")
	ftz.WriteString("var TimeZones = map[string]string {")
	for _, name := range names {
		fmt.Fprintf(ftz, "\n\t\"tz.%v\": \"[%v]\",", name, zoneMap[name])
	}
	ftz.WriteString("\n}")

	fdoc, err := os.Create(TzMd)
	if err != nil {
		log.Panic(err)
	}
	defer fdoc.Close()

	fmt.Fprintf(fdoc, `
# tz

Time zone database.

| Operation | Description
|-----------|------------
`)

	for _, name := range names {
		fmt.Fprintf(fdoc, "| `tz.%v` | %v\n", name, zoneMap[name])
	}
}

func processDir(zones map[string]string, parent []string, dir string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Panic(err)
	}
	for _, entry := range entries {
		here := append(parent, entry.Name())
		if entry.IsDir() {
			childDir := path.Join(dir, entry.Name())
			processDir(zones, here, childDir)
			continue
		}
		first := entry.Name()[0]
		if first == '+' || unicode.IsLower(rune(first)) {
			continue
		}
		zone := strings.Join(here, "/")

		var parts []string
		for i := len(here) - 1; i >= 0; i-- {
			part := strings.ToLower(here[i])
			part = strings.ReplaceAll(part, "_", "-")
			parts = append(parts, part)
		}
		word := strings.Join(parts, ".")
		zones[word] = zone
	}
}