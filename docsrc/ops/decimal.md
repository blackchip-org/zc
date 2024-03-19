A `Decimal` value is a real number using fixed-point math and support is
provided by the [shopspring/decimal](https://github.com/shopspring/decimal)
library. The calculator typically prefers working with fixed-point math over
floating-point math whenever an operation can use a function in this library.

Floating-point, `Rational` numbers, and numbers in other bases can be converted
to a `Decimal` using the `dec` operation:

<!-- test: decimal-convert -->

| Input               | Stack
|---------------------|-------------
| `c 1e3 dec`         | `1000`
| `c 1/2 dec`         | `0.5`
| `c 0xff dec`        | `255`

When parsing a `Decimal` number formatting characters are first removed. Those characters are:

- Thousand separators (`','`, `'_'`, `' '`)
- Currency symbols ('`$'`, `'€'`, `'¥'`)

The following values all parse to the same number:

<!-- test: parse-formatting-decimal -->

| Input               | Stack
|---------------------|-------------
| `c 12,345.67 dec`   | `12345.67`
| `c 12_345.67 dec  ` | `12345.67`
| `c '12_345.67' dec` | `12345.67`
| `c $12,345.67 dec`  | `12345.67`
| `c 12,345.67$ dec`  | `12345.67`

By default, values containing exponent notation are not parsed as a decimal.
To parse as a decimal, add a `d` suffix to the number. Example:

<!-- test: decimal-parse -->

| Input               | Stack
|---------------------|-------------
| `c 1e10 1e10 add`   | `2e10`
| `c 1e10d 1e10d add` | `20000000000`
