package pretty

import (
	"regexp"
	"strings"
	"testing"
)

func TestMarkdownTable(t *testing.T) {
	tab := NewMarkdownTable(3)
	tab.Heading("Alpha", "Bravo", "Charlie")
	tab.Row("One", "Two", "Three")
	tab.Row("Four", "Five", "Six")
	tab.Row("Seven", "Eight", "Nine")

	have := strings.TrimSpace(tab.Format())
	want := strings.TrimSpace(`
| Alpha | Bravo | Charlie
|-------|-------|--------
| One   | Two   | Three
| Four  | Five  | Six
| Seven | Eight | Nine
`)
	spaces := regexp.MustCompile(` +\n`)
	have = spaces.ReplaceAllString(have, "\n")

	if have != want {
		t.Errorf("\n have:\n%v\n want: \n%v\n", have, want)
	}
}
