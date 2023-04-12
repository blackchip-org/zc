package types

type BigInt interface {
	Value
	Add(BigInt) BigInt
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

type BigIntType struct{}

func (t BigIntType) String() string { return "BigInt" }

func (t BigIntType) Parse(s string) (BigInt, error) {
	return implBigInt.Parse(cleanNumber(s))
}

func (t BigIntType) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return nil, err
	}
	return t.Value(v), nil
}

func (t BigIntType) Format(i BigInt) string {
	return i.String()
}

func (t BigIntType) Value(i BigInt) Value {
	return i
}

func (t BigIntType) Native(v Value) BigInt {
	return v.Native().(BigInt)
}
