package scanner

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Reader struct {
	src *bufio.Reader
	ch  []rune
	pos int
	err error
}

func NewReader(src io.Reader, n int) *Reader {
	r := &Reader{
		src: bufio.NewReader(src),
		ch:  make([]rune, n+1),
	}
	for i := 0; i < n+1; i++ {
		r.Read()
	}
	return r
}

func NewStringReader(src string, n int) *Reader {
	return NewReader(strings.NewReader(src), n)
}

func (r *Reader) Read() rune {
	if r.ch[r.pos] == EndCh {
		return EndCh
	}

	ch := r.ch[r.pos]
	la, _, err := r.src.ReadRune()
	if err != nil {
		r.err = err
		la = EndCh
	}
	r.ch[r.pos] = la
	r.pos++
	if r.pos >= len(r.ch) {
		r.pos = 0
	}
	return ch
}

func (r *Reader) Peek(i int) rune {
	if i > len(r.ch) {
		panic(fmt.Sprintf("peek of %v exceeds lookahead of %v", i, len(r.ch)-1))
	}
	offset := r.pos + i
	if offset >= len(r.ch) {
		offset = offset - len(r.ch)
	}
	return r.ch[offset]
}

func (r *Reader) PeekN(n int) string {
	var s []rune
	for i := 1; i <= n; i++ {
		ch := r.Peek(i)
		if ch == EndCh {
			break
		}
		s = append(s, ch)
	}
	return string(s)
}
