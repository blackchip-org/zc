package calc

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/zlib/zlib_math"
)

func zlibLoad() zc.Library {
	l := NewLibrary()
	l.Define(zlib_math.Def)
	return l
}

var zlib = zlibLoad()
