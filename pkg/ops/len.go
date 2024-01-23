package ops

/*
oper	ft-m
func	- p0:Decimal -- Decimal
macro	0.3048 mul /m anno
title	Feet to meters

desc
Convert *p0* in feet to meters.
end

example
1000 ft-m -- 304.8 # m
end
*/

/*
oper	ft-mi
func	- p0:Decimal -- Decimal
macro	5280 div /mi anno
title	Miles to feet

desc
Convert *p0* in feet to miles
end

example
2640 ft-mi -- 0.5 # mi
end
*/

/*
oper	ft-yd
func	- p0:Decimal -- Decimal
macro	3 div /yd anno
title	Feet to yards

desc
Convert *p0* in feet to yards.
end

example
300 ft-yd -- 100 # yd
end
*/

/*
oper	in-mm
func	- p0:Decimal -- Decimal
macro 	25.4 mul /mm anno
title 	Inches to millimeters

desc
Convert *p0* in inches to millimeters.
end

example
100 in-mm -- 2540 # mm
end
*/

/*
oper	km-m
func	- p0:Decimal -- Decimal
macro	1000 mul /m anno
title   Kilometers to meters

desc
Convert *p0* in kilometers to meters.
end

example
6378.137 km-m -- 6378137 # m
end
*/

/*
oper	km-mi
func	- p0:Decimal -- Decimal
macro	0.62137119 mul /mi anno
title	Kilometers to miles

desc
Convert *p0* in kilometers to miles.
end

example
100 km-mi 2 round -- 62.14 # mi
end
*/

/*
oper	km-nmi
func	- p0:Decimal -- Decimal
macro	0.539957 mul /nmi anno
title	Kilometers to nautical miles

desc
Convert *p0* in kilometers to nautical miles.
end

example
100 km-nmi 2 round -- 54 # nmi
end
*/

/*
oper	m-ft
func	- p0:Decimal -- Decimal
macro	0.3048 div /ft anno
title	Meters to feet

desc
Convert *p0* in meters to feet.
end

example
304.8 m-ft 2 round -- 1000 # ft
end
*/

/*
oper	m-km
func	- p0:Decimal -- Decimal
macro	1000 div /km anno
title   Meters to kilometers

desc
Convert *p0* in meters to kilometers.
end

example
earth-equatorial-radius -- 6378137 # m
m-km -- 6378.137 # km
end
*/

/*
oper	m-nmi
func	- p0:Decimal -- Decimal
macro	0.000539957 mul /nmi anno
title	Meters to nautical miles

desc
Convert *p0* in meters to nautical miles.
end

example
100,000 m-nmi 2 round -- 54 # nmi
end
*/

/*
oper	m-yd
func	- p0:Decimal -- Decimal
macro	0.9144 div /yd anno
title	Meters to yards

desc
Convert *p0* in meters to yards
end

example
91.44 m-yd 2 round -- 100 # yd
end
*/

/*
oper	mi-ft
func	- p0:Decimal -- Decimal
macro	5280 mul /ft anno
title	Miles to feet

desc
Convert *p0* in miles to feet
end

example
0.5 mi-ft -- 2640 # ft
end
*/

/*
oper	mi-km
func	- p0:Decimal -- Decimal
macro	1.609344 mul /km anno
title	Miles to kilometers

desc
Convert *p0* in miles to kilometers
end

example
100 mi-km 2 round -- 160.93 # km
end
*/

/*
oper	mi-nmi
func	- p0:Decimal -- Decimal
macro	0.868976 mul /nmi anno
title	Miles to nautical miles

desc
Convert *p0* in miles to nautical miles
end

example
100 mi-nmi 2 round -- 86.9 # nmi
end
*/

/*
oper	mm-in
func	- p0:Decimal -- Decimal
macro 	25.4 div /in anno
title 	Millimeters to inches

desc
Convert *p0* in millimeters to inches.
end

example
2540 mm-in 2 round -- 100 # in
end
*/

/*
oper	nmi-km
func	- p0:Decimal -- Decimal
macro	1.852 mul /km anno
title 	Nautical miles to kilometers

desc
Convert *p0* in nautical miles to kilometers
end

example
100 nmi-km 2 round -- 185.2 # km
end
*/

/*
oper	nmi-m
func	- p0:Decimal -- Decimal
macro	1852 mul /m anno
title	Nautical miles to meters

desc
Convert *p0* in nautical miles to meters
end

example
100 nmi-m -- 185200 # m
end
*/

/*
oper	nmi-mi
func	- p0:Decimal -- Decimal
macro	1.15078 mul /mi anno
title	Nautical miles to miles

desc
Convert *p0* in nautical miles to miles
end

example
100 nmi-mi 2 round -- 115.08 # mi
end
*/

/*
oper	yd-ft
func	- p0:Decimal -- Decimal
macro	3 mul /ft anno
title	Yards to feet

desc
Convert *p0* in yards to feet
end

example
100 yd-ft -- 300 # ft
end
*/

/*
oper	yd-m
func	- p0:Decimal -- Decimal
macro	0.9144 mul /m anno
title	Yards to meters

desc
Convert *p0* in yards to meters
end

example
100 yd-m -- 91.44 # m
end
*/
