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
	evalDirective = regexp.MustCompile(`<!-- eval: (.*) -->`)
	testBanner    = regexp.MustCompile(`<!-- test: (.*) *-->`)
	tableHeader   = regexp.MustCompile(`.*Input.*Stack`)
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
	var setup []string

	for scanner.Scan() {
		if scanner.Err() != nil {
			t.Fatal(scanner.Err())
		}
		line := scanner.Text()

		matches := evalDirective.FindStringSubmatch(line)
		if matches != nil {
			setup = append(setup, matches[1])
			continue
		}

		matches = testBanner.FindStringSubmatch(line)
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
			testTable(t, setup, scanner)
		})

	}
}

func testTable(t *testing.T, setup []string, scanner *bufio.Scanner) {
	c := app.NewDefaultCalc()

	for _, l := range setup {
		if err := c.EvalString("", l); err != nil {
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

		if err := c.EvalString("", in); err != nil {
			t.Fatal(err)
		}
		if c.Info != "" {
			have := "*" + c.Info + "*"
			if have != info {
				t.Fatalf("\n have: %v \n want: %v", have, info)
			}
		} else {
			if c.Env.Stack.String() != out {
				t.Fatalf("\n have: %v \n want: %v", c.Env.Stack.String(), out)
			}
		}
		scanner.Scan()
		if scanner.Err() != nil {
			t.Fatal(scanner.Err())
		}
	}
}
