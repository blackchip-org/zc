package types

import (
	"fmt"
	"unicode"

	"github.com/blackchip-org/zc/pkg/scanner"
)

var (
	d60   = NewDecimalFromInt(60)
	d3600 = NewDecimalFromInt(3600)
)

type DMS struct {
	v Decimal
}

func NewDMS(v Decimal) DMS {
	return DMS{v: v}
}

func (d DMS) Degrees() Decimal {
	return d.v
}

func (d DMS) Minutes() Decimal {
	return d.v.Mul(d60)
}

func (d DMS) Seconds() Decimal {
	return d.v.Mul(d3600)
}

func (d DMS) FormatDMS(places int32) string {
	deg := d.v.Truncate(0)
	aDeg := deg.Abs()

	min := d.Minutes().Abs().Sub(aDeg.Mul(d60)).Truncate(0)
	//min := math.Trunc(math.Abs(a.Minutes()) - (aDeg * 60))

	sec := d.Seconds().Abs().Sub(aDeg.Mul(d3600)).Sub(min.Mul(d60))
	// fSec := math.Abs(a.Seconds()) - (aDeg * 3600) - (min * 60)
	// dSec := decimal.NewFromFloat(fSec)
	if places >= 0 {
		sec = sec.Round(places)
	}
	return fmt.Sprintf("%v° %v′ %v″", deg, min, sec)
}

func (d DMS) FormatDM(places int32) string {
	deg := d.v.Truncate(0)
	aDeg := deg.Abs()

	min := d.Minutes().Abs().Sub(aDeg.Mul(d60))
	if places >= 0 {
		min = min.Round(places)
	}
	return fmt.Sprintf("%v° %v′", deg, min)
}

func (d DMS) String() string {
	return d.FormatDMS(-1)
}

func ParseDMS(str string) (DMS, bool) {
	var s scanner.Scanner
	s.SetString(str)
	s.TrimText = true

	sDeg := s.ScanUntil(scanner.Runes('°', 'd'))
	s.Next()
	sMin := s.ScanUntil(scanner.Runes('m', '\'', '′'))
	s.Next()
	sSec := s.ScanUntil(scanner.Runes('s', '"', '″'))
	s.Next()

	v := DecimalZero
	var deg, min, sec Decimal
	var err error

	if sDeg != "" {
		deg, err = NewDecimalFromString(sDeg)
		if err != nil {
			return DMS{}, false
		}
		v = v.Add(deg)
	}
	if sMin != "" {
		min, err = NewDecimalFromString(sMin)
		if err != nil {
			return DMS{}, false
		}
		v = v.Add(min.Div(d60))
	}
	if sSec != "" {
		sec, err = NewDecimalFromString(sSec)
		if err != nil {
			return DMS{}, false
		}
		v = v.Add(sec.Div(d3600))
	}

	if sDeg != "" && (sMin != "" || sSec != "") {
		if !deg.IsInteger() {
			return DMS{}, false
		}
	}
	if sMin != "" && sSec != "" {
		if !min.IsInteger() {
			return DMS{}, false
		}
	}

	s.ScanWhile(unicode.IsSpace)
	if !s.End() {
		return DMS{}, false
	}
	return DMS{v: v}, true
}
