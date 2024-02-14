package types

import (
	"testing"

	"github.com/blackchip-org/dms"
)

func TestDMS(t *testing.T) {
	tests := []struct {
		fields                 dms.Fields
		outDeg, outMin, outSec string
	}{
		{dms.Fields{Deg: "1", Min: "0", Sec: "0"}, "1", "0", "0"},
		{dms.Fields{Deg: "1.5", Min: "0", Sec: "0"}, "1", "30", "0"},
		{dms.Fields{Deg: "1.5", Min: "1.5", Sec: "0"}, "1", "31", "30"},
		{dms.Fields{Deg: "1.5", Min: "1.5", Sec: "1.5"}, "1", "31", "31.5"},
		{dms.Fields{Deg: "0", Min: "0", Sec: "3600"}, "1", "0", "0"},
		{dms.Fields{Deg: "0", Min: "0", Sec: "3665"}, "1", "1", "5"},
		{dms.Fields{Deg: "0", Min: "65", Sec: "0"}, "1", "5", "0"},
	}

	for _, test := range tests {
		t.Run(test.fields.String(), func(t *testing.T) {
			dms, err := NewDMS(test.fields)
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
		fields dms.Fields
		deg    string
	}{
		{dms.Fields{Deg: "1", Min: "0", Sec: "0"}, "1"},
		{dms.Fields{Deg: "1", Min: "3", Sec: "0"}, "1.05"},
		{dms.Fields{Deg: "1", Min: "3", Sec: "9"}, "1.0525"},
		{dms.Fields{Deg: "0", Min: "3", Sec: "0"}, "0.05"},
		{dms.Fields{Deg: "0", Min: "3", Sec: "9"}, "0.0525"},
		{dms.Fields{Deg: "0", Min: "0", Sec: "9"}, "0.0025"},
		{dms.Fields{Hemi: "-", Deg: "1", Min: "3", Sec: "9"}, "-1.0525"},
	}

	for _, test := range tests {
		t.Run(test.fields.String(), func(t *testing.T) {
			dms, err := NewDMS(test.fields)
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
