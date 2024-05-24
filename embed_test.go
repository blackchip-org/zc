package zc

import "testing"

func TestEmbed(t *testing.T) {
	_, err := LoadDefs()
	if err != nil {
		t.Fatal(err)
	}
}
