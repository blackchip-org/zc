For integer operations, the calculator prefers using arbitrary-precision
functions provided by the [math/big](https://pkg.go.dev/math/big) library.
This type is simply called `BigInt` in the documentation.

Formatting characters are first removed when parsing an integer value. Those
characters are:

- Thousand separators (`','`, `'_'`, `' '`)
- Currency symbols ('`$'`, `'€'`, `'¥'`)

The following values all parse to the same number:

<!-- test: parse-formatting-int -->

| Input            | Stack
|------------------|-------------
| `c 12,345 int`   | `12345`
| `c 12_345 int`   | `12345`
| `c '12_345' int` | `12345`
| `c $12,345 int`  | `12345`
| `c 12,345$ int`  | `12345`

Certain operations may call functions that rely on using native integer types
defined by the implementing language.

- Signed integers of various bit sizes: `Int64`, `Int32`, `Int16`, and `Int8`
- Unsigned integers of various bit sizes: `Uint64`, `Unt32`, `Uint16`, and `Uint8`
- Integers of a size that is architecture-dependent: `Int` and `Uint`

For the `add`, `mul`, and `sub` operations, there is a corresponding
operation specifically for each native type (e.g., `add-i64` or
`mul-u16`). These can be useful to see if an calculation would cause an
overflow condition and to see what the result is in that case.

Use `int?` to see if a value can be parsed as an integer or use one of the
more specific types such as `i64?`

Integer division can either use Euclidean division with `div` and `div-mod`
or truncated division with `quo` and `quo-rem`.


