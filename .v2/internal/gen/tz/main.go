package main

//go:generate go run main.go ../../zlib/tz.zc  ../../../doc/zlib/tz.md

import (
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strings"
	"unicode"
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

	ftz, err := os.Create(os.Args[1])
	if err != nil {
		log.Panic(err)
	}
	defer ftz.Close()

	for _, name := range names {
		fmt.Fprintf(ftz, "macro %v '%v'\n", name, zoneMap[name])
	}

	fdoc, err := os.Create(os.Args[2])
	if err != nil {
		log.Panic(err)
	}
	defer fdoc.Close()

	fmt.Fprintf(fdoc, `
# tz

Time zone database.

- Use: import

| Operation | Description
|-----------|------------
`)

	for _, name := range names {
		fmt.Fprintf(fdoc, "| `%v` | %v\n", name, zoneMap[name])
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
