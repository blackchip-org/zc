package main

//go:generate go run main.go ../../../doc/index.md

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	Root = "../../.."
)

var (
	modDirective   = regexp.MustCompile(`^# *(.+) *$`)
	indexDirective = regexp.MustCompile(`<!-- *index *-->`)
	//mainDirective  = regexp.MustCompile(`<!-- *main: ([^ ]+) -->`)
	tableHeader   = regexp.MustCompile(`.*Operation.*Description`)
	operationName = regexp.MustCompile(`\[(.*)\]\((.*)\)`)
	aliasFormat   = regexp.MustCompile("\\`[^\\`]+\\`")
)

type Entry struct {
	Func        string
	Module      string
	Description string
	File        string
	Anchor      string
}

var EntryMap map[string]Entry

func init() {
	EntryMap = make(map[string]Entry)
}

func main() {
	outFile := os.Args[1]

	var files []string
	dirs := []string{
		"doc/ops",
	}

	for _, dir := range dirs {
		entries, err := os.ReadDir(path.Join(Root, dir))
		if err != nil {
			log.Fatal(err)
		}
		for _, entry := range entries {
			files = append(files, path.Join(Root, dir, entry.Name()))
		}
	}

	for _, name := range files {
		file, err := os.Open(name)
		if err != nil {
			log.Fatal(err)
		}
		indexFile(name, file)
	}

	var keys []string
	for key := range EntryMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	out, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	fmt.Fprint(out, "# index\n\n")

	thisHeading := rune(0)
	for _, key := range keys {
		heading, _ := utf8.DecodeRuneInString(key)
		if unicode.IsLetter(heading) && heading != thisHeading {
			fmt.Fprintf(out, "\n## %v\n\n", string(heading))
			thisHeading = heading
		}
		entry := EntryMap[key]
		fmt.Fprintf(out, "- [%v](%v%v): %v\n", entry.Func, entry.File[4:], entry.Anchor, entry.Description)

	}
}

func indexFile(name string, file fs.File) {
	scanner := bufio.NewScanner(file)
	mod := ""

	for scanner.Scan() {
		if scanner.Err() != nil {
			log.Fatal(scanner.Err())
		}
		line := scanner.Text()

		matches := modDirective.FindStringSubmatch(line)
		if matches != nil {
			mod = matches[1]
			continue
		}

		if !indexDirective.MatchString(line) {
			continue
		}

		for scanner.Scan() {
			line = scanner.Text()
			if strings.TrimSpace(line) != "" || scanner.Err() != nil {
				break
			}
		}
		line = scanner.Text()
		if !tableHeader.MatchString(line) {
			log.Fatalf("expected table header but got: %v", line)
		}
		scanner.Scan()
		indexTable(name, mod, scanner)
	}
}

func indexTable(file string, mod string, scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "|")
		if len(fields) < 3 {
			break
		}
		var name, anchor, aliases, description string
		matches := operationName.FindStringSubmatch(fields[1])
		if matches == nil {
			log.Printf("unexpected operation format: %v", fields[1])
			continue
		}
		name = matches[1]
		anchor = matches[2]
		if len(fields) == 3 {
			description = strings.TrimSpace(fields[2])
		} else {
			aliases = fields[2]
			description = strings.TrimSpace(fields[3])
		}

		names := []string{name}
		matches = aliasFormat.FindAllString(aliases, -1)
		for _, match := range matches {
			names = append(names, match[1:len(match)-1])
		}

		for _, name := range names {
			entry := Entry{
				Func:        name,
				Anchor:      anchor,
				Description: description,
				File:        file,
				Module:      mod,
			}
			addEntry(entry)
		}
	}
}

func addEntry(e Entry) {
	_, exists := EntryMap[e.Func]
	if exists {
		panic(fmt.Sprintf("entry already exists: %v", e.Func))
	}
	EntryMap[e.Func] = e
}
