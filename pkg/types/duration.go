package types

import (
	"strings"
	"time"

	"github.com/blackchip-org/scan"
)

func ParseDuration(str string) (time.Duration, bool) {
	s := scan.NewScannerFromString("", str)
	for s.HasMore() {
		if scan.IsSpace(s.This) {
			s.Skip()
		} else {
			s.Keep()
		}
	}
	d, err := time.ParseDuration(s.Emit().Val)
	if err != nil {
		return time.Duration(0), false
	}
	return d, true
}

func FormatDuration(d time.Duration) string {
	s := scan.NewScannerFromString("", d.String())
	var hrs, min, sec string
	for s.HasMore() {
		switch s.This {
		case 'h':
			hrs = strings.TrimSpace(s.Emit().Val)
			s.Discard()
		case 'm':
			min = strings.TrimSpace(s.Emit().Val)
			s.Discard()
		case 's':
			sec = strings.TrimSpace(s.Emit().Val)
			s.Discard()
		default:
			s.Keep()
		}
	}
	var fields []string
	if hrs != "" && hrs != "0" {
		fields = append(fields, hrs+"h")
	}
	if min != "" && min != "0" {
		fields = append(fields, min+"m")
	}
	if sec != "" && sec != "0" {
		fields = append(fields, sec+"s")
	}
	return strings.Join(fields, " ")
}
