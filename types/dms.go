package types

// type DMS struct {
// 	deg *apd.Decimal
// 	min *apd.Decimal
// 	sec *apd.Decimal
// }

// var (
// 	d60   = apd.New(60, 0)
// 	d3600 = apd.New(3600, 0)
// )

// func NewDMS(deg, min, sec *apd.Decimal) DMS {
// 	sign := apd.New(int64(deg.Sign()), 0)
// 	if sign.IsZero() {
// 		sign = apd.New(1, 0)
// 	}
// 	deg, min, sec = deg.Abs(), min.Abs(), sec.Abs()

// 	// Normalize floats
// 	ideg := deg.Int()
// 	if !deg.Sub(ideg).IsZero() {
// 		fdeg := deg.Sub(ideg)
// 		deg = ideg
// 		min = min.Add(fdeg.Mul(d60))
// 	}
// 	imin := min.Int()
// 	if !min.Sub(imin).IsZero() {
// 		fmin := min.Sub(imin)
// 		min = imin
// 		sec = sec.Add(fmin.Mul(d60))
// 	}

// 	return DMS{}.Add(DMS{
// 		deg: deg.Mul(sign),
// 		min: min.Mul(sign),
// 		sec: sec.Mul(sign),
// 	})
// }
