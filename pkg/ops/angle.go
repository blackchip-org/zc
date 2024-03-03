package ops

import (
	"github.com/blackchip-org/dms"
	"github.com/blackchip-org/zc/v5/pkg/types"
	"github.com/blackchip-org/zc/v5/pkg/zc"
)

/*
oper	dec
func	DecDMS p0:DMS -- Float
title	DMS angle to decimal degrees

desc
Convert the DMS angle *p0* to decimal degrees.
end

example
10d7m24.24s dec -- 10.1234
end
*/
func DecDMS(c zc.Calc) {
	p0 := zc.PopDMS(c)
	r0 := p0.Degrees()
	zc.PushDecimal(c, r0)
}

/*
oper	deg-min
func	DM p0:DMS -- DMS
alias	dm
title	Angle in degrees and minutes

desc
Reformat the angle *p0* to degrees and minutes
example
10.1234 dm -- 10° 7.404000000000011′
end
*/
func DM(c zc.Calc) {
	p0 := zc.PopDMS(c)
	r0 := types.FormatDMS(p0, dms.MinUnit, -1)
	zc.PushString(c, r0)
}

/*
oper	deg-min-round
func	DMRound p0:DMS n:Int -- DMS
alias	dmr
title	Rounded angle in degrees and minutes

desc
Reformat the angle *p0* to degrees and minutes and round the minutes to
*n* places.
end

example
12.57611 3 dmr -- 12° 34.567′
end
*/
func DMRound(c zc.Calc) {
	places := zc.PopInt(c)
	p0 := zc.PopDMS(c)
	r0 := types.FormatDMS(p0, dms.MinUnit, places)
	zc.PushString(c, r0)
}

/*
oper	deg-min-sec
func	DMS p0:DMS -- DMS
alias	dms
title	Angle in degrees and minutes

desc
Reformat the angle *p0* to degrees, minutes, and seconds.
end

example
-76.856944 dms -- -76° 51′ 24.9984″
end
*/
func DMS(c zc.Calc) {
	p0 := zc.PopDMS(c)
	zc.PushDMS(c, p0)
}

/*
oper	dms?
func 	DMSIs p0:Str -- Bool
alias	dec-min-sec?
title	Checks value can be parsed as degree, minutes, seconds

desc
Returns `true` if the value *p0* can be parsed as an angle with degrees,
minutes, and seconds.
end

example
c [10° 30′ 45″] dms? -- true
c [10  30  45 ] dms? -- false
end
*/
func DMSIs(c zc.Calc) {
	p0 := zc.PopString(c)
	r0 := zc.DMS.Is(p0)
	zc.PushBool(c, r0)
}

/*
oper	deg-min-sec-round
func	DMSRound p0:DMS n:Int -- DMS
alias	dmsr
title	Rounded angle in degrees, minutes, seconds

desc
Reformat the angle *p0* to degrees, minutes, and seconds, and round the
seconds to *n* places.
end

example
-76.856944 0 dmsr -- -76° 51′ 25″
end
*/
func DMSRound(c zc.Calc) {
	places := zc.PopInt(c)
	p0 := zc.PopDMS(c)
	r0 := types.FormatDMS(p0, dms.SecUnit, places)
	zc.PushString(c, r0)
}

/*
oper	deg-rad
func 	- p0:DMS -- Decimal
macro	pi 180 div mul
title	Degrees to radians

desc
Converts the angle *p0* in degrees to radians
end

example
90 deg-rad -- 1.5707963267948966
end
*/

/*
oper	minutes
func	MinutesDMS p0:DMS -- Float
title	Angle in minutes

desc
Converts the angle *p0* to minutes.
end

example
10d30m minutes -- 630
end
*/
func MinutesDMS(c zc.Calc) {
	p0 := zc.PopDMS(c)
	r0 := p0.Minutes()
	zc.PushDecimal(c, r0)
}

/*
oper	rad-deg
func 	- p0:Decimal -- DMS
macro	180 pi div mul
title	Degrees to radians

desc
Converts the angle *p0* in radians to degrees
end

example
1.570796326794897 rad-deg 2 round -- 90
end
*/

/*
oper	seconds
func	SecondsDMS p0:DMS -- Float
title	Angle in seconds

desc
Converts the angle *p0* to seconds.
end

example
10d30m seconds -- 37800
end
*/
func SecondsDMS(c zc.Calc) {
	p0 := zc.PopDMS(c)
	r0 := p0.Seconds()
	zc.PushDecimal(c, r0)
}
