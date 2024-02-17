package types

import (
	"fmt"
	"strings"

	"github.com/blackchip-org/dms"
)

type DMS struct {
	deg Decimal
	min Decimal
	sec Decimal
}

var (
	d60   = NewDecimalFromInt(60)
	d3600 = NewDecimalFromInt(3600)
)

func NewDMS(f dms.Fields) (DMS, error) {
	if f.Deg == "" {
		f.Deg = "0"
	}
	deg, err := NewDecimalFromString(f.Deg)
	if err != nil {
		return DMS{}, fmt.Errorf("invalid degrees: %v", f.Deg)
	}
	if f.Min == "" {
		f.Min = "0"
	}
	min, err := NewDecimalFromString(f.Min)
	if err != nil {
		return DMS{}, fmt.Errorf("invalid minutes: %v", f.Min)
	}
	if f.Sec == "" {
		f.Sec = "0"
	}
	sec, err := NewDecimalFromString(f.Sec)
	if err != nil {
		return DMS{}, fmt.Errorf("invalid seconds: %v", f.Sec)
	}

	sign := NewDecimalFromInt(int64(dms.Sign(f.Hemi)))
	deg, min, sec = deg.Abs(), min.Abs(), sec.Abs()

	// Normalize floats
	ideg := deg.Int()
	if !deg.Sub(ideg).IsZero() {
		fdeg := deg.Sub(ideg)
		deg = ideg
		min = min.Add(fdeg.Mul(d60))
	}
	imin := min.Int()
	if !min.Sub(imin).IsZero() {
		fmin := min.Sub(imin)
		min = imin
		sec = sec.Add(fmin.Mul(d60))
	}

	return DMS{}.Add(DMS{
		deg: deg.Mul(sign),
		min: min.Mul(sign),
		sec: sec.Mul(sign),
	}), nil
}

func MustNewDMS(f dms.Fields) DMS {
	d, err := NewDMS(f)
	if err != nil {
		panic(err)
	}
	return d
}

func (d DMS) String() string {
	return fmt.Sprintf("(%v,%v,%v)", d.deg, d.min.Abs(), d.sec.Abs())
}

func (d DMS) Add(d2 DMS) DMS {
	var carry Decimal

	sec := d.sec.Add(d2.sec)
	carry = sec.Div(d60).Int()
	d.sec = sec.Mod(d60)

	min := d.min.Add(d2.min).Add(carry)
	carry = min.Div(d60).Int()
	d.min = min.Mod(d60)

	d.deg = d.deg.Add(d2.deg).Add(carry)
	return d
}

func (d DMS) Degrees() Decimal {
	return d.deg.Add(d.min.Div(d60)).Add(d.sec.Div(d3600))
}

func (d DMS) Minutes() Decimal {
	return d.deg.Mul(d60).Add(d.min).Add(d.sec.Div(d60))
}

func (d DMS) Seconds() Decimal {
	return d.deg.Mul(d3600).Add(d.min.Mul(d60)).Add(d.sec)
}

func (d DMS) DMS() (deg, min, sec Decimal) {
	deg = d.deg
	min = d.min.Abs()
	sec = d.sec.Abs()
	return
}

func FormatDMS(d DMS, to string, places int, axis dms.Axis) string {
	deg, min, sec := d.DMS()
	sign := deg.Sign()
	if sign >= 0 {
		sign = 1
	}

	var buf strings.Builder
	if axis != dms.NoAxis {
		deg = deg.Abs()
	}

	func() {
		if to == dms.DegType {
			degs := deg.Add(min.Div(d60)).Add(sec.Div(d3600))
			buf.WriteString(degs.StringRound(places))
			buf.WriteRune('°')
			return
		}
		buf.WriteString(deg.String())
		buf.WriteString("° ")
		if to == dms.MinType {
			mins := min.Add(sec.Div(d60))
			buf.WriteString(mins.StringRound(places))
			buf.WriteRune('′')
			return
		}
		buf.WriteString(min.String())
		buf.WriteString("′ ")
		buf.WriteString(sec.StringRound(places))
		buf.WriteRune('″')
	}()
	if axis != dms.NoAxis {
		hemi := dms.Hemi(axis, sign)
		buf.WriteRune(' ')
		buf.WriteString(hemi)
	}
	return buf.String()
}
