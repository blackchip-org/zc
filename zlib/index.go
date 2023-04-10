package zlib

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/zlib/char/unicode_"
	"github.com/blackchip-org/zc/zlib/color_"
	"github.com/blackchip-org/zc/zlib/crypto_/rot"
	"github.com/blackchip-org/zc/zlib/dev"
	"github.com/blackchip-org/zc/zlib/dev/assert"
	"github.com/blackchip-org/zc/zlib/dev/bool_"
	"github.com/blackchip-org/zc/zlib/dev/dict"
	"github.com/blackchip-org/zc/zlib/dev/io_"
	"github.com/blackchip-org/zc/zlib/dev/runtime"
	"github.com/blackchip-org/zc/zlib/dev/test"
	"github.com/blackchip-org/zc/zlib/geo"
	"github.com/blackchip-org/zc/zlib/geo/epsg"
	"github.com/blackchip-org/zc/zlib/rand_"
	"github.com/blackchip-org/zc/zlib/rand_/dice"
	"github.com/blackchip-org/zc/zlib/time_"
	"github.com/blackchip-org/zc/zlib/time_/tz"
	"github.com/blackchip-org/zc/zlib/unit"
	"github.com/blackchip-org/zc/zlib/unit/si"
	"github.com/blackchip-org/zc/zlib/user/fn"
	"github.com/blackchip-org/zc/zlib/user/format"
	"github.com/blackchip-org/zc/zlib/user/math_"
	"github.com/blackchip-org/zc/zlib/user/prog"
	"github.com/blackchip-org/zc/zlib/user/sci"
	"github.com/blackchip-org/zc/zlib/user/stack"
	"github.com/blackchip-org/zc/zlib/user/str"
	"github.com/blackchip-org/zc/zlib/user/zc_"
)

var All = []zc.ModuleDef{
	assert.Mod,
	bool_.Mod,
	color_.Mod,
	dev.Mod,
	dice.Mod,
	dict.Mod,
	epsg.Mod,
	fn.Mod,
	format.Mod,
	geo.Mod,
	math_.Module,
	io_.Mod,
	prog.Mod,
	rand_.Mod,
	runtime.Mod,
	rot.Mod,
	sci.Mod,
	si.Mod,
	stack.Mod,
	str.Mod,
	test.Mod,
	time_.Mod,
	tz.Mod,
	unit.Mod,
	unicode_.Mod,
	zc_.Mod,
}

var (
	Preload = []string{
		"zc",
		"math",
		"bool",
		"assert",
		"dev",
		"format",
		"stack",
		"dict",
		"fn",
	}

	PreludeUser = []string{
		"bool",
		"format",
		"math",
		"stack",
		"str",
		"fn",
		"zc",
	}
	// Order here needs to be sorted based on dependencies. Do not put in
	// alphabetical order. Modules to the top of the list do not have access
	// to functions found at the bottom of the list.
	//
	// Also, update dev mode when this changes
	PreludeDev = []string{
		"assert",
		"bool",
		"format",
		"dev",
		"stack",
		"conf",
		"math",
		"str",
		"dict",
		"fn",
	}
)
