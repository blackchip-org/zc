package ints

import (
	"github.com/blackchip-org/zc/v6"
)

func newCalc() *zc.Calc {
	cat := zc.NewCatalogBuilder()
	cat.AddVolume(Volume)
	return zc.NewCalc(cat.Build())
}

// func TestConversions(t *testing.T) {
// 	var (
// 		rint   int
// 		rint8  int8
// 		rint16 int16
// 	)

// 	tests = []struct {
// 		in any
// 	}
// }
