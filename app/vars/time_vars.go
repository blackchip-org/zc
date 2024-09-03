package vars

import (
	"time"

	"github.com/blackchip-org/zc/v6/pkg/coll"
	"github.com/blackchip-org/zc/v6/pkg/ptime/locale"
)

const TimeID = "time"

type Time struct {
	Locale         *locale.Locale
	Zone           *time.Location
	ZoneName       string
	DateLayout     string
	TimeLayout     string
	DateTimeLayout string
	Now            func() time.Time
}

const (
	DefaultDateLayout     = "[weekday/abbr] [month/abbr] [day] [year]"
	DefaultTimeLayout     = "[hour/12]:[minute]:[second][period/alt] [offset-zone]"
	DefaultDateTimeLayout = DefaultDateLayout + " " + DefaultTimeLayout
)

func ForTime(state coll.State) *Time {
	v, ok := state.Var(TimeID)
	if !ok {
		loc := time.Now().Location()
		tz, _ := time.Now().Zone()

		v = &Time{
			Locale:         locale.EnUS,
			Zone:           loc,
			ZoneName:       tz,
			DateLayout:     DefaultDateLayout,
			TimeLayout:     DefaultTimeLayout,
			DateTimeLayout: DefaultDateTimeLayout,
			Now: func() time.Time {
				return time.Now().In(loc)
			},
		}
		state.NewVar(TimeID, v)
	}
	return v.(*Time)
}
