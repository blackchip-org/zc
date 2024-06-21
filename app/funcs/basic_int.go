package funcs

import "github.com/blackchip-org/zc/v6"

func add8(x, y uint8, carry bool) (z uint8, c, v bool) {
	// https://stackoverflow.com/questions/8034566/overflow-and-carry-flags-on-z80/8037485#8037485
	var carryOut uint8

	if carry {
		if x >= 0xff-y {
			carryOut = 1
		}
		z = x + y + 1
	} else {
		if x > 0xff-y {
			carryOut = 1
		}
		z = x + y
	}
	carryIns := z ^ x ^ y

	c = carryOut != 0
	v = (carryIns>>7)^carryOut != 0
	return
}

func intFlags(c, v bool) (flags uint64) {
	if c {
		flags |= zc.FlagCarry
	}
	if v {
		flags |= zc.FlagOverflow
	}
	return
}

func AddInt8(e *zc.OpEnv) {
	y := zc.Int8.Pop(e)
	x := zc.Int8.Pop(e)
	z, c, v := add8(uint8(x), uint8(y), false)
	zc.Int8.Push(e, int8(z))
	e.Flags(intFlags(c, v))
}

func AddUint8(e *zc.OpEnv) {
	y := zc.Uint8.Pop(e)
	x := zc.Uint8.Pop(e)
	z, c, v := add8(x, y, false)
	zc.Uint8.Push(e, z)
	e.Flags(intFlags(c, v))
}

func NegUint8(e *zc.OpEnv) {
	x := zc.Uint8.Pop(e)
	x = -x
	zc.Uint8.Push(e, x)
}

func SubUint8(e *zc.OpEnv) {
	y := zc.Uint8.Pop(e)
	x := zc.Uint8.Pop(e)
	z, c, v := add8(x, ^y, true)
	zc.Uint8.Push(e, z)
	e.Flags(intFlags(!c, !v))
}
