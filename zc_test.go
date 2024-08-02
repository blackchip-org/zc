package zc

import "testing"

func TestItemString(t *testing.T) {
	tests := []struct {
		item Item
		str  string
	}{
		{Item{TypeVal: "123", Type: String}, "123"},
		{Item{TypeVal: "123\n456", Type: String}, "123\\n456"},
		{Item{TypeVal: "123", Unit: "m", Type: String}, "123m"},
		{Item{TypeVal: "123", Label: "label", Type: String}, "123 :label"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			str := test.item.String()
			if str != test.str {
				t.Fatalf("\n have: %v \n want: %v", str, test.str)
			}
		})
	}
}
