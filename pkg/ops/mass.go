package ops

/*
oper	g-oz
func    - p0:Decimal -- Decimal
macro   1 28.349523125 div mul /oz anno
title   Grams to ounces

desc
Convert *p0* in grams to ounces.
end

example
2834.95 g-oz 2 round -- 100
end
*/

/*
oper	kg-lb
func	- p0:Decimal -- Decimal
macro	1 0.45359237 div mul /lb anno
title	Kilograms to pounds

desc
Convert *p0* in kilograms to pounds.
end

example
45.36 kg-lb 2 round -- 100
end
*/

/*
oper	lb-kg
func	- p0:Decimal -- Decimal
macro	0.45359237 mul /kg anno
title	Pounds to kilograms

desc
Convert *p0* in pounds to kilograms.
end

example
100 lb-kg 2 round -- 45.36
end
*/

/*
oper	oz-g
func    - p0:Decimal -- Decimal
macro   28.349523125 mul /g anno
title   Ounces to grams

desc
Convert *p0* in ounces to grams.
end

example
100 oz-g 2 round -- 2834.95
end
*/
