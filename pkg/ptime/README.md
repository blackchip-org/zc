# ptime

[![Go Reference](https://pkg.go.dev/badge/github.com/blackchip-org/ptime.svg)](https://pkg.go.dev/github.com/blackchip-org/ptime)

Parse date and times without knowing the layout ahead of time.

Experimental and a work in progress. Use at your own risk.


## Overview

This library is for parsing, on a best effort basis, input that is provided by
a user without having to be strict on the layout. The goal is to accept
commonly used layouts and try to reject those that don't make sense. Uncommon
layouts may parse, may fail, or may parse incorrectly. The expectation is that
the user is able to review the input and re-enter the date or time if it wasn't
parsed correctly. This library is not appropriate for parsing in bulk with no
supervision or when there needs to be a 100% guarantee that the parsing is
correct.

## Locales

A locale is necessary for parsing. The only locales pre-configured at the moment
are [en-US](https://github.com/blackchip-org/ptime/blob/main/locale/en.go) and
[fr-FR](https://github.com/blackchip-org/ptime/blob/main/locale/fr.go). The
[CLDR](https://cldr.unicode.org/) may be included at some point. In the mean
time, structures can be constructed manually with the needed locale
information.

Create a `ptime.P` structure with a locale:

```go
p := ptime.For(locale.EnUS)
parsed, err := p.Parse("2006-01-02")
```

or call each function with a locale:

```go
parsed, err := ptime.Parse(locale.EnUS, "2006-01-02")
```

## Parsing

The result of a parse is a `Parsed` structure of strings that contains the
fields that were identified. For example:

```go
ptime.Parse(locale.EnUS, "3:05pm")
```

returns this structure:

```json
{
  "Hour": "3",
  "Minute": "05",
  "Period": "PM",
  "TimeSep": ":"
}
```

Names are normalized to an abbreviated format. For example:

```go
ptime.Parse(locale.EnUS, "Friday, April 15 2014")
```

results this:

```json
{
  "Weekday": "Fri",
  "Year": "2014",
  "Month": "Apr",
  "Day": "15",
  "DateSep": " "
}
```

and:

```go
ptime.Parse(locale.FrFR, "vendredi, 15 avril 2014")
```

returns this:

```json
{
  "Weekday": "ven.",
  "Year": "2014",
  "Month": "avr.",
  "Day": "15",
  "DateSep": " "
}
```

Besides `Parse`, there are `ParseDate` and `ParseTime` functions that can
be used to restrict the parse as needed.

Layouts known to work can be found by reviewing the test cases here:

https://github.com/blackchip-org/ptime/blob/main/parser_test.go

The results of parsing any other layout are undefined.

## Time

A `time.Time` can be created from a `Parsed` structure given a reference
time. For example:

```go
parsed, err := p.Parse("3:04:05pm MST")
if err != nil {
    log.Panic(err)
}
t, err := ptime.Time(parsed, time.Now())
```

In this case, we don't want the `time.Time` value to be in the year 0, so the
year, month, and day is used from the reference time given as `time.Now()`.
Use `time.Time{}` if you really want year 0 but be aware that times can
be weird there.

If a parsed value contains a 2 digit year, the century will be set to the
year found in the reference time. If now is the year 2023 and the 2 digit year
is 99, the year will evaluate to 2099. To get 1999, set the reference year
to 1900 or use an explicit 4 digit year.

## Formatting

Use the `Format` function to format a `time.Time` with an alterative syntax to
the one provided in the standard library. The function takes a layout and the
time to format. The contents of the layout are copied as-is and date/time
fields denoted in square brackets are evaluated. For example, the following
date:

    2006-01-02

Can be formatted with:

    [year]-[momth]-[day]

Field names can be followed by a `/` and a formatting directive. For example,
the following date:

    2 Jan 06

Can be formatted with:

    [day] [month/abbr] [year/2]

The available fields and formats are as follows:

| Field/Format     | Example
|-------------------|-------------------------
| `weekday`         | `"Monday"`
| `weekday/abbr`    | `"Mon"`
| `weekday/wide`    | `"Monday"`
| `year`            | `"2006"`
| `year/2`          | `"06"`
| `month`           | `"1"`
| `month/2`         | `" 1"`
| `month/02`        | `"01"`
| `month/abbr`      | `"Jan"`
| `month/name`      | `"January"`
| `month/wide`      | `"January"`
| `day`             | `"2"`
| `day/2`           | `" 2"`
| `day/02`          | `"02"`
| `day/year`        | `"002"`
| `hour`            | `"15"`
| `hour/12`         | `"3"`
| `hour/24`         | `"15"`
| `minute`          | `"04"`
| `second`          | `"05"`
| `second/4`        | `"05.9999"`
| `period`          | `"AM"`
| `period/abbr`     | `"AM"`
| `period/alt`      | `"am"`
| `period/abbr-alt` | `"am"`
| `period/narrow`   | `"a"`
| `zone`            | `"MST"`
| `offset`          | `"-0700"`
| `offset/:`        | `"-07:00"`
| `offset-zone`     | `"-0700 MST"` or `"UTC"`
| `offset-zone/:`   | `"-07:00 MST"` or `"UTC"`
| `zone-offset`     | `"MST -0700"` or `"UTC"`
| `zone-offset/:`   | `"MST -07:00"` or `"UTC"`


## Installation

Install [go](https://go.dev/dl/).

Install with:

    go install github.com/blackchip-org/ptime/cmd/ptime@latest

## Command line

A command line interface is provided for quick parsing and formatting.

```
Usage: ptime [options] field ...
  -d	only parse date
  -f layout
    	format the result with layout
  -l locale
    	set locale (default "en-US")
  -t	only parse time
  -v	verbose
```

Example:

```bash
ptime -l en-US Mon Jan 2 2006 3:04:05pm MST
```

Output:

```json
{
  "Weekday": "Mon",
  "Year": "2006",
  "Month": "Jan",
  "Day": "2",
  "Hour": "3",
  "Minute": "04",
  "Second": "05",
  "Period": "PM",
  "Zone": "MST",
  "Offset": "-0700",
  "DateSep": " ",
  "TimeSep": ":"
}
```

Example:

```bash
ptime -l fr-FR lundi, 2/1/06 15:04:05,9999
```

Output:

```json
{
  "Weekday": "lun.",
  "Year": "06",
  "Month": "1",
  "Day": "2",
  "Hour": "15",
  "Minute": "04",
  "Second": "05",
  "FracSecond": "9999",
  "DateSep": "/",
  "TimeSep": ":"
}
```

Example:

```bash
ptime -f "[day] [month/abbr] [year/2]" 2006-01-02
```

Output:

```
2 Jan 06
```

## Code Example

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/blackchip-org/ptime"
	"github.com/blackchip-org/ptime/locale"
)

func main() {
	p := ptime.ForLocale(locale.EnUS)

	parsed, err := p.Parse("3:04:05pm MST")
	if err != nil {
		log.Panic(err)
	}
	t, err := p.Time(parsed, time.Now())
	if err != nil {
		log.Panic(err)
	}
	f := p.Format("[hour]:[minute]:[second] [offset]", t)
	fmt.Println(f)
}
```

Output:

```
15:04:05 -0700
```

Also found here:

https://github.com/blackchip-org/ptime/blob/main/cmd/ptime-example/main.go

## License

MIT

## Feedback

Contact me at zc@blackchip.org



