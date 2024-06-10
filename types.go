package zc

type ConvFunc func(Type, Type, any) (any, bool)

type Type interface {
	Name() string
	TypeConvs() []TypeConv
	Conv(Type, Type, any) (any, bool)
}

type TypeConv struct {
	From Type
	To   Type
}

var BigInt = bigIntType{}

type bigIntType struct{}

func (t bigIntType) Name() string {
	return "Int"
}

func (t bigIntType) ConvFunc
