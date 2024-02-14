package types

import (
	"fmt"

	"github.com/blackchip-org/dms"
)

type DMS struct {
	deg Decimal
	min Decimal
	sec Decimal
}

var (
	dec60   = NewDecimalFromInt(60)
	dec3600 = NewDecimalFromInt(3600)
)

func NewDMS(f dms.Fields) (DMS, error) {
	deg, err := NewDecimalFromString(f.Deg)
	if err != nil {
		return DMS{}, fmt.Errorf("invalid degrees: %v", f.Deg)
	}
	min, err := NewDecimalFromString(f.Min)
	if err != nil {
		return DMS{}, fmt.Errorf("invalid minutes: %v", f.Min)
	}
	sec, err := NewDecimalFromString(f.Sec)
	if err != nil {
		return DMS{}, fmt.Errorf("invalid seconds: %v", f.Sec)
	}

	sign := NewDecimalFromInt(int64(f.Sign()))
	deg, min, sec = deg.Abs(), min.Abs(), sec.Abs()

	// Normalize floats
	ideg := deg.Int()
	if !deg.Sub(ideg).IsZero() {
		fdeg := deg.Sub(ideg)
		deg = ideg
		min = min.Add(fdeg.Mul(dec60))
	}
	imin := min.Int()
	if !min.Sub(imin).IsZero() {
		fmin := min.Sub(imin)
		min = imin
		sec = sec.Add(fmin.Mul(dec60))
	}

	return DMS{}.Add(DMS{
		deg: deg.Mul(sign),
		min: min.Mul(sign),
		sec: sec.Mul(sign),
	}), nil
}

func (d DMS) Add(d2 DMS) DMS {
	var carry Decimal

	sec := d.sec.Add(d2.sec)
	carry = sec.Div(dec60).Int()
	d.sec = sec.Mod(dec60)

	min := d.min.Add(d2.min).Add(carry)
	carry = min.Div(dec60).Int()
	d.min = min.Mod(dec60)

	d.deg = d.deg.Add(d2.deg).Add(carry)
	return d
}

func (d DMS) Degrees() Decimal {
	return d.deg.Add(d.min.Div(dec60)).Add(d.sec.Div(dec3600))
}

func (d DMS) Minutes() Decimal {
	return d.deg.Mul(dec60).Add(d.min).Add(d.sec.Div(dec60))
}

func (d DMS) Seconds() Decimal {
	return d.deg.Mul(dec3600).Add(d.min.Mul(dec60)).Add(d.sec)
}

func (d DMS) DMS() (deg, min, sec Decimal) {
	deg = d.deg
	min = d.min.Abs()
	sec = d.sec.Abs()
	return
}
