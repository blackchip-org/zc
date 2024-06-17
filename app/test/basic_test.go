package test

import (
	"testing"

	"github.com/blackchip-org/zc/v6/app"
	"github.com/blackchip-org/zc/v6/app/ops"
	"github.com/blackchip-org/zc/v6/app/types"
)

func TestAdd(t *testing.T) {
	var c app.Calc

	c.PushVal(6, 2)
	c.Do(ops.AddBigInt)

	have := types.BigInt.As(c.Top()).String()
	want := "8"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}
