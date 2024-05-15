package zc

import "github.com/blackchip-org/scan"

type PreParser func(*scan.Scanner) string

func PreParseNumber(s *scan.Scanner) string {
	for s.HasMore() {
		switch {
		case scan.IsCurrency(s.This):
			s.Skip()
		case s.This == ',':
			s.Skip()
		default:
			s.Keep()
		}
	}
	return s.Emit().Val
}
