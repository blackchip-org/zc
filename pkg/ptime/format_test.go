package ptime

import (
	"testing"
	"time"

	"github.com/blackchip-org/zc/pkg/ptime/locale"
)

func TestFormatEnUS(t *testing.T) {
	tests := []struct {
		in     string
		layout string
		out    string
	}{
		{
			"2016-11-22",
			"[year]-[month]-[day]",
			"2016-11-22",
		},
		{
			"2016-11-22",
			"[year]-[day/year]",
			"2016-327",
		},
		{
			"2016-05-06",
			"[month]/[day]",
			"5/6",
		},
		{
			"2016-05-06",
			"[month]/[day/02]",
			"5/06",
		},
		{
			"2016-05-06",
			"[month/02]/[day/02]",
			"05/06",
		},
		{
			"2016-05-06",
			"[month/abbr] [day] [year]",
			"May 6 2016",
		},
		{
			"2016-05-06",
			"[month/abbr] [day/02] [year]",
			"May 06 2016",
		},
		{
			"2016-05-06",
			"[month/abbr] [day/2] [year]",
			"May  6 2016",
		},
		{
			"2016-05-06",
			"[month/abbr] [day] [year/2]",
			"May 6 16",
		},
		{
			"2016-05-06",
			"[weekday], [month/abbr] [day] [year/2]",
			"Friday, May 6 16",
		},
		{
			"2016-05-06",
			"[weekday], [month/abbr] [day] [year/2]",
			"Friday, May 6 16",
		},
		{
			"2016-05-06",
			"[weekday/abbr], [month/abbr] [day] [year/2]",
			"Fri, May 6 16",
		},
		{
			"17:30:25 MST -0700",
			"[hour]:[minute]:[second] [zone-offset]",
			"17:30:25 MST -0700",
		},
		{
			"17:30:25 UTC +0000",
			"[hour]:[minute]:[second] [zone-offset]",
			"17:30:25 UTC",
		},
		{
			"17:30:25 MST -0700",
			"[hour]:[minute]:[second] [zone-offset/:]",
			"17:30:25 MST -07:00",
		},
		{
			"17:30:25 UTC +0000",
			"[hour]:[minute]:[second] [zone-offset/:]",
			"17:30:25 UTC",
		},
		{
			"17:30:25 MST -0700",
			"[hour]:[minute]:[second] [offset-zone]",
			"17:30:25 -0700 MST",
		},
		{
			"17:30:25 UTC +0000",
			"[hour]:[minute]:[second] [offset-zone]",
			"17:30:25 UTC",
		},

		{
			"17:30:25",
			"[hour/12]:[minute][period]",
			"5:30PM",
		},
		{
			"17:30:25",
			"[hour/12]:[minute][period/narrow]",
			"5:30p",
		},
		{
			"17:30:25",
			"[hour/12]:[minute][period/alt]",
			"5:30pm",
		},
		{
			"17:30:25.1234",
			"[hour]:[minute]:[second/4]",
			"17:30:25.1234",
		},
		{
			"17:30:25.1234",
			"[hour]:[minute]:[second/2]",
			"17:30:25.12",
		},
		{
			"17:30:25",
			"[hour]:[minute]:[second]",
			"17:30:25",
		},
		{
			"17:30:00",
			"[hour]:[minute]:[second]",
			"17:30:00",
		},
	}

	l, err := time.LoadLocation("America/Denver")
	if err != nil {
		t.Fatalf("unable to load location: %v", err)
	}
	now := time.Date(2006, 01, 02, 15, 04, 05, 0, l)

	p := For(locale.EnUS)
	for _, test := range tests {
		t.Run(test.layout, func(t *testing.T) {
			parsed, err := p.Parse(test.in)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			pt, err := p.Time(parsed, now)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			out := p.Format(test.layout, pt)
			if out != test.out {
				t.Errorf("\n have: %v \n want: %v", out, test.out)
			}
		})
	}
}

func TestFormatFrFR(t *testing.T) {
	tests := []struct {
		in     string
		layout string
		out    string
	}{
		{
			"2016-11-22",
			"[day] [month/abbr] [year]",
			"22 nov. 2016",
		},
		{
			"2016-11-22",
			"[day] [month/name] [year]",
			"22 novembre 2016",
		},
		{
			"2016-01-02",
			"[weekday], [day]/[month]",
			"samedi, 2/1",
		},
		{
			"2016-01-02",
			"[weekday/abbr] [day]/[month]",
			"sam. 2/1",
		},
	}

	l := time.UTC
	now := time.Date(2006, 01, 02, 15, 04, 05, 0, l)

	p := For(locale.FrFR)
	for _, test := range tests {
		t.Run(test.layout, func(t *testing.T) {
			parsed, err := p.Parse(test.in)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			pt, err := p.Time(parsed, now)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			out := p.Format(test.layout, pt)
			if out != test.out {
				t.Errorf("\n have: %v \n want: %v", out, test.out)
			}
		})
	}
}
