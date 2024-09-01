package ptime

import (
	"testing"
	"time"

	"github.com/blackchip-org/zc/v6/pkg/ptime/locale"
)

func TestTimeEnUS(t *testing.T) {
	nowZ := time.FixedZone("MST", -7*3600)
	estZ := time.FixedZone("EST", -5*3600)
	now := time.Date(2006, 01, 02, 15, 04, 05, 00, nowZ)

	tests := []struct {
		name   string
		parsed Parsed
		time   time.Time
	}{
		{
			"2016-11-22",
			Parsed{Year: "2016", Month: "11", Day: "22"},
			time.Date(2016, 11, 22, 0, 0, 0, 0, nowZ),
		},
		{
			"2016-11",
			Parsed{Year: "2016", Month: "11"},
			time.Date(2016, 11, 1, 0, 0, 0, 0, nowZ),
		},
		{
			"2016-075",
			Parsed{Year: "2016", Day: "075"},
			time.Date(2016, 3, 15, 0, 0, 0, 0, nowZ),
		},
		{
			"22 Nov 2016",
			Parsed{Day: "22", Month: "Nov", Year: "2016"},
			time.Date(2016, 11, 22, 0, 0, 0, 0, nowZ),
		},
		{
			"22 Nov 16",
			Parsed{Day: "22", Month: "Nov", Year: "16"},
			time.Date(2016, 11, 22, 0, 0, 0, 0, nowZ),
		},
		{
			"22:33",
			Parsed{Hour: "22", Minute: "33"},
			time.Date(2006, 01, 02, 22, 33, 0, 0, nowZ),
		},
		{
			"10:33pm",
			Parsed{Hour: "10", Minute: "33", Period: "pm"},
			time.Date(2006, 01, 02, 22, 33, 0, 0, nowZ),
		},
		{
			"22:33:44",
			Parsed{Hour: "22", Minute: "33", Second: "44"},
			time.Date(2006, 01, 02, 22, 33, 44, 0, nowZ),
		},
		{
			"22:33:44.55",
			Parsed{Hour: "22", Minute: "33", Second: "44", FracSecond: "55"},
			time.Date(2006, 01, 02, 22, 33, 44, fsecToNsec(55), nowZ),
		},
		{
			"22:33:44 MST -0700",
			Parsed{Hour: "22", Minute: "33", Second: "44", Zone: "MST", Offset: "-0700"},
			time.Date(2006, 01, 02, 22, 33, 44, 0, nowZ),
		},
		{
			"22:33:44 EST -0500",
			Parsed{Hour: "22", Minute: "33", Second: "44", Zone: "EST", Offset: "-0500"},
			time.Date(2006, 01, 02, 22, 33, 44, 0, estZ),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tt, err := Time(locale.EnUS, test.parsed, now)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tt.Equal(test.time) {
				t.Errorf("\n have: %v \n want: %v", tt, test.time)
			}
		})
	}
}

func TestTimeFrFR(t *testing.T) {
	nowZ := time.UTC
	now := time.Date(2006, 01, 02, 15, 04, 05, 00, nowZ)

	tests := []struct {
		name   string
		parsed Parsed
		time   time.Time
	}{
		{
			"2 janvier 2016",
			Parsed{Year: "2016", Month: "janvier", Day: "2"},
			time.Date(2016, 1, 2, 0, 0, 0, 0, nowZ),
		},
		{
			"2 janv. 2016",
			Parsed{Year: "2016", Month: "janv.", Day: "2"},
			time.Date(2016, 1, 2, 0, 0, 0, 0, nowZ),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tt, err := Time(locale.FrFR, test.parsed, now)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tt.Equal(test.time) {
				t.Errorf("\n have: %v \n want: %v", tt, test.time)
			}
		})
	}
}
