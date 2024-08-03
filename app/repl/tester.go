package repl

import (
	"testing"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app"
	"github.com/blackchip-org/zc/v6/pkg/ansi"
)

type ReplTester struct {
	Calc *app.Calc
	Repl *Repl
	ct   *app.CalcTester
	t    *testing.T
}

func NewReplTester(t *testing.T) *ReplTester {
	ansi.Enabled = false
	ct := app.NewCalcTester(t)
	return &ReplTester{
		Calc: ct.Calc,
		Repl: New(ct.Calc),
		ct:   ct,
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
	if r.Calc.Error != nil {
		r.t.Logf("(!) %v\n", r.Calc.Error)
	}
	if r.Calc.Notice != "" {
		r.t.Logf("(?) %v", r.Calc.Notice)
	}
	r.t.Logf("%v\n", r.Calc.String())
}

func (r *ReplTester) AssertStack(vals ...any) {
	r.t.Helper()
	r.ct.AssertStack(vals...)
}

func (r *ReplTester) AssertError(err string) {
	r.t.Helper()
	if r.Repl.Error().Error() == err {
		return
	}
	r.ct.AssertError(err)
}

func (r *ReplTester) AssertNotice(notice string) {
	r.t.Helper()
	r.ct.AssertNotice(notice)
}
