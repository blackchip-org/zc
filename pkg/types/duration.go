package types

import (
	"strings"
	"time"
	"unicode"

	"github.com/blackchip-org/zc/pkg/scanner"
)

func ParseDuration(str string) (time.Duration, bool) {
	s := scanner.NewString(str)
	for s.Ok() {
		s.SkipIf(unicode.IsSpace)
	}
	d, err := time.ParseDuration(s.Token())
	if err != nil {
		return time.Duration(0), false
	}
	return d, true
}

func FormatDuration(d time.Duration) string {
	s := scanner.NewString(d.String())
	var hrs, min, sec string
	for s.Ok() {
		switch s.Ch {
		case 'h':
			hrs = s.TrimToken()
			s.Next()
		case 'm':
			min = s.TrimToken()
			s.Next()
		case 's':
			sec = s.TrimToken()
			s.Next()
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
