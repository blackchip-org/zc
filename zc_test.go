package zc

import "testing"

func TestIsValue(t *testing.T) {
	tests := []struct {
		word    string
		isValue bool
	}{
		{"", false},
		{"1", true},
		{"11", true},
		{"$1", true},
		{"a", false},
		{"abc", false},
		{"/a", true},
		{"/abc", true},
	}

	for _, test := range tests {
		t.Run(test.word, func(t *testing.T) {
			isValue := IsValue(test.word)
			if isValue != test.isValue {
				t.Errorf("\n have: %v \n want: %v", isValue, test.isValue)
			}
		})
	}
}
