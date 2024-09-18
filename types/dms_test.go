package types

import (
	"strconv"
	"testing"

	"github.com/blackchip-org/dms"
	"github.com/blackchip-org/zc/v6/app/vars"
	"github.com/cockroachdb/apd/v3"
)

func TestDMS(t *testing.T) {
	c := apd.BaseContext.WithPrecision(vars.DefaultDecPrec)

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
			dms, err := NewAngleDMSFromFields(c, test.fields)
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
	c := apd.BaseContext.WithPrecision(vars.DefaultDecPrec)

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
			dms, err := NewAngleDMSFromFields(c, test.fields)
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

func TestFormat(t *testing.T) {
	c := apd.BaseContext.WithPrecision(vars.DefaultDecPrec)

	tests := []struct {
		angle  AngleDMS
		to     dms.Unit
		places int
		result string
	}{
		{NewAngleDMSFromFloat(c, 1, 0, 0), dms.SecUnit, -1, "1° 0′ 0″"},
		{NewAngleDMSFromFloat(c, 1, 2, 0), dms.SecUnit, -1, "1° 2′ 0″"},
		{NewAngleDMSFromFloat(c, 1, 2, 3.33), dms.SecUnit, -1, "1° 2′ 3.33″"},
		{NewAngleDMSFromFloat(c, -1, 2, 3.33), dms.SecUnit, -1, "-1° 2′ 3.33″"},
		{NewAngleDMSFromFloat(c, 1, -2, -3.33), dms.SecUnit, -1, "1° 2′ 3.33″"},
		{NewAngleDMSFromFloat(c, 1, -2, -3.36), dms.SecUnit, 1, "1° 2′ 3.4″"},
		{NewAngleDMSFromFloat(c, 1.051667, 0, 0), dms.SecUnit, -1, "1° 3′ 6.0012″"},
		{NewAngleDMSFromFloat(c, -1.051667, 0, 0), dms.SecUnit, -1, "-1° 3′ 6.0012″"},
		{NewAngleDMSFromFloat(c, 1.5, 0, 0), dms.SecUnit, -1, "1° 30′ 0″"},
		{NewAngleDMSFromFloat(c, 1.5, 10, 0), dms.SecUnit, -1, "1° 40′ 0″"},
		{NewAngleDMSFromFloat(c, -1.5, 10, 0), dms.SecUnit, -1, "-1° 40′ 0″"},
		{NewAngleDMSFromFloat(c, -1.5, -10, 0), dms.SecUnit, -1, "-1° 40′ 0″"},
		{NewAngleDMSFromFloat(c, 1.5, 10.5, 10), dms.SecUnit, -1, "1° 40′ 40″"},
		{NewAngleDMSFromFloat(c, 0, 0, 75), dms.SecUnit, -1, "0° 1′ 15″"},
		{NewAngleDMSFromFloat(c, 0, 10, 135), dms.SecUnit, -1, "0° 12′ 15″"},
		{NewAngleDMSFromFloat(c, 0, 59, 135), dms.SecUnit, -1, "1° 1′ 15″"},

		{NewAngleDMSFromFloat(c, 1, 0, 0), dms.MinUnit, 3, "1° 0.000′"},
		{NewAngleDMSFromFloat(c, 1, 2, 0), dms.MinUnit, 3, "1° 2.000′"},
		{NewAngleDMSFromFloat(c, 1, 2, 6), dms.MinUnit, 3, "1° 2.100′"},

		{NewAngleDMSFromFloat(c, 1, 0, 0), dms.DegUnit, 6, "1.000000°"},
		{NewAngleDMSFromFloat(c, 1, 3, 0), dms.DegUnit, 6, "1.050000°"},
		{NewAngleDMSFromFloat(c, 1, 3, 9), dms.DegUnit, 6, "1.052500°"},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := FormatAngleDMS(test.angle, test.to, test.places)
			if result != test.result {
				t.Errorf("\n have: %v \n want: %v", result, test.result)
			}
		})
	}
}
