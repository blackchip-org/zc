package test

import (
	"bufio"
	"io/fs"
	"strings"
	"testing"

	"github.com/blackchip-org/zc/app"
	"github.com/blackchip-org/zc/doc"
)

func TestDoc(t *testing.T) {
	files, err := doc.Files.ReadDir("zlib")
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		t.Run(file.Name(), func(t *testing.T) {
			f, err := doc.Files.Open("zlib/" + file.Name())
			if err != nil {
				t.Fatal(err)
			}
			testFile(t, f)
		})
	}
}

func testFile(t *testing.T, file fs.File) {
	scanner := bufio.NewScanner(file)
	mod := ""
	fn := ""

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "## ") {
			fn = line[3:]
			continue
		}
		if strings.HasPrefix(line, "# ") {
			mod = line[2:]
			continue
		}
		fields := strings.Split(line, "|")
		if len(fields) == 3 &&
			strings.TrimSpace(fields[1]) == "Input" &&
			strings.TrimSpace(fields[2]) == "Stack" {
			scanner.Scan()
			scanner.Scan()
			t.Run(fn, func(t *testing.T) {
				testTable(t, mod, scanner)
			})
		}
	}
}

func testTable(t *testing.T, mod string, scanner *bufio.Scanner) {
	c := app.NewDefaultCalc()

	if err := c.Include(mod); err != nil {
		t.Fatal(err)
	}

	for {
		line := scanner.Text()
		if scanner.Err() != nil {
			t.Fatal(scanner.Err())
		}
		fields := strings.SplitN(line, "|", 3)
		if len(fields) != 3 {
			break
		}

		in := strings.ReplaceAll(fields[1], "`", "")
		in = strings.TrimSpace(in)

		out := strings.ReplaceAll(fields[2], "`", "")
		out = strings.ReplaceAll(out, "\\|", "|")
		out = strings.TrimSpace(out)

		if err := c.EvalString("", in); err != nil {
			t.Fatal(err)
		}
		if c.Stack.String() != out {
			t.Fatalf("\n have: %v \n want: %v", c.Stack.String(), out)
		}
		scanner.Scan()
	}
}
