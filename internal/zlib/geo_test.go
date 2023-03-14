package zlib

import (
	"testing"

	"github.com/blackchip-org/zc"
)

func TestParseCoordinate(t *testing.T) {
	tests := []struct {
		text  string
		coord float64
	}{
		{"123° 27' 24.4404\"", 123.456789},
		{"123° 27' 24.4404\"N", 123.456789},
		{"123° 27' 24.4404\"S", -123.456789},
		{"123° 27' 24.4404\" S", -123.456789},
		{"123°27'24.4404\"", 123.456789},
		{"123d27'24.4404\"", 123.456789},
		{"123 27'24.4404\"", 123.456789},
		{"123", 123},
		{"123.45", 123.45},
		{"-123.45", -123.45},
		{"123d 45'", 123.75},
		{"123d 45.678'", 123.7613},
	}

	calc, err := zc.NewCalc(zc.Config{Point: '.'})
	if err != nil {
		panic(err)
	}
	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {
			coord, ok := parseDD(calc, test.text)
			if !ok {
				t.Fatalf("not a coordinate: %v", test.text)
			}
			if coord != test.coord {
				t.Errorf(" have: %v \n want: %v", coord, test.coord)
			}
		})
	}
}

func TestParseCoordinateBad(t *testing.T) {
	tests := []string{
		"123d -45'",
		"123x",
		"123d 45",
		"123d 45' 56",
		"123.45d 67.89'",
		"123d 45.67' 89\"",
		"123° 27' 24.4404\"X",
	}

	calc, err := zc.NewCalc(zc.Config{Point: '.'})
	if err != nil {
		panic(err)
	}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			_, ok := parseDD(calc, test)
			if ok {
				t.Errorf("expected error: %v", test)
			}
		})
	}
}
