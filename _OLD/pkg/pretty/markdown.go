package pretty

import (
	"fmt"
	"strings"
)

type MarkdownTable struct {
	cols int
	rows [][]string
}

func NewMarkdownTable(cols int) *MarkdownTable {
	return &MarkdownTable{cols: cols}
}

func (t *MarkdownTable) Heading(h ...string) {
	t.checkCols(h)
	t.rows = append([][]string{h}, t.rows...)
}

func (t *MarkdownTable) Row(c ...string) {
	t.checkCols(c)
	t.rows = append(t.rows, c)
}

func (t *MarkdownTable) Format() string {
	var b strings.Builder

	widths := make([]int, t.cols)
	for r := 0; r < len(t.rows); r++ {
		for c := 0; c < len(t.rows[r]); c++ {
			widths[c] = Max(widths[c], len(t.rows[r][c]))
		}
	}
	fmt.Fprint(&b, "| ")
	for i := 0; i < t.cols; i++ {
		fmt.Fprintf(&b, "%-*v", widths[i], t.rows[0][i])
		if i < t.cols-1 {
			fmt.Fprint(&b, " | ")
		}
	}
	fmt.Fprintln(&b)
	fmt.Fprint(&b, "|-")
	for i := 0; i < t.cols; i++ {
		fmt.Fprint(&b, RepeatString(widths[i], "-"))
		if i < t.cols-1 {
			fmt.Fprint(&b, "-|-")
		}
	}
	fmt.Fprintln(&b)
	for _, r := range t.rows[1:] {
		fmt.Fprint(&b, "| ")
		for i := 0; i < t.cols; i++ {
			fmt.Fprintf(&b, "%-*v", widths[i], r[i])
			if i < t.cols-1 {
				fmt.Fprint(&b, " | ")
			}
		}
		fmt.Fprintln(&b)
	}
	return b.String()
}

func (t *MarkdownTable) checkCols(v []string) {
	if len(v) != t.cols {
		panic(fmt.Sprintf("expected %v columns but got %v", t.cols, len(v)))
	}
}
