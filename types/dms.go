package types

import (
	"github.com/cockroachdb/apd/v3"
)

type DMS struct {
	deg *apd.Decimal
	min *apd.Decimal
	sec *apd.Decimal
}

// func NewDMS(deg, min, sec *apd.Decimal) DMS {
// 	c := calc.NewDecimal()

// 	c.Push(deg)
// 	c.Abs()
// 	c.Save("deg")

// 	c.Push(min)
// 	c.Abs()
// 	c.Save("min")

// 	c.Push(sec)
// 	c.Abs()
// 	c.Save("sec")

// }
