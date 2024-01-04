package ops

/*
oper 	c-f
func	- p0:Decimal -- Decimal
macro	9 5 div mul 32 add /°F anno
title	Celsius to Fahrenheit

desc
Converts the temperature *p0* in Celsius to Fahrenheit.
end

example
20 c-f -- 68 # °F
end
*/

/*
oper	c-k
func	- p0:Decimal -- Decimal
macro	273.15 add /K anno
title	Celsius to Kelvin

desc
Converts the temperature *p0* in Celsius to Kelvin.
end

example
100 c-k -- 373.15 # K
end
*/

/*
oper	f-c
func	- p0:Decimal -- Decimal
macro	32 sub 5 9 div mul /°C anno
title	Fahrenheit to Celsius

desc
Converts the temperature *p0* in Fahrenheit to Celsius.
end

example
68 f-c 2 round -- 20
end
*/

/*
oper 	k-c
func	- p0:Decimal -- Decimal
macro	273.15 sub /°C anno
title	Kelvin to Celsius

desc
Converts the temperature *p0* in Kelvin to Celsius.
end

example
373.15 k-c -- 100 # °C
end
*/
