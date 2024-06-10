package types

import (
	"github.com/blackchip-org/zc/v6"
)

const (
	Nil zc.TypeID = iota
	Int
	IntA
	Int64
	Int32
	Int16
	Int8
	IntAU
	Int64U
	Int32U
	Int16U
	Int8U
	Dec
	Float
	Float64
	Float32
	Rational
	Complex64
	Complex128
	Bool
	Text
	Char
	Val
)
