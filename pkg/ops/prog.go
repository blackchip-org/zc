package ops

import (
	"fmt"
	"math"
	"math/big"

	"github.com/blackchip-org/zc/v5/pkg/zc"
)

/*
oper	and
func	AndBigInt p0:BigInt p1:BigInt -- BigInt
title	Bitwise and

desc
The bitwise and of `p0` and `p1`
end

example
0b1100 -- 0b1100
0b1010 -- 0b1100 | 0b1010
and bin -- 0b1000
end
*/
func AndBigInt(c zc.Calc) {
	var r0 big.Int
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0.And(a0, a1)
	zc.PushBigInt(c, &r0)
}

/*
oper	bin
func	Bin p0:BigInt -- BigInt
title	Binary conversion

desc
Convert the value of *p0* to a binary number.
end

example
0x0f -- 0x0f
bin -- 0b1111
end
*/
func Bin(c zc.Calc) {
	var zero big.Int
	var r0 string
	a0 := zc.PopBigInt(c)
	if a0.Cmp(&zero) < 0 {
		var t0 big.Int
		t0.Abs(a0)
		r0 = fmt.Sprintf("-0b%b", &t0)
	} else {
		r0 = fmt.Sprintf("0b%b", a0)
	}
	zc.PushString(c, r0)
}

/*
oper	bit
func	Bit p0:BigInt n:Int -- Uint
title	Bit value

desc
The value of the *n*th bit of *p0*.
end

example
0b100 -- 0b100
2 bit -- 1
end
*/
func Bit(c zc.Calc) {
	i := zc.PopInt(c)
	a0 := zc.PopBigInt(c)
	r0 := a0.Bit(i)
	zc.PushUint(c, r0)
}

/*
oper	bits
func	Bits p0:BigInt -- Int
title	Length in bits

desc
The length of *p0* in bits.
end

example
0b11111 -- 0b11111
bits -- 5
end
*/
func Bits(c zc.Calc) {
	a0 := zc.PopBigInt(c)
	r0 := a0.BitLen()
	zc.PushInt(c, r0)
}

/*
oper	bytes
func	Bytes p0:BigInt -- Int
title	Length in bytes

desc
The length of *p0* in bytes.
end

example
0x1ff -- 0x1ff
bytes -- 2
end
*/
func Bytes(c zc.Calc) {
	a0 := zc.PopBigInt(c)
	r0 := int(math.Ceil(float64(a0.BitLen()) / 8.0))
	zc.PushInt(c, r0)
}

/*
oper	dec
func	DecBigInt p0:BigInt -- BigInt
title	Decimal conversion

desc
Convert the value of *p0* to a decimal number.
end

example
0xf -- 0xf
dec -- 15
end
*/
func DecBigInt(c zc.Calc) {
	a0 := zc.PopBigInt(c)
	zc.PushBigInt(c, a0)
}

/*
oper	hex
func	HexBigInt p0:BigInt -- BigInt
title	Hexadecimal conversion

desc
Convert the value of *p0* to a hexadecimal number.
end

example
0b1111 -- 0b1111
hex -- 0xf
end
*/
func HexBigInt(c zc.Calc) {
	var zero big.Int
	var r0 string
	a0 := zc.PopBigInt(c)
	if a0.Cmp(&zero) < 0 {
		var t0 big.Int
		t0.Abs(a0)
		r0 = fmt.Sprintf("-0x%x", &t0)
	} else {
		r0 = fmt.Sprintf("0x%x", a0)
	}
	zc.PushString(c, r0)
}

func Int(c zc.Calc) {
	a0 := zc.PopBigInt(c)
	zc.PushBigInt(c, a0)
}

/*
oper	lsh
func	Lsh p0:BigInt n:Uint -- BigInt
alias 	left-shift
title	Left shift bits

desc
Shifts all bits in *p0* to the left by *n*.
end

example
0b10 -- 0b10
2 lsh bin -- 0b1000
end
*/
func Lsh(c zc.Calc) {
	var r0 big.Int
	n := zc.PopUint(c)
	a0 := zc.PopBigInt(c)
	r0.Lsh(a0, n)
	zc.PushBigInt(c, &r0)
}

/*
oper	not
func	NotBigInt p0:BigInt -- BigInt
title	Bitwise not

desc
The bitwise not of *p0*.
end

example
0b101 -- 0b101
not bin -- -0b110
end
*/
func NotBigInt(c zc.Calc) {
	var r0 big.Int
	a0 := zc.PopBigInt(c)
	r0.Not(a0)
	zc.PushBigInt(c, &r0)
}

/*
oper	oct
func	Oct p0:BigInt -- BigInt
title	Octal conversion

desc
Convert the value of *p0* to an octal number.
end

example
0b1111 -- 0b1111
oct -- 0o17
end
*/
func Oct(c zc.Calc) {
	var zero big.Int
	var r0 string
	a0 := zc.PopBigInt(c)
	if a0.Cmp(&zero) < 0 {
		var t0 big.Int
		t0.Abs(a0)
		r0 = fmt.Sprintf("-0o%o", &t0)
	} else {
		r0 = fmt.Sprintf("0o%o", a0)
	}
	zc.PushString(c, r0)
}

/*
oper	or
func	OrBigInt p0:BigInt p1:BigInt -- BigInt
title	Bitwise or

desc
The bitwise or of *p0* and *p1*.
end

example
0b1100 -- 0b1100
0b1010 -- 0b1100 | 0b1010
or bin -- 0b1110
end
*/
func OrBigInt(c zc.Calc) {
	var r0 big.Int
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0.Or(a0, a1)
	zc.PushBigInt(c, &r0)
}

/*
oper	rsh
func	Rsh p0:BigInt n:Uint -- BigInt
alias	right-shift
title	Right shift bits

desc
Shifts all bits in *p0* to the right by *n*.
end

example
0b1000 -- 0b1000
2 rsh bin -- 0b10
end
*/
func Rsh(c zc.Calc) {
	var r0 big.Int
	n := zc.PopUint(c)
	a0 := zc.PopBigInt(c)
	r0.Rsh(a0, n)
	zc.PushBigInt(c, &r0)
}

/*
oper	xor
func	Xor p0:BigInt p1:BigInt -- BigInt
title	Bitwise exclusive or

desc
The bitwise exclusive or of `p0` and `p1`.
end

example
0b1100 -- 0b1100
0b1010 -- 0b1100 | 0b1010
xor bin -- 0b110
end
*/
func Xor(c zc.Calc) {
	var r0 big.Int
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0.Xor(a0, a1)
	zc.PushBigInt(c, &r0)
}
