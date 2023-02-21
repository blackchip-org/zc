# time

<!-- eval: use time -->
<!-- eval: 'Jan 2 2006 15:04:05 -0700 MST' travel -->
<!-- eval: 'MST' local -->

Date, time, and duration operations

<!-- index -->

| Operation                               | Alias   | Description
|-----------------------------------------|---------|------------
| [add](#add-duration)                    | `add-d` | Time after a duration
| [date](#date)                           |         | Formats a date with the default layout
| [date-layout](#date-layout)             |         | Default date layout, set
| [date-layout=](#date-layout=)           |         | Default date layout, get
| [day-year](#date-year)                  |         | Day of year for a given date
| [date-time](#date-time)                 | `dt`    | Formats a datetime with the default layout
| [date-time-layout](#date-time-layout)   |         | Default datetime layout, set
| [date-time-layout=](#date-time-layout=) |         | Default datetime layout, get
| [local](#local)                         |         | Local time zone, set
| [local=](#local=)                       |         | Local time zone, get
| [now](#now)                             |         | Current date and time
| [subtract-time](#subtact-time)          | `sub-t` | Duration between two times
| [time](#time)                           |         | Formats a time with the default layout
| [time-layout](#time-layout)             |         | Default time layout, set
| [time-layout](#time-layout=)            |         | Default time layout, get
| [time-zone](#time-zone)                 | `tz`    | Convert time to a given time zone
| [travel](#travel)                       |         | Override the value returned by `now`
| [travel-end](#travel-end)               |         | Cancel override of the value returned by now


## add-duration

Adds the duration `d` to time `t`.

    ( t:DateTime d:Duration -- add:DateTime); or
    ( d:Duration t:DateTime -- add:DateTime)

Aliases: `add-d`

Example:

<!-- test: add-duration -->

| Input      | Stack
|------------|------------------
| `3:30pm`   | `3:30pm`
| `2h add-d` | `Mon Jan 2 2006 5:30:00pm -0700 MST`

## date

Formats a the date `d` with the default layout

    ( d:DateTime -- date:Date )

Example:

<!-- test: date -->

| Input                      | Stack
|----------------------------|------------------
| `'2006-01-02T15:04:05 UTC` | `2006-01-02T15:04:05 UTC`
| `date`                     | `Mon Jan 2 2006`

## date-layout

Sets the default date layout.

    ( layout:String -- )

Example:

<!-- test: date-layout -->

| Input                      | Stack
|----------------------------|------------------
| `'[year]-[day/year]`       | `[year]-[day/year]`
| `date-layout`              |
| `'Mar 15 2006' date`       | `2006-074`


## day-year

Day of year for a given date.

    ( dt:DateTime -- Int )

Example:

<!-- test: day-year -->

| Input                      | Stack
|----------------------------|------------------
| `2006-03-15`               | `2006-03-15`
| `day-year`                 | `74`


## date-layout=

Gets the default layout.

    ( -- layout:String )

Example:

<!-- test: date-layout= -->

| Input                      | Stack
|----------------------------|------------------
| `date-layout=`             | `[weekday/abbr] [month/abbr] [day] [year]`

## date-time

Formats a the date `d` with the default layout

    ( d:DateTime -- date:Date )

Aliases: `dt`

Example:

<!-- test: datetime -->

| Input                      | Stack
|----------------------------|------------------
| `'2006-01-02T15:04:05 UTC` | `2006-01-02T15:04:05 UTC`
| `dt`                       | `Mon Jan 2 2006 3:04:05pm UTC`


## date-time-layout

Sets the default datetime layout.

    ( layout:String -- )

Example:

<!-- test: date-time-layout -->

| Input                               | Stack
|-------------------------------------|------------------
| `'[year]-[day/year] [hour]:[minute]`| `[year]-[day/year] [hour]:[minute]`
| `date-time-layout`                  |
| `'Mar 15 2006 3:04pm' date-time`    | `2006-074 15:04`


## date-time-layout=

Gets the default datetime layout.

    ( -- layout:String )

Example:

<!-- test: date-time-layout= -->

| Input                      | Stack
|----------------------------|------------------
| `date-time-layout=`        | `[weekday/abbr] [month/abbr] [day] [year] [hour/12]:[minute]:[second][period/alt] [offset-zone]`


## local

Sets the local time zone.

    ( zone:String -- )

Example:

<!-- test: local -->

| Input                           | Stack
|---------------------------------|------------------
| `now time`                      | `3:04:05pm -0700 MST`
| `clear`                         |
| `'est' local now time`          | `5:04:05pm -0500 EST`
| `clear`                         |
| `'Asia/Jakarta' local now time` | `5:04:05am +0700`


## local=

Gets the local time zone.

    ( -- zone:String )

Example:

<!-- test: local -->

| Input                  | Stack
|------------------------|------------------
| `local=`               | `MST`


## now

The current date and time. If `travel` has been called, that date and
time will be returned instead.

    ( -- now:DateTime )

<!-- test: now -->

| Input                  | Stack
|------------------------|------------------
| `now`                  | `Mon Jan 2 2006 3:04:05pm -0700 MST`


## subtract-time

The duration in time by subtracting `b` from `a`.

    ( a:DateTime b:DateTime -- sub-t:Duration )

Aliases: `sub-t`

<!-- test: subtract-time -->

| Input                        | Stack
|------------------------------|------------------
| `'Jan 2 2006 10:00am`        | `Jan 2 2006 10:00am`
| `'Dec 31 2005 5:30pm' sub-t` | `40h30m`


## time

Formats a time with the default layout

    ( t:DateTime -- time:Time )

Example:

<!-- test: date -->

| Input                      | Stack
|----------------------------|------------------
| `'2006-01-02T15:04:05 UTC` | `2006-01-02T15:04:05 UTC`
| `time`                     | `3:04:05pm UTC`


## time-layout

Sets the default time layout.

    ( layout:String -- )

Example:

<!-- test: time-layout -->

| Input                      | Stack
|----------------------------|------------------
| `'[hour/24]:[minute]`      | `[hour/24]:[minute]`
| `time-layout`              |
| `'3:15:45pm' time`         | `15:15`


## time-layout=

Gets the default time layout.

    ( -- layout:String )

Example:

<!-- test: time-layout= -->

| Input                      | Stack
|----------------------------|------------------
| `time-layout=`             | `[hour/12]:[minute]:[second][period/alt] [offset-zone]`


## time-zone

Convert time to a given time zone

    ( dt:DateTime zone:String -- dt:DateTime )

Aliases: `tz`

Example:

<!-- test: time-zone -->

| Input                      | Stack
|----------------------------|------------------
| `now`                      | `Mon Jan 2 2006 3:04:05pm -0700 MST`
| `'PST' time-zone`          | `Mon Jan 2 2006 2:04:05pm -0800 PST`
| `'Asia/Jakarta' time-zone` | `Tue Jan 3 2006 5:04:05am +0700 WIB`


## travel

Override the value returned by `now`. Useful for to mock current time while
testing.

    ( dt:DateTime -- )

Example:

<!-- test: travel -->

| Input                       | Stack
|-----------------------------|------------------
| `'Nov 5 1955 01:22' travel` |
| `now`                       | `Sat Nov 5 1955 1:22:00am -0700 MST`

## travel-end

Cancel override of the value returned by now

