package types

import (
	"fmt"
	"strings"

	"github.com/blackchip-org/dms"
	"github.com/blackchip-org/zc/v6/pkg/calc"
	"github.com/cockroachdb/apd/v3"
)

type DMS struct {
	ctx *apd.Context
	deg apd.Decimal
	min apd.Decimal
	sec apd.Decimal
}

func NewDMS(ctx *apd.Context, deg, min, sec *apd.Decimal) DMS {
	dc := calc.NewDecimal()
	dc.Ctx = ctx

	sign := deg.Sign()
	if sign == 0 {
		sign = 1
	}

	dc.Push(deg).Abs().Save("deg")
	dc.Push(min).Abs().Save("min")
	dc.Push(sec).Abs().Save("sec")

	// Normalize values
	dc.Load("deg")
	dc.Dup().Trunc().Dup().Save("ideg")
	dc.Sub()

	dc.Dup().PushInt(0)
	if dc.Cmp() != 0 {
		dc.Load("ideg").Save("deg")
		dc.PushInt(60).Mul()
		dc.Load("min").Add().Save("min")
	} else {
		dc.Drop()
	}

	dc.Load("min")
	dc.Dup().Trunc().Dup().Save("imin")
	dc.Sub()

	dc.Dup().PushInt(0)
	if dc.Cmp() != 0 {
		dc.Load("imin").Save("min")
		dc.PushInt(60).Mul()
		dc.Load("sec").Add().Save("sec")
	} else {
		dc.Drop()
	}

	ndeg := dc.Load("deg").PushInt(sign).Mul().Reduce().Pop()
	nmin := dc.Load("min").PushInt(sign).Mul().Reduce().Pop()
	nsec := dc.Load("sec").PushInt(sign).Mul().Reduce().Pop()

	return DMS{ctx: ctx}.Add(DMS{ctx: ctx, deg: *ndeg, min: *nmin, sec: *nsec})
}

func NewDMSFromFields(c *apd.Context, f dms.Fields) (DMS, error) {
	if f.Deg == "" {
		f.Deg = "0"
	}
	deg, _, err := apd.NewFromString(f.Deg)
	if err != nil {
		return DMS{}, fmt.Errorf("invalid degrees: %v", f.Deg)
	}
	if f.Min == "" {
		f.Min = "0"
	}
	min, _, err := apd.NewFromString(f.Min)
	if err != nil {
		return DMS{}, fmt.Errorf("invalid minutes: %v", f.Min)
	}
	if f.Sec == "" {
		f.Sec = "0"
	}
	sec, _, err := apd.NewFromString(f.Sec)
	if err != nil {
		return DMS{}, fmt.Errorf("invalid seconds: %v", f.Sec)
	}

	if dms.Sign(f.Hemi) < 0 {
		deg.Neg(deg)
	}
	return NewDMS(c, deg, min, sec), nil
}

func NewDMSFromFloat(c *apd.Context, d, m, s float64) DMS {
	var deg, min, sec apd.Decimal
	if _, err := deg.SetFloat64(d); err != nil {
		panic(err)
	}
	if _, err := min.SetFloat64(m); err != nil {
		panic(err)
	}
	if _, err := sec.SetFloat64(s); err != nil {
		panic(err)
	}
	return NewDMS(c, &deg, &min, &sec)
}

func (d DMS) String() string {
	var amin, asec apd.Decimal
	amin.Abs(&d.min)
	asec.Abs(&d.sec)
	return fmt.Sprintf("(%v,%v,%v)", d.deg, amin, asec)
}

func (d DMS) Add(d2 DMS) DMS {
	dc := calc.NewDecimal()
	dc.Ctx = d.ctx

	dc.Push(&d.sec, &d2.sec).Add()
	dc.Dup().PushInt(60).Quo().Trunc().Save("carry")
	d.sec = *dc.PushInt(60).Rem().Pop()

	dc.Load("carry").Push(&d.min, &d2.min).Add().Add()
	dc.Dup().PushInt(60).Quo().Trunc().Save("carry")
	d.min = *dc.PushInt(60).Rem().Pop()

	dc.Load("carry").Push(&d.deg, &d2.deg).Add().Add()
	d.deg = *dc.Pop()

	return d
}

func (d DMS) DMS() (*apd.Decimal, *apd.Decimal, *apd.Decimal) {
	var deg, min, sec apd.Decimal

	deg.Set(&d.deg)
	min.Abs(&d.min)
	sec.Abs(&d.sec)
	return &deg, &min, &sec
}

func (d DMS) Degrees() *apd.Decimal {
	dc := calc.NewDecimal()
	dc.Ctx = d.ctx

	dc.Push(&d.deg)
	dc.Push(&d.min).PushInt(60).Quo()
	dc.Push(&d.sec).PushInt(3600).Quo()
	dc.Add().Add().Reduce()
	return dc.Pop()
}

func (d DMS) Minutes() *apd.Decimal {
	dc := calc.NewDecimal()
	dc.Ctx = d.ctx

	dc.Push(&d.deg).PushInt(60).Mul()
	dc.Push(&d.min)
	dc.Push(&d.sec).PushInt(60).Quo()
	dc.Add().Add().Reduce()
	return dc.Pop()
}

func (d DMS) Seconds() *apd.Decimal {
	dc := calc.NewDecimal()
	dc.Ctx = d.ctx

	dc.Push(&d.deg).PushInt(3600).Mul()
	dc.Push(&d.min).PushInt(60).Mul()
	dc.Push(&d.sec)
	dc.Add().Add().Reduce()
	return dc.Pop()
}

func FormatDMS(d DMS, to dms.Unit, places int) string {
	deg, min, sec := d.DMS()
	var buf strings.Builder

	dc := calc.NewDecimal()
	dc.Ctx = d.ctx

	func() {
		if to == dms.DegUnit {
			dc.Push(d.Degrees())
			dc.Reduce()
			if places >= 0 {
				dc.Quantize(-int32(places))
			}
			buf.WriteString(dc.Pop().Text('f'))
			buf.WriteRune('°')
			return
		}
		buf.WriteString(deg.String())
		buf.WriteString("° ")
		if to == dms.MinUnit {
			dc.Push(min)
			dc.Push(sec).PushInt(60).Quo()
			dc.Add()
			dc.Reduce()
			if places >= 0 {
				dc.Quantize(-int32(places))
			}
			buf.WriteString(dc.Pop().Text('f'))
			buf.WriteRune('′')
			return
		}
		buf.WriteString(min.String())
		buf.WriteString("′ ")

		dc.Push(sec)
		dc.Reduce()
		if places >= 0 {
			dc.Quantize(-int32(places))
		}
		buf.WriteString(dc.Pop().Text('f'))
		buf.WriteRune('″')
	}()
	return buf.String()
}
