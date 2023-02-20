<!-- eval: use time -->
<!-- eval: 'Jan 2 2006 15:04:05 -0700' travel -->

# time

Date, time, and duration operations

<!-- index -->

| Operation               | Alias    | Description
|-------------------------|----------|------------
| [add](#add-duration)    | `add-d`  | Time after a duration
| [date](#date)           |          | Formats a date with the default layout
| [date-layout]           |          | Set the default date layout
| [date-layout=]          |          | Get the default date layout
| [date-time]             | `dt`     | Formats a datetime with the default layout
| [date-time-layout]      |          | Set the default datetime layout
| [date-time-layout=]     |          | Get the default datetime layout

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
| `2h add-d` | `Mon Jan 2 2006 5:30:00pm -0700`

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
| `date-layout=`             | `[weekday/abbr] [month/abbr] [day] [year] [hour/12]:[minute]:[second][period/alt] [offset-zone]`


## local

Set the local time zone.

    ( zone:String -- )

