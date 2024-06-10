package zc

import "testing"

func AssertStack(t *testing.T, s Stack, vals ...any) {
	t.Helper()

	fmtHave := s.String()
	fmtWant := FormatList(vals...)

	if fmtHave != fmtWant {
		t.Fatalf("\n have: %v \n want: %v", fmtHave, fmtWant)
	}
}
