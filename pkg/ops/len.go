package ops

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
