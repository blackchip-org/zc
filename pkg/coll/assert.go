package coll

import (
	"fmt"
	"testing"
)

func assertPanic(t *testing.T, expected string) {
	t.Helper()
	if have := recover(); have != nil {
		have := fmt.Sprintf("%v", have)
		want := expected
		if have != want {
			t.Fatalf("\n have panic: %v \n want panic: %v", have, want)
		}
	} else {
		t.Fatal("expected panic")
	}
}
