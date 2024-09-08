package repl

import (
	"reflect"
	"testing"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app"
	"github.com/blackchip-org/zc/v6/pkg/ansi"
)

type ReplTester struct {
	Calc *app.Calc
	Repl *Repl
	t    *testing.T
}

func NewReplTester(t *testing.T) *ReplTester {
	ansi.Enabled = false
	calc := app.NewCalc()
	return &ReplTester{
		Calc: calc,
		Repl: New(calc),
		t:    t,
	}
}

func (r *ReplTester) Eval(line string) {
	r.t.Helper()
	r.t.Logf("%v> %v\n", zc.ProgName, line)
	r.Repl.Eval(line)
	if r.Repl.Error() != nil {
		r.t.Logf("(!) %v\n", r.Repl.Error())
	}
	if r.Repl.Notice() != "" {
		r.t.Logf("(?) %v", r.Repl.Notice())
	}
	r.t.Logf("%v\n", r.Calc.String())
}

func (r *ReplTester) AssertStack(vals ...any) {
	r.t.Helper()

	fmtWant := zc.FormatList(zc.Strings(vals...))
	fmtHave := r.Calc.String()

	if r.Repl.Error() != nil {
		r.t.Fatalf("(FAIL) unexpected error")
	}
	if r.Repl.Notice() != "" {
		r.t.Fatalf("(FAIL) unexpected notice")
	}
	if !reflect.DeepEqual(fmtHave, fmtWant) {
		r.t.Fatalf("(FAIL) expected: %v", fmtWant)
	}
}

func (r *ReplTester) AssertError(msg string) {
	r.t.Helper()
	err := r.Repl.Error()
	if err == nil || err.Error() != msg {
		r.t.Fatalf("(FAIL) expected error: %v", msg)
	}
}

func (r *ReplTester) AssertNotice(notice string) {
	r.t.Helper()
	if r.Repl.notice != notice {
		r.t.Fatalf("(FAIL) expected notice: %v", notice)
	}
	r.Repl.notice = ""
}
