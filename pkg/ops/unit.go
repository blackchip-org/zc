package ops

// Length
const (
	KmToMi  = "0.62137119  mul"
	KmToNmi = "0.539957    mul"
	MToNmi  = "0.000539957 mul"
	MiToKm  = "1.609344    mul"
	MiToNmi = "0.868976    mul"
	NmiToKm = "1.852       mul"
	NmiToM  = "1852        mul"
	NmiToMi = "1.15078     mul"
)

// SI Prefixes
const (
	Quetta = "1e30"
	Ronna  = "1e27"
	Yotta  = "1e24"
	Zetta  = "1e21"
	Exa    = "1e18"
	Peta   = "1e15"
	Tera   = "1e12"
	Giga   = "1e09"
	Mega   = "1e06"
	Kilo   = "1e03"
	Hecto  = "1e02"
	Deca   = "1e01"
	Deci   = "1e-01"
	Centi  = "1e-02"
	Milli  = "1e-03"
	Micro  = "1e-06"
	Nano   = "1e-09"
	Pico   = "1e-12"
	Femto  = "1e-15"
	Atto   = "1e-18"
	Zepto  = "1e-21"
	Yocto  = "1e-24"
	Ronto  = "1e-27"
	Quecto = "1e-30"
)

// Temperature
const (
	CToF = "9 5 div mul 32 add"
	CToK = "273.15 add"
	FToC = "32 sub 5 9 div mul"
	KToC = "273.15 sub"
)
