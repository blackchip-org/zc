package doc

import (
	"cmp"
	"slices"
	"unicode"
	"unicode/utf8"

	"golang.org/x/exp/maps"
)

type IndexEntry struct {
	Name string
	Ops  []IndexOp
}

func SortIndexEntries(ies []IndexEntry) {
	slices.SortStableFunc(ies, func(a IndexEntry, b IndexEntry) int {
		return cmp.Compare(a.Name, b.Name)
	})
}

func (e IndexEntry) Heading() rune {
	r, _ := utf8.DecodeRuneInString(e.Name)
	if !unicode.IsLetter(r) {
		return rune(0)
	}
	return unicode.ToUpper(r)
}

type IndexOp struct {
	VolName string
	Op      Op
}

func SortIndexOps(ops []IndexOp) {
	slices.SortStableFunc(ops, func(a IndexOp, b IndexOp) int {
		return cmp.Compare(a.VolName, b.VolName)
	})
}

func Index(vols []Vol) []IndexEntry {
	index := make(map[string]IndexEntry)
	for _, vol := range vols {
		for _, op := range vol.Ops {
			names := append([]string{op.Name}, op.Aliases...)
			for _, name := range names {
				e, ok := index[name]
				if !ok {
					e.Name = name
				}
				iop := IndexOp{vol.Name, op}
				e.Ops = append(e.Ops, iop)
				SortIndexOps(e.Ops)
				index[name] = e
			}
		}
	}
	entries := maps.Values(index)
	SortIndexEntries(entries)
	return entries
}
