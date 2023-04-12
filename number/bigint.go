package number

type BigInt interface {
	Add(BigInt) BigInt
	String() string
}

type BigIntImpl struct {
	Parse func(string) (BigInt, error)
	New   func(int) BigInt
}

var implBigInt = BigIntImpl{
	Parse: func(string) (BigInt, error) { return nil, ErrNotSupported },
	New:   func(int) BigInt { return nil },
}

func UseBigInt(impl BigIntImpl) {
	implBigInt = impl
}

func ParseBigInt(s string) (BigInt, error) {
	return implBigInt.Parse(s)
}

func MustParseBigInt(s string) BigInt {
	d, err := implBigInt.Parse(s)
	if err != nil {
		panic(err)
	}
	return d
}

func NewBigInt(i int) BigInt {
	return implBigInt.New(i)
}
