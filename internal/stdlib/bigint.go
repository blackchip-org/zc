package stdlib

import (
	"fmt"
	"math/big"
	sl "math/big"

	"github.com/blackchip-org/zc/types"
)

type slBigInt struct {
	val *sl.Int
}

func unwrapI(bi types.BigInt) *sl.Int {
	return bi.(slBigInt).val
}

func wrapI(bi *sl.Int) slBigInt {
	return slBigInt{val: bi}
}

func (i slBigInt) Add(i2 types.BigInt) types.BigInt {
	var z big.Int
	return wrapI(z.Add(i.val, unwrapI(i2)))
}

func (i slBigInt) Type() types.Type {
	return types.BigIntType{}
}

func (i slBigInt) Value() types.Value {
	return i
}

func (i slBigInt) Native() any {
	return i
}

func (i slBigInt) String() string {
	return i.val.String()
}

func ParseBigInt(s string) (types.BigInt, error) {
	var bi big.Int
	z, ok := bi.SetString(s, 0)
	if !ok {
		return wrapI(z), fmt.Errorf("unable to parse: %v", s)
	}
	return wrapI(z), nil
}

func NewBigInt(i int) types.BigInt {
	bi := big.NewInt(int64(i))
	return wrapI(bi)
}

func UseBigInt() {
	types.UseBigInt(types.BigIntImpl{
		Parse: ParseBigInt,
		New:   NewBigInt,
	})
}
