# Types

<!-- eval: 'Jan 2 2006 15:04:05 -0700 MST' now= -->
<!-- eval: 'MST' local-zone -->

Each value on the calculator stack is a string of bytes. When an operation
needs to pop a value off the stack, it must first parse the value into the
desired type. Results are then formatted to a string before being pushed
back onto the stack.

Each type has a function for parsing and formatting. A value is considered to
be of a type if it can be successfully parsed by that type's parse function.
The parse function for a floating point number can parse values such as
`6`, `6.4`, `6e4` but not `'6 2/5'` which requires a conversion function.

Table of contents:

- [BigFloat](#float)
- [BigInt](#integer)
- [Bool](#bool)
- [Complex](#complex)
- [Date](#datetime)
- [DateTime](#datetime)
- [Decimal](#decimal)
- [DMS](#dms)
- [Duration](#duration)
- [Float](float)
- [Int](#integer), [Int64](#integer), [Int32](#integer), [Int16](#integer), [Int8](#integer)
- [Rational](#rational)
- [Str](#strval)
- [Time](#datetime)
- [Uint](#integer), [Uint64](#integer), [Uint32](#integer), [Uint16](#integer), [Uint8](#integer)
- [Val](#strval)

## Bool

A `Boolean` is a value that is either true or false.

An item on the stack can be parsed as a boolean if it is equal to `true`
or `false` when all characters are converted to lowercase. Operations
are defined for `true` and `false` that simply return that string.

Example:

<!-- test: types-bool -->

| Input           | Stack
|-----------------|-------------
| `true true and` | `true`
| `'FALSE' and`   | `false`
| `1 and`         | `no operation for: 'false' 1 and`

## Complex

A `Complex` value is a number that has a floating point real number, *real*,
and a floating point imaginary number, *imag* in the form of
*real*`+`*imag*`i`.

<!-- test: complex -->

| Input               | Stack
|---------------------|-------------
| `-16 sqrt`          | `0+4i`
| `2+2i add`          | `2+6i`

## DateTime

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

<!-- test: types-date -->

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

## Decimal

A `Decimal` value is a number using fixed-point math and support is
provided by the [shopspring/decimal](https://github.com/shopspring/decimal)
library. The calculator prefers working with `Decimal` values whenever an
operation can use a function in this library.

Decimal values first remove [formatting characters](#formatting-characters).
By default, values containing exponent notation are not parsed as a decimal.
To parse as a decimal, add a `d` suffix to the number. Example:

<!-- test: decimal-parse -->

| Input               | Stack
|---------------------|-------------
| `c 1e10 1e10 add`   | `2e10`
| `c 1e10d 1e10d add` | `20000000000`

## DMS

A `DMS` value is an angle that can be expressed as:

- decimal degrees
- degrees and minutes
- degrees, minutes, and seconds

Any valid decimal number, such as `12.345` can be parsed as a DMS value.
Use unit markers to designate each part of the DMS by using:

- degrees: `d`, `°`
- minutes: `m`, `'`, `′`
- seconds: `s`, `"`, `″`

Using the letter unit markers with no whitespace is the easiest to use when entering values manually. All of the following parse to the same
value:

<!-- test: dms-parsing -->

| Input                 | Stack
|-----------------------|-------------
| `c 10.5125 dec`       | `10.5125`
| `c 10.5125d dec`      | `10.5125`
| `c 10.5125° dec`      | `10.5125`
| `c 10d30.75m dec`     | `10.5125`
| `c 10d30.75' dec`     | `10.5125`
| `c [10° 30′ 45″] dec` | `10.5125`

## Duration

A `Duration` is a value with *hours*, *minutes*, and *seconds* in the form
of *hours*`h`*minutes*`m`*seconds*`s`. Zero values may be omitted. Examples:

<!-- test: types-duration -->

| Input                   | Stack
|-------------------------|-------------
| `4h15m30s 10m20s add`   | `4h 25m 50s`
| `10s add`               | `4h 26m`
| `34m add`               | `5h`

## Float

A `Float` value is a number using floating-point math. For operations supported
by `Decimal`, those are preferred over using a `Float`. To force the use of a
float, add a `f` suffix to the number.

<!-- test: types-float -->

| Input              | Stack
|--------------------|-------------
| `c 1.1 2.2 add`    | `3.3`
| `c 1.1f 2.2f add`  | `3.3000000000000003`

Float values first remove [formatting characters](#formatting-characters)
when parsing. Exponents can follow a number using an `e`, or a `×10`. The
latter form is useful in copy and paste operations (e.g., from Wikipedia).

Example:

3×102

<!-- test: float-exponent -->

| Input               | Stack
|---------------------|-------------
| `c 3e2 2e1 add`     | `320`
| `c 3×102 2×101 add` | `320`

## Integer

An `Integer` value is a number that can either be a:

- `BigInt`; or
- `Int`, `Int64`, `Int32`, `Int16`, `Int8`; or
- `Uint`, `Uint64`, `Unit32`, `Uint16`, `Uint8`

A `BigInt` is an integer of an arbitrary size and support is provided by the
[math/big](https://pkg.go.dev/math/big) library. The calculator prefers working
with `BigInt` values whenever an operation can use a function in this library.

The `Int` and `Uint` series of types are signed and unsigned integers of
a specific size and are used when an underlying implementation of an
operation needs that type.

Integer values first remove [formatting characters](#formatting-characters)
when parsing. Integers may have a radix prefix of:

- `0b`: binary number, base 2
- `0o`: octal number, base 8
- `0x`: hexadecimal number, base 16

<!-- test: radix -->

| Input               | Stack
|---------------------|-------------
| `c 0b11111111 dec`  | `255`
| `c 0o377 dec`       | `255`
| `c 0xff dec`        | `255`

## Rational

A `Rational`value is a number that has a numerator *n* and a denominator *d* in
the form of *n*`/`*d*. A whole number can prefix a rational using a ` `, `_` or
`-` character. Examples:

<!-- test: types-rational -->

| Input                   | Stack
|-------------------------|-------------
| `c 1/2 1/4 add`         | `3/4`
| `c 2_1/2 3_1/4 add`     | `5 3/4`
| `c [2 1/2] [3 1/4] add` | `5 3/4`
| `c 2-1/2 3-1/4 add`     | `5 3/4`

## Str/Val

The type `Str` is a string of bytes is the native type of values stored on the
stack. The parse function accepts any value and the formatting function uses
the string as-is.

The type `Val` is used when an operation doesn't depend on the type of the
value but is otherwise is the same as `Str`. It is used to notate operations
such as `swap` where the types of the values being swapped is irrelevant.

### Formatting Characters

Formatting characters are first removed when parsing numbers that are either
an `Integer`, `Decimal`, or `Float`. Those characters are:

- Thousand separators (`','`, `'_'`, `' '`)
- Currency symbols ('`$'`, `'€'`, `'¥'`)

The following strings all parse to the same number:

<!-- test: parse-formatting -->

| Input               | Stack
|---------------------|-------------
| `c 12,345.67 dec`   | `12345.67`
| `c 12_345.67 dec  ` | `12345.67`
| `c '12_345.67' dec` | `12345.67`
| `c $12,345.67 dec`  | `12345.67`
| `c 12,345.67$ dec`  | `12345.67`
