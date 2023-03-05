# format

<!-- eval: use format -->

Numeric formatting

    use format -- user and dev prelude

<!-- index -->

| Operation                         | Description
|-----------------------------------|----------------
| [auto-currency](#auto-currency)   | Retain currency symbols when true, set
| [auto-currency=](#auto-currency=) | Retain currency symbols when true, get
| [auto-format](#auto-format)       | Format numbers by default when true, set
| [auto-format=](#auto-format=)     | Format numbers by default when true, get
| [format](#format)                 | Apply a numeric layout
| [int-layout](#int-layout)         | Layout for integers, get
| [int-layout=](#int-layout=)       | Layout for integers, set
| [min-digits](#min-digits)         | Minimum number of digits after the decimal point to display, set
| [min-digits](#min-digits)         | Minimum number of digits after the decimal point to display, get
| [precision](#precision)           | Precision for fixed-point math, set
| [precision=](#precision=)         | Prevision for fixed-point math, get
| [point](#point)                   | Character to use as the decimal point, set
| [point=](#point=)                 | Character to use as the decimal point, get
| [round](#round)                   | Round to a given precision
| [rounding-mode](#rounding-mode)   | Method to use in rounding, set
| [rounding-mode=](#rounding-mode=) | Method to use in rounding, get


## auto-currency

When true, currency symbols that are a prefix or a suffix to a number are
retained when pushing the value to the stack.

    ( b:Bool -- )

Example:

<!-- test: auto-currency -->

| Input                | Stack
|----------------------|---------------------
| `$1234`              | `1234`
| `clear`              |
| `true auto-currency` | *auto-currency set to true*
| `$1234`              | `$1234`


## auto-currency=

Gets the current auto-currency setting

    ( -- b:Bool )

Example:

<!-- test: auto-currency -->

| Input                | Stack
|----------------------|---------------------
| `auto-currency=`     | `false`


## auto-format

When true, the `int-layout` is automatically applied to numbers when they
are pushed onto the stack.

    ( b:Bool -- )

Example:

<!-- test: auto-format -->

| Input                | Stack
|----------------------|---------------------
| `$1234`              | `1234`
| `clear`              |
| `true auto-format`   | *auto-format set to true*
| `$1234`              | `1,234`


## auto-format=

Gets the current auto-format setting

    ( -- b:Bool )

Example:

<!-- test: auto-format -->

| Input                | Stack
|----------------------|---------------------
| `auto-format=`       | `false`


## format

Format a number with the current layout.

    ( n:Num -- format:Num )


Example:

<!-- test: format -->

| Input                | Stack
|----------------------|---------------------
| `1234`               | `1234`
| `format`             | `1,234`


## int-layout

Layout to use for integers and for integer parts of decimal numbers. This
layout is subject to change.

    ( layout:Str -- )

Example:

<!-- test: int-layout -->

| Input                | Stack
|----------------------|---------------------
| `' 000' int-layout`  | *int-layout set to ' 000'*
| `12345 format`       | `12 345`


## int-layout=

Gets the current integer layout.

    ( -- layout:Str )

Example:

<!-- test: int-layout -->

| Input                | Stack
|----------------------|---------------------
| `int-layout=`        | `,000`


## min-digits

The minimum number of digits to display after the decimal point.

    ( digits:Int -- )

Example:

<!-- test: min-digits -->

| Input                | Stack
|----------------------|---------------------
| `1234`               | `1234`
| `2 min-digits`       | *min-digits set to 2*
| `0 add`              | `1234.00`


## min-digits=

Gets the current minimum digits setting.

    ( -- digits:Int )

Example:

<!-- test: min-digits -->

| Input                | Stack
|----------------------|---------------------
| `min-digits=`        | `0`


## precision

Sets the precision to `a` for fixed-point math which is the number of digits
used after the decimal point. Extra digits are rounded using the current
rounding mode. If set to zero, no rounding is performed.

    ( a:Int32 -- )

Example:

<!-- test: places -->

| Input               | Stack
|---------------------|---------------------|
| `2 3 div`           | `0.6666666666666667`
| `clear`             |
| `2 precision`       | *precision set to 2*
| `2 3 div`           | `0.67`


## precision=

Gets the precision for fixed-point math.

    ( -- places:Int32 )

Example:

<!-- test: places= -->

| Input               | Stack
|---------------------|---------------------|
| `precision=`        | `0`


## point

Sets the character to use as the decimal point.

    ( point:Char -- )

Example:

<!-- test: point -->

| Input               | Stack
|---------------------|---------------------|
| `1234 10 div`       | `123.4`
| `',' point`         | *point set to ','*
| `clear`             |
| `1234 10 div`       | `123,4`


## point=

Gets the character used as the decimal point.

    ( -- point:Char )

Example:

<!-- test: point= -->

| Input               | Stack
|---------------------|---------------------|
| `point=`            | `.`


## round

Rounds the number `n` to `d` digits using the current rounding mode.

    ( n:Num d:Int -- round:Num )

Example:

<!-- test: round -->

| Input               | Stack
|---------------------|---------------------|
| `2 3 div`           | `0.6666666666666667`
| `2 round`           | `0.67`


## rounding-mode

Sets the mode to be used when rounding. Valid modes are:

- `half-up`
- `ceil`
- `down`
- `floor`
- `half-even`
- `up`

    ( m:Str -- )

Example:

<!-- test: rounding-mode -->

| Input                | Stack
|----------------------|---------------------
| `1.01 0.05 mul`      | `0.0505`
| `2 round`            | `0.05`
| `'up' rounding-mode` | *rounding-mode set to 'up'*
| `clear`              |
| `1.01 0.05 mul`      | `0.0505`
| `2 round`            | `0.06`


## rounding-mode=

Gets the current rounding mode

    ( -- m:Str )

Example:

<!-- test rounding-mode= -->

| Input                | Stack
|----------------------|---------------------
| `rounding-mode=`     | `half-up`
