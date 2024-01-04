package ops

/*
oper	iec.convert
func	- p0:Val u0:BigInt u1:BigInt -- Val
macro   div mul
title	Converts between IEC units

desc
Convert *p0* in *u0* IEC units to *u1* IEC units.
end

example
145 -- 145
iec.mebi iec.kibi iec.convert -- 148480
end
*/

/*
oper	iec.unit
func    - -- BigInt
macro 	1
title   IEC unit value, 2^0

desc
IEC unit value, 2^0
end
*/

//tab iec.kibi	-- 2 10 pow	-- IEC prefix Ki, 2^10
//tab iec.mebi  -- 2 20 pow -- IEC prefix Mi, 2^20
//tab iec.gibi  -- 2 30 pow -- IEC prefix Gi, 2^30
//tab iec.tebi  -- 2 40 pow -- IEC prefix Ti, 2^40
//tab iec.pebi  -- 2 50 pow -- IEC prefix Pi, 2^50
//tab iec.exbi  -- 2 60 pow -- IEC prefix EI, 2^60
