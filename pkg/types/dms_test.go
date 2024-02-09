package types

import (
	"fmt"
	"testing"
)

func TestDMS(t *testing.T) {
	tests := []struct {
		inDeg, inMin, inSec    any
		outDeg, outMin, outSec string
	}{
		{1, 0, 0, "1", "0", "0"},
		{1.5, 0, 0, "1", "30", "0"},
		{1.5, 1.5, 0, "1", "31", "30"},
		{1.5, 1.5, 1.5, "1", "31", "31.5"},
		{0, 0, 3600, "1", "0", "0"},
		{0, 0, 3665, "1", "1", "5"},
		{0, 65, 0, "1", "5", "0"},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v:%v:%v", test.inDeg, test.inMin, test.inSec)
		t.Run(name, func(t *testing.T) {
			dms, err := NewDMS(test.inDeg, test.inMin, test.inSec)
			if err != nil {
				t.Fatal(err)
			}
			deg, min, sec := dms.DMS()
			if deg.String() != test.outDeg || min.String() != test.outMin || sec.String() != test.outSec {
				t.Errorf("\n have: %v %v %v \n want: %v %v %v", deg, min, sec, test.outDeg, test.outMin, test.outSec)
			}
		})
	}
}

func TestDegrees(t *testing.T) {
	tests := []struct {
		d, m, s string
		deg     string
	}{
		{"1", "0", "0", "1"},
		{"1", "3", "0", "1.05"},
		{"1", "3", "9", "1.0525"},
		{"0", "3", "0", "0.05"},
		{"0", "3", "9", "0.0525"},
		{"0", "0", "9", "0.0025"},
		{"-1", "3", "9", "-1.0525"},
	}

	for _, test := range tests {
		t.Run(test.deg, func(t *testing.T) {
			dms, err := NewDMS(test.d, test.m, test.s)
			if err != nil {
				t.Fatal(err)
			}
			deg := dms.Degrees().String()
			if deg != test.deg {
				t.Errorf("\n have: %v \n want: %v", deg, test.deg)
			}
		})
	}
}
