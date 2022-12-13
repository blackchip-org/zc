package zc

import "testing"

func TestFormatNumberString(t *testing.T) {
	optDefault := DefaultNumberFormatOptions()
	optFR := NumberFormatOptions{
		IntPat: ".000",
		Point:  ',',
	}
	optBin := NumberFormatOptions{
		IntPat: "__0000_0000",
	}
	optEmpty := NumberFormatOptions{}

	tests := []struct {
		in   string
		want string
		opts NumberFormatOptions
	}{
		{"1", "1", optDefault},
		{"123", "123", optDefault},
		{"1234", "1,234", optDefault},
		{"123456", "123,456", optDefault},
		{"1234567", "1,234,567", optDefault},
		{"-123", "-123", optDefault},
		{".123", ".123", optDefault},
		{".12345", ".12345", optDefault},
		{"1234567.8901", "1,234,567.8901", optDefault},
		{"1234567.8901", "1234567.8901", optEmpty},
		{"1234567.8901", "1.234.567,8901", optFR},
		{"11110000", "1111_0000", optBin},
		{"1111000011110000", "1111_0000__1111_0000", optBin},
	}

	c := NewCalc()
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			have := c.FormatNumberString(test.in, test.opts)
			if have != test.want {
				t.Errorf("\n have: %v \n want: %v", have, test.want)
			}
		})
	}
}
