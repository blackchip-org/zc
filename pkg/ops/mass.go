package ops

/*
oper	g-kg
func 	- p0:Decimal -- Decimal
macro	1000 div /kg anno
title	Grams to kilograms

desc
Convert *p0* in grams to kilograms.
end

example
52124 g-kg -- 52.124 # kg
end
*/

/*
oper	g-oz
func    - p0:Decimal -- Decimal
macro   1 28.349523125 div mul /oz anno
title   Grams to ounces

desc
Convert *p0* in grams to ounces.
end

example
2834.95 g-oz 2 round -- 100 # oz
end
*/

/*
oper	g-ozt
func	- p0:Decimal -- Decimal
macro	1 31.1034768 div mul [oz t] anno
title	Grams to troy ounces

desc
Convert *p0* in grams to troy ounces.
end

example
326.59 g-ozt 2 round -- 10.5 # oz t
end
*/

/*
oper	kg-g
func 	- p0:Decimal -- Decimal
macro	1000 mul /kg anno
title	Kilograms to grams

desc
Convert *p0* in kilograms to grams.
end

example
52.124 kg-g -- 52124 # kg
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
45.36 kg-lb 2 round -- 100 # lb
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
100 lb-kg 2 round -- 45.36 # kg
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
100 oz-g 2 round -- 2834.95 # g
end
*/

/*
oper	ozt-g
func	- p0:Decimal -- Decimal
macro	31.1034768 mul /g anno
title	Troy ounces to grams

desc
Convert *p0* in troy ounces to grams.
end

example
10.5 ozt-g 2 round -- 326.59 # g
end
*/
