package test

import (
	"bufio"
	"os"
	"path"
	"reflect"
	"strings"
	"testing"

	"github.com/blackchip-org/zc/calc"
)

func TestOps(t *testing.T) {
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
		t.Run(file.Name(), func(*testing.T) {
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

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "" || strings.HasPrefix(input, "#") {
			continue
		}
		var output []string
		for scanner.Scan() {
			o := strings.TrimSpace(scanner.Text())
			if o == "" {
				break
			}
			output = append(output, o)
		}

		t.Run(input, func(t *testing.T) {
			doTest(t, input, output)
		})
	}
}

func doTest(t *testing.T, input string, want []string) {
	c := calc.New()
	err := c.Eval(input)
	if err != nil {
		if len(want) == 0 || err.Error() != want[0] {
			t.Fatalf("unexpected error: %v", err)
		}
	} else {
		have := c.Stack()
		if !reflect.DeepEqual(have, want) {
			t.Fatalf("\n have: %v \n want: %v", have, want)
		}
	}
}
