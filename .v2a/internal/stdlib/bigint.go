package stdlib

import (
	"errors"
	"math/big"
	sl "math/big"

	"github.com/blackchip-org/zc/number"
)

type slBigInt struct {
	val *sl.Int
}

func unwrapI(bi number.BigInt) *sl.Int {
	return bi.(slBigInt).val
}

func wrapI(bi *sl.Int) slBigInt {
	return slBigInt{val: bi}
}

func (i slBigInt) Add(i2 number.BigInt) number.BigInt {
	var z big.Int
	return wrapI(z.Add(i.val, unwrapI(i2)))
}

func (i slBigInt) String() string {
	return i.val.String()
}

func ParseBigInt(s string) (number.BigInt, error) {
	var bi big.Int
	z, ok := bi.SetString(s, 0)
	if !ok {
		return wrapI(z), errors.New("unable to parse")
	}
	return wrapI(z), nil
}

func NewBigInt(i int) number.BigInt {
	bi := big.NewInt(int64(i))
	return wrapI(bi)
}

func UseBigInt() {
	number.UseBigInt(number.BigIntImpl{
		Parse: ParseBigInt,
		New:   NewBigInt,
	})
}
