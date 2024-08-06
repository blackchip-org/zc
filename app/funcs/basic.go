package funcs

import (
	"math"
	"math/big"
	"math/cmplx"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/vars"
	"github.com/cockroachdb/apd/v3"
)

// ----------------------------------------------------------------------------
func AbsBigInt(c zc.Calc) {
	x := zc.BigInt.Pop(c)
	x.Abs(x)
	zc.BigInt.Push(c, x)
}

func AbsComplex(c zc.Calc) {
	x := zc.Complex.Pop(c)
	z := cmplx.Abs(x)
	zc.Float64.Push(c, z)
}

func AbsDecimal(c zc.Calc) {
	x := zc.Decimal.Pop(c)
	x.Abs(x)
	zc.Decimal.Push(c, x)
}

// ----------------------------------------------------------------------------
func AddBigInt(c zc.Calc) {
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	x.Add(x, y)
	zc.BigInt.Push(c, x)
	zc.BigInt.Recycle(y)
}

func AddBigFloat(c zc.Calc) {
	y := zc.BigFloat.Pop(c)
	x := zc.BigFloat.Pop(c)
	x.Add(x, y)
	zc.BigFloat.Push(c, x)
	zc.BigFloat.Recycle(y)
}

func AddComplex(c zc.Calc) {
	y := zc.Complex.Pop(c)
	x := zc.Complex.Pop(c)
	z := x + y
	zc.Complex.Push(c, z)
}

func AddDecimal(c zc.Calc) {
	d := vars.ForConf(c).DecMath
	y := zc.Decimal.Pop(c)
	x := zc.Decimal.Pop(c)
	defer zc.Decimal.Recycle(y)

	_, err := d.Add(x, x, y)
	if err != nil {
		c.Raise(err)
		return
	}
	zc.Decimal.Push(c, x)
}

func AddRat(c zc.Calc) {
	y := zc.Rat.Pop(c)
	x := zc.Rat.Pop(c)
	x.Add(x, y)
	zc.Rat.Push(c, x)
	zc.Rat.Recycle(y)
}

// ----------------------------------------------------------------------------
func CbrtDecimal(c zc.Calc) {
	var zero apd.Decimal
	d := vars.ForConf(c).DecMath
	x := zc.Decimal.Pop(c)

	if x.Cmp(&zero) < 0 {
		c.Raise(zc.ErrInvalidArg("%v < 0", x))
		return
	}
	_, err := d.Cbrt(x, x)
	if err != nil {
		c.Raise(err)
		return
	}
	zc.Decimal.Push(c, x)
}

func CbrtFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	if x < 0 {
		c.Raise(zc.ErrInvalidArg("%v < 0", x))
		return
	}
	z := math.Cbrt(x)
	zc.Float64.Push(c, z)
}

// ----------------------------------------------------------------------------
func CeilDecimal(c zc.Calc) {
	d := vars.ForConf(c).DecMath
	x := zc.Decimal.Pop(c)
	_, err := d.Ceil(x, x)
	if err != nil {
		c.Raise(err)
		return
	}
	zc.Decimal.Push(c, x)
}

// ----------------------------------------------------------------------------
func DivBigInt(c zc.Calc) {
	var zero big.Int
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		c.Raise(zc.ErrDivisionByZero)
		return
	}
	x.Div(x, y)
	zc.BigInt.Push(c, x)
}

func DivBigFloat(c zc.Calc) {
	var zero big.Float
	y := zc.BigFloat.Pop(c)
	x := zc.BigFloat.Pop(c)

	if y.Cmp(&zero) == 0 {
		c.Raise(zc.ErrDivisionByZero)
		return
	}

	x.Quo(x, y)
	zc.BigFloat.Push(c, x)
	zc.BigFloat.Recycle(y)
}

func DivComplex(c zc.Calc) {
	y := zc.Complex.Pop(c)
	x := zc.Complex.Pop(c)
	z := x / y
	zc.Complex.Push(c, z)
}

func DivDecimal(c zc.Calc) {
	d := vars.ForConf(c).DecMath
	y := zc.Decimal.Pop(c)
	x := zc.Decimal.Pop(c)
	defer zc.Decimal.Recycle(y)

	cond, err := d.Quo(x, x, y)
	if cond.DivisionByZero() {
		c.Raise(zc.ErrDivisionByZero)
		return
	}
	if err != nil {
		c.Raise(err)
		return
	}
	zc.Decimal.Push(c, x)
}

func DivRat(c zc.Calc) {
	var zero big.Rat
	y := zc.Rat.Pop(c)
	x := zc.Rat.Pop(c)
	defer zc.Rat.Recycle(y)

	if y.Cmp(&zero) == 0 {
		c.Raise(zc.ErrDivisionByZero)
		return
	}

	x.Quo(x, y)
	zc.Rat.Push(c, x)
}

// ----------------------------------------------------------------------------
func DivModBigInt(c zc.Calc) {
	var zero big.Int
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	m := zc.BigInt.New()
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		c.Raise(zc.ErrDivisionByZero)
		return
	}
	x.DivMod(x, y, m)
	zc.BigInt.Push(c, x)
	c.SetLabel("quo")
	zc.BigInt.Push(c, m)
	c.SetLabel("mod")
}

// ----------------------------------------------------------------------------
func ExpDecimal(c zc.Calc) {
	d := vars.ForConf(c).DecMath
	x := zc.Decimal.Pop(c)
	_, err := d.Exp(x, x)
	if err != nil {
		c.Raise(err)
		return
	}
	zc.Decimal.Push(c, x)
}

func ExpComplex(c zc.Calc) {
	x := zc.Complex.Pop(c)
	z := cmplx.Exp(x)
	zc.Complex.Push(c, z)
}

func ExpFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	z := math.Exp(x)
	zc.Float64.Push(c, z)
}

// ----------------------------------------------------------------------------
func FloorDecimal(c zc.Calc) {
	d := vars.ForConf(c).DecMath
	x := zc.Decimal.Pop(c)
	_, err := d.Floor(x, x)
	if err != nil {
		c.Raise(err)
		return
	}
	zc.Decimal.Push(c, x)
}

// ----------------------------------------------------------------------------
func LnDecimal(c zc.Calc) {
	d := vars.ForConf(c).DecMath
	x := zc.Decimal.Pop(c)
	_, err := d.Ln(x, x)
	if err != nil {
		c.Raise(err)
		return
	}
	zc.Decimal.Push(c, x)
}

// ----------------------------------------------------------------------------
func Log10Decimal(c zc.Calc) {
	d := vars.ForConf(c).DecMath
	x := zc.Decimal.Pop(c)
	_, err := d.Log10(x, x)
	if err != nil {
		c.Raise(err)
		return
	}
	zc.Decimal.Push(c, x)
}

func Log10Complex(c zc.Calc) {
	x := zc.Complex.Pop(c)
	z := cmplx.Log10(x)
	zc.Complex.Push(c, z)
}

func Log10Float64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	z := math.Log10(x)
	zc.Float64.Push(c, z)
}

// ----------------------------------------------------------------------------
func ModBigInt(c zc.Calc) {
	var zero big.Int
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		c.Raise(zc.ErrDivisionByZero)
		return
	}
	x.Mod(x, y)
	zc.BigInt.Push(c, x)
}

// ----------------------------------------------------------------------------
func MulBigInt(c zc.Calc) {
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	x.Mul(x, y)
	zc.BigInt.Push(c, x)
	zc.BigInt.Recycle(y)
}

func MulBigFloat(c zc.Calc) {
	y := zc.BigFloat.Pop(c)
	x := zc.BigFloat.Pop(c)
	x.Mul(x, y)
	zc.BigFloat.Push(c, x)
	zc.BigFloat.Recycle(y)
}

func MulComplex(c zc.Calc) {
	y := zc.Complex.Pop(c)
	x := zc.Complex.Pop(c)
	z := x * y
	zc.Complex.Push(c, z)
}

func MulDecimal(c zc.Calc) {
	d := vars.ForConf(c).DecMath
	y := zc.Decimal.Pop(c)
	x := zc.Decimal.Pop(c)
	defer zc.Decimal.Recycle(y)

	_, err := d.Mul(x, x, y)
	if err != nil {
		c.Raise(err)
		return
	}
	zc.Decimal.Push(c, x)
}

func MulRat(c zc.Calc) {
	y := zc.Rat.Pop(c)
	x := zc.Rat.Pop(c)
	x.Mul(x, y)
	zc.Rat.Push(c, x)
	zc.Rat.Recycle(y)
}

// ----------------------------------------------------------------------------
func NegBigInt(c zc.Calc) {
	x := zc.BigInt.Pop(c)
	x.Neg(x)
	zc.BigInt.Push(c, x)
}

func NegBigFloat(c zc.Calc) {
	x := zc.BigFloat.Pop(c)
	x.Neg(x)
	zc.BigFloat.Push(c, x)
}

func NegComplex(c zc.Calc) {
	x := zc.Complex.Pop(c)
	z := -x
	zc.Complex.Push(c, z)
}

func NegDecimal(c zc.Calc) {
	d := vars.ForConf(c).DecMath
	x := zc.Decimal.Pop(c)
	_, err := d.Neg(x, x)
	if err != nil {
		c.Raise(err)
		return
	}
	zc.Decimal.Push(c, x)
}

func NegRat(c zc.Calc) {
	x := zc.Rat.Pop(c)
	x.Neg(x)
	zc.Rat.Push(c, x)
}

// ----------------------------------------------------------------------------
func PowBigInt(c zc.Calc) {
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	x.Exp(x, y, nil)
	zc.BigInt.Push(c, x)
	zc.BigInt.Recycle(y)
}

func PowComplex(c zc.Calc) {
	y := zc.Complex.Pop(c)
	x := zc.Complex.Pop(c)
	z := cmplx.Pow(x, y)
	zc.Complex.Push(c, z)
}

func PowDecimal(c zc.Calc) {
	d := vars.ForConf(c).DecMath
	y := zc.Decimal.Pop(c)
	x := zc.Decimal.Pop(c)
	defer zc.Decimal.Recycle(y)

	_, err := d.Pow(x, x, y)
	if err != nil {
		c.Raise(err)
		return
	}
	zc.Decimal.Push(c, x)
}

// ----------------------------------------------------------------------------
func RemBigInt(c zc.Calc) {
	var zero big.Int
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		c.Raise(zc.ErrDivisionByZero)
		return
	}
	x.Rem(x, y)
	zc.BigInt.Push(c, x)
}

func RemDecimal(c zc.Calc) {
	d := vars.ForConf(c).DecMath
	y := zc.Decimal.Pop(c)
	x := zc.Decimal.Pop(c)
	defer zc.Decimal.Recycle(y)

	if y.IsZero() {
		c.Raise(zc.ErrDivisionByZero)
		return
	}
	_, err := d.Rem(x, x, y)
	if err != nil {
		c.Raise(err)
		return
	}
	zc.Decimal.Push(c, x)
}

// ----------------------------------------------------------------------------
func SignBigInt(c zc.Calc) {
	x := zc.BigInt.Pop(c)
	zc.Int.Push(c, x.Sign())
	zc.BigInt.Recycle(x)
}

func SignBigFloat(c zc.Calc) {
	x := zc.BigFloat.Pop(c)
	zc.Int.Push(c, x.Sign())
	zc.BigFloat.Recycle(x)
}

func SignDecimal(c zc.Calc) {
	x := zc.Decimal.Pop(c)
	zc.Int.Push(c, x.Sign())
	zc.Decimal.Recycle(x)
}

func SignRat(c zc.Calc) {
	x := zc.Rat.Pop(c)
	s := x.Sign()
	zc.Int.Push(c, s)
	zc.Rat.Recycle(x)
}

// ----------------------------------------------------------------------------
func SqrtBigInt(c zc.Calc) {
	var zero big.Int
	x := zc.BigInt.Pop(c)
	if x.Cmp(&zero) < 0 {
		c.Raise(zc.ErrInvalidArg("%v < 0", x.String()))
		return
	}
	x.Sqrt(x)
	zc.BigInt.Push(c, x)
}

func SqrtBigFloat(c zc.Calc) {
	var zero big.Float
	x := zc.BigFloat.Pop(c)
	if x.Cmp(&zero) < 0 {
		c.Raise(zc.ErrInvalidArg("%v < 0", x.String()))
		return
	}
	x.Sqrt(x)
	zc.BigFloat.Push(c, x)
}

func SqrtComplex(c zc.Calc) {
	x := zc.Complex.Pop(c)
	z := cmplx.Sqrt(x)
	zc.Complex.Push(c, z)
}

func SqrtDecimal(c zc.Calc) {
	var zero apd.Decimal
	d := vars.ForConf(c).DecMath
	x := zc.Decimal.Pop(c)
	if x.Cmp(&zero) < 0 {
		c.Raise(zc.ErrInvalidArg("%v < 0", x.String()))
		return
	}

	_, err := d.Sqrt(x, x)
	if err != nil {
		c.Raise(err)
		return
	}
	zc.Decimal.Push(c, x)
}

// ----------------------------------------------------------------------------
func SubBigInt(c zc.Calc) {
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	x.Sub(x, y)
	zc.BigInt.Push(c, x)
	zc.BigInt.Recycle(y)
}

func SubBigFloat(c zc.Calc) {
	y := zc.BigFloat.Pop(c)
	x := zc.BigFloat.Pop(c)
	x.Sub(x, y)
	zc.BigFloat.Push(c, x)
	zc.BigFloat.Recycle(y)
}

func SubComplex(c zc.Calc) {
	y := zc.Complex.Pop(c)
	x := zc.Complex.Pop(c)
	z := x - y
	zc.Complex.Push(c, z)
}

func SubDecimal(c zc.Calc) {
	d := vars.ForConf(c).DecMath
	y := zc.Decimal.Pop(c)
	x := zc.Decimal.Pop(c)
	defer zc.Decimal.Recycle(y)

	_, err := d.Sub(x, x, y)
	if err != nil {
		c.Raise(err)
		return
	}
	zc.Decimal.Push(c, x)
}

func SubRat(c zc.Calc) {
	y := zc.Rat.Pop(c)
	x := zc.Rat.Pop(c)
	x.Sub(x, y)
	zc.Rat.Push(c, x)
	zc.Rat.Recycle(y)
}

// ----------------------------------------------------------------------------
func QuoBigInt(c zc.Calc) {
	var zero big.Int
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		c.Raise(zc.ErrDivisionByZero)
		return
	}
	x.Quo(x, y)
	zc.BigInt.Push(c, x)
}

// ----------------------------------------------------------------------------
func QuoRemBigInt(c zc.Calc) {
	var zero big.Int
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	r := zc.BigInt.New()
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		c.Raise(zc.ErrDivisionByZero)
		return
	}
	x.QuoRem(x, y, r)
	zc.BigInt.Push(c, x)
	c.SetLabel("quo")
	zc.BigInt.Push(c, r)
	c.SetLabel("rem")
}
