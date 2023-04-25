# time

<!-- eval: 'Jan 2 2006 15:04:05 -0700 MST' now-set -->
<!-- eval: 'MST' local-zone -->

Date, time, and duration operations

<!-- index -->

| Operation                    | Alias    | Description
|------------------------------|----------|------------
| [add](#add)                  | `a`, `+` | Time after a duration
| [date](#date)                |          | Formats a date with the default layout
| [day-year](#date-year)       |          | Day of year for a given date
| [date-time](#date-time)      | `dt`     | Formats a datetime with the default layout
| [hours](#hours)              |          | Hours in a duration
| [local-zone](#local-zone)    |          | Local time zone, set
| [local-zone=](#local=)       |          | Local time zone, get
| [minutes](#minutes)          |          | Minutes in a duration
| [now](#now)                  |          | Current date and time
| [now-set](#now-set)          |          | Override the value returned by `now`
| [now-restore](#now-restore)  |          | Cancel override of the value
| [seconds](#seconds)          |          | Seconds in a duration
| [sub](#sub)                  | `s`, `-` | Duration between two times
| [time](#time)                |          | Formats a time with the default layout
| [time-zone](#time-zone)      | `tz`     | Convert time to a given time zone


## add

Adds the duration *d* to time *t*.

Aliases: `a`, `+`

    ( t:DateTime d:Duration -- DateTime); or
    ( d:Duration t:DateTime -- DateTime)

Example:

<!-- test: add-duration -->

| Input      | Stack
|------------|------------------
| `3:30pm`   | `3:30pm`
| `2h add`   | `Mon Jan 2 2006 5:30:00pm -0700 MST`

Or, adds the duration of *p0* to *p1*.

    ( p0:Duration p1:Duration -- Duration )

Example:

<!-- test: add-duration-2 -->

| Input      | Stack
|------------|------------------
| `2h30m`    | `2h30m`
| `45m add`  | `3h15m`

## date

Formats a date/time to a common date layout. Time information, if any, is
removed.

    ( DateTime -- date:Date )

Example:

<!-- test: date -->

| Input                      | Stack
|----------------------------|------------------
| `'2006-01-02T15:04:05 UTC` | `2006-01-02T15:04:05 UTC`
| `date`                     | `Mon Jan 2 2006`

## datetime

Formats a date/time to a common layout.

    ( DateTime -- DateTime )

Aliases: `dt`

Example:

<!-- test: datetime -->

| Input                      | Stack
|----------------------------|------------------
| `'2006-01-02T15:04:05 UTC` | `2006-01-02T15:04:05 UTC`
| `dt`                       | `Mon Jan 2 2006 3:04:05pm UTC`

## day-year

Day of year for a given date.

    ( DateTime -- Int )

Example:

<!-- test: day-year -->

| Input                      | Stack
|----------------------------|------------------
| `2006-03-15`               | `2006-03-15`
| `day-year`                 | `74`

## hours

Hours in duration.

    ( Duration -- Float )

Example:

<!-- test: hours -->

| Input           | Stack
|-----------------|------------------
| `10h20m30s`     | `10h20m30s`
| `hours 2 round` | `10.34`


## local-zone

Sets the local time zone.

    ( String -- )

Example:

<!-- test: local-zone -->

| Input                           | Stack
|---------------------------------|------------------
| `now time`                      | `3:04:05pm -0700 MST`
| `c 'est' local-zone`            | *local time zone is now 'EST'*
| `now time`                      | `5:04:05pm -0500 EST`
| `c 'Asia/Jakarta' local-zone`   | *local time zone is now 'Asia/Jakarta'*
| `now time`                      | `5:04:05am +0700`

## local-zone=

Gets the local time zone.

    ( -- String )

Example:

<!-- test: local-zone -->

| Input                  | Stack
|------------------------|------------------
| `local-zone=`          | `MST`

## minutes

Minutes in duration.

    ( Duration -- Float )

Example:

<!-- test: minutes -->

| Input             | Stack
|-------------------|------------------
| `10h20m30s`       | `10h20m30s`
| `minutes 2 round` | `620.5`

## now

The current date and time. If `now-set` has been called, that date and
time will be returned instead.

    ( -- DateTime )

<!-- test: now -->

| Input                  | Stack
|------------------------|------------------
| `now`                  | `Mon Jan 2 2006 3:04:05pm -0700 MST`

## now-set

Override the value returned by `now`. Useful for to mock current time while
testing.

    ( DateTime -- )

Example:

<!-- test: now-set -->

| Input                       | Stack
|-----------------------------|------------------
| `'Nov 5 1955 01:22'`        | `Nov 5 1955 01:22`
| `now-set`                   | *now set to 'Sat Nov 5 1955 1:22:00am -0700 MST'*
| `now`                       | `Sat Nov 5 1955 1:22:00am -0700 MST`

## now-restore

Cancel override of the value returned by now

## seconds

Seconds in duration.

    ( Duration -- Float )

Example:

<!-- test: minutes -->

| Input             | Stack
|-------------------|------------------
| `10h20m30s`       | `10h20m30s`
| `seconds`         | `37230`

## sub

The duration in time by subtracting `p1` from `p0`.

    ( p0:DateTime p1:DateTime -- Duration )

Aliases: `s`, `-`

<!-- test: sub-time -->

| Input                        | Stack
|------------------------------|------------------
| `'Jan 2 2006 10:00am`        | `Jan 2 2006 10:00am`
| `'Dec 31 2005 5:30pm' sub`   | `40h30m`

## time

Formats a date/time with the common time layout. Date information, if any,
is discarded.

    ( DateTime -- Time )

Example:

<!-- test: time -->

| Input                      | Stack
|----------------------------|------------------
| `'2006-01-02T15:04:05 UTC` | `2006-01-02T15:04:05 UTC`
| `time`                     | `3:04:05pm UTC`

## time-zone

Convert *time** to a given time *zone*

    ( time:DateTime zone:String -- DateTime )

Aliases: `tz`

Example:

<!-- test: time-zone -->

| Input                      | Stack
|----------------------------|------------------
| `now`                      | `Mon Jan 2 2006 3:04:05pm -0700 MST`
| `'PST' time-zone`          | `Mon Jan 2 2006 2:04:05pm -0800 PST`
| `'Asia/Jakarta' time-zone` | `Tue Jan 3 2006 5:04:05am +0700 WIB`

