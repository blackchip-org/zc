package test

import (
	"bufio"
	"io/fs"
	"os"
	"path"
	"regexp"
	"strings"
	"testing"

	"github.com/blackchip-org/zc/pkg/calc"
)

var (
	testBanner  = regexp.MustCompile(`<!-- test: (.*) *-->`)
	tableHeader = regexp.MustCompile(`.*Input.*Stack`)
)

func TestDoc(t *testing.T) {
	files := []string{
		//"../README.md",
	}
	dirs := []string{
		"../doc/ops",
	}

	for _, dir := range dirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			t.Fatal(err)
		}
		for _, entry := range entries {
			files = append(files, path.Join(dir, entry.Name()))
		}
	}

	for _, name := range files {
		t.Run(name, func(t *testing.T) {
			file, err := os.Open(name)
			if err != nil {
				t.Fatal(err)
			}
			defer file.Close()
			testDocFile(t, file)
		})
	}
}

func testDocFile(t *testing.T, file fs.File) {
	scanner := bufio.NewScanner(file)
	testName := ""

	for scanner.Scan() {
		if scanner.Err() != nil {
			t.Fatal(scanner.Err())
		}
		line := scanner.Text()

		matches := testBanner.FindStringSubmatch(line)
		if matches == nil {
			continue
		}
		testName = strings.TrimSpace(matches[1])

		t.Run(testName, func(t *testing.T) {
			for scanner.Scan() {
				line = scanner.Text()
				if strings.TrimSpace(line) != "" || scanner.Err() != nil {
					break
				}
			}
			line = scanner.Text()
			if !tableHeader.MatchString(line) {
				t.Fatalf("expected table header but got: %v", line)
			}
			scanner.Scan()
			scanner.Scan()
			testTable(t, scanner)
		})

	}
}

func testTable(t *testing.T, scanner *bufio.Scanner) {
	c := calc.New()

	for {
		line := scanner.Text()
		fields := strings.SplitN(line, "|", 3)
		if len(fields) != 3 {
			break
		}

		in := strings.ReplaceAll(fields[1], "`", "")
		in = strings.TrimSpace(in)

		out := ""
		info := ""

		f2 := strings.TrimSpace(fields[2])
		if strings.HasPrefix(f2, "*") {
			info = f2
		} else {
			out = strings.ReplaceAll(f2, "`", "")
			out = strings.ReplaceAll(out, "\\|", "|")
			out = strings.TrimSpace(out)
		}

		t.Log(in)

		if err := c.Eval(in); err != nil {
			t.Fatal(err)
		}
		if c.Info() != "" || info != "" {
			have := "*" + c.Info() + "*"
			if have != info {
				t.Fatalf("\n have: %v \n want: %v", have, info)
			}
		} else {
			actualOut := strings.Join(c.Stack(), " | ")
			if actualOut != out {
				t.Fatalf("\n have: %v \n want: %v", actualOut, out)
			}
		}
		scanner.Scan()
		if scanner.Err() != nil {
			t.Fatal(scanner.Err())
		}
	}
}
