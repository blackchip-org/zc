package test

import (
	"bufio"
	"io/fs"
	"path"
	"regexp"
	"strings"
	"testing"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/app"
)

var (
	modDirective = regexp.MustCompile(`<!-- mod: *([\w-\.]+) *-->`)
	testBanner   = regexp.MustCompile(`<!-- test: *(\w+) *-->`)
	tableHeader  = regexp.MustCompile(`.*Input.*Stack`)
)

func TestDoc(t *testing.T) {
	files := []string{
		"README.md",
	}
	dirs := []string{
		"doc/zlib",
	}

	for _, dir := range dirs {
		entries, err := zc.Files.ReadDir(dir)
		if err != nil {
			t.Fatal(err)
		}
		for _, entry := range entries {
			files = append(files, path.Join(dir, entry.Name()))
		}
	}

	for _, name := range files {
		t.Run(name, func(t *testing.T) {
			file, err := zc.Files.Open(name)
			if err != nil {
				t.Fatal(err)
			}
			testFile(t, file)
		})
	}
}

func testFile(t *testing.T, file fs.File) {
	scanner := bufio.NewScanner(file)
	testName := ""
	mod := ""

	for scanner.Scan() {
		if scanner.Err() != nil {
			t.Fatal(scanner.Err())
		}
		line := scanner.Text()

		matches := modDirective.FindStringSubmatch(line)
		if matches != nil {
			mod = matches[1]
			continue
		}

		matches = testBanner.FindStringSubmatch(line)
		if matches == nil {
			continue
		}
		testName = matches[1]

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
			testTable(t, mod, scanner)
		})

	}
}

func testTable(t *testing.T, mod string, scanner *bufio.Scanner) {
	c := app.NewDefaultCalc()

	if mod != "" {
		if err := c.EvalString("", "use "+mod); err != nil {
			t.Fatal(err)
		}
	}

	for {
		line := scanner.Text()
		fields := strings.SplitN(line, "|", 3)
		if len(fields) != 3 {
			break
		}

		in := strings.ReplaceAll(fields[1], "`", "")
		in = strings.TrimSpace(in)

		out := strings.ReplaceAll(fields[2], "`", "")
		out = strings.ReplaceAll(out, "\\|", "|")
		out = strings.TrimSpace(out)

		t.Log(in)

		if err := c.EvalString("", in); err != nil {
			t.Fatal(err)
		}
		if c.Env.Stack.String() != out {
			t.Fatalf("\n have: %v \n want: %v", c.Env.Stack.String(), out)
		}
		scanner.Scan()
		if scanner.Err() != nil {
			t.Fatal(scanner.Err())
		}
	}
}
