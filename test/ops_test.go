package test

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/blackchip-org/zc/pkg/ansi"
	"github.com/blackchip-org/zc/pkg/calc"
	"github.com/blackchip-org/zc/pkg/repl"
	"github.com/blackchip-org/zc/pkg/scanner"
	"github.com/blackchip-org/zc/pkg/zc"
)

func TestOps(t *testing.T) {
	ansi.Enabled = false
	testDir(t, ".")
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

	var c zc.Calc
	var r *repl.REPL
	var out strings.Builder

	var reset = func() {
		c = calc.New()
		r = repl.New(c)
		r.Out = &out
		out.Reset()
	}

	s := scanner.New(f)
	space4 := "    "
	space8 := space4 + space4

	for s.Ok() {
		space := s.ScanWhile(scanner.Rune(' '))
		switch space {
		case space4:
			if c == nil {
				t.Log("reset")
				reset()
			}
			expr := s.ScanUntil(scanner.IsNewline)
			s.Next()
			t.Logf("-> %v", expr)
			r.Eval(expr)
			if c.Error() != nil {
				c.Push(c.Error().Error())
			}
		case space8:
			rem := s.ScanUntil(scanner.IsNewline)
			s.Next()
			t.Logf("<- %v", zc.StackString(c))
			wants := strings.Split(rem, "|")
			for i := len(wants) - 1; i >= 0; i-- {
				want := strings.TrimSpace(wants[i])
				have, ok := c.Pop()
				if !ok {
					t.Fatalf("\n have: empty stack \n want: %v", zc.Quote(want))
				}
				if have != want {
					t.Fatalf("\n have: %v \n want: %v\n", zc.Quote(have), zc.Quote(want))
				}
			}
		default:
			c = nil
			s.ScanUntil(scanner.IsNewline)
			s.Next()
		}
	}
}
