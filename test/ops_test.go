package test

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v5/pkg/ansi"
	"github.com/blackchip-org/zc/v5/pkg/calc"
	"github.com/blackchip-org/zc/v5/pkg/repl"
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

	s := scan.NewScanner(file, f)

	var blockName string
	for s.HasMore() {
		if s.This == '#' && s.Next == '#' {
			s.Discard()
			s.Discard()
			blockName = strings.TrimSpace(scan.Line(s))
			s.Discard()
		} else {
			scan.While(s, scan.Rune(' '), s.Keep)
			space := s.Emit().Val
			if space == space4 {
				t.Run(blockName, func(t *testing.T) {
					testBlock(t, s)
				})
			} else {
				scan.Until(s, scan.Rune('\n'), s.Discard)
				s.Discard()
			}
		}
	}
}

func testBlock(t *testing.T, s *scan.Scanner) {
	c := calc.New()
	r := repl.New(c)
	out := &strings.Builder{}
	r.Out = out

	for s.HasMore() {
		scan.Until(s, scan.Rune('\n'), s.Keep)
		test := strings.TrimSpace(s.Emit().Val)
		s.Discard()

		if !repl.Test(r, test) {
			t.Errorf("\n output: %v (error %v) \n  input: %v\n", zc.StackString(c), c.Error(), test)
		}
		scan.While(s, scan.Rune(' '), s.Keep)
		space := s.Emit().Val
		if space != space4 {
			break
		}
	}
}
