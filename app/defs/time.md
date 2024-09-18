<!-- eval: 'Jan 2 2006 15:04:05 -0700 MST' now= -->
<!-- eval: 'MST' local.zone -->

Time is represented by either a:

- `DateTime`
- `Date`
- `Time`

An interval of time is represented by a `Duration`.

### `DateTime`, `Date`, `Time`

Parsing of `DateTime` values tries to be as lenient as possible to allow easy
entry by hand or from various sources via cut and paste. Parsing uses the
following rules:

- If a date and time is needed but the value only contains a time, the date
portion is set to today's date.
- Days of week such as `Monday` are parsed but ignored. The day of week
is computed from the actual date.
- If a two digit year is used it is assumed to apply to the current century.
The value of 23 is set to the year 2023 and the value of 99 is set to 2099.
Use a four digit year to use 1999.

The types of `Date` and `Time` are used when only those portions of a
`DateTime` are necessary.

All of the following formats can be parsed:

<!-- test: TypesDate -->

| Input                      | Stack
|----------------------------|-------------
| `c 2006-01 dt`             | `Sun Jan 1 2006 12:00:00am -0700 MST`
| `c 2006-01-02 dt`          | `Mon Jan 2 2006 12:00:00am -0700 MST`
| `c 2006-032 dt`            | `Wed Feb 1 2006 12:00:00am -0700 MST`
| `c 1/2 dt`                 | `Mon Jan 2 2006 12:00:00am -0700 MST`
| `c 1/2/2006 dt`            | `Mon Jan 2 2006 12:00:00am -0700 MST`
| `c 1/2/06 dt`              | `Mon Jan 2 2006 12:00:00am -0700 MST`
| `c 'Jan 2 2006' dt`        | `Mon Jan 2 2006 12:00:00am -0700 MST`
| `c 'Jan 2 06' dt`          | `Mon Jan 2 2006 12:00:00am -0700 MST`
| `c 'Fri Jan 2 2006' dt`    | `Mon Jan 2 2006 12:00:00am -0700 MST`
| `c 'Friday Jan 2 2006' dt` | `Mon Jan 2 2006 12:00:00am -0700 MST`
| `c 'Jan 2' dt`             | `Mon Jan 2 2006 12:00:00am -0700 MST`
| `c 'Fri, Jan 2' dt`        | `Mon Jan 2 2006 12:00:00am -0700 MST`
| `c 15:04:05 dt`            | `Mon Jan 2 2006 3:04:05pm -0700 MST`
| `c '15:04:05 PDT' dt`      | `Mon Jan 2 2006 3:04:05pm -0700 PDT`
| `c 15:04 dt`               | `Mon Jan 2 2006 3:04:00pm -0700 MST`
| `c 3:04PM dt`              | `Mon Jan 2 2006 3:04:00pm -0700 MST`
| `c '3:04 PM' dt`           | `Mon Jan 2 2006 3:04:00pm -0700 MST`
| `c 3:04a dt`               | `Mon Jan 2 2006 3:04:00am -0700 MST`
| `c '3:04a EST' dt`         | `Mon Jan 2 2006 3:04:00am -0500 EST`
| `c '3:04a -0500' dt`       | `Mon Jan 2 2006 3:04:00am -0500`
| `c '3:04a EST -0500' dt`   | `Mon Jan 2 2006 3:04:00am -0500 EST`

The examples above are the result if the current time is
`Jan 2 2006 15:04:05 -0700 MST`.

### Duration 

A `Duration` is a value with *hours*, *minutes*, and *seconds* in the form
of *hours*`h`*minutes*`m`*seconds*`s`. Zero values may be omitted. Examples:

<!-- test: TypesDuration -->

| Input                   | Stack
|-------------------------|-------------
| `4h15m30s 10m20s add`   | `4h 25m 50s`
| `10s add`               | `4h 26m`
| `34m add`               | `5h`

