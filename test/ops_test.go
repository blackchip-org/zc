package test

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/blackchip-org/zc/v5/pkg/ansi"
	"github.com/blackchip-org/zc/v5/pkg/calc"
	"github.com/blackchip-org/zc/v5/pkg/repl"
	"github.com/blackchip-org/zc/v5/pkg/scanner"
	"github.com/blackchip-org/zc/v5/pkg/zc"
)

const space4 = "    "

func TestOps(t *testing.T) {
	ansi.Enabled = false
	testDir(t, "./ops")
}

func testDir(t *testing.T, dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".md") {
			continue
		}
		t.Run(file.Name(), func(t *testing.T) {
			name := path.Join(dir, file.Name())
			if file.IsDir() {
				testDir(t, name)
			} else {
				testFile(t, name)
			}
		})
	}
}

func testFile(t *testing.T, file string) {
	f, err := os.Open(file)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	s := scanner.New(f)

	var blockName string
	for s.Ok() {
		if s.Ch == '#' && s.Lookahead == '#' {
			s.Next()
			s.Next()
			blockName = s.Scan(scanner.LineTrimSpace)
		} else {
			space := s.ScanWhile(scanner.Rune(' '))
			if space == space4 {
				t.Run(blockName, func(t *testing.T) {
					testBlock(t, s)
				})
			} else {
				s.Scan(scanner.Line)
			}
		}
	}
}

func testBlock(t *testing.T, s *scanner.Scanner) {
	c := calc.New()
	r := repl.New(c)
	out := &strings.Builder{}
	r.Out = out

	for s.Ok() {
		test := s.Scan(scanner.LineTrimSpace)
		if !repl.Test(r, test) {
			t.Errorf("\n output: %v (error %v) \n  input: %v\n", zc.StackString(c), c.Error(), test)
		}
		space := s.ScanWhile(scanner.Rune(' '))
		if space != space4 {
			break
		}
	}
}
