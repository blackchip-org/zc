package scanner

import (
	"strconv"
	"testing"
)

func TestReaderN(t *testing.T) {
	for i := 0; i <= 3; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			testReaderN(t, i)
		})
	}
}

func testReaderN(t *testing.T, n int) {
	out := make([]rune, 3)
	want := "abc"
	r := NewStringReader(want, n)
	out[0] = r.Read()
	out[1] = r.Read()
	out[2] = r.Read()
	have := string(out)
	if have != want {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}
	if r.Read() != EndCh {
		t.Errorf("expected end")
	}
}

func TestLA(t *testing.T) {
	r := NewStringReader("1234", 3)
	want := "23"
	have := r.PeekN(2)
	if have != want {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}

	r.Read()
	want = "34"
	have = r.PeekN(2)
	if have != want {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}

	r.Read()
	want = "4"
	have = r.PeekN(2)
	if have != want {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}

	r.Read()
	want = ""
	have = r.PeekN(2)
	if have != want {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}
}
